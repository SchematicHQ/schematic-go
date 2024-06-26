/*
Schematic API

Testing CrmAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"testing"

	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_CrmAPIService(t *testing.T) {

	configuration := schematicapi.NewConfiguration()
	apiClient := schematicapi.NewAPIClient(configuration)

	t.Run("Test CrmAPIService ListCRMProducts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CrmAPI.ListCrmProducts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CrmAPIService UpsertCRMProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CrmAPI.UpsertCrmProduct(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
