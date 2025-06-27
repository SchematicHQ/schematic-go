package datastream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/cache"
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

	// Create flag cache if not provided in options
	flagCacheProvider := options.FlagCache
	if flagCacheProvider == nil {
		flagCacheProvider = cache.NewLocalCache[*rulesengine.Flag](1000, -1)
	}

	dataStreamUrl, err := getBaseURL(options.BaseURL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse base URL: %v", err))
	}

	return &DataStreamClient{
		apiKey:               options.ApiKey,
		cacheTTL:             configurationOptions.CacheTTL,
		done:                 make(chan bool, 1),
		ctxErrors:            make(chan *core.CtxError, 100),
		reconnect:            make(chan bool, 1),
		monitorChannel:       options.MonitorChannel,
		logger:               options.Logger,
		flagsCacheProvider:   flagCacheProvider,
		companyCacheProvider: companyCacheProvider,
		companyCache:         make(map[string]*rulesengine.Company),
		url:                  dataStreamUrl,
		userCacheProvider:    userCacheProvider,

		pendingCompanyRequests: make(map[string][]chan *rulesengine.Company),
		pendingUserRequests:    make(map[string][]chan *rulesengine.User),
	}
}

func (c *DataStreamClient) connect() (*websocket.Conn, error) {
	client, _, err := websocket.DefaultDialer.Dial(c.url.String(), http.Header{
		"X-Schematic-Api-Key": []string{c.apiKey},
	})
	return client, err
}

func (c *DataStreamClient) Start() {
	go c.ConnectAndRead()
}

func (c *DataStreamClient) ConnectAndRead() {
	ctx := context.Background()

	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, fmt.Sprintf("Fatal error occurred in WebSocket handler: %v", r))
		}
	}()

	attempts := 0
	for {
		conn, err := c.connect()
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to connect to WebSocket: %v", err))
			attempts += 1
			c.monitorChannel <- false
			if attempts >= maxReconnectAttempts {
				c.logger.Error(ctx, "Unable to connect to server")
				return
			}
			time.Sleep(calculateBackoffDelay(attempts))
			continue
		}

		c.logger.Info(ctx, "Connected to Schematic WebSocket")
		attempts = 0
		c.conn = conn

		c.monitorChannel <- true

		err = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to set read deadline: %v", err))
		}

		go c.readMessages(ctx)
		closed := c.handleWebSocketConnection(ctx)

		if closed {
			c.monitorChannel <- false
			return
		}

		c.logger.Info(ctx, "Reconnecting to WebSocket...")
	}
}

func calculateBackoffDelay(attempt int) time.Duration {
	// Add jitter to prevent synchronized reconnection attempts
	jitter := time.Duration(rand.Int63n(int64(minReconnectDelay)))

	// Exponential backoff with a cap
	delay := time.Duration(math.Pow(2, float64(attempt-1)))*minReconnectDelay + jitter
	if delay > maxReconnectDelay {
		delay = maxReconnectDelay + jitter // Ensure jitter is added even when capped
	}
	return delay
}

func (c *DataStreamClient) handleWebSocketConnection(ctx context.Context) bool {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()

	c.conn.SetPongHandler(c.handlePong)
	err := c.getAllFlags(ctx)
	if err != nil {
		c.logger.Error(ctx, fmt.Sprintf("Failed to get all flags: %v", err))
		return true
	}

	for {
		select {
		case <-c.done:
			return true
		case <-c.reconnect:
			return false
		case err := <-c.ctxErrors:
			c.logger.Error(err.Ctx, fmt.Sprintf("%v", err.Err))
		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				c.logger.Error(ctx, fmt.Sprintf("Failed to send ping message: %v", err))
				return false
			}
		}
	}
}

func (c *DataStreamClient) readMessages(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.ctxErrors <- &core.CtxError{
				Ctx: ctx,
				Err: errors.New(fmt.Sprintf("Fatal error occurred in WebSocket reader: %v", r)),
			}
			return
		}
	}()
	for {
		var message DataStreamResp
		_, m, err := c.conn.ReadMessage()

		if err != nil {
			var opErr *net.OpError
			if errors.As(err, &opErr) {
				return
			}

			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				return
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Debug(ctx, fmt.Sprintf("Failed to read WebSocket message: %v", err))
				return
			}

			c.reconnect <- true
			return
		}

		err = json.Unmarshal(m, &message)
		if message.Data == nil {
			c.logger.Debug(ctx, "Received empty message from WebSocket")
			c.reconnect <- true
			return
		}

		err = c.handleMessageResponse(ctx, &message)
		if err != nil {
			c.ctxErrors <- &core.CtxError{
				Ctx: ctx,
				Err: errors.New(fmt.Sprintf("Failed to handle WebSocket message: %v", err)),
			}
		}
	}
}

