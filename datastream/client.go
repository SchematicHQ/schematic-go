package datastream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

func NewDataStream(baseUrl string, logger core.Logger, apiKey string, options *core.DatastreamOptions) *DataStreamClient {
	var flagCacheProvider FlagCacheProvider
	var companyCacheProvider CompanyCacheProvider
	var userCacheProvider UserCacheProvider
	if options.CacheProvider == defaultCacheProvider {
		companyCacheProvider = cache.NewDefaultCache[*rulesengine.Company]()
		userCacheProvider = cache.NewDefaultCache[*rulesengine.User]()
	}

	flagCacheProvider = cache.NewLocalCache[*rulesengine.Flag](1000, 1000*time.Hour)

	dataStreamUrl, err := getBaseURL(baseUrl)
	if err != nil {
		panic(fmt.Sprintf("failed to parse base URL: %v", err))
	}

	return &DataStreamClient{
		apiKey:               apiKey,
		cacheTTL:             options.CacheTTL,
		done:                 make(chan bool, 1),
		reconnect:            make(chan bool, 1),
		logger:               logger,
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
		if attempts >= maxReconnectAttempts {
			c.logger.Error(ctx, "Unable to connect to server")
			return
		}
		conn, err := c.connect()
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to connect to WebSocket: %v", err))
			attempts += 1
			time.Sleep(reconnectDelay)
			continue
		}

		c.logger.Info(ctx, "Connected to Schematic WebSocket")
		attempts = 0
		c.conn = conn

		err = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to set read deadline: %v", err))
		}

		go c.readMessages(ctx)
		closed := c.handleWebSocketConnection(ctx)

		if closed {
			return
		}

		c.logger.Info(ctx, "Reconnecting to WebSocket...")
	}
}

func (c *DataStreamClient) handleWebSocketConnection(ctx context.Context) bool {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()

	c.conn.SetPongHandler(c.handlePong)
	err := c.GetAllFlags(ctx)
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
			c.logger.Error(ctx, fmt.Sprintf("Fatal error occurred in WebSocket reader: %v", r))
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

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				return
			}
			c.logger.Debug(ctx, fmt.Sprintf("Failed to read WebSocket message: %v", err))
			c.reconnect <- true
		}

		err = json.Unmarshal(m, &message)
		if message.Data == nil {
			c.logger.Debug(ctx, "Received empty message from WebSocket")
			c.reconnect <- true
			return
		}

		err = c.handleMessageResponse(ctx, &message)
		if err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to handle WebSocket message: %v", err))
		}
	}
}

func (c *DataStreamClient) SendWebSocketMessage(ctx context.Context, req *DataStreamReq) error {
	if c.conn == nil {
		return fmt.Errorf("WebSocket connection is not initialized")
	}

	message := c.packageMessage(req)

	err := c.conn.WriteJSON(message)
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

func (c *DataStreamClient) GetAllFlags(ctx context.Context) error {
	// Check if there is already a pending request for flags
	if c.pendingFlagRequest != nil {
		return nil
	}

	waitCh := make(chan bool, 1)
	c.mu.Lock()
	c.pendingFlagRequest = waitCh
	c.mu.Unlock()
	defer func() {
		c.pendingFlagRequest = nil
		close(waitCh)
	}()

	req := &DataStreamReq{
		EntityType: EntityTypeFlags,
	}
	err := c.SendWebSocketMessage(ctx, req)
	if err != nil {
		c.logger.Error(ctx, fmt.Sprintf("Failed to send WebSocket message: %v", err))
		return err
	}

	select {
	case <-waitCh:
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout while waiting for flags data")
	case <-ctx.Done():
		c.pendingFlagRequest = nil
		return ctx.Err()
	}
	return nil
}

func (c *DataStreamClient) handleFlagsMessage(ctx context.Context, resp *DataStreamResp) error {
	flags := resp.Data

	var flagsData []*rulesengine.Flag
	err := json.Unmarshal(flags, &flagsData)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to unmarshal flags data: %v", err))
	}

	for _, flag := range flagsData {
		c.flagsCacheProvider.Set(ctx, flagCacheKey(flag.Key), flag, nil)
	}

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

	for key, value := range company.Keys {
		companyKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		ttl := c.cacheTTL
		c.companyCacheProvider.Set(ctx, companyKey, company, &ttl)
	}

	// Notify all goroutines waiting for this company
	c.mu.Lock()
	defer c.mu.Unlock()
	// For each relevant key, notify all waiting channels
	for key, value := range company.Keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		if channels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			for _, ch := range channels {
				// Non-blocking send - if the channel is full, we'll just continue
				select {
				case ch <- company:
				default:
				}
			} // Remove this set of channels as they've been notified
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
			return nil
		}
	}

	for key, value := range user.Keys {
		companyKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		ttl := c.cacheTTL
		c.userCacheProvider.Set(ctx, companyKey, user, &ttl)
	}

	// Notify all goroutines waiting for this company
	c.mu.Lock()
	defer c.mu.Unlock()
	// For each relevant key, notify all waiting channels
	for key, value := range user.Keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
		if channels, ok := c.pendingUserRequests[cacheKey]; ok {
			for _, ch := range channels {
				// Non-blocking send - if the channel is full, we'll just continue
				select {
				case ch <- user:
				default:
				}
			} // Remove this set of channels as they've been notified
			delete(c.pendingUserRequests, cacheKey)
		}
	}

	return nil
}

