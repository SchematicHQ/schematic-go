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
	assert.True(t, result, "Flag value should be true")

	// Test checking a flag that exists and is false
	result = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.False(t, result, "Flag value should be false")

	// Test checking a non-existent flag
	result = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.False(t, result, "Non-existent flag should return false")
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
	assert.True(t, result, "Flag value should be true")

	// Test checking a flag that exists and is false
	result = client.CheckFlag(ctx, evalCtx, "test-flag-2")
	assert.False(t, result, "Flag value should be false")

	// Test checking a non-existent flag
	result = client.CheckFlag(ctx, evalCtx, "non-existent-flag")
	assert.False(t, result, "Non-existent flag should return false")
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
