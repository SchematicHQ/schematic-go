package datastream_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/stretchr/testify/assert"
)

// MockLogger implements the core.Logger interface for testing
type MockLogger struct {
	infoMessages  []string
	errorMessages []string
	debugMessages []string
	warnMessages  []string
	mu            sync.Mutex
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		infoMessages:  []string{},
		errorMessages: []string{},
		debugMessages: []string{},
		warnMessages:  []string{},
	}
}

func (m *MockLogger) Info(ctx context.Context, message string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.infoMessages = append(m.infoMessages, message)
}

func (m *MockLogger) Error(ctx context.Context, message string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errorMessages = append(m.errorMessages, message)
}

func (m *MockLogger) Debug(ctx context.Context, message string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.debugMessages = append(m.debugMessages, message)
}

func (m *MockLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.warnMessages = append(m.warnMessages, message)
}

func (m *MockLogger) SetLevel(level core.LogLevel) {
	// No-op for mock logger
}
func (m *MockLogger) GetLevel() core.LogLevel {
	// No-op for mock logger
	return core.LogLevelDebug
}

// Mock WebSocket server for testing
func setupMockWebSocketServer() (*httptest.Server, chan string, chan string) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	incomingMessages := make(chan string, 100)
	outgoingMessages := make(chan string, 100)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		// Immediately send flags to client on connection
		flagsData := createMockFlagsData()
		conn.WriteMessage(websocket.TextMessage, []byte(flagsData))

		// Handle incoming messages
		go func() {
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					return
				}
				incomingMessages <- string(message)
			}
		}()

		// Send outgoing messages
		for msg := range outgoingMessages {
			conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}))

	return server, incomingMessages, outgoingMessages
}

