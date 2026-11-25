package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func whoAmIResponse(t *testing.T, status int) *http.Response {
	t.Helper()
	data, err := json.Marshal(&schematicgo.GetWhoAmIResponse{
		Data: &schematicgo.WhoAmIResponseData{},
	})
	require.NoError(t, err)
	return &http.Response{
		Status:     http.StatusText(status),
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(bytes.NewReader(data)),
	}
}

func countingWhoAmIClient(t *testing.T, status int, whoAmICalls *atomic.Int32) *mocks.MockHTTPClient {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		if strings.HasSuffix(req.URL.Path, "/whoami") {
			whoAmICalls.Add(1)
			return whoAmIResponse(t, status), nil
		}
		return whoAmIResponse(t, http.StatusOK), nil
	}).AnyTimes()
	return mockHTTPClient
}

func newReadinessReplicatorClient(t *testing.T, ready bool, whoAmICalls *atomic.Int32) *schematicclient.SchematicClient {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	var polls atomic.Int32
	health := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer polls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"ready": ready, "cache_version": "v1"})
	}))
	t.Cleanup(health.Close)

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(countingWhoAmIClient(t, http.StatusOK, whoAmICalls)),
		option.WithDatastream(
			option.WithReplicatorMode(),
			option.WithReplicatorHealthURL(health.URL),
			option.WithReplicatorHealthInterval(10*time.Millisecond),
			option.WithRedisClient(rdb),
		),
	)
	t.Cleanup(client.Close)

	require.Eventually(t, func() bool { return polls.Load() >= 2 }, 2*time.Second, 5*time.Millisecond)
	return client
}

func TestIsReadyOffline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithOfflineMode(), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.True(t, client.IsReady(context.Background()))
}

func TestIsReadyReplicatorReadyMakesNoNetworkCall(t *testing.T) {
	var whoAmICalls atomic.Int32
	client := newReadinessReplicatorClient(t, true, &whoAmICalls)

	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, int32(0), whoAmICalls.Load())
}

func TestIsReadyReplicatorNotReadyCallsWhoAmI(t *testing.T) {
	var whoAmICalls atomic.Int32
	client := newReadinessReplicatorClient(t, false, &whoAmICalls)

	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, int32(1), whoAmICalls.Load())
}

func TestIsReadyWebSocketFlagsLoadedMakesNoNetworkCall(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true})
	var whoAmICalls atomic.Int32
	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(countingWhoAmIClient(t, http.StatusInternalServerError, &whoAmICalls)),
		option.WithBaseURL(server.URL),
		option.WithDatastream(),
	)
	t.Cleanup(client.Close)

	require.Eventually(t, func() bool { return client.IsReady(context.Background()) }, 3*time.Second, 10*time.Millisecond)

	before := whoAmICalls.Load()
	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, before, whoAmICalls.Load())
}

func TestIsReadyWebSocketFlagsNotLoadedCallsWhoAmI(t *testing.T) {
	var connections atomic.Int32
	server := startDataStreamServer(t, dataStreamFixture{connections: &connections})
	var whoAmICalls atomic.Int32
	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(countingWhoAmIClient(t, http.StatusOK, &whoAmICalls)),
		option.WithBaseURL(server.URL),
		option.WithDatastream(),
	)
	t.Cleanup(client.Close)

	require.Eventually(t, func() bool { return connections.Load() >= 1 }, 2*time.Second, 5*time.Millisecond)

	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, int32(1), whoAmICalls.Load())
}

func TestIsReadyWebSocketNotConnectedCallsWhoAmI(t *testing.T) {
	closed := httptest.NewServer(http.NotFoundHandler())
	closed.Close()

	var whoAmICalls atomic.Int32
	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(countingWhoAmIClient(t, http.StatusOK, &whoAmICalls)),
		option.WithBaseURL(closed.URL),
		option.WithDatastream(),
	)
	t.Cleanup(client.Close)

	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, int32(1), whoAmICalls.Load())
}

func TestIsReadyAPISuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.True(t, strings.HasSuffix(req.URL.Path, "/whoami"), req.URL.Path)
		return whoAmIResponse(t, http.StatusOK), nil
	}).Times(1)

	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.True(t, client.IsReady(context.Background()))
}

func TestIsReadyAPIFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(whoAmIResponse(t, http.StatusUnauthorized), nil).Times(1)

	client := schematicclient.NewSchematicClient(option.WithAPIKey("bad-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.False(t, client.IsReady(context.Background()))
}

func TestIsReadyAPIServerErrorIsNotRetried(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(whoAmIResponse(t, http.StatusServiceUnavailable), nil).Times(1)

	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.False(t, client.IsReady(context.Background()))
}

func TestIsReadyAPIContextExpired(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		return nil, req.Context().Err()
	}).AnyTimes()

	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	defer cancel()

	assert.False(t, client.IsReady(ctx))
}

func TestIsReadyAPIRequestHasFiveSecondDeadline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		deadline, ok := req.Context().Deadline()
		require.True(t, ok)
		assert.InDelta(t, float64(5*time.Second), float64(time.Until(deadline)), float64(time.Second))
		return whoAmIResponse(t, http.StatusOK), nil
	}).Times(1)

	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.True(t, client.IsReady(context.Background()))
}

func TestIsReadyAPIShorterCallerDeadlineWins(t *testing.T) {
	callerDeadline := time.Now().Add(time.Second)
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		deadline, ok := req.Context().Deadline()
		require.True(t, ok)
		assert.True(t, deadline.Equal(callerDeadline))
		return whoAmIResponse(t, http.StatusOK), nil
	}).Times(1)

	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	ctx, cancel := context.WithDeadline(context.Background(), callerDeadline)
	defer cancel()

	assert.True(t, client.IsReady(ctx))
}
