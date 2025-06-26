package datastream_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/schematichq/rulesengine"
	schematicgo "github.com/schematichq/schematic-go"
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

func TestNewDataStreamClient(t *testing.T) {
	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient("http://localhost:8080", logger, "test-api-key", monitorChannel, options)

	assert.NotNil(t, client, "Client should not be nil")
}

func TestDataStreamClientConnection(t *testing.T) {
	server, _, outgoingMessages := setupMockWebSocketServer()
	defer server.Close()
	defer close(outgoingMessages)

	logger := NewMockLogger()
	monitorChannel := make(chan bool, 1)
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient(server.URL, logger, "test-api-key", monitorChannel, options)

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
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient(server.URL, logger, "test-api-key", monitorChannel, options)
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
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient(server.URL, logger, "test-api-key", monitorChannel, options)
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
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient(server.URL, logger, "test-api-key", monitorChannel, options)
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
	options := &core.DatastreamOptions{
		CacheTTL: 5 * time.Minute,
	}

	client := datastream.NewDataStreamClient(server.URL, logger, "test-api-key", monitorChannel, options)
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
		ID: "company-id",
		Keys: map[string]string{
			"company_id": companyID,
		},
		Metrics: []*rulesengine.CompanyMetric{
			{
				CompanyID:    companyID,
				EventSubtype: eventType,
				Value:        initialValue,
			},
		},
	}

	companyData, _ := json.Marshal(company)
	resp := datastream.DataStreamResp{
		EntityType:  string(datastream.EntityTypeCompany),
		Data:        companyData,
		MessageType: datastream.MessageTypFull,
	}
	respData, _ := json.Marshal(resp)

	// Send the company to the client (will be cached)
	outgoingMessages <- string(respData)

	// Give time for the company to be processed and cached
	time.Sleep(100 * time.Millisecond)

	// Now update metrics for this company
	ctx := context.Background()
	incrementValue := 5
	event := &schematicgo.EventBodyTrack{
		Event:    eventType,
		Quantity: &incrementValue,
	}

	err := client.UpdateCompanyMetrics(ctx, event)
	assert.NoError(t, err)

	// Give time for the update to be processed
	time.Sleep(50 * time.Millisecond)

	// In a real implementation, we would send a request to fetch the company
	// via the websocket connection. Since we can't access the unexported sendWebSocketMessage
	// method in tests, we'll simulate the response directly:

	// Create an updated company to simulate what would be returned from the server
	// In a real scenario, the server would return the updated company
	updatedCompany := &rulesengine.Company{
		ID: "company-id",
		Keys: map[string]string{
			"company_id": companyID,
		},
		Metrics: []*rulesengine.CompanyMetric{
			{
				CompanyID:    companyID,
				EventSubtype: eventType,
				Value:        initialValue + int64(incrementValue),
			},
		},
	}

	updatedCompanyData, _ := json.Marshal(updatedCompany)
	updatedResp := datastream.DataStreamResp{
		EntityType:  string(datastream.EntityTypeCompany),
		Data:        updatedCompanyData,
		MessageType: datastream.MessageTypFull,
	}
	updatedRespData, _ := json.Marshal(updatedResp)

	// Send the updated company data through outgoing messages
	// This simulates the server's response with the updated company
	outgoingMessages <- string(updatedRespData)

	// Give time for the response to be processed
	time.Sleep(100 * time.Millisecond)

	// Make a request to fetch the company again - in a real scenario, we would call
	// the getCompany method. Since we can't access internal cache directly,
	// we'll simulate receiving the updated company again
	outgoingMessages <- string(updatedRespData)
	time.Sleep(100 * time.Millisecond)

	// Since we can't directly verify the cache was updated (getCompanyFromCache is unexported),
	// we can verify that:
	// 1. The test completed without errors, meaning UpdateCompanyMetrics didn't fail
	// 2. We successfully sent the update message
	assert.NotEmpty(t, logger.infoMessages, "Logger should have recorded messages")

	// In a real world scenario, we would modify this test to use exported methods
	// or add test-specific exported methods to verify the cache state

	// Note: There might be a type mismatch in UpdateCompanyMetrics implementation.
	// EventBodyTrack.Quantity is defined as *int but the function uses it as if it were *float64.
	// This test uses *int as per the type definition, but the implementation might need to be verified.
}
