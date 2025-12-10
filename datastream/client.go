package datastream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/schematichq/rulesengine"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
)

func NewDataStreamClient(options DataStreamClientOptions, configurationOptions *core.DatastreamOptions) *DataStreamClient {
	// Provide default configuration options if nil is passed
	if configurationOptions == nil {
		configurationOptions = &core.DatastreamOptions{
			CacheTTL: defaultTTL,
		}
	}

	// Get or create cache providers based on options
	companyCacheProvider, userCacheProvider := getCacheProviders(options, configurationOptions)

	if companyCacheProvider == nil || userCacheProvider == nil {
		panic("failed to create cache providers")
	}

	// Create flag cache based on configuration - use Redis if configured, otherwise local cache
	flagCacheProvider := options.FlagCache
	if flagCacheProvider == nil {
		flagCacheProvider = buildFlagCacheProvider(configurationOptions)
	}

	client := &DataStreamClient{
		apiKey:               options.ApiKey,
		cacheTTL:             configurationOptions.CacheTTL,
		logger:               options.Logger,
		flagsCacheProvider:   flagCacheProvider,
		companyCacheProvider: companyCacheProvider,
		companyCache:         make(map[string]*rulesengine.Company),
		userCacheProvider:    userCacheProvider,

		pendingCompanyRequests: make(map[string][]chan *rulesengine.Company),
		pendingUserRequests:    make(map[string][]chan *rulesengine.User),

		// Replicator mode configuration
		replicatorMode:        configurationOptions.ReplicatorMode,
		replicatorHealthURL:   configurationOptions.ReplicatorHealthURL,
		replicatorHealthCheck: configurationOptions.ReplicatorHealthCheck,
		replicatorReady:       false,
		replicatorHealthDone:  make(chan bool, 1),
	}

	// Initialize WebSocket client if not in replicator mode and if BaseURL is provided
	if !client.replicatorMode && options.BaseURL != "" {
		wsClient, err := schematicdatastreamws.NewClient(schematicdatastreamws.ClientOptions{
			URL:                    options.BaseURL,
			ApiKey:                 options.ApiKey,
			MessageHandler:         client.handleWebSocketMessage,
			ConnectionReadyHandler: client.handleConnectionReady,
			Logger:                 client.logger,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to create WebSocket client: %v", err))
		}
		client.wsClient = wsClient
	}

	// Start replicator health checking if enabled
	if client.replicatorMode {
		go client.startReplicatorHealthCheck()
	}

	return client
}

// handleWebSocketMessage processes incoming WebSocket messages
func (c *DataStreamClient) handleWebSocketMessage(ctx context.Context, message *schematicdatastreamws.DataStreamResp) error {
	return c.handleMessageResponse(ctx, message)
}

// handleConnectionReady is called when the WebSocket connection is ready
func (c *DataStreamClient) handleConnectionReady(ctx context.Context) error {
	return c.getAllFlags(ctx)
}

func (c *DataStreamClient) Start() {
	// In replicator mode, we don't establish WebSocket connections
	// The external replicator handles all data streaming
	if c.replicatorMode {
		c.logger.Info(context.Background(), "Replicator mode enabled - skipping WebSocket connection")
		return
	}

	// Start the WebSocket client
	if c.wsClient != nil {
		c.wsClient.Start()
	}
}

func (c *DataStreamClient) sendWebSocketMessage(ctx context.Context, req *schematicdatastreamws.DataStreamReq) error {
	if c.wsClient == nil {
		return fmt.Errorf("WebSocket client is not initialized")
	}

	message := c.packageMessage(req)
	return c.wsClient.SendMessage(message)
}

func (c *DataStreamClient) Close() {
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, fmt.Sprintf("Fatal error occurred while closing client: %v", r))
		}
	}()

	// Stop replicator health checking if enabled
	if c.replicatorMode {
		c.logger.Info(ctx, "Stopping replicator health check")
		close(c.replicatorHealthDone)
		return
	}

	// Close WebSocket client if it exists
	if c.wsClient != nil {
		c.logger.Info(ctx, "Closing WebSocket connection")
		c.wsClient.Close()
	}

	c.logger.Info(ctx, "WebSocket connection closed")
}

