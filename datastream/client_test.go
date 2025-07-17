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
	resp := datastream.DataStreamResp{
		EntityType: string(datastream.EntityTypeFlags),
		Data:       flagsData,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function to create mock company data
func createMockCompanyData(companyID string, messageType datastream.MessageType) string {
	company := &rulesengine.Company{
		Keys: map[string]string{
			"company_id": companyID,
		},
	}

	companyData, _ := json.Marshal(company)
	resp := datastream.DataStreamResp{
		EntityType:  string(datastream.EntityTypeCompany),
		Data:        companyData,
		MessageType: messageType,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function to create mock user data
func createMockUserData(userID string, messageType datastream.MessageType) string {
	user := &rulesengine.User{
		Keys: map[string]string{
			"user_id": userID,
		},
	}

	userData, _ := json.Marshal(user)
	resp := datastream.DataStreamResp{
		EntityType:  string(datastream.EntityTypeUser),
		Data:        userData,
		MessageType: messageType,
	}
	respData, _ := json.Marshal(resp)

	return string(respData)
}

// Helper function for creating DataStreamClientOptions for tests
func createTestClientOptions(baseURL string, logger core.Logger, apiKey string, monitorChannel chan bool) datastream.DataStreamClientOptions {
	return datastream.DataStreamClientOptions{
		ApiKey:         apiKey,
		BaseURL:        baseURL,
		Logger:         logger,
		MonitorChannel: monitorChannel,
	}
}

func TestNewDataStreamClient(t *testing.T) {
	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions("http://localhost:8080", logger, "test-api-key", monitorChannel)

	client := datastream.NewDataStreamClient(clientOptions, configOptions)

	assert.NotNil(t, client, "Client should not be nil")
}

func TestDataStreamClientConnection(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key", monitorChannel)
	client := datastream.NewDataStreamClient(clientOptions, configOptions)

	// Start the client and wait for connection
	client.Start()

	// Wait for monitor channel to receive connection status
	status := <-monitorChannel
	assert.True(t, status, "Connection should be successful")

	// Allow time for connection to establish
	time.Sleep(100 * time.Millisecond)

	// Clean up
	client.Close()
}

func TestCheckFlagCompany(t *testing.T) {
	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key", monitorChannel)
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection and initial flags to be received
	<-monitorChannel
	time.Sleep(100 * time.Millisecond)

	// Setup a goroutine to handle incoming company and user requests
	go func() {
		for msg := range incomingMessages {
			var req datastream.DataStreamBaseReq
			json.Unmarshal([]byte(msg), &req)

			if req.Data.EntityType == datastream.EntityTypeCompany {
				companyID := req.Data.Keys["company_id"]
				outgoingMessages <- createMockCompanyData(companyID, datastream.MessageTypFull)
			}
		}
	}()

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}

	// Test checking a flag that exists and is true
	result := client.CheckFlag(ctx, evalCtx, "test-flag-1")
	assert.True(t, result.Value, "Flag value should be true")

	// Test checking a flag that exists and is false
	result = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.False(t, result.Value, "Flag value should be false")

	// Test checking a non-existent flag
	result = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.Nil(t, result, "Non-existent flag should return nil flag check result")
}

func TestCheckFlagUser(t *testing.T) {
	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key", monitorChannel)
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection and initial flags to be received
	<-monitorChannel
	time.Sleep(100 * time.Millisecond)

	// Setup a goroutine to handle incoming company and user requests
	go func() {
		for msg := range incomingMessages {
			var req datastream.DataStreamBaseReq
			json.Unmarshal([]byte(msg), &req)

			if req.Data.EntityType == datastream.EntityTypeUser {
				companyID := req.Data.Keys["user_id"]
				outgoingMessages <- createMockUserData(companyID, datastream.MessageTypFull)
			}
		}
	}()

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		User: map[string]string{"user_id": "123"},
	}

	// Test checking a flag that exists and is true
	result := client.CheckFlag(ctx, evalCtx, "test-flag-1")
	assert.True(t, result.Value, "Flag value should be true")

	// Test checking a flag that exists and is false
	result = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.False(t, result.Value, "Flag value should be false")

	// Test checking a non-existent flag
	result = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.Nil(t, result, "Non-existent flag should return nil flag check result")
}

func TestDeleteMessage(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	configOptions := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key", monitorChannel)
	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	defer client.Close()

	// Wait for connection
	<-monitorChannel
	time.Sleep(100 * time.Millisecond)

	// First add a company to the cache
	outgoingMessages <- createMockCompanyData("123", datastream.MessageTypeDelete)
	time.Sleep(100 * time.Millisecond)

}

func TestUpdateCompanyMetrics(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)

	// Create a local cache provider for companies with a 1 minute TTL
	localCompanyCache := cache.NewLocalCache[*rulesengine.Company](100, time.Minute)

	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	// Create client options with the local cache provider
	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key", monitorChannel)
	// Add the cache to the options
	clientOptions.CompanyCache = localCompanyCache

	client := datastream.NewDataStreamClient(clientOptions, options)
	client.Start()
	defer client.Close()

	// Wait for connection
	<-monitorChannel
	time.Sleep(100 * time.Millisecond)

	// Create a mock company with metrics
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
				Period:       "lifetime", // Set period to ensure proper metric matching
			},
		},
	}

	// Add the company to the cache using each key in the company
	ctx := context.Background()
	for key, value := range company.Keys {
		cacheKey := createTestCacheKey("company", key, value)
		localCompanyCache.Set(ctx, cacheKey, company, nil)
	}

	// Verify the company was added to the cache
	cacheKey := createTestCacheKey("company", "company_id", companyID)
	cachedCompany, exists := localCompanyCache.Get(ctx, cacheKey)
	assert.True(t, exists, "Company should exist in cache")
	assert.NotNil(t, cachedCompany, "Cached company should not be nil")

	// Now update metrics for this company
	incrementValue := 5
	event := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Company:  map[string]string{"company_id": companyID},
		Quantity: &incrementValue,
	}

	// Set the company field in the event
	event.Company = map[string]string{"company_id": companyID}
	err := client.UpdateCompanyMetrics(ctx, event)
	assert.NoError(t, err)

	// Allow time for the update to complete
	time.Sleep(100 * time.Millisecond)

	// Now we can directly check the cache to verify the update
	updatedCompany, exists := localCompanyCache.Get(ctx, cacheKey)
	assert.True(t, exists, "Company should still exist in cache")
	assert.NotNil(t, updatedCompany, "Updated company should not be nil")

	// Find the metric and verify its value
	var foundMetric *rulesengine.CompanyMetric
	for _, metric := range updatedCompany.Metrics {
		if metric != nil && metric.EventSubtype == eventType {
			foundMetric = metric
			break
		}
	}

	assert.NotNil(t, foundMetric, "Metric should exist for event type %s", eventType)
	expectedValue := initialValue + int64(incrementValue)
	assert.Equal(t, expectedValue, foundMetric.Value, "Metric value should be incremented correctly from %d to %d",
		initialValue, expectedValue)

	// Make a second update with a different increment
	secondIncrementValue := 7
	secondEvent := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Company:  map[string]string{"company_id": companyID},
		Quantity: &secondIncrementValue,
	}

	err = client.UpdateCompanyMetrics(ctx, secondEvent)
	assert.NoError(t, err)

	// Allow more time for the update to complete
	time.Sleep(100 * time.Millisecond)

	// Check the cache again for the final update
	finalCompany, exists := localCompanyCache.Get(ctx, cacheKey)
	assert.True(t, exists, "Company should still exist in cache after second update")
	assert.NotNil(t, finalCompany, "Final company should not be nil")

	// Find the metric again and verify its final value
	var finalMetric *rulesengine.CompanyMetric
	for _, metric := range finalCompany.Metrics {
		if metric != nil && metric.EventSubtype == eventType {
			finalMetric = metric
			break
		}
	}

	assert.NotNil(t, finalMetric, "Metric should exist after second update")
	finalExpectedValue := initialValue + int64(incrementValue) + int64(secondIncrementValue)
	assert.Equal(t, finalExpectedValue, finalMetric.Value,
		"Metric value should include both increments (initial %d + first %d + second %d = %d)",
		initialValue, incrementValue, secondIncrementValue, finalExpectedValue)
}

// Helper function to create cache keys in the same format as the datastream package
func createTestCacheKey(resourceType string, key string, value string) string {
	return fmt.Sprintf("schematic:%s:%s:%s:%s", resourceType, datastream.RulesEngineVersionKey, strings.ToLower(key), strings.ToLower(value))
}
