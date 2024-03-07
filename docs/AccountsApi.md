# \AccountsApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountApiKeys**](AccountsApi.md#CountApiKeys) | **Get** /api-keys/count | Count api keys
[**CountApiRequests**](AccountsApi.md#CountApiRequests) | **Get** /api-requests/count | Count api requests
[**CreateApiKey**](AccountsApi.md#CreateApiKey) | **Post** /api-keys | Create api key
[**CreateEnvironment**](AccountsApi.md#CreateEnvironment) | **Post** /environments | Create environment
[**DeleteApiKey**](AccountsApi.md#DeleteApiKey) | **Delete** /api-keys/{api_key_id} | Delete api key
[**DeleteEnvironment**](AccountsApi.md#DeleteEnvironment) | **Delete** /environments/{environment_id} | Delete environment
[**GetApiKey**](AccountsApi.md#GetApiKey) | **Get** /api-keys/{api_key_id} | Get api key
[**GetApiRequest**](AccountsApi.md#GetApiRequest) | **Get** /api-requests/{key} | Get api request
[**GetEnvironment**](AccountsApi.md#GetEnvironment) | **Get** /environments/{environment_id} | Get environment
[**ListApiKeys**](AccountsApi.md#ListApiKeys) | **Get** /api-keys | List api keys
[**ListApiRequests**](AccountsApi.md#ListApiRequests) | **Get** /api-requests | List api requests
[**UpdateApiKey**](AccountsApi.md#UpdateApiKey) | **Put** /api-keys/{api_key_id} | Update api key
[**UpdateEnvironment**](AccountsApi.md#UpdateEnvironment) | **Put** /environments/{environment_id} | Update environment



## CountApiKeys

> CountApiKeysResponse CountApiKeys(ctx).RequireEnvironment(requireEnvironment).XSchematicEnvironmentId(xSchematicEnvironmentId).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()

Count api keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    requireEnvironment := true // bool | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    environmentId := "environmentId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CountApiKeys(context.Background()).RequireEnvironment(requireEnvironment).XSchematicEnvironmentId(xSchematicEnvironmentId).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CountApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountApiKeys`: CountApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CountApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requireEnvironment** | **bool** |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **environmentId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountApiKeysResponse**](CountApiKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountApiRequests

> CountApiRequestsResponse CountApiRequests(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()

Count api requests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CountApiRequests(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CountApiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountApiRequests`: CountApiRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CountApiRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountApiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountApiRequestsResponse**](CountApiRequestsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiKey

> CreateApiKeyResponse CreateApiKey(ctx).CreateApiKeyRequestBody(createApiKeyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create api key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    createApiKeyRequestBody := *openapiclient.NewCreateApiKeyRequestBody("Name_example") // CreateApiKeyRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateApiKey(context.Background()).CreateApiKeyRequestBody(createApiKeyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: CreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKeyRequestBody** | [**CreateApiKeyRequestBody**](CreateApiKeyRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnvironment

> CreateEnvironmentResponse CreateEnvironment(ctx).CreateEnvironmentRequestBody(createEnvironmentRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create environment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    createEnvironmentRequestBody := *openapiclient.NewCreateEnvironmentRequestBody("EnvironmentType_example", "Name_example") // CreateEnvironmentRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateEnvironment(context.Background()).CreateEnvironmentRequestBody(createEnvironmentRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: CreateEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEnvironmentRequestBody** | [**CreateEnvironmentRequestBody**](CreateEnvironmentRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateEnvironmentResponse**](CreateEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> DeleteApiKeyResponse DeleteApiKey(ctx, apiKeyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete api key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    apiKeyId := "apiKeyId_example" // string | api_key_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DeleteApiKey(context.Background(), apiKeyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKey`: DeleteApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | api_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteApiKeyResponse**](DeleteApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironmentResponse DeleteEnvironment(ctx, environmentId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete environment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    environmentId := "environmentId_example" // string | environment_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DeleteEnvironment(context.Background(), environmentId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironment`: DeleteEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | environment_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteEnvironmentResponse**](DeleteEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKey

> GetApiKeyResponse GetApiKey(ctx, apiKeyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get api key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    apiKeyId := "apiKeyId_example" // string | api_key_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetApiKey(context.Background(), apiKeyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: GetApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | api_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetApiKeyResponse**](GetApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiRequest

> GetApiRequestResponse GetApiRequest(ctx, key).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get api request

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    key := "key_example" // string | key
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetApiRequest(context.Background(), key).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetApiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiRequest`: GetApiRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetApiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetApiRequestResponse**](GetApiRequestResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> GetEnvironmentResponse GetEnvironment(ctx, environmentId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get environment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    environmentId := "environmentId_example" // string | environment_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetEnvironment(context.Background(), environmentId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: GetEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | environment_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetEnvironmentResponse**](GetEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeys

> ListApiKeysResponse ListApiKeys(ctx).RequireEnvironment(requireEnvironment).XSchematicEnvironmentId(xSchematicEnvironmentId).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()

List api keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    requireEnvironment := true // bool | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    environmentId := "environmentId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ListApiKeys(context.Background()).RequireEnvironment(requireEnvironment).XSchematicEnvironmentId(xSchematicEnvironmentId).EnvironmentId(environmentId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ListApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requireEnvironment** | **bool** |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **environmentId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListApiKeysResponse**](ListApiKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiRequests

> ListApiRequestsResponse ListApiRequests(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()

List api requests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.ListApiRequests(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListApiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiRequests`: ListApiRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListApiRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListApiRequestsResponse**](ListApiRequestsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKey

> UpdateApiKeyResponse UpdateApiKey(ctx, apiKeyId).UpdateApiKeyRequestBody(updateApiKeyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update api key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    apiKeyId := "apiKeyId_example" // string | api_key_id
    updateApiKeyRequestBody := *openapiclient.NewUpdateApiKeyRequestBody() // UpdateApiKeyRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateApiKey(context.Background(), apiKeyId).UpdateApiKeyRequestBody(updateApiKeyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKey`: UpdateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | api_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApiKeyRequestBody** | [**UpdateApiKeyRequestBody**](UpdateApiKeyRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateApiKeyResponse**](UpdateApiKeyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> UpdateEnvironmentResponse UpdateEnvironment(ctx, environmentId).UpdateEnvironmentRequestBody(updateEnvironmentRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update environment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    environmentId := "environmentId_example" // string | environment_id
    updateEnvironmentRequestBody := *openapiclient.NewUpdateEnvironmentRequestBody() // UpdateEnvironmentRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.UpdateEnvironment(context.Background(), environmentId).UpdateEnvironmentRequestBody(updateEnvironmentRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: UpdateEnvironmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | environment_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvironmentRequestBody** | [**UpdateEnvironmentRequestBody**](UpdateEnvironmentRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateEnvironmentResponse**](UpdateEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