// IsConnected checks if the WebSocket connection is active
// In replicator mode, returns true if the external replicator is ready
func (c *DataStreamClient) IsConnected() bool {
	if c.replicatorMode {
		return c.IsReplicatorReady()
	}

	if c.wsClient == nil {
		return false
	}

	return c.wsClient.IsConnected()
}

func (c *DataStreamClient) handleMessageResponse(ctx context.Context, message *schematicdatastreamws.DataStreamResp) error {
	if message == nil {
		return nil
	}

	if message.MessageType == schematicdatastreamws.MessageTypeError {
		return c.handleErrorMessage(ctx, message)
	}

	switch message.EntityType {
	case string(schematicdatastreamws.EntityTypeCompany):
		return c.handleCompanyMessage(ctx, message)
	case string(schematicdatastreamws.EntityTypeFlags):
		return c.handleFlagsMessage(ctx, message)
	case string(schematicdatastreamws.EntityTypeFlag):
		return c.handleFlagMessage(ctx, message)
	case string(schematicdatastreamws.EntityTypeUser):
		return c.handleUserMessage(ctx, message)
	default:
		return fmt.Errorf("Received unknown entity type: %s", message.EntityType)
	}
}

func (c *DataStreamClient) handleFlagsMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	flags := resp.Data

	var flagsData []*rulesengine.Flag
	err := json.Unmarshal(flags, &flagsData)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal flags data: %v", err)
	}

	c.flagsMu.Lock()
	var cacheKeys []string
	for _, flag := range flagsData {
		cacheKey := flagCacheKey(flag.Key)
		if err := c.flagsCacheProvider.Set(ctx, cacheKey, flag, nil); err != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Failed to cache flag '%s': %v", flag.Key, err))
		}
		cacheKeys = append(cacheKeys, cacheKey)
	}

	c.flagsCacheProvider.DeleteMissing(ctx, cacheKeys)
	c.flagsMu.Unlock()

	if c.pendingFlagRequest != nil {
		c.pendingFlagRequest <- true
	}

	return nil
}

func (c *DataStreamClient) handleFlagMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	flags := resp.Data

	var flag *rulesengine.Flag
	if err := json.Unmarshal(flags, &flag); err != nil {
		return fmt.Errorf("Failed to unmarshal flags data: %v", err)
	}

	c.flagsMu.Lock()
	defer func() {
		c.flagsMu.Unlock()

		if c.pendingFlagRequest != nil {
			c.pendingFlagRequest <- true
		}
	}()

	cacheKey := flagCacheKey(flag.Key)
	switch resp.MessageType {
	case schematicdatastreamws.MessageTypeDelete:
		if err := c.flagsCacheProvider.Delete(ctx, cacheKey); err != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Failed to delete flag from cache '%s': %v", cacheKey, err))
		}
	case schematicdatastreamws.MessageTypeFull:
		// For full messages, we replace the existing flag in the cache
		if err := c.flagsCacheProvider.Set(ctx, cacheKey, flag, nil); err != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Failed to cache flag '%s': %v", flag.Key, err))
		}
	default:
		// Other message types are unhandled
		c.logger.Warn(ctx, fmt.Sprintf("Unhandled message type for flag: %s", resp.MessageType))
	}

	return nil
}

// Helper function to notify pending company requests
func (c *DataStreamClient) notifyPendingCompanyRequests(_ context.Context, keys map[string]string, company *rulesengine.Company) {
	// Notify all goroutines waiting for this company
	c.pendingCompReqMu.Lock()
	defer c.pendingCompReqMu.Unlock()

	// For each relevant key, notify all waiting channels
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)

		// Check if there are channels waiting for this key
		if channels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			for _, ch := range channels {
				// Non-blocking send - if the channel is full, we'll just continue
				select {
				case ch <- company:
				default:
				}
			}

			// Remove this set of channels as they've been notified
			delete(c.pendingCompanyRequests, cacheKey)
		}
	}
}

