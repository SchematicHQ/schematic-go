/*
Schematic API

Testing AccesstokensAPIService

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

func Test_api_AccesstokensAPIService(t *testing.T) {

	configuration := schematicapi.NewConfiguration()
	apiClient := schematicapi.NewAPIClient(configuration)

	t.Run("Test AccesstokensAPIService IssueTemporaryAccessToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccesstokensAPI.IssueTemporaryAccessToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
