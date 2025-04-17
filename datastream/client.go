package datastream

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/logger"
)

func NewDataStream(baseUrl string, apiKey string) *DataStreamClient {

	flagCacheProvider := cache.NewDefaultCache[*rulesengine.Flag]()

	companyCacheProvider := cache.NewDefaultCache[*rulesengine.Company]()

	userCacheProvider := cache.NewDefaultCache[*rulesengine.User]()

	return &DataStreamClient{
		apiKey:               apiKey,
		done:                 make(chan bool),
		reconnect:            make(chan bool),
		logger:               logger.NewDefaultLogger(),
		flagsCacheProvider:   flagCacheProvider,
		companyCacheProvider: companyCacheProvider,
		companyCache:         make(map[string]*rulesengine.Company),
		url:                  getBaseURL(baseUrl),
		userCacheProvider:    userCacheProvider,
	}
}

func connect(addr string, apiKey string) (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/datastream"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{
		"X-Schematic-Api-Key": []string{apiKey},
	})
	return c, err
}

func (c *DataStreamClient) Start() {
	go c.ConnectAndRead()
}

func (c *DataStreamClient) ConnectAndRead() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred in WebSocket handler %v", r)
		}
	}()
	defer c.Close()

	ctx := context.Background()

	attempts := 0
	for {
		if attempts >= maxReconnectAttempts {
			c.logger.Printf("ERROR: Unable to connect to server")
			return
		}
		conn, err := connect(c.url, c.apiKey)
		if err != nil {
			c.logger.Printf("ERROR: Failed to connect to WebSocket: %v", err)
			attempts += 1
			time.Sleep(reconnectDelay)
			continue
		}

		c.logger.Printf("INFO: Connected to Schematic WebSocket")
		attempts = 0
		c.conn = conn

		go c.readMessages(ctx)
		closed := c.handleWebSocketConnection(ctx)

		if closed {
			return
		}

		c.logger.Printf("INFO: Reconnecting to WebSocket...")
	}
}

func (c *DataStreamClient) handleWebSocketConnection(ctx context.Context) bool {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()

	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(c.handlePong)
	c.GetAllFlags(ctx)

	for {
		select {
		case <-c.done:
			return true
		case <-c.reconnect:
			return false
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				c.logger.Printf("ERROR: Failed to send ping message: %v", err)
				return false
			}
		}
	}
}

func (c *DataStreamClient) readMessages(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred in WebSocket reader %v", r)
			c.reconnect <- true
		}
	}()
	for {
		var message DataStreamResp
		err := c.conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Printf("ERROR: Failed to read WebSocket message: %v", err)
				c.reconnect <- true
				return
			}
		}

		err = c.handleMessageResponse(ctx, &message)
		if err != nil {
			c.logger.Printf("ERROR: Failed to handle WebSocket message: %v", err)
			return
		}
	}
}

func (c *DataStreamClient) SendWebSocketMessage(ctx context.Context, req *DataStreamReq) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while sending WebSocket message %v", r)
		}
	}()

	if c.conn == nil {
		return fmt.Errorf("WebSocket connection is not initialized")
	}

	message := c.packageMessage(req)

	err := c.conn.WriteJSON(message)
	if err != nil {
		c.logger.Printf("ERROR: Failed to send WebSocket message: %v", err)
		return err
	}

	return nil
}

func (c *DataStreamClient) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while closing WebSocket %v", r)
		}
	}()

	if c.conn != nil {
		c.logger.Printf("INFO: Closing WebSocket connection")
		c.conn.Close()
	}

	close(c.done)
	c.logger.Printf("INFO: WebSocket connection closed")
}

func getBaseURL(baseUrl string) string {
	if baseUrl == "" {
		return defaultBaseURL
	}

	url, err := url.Parse(baseUrl)
	if err != nil {
		fmt.Printf("ERROR: Invalid base URL: %v", err)
		return defaultBaseURL
	}

	return url.Host
}

func (c *DataStreamClient) handleMessageResponse(ctx context.Context, message *DataStreamResp) error {
	if message == nil {
		return nil
	}

	switch message.EntityType {
	case string(EntityTypeCompany):
		c.handleCompanyMessage(ctx, message)
	case string(EntityTypeFlags):
		c.handleFlagsMessage(ctx, message)
	case string(EntityTypeUser):
		c.handleUserMessage(ctx, message)
	default:
		c.logger.Printf("ERROR: Received unknown entity type: %s", message.EntityType)
	}

	return nil
}