func (c *DataStreamClient) handleCompanyMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	var company *rulesengine.Company
	err := json.Unmarshal(resp.Data, &company)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal company data: %v", err)
	}

	if company == nil {
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypeDelete {
		// Remove the company from the cache
		for key, value := range company.Keys {
			companyKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
			if err := c.companyCacheProvider.Delete(ctx, companyKey); err != nil {
				c.logger.Warn(ctx, fmt.Sprintf("Failed to delete company from cache '%s': %v", companyKey, err))
			}
		}
		return nil
	}

	c.companyMu.Lock()
	cacheResults, err := c.cacheCompanyForKeys(ctx, company)
	c.companyMu.Unlock()

	if err != nil {
		return err
	}

	// Check for cache errors and log warnings
	for cacheKey, cacheErr := range cacheResults {
		if cacheErr != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Cache error for company key '%s': %v", cacheKey, cacheErr))
		}
	}

	// Notify pending requests with the company data
	c.notifyPendingCompanyRequests(ctx, company.Keys, company)

	return nil
}

// Helper function to notify pending user requests
func (c *DataStreamClient) notifyPendingUserRequests(_ context.Context, keys map[string]string, user *rulesengine.User) {
	// Notify all goroutines waiting for this user
	c.pendingUserReqMu.Lock()
	defer c.pendingUserReqMu.Unlock()

	// For each relevant key, notify all waiting channels
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		if channels, ok := c.pendingUserRequests[cacheKey]; ok {
			for _, ch := range channels {
				// Non-blocking send - if the channel is full, we'll just continue
				select {
				case ch <- user:
				default:
				}
			}
			// Remove this set of channels as they've been notified
			delete(c.pendingUserRequests, cacheKey)
		}
	}
}

func (c *DataStreamClient) handleUserMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	var user *rulesengine.User
	err := json.Unmarshal(resp.Data, &user)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal user data: %v", err)
	}

	if user == nil {
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypeDelete {
		// Remove the user from the cache
		for key, value := range user.Keys {
			userKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
			if err := c.userCacheProvider.Delete(ctx, userKey); err != nil {
				c.logger.Warn(ctx, fmt.Sprintf("Failed to delete user from cache '%s': %v", userKey, err))
			}
		}

		return nil
	}

	// Map to track which cache keys were successfully cached
	cacheResults := make(map[string]error)

	for key, value := range user.Keys {
		userKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		ttl := c.cacheTTL
		err := c.userCacheProvider.Set(ctx, userKey, user, &ttl)

		// Store the result for each cache key
		cacheResults[userKey] = err
	}

	// Check for cache errors and log warnings
	for cacheKey, cacheErr := range cacheResults {
		if cacheErr != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Cache error for user key '%s': %v", cacheKey, cacheErr))
		}
	}

	// Notify pending requests with the user data
	c.notifyPendingUserRequests(ctx, user.Keys, user)

	return nil
}

func (c *DataStreamClient) handleErrorMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	var respError schematicdatastreamws.DataStreamError
	err := json.Unmarshal(resp.Data, &respError)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal error message: %v", err)
	}

	// Check if we have both keys and entity type in the error response
	if len(respError.Keys) > 0 && respError.EntityType != nil {
		// Based on the entity type, notify the appropriate pending channels with nil
		switch *respError.EntityType {
		case schematicdatastreamws.EntityTypeCompany:
			c.notifyPendingCompanyRequests(ctx, respError.Keys, nil)
		case schematicdatastreamws.EntityTypeUser:
			c.notifyPendingUserRequests(ctx, respError.Keys, nil)
		default:
			// Unsupported entity type
			c.logger.Warn(ctx, fmt.Sprintf("Received error for unsupported entity type: %s", *respError.EntityType))
		}
	}

	return fmt.Errorf("%s", respError.Error)
}

