/*
Schematic API

Testing FeaturesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package schematic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_schematic_FeaturesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeaturesApiService CheckFlag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var key string

		resp, httpRes, err := apiClient.FeaturesApi.CheckFlag(context.Background(), key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CheckFlags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CheckFlags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CountCompaniesAudience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CountCompaniesAudience(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CountFlagChecks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CountFlagChecks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CountUsersAudience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CountUsersAudience(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CreateFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CreateFeature(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CreateFlag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CreateFlag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService CreateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.CreateRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService DeleteFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesApi.DeleteFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService DeleteFlag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesApi.DeleteFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetCompaniesAudience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.GetCompaniesAudience(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesApi.GetFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetFlag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesApi.GetFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetFlagCheck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var flagCheckId string

		resp, httpRes, err := apiClient.FeaturesApi.GetFlagCheck(context.Background(), flagCheckId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ruleId string

		resp, httpRes, err := apiClient.FeaturesApi.GetRule(context.Background(), ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService GetUsersAudience", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.GetUsersAudience(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService LatestFlagChecks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.LatestFlagChecks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService ListFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.ListFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService ListFlagChecks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.ListFlagChecks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService ListFlags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FeaturesApi.ListFlags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService UpdateFeature", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var featureId string

		resp, httpRes, err := apiClient.FeaturesApi.UpdateFeature(context.Background(), featureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService UpdateFlag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var flagId string

		resp, httpRes, err := apiClient.FeaturesApi.UpdateFlag(context.Background(), flagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeaturesApiService UpdateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ruleId string

		resp, httpRes, err := apiClient.FeaturesApi.UpdateRule(context.Background(), ruleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
