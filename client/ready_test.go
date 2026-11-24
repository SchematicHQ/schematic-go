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

func TestIsReadyOffline(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithOfflineMode(), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.True(t, client.IsReady(context.Background()))
}

func TestIsReadyDataStreamCacheMakesNoNetworkCall(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	var polls atomic.Int32
	health := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer polls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"ready": true, "cache_version": "v1"})
	}))
	t.Cleanup(health.Close)

	var whoAmICalls atomic.Int32
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		if strings.HasSuffix(req.URL.Path, "/whoami") {
			whoAmICalls.Add(1)
		}
		return whoAmIResponse(t, http.StatusOK), nil
	}).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithDatastream(
			option.WithReplicatorMode(),
			option.WithReplicatorHealthURL(health.URL),
			option.WithReplicatorHealthInterval(10*time.Millisecond),
			option.WithRedisClient(rdb),
		),
	)
	t.Cleanup(client.Close)

	require.Eventually(t, func() bool { return polls.Load() >= 2 }, 2*time.Second, 5*time.Millisecond)

	assert.True(t, client.IsReady(context.Background()))
	assert.Equal(t, int32(0), whoAmICalls.Load())
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