func (c *DataStreamClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) (*rulesengine.CheckFlagResult, error) {
	// Get flag first - return error if not found
	flag, found := c.getFlag(ctx, flagKey)
	if !found {
		return nil, fmt.Errorf("flag not found: %s", flagKey)
	}

	needsCompany := len(evalCtx.Company) > 0
	needsUser := len(evalCtx.User) > 0

	var company *rulesengine.Company
	var user *rulesengine.User
	var err error

	// Try to get cached data first
	if needsCompany {
		company = c.getCompanyFromCache(evalCtx.Company)
	}
	if needsUser {
		user = c.getUserFromCache(evalCtx.User)
	}

	// Check if we have all the data we need
	hasAllData := (!needsCompany || company != nil) && (!needsUser || user != nil)

	// If we don't have all cached data, try to fetch missing data
	if !hasAllData {
		// In replicator mode, we don't fetch data - we expect it to be cached
		if c.replicatorMode {
			if !c.IsReplicatorReady() {
				return nil, fmt.Errorf("replicator not ready and missing cached data")
			}
			// Continue with whatever cached data we have - replicator should have populated it
		} else {
			// Check if we're connected to datastream for fetching
			if !c.IsConnected() {
				return nil, fmt.Errorf("datastream not connected and missing cached data")
			} // Fetch missing company data
			if needsCompany && company == nil {
				company, err = c.getCompany(ctx, evalCtx.Company)
				if err != nil {
					return nil, fmt.Errorf("failed to get company data: %w", err)
				}
			}

			// Fetch missing user data
			if needsUser && user == nil {
				user, err = c.getUser(ctx, evalCtx.User)
				if err != nil {
					return nil, fmt.Errorf("failed to get user data: %w", err)
				}
			}
		}
	}

	// Evaluate against the rules engine
	resp, err := rulesengine.CheckFlag(ctx, company, user, flag)
	if err != nil {
		return nil, fmt.Errorf("rules engine error: %w", err)
	}

	return resp, nil
}

func (c *DataStreamClient) getAllFlags(ctx context.Context) error {
	// Check if there is already a pending request for flags
	if c.pendingFlagRequest != nil {
		return nil
	}

	waitCh := make(chan bool, 1)
	c.pendingFlagReqMu.Lock()
	c.pendingFlagRequest = waitCh
	c.pendingFlagReqMu.Unlock()
	defer func() {
		c.pendingFlagRequest = nil
		close(waitCh)
	}()

	req := &schematicdatastreamws.DataStreamReq{
		EntityType: schematicdatastreamws.EntityTypeFlags,
	}
	err := c.sendWebSocketMessage(ctx, req)
	if err != nil {
		c.logger.Error(ctx, fmt.Sprintf("Failed to send WebSocket message: %v", err))
		return err
	}

	select {
	case <-waitCh:
	case <-time.After(resourceTimeout):
		return fmt.Errorf("timeout while waiting for flags data")
	case <-ctx.Done():
		c.pendingFlagRequest = nil
		return ctx.Err()
	}
	return nil
}

func (c *DataStreamClient) getCompany(ctx context.Context, keys map[string]string) (*rulesengine.Company, error) {
	company := c.getCompanyFromCache(keys)
	if company != nil {
		return company, nil
	}

	waitCh := make(chan *rulesengine.Company, 1) // Register the wait channel for each key
	cacheKeys := []string{}
	c.companyMu.Lock()
	shouldSendRequest := true
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		cacheKeys = append(cacheKeys, cacheKey)
		if existingChannels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			c.pendingCompanyRequests[cacheKey] = append(existingChannels, waitCh)
			shouldSendRequest = false // Someone else already requested this
		}
	}
	c.companyMu.Unlock()

	if shouldSendRequest {
		req := &schematicdatastreamws.DataStreamReq{
			EntityType: schematicdatastreamws.EntityTypeCompany,
			Keys:       keys,
		}

		err := c.sendWebSocketMessage(ctx, req)
		if err != nil {
			return nil, err
		}

		c.pendingCompReqMu.Lock()
		for key, value := range keys {
			cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
			c.pendingCompanyRequests[cacheKey] = []chan *rulesengine.Company{waitCh}
		}
		c.pendingCompReqMu.Unlock()
	}

	var err error
	select {
	case company := <-waitCh:
		return company, nil
	case <-time.After(resourceTimeout):
		err = fmt.Errorf("timeout while waiting for company data")
	case <-ctx.Done():
		err = ctx.Err()
	}

	// Clean up all channels involved in this request
	c.cleanupPendingCompanyRequests(cacheKeys, waitCh)

	return nil, err
}

