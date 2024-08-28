/*
Schematic API

Testing PlansAPIService

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

func Test_api_PlansAPIService(t *testing.T) {

	configuration := schematicapi.NewConfiguration()
	apiClient := schematicapi.NewAPIClient(configuration)

	t.Run("Test PlansAPIService CountPlans", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PlansAPI.CountPlans(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService CreatePlan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PlansAPI.CreatePlan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService DeleteAudience", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planAudienceId string

		resp, httpRes, err := apiClient.PlansAPI.DeleteAudience(context.Background(), planAudienceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService DeletePlan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planId string

		resp, httpRes, err := apiClient.PlansAPI.DeletePlan(context.Background(), planId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService GetAudience", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planAudienceId string

		resp, httpRes, err := apiClient.PlansAPI.GetAudience(context.Background(), planAudienceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService GetPlan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planId string

		resp, httpRes, err := apiClient.PlansAPI.GetPlan(context.Background(), planId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService ListPlans", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PlansAPI.ListPlans(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService UpdateAudience", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planAudienceId string

		resp, httpRes, err := apiClient.PlansAPI.UpdateAudience(context.Background(), planAudienceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService UpdatePlan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planId string

		resp, httpRes, err := apiClient.PlansAPI.UpdatePlan(context.Background(), planId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService UpsertBillingProductPlan", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var planId string

		resp, httpRes, err := apiClient.PlansAPI.UpsertBillingProductPlan(context.Background(), planId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}