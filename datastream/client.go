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

	// Build a shared Redis client (if configured) for all cache providers
	redisClient := buildRedisClient(configurationOptions)

	// Get or create cache providers based on options
	companyCacheProvider, userCacheProvider := getCacheProviders(options, configurationOptions, redisClient)

	if companyCacheProvider == nil || userCacheProvider == nil {
		panic("failed to create cache providers")
	}

	// Create flag cache based on configuration - use Redis if configured, otherwise local cache
	flagCacheProvider := options.FlagCache
	if flagCacheProvider == nil {
		flagCacheProvider = buildFlagCacheProvider(configurationOptions, redisClient)
	}

	// Build lookup cache providers for two-step ID resolution
	companyLookupCacheProvider := buildLookupCacheProvider(configurationOptions, redisClient)
	userLookupCacheProvider := buildLookupCacheProvider(configurationOptions, redisClient)

	client := &DataStreamClient{
		apiKey:             options.ApiKey,
		cacheTTL:           configurationOptions.CacheTTL,
		logger:             options.Logger,
		flagsCacheProvider: flagCacheProvider,

		pendingCompanyRequests: make(map[string][]chan *rulesengine.Company),
		pendingUserRequests:    make(map[string][]chan *rulesengine.User),

		// Replicator mode configuration
		replicatorMode:        configurationOptions.ReplicatorMode,
		replicatorHealthURL:   configurationOptions.ReplicatorHealthURL,
		replicatorHealthCheck: configurationOptions.ReplicatorHealthCheck,
		replicatorReady:       false,
		replicatorHealthDone:  make(chan bool, 1),
	}

	// Initialize resourceCache instances (cacheVersion requires client to exist)
	client.companyCache = &resourceCache[*rulesengine.Company]{
		primaryCache: companyCacheProvider,
		lookupCache:  companyLookupCacheProvider,
		keyPrefix:    cacheKeyPrefixCompany,
		getID:        func(c *rulesengine.Company) string { return c.ID },
		getKeys:      func(c *rulesengine.Company) map[string]string { return c.Keys },
		cacheVersion: client.cacheVersion,
	}
	client.userCache = &resourceCache[*rulesengine.User]{
		primaryCache: userCacheProvider,
		lookupCache:  userLookupCacheProvider,
		keyPrefix:    cacheKeyPrefixUser,
		getID:        func(u *rulesengine.User) string { return u.ID },
		getKeys:      func(u *rulesengine.User) map[string]string { return u.Keys },
		cacheVersion: client.cacheVersion,
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
		return fmt.Errorf("webSocket client is not initialized")
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
		return fmt.Errorf("received unknown entity type: %s", message.EntityType)
	}
}