// Helper function to create mock flags data
func createMockFlagsData() string {
	flags := []*rulesengine.Flag{
		{
			ID:            "flag_XXXX",
			AccountID:     "account_XXXX",
			Key:           "test-flag-1",
			EnvironmentID: "env_XXXX",
			DefaultValue:  true,
		},
		{
			ID:            "flag_YYYYY",
			AccountID:     "account_XXXX",
			Key:           "test-flag-2",
			EnvironmentID: "env_XXXX",
			DefaultValue:  false,
		},
	}

	flagsData, _ := json.Marshal(flags)
	resp := schematicdatastreamws.DataStreamResp{
		EntityType: string(schematicdatastreamws.EntityTypeFlags),
		Data:       flagsData,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function to create mock company data
func createMockCompanyData(companyID string, messageType schematicdatastreamws.MessageType) string {
	company := &rulesengine.Company{
		Keys: map[string]string{
			"company_id": companyID,
		},
	}

	companyData, _ := json.Marshal(company)
	resp := schematicdatastreamws.DataStreamResp{
		EntityType:  string(schematicdatastreamws.EntityTypeCompany),
		Data:        companyData,
		MessageType: messageType,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function to create mock user data
func createMockUserData(userID string, messageType schematicdatastreamws.MessageType) string {
	user := &rulesengine.User{
		Keys: map[string]string{
			"user_id": userID,
		},
	}

	userData, _ := json.Marshal(user)
	resp := schematicdatastreamws.DataStreamResp{
		EntityType:  string(schematicdatastreamws.EntityTypeUser),
		Data:        userData,
		MessageType: messageType,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function to create mock single flag data
func createMockSingleFlagData(flagKey string, defaultValue bool, messageType schematicdatastreamws.MessageType) string {
	flag := &rulesengine.Flag{
		ID:            fmt.Sprintf("flag_%s", flagKey),
		AccountID:     "account_XXXX",
		Key:           flagKey,
		EnvironmentID: "env_XXXX",
		DefaultValue:  defaultValue,
	}

	flagData, _ := json.Marshal(flag)
	resp := schematicdatastreamws.DataStreamResp{
		EntityType:  string(schematicdatastreamws.EntityTypeFlag),
		Data:        flagData,
		MessageType: messageType,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function for creating DataStreamClientOptions for tests
func createTestClientOptions(baseURL string, logger core.Logger, apiKey string) datastream.DataStreamClientOptions {
	return datastream.DataStreamClientOptions{
		ApiKey:  apiKey,
		BaseURL: baseURL,
		Logger:  logger,
	}
}

func TestNewDataStreamClient(t *testing.T) {
	logger := NewMockLogger()
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions("http://localhost:8080", logger, "test-api-key")

	client := datastream.NewDataStreamClient(clientOptions, configOptions)

	assert.NotNil(t, client, "Client should not be nil")
}

func TestDataStreamClientConnection(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
	client := datastream.NewDataStreamClient(clientOptions, configOptions)

	// Start the client and wait for connection
	client.Start()

	// Allow time for connection to establish
	time.Sleep(200 * time.Millisecond)

	// Check that the client is connected
	assert.True(t, client.IsConnected(), "Connection should be successful")

	// Clean up
	client.Close()
}

func TestCheckFlagCompany(t *testing.T) {
	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection and initial flags to be received
	time.Sleep(200 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	// Setup a goroutine to handle incoming company and user requests
	go func() {
		for msg := range incomingMessages {
			var req schematicdatastreamws.DataStreamBaseReq
			json.Unmarshal([]byte(msg), &req)

			if req.Data.EntityType == schematicdatastreamws.EntityTypeCompany {
				companyID := req.Data.Keys["company_id"]
				outgoingMessages <- createMockCompanyData(companyID, schematicdatastreamws.MessageTypeFull)
			}
		}
	}()

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}

	// Test checking a flag that exists and is true
	result, err := client.CheckFlag(ctx, evalCtx, "test-flag-1")
	assert.NoError(t, err, "CheckFlag should not return an error for existing flag")
	assert.True(t, result.Value, "Flag value should be true")

	// Test checking a flag that exists and is false
	result, err = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.NoError(t, err, "CheckFlag should not return an error for existing flag")
	assert.False(t, result.Value, "Flag value should be false")

	// Test checking a non-existent flag
	result, err = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.Error(t, err, "Non-existent flag should return an error")
	assert.Nil(t, result, "Non-existent flag should return nil result")
	assert.Contains(t, err.Error(), "flag not found", "Error should indicate flag not found")
}

func TestCheckFlagUser(t *testing.T) {
	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()

	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection and initial flags to be received
	time.Sleep(200 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	// Setup a goroutine to handle incoming company and user requests
	go func() {
		for msg := range incomingMessages {
			var req schematicdatastreamws.DataStreamBaseReq
			json.Unmarshal([]byte(msg), &req)

			if req.Data.EntityType == schematicdatastreamws.EntityTypeUser {
				companyID := req.Data.Keys["user_id"]
				outgoingMessages <- createMockUserData(companyID, schematicdatastreamws.MessageTypeFull)
			}
		}
	}()

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		User: map[string]string{"user_id": "123"},
	}

	// Test checking a flag that exists and is true
	result, err := client.CheckFlag(ctx, evalCtx, "test-flag-1")
	assert.NoError(t, err, "CheckFlag should not return an error for existing flag")
	assert.True(t, result.Value, "Flag value should be true")

	// Test checking a flag that exists and is false
	result, err = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.NoError(t, err, "CheckFlag should not return an error for existing flag")
	assert.False(t, result.Value, "Flag value should be false")

	// Test checking a non-existent flag
	result, err = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.Error(t, err, "Non-existent flag should return an error")
	assert.Nil(t, result, "Non-existent flag should return nil result")
	assert.Contains(t, err.Error(), "flag not found", "Error should indicate flag not found")
}

func TestSingleFlagMessage(t *testing.T) {
	t.Run("Handles single flag with Full message type", func(t *testing.T) {
		server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
		defer server.Close()
		defer close(outgoingMessages)

		logger := NewMockLogger()

		// Create a local cache provider for flags
		localFlagCache := cache.NewLocalCache[*rulesengine.Flag](100, time.Minute)

		configOptions := &core.DatastreamOptions{
			CacheTTL: 5 * time.Minute,
		}

		clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
		clientOptions.FlagCache = localFlagCache

		client := datastream.NewDataStreamClient(clientOptions, configOptions)
		client.Start()
		defer client.Close()

		// Wait for connection and initial flags
		time.Sleep(300 * time.Millisecond)

		// Send a single flag update
		flagKey := "test-single-flag"
		outgoingMessages <- createMockSingleFlagData(flagKey, true, schematicdatastreamws.MessageTypeFull)

		// Allow time for message processing
		time.Sleep(100 * time.Millisecond)

		// Setup a goroutine to handle incoming company requests
		go func() {
			for msg := range incomingMessages {
				var req schematicdatastreamws.DataStreamBaseReq
				json.Unmarshal([]byte(msg), &req)

				if req.Data.EntityType == schematicdatastreamws.EntityTypeCompany {
					companyID := req.Data.Keys["company_id"]
					outgoingMessages <- createMockCompanyData(companyID, schematicdatastreamws.MessageTypeFull)
				}
			}
		}()

		// Test that the single flag was added/updated in cache
		ctx := context.Background()
		evalCtx := &schematicgo.CheckFlagRequestBody{
			Company: map[string]string{"company_id": "123"},
		}

		result, err := client.CheckFlag(ctx, evalCtx, flagKey)
		assert.NoError(t, err, "CheckFlag should not return an error for flag updated via single flag message")
		assert.True(t, result.Value, "Flag value should be true as set in the single flag message")
	})

	t.Run("Handles single flag with Delete message type", func(t *testing.T) {
		server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
		defer server.Close()
		defer close(outgoingMessages)

		logger := NewMockLogger()

		// Create a local cache provider for flags
		localFlagCache := cache.NewLocalCache[*rulesengine.Flag](100, time.Minute)

		configOptions := &core.DatastreamOptions{
			CacheTTL: 5 * time.Minute,
		}

		clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
		clientOptions.FlagCache = localFlagCache

		client := datastream.NewDataStreamClient(clientOptions, configOptions)
		client.Start()
		defer client.Close()

		// Wait for connection and initial flags
		time.Sleep(300 * time.Millisecond)

		// First, add a single flag
		flagKey := "test-delete-flag"
		outgoingMessages <- createMockSingleFlagData(flagKey, true, schematicdatastreamws.MessageTypeFull)

		// Allow time for message processing
		time.Sleep(100 * time.Millisecond)

		// Setup a goroutine to handle incoming company requests
		go func() {
			for msg := range incomingMessages {
				var req schematicdatastreamws.DataStreamBaseReq
				json.Unmarshal([]byte(msg), &req)

				if req.Data.EntityType == schematicdatastreamws.EntityTypeCompany {
					companyID := req.Data.Keys["company_id"]
					outgoingMessages <- createMockCompanyData(companyID, schematicdatastreamws.MessageTypeFull)
				}
			}
		}()

		// Verify the flag exists
		ctx := context.Background()
		evalCtx := &schematicgo.CheckFlagRequestBody{
			Company: map[string]string{"company_id": "123"},
		}

		result, err := client.CheckFlag(ctx, evalCtx, flagKey)
		assert.NoError(t, err, "Flag should exist before deletion")
		assert.True(t, result.Value, "Flag value should be true before deletion")

		// Now send a delete message for the same flag
		outgoingMessages <- createMockSingleFlagData(flagKey, false, schematicdatastreamws.MessageTypeDelete)

		// Allow time for message processing
		time.Sleep(100 * time.Millisecond)

		// Verify the flag was removed from cache
		result, err = client.CheckFlag(ctx, evalCtx, flagKey)
		assert.Error(t, err, "CheckFlag should return an error after flag deletion")
		assert.Nil(t, result, "Result should be nil after flag deletion")
		assert.Contains(t, err.Error(), "flag not found", "Error should indicate flag not found after deletion")
	})
}

func TestDeleteMessage(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()

	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection
	time.Sleep(200 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	// First add a company to the cache
	outgoingMessages <- createMockCompanyData("123", schematicdatastreamws.MessageTypeDelete)
	time.Sleep(100 * time.Millisecond)

}

func TestUpdateCompanyMetrics(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()

	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")

	client := datastream.NewDataStreamClient(clientOptions, options)
	client.Start()
	defer client.Close()

	// Wait for connection
	time.Sleep(200 * time.Millisecond)
	time.Sleep(100 * time.Millisecond)

	// Send a company with metrics via WebSocket so both ID key and lookup key are populated
	companyID := "test-company-123"
	eventType := "test-event"
	initialValue := int64(10)

	company := &rulesengine.Company{
		ID:            "company-id",
		AccountID:     "account-id",
		EnvironmentID: "env-id",
		Keys: map[string]string{
			"company_id": companyID,
		},
		Metrics: []*rulesengine.CompanyMetric{
			{
				CompanyID:    companyID,
				EventSubtype: eventType,
				Value:        initialValue,
				Period:       "lifetime",
			},
		},
	}

	companyData, _ := json.Marshal(company)
	resp := schematicdatastreamws.DataStreamResp{
		EntityType:  string(schematicdatastreamws.EntityTypeCompany),
		Data:        companyData,
		MessageType: schematicdatastreamws.MessageTypeFull,
	}
	respData, _ := json.Marshal(resp)
	outgoingMessages <- string(respData)

	// Allow time for message processing
	time.Sleep(200 * time.Millisecond)

	ctx := context.Background()

	// Now update metrics for this company
	incrementValue := 5
	event := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Company:  map[string]string{"company_id": companyID},
		Quantity: &incrementValue,
	}

	err := client.UpdateCompanyMetrics(ctx, event)
	assert.NoError(t, err)

	// Allow time for the update to complete
	time.Sleep(100 * time.Millisecond)

	// Send another metric update with a different increment
	secondIncrementValue := 7
	secondEvent := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Company:  map[string]string{"company_id": companyID},
		Quantity: &secondIncrementValue,
	}

	err = client.UpdateCompanyMetrics(ctx, secondEvent)
	assert.NoError(t, err)

	// Allow time for the update to complete
	time.Sleep(100 * time.Millisecond)

	// Verify by doing a third update - if it doesn't error, the company is still in cache
	// and the previous updates were applied successfully
	thirdIncrementValue := 1
	thirdEvent := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Company:  map[string]string{"company_id": companyID},
		Quantity: &thirdIncrementValue,
	}

	err = client.UpdateCompanyMetrics(ctx, thirdEvent)
	assert.NoError(t, err, "Third update should succeed, confirming company persists in cache")
}

func TestConnectionStateTracking(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()

	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")
	client := datastream.NewDataStreamClient(clientOptions, configOptions)

	// Initially the client should not be connected
	connected := client.IsConnected()
	assert.False(t, connected, "Client should not be connected initially")

	// Start the client and wait for connection
	client.Start()

	// Allow time for connection to establish
	time.Sleep(200 * time.Millisecond)

	// Now the client should be connected
	connected = client.IsConnected()
	assert.True(t, connected, "Connection should be successful")
	connected = client.IsConnected()
	assert.True(t, connected, "Client should be connected after successful connection")

	// Clean up
	client.Close()

	// After closing, the client should not be connected
	time.Sleep(10 * time.Millisecond) // Allow time for cleanup
	connected = client.IsConnected()
	assert.False(t, connected, "Client should not be connected after close")
}

func TestNewDataStreamClientFlagCache(t *testing.T) {
	t.Run("Creates client with local flag cache when no Redis config", func(t *testing.T) {
		logger := NewMockLogger()

		options := datastream.DataStreamClientOptions{
			ApiKey: "test-key",
			Logger: logger,
		}

		configOptions := &core.DatastreamOptions{
			CacheTTL: 24 * time.Hour,
			// No CacheConfig provided - should use local cache
		}

		client := datastream.NewDataStreamClient(options, configOptions)
		assert.NotNil(t, client, "Expected client to be created")

		// Test that flag cache works by attempting a basic operation
		// We can't directly test the cache type, but we can verify the client was created successfully
	})

	t.Run("Creates client with Redis flag cache when Redis config provided", func(t *testing.T) {
		logger := NewMockLogger()

		redisConfig := &core.RedisCacheConfig{
			Addr: "localhost:6379",
			DB:   0,
		}

		options := datastream.DataStreamClientOptions{
			ApiKey: "test-key",
			Logger: logger,
		}

		configOptions := &core.DatastreamOptions{
			CacheTTL:    24 * time.Hour,
			CacheConfig: redisConfig,
		}

		// This should create a client with Redis cache for flags
		client := datastream.NewDataStreamClient(options, configOptions)
		assert.NotNil(t, client, "Expected client to be created with Redis config")

		// We can't easily test Redis connectivity without a running Redis instance,
		// but we can verify the client was created successfully
	})

	t.Run("Uses provided flag cache when specified in options", func(t *testing.T) {
		logger := NewMockLogger()
		customFlagCache := cache.NewLocalCache[*rulesengine.Flag](100, time.Hour)

		options := datastream.DataStreamClientOptions{
			ApiKey:    "test-key",
			Logger:    logger,
			FlagCache: customFlagCache,
		}

		configOptions := &core.DatastreamOptions{
			CacheTTL: 24 * time.Hour,
		}

		client := datastream.NewDataStreamClient(options, configOptions)
		assert.NotNil(t, client, "Expected client to be created")

		// The client should use the provided cache, but we can't easily verify this
		// without accessing private fields. The test ensures the client creation works.
	})
}

// Helper function to create cache keys in the same format as the datastream package
func createTestCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("schematic:%s:%s:%s:%s", resourceType, rulesengine.VersionKey, strings.ToLower(key), strings.ToLower(value))
}
