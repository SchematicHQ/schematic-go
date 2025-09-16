package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	http "net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"

	"github.com/stretchr/testify/assert"
)

func TestNewSchematicClient(t *testing.T) {
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"))
	defer client.Close()

	assert.NotNil(t, client)
}

func TestCheckFlagWithCacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.NotNil(t, client)

	responseBody := &schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Value: true,
		},
	}

	data, err := json.Marshal(responseBody)

	assert.Nil(t, err)
	body := bytes.NewReader(data)

	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(body),
	}, nil).Times(1)

	// First call should hit the API and set the cache
	resp := client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag")
	assert.True(t, resp)
	// Wait enough time for the cache to be set (since this happens in a goroutine)
	time.Sleep(10 * time.Millisecond)
	// Second call should hit the cache
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagWithNoCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithDisableFlagCheckCache(), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.NotNil(t, client)

	responseBody := &schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Value: true,
		},
	}

	for i := 0; i < 2; i++ {
		data, err := json.Marshal(responseBody)
		assert.Nil(t, err)
		mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
			Status:     "200",
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(data)),
		}, nil)
	}

	// Both calls should hit the API
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time in between calls for the cache to be set, were it to be active (since this happens in a goroutine)
	time.Sleep(10 * time.Millisecond)
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagWithCacheOptions(t *testing.T) {
	cacheTTL := 50 * time.Millisecond
	cacheMaxSize := 1000
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithLocalFlagCheckCache(cacheMaxSize, cacheTTL), option.WithHTTPClient(mockHTTPClient))
	defer client.Close()

	assert.NotNil(t, client)

	responseBody := &schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Value: true,
		},
	}

	for i := 0; i < 2; i++ {
		data, err := json.Marshal(responseBody)
		assert.Nil(t, err)
		mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
			Status:     "200",
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(data)),
		}, nil)
	}

	// First call should hit the API and set the cache
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time for the cache to finish being set in the goroutine
	time.Sleep(10 * time.Millisecond)
	// Cache hit
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
	// Wait enough time for the cache to expire
	time.Sleep(50 * time.Millisecond)
	// Cache miss
	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestTrackEventBatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithEventBufferPeriod(10*time.Millisecond), option.WithHTTPClient(mockHTTPClient))
	assert.NotNil(t, client)
	defer client.Close()

	// expect a single call to the API for the batch
	responseBody := &schematicgo.CreateEventBatchResponse{
		Data: &schematicgo.RawEventBatchResponseData{
			Events: []*schematicgo.RawEventResponseData{},
		},
	}
	data, err := json.Marshal(responseBody)
	assert.Nil(t, err)
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(data)),
	}, nil)

	ctx := context.Background()
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "foo", Company: map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "bar", Company: map[string]string{"foo": "bar"}})
	time.Sleep(20 * time.Millisecond)
}

func TestTrackEventWithQuantity(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithEventBufferPeriod(10*time.Millisecond), option.WithHTTPClient(mockHTTPClient))
	assert.NotNil(t, client)
	defer client.Close()

	responseBody := &schematicgo.CreateEventBatchResponse{
		Data: &schematicgo.RawEventBatchResponseData{
			Events: []*schematicgo.RawEventResponseData{},
		},
	}
	data, err := json.Marshal(responseBody)
	assert.Nil(t, err)
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(data)),
	}, nil)

	ctx := context.Background()
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "bar", Company: map[string]string{"foo": "bar"}, Quantity: schematicgo.Int(100)})
	time.Sleep(20 * time.Millisecond)
}

func TestMultipleEventBatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithEventBufferPeriod(10*time.Millisecond), option.WithHTTPClient(mockHTTPClient))
	assert.NotNil(t, client)
	defer client.Close()

	responseBody := &schematicgo.CreateEventBatchResponse{
		Data: &schematicgo.RawEventBatchResponseData{
			Events: []*schematicgo.RawEventResponseData{},
		},
	}

	for i := 0; i < 2; i++ {
		data, err := json.Marshal(responseBody)
		assert.Nil(t, err)
		mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
			Status:     "200",
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(data)),
		}, nil)
	}

	ctx := context.Background()
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "foo", Company: map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "bar", Company: map[string]string{"foo": "bar"}})
	// Wait for event buffer to flush and start submitting second batch
	time.Sleep(20 * time.Millisecond)
	client.Identify(ctx, &schematicgo.EventBodyIdentify{Keys: map[string]string{"foo": "bar"}})
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "baz", Company: map[string]string{"foo": "bar"}})
	// Wait for event buffer to flush
	time.Sleep(20 * time.Millisecond)
}

func TestCheckFlagOfflineModeNoAPIKeyOption(t *testing.T) {
	// Passing mock HTTP client, which we expect to be replaced by a NoopClient internally. This test will fail if any calls are made to the mock client.
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithHTTPClient(mockHTTPClient))
	client.SetFlagDefault("test-flag", true)
	defer client.Close()

	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagOfflineModeEmptyAPIKey(t *testing.T) {
	// Passing mock HTTP client, which we expect to be replaced by a NoopClient internally. This test will fail if any calls are made to the mock client.
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithAPIKey(""), option.WithHTTPClient(mockHTTPClient))
	client.SetFlagDefault("test-flag", true)
	defer client.Close()

	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagOfflineModeOption(t *testing.T) {
	// Passing mock HTTP client, which we expect to be replaced by a NoopClient internally. This test will fail if any calls are made to the mock client.
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	client := schematicclient.NewSchematicClient(option.WithOfflineMode(), option.WithHTTPClient(mockHTTPClient))
	client.SetFlagDefault("test-flag", true)
	defer client.Close()

	assert.True(t, client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "test-flag"))
}

func TestCheckFlagDatastreamFallbackToAPI(t *testing.T) {
	// Test that when datastream is enabled but fails/is not connected,
	// the client falls back to the API
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)

	// Set up mock API response for the fallback
	responseBody := &schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Value: true,
		},
	}

	data, err := json.Marshal(responseBody)
	assert.Nil(t, err)
	body := bytes.NewReader(data)

	// Expect the API call when datastream fails
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(body),
	}, nil)

	// Create client with datastream enabled but with invalid base URL to force failure
	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithBaseURL("https://invalid.example.com"), // This will cause datastream to fail
		option.WithDatastream(),
	)
	defer client.Close()

	// Give some time for datastream connection attempt to fail
	time.Sleep(50 * time.Millisecond)

	// Check flag - should fallback to API since datastream failed
	result := client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": "test-company"},
	}, "test-flag")

	assert.True(t, result)
}