func (c *DataStreamClient) getUser(ctx context.Context, keys map[string]string) (*rulesengine.User, error) {
	user := c.getUserFromCache(keys)
	if user != nil {
		return user, nil
	}

	waitCh := make(chan *rulesengine.User, 1) // Register the wait channel for each key
	cacheKeys := []string{}
	c.pendingUserReqMu.Lock()
	shouldSendRequest := true
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value) // If there are already pending requests for this key, just add our channel
		cacheKeys = append(cacheKeys, cacheKey)
		if existingChannels, ok := c.pendingUserRequests[cacheKey]; ok {
			c.pendingUserRequests[cacheKey] = append(existingChannels, waitCh)
			shouldSendRequest = false // Someone else already requested this
		}
	}
	c.pendingUserReqMu.Unlock()

	if shouldSendRequest {
		req := &schematicdatastreamws.DataStreamReq{
			EntityType: schematicdatastreamws.EntityTypeUser,
			Keys:       keys,
		}

		err := c.sendWebSocketMessage(ctx, req)
		if err != nil {
			return nil, err
		}

		c.pendingUserReqMu.Lock()
		for key, value := range keys {
			cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
			c.pendingUserRequests[cacheKey] = []chan *rulesengine.User{waitCh}
		}
		c.pendingUserReqMu.Unlock()
	}

	var err error
	select {
	case user := <-waitCh:
		return user, nil
	case <-time.After(resourceTimeout):
		err = fmt.Errorf("timeout while waiting for user data")
	case <-ctx.Done():
		err = ctx.Err()
	}

	// Clean up all channels involved in this request
	c.cleanupPendingUserRequests(cacheKeys, waitCh)

	return nil, err
}

func (c *DataStreamClient) getFlag(ctx context.Context, key string) (*rulesengine.Flag, bool) {
	c.flagsMu.RLock()
	defer c.flagsMu.RUnlock()
	flag, exists := c.flagsCacheProvider.Get(ctx, flagCacheKey(key))
	if !exists {
		return nil, false
	}
	return flag, true
}

func (c *DataStreamClient) packageMessage(req *schematicdatastreamws.DataStreamReq) *schematicdatastreamws.DataStreamBaseReq {
	return &schematicdatastreamws.DataStreamBaseReq{
		Data: *req,
	}
}

func (c *DataStreamClient) getCompanyFromCache(keys map[string]string) *rulesengine.Company {
	c.companyMu.RLock()
	defer c.companyMu.RUnlock()
	for key, value := range keys {
		companyKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		company, exists := c.companyCacheProvider.Get(context.Background(), companyKey)
		if exists {
			return company
		}
	}

	return nil
}

func (c *DataStreamClient) getUserFromCache(keys map[string]string) *rulesengine.User {
	c.userMu.RLock()
	defer c.userMu.RUnlock()
	for key, value := range keys {
		userKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		user, exists := c.userCacheProvider.Get(context.Background(), userKey)
		if exists {
			return user
		}
	}

	return nil
}

func flagCacheKey(key string) string {
	return fmt.Sprintf("%s:%s:%s:%s", cacheKeyPrefix, cacheKeyPrefixFlags, rulesengine.VersionKey, strings.ToLower(key))
}

func resourceKeyToCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", cacheKeyPrefix, resourceType, rulesengine.VersionKey, strings.ToLower(key), strings.ToLower(value))
}

// Helper function to clean up pending company requests
func (c *DataStreamClient) cleanupPendingCompanyRequests(cacheKeys []string, waitCh chan *rulesengine.Company) {
	c.pendingCompReqMu.Lock()
	defer c.pendingCompReqMu.Unlock()
	for _, cacheKey := range cacheKeys {
		if channels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			// Remove the specific channel from the list
			filteredChannels := []chan *rulesengine.Company{}
			for _, ch := range channels {
				if ch != waitCh {
					filteredChannels = append(filteredChannels, ch)
				}
			}
			if len(filteredChannels) > 0 {
				c.pendingCompanyRequests[cacheKey] = filteredChannels
			} else {
				delete(c.pendingCompanyRequests, cacheKey)
			}
		}
	}
	close(waitCh) // Close the wait channel to signal that we're done with it
}