func (c *DataStreamClient) handleFlagsMessage(ctx context.Context, resp *schematicdatastreamws.DataStreamResp) error {
	flags := resp.Data

	var flagsData []*rulesengine.Flag
	err := json.Unmarshal(flags, &flagsData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal flags data: %v", err)
	}

	c.flagsMu.Lock()
	var cacheKeys []string
	for _, flag := range flagsData {
		cacheKey := c.flagCacheKey(flag.Key)
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
		return fmt.Errorf("failed to unmarshal flags data: %v", err)
	}

	c.flagsMu.Lock()
	defer func() {
		c.flagsMu.Unlock()

		if c.pendingFlagRequest != nil {
			c.pendingFlagRequest <- true
		}
	}()

	cacheKey := c.flagCacheKey(flag.Key)
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
		cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)

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
		return fmt.Errorf("failed to unmarshal company data: %v", err)
	}

	if company == nil {
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypeDelete {
		c.companyMu.Lock()
		c.companyCache.deleteEntity(ctx, company, c.logger)
		c.companyMu.Unlock()
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypePartial {
		id, err := ExtractIDFromJSON(resp.Data)
		if err != nil {
			return fmt.Errorf("failed to extract company ID from partial message: %v", err)
		}

		c.companyMu.Lock()
		existing, found := c.companyCacheProvider.Get(ctx, c.companyIDCacheKey(id))
		if !found || existing == nil {
			c.companyMu.Unlock()
			c.logger.Warn(ctx, fmt.Sprintf("Cache miss for partial company '%s', skipping", id))
			return nil
		}
		merged, mergeErr := PartialCompany(existing, resp.Data)
		if mergeErr != nil {
			c.companyMu.Unlock()
			return fmt.Errorf("failed to merge partial company: %v", mergeErr)
		}
		company = merged
		cacheResults, cacheErr := c.cacheCompanyForKeys(ctx, company)
		c.companyMu.Unlock()

		if cacheErr != nil {
			return cacheErr
		}
		for cacheKey, err := range cacheResults {
			if err != nil {
				c.logger.Warn(ctx, fmt.Sprintf("Cache error for company key '%s': %v", cacheKey, err))
			}
		}
		c.notifyPendingCompanyRequests(ctx, company.Keys, company)
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
		cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
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
		return fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	if user == nil {
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypeDelete {
		c.userMu.Lock()
		c.userCache.deleteEntity(ctx, user, c.logger)
		c.userMu.Unlock()
		return nil
	}

	if resp.MessageType == schematicdatastreamws.MessageTypePartial {
		id, err := ExtractIDFromJSON(resp.Data)
		if err != nil {
			return fmt.Errorf("failed to extract user ID from partial message: %v", err)
		}

		c.userMu.Lock()
		existing, found := c.userCacheProvider.Get(ctx, c.userIDCacheKey(id))
		if !found || existing == nil {
			c.userMu.Unlock()
			c.logger.Warn(ctx, fmt.Sprintf("Cache miss for partial user '%s', skipping", id))
			return nil
		}
		merged, mergeErr := PartialUser(existing, resp.Data)
		if mergeErr != nil {
			c.userMu.Unlock()
			return fmt.Errorf("failed to merge partial user: %v", mergeErr)
		}
		user = merged
		cacheResults := c.cacheUserForKeys(ctx, user)
		c.userMu.Unlock()

		for cacheKey, err := range cacheResults {
			if err != nil {
				c.logger.Warn(ctx, fmt.Sprintf("Cache error for user key '%s': %v", cacheKey, err))
			}
		}
		c.notifyPendingUserRequests(ctx, user.Keys, user)
		return nil
	}

	c.userMu.Lock()
	cacheResults := c.cacheUserForKeys(ctx, user)
	c.userMu.Unlock()

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
		return fmt.Errorf("failed to unmarshal error message: %v", err)
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

	var cachedCompany *rulesengine.Company
	var cachedUser *rulesengine.User

	// Try to get cached data first
	if needsCompany {
		cachedCompany = c.getCompanyFromCache(evalCtx.Company)
	}
	if needsUser {
		cachedUser = c.getUserFromCache(evalCtx.User)
	}

	// If we have all cached data we need, use it
	if (!needsCompany || cachedCompany != nil) && (!needsUser || cachedUser != nil) {
		// Evaluate against the rules engine with cached data
		resp, err := rulesengine.CheckFlag(ctx, cachedCompany, cachedUser, flag)
		if err != nil {
			return nil, fmt.Errorf("rules engine error: %w", err)
		}
		return resp, nil
	}

	// Handle replicator mode behavior
	if c.replicatorMode {
		// In replicator mode, if we don't have all cached data, evaluate with nil values instead of fetching
		// The external replicator should have populated the cache with all necessary data
		resp, err := rulesengine.CheckFlag(ctx, cachedCompany, cachedUser, flag)
		if err != nil {
			return nil, fmt.Errorf("rules engine error: %w", err)
		}
		return resp, nil
	}

	// Otherwise, check if we're connected to datastream
	if !c.IsConnected() {
		return nil, fmt.Errorf("datastream not connected")
	}

	// Fetch missing data from datastream
	var company *rulesengine.Company
	var user *rulesengine.User
	var err error

	if needsCompany {
		if cachedCompany != nil {
			company = cachedCompany
		} else {
			company, err = c.getCompany(ctx, evalCtx.Company)
			if err != nil {
				return nil, fmt.Errorf("failed to get company data: %w", err)
			}
		}
	}

	if needsUser {
		if cachedUser != nil {
			user = cachedUser
		} else {
			user, err = c.getUser(ctx, evalCtx.User)
			if err != nil {
				return nil, fmt.Errorf("failed to get user data: %w", err)
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
		cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
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
			cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
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
		cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixUser, key, value) // If there are already pending requests for this key, just add our channel
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
			cacheKey := c.resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
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
	flag, exists := c.flagsCacheProvider.Get(ctx, c.flagCacheKey(key))
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
	ctx := context.Background()

	company, _ := c.companyCache.getFromCache(ctx, keys)
	return company
}

func (c *DataStreamClient) getUserFromCache(keys map[string]string) *rulesengine.User {
	c.userMu.RLock()
	defer c.userMu.RUnlock()
	ctx := context.Background()

	user, _ := c.userCache.getFromCache(ctx, keys)
	return user
}

func (c *DataStreamClient) cacheVersion() string {
	versionKey := c.GetCacheVersion()
	if versionKey == "" {
		versionKey = rulesengine.VersionKey
	}
	return versionKey
}

func (c *DataStreamClient) flagCacheKey(key string) string {
	return fmt.Sprintf("%s:%s:%s:%s", cacheKeyPrefix, cacheKeyPrefixFlags, c.cacheVersion(), strings.ToLower(key))
}

func (c *DataStreamClient) resourceKeyToCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", cacheKeyPrefix, resourceType, c.cacheVersion(), strings.ToLower(key), strings.ToLower(value))
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

func (c *DataStreamClient) cacheUserForKeys(ctx context.Context, user *rulesengine.User) map[string]error {
	if user == nil || len(user.Keys) == 0 {
		return nil
	}
	return c.userCache.cacheForKeys(ctx, user, c.cacheTTL)
}

func (c *DataStreamClient) cacheCompanyForKeys(ctx context.Context, company *rulesengine.Company) (map[string]error, error) {
	if company == nil {
		return nil, errors.New("company cannot be nil")
	}
	if len(company.Keys) == 0 {
		return nil, errors.New("no keys provided for company lookup")
	}
	return c.companyCache.cacheForKeys(ctx, company, c.cacheTTL), nil
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
	companyCopy := DeepCopyCompany(company)

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
		c.setReplicatorReady(false, "")
		c.logger.Debug(ctx, fmt.Sprintf("Replicator health check failed: %v", err))
		return
	}
	defer func() { _ = resp.Body.Close() }()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.setReplicatorReady(false, "")
		c.logger.Warn(ctx, fmt.Sprintf("Failed to read replicator health response: %v", err))
		return
	}

	// Parse the JSON response
	var healthResp struct {
		Ready        bool   `json:"ready"`
		CacheVersion string `json:"cache_version"`
	}

	if err := json.Unmarshal(body, &healthResp); err != nil {
		c.setReplicatorReady(false, "")
		c.logger.Warn(ctx, fmt.Sprintf("Failed to parse replicator health response: %v", err))
		return
	}

	// Update replicator ready state and cache version
	wasReady := c.IsReplicatorReady()
	oldCacheVersion := c.GetCacheVersion()
	c.setReplicatorReady(healthResp.Ready, healthResp.CacheVersion)

	// Log state changes
	if healthResp.Ready && !wasReady {
		c.logger.Info(ctx, fmt.Sprintf("External replicator is now ready (cache_version: %s)", healthResp.CacheVersion))
	} else if !healthResp.Ready && wasReady {
		c.logger.Info(ctx, "External replicator is no longer ready")
	}

	// Log cache version changes
	if healthResp.CacheVersion != "" && healthResp.CacheVersion != oldCacheVersion {
		c.logger.Info(ctx, fmt.Sprintf("Replicator cache version changed: %s -> %s", oldCacheVersion, healthResp.CacheVersion))
	}
}

// IsReplicatorReady returns whether the external replicator is ready
func (c *DataStreamClient) IsReplicatorReady() bool {
	c.replicatorMu.RLock()
	defer c.replicatorMu.RUnlock()
	return c.replicatorReady
}

// setReplicatorReady updates the replicator ready state and cache version
func (c *DataStreamClient) setReplicatorReady(ready bool, cacheVersion string) {
	c.replicatorMu.Lock()
	defer c.replicatorMu.Unlock()
	c.replicatorReady = ready
	// Only update cache version if a non-empty value is provided
	// This preserves the existing version when the replicator is temporarily unavailable
	if cacheVersion != "" {
		c.replicatorCacheVersion = cacheVersion
	}
}

// GetCacheVersion returns the cache version from the replicator (in replicator mode)
// or the local rulesengine version key (in direct WebSocket mode)
func (c *DataStreamClient) GetCacheVersion() string {
	if c.replicatorMode {
		c.replicatorMu.RLock()
		defer c.replicatorMu.RUnlock()
		return c.replicatorCacheVersion
	}
	return rulesengine.VersionKey
}

// IsReplicatorMode returns whether the client is running in replicator mode
func (c *DataStreamClient) IsReplicatorMode() bool {
	return c.replicatorMode
}