func (c *DataStreamClient) sendWebSocketMessage(ctx context.Context, req *DataStreamReq) error {
	if c.conn == nil {
		return fmt.Errorf("WebSocket connection is not initialized")
	}

	message := c.packageMessage(req)

	c.writeMu.Lock()
	err := c.conn.WriteJSON(message)
	c.writeMu.Unlock()
	if err != nil {
		c.logger.Error(ctx, fmt.Sprintf("Failed to send WebSocket message: %v", err))
		return err
	}

	return nil
}

func (c *DataStreamClient) Close() {
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, fmt.Sprintf("Fatal error occurred while closing WebSocket: %v", r))
		}
	}()

	close(c.done)
	close(c.reconnect)

	if c.conn != nil {
		c.logger.Info(ctx, "Closing WebSocket connection")
		c.conn.Close()
	}

	c.logger.Info(ctx, "WebSocket connection closed")
}

func getBaseURL(baseUrl string) (*url.URL, error) {
	if baseUrl == "" {
		baseUrl = defaultBaseURL
	}

	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	if url.Scheme == "https" {
		url.Scheme = "wss"
	} else if url.Scheme == "http" {
		url.Scheme = "ws"
	}

	url.Path = "/datastream"

	return url, nil
}

func (c *DataStreamClient) handleMessageResponse(ctx context.Context, message *DataStreamResp) error {
	if message == nil {
		return nil
	}

	if message.MessageType == MessageTypeError {
		return c.handleErrorMessage(ctx, message)
	}

	switch message.EntityType {
	case string(EntityTypeCompany):
		return c.handleCompanyMessage(ctx, message)
	case string(EntityTypeFlags):
		return c.handleFlagsMessage(ctx, message)
	case string(EntityTypeUser):
		return c.handleUserMessage(ctx, message)
	default:
		return errors.New(fmt.Sprintf("Received unknown entity type: %s", message.EntityType))
	}
}

func (c *DataStreamClient) handleFlagsMessage(ctx context.Context, resp *DataStreamResp) error {
	flags := resp.Data

	var flagsData []*rulesengine.Flag
	err := json.Unmarshal(flags, &flagsData)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal flags data: %v", err))
	}

	var cacheKeys []string
	for _, flag := range flagsData {
		cacheKey := flagCacheKey(flag.Key)
		c.flagsCacheProvider.Set(ctx, cacheKey, flag, nil)
		cacheKeys = append(cacheKeys, cacheKey)
	}

	c.flagsCacheProvider.DeleteMissing(ctx, cacheKeys)

	if c.pendingFlagRequest != nil {
		c.pendingFlagRequest <- true
	}

	return nil
}

func (c *DataStreamClient) handleCompanyMessage(ctx context.Context, resp *DataStreamResp) error {
	var company *rulesengine.Company
	err := json.Unmarshal(resp.Data, &company)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal company data: %v", err))
	}

	if company == nil {
		return nil
	}

	if resp.MessageType == MessageTypeDelete {
		// Remove the company from the cache
		for key, value := range company.Keys {
			companyKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
			c.companyCacheProvider.Delete(ctx, companyKey)
		}
		return nil
	}

	cached := true

	c.companyMu.Lock()
	cacheResults, err := c.cacheCompanyForKeys(ctx, company)
	c.companyMu.Unlock()

	if err != nil {
		return err
	}

	// Notify all goroutines waiting for this company
	c.pendingCompReqMu.Lock()
	defer c.pendingCompReqMu.Unlock()

	// For each relevant key, notify all waiting channels
	for key, value := range company.Keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)

		// Check if there are channels waiting for this key
		if channels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			// Only notify channels if the cache operation for this key was successful
			if cacheErr, exists := cacheResults[cacheKey]; exists && cacheErr == nil {
				for _, ch := range channels {
					// Non-blocking send - if the channel is full, we'll just continue
					select {
					case ch <- company:
					default:
					}
				}
			} else {
				c.logger.Warn(ctx, fmt.Sprintf("Not notifying pending requests for key '%s' due to cache error", cacheKey))
			}

			// Remove this set of channels as they've been notified or informed of failure
			delete(c.pendingCompanyRequests, cacheKey)
		}
	}

	return nil
}

