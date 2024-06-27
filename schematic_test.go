package schematic_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/SchematicHQ/schematic-go"
	"github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewClient(t *testing.T) {
	client := schematic.NewClient("test-api-key")
	defer client.Close()
	assert.NotNil(t, client)
	assert.NotNil(t, client.API())
}

func TestCheckFlagWithCacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiClient, mocks := mocks.NewMockAPIClient(ctrl)
	client := schematic.NewClient("test-api-key")
	client.SetAPIClient(apiClient)
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, client.API())

	mocks.FeaturesAPI.EXPECT().CheckFlag(gomock.Any(), "test-flag").Return(api.ApiCheckFlagRequest{ApiService: mocks.FeaturesAPI}).Times(1)
	mocks.FeaturesAPI.EXPECT().CheckFlagExecute(gomock.Any()).Return(&api.CheckFlagResponse{
		Data: api.CheckFlagResponseData{
			Value: true,
		},
	}, nil, nil).Times(1)

	// First call should hit the API and set the cache
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time for the cache to be set (since this happens in a goroutine)
	time.Sleep(10 * time.Millisecond)
	// Second call should hit the cache
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagWithNoCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiClient, mocks := mocks.NewMockAPIClient(ctrl)
	client := schematic.NewClient("test-api-key", schematic.WithDisableFlagCheckCache())
	client.SetAPIClient(apiClient)
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, client.API())

	mocks.FeaturesAPI.EXPECT().CheckFlag(gomock.Any(), "test-flag").Return(api.ApiCheckFlagRequest{ApiService: mocks.FeaturesAPI}).Times(2)
	mocks.FeaturesAPI.EXPECT().CheckFlagExecute(gomock.Any()).Return(&api.CheckFlagResponse{
		Data: api.CheckFlagResponseData{
			Value: true,
		},
	}, nil, nil).Times(2)

	// Both calls should hit the API
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time in between calls for the cache to be set, were it to be active (since this happens in a goroutine)
	time.Sleep(10 * time.Millisecond)
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagWithCacheOptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiClient, mocks := mocks.NewMockAPIClient(ctrl)
	cacheTTL := 50 * time.Millisecond
	cacheMaxSize := 1000
	client := schematic.NewClient("test-api-key", schematic.WithLocalFlagCheckCache(cacheMaxSize, cacheTTL))
	client.SetAPIClient(apiClient)
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, client.API())

	mocks.FeaturesAPI.EXPECT().CheckFlag(gomock.Any(), "test-flag").Return(api.ApiCheckFlagRequest{ApiService: mocks.FeaturesAPI}).Times(2)
	mocks.FeaturesAPI.EXPECT().CheckFlagExecute(gomock.Any()).Return(&api.CheckFlagResponse{
		Data: api.CheckFlagResponseData{
			Value: true,
		},
	}, nil, nil).Times(2)

	// First call should hit the API and set the cache
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time for the cache to finish being set in the goroutine
	time.Sleep(10 * time.Millisecond)
	// Cache hit
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time for the cache to expire
	time.Sleep(50 * time.Millisecond)
	// Cache miss
	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
}

func TestTrackEventBatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiClient, mocks := mocks.NewMockAPIClient(ctrl)
	client := schematic.NewClient("test-api-key", schematic.WithEventBufferPeriod(10*time.Millisecond))
	client.SetAPIClient(apiClient)
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, client.API())

	mocks.EventsAPI.EXPECT().CreateEventBatch(gomock.Any()).Return(api.ApiCreateEventBatchRequest{ApiService: mocks.EventsAPI}).Times(1)
	mocks.EventsAPI.EXPECT().CreateEventBatchExecute(gomock.Any()).Return(nil, nil, nil).Times(1)

	ctx := context.Background()
	client.Track(ctx, &schematic.EventBodyTrack{Event: "foo", Company: &map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematic.EventBodyTrack{Event: "bar", Company: &map[string]string{"foo": "bar"}})
	time.Sleep(11 * time.Millisecond)
}

func TestMultipleEventBatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiClient, mocks := mocks.NewMockAPIClient(ctrl)
	client := schematic.NewClient("test-api-key", schematic.WithEventBufferPeriod(10*time.Millisecond))
	client.SetAPIClient(apiClient)
	defer client.Close()

	assert.NotNil(t, client)
	assert.NotNil(t, client.API())

	mocks.EventsAPI.EXPECT().CreateEventBatch(gomock.Any()).Return(api.ApiCreateEventBatchRequest{ApiService: mocks.EventsAPI}).Times(2)
	mocks.EventsAPI.EXPECT().CreateEventBatchExecute(gomock.Any()).Return(nil, nil, nil).Times(2)

	ctx := context.Background()
	client.Track(ctx, &schematic.EventBodyTrack{Event: "foo", Company: &map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematic.EventBodyTrack{Event: "bar", Company: &map[string]string{"foo": "bar"}})
	// Wait for event buffer to flush and start submitting second batch
	time.Sleep(20 * time.Millisecond)
	client.Identify(ctx, &schematic.EventBodyIdentify{Keys: map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematic.EventBodyTrack{Event: "baz", Company: &map[string]string{"foo": "bar"}})
	// Wait for event buffer to flush
	time.Sleep(20 * time.Millisecond)
}

func TestCheckFlagOfflineMode(t *testing.T) {
	client, _ := mocks.NewClientWithMockHTTP("", gomock.NewController(t))
	client.SetFlagDefault("test-flag", true)
	defer client.Close()

	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagLiveMode(t *testing.T) {
	client, mockObjs := mocks.NewClientWithMockHTTP("api_foo", gomock.NewController(t))
	defer client.Close()

	respBody, err := json.Marshal(&api.CheckFlagResponse{
		Data:   api.CheckFlagResponseData{Value: true},
		Params: map[string]any{},
	})
	assert.Nil(t, err)
	mockObjs.HTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		Body:       io.NopCloser(bytes.NewReader(respBody)),
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}, nil).Times(1)

	assert.True(t, client.CheckFlag(context.Background(), &schematic.CheckFlagRequestBody{}, "test-flag"))
}