func (c *DataStreamClient) GetAllFlags(ctx context.Context) error {
	req := &DataStreamReq{
		EntityType: EntityTypeFlags,
	}

	err := c.SendWebSocketMessage(ctx, req)
	if err != nil {
		c.logger.Printf("ERROR: Failed to send WebSocket message: %v", err)
		return err
	}

	return nil
}

func (c *DataStreamClient) handleFlagsMessage(ctx context.Context, resp *DataStreamResp) {
	flags := resp.Data

	var flagsData []*rulesengine.Flag
	err := json.Unmarshal(flags, &flagsData)
	if err != nil {
		c.logger.Printf("ERROR: Failed to unmarshal flags data: %v", err)
		return
	}

	for _, flag := range flagsData {
		ttl := defaultExpiration
		c.flagsCacheProvider.Set(ctx, flag.Key, flag, &ttl)
	}
}

func (c *DataStreamClient) handleCompanyMessage(ctx context.Context, resp *DataStreamResp) {
	var company *rulesengine.Company
	err := json.Unmarshal(resp.Data, &company)
	if err != nil {
		c.logger.Printf("ERROR: Failed to unmarshal company data: %v", err)
		return
	}

	if company == nil {
		return
	}

	for key, value := range company.Keys {
		companyKey := resourceKeyToCacheKey("company", key, value)
		ttl := defaultExpiration
		c.companyCacheProvider.Set(ctx, companyKey, company, &ttl)
	}
}

func (c *DataStreamClient) handleUserMessage(ctx context.Context, resp *DataStreamResp) {
	var user *rulesengine.User
	err := json.Unmarshal(resp.Data, &user)
	if err != nil {
		c.logger.Printf("ERROR: Failed to unmarshal user data: %v", err)
		return
	}

	for key, value := range user.Keys {
		companyKey := resourceKeyToCacheKey("user", key, value)
		ttl := defaultExpiration
		c.userCacheProvider.Set(ctx, companyKey, user, &ttl)
	}
}

func (c *DataStreamClient) GetCompany(ctx context.Context, keys map[string]string) (*rulesengine.Company, error) {

	company := c.getCompanyFromCache(keys)
	if company != nil {
		return company, nil
	}
	// If not found in cache, send a request to the server
	// and wait for the response

	req := &DataStreamReq{
		EntityType: EntityTypeCompany,
		Keys:       keys,
	}

	err := c.SendWebSocketMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	done := make(chan *rulesengine.Company, 1)
	go func() {
		// TODO: Handle panic
		for {
			company := c.getCompanyFromCache(keys)
			if company != nil {
				done <- company
			}
		}
	}()

	for {
		select {
		case company := <-done:
			return company, nil
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("timeout while waiting for company data")
		}
	}
}

func (c *DataStreamClient) GetUser(ctx context.Context, keys map[string]string) (*rulesengine.User, error) {

	user := c.getUserFromCache(keys)
	if user != nil {
		return user, nil
	}
	// If not found in cache, send a request to the server
	// and wait for the response

	req := &DataStreamReq{
		EntityType: EntityTypeUser,
		Keys:       keys,
	}

	err := c.SendWebSocketMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	done := make(chan *rulesengine.User, 1)
	go func() {
		// TODO: Handle panic
		for {
			user := c.getUserFromCache(keys)
			if user != nil {
				done <- user
			}
		}
	}()

	for {
		select {
		case user := <-done:
			return user, nil
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("timeout while waiting for user data")
		}
	}
}

func (c *DataStreamClient) GetFlag(ctx context.Context, key string) (*rulesengine.Flag, bool) {
	flag, exists := c.flagsCacheProvider.Get(ctx, key)
	if !exists {
		return nil, false
	}
	return flag, true
}

func (c *DataStreamClient) handlePong(string) error {
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.logger.Printf("setting read deadline %v", time.Now().Add(pongWait))
	return nil
}

func (c *DataStreamClient) packageMessage(req *DataStreamReq) *DataStreamBaseReq {
	return &DataStreamBaseReq{
		Data: *req,
	}
}

func (c *DataStreamClient) getCompanyFromCache(keys map[string]string) *rulesengine.Company {
	for key, value := range keys {
		companyKey := resourceKeyToCacheKey("company", key, value)
		company, exists := c.companyCacheProvider.Get(context.Background(), companyKey)
		if exists {
			return company
		}
	}

	return nil
}

func (c *DataStreamClient) getUserFromCache(keys map[string]string) *rulesengine.User {
	for key, value := range keys {
		userKey := resourceKeyToCacheKey("user", key, value)
		user, exists := c.userCacheProvider.Get(context.Background(), userKey)
		if exists {
			return user
		}
	}

	return nil
}

func resourceKeyToCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("%s:%s:%s", resourceType, key, value)
}