func (c *DataStreamClient) handleUserMessage(ctx context.Context, resp *DataStreamResp) error {
	var user *rulesengine.User
	err := json.Unmarshal(resp.Data, &user)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal user data: %v", err))
	}

	if user == nil {
		return nil
	}

	if resp.MessageType == MessageTypeDelete {
		// Remove the user from the cache
		for key, value := range user.Keys {
			userKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
			c.userCacheProvider.Delete(ctx, userKey)
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

	// Notify all goroutines waiting for this user
	c.pendingUserReqMu.Lock()
	defer c.pendingUserReqMu.Unlock()

	// For each relevant key, notify all waiting channels
	for key, value := range user.Keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		if channels, ok := c.pendingUserRequests[cacheKey]; ok {
			// Only notify channels if the cache operation for this key was successful
			if cacheErr, exists := cacheResults[cacheKey]; exists && cacheErr == nil {
				for _, ch := range channels {
					// Non-blocking send - if the channel is full, we'll just continue
					select {
					case ch <- user:
					default:
					}
				}
			} else {
				c.logger.Warn(ctx, fmt.Sprintf("Not notifying pending requests for user key '%s' due to cache error", cacheKey))
			}
			// Remove this set of channels as they've been notified or informed of failure
			delete(c.pendingUserRequests, cacheKey)
		}
	}

	return nil
}

func (c *DataStreamClient) handleErrorMessage(_ context.Context, resp *DataStreamResp) error {
	var respError DataStreamError
	err := json.Unmarshal(resp.Data, &respError)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal error message: %v", err))
	}
	return errors.New(fmt.Sprintf("%s", respError.Error))
}

func (c *DataStreamClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) *rulesengine.CheckFlagResult {
	var company *rulesengine.Company
	var err error
	if evalCtx.Company != nil {
		company, err = c.getCompany(ctx, evalCtx.Company)
		if err != nil {
			c.ctxErrors <- &core.CtxError{
				Ctx: ctx,
				Err: errors.New(fmt.Sprintf("Failed to get company from cache: %v", err)),
			}
			return nil
		}
	}

	var user *rulesengine.User
	if evalCtx.User != nil {
		user, err = c.getUser(ctx, evalCtx.User)
		if err != nil {
			c.ctxErrors <- &core.CtxError{
				Ctx: ctx,
				Err: errors.New(fmt.Sprintf("Failed to get user from cache: %v", err)),
			}
			return nil
		}
	}

	// Get flag here
	flag, found := c.getFlag(ctx, flagKey)
	if !found {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: errors.New(fmt.Sprintf("Flag %s not found", flagKey)),
		}
		return nil
	}

	// Evaluate against the rules engine
	resp, err := rulesengine.CheckFlag(ctx, company, user, flag)
	if err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}
		return nil
	}

	return resp
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

	req := &DataStreamReq{
		EntityType: EntityTypeFlags,
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
		req := &DataStreamReq{
			EntityType: EntityTypeCompany,
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
		req := &DataStreamReq{
			EntityType: EntityTypeUser,
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
	flag, exists := c.flagsCacheProvider.Get(ctx, flagCacheKey(key))
	if !exists {
		return nil, false
	}
	return flag, true
}

func (c *DataStreamClient) handlePong(string) error {
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return err
	}
	return nil
}

func (c *DataStreamClient) packageMessage(req *DataStreamReq) *DataStreamBaseReq {
	return &DataStreamBaseReq{
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
	return fmt.Sprintf("%s:%s:%s:%s", cacheKeyPrefix, cacheKeyPrefixFlags, RulesEngineVersionKey, strings.ToLower(key))
}

func resourceKeyToCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", cacheKeyPrefix, resourceType, RulesEngineVersionKey, strings.ToLower(key), strings.ToLower(value))
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
	keys := event.Company
	if len(keys) == 0 {
		return errors.New("no keys provided for company lookup")
	}

	if event == nil {
		return errors.New("event body cannot be nil")
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