func (c *DataStreamClient) GetCompany(ctx context.Context, keys map[string]string) (*rulesengine.Company, error) {
	company := c.getCompanyFromCache(keys)
	if company != nil {
		return company, nil
	}

	waitCh := make(chan *rulesengine.Company, 1) // Register the wait channel for each key
	cacheKeys := []string{}
	c.mu.Lock()
	shouldSendRequest := true
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
		cacheKeys = append(cacheKeys, cacheKey)
		if existingChannels, ok := c.pendingCompanyRequests[cacheKey]; ok {
			c.pendingCompanyRequests[cacheKey] = append(existingChannels, waitCh)
			shouldSendRequest = false // Someone else already requested this
		}
	}
	c.mu.Unlock()

	if shouldSendRequest {
		req := &DataStreamReq{
			EntityType: EntityTypeCompany,
			Keys:       keys,
		}

		err := c.SendWebSocketMessage(ctx, req)
		if err != nil {
			return nil, err
		}

		c.mu.Lock()
		for key, value := range keys {
			cacheKey := resourceKeyToCacheKey(cacheKeyPrefixCompany, key, value)
			c.pendingCompanyRequests[cacheKey] = []chan *rulesengine.Company{waitCh}
		}
		c.mu.Unlock()
	}

	var err error
	select {
	case company := <-waitCh:
		return company, nil
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("timeout while waiting for company data")
	case <-ctx.Done():
		err = ctx.Err()
	}

	// Clean up all channels involved in this request
	c.cleanupPendingCompanyRequests(cacheKeys, waitCh)

	return nil, err
}

func (c *DataStreamClient) GetUser(ctx context.Context, keys map[string]string) (*rulesengine.User, error) {
	user := c.getUserFromCache(keys)
	if user != nil {
		return user, nil
	}

	waitCh := make(chan *rulesengine.User, 1) // Register the wait channel for each key
	cacheKeys := []string{}
	c.mu.Lock()
	shouldSendRequest := true
	for key, value := range keys {
		cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value) // If there are already pending requests for this key, just add our channel
		cacheKeys = append(cacheKeys, cacheKey)
		if existingChannels, ok := c.pendingUserRequests[cacheKey]; ok {
			c.pendingUserRequests[cacheKey] = append(existingChannels, waitCh)
			shouldSendRequest = false // Someone else already requested this
		}
	}
	c.mu.Unlock()

	if shouldSendRequest {
		req := &DataStreamReq{
			EntityType: EntityTypeUser,
			Keys:       keys,
		}

		err := c.SendWebSocketMessage(ctx, req)
		if err != nil {
			return nil, err
		}

		c.mu.Lock()
		for key, value := range keys {
			cacheKey := resourceKeyToCacheKey(cacheKeyPrefixUser, key, value)
			c.pendingUserRequests[cacheKey] = []chan *rulesengine.User{waitCh}
		}
		c.mu.Unlock()
	}

	var err error
	select {
	case user := <-waitCh:
		return user, nil
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("timeout while waiting for user data")
	case <-ctx.Done():
		err = ctx.Err()
	}

	// Clean up all channels involved in this request
	c.cleanupPendingUserRequests(cacheKeys, waitCh)

	return nil, err
}

func (c *DataStreamClient) GetFlag(ctx context.Context, key string) (*rulesengine.Flag, bool) {
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
	return fmt.Sprintf("%s:%s", cacheKeyPrefixFlags, key)
}

func resourceKeyToCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("%s:%s:%s", resourceType, key, value)
}

// Helper function to clean up pending company requests
func (c *DataStreamClient) cleanupPendingCompanyRequests(cacheKeys []string, waitCh chan *rulesengine.Company) {
	c.mu.Lock()
	defer c.mu.Unlock()
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
	c.mu.Lock()
	defer c.mu.Unlock()
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
				delete(c.pendingCompanyRequests, cacheKey)
			}
		}
	}

	close(waitCh) // Close the wait channel to signal that we're done with it
}
