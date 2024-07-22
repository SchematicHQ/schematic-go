# \AccesstokensAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueTemporaryAccessToken**](AccesstokensAPI.md#IssueTemporaryAccessToken) | **Post** /temporary-access-tokens | Issue temporary access token



## IssueTemporaryAccessToken

> IssueTemporaryAccessTokenResponse IssueTemporaryAccessToken(ctx).IssueTemporaryAccessTokenRequestBody(issueTemporaryAccessTokenRequestBody).Execute()

Issue temporary access token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	issueTemporaryAccessTokenRequestBody := *schematicapi.NewIssueTemporaryAccessTokenRequestBody(map[string]string{"key": "Inner_example"}, "ResourceType_example") // IssueTemporaryAccessTokenRequestBody | 

	resp, r, err := client.API().AccesstokensAPI.IssueTemporaryAccessToken(context.Background()).IssueTemporaryAccessTokenRequestBody(issueTemporaryAccessTokenRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesstokensAPI.IssueTemporaryAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueTemporaryAccessToken`: IssueTemporaryAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AccesstokensAPI.IssueTemporaryAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueTemporaryAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTemporaryAccessTokenRequestBody** | [**IssueTemporaryAccessTokenRequestBody**](IssueTemporaryAccessTokenRequestBody.md) |  | 

### Return type

[**IssueTemporaryAccessTokenResponse**](IssueTemporaryAccessTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