func (c *DataStreamClient) cleanupPendingUserRequests(cacheKeys []string, waitCh chan *rulesengine.User) {
	c.pendingUserReqMu.Lock()
	defer c.pendingUserReqMu.Unlock()
	for _, cacheKey := range cacheKeys {
		if channels, ok := c.pendingUserRequests[cacheKey]; ok {
			// Remove the specific channel from the list
			filteredChannels := []chan *rulesengine.User{}
			for _, ch := range channels {
				if ch != waitCh {
					filteredChannels = append(filteredChannels, ch)
				}
			}
			if len(filteredChannels) > 0 {
				c.pendingUserRequests[cacheKey] = filteredChannels
			} else {
				delete(c.pendingUserRequests, cacheKey)
			}
		}
	}

	close(waitCh) // Close the wait channel to signal that we're done with it
}

func (c *DataStreamClient) cacheCompanyForKeys(ctx context.Context, company *rulesengine.Company) (map[string]error, error) {
	if len(company.Keys) == 0 {
		return nil, errors.New("no keys provided for company lookup")
	}

	if company == nil {
		return nil, errors.New("company cannot be nil")
	}

	// Map to track which cache keys were successfully cached and which ones failed
	cacheResults := make(map[string]error)

	// Try to cache the company for all keys
	for key, value := range company.Keys {
		companyKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		ttl := c.cacheTTL
		err := c.companyCacheProvider.Set(ctx, companyKey, company, &ttl)

		// Store the result for each cache key
		cacheResults[companyKey] = err
	}

	return cacheResults, nil
}

func (c *DataStreamClient) UpdateCompanyMetrics(ctx context.Context, event *schematicgo.EventBodyTrack) error {
	if event == nil {
		return errors.New("event body cannot be nil")
	}

	keys := event.Company
	if len(keys) == 0 {
		return errors.New("no keys provided for company lookup")
	}

	company := c.getCompanyFromCache(keys)
	if company == nil {
		return nil
	}

	// Create a deep copy of the company to avoid modifying the cached object directly
	companyCopy := deepCopyCompany(company)

	// Update the metric value if it matches the event
	for _, metric := range companyCopy.Metrics {
		if metric == nil {
			continue
		}
		if metric.EventSubtype == event.Event {
			metric.Value += int64(*event.Quantity)
		}
	}

	c.companyMu.Lock()
	cacheResults, err := c.cacheCompanyForKeys(ctx, companyCopy)
	c.companyMu.Unlock()

	if err != nil {
		return fmt.Errorf("failed to cache company metrics: %v", err)
	}

	// Log specific cache key failures
	for cacheKey, cacheErr := range cacheResults {
		if cacheErr != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Failed to cache company metric for key '%s': %v", cacheKey, cacheErr))
		}
	}
	return nil
}

