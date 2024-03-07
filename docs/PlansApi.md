# \PlansApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlan**](PlansApi.md#CreatePlan) | **Post** /plans | Create plan
[**DeletePlan**](PlansApi.md#DeletePlan) | **Delete** /plans/{plan_id} | Delete plan
[**GetPlan**](PlansApi.md#GetPlan) | **Get** /plans/{plan_id} | Get plan
[**ListPlans**](PlansApi.md#ListPlans) | **Get** /plans | List plans
[**UpdatePlan**](PlansApi.md#UpdatePlan) | **Put** /plans/{plan_id} | Update plan
[**UpdatePlanAudience**](PlansApi.md#UpdatePlanAudience) | **Put** /plan-audiences/{plan_audience_id} | Update plan audience



## CreatePlan

> CreatePlanResponse CreatePlan(ctx).CreatePlanRequestBody(createPlanRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create plan

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
    createPlanRequestBody := *openapiclient.NewCreatePlanRequestBody("Name_example") // CreatePlanRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.CreatePlan(context.Background()).CreatePlanRequestBody(createPlanRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.CreatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlan`: CreatePlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanRequestBody** | [**CreatePlanRequestBody**](CreatePlanRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreatePlanResponse**](CreatePlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlanResponse DeletePlan(ctx, planId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete plan

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
    planId := "planId_example" // string | plan_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.DeletePlan(context.Background(), planId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.DeletePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlan`: DeletePlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.DeletePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | plan_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeletePlanResponse**](DeletePlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> GetPlanResponse GetPlan(ctx, planId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get plan

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
    planId := "planId_example" // string | plan_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.GetPlan(context.Background(), planId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.GetPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlan`: GetPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | plan_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetPlanResponse**](GetPlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlans

> ListPlansResponse ListPlans(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Limit(limit).Offset(offset).Execute()

List plans

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
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.ListPlans(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.ListPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlans`: ListPlansResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.ListPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListPlansResponse**](ListPlansResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> UpdatePlanResponse UpdatePlan(ctx, planId).UpdatePlanRequestBody(updatePlanRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update plan

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
    planId := "planId_example" // string | plan_id
    updatePlanRequestBody := *openapiclient.NewUpdatePlanRequestBody("Name_example") // UpdatePlanRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.UpdatePlan(context.Background(), planId).UpdatePlanRequestBody(updatePlanRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.UpdatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlan`: UpdatePlanResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | plan_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlanRequestBody** | [**UpdatePlanRequestBody**](UpdatePlanRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdatePlanResponse**](UpdatePlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanAudience

> UpdatePlanAudienceResponse UpdatePlanAudience(ctx, planAudienceId).UpdateAudienceRequestBody(updateAudienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update plan audience

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
    planAudienceId := "planAudienceId_example" // string | plan_audience_id
    updateAudienceRequestBody := *openapiclient.NewUpdateAudienceRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // UpdateAudienceRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansApi.UpdatePlanAudience(context.Background(), planAudienceId).UpdateAudienceRequestBody(updateAudienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansApi.UpdatePlanAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlanAudience`: UpdatePlanAudienceResponse
    fmt.Fprintf(os.Stdout, "Response from `PlansApi.UpdatePlanAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planAudienceId** | **string** | plan_audience_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAudienceRequestBody** | [**UpdateAudienceRequestBody**](UpdateAudienceRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdatePlanAudienceResponse**](UpdatePlanAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

