# \FeaturesApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckFlag**](FeaturesApi.md#CheckFlag) | **Post** /flags/{key}/check | Check flag
[**CheckFlags**](FeaturesApi.md#CheckFlags) | **Post** /flags/check | Check flags
[**CountCompaniesAudience**](FeaturesApi.md#CountCompaniesAudience) | **Post** /audience/count-companies | Count Companies audience
[**CountFlagChecks**](FeaturesApi.md#CountFlagChecks) | **Get** /flag-checks/count | Count flag checks
[**CountUsersAudience**](FeaturesApi.md#CountUsersAudience) | **Post** /audience/count-users | Count Users audience
[**CreateFeature**](FeaturesApi.md#CreateFeature) | **Post** /features | Create feature
[**CreateFlag**](FeaturesApi.md#CreateFlag) | **Post** /flags | Create flag
[**CreateRule**](FeaturesApi.md#CreateRule) | **Post** /rules | Create rule
[**DeleteFeature**](FeaturesApi.md#DeleteFeature) | **Delete** /features/{feature_id} | Delete feature
[**DeleteFlag**](FeaturesApi.md#DeleteFlag) | **Delete** /flags/{flag_id} | Delete flag
[**GetCompaniesAudience**](FeaturesApi.md#GetCompaniesAudience) | **Post** /audience/get-companies | Get Companies audience
[**GetFeature**](FeaturesApi.md#GetFeature) | **Get** /features/{feature_id} | Get feature
[**GetFlag**](FeaturesApi.md#GetFlag) | **Get** /flags/{flag_id} | Get flag
[**GetFlagCheck**](FeaturesApi.md#GetFlagCheck) | **Get** /flag-checks/{flag_check_id} | Get flag check
[**GetRule**](FeaturesApi.md#GetRule) | **Get** /rules/{rule_id} | Get rule
[**GetUsersAudience**](FeaturesApi.md#GetUsersAudience) | **Post** /audience/get-users | Get Users audience
[**LatestFlagChecks**](FeaturesApi.md#LatestFlagChecks) | **Get** /flag-checks/latest | Latest flag checks
[**ListFeatures**](FeaturesApi.md#ListFeatures) | **Get** /features | List features
[**ListFlagChecks**](FeaturesApi.md#ListFlagChecks) | **Get** /flag-checks | List flag checks
[**ListFlags**](FeaturesApi.md#ListFlags) | **Get** /flags | List flags
[**UpdateFeature**](FeaturesApi.md#UpdateFeature) | **Put** /features/{feature_id} | Update feature
[**UpdateFlag**](FeaturesApi.md#UpdateFlag) | **Put** /flags/{flag_id} | Update flag
[**UpdateRule**](FeaturesApi.md#UpdateRule) | **Put** /rules/{rule_id} | Update rule



## CheckFlag

> CheckFlagResponse CheckFlag(ctx, key).CheckFlagRequestBody(checkFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Check flag

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
    checkFlagRequestBody := *openapiclient.NewCheckFlagRequestBody() // CheckFlagRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CheckFlag(context.Background(), key).CheckFlagRequestBody(checkFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CheckFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckFlag`: CheckFlagResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CheckFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFlagRequestBody** | [**CheckFlagRequestBody**](CheckFlagRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CheckFlagResponse**](CheckFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFlags

> CheckFlagsResponse CheckFlags(ctx).CheckFlagRequestBody(checkFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Check flags

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
    checkFlagRequestBody := *openapiclient.NewCheckFlagRequestBody() // CheckFlagRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CheckFlags(context.Background()).CheckFlagRequestBody(checkFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CheckFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckFlags`: CheckFlagsResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CheckFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkFlagRequestBody** | [**CheckFlagRequestBody**](CheckFlagRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CheckFlagsResponse**](CheckFlagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCompaniesAudience

> CountCompaniesAudienceResponse CountCompaniesAudience(ctx).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Count Companies audience

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
    audienceRequestBody := *openapiclient.NewAudienceRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CountCompaniesAudience(context.Background()).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CountCompaniesAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCompaniesAudience`: CountCompaniesAudienceResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CountCompaniesAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompaniesAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CountCompaniesAudienceResponse**](CountCompaniesAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountFlagChecks

> CountFlagChecksResponse CountFlagChecks(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

Count flag checks

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
    flagId := "flagId_example" // string |  (optional)
    flagIds := []string{"Inner_example"} // []string |  (optional)
    id := "id_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CountFlagChecks(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CountFlagChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountFlagChecks`: CountFlagChecksResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CountFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountFlagChecksResponse**](CountFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountUsersAudience

> CountUsersAudienceResponse CountUsersAudience(ctx).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Count Users audience

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
    audienceRequestBody := *openapiclient.NewAudienceRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CountUsersAudience(context.Background()).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CountUsersAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountUsersAudience`: CountUsersAudienceResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CountUsersAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountUsersAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CountUsersAudienceResponse**](CountUsersAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeature

> CreateFeatureResponse CreateFeature(ctx).CreateFeatureRequestBody(createFeatureRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create feature

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
    createFeatureRequestBody := *openapiclient.NewCreateFeatureRequestBody("Description_example", "FeatureType_example", "Name_example") // CreateFeatureRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CreateFeature(context.Background()).CreateFeatureRequestBody(createFeatureRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CreateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeature`: CreateFeatureResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CreateFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFeatureRequestBody** | [**CreateFeatureRequestBody**](CreateFeatureRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateFeatureResponse**](CreateFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlag

> CreateFlagResponse CreateFlag(ctx).CreateFlagRequestBody(createFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create flag

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
    createFlagRequestBody := *openapiclient.NewCreateFlagRequestBody(false, "Description_example", "FlagType_example", "Key_example", "Name_example", []openapiclient.CreateOrUpdateRuleRequestBody{*openapiclient.NewCreateOrUpdateRuleRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}, "Name_example", int32(123), false)}) // CreateFlagRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CreateFlag(context.Background()).CreateFlagRequestBody(createFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CreateFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlag`: CreateFlagResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CreateFlag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFlagRequestBody** | [**CreateFlagRequestBody**](CreateFlagRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateFlagResponse**](CreateFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> CreateRuleResponse CreateRule(ctx).CreateRuleRequestBody(createRuleRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create rule

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
    createRuleRequestBody := *openapiclient.NewCreateRuleRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}, "Name_example", int32(123), false) // CreateRuleRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.CreateRule(context.Background()).CreateRuleRequestBody(createRuleRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: CreateRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRuleRequestBody** | [**CreateRuleRequestBody**](CreateRuleRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateRuleResponse**](CreateRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeature

> DeleteFeatureResponse DeleteFeature(ctx, featureId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete feature

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
    featureId := "featureId_example" // string | feature_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.DeleteFeature(context.Background(), featureId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DeleteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFeature`: DeleteFeatureResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.DeleteFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteFeatureResponse**](DeleteFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlag

> DeleteFlagResponse DeleteFlag(ctx, flagId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete flag

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
    flagId := "flagId_example" // string | flag_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.DeleteFlag(context.Background(), flagId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.DeleteFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFlag`: DeleteFlagResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.DeleteFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteFlagResponse**](DeleteFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompaniesAudience

> GetCompaniesAudienceResponse GetCompaniesAudience(ctx).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get Companies audience

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
    audienceRequestBody := *openapiclient.NewAudienceRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetCompaniesAudience(context.Background()).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetCompaniesAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompaniesAudience`: GetCompaniesAudienceResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetCompaniesAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetCompaniesAudienceResponse**](GetCompaniesAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeature

> GetFeatureResponse GetFeature(ctx, featureId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get feature

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
    featureId := "featureId_example" // string | feature_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFeature(context.Background(), featureId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeature`: GetFeatureResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetFeatureResponse**](GetFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlag

> GetFlagResponse GetFlag(ctx, flagId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get flag

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
    flagId := "flagId_example" // string | flag_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFlag(context.Background(), flagId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlag`: GetFlagResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetFlagResponse**](GetFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagCheck

> GetFlagCheckResponse GetFlagCheck(ctx, flagCheckId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get flag check

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
    flagCheckId := "flagCheckId_example" // string | flag_check_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetFlagCheck(context.Background(), flagCheckId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetFlagCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFlagCheck`: GetFlagCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetFlagCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagCheckId** | **string** | flag_check_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetFlagCheckResponse**](GetFlagCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRule

> GetRuleResponse GetRule(ctx, ruleId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get rule

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
    ruleId := "ruleId_example" // string | rule_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetRule(context.Background(), ruleId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: GetRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | rule_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetRuleResponse**](GetRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAudience

> GetUsersAudienceResponse GetUsersAudience(ctx).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get Users audience

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
    audienceRequestBody := *openapiclient.NewAudienceRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}) // AudienceRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.GetUsersAudience(context.Background()).AudienceRequestBody(audienceRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.GetUsersAudience``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersAudience`: GetUsersAudienceResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.GetUsersAudience`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceRequestBody** | [**AudienceRequestBody**](AudienceRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetUsersAudienceResponse**](GetUsersAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LatestFlagChecks

> LatestFlagChecksResponse LatestFlagChecks(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

Latest flag checks

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
    flagId := "flagId_example" // string |  (optional)
    flagIds := []string{"Inner_example"} // []string |  (optional)
    id := "id_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.LatestFlagChecks(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.LatestFlagChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LatestFlagChecks`: LatestFlagChecksResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.LatestFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLatestFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**LatestFlagChecksResponse**](LatestFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatures

> ListFeaturesResponse ListFeatures(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Limit(limit).Offset(offset).Execute()

List features

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
    resp, r, err := apiClient.FeaturesApi.ListFeatures(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ListFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFeatures`: ListFeaturesResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ListFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFeaturesResponse**](ListFeaturesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlagChecks

> ListFlagChecksResponse ListFlagChecks(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()

List flag checks

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
    flagId := "flagId_example" // string |  (optional)
    flagIds := []string{"Inner_example"} // []string |  (optional)
    id := "id_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ListFlagChecks(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).FlagId(flagId).FlagIds(flagIds).Id(id).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ListFlagChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlagChecks`: ListFlagChecksResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ListFlagChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlagChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **flagId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **id** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFlagChecksResponse**](ListFlagChecksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlags

> ListFlagsResponse ListFlags(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).FeatureId(featureId).FlagIds(flagIds).Limit(limit).Offset(offset).Execute()

List flags

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
    featureId := "featureId_example" // string |  (optional)
    flagIds := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.ListFlags(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).FeatureId(featureId).FlagIds(flagIds).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.ListFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFlags`: ListFlagsResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.ListFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **featureId** | **string** |  | 
 **flagIds** | **[]string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFlagsResponse**](ListFlagsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> UpdateFeatureResponse UpdateFeature(ctx, featureId).UpdateFeatureRequestBody(updateFeatureRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update feature

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
    featureId := "featureId_example" // string | feature_id
    updateFeatureRequestBody := *openapiclient.NewUpdateFeatureRequestBody() // UpdateFeatureRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateFeature(context.Background(), featureId).UpdateFeatureRequestBody(updateFeatureRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeature`: UpdateFeatureResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureId** | **string** | feature_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFeatureRequestBody** | [**UpdateFeatureRequestBody**](UpdateFeatureRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateFeatureResponse**](UpdateFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlag

> UpdateFlagResponse UpdateFlag(ctx, flagId).CreateFlagRequestBody(createFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update flag

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
    flagId := "flagId_example" // string | flag_id
    createFlagRequestBody := *openapiclient.NewCreateFlagRequestBody(false, "Description_example", "FlagType_example", "Key_example", "Name_example", []openapiclient.CreateOrUpdateRuleRequestBody{*openapiclient.NewCreateOrUpdateRuleRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}, "Name_example", int32(123), false)}) // CreateFlagRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateFlag(context.Background(), flagId).CreateFlagRequestBody(createFlagRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlag`: UpdateFlagResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flagId** | **string** | flag_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFlagRequestBody** | [**CreateFlagRequestBody**](CreateFlagRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateFlagResponse**](UpdateFlagResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> UpdateRuleResponse UpdateRule(ctx, ruleId).UpdateRuleRequestBody(updateRuleRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update rule

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
    ruleId := "ruleId_example" // string | rule_id
    updateRuleRequestBody := *openapiclient.NewUpdateRuleRequestBody([]openapiclient.CreateOrUpdateConditionGroupRequestBody{*openapiclient.NewCreateOrUpdateConditionGroupRequestBody([]openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})})}, []openapiclient.CreateOrUpdateConditionRequestBody{*openapiclient.NewCreateOrUpdateConditionRequestBody("ConditionType_example", int32(123), "Operator_example", []string{"ResourceIds_example"})}, "Name_example", int32(123), false) // UpdateRuleRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeaturesApi.UpdateRule(context.Background(), ruleId).UpdateRuleRequestBody(updateRuleRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeaturesApi.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: UpdateRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `FeaturesApi.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | rule_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRuleRequestBody** | [**UpdateRuleRequestBody**](UpdateRuleRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateRuleResponse**](UpdateRuleResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