// deepCopyCompany creates a complete deep copy of a Company struct and all its nested fields.
// This ensures that modifying the returned company won't affect the original object.
// Use this function when you need to make changes to a company object without affecting the cached version.
func deepCopyCompany(company *rulesengine.Company) *rulesengine.Company {
	if company == nil {
		return nil
	}

	// Create a deep copy of the company
	companyCopy := &rulesengine.Company{
		ID:                company.ID,
		AccountID:         company.AccountID,
		EnvironmentID:     company.EnvironmentID,
		BasePlanID:        company.BasePlanID,
		BillingProductIDs: append([]string{}, company.BillingProductIDs...),
		CRMProductIDs:     append([]string{}, company.CRMProductIDs...),
		PlanIDs:           append([]string{}, company.PlanIDs...),
		Subscription:      company.Subscription, // Note: this is a shallow copy, as Subscription isn't modified in our case
		Keys:              make(map[string]string),
		Metrics:           make([]*rulesengine.CompanyMetric, 0, len(company.Metrics)),
		Traits:            make([]*rulesengine.Trait, 0, len(company.Traits)),
	}

	// Copy the keys map
	for k, v := range company.Keys {
		companyCopy.Keys[k] = v
	}

	// Copy the metrics slice and each metric inside it
	for _, metric := range company.Metrics {
		if metric == nil {
			companyCopy.Metrics = append(companyCopy.Metrics, nil)
			continue
		}

		metricCopy := &rulesengine.CompanyMetric{
			AccountID:     metric.AccountID,
			EnvironmentID: metric.EnvironmentID,
			CompanyID:     metric.CompanyID,
			EventSubtype:  metric.EventSubtype,
			Period:        metric.Period,
			MonthReset:    metric.MonthReset,
			Value:         metric.Value,
			CreatedAt:     metric.CreatedAt,
		}

		if metric.ValidUntil != nil {
			validUntil := *metric.ValidUntil
			metricCopy.ValidUntil = &validUntil
		}

		companyCopy.Metrics = append(companyCopy.Metrics, metricCopy)
	}

	// Copy traits if needed
	for _, trait := range company.Traits {
		if trait != nil {
			traitCopy := *trait // shallow copy is sufficient for traits since we don't modify them deeply
			companyCopy.Traits = append(companyCopy.Traits, &traitCopy)
		} else {
			companyCopy.Traits = append(companyCopy.Traits, nil)
		}
	}

	return companyCopy
}

// startReplicatorHealthCheck starts a background goroutine that periodically checks
// the external replicator's readiness endpoint
func (c *DataStreamClient) startReplicatorHealthCheck() {
	ctx := context.Background()
	ticker := time.NewTicker(c.replicatorHealthCheck)
	defer ticker.Stop()

	c.logger.Info(ctx, fmt.Sprintf("Starting replicator health check with URL: %s, interval: %v", c.replicatorHealthURL, c.replicatorHealthCheck))

	// Initial health check
	c.checkReplicatorHealth(ctx)

	for {
		select {
		case <-c.replicatorHealthDone:
			c.logger.Info(ctx, "Replicator health check stopped")
			return
		case <-ticker.C:
			c.checkReplicatorHealth(ctx)
		}
	}
}

// checkReplicatorHealth performs a single health check against the external replicator
func (c *DataStreamClient) checkReplicatorHealth(ctx context.Context) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(c.replicatorHealthURL)
	if err != nil {
		c.setReplicatorReady(false)
		c.logger.Debug(ctx, fmt.Sprintf("Replicator health check failed: %v", err))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.setReplicatorReady(false)
		c.logger.Warn(ctx, fmt.Sprintf("Failed to read replicator health response: %v", err))
		return
	}

	// Parse the JSON response
	var healthResp struct {
		Ready bool `json:"ready"`
	}

	if err := json.Unmarshal(body, &healthResp); err != nil {
		c.setReplicatorReady(false)
		c.logger.Warn(ctx, fmt.Sprintf("Failed to parse replicator health response: %v", err))
		return
	}

	// Update replicator ready state
	wasReady := c.IsReplicatorReady()
	c.setReplicatorReady(healthResp.Ready)

	// Log state changes
	if healthResp.Ready && !wasReady {
		c.logger.Info(ctx, "External replicator is now ready")
	} else if !healthResp.Ready && wasReady {
		c.logger.Info(ctx, "External replicator is no longer ready")
	}
}

// IsReplicatorReady returns whether the external replicator is ready
func (c *DataStreamClient) IsReplicatorReady() bool {
	c.replicatorMu.RLock()
	defer c.replicatorMu.RUnlock()
	return c.replicatorReady
}

// setReplicatorReady updates the replicator ready state
func (c *DataStreamClient) setReplicatorReady(ready bool) {
	c.replicatorMu.Lock()
	defer c.replicatorMu.Unlock()
	c.replicatorReady = ready
}

// IsReplicatorMode returns whether the client is running in replicator mode
func (c *DataStreamClient) IsReplicatorMode() bool {
	return c.replicatorMode
}
