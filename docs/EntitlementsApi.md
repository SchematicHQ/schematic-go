# \EntitlementsApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompanyOverride**](EntitlementsApi.md#CreateCompanyOverride) | **Post** /company-overrides | Create company override
[**CreatePlanEntitlement**](EntitlementsApi.md#CreatePlanEntitlement) | **Post** /plan-entitlements | Create plan entitlement
[**DeleteCompanyOverride**](EntitlementsApi.md#DeleteCompanyOverride) | **Delete** /company-overrides/{company_override_id} | Delete company override
[**DeletePlanEntitlement**](EntitlementsApi.md#DeletePlanEntitlement) | **Delete** /plan-entitlements/{plan_entitlement_id} | Delete plan entitlement
[**GetCompanyOverride**](EntitlementsApi.md#GetCompanyOverride) | **Get** /company-overrides/{company_override_id} | Get company override
[**GetPlanEntitlement**](EntitlementsApi.md#GetPlanEntitlement) | **Get** /plan-entitlements/{plan_entitlement_id} | Get plan entitlement
[**ListCompanyOverrides**](EntitlementsApi.md#ListCompanyOverrides) | **Get** /company-overrides | List company overrides
[**ListPlanEntitlements**](EntitlementsApi.md#ListPlanEntitlements) | **Get** /plan-entitlements | List plan entitlements
[**UpdateCompanyOverride**](EntitlementsApi.md#UpdateCompanyOverride) | **Put** /company-overrides/{company_override_id} | Update company override
[**UpdatePlanEntitlement**](EntitlementsApi.md#UpdatePlanEntitlement) | **Put** /plan-entitlements/{plan_entitlement_id} | Update plan entitlement



## CreateCompanyOverride

> CreateCompanyOverrideResponse CreateCompanyOverride(ctx).CreateCompanyOverrideRequestBody(createCompanyOverrideRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create company override

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
    createCompanyOverrideRequestBody := *openapiclient.NewCreateCompanyOverrideRequestBody("CompanyId_example", "FeatureId_example", "ValueType_example") // CreateCompanyOverrideRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.CreateCompanyOverride(context.Background()).CreateCompanyOverrideRequestBody(createCompanyOverrideRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.CreateCompanyOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompanyOverride`: CreateCompanyOverrideResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.CreateCompanyOverride`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCompanyOverrideRequestBody** | [**CreateCompanyOverrideRequestBody**](CreateCompanyOverrideRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateCompanyOverrideResponse**](CreateCompanyOverrideResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlanEntitlement

> CreatePlanEntitlementResponse CreatePlanEntitlement(ctx).CreatePlanEntitlementRequestBody(createPlanEntitlementRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create plan entitlement

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
    createPlanEntitlementRequestBody := *openapiclient.NewCreatePlanEntitlementRequestBody("FeatureId_example", "PlanId_example", "ValueType_example") // CreatePlanEntitlementRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.CreatePlanEntitlement(context.Background()).CreatePlanEntitlementRequestBody(createPlanEntitlementRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.CreatePlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlanEntitlement`: CreatePlanEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.CreatePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanEntitlementRequestBody** | [**CreatePlanEntitlementRequestBody**](CreatePlanEntitlementRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreatePlanEntitlementResponse**](CreatePlanEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyOverride

> DeleteCompanyOverrideResponse DeleteCompanyOverride(ctx, companyOverrideId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete company override

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
    companyOverrideId := "companyOverrideId_example" // string | company_override_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.DeleteCompanyOverride(context.Background(), companyOverrideId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.DeleteCompanyOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCompanyOverride`: DeleteCompanyOverrideResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.DeleteCompanyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyOverrideId** | **string** | company_override_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteCompanyOverrideResponse**](DeleteCompanyOverrideResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlanEntitlement

> DeletePlanEntitlementResponse DeletePlanEntitlement(ctx, planEntitlementId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete plan entitlement

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
    planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.DeletePlanEntitlement(context.Background(), planEntitlementId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.DeletePlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlanEntitlement`: DeletePlanEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.DeletePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planEntitlementId** | **string** | plan_entitlement_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeletePlanEntitlementResponse**](DeletePlanEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyOverride

> GetCompanyOverrideResponse GetCompanyOverride(ctx, companyOverrideId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get company override

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
    companyOverrideId := "companyOverrideId_example" // string | company_override_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.GetCompanyOverride(context.Background(), companyOverrideId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetCompanyOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyOverride`: GetCompanyOverrideResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetCompanyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyOverrideId** | **string** | company_override_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetCompanyOverrideResponse**](GetCompanyOverrideResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanEntitlement

> GetPlanEntitlementResponse GetPlanEntitlement(ctx, planEntitlementId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get plan entitlement

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
    planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.GetPlanEntitlement(context.Background(), planEntitlementId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetPlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanEntitlement`: GetPlanEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetPlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planEntitlementId** | **string** | plan_entitlement_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetPlanEntitlementResponse**](GetPlanEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyOverrides

> ListCompanyOverridesResponse ListCompanyOverrides(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()

List company overrides

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
    companyId := "companyId_example" // string |  (optional)
    featureId := "featureId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.ListCompanyOverrides(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListCompanyOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyOverrides`: ListCompanyOverridesResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListCompanyOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **companyId** | **string** |  | 
 **featureId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCompanyOverridesResponse**](ListCompanyOverridesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlanEntitlements

> ListPlanEntitlementsResponse ListPlanEntitlements(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).PlanId(planId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()

List plan entitlements

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
    planId := "planId_example" // string |  (optional)
    featureId := "featureId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.ListPlanEntitlements(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).PlanId(planId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListPlanEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlanEntitlements`: ListPlanEntitlementsResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListPlanEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlanEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **planId** | **string** |  | 
 **featureId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListPlanEntitlementsResponse**](ListPlanEntitlementsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyOverride

> UpdateCompanyOverrideResponse UpdateCompanyOverride(ctx, companyOverrideId).UpdateCompanyOverrideRequestBody(updateCompanyOverrideRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update company override

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
    companyOverrideId := "companyOverrideId_example" // string | company_override_id
    updateCompanyOverrideRequestBody := *openapiclient.NewUpdateCompanyOverrideRequestBody("ValueType_example") // UpdateCompanyOverrideRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.UpdateCompanyOverride(context.Background(), companyOverrideId).UpdateCompanyOverrideRequestBody(updateCompanyOverrideRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.UpdateCompanyOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompanyOverride`: UpdateCompanyOverrideResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.UpdateCompanyOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyOverrideId** | **string** | company_override_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCompanyOverrideRequestBody** | [**UpdateCompanyOverrideRequestBody**](UpdateCompanyOverrideRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateCompanyOverrideResponse**](UpdateCompanyOverrideResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanEntitlement

> UpdatePlanEntitlementResponse UpdatePlanEntitlement(ctx, planEntitlementId).UpdatePlanEntitlementRequestBody(updatePlanEntitlementRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update plan entitlement

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
    planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id
    updatePlanEntitlementRequestBody := *openapiclient.NewUpdatePlanEntitlementRequestBody("ValueType_example") // UpdatePlanEntitlementRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.UpdatePlanEntitlement(context.Background(), planEntitlementId).UpdatePlanEntitlementRequestBody(updatePlanEntitlementRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.UpdatePlanEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlanEntitlement`: UpdatePlanEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.UpdatePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planEntitlementId** | **string** | plan_entitlement_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlanEntitlementRequestBody** | [**UpdatePlanEntitlementRequestBody**](UpdatePlanEntitlementRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdatePlanEntitlementResponse**](UpdatePlanEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

