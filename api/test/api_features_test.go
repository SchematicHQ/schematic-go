/*
Schematic API

Testing FeaturesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_FeaturesAPIService(t *testing.T) {

	configuration := schematicapi.NewConfiguration()
	apiClient := schematicapi.NewAPIClient(configuration)

	t.Run("Test FeaturesAPIService CheckFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var key string

		resp, httpRes, err := apiClient.FeaturesAPI.CheckFlag(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CheckFlags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CheckFlags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CountAudienceCompanies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CountAudienceCompanies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CountAudienceUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CountAudienceUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CountFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CountFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CountFlags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CountFlags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CreateFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CreateFeature(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService CreateFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.CreateFlag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService DeleteFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesAPI.DeleteFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService DeleteFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesAPI.DeleteFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService GetFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesAPI.GetFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService GetFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesAPI.GetFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService ListAudienceCompanies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.ListAudienceCompanies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService ListAudienceUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.ListAudienceUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService ListFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.ListFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService ListFlags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeaturesAPI.ListFlags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService UpdateFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesAPI.UpdateFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService UpdateFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesAPI.UpdateFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesAPIService UpdateFlagRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesAPI.UpdateFlagRules(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
