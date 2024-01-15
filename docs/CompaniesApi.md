# \CompaniesApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompaniesApi.md#CreateCompany) | **Post** /companies | Create company
[**CreateCompanyMembership**](CompaniesApi.md#CreateCompanyMembership) | **Post** /company-memberships | Create company membership
[**CreateUser**](CompaniesApi.md#CreateUser) | **Post** /users | Create user
[**DeleteCompanyMembership**](CompaniesApi.md#DeleteCompanyMembership) | **Delete** /company-memberships/{company_membership_id} | Delete company membership
[**GetCompany**](CompaniesApi.md#GetCompany) | **Get** /companies/{company_id} | Get company
[**GetUser**](CompaniesApi.md#GetUser) | **Get** /users/{user_id} | Get user
[**ListCompanies**](CompaniesApi.md#ListCompanies) | **Get** /companies | List companies
[**ListCompanyMemberships**](CompaniesApi.md#ListCompanyMemberships) | **Get** /company-memberships | List company memberships
[**ListCompanyPlans**](CompaniesApi.md#ListCompanyPlans) | **Get** /company-plans | List company plans
[**ListUsers**](CompaniesApi.md#ListUsers) | **Get** /users | List users
[**UpdateEntityTraitDefinition**](CompaniesApi.md#UpdateEntityTraitDefinition) | **Put** /entity-trait-definitions/{entity_trait_definition_id} | Update entity trait definition



## CreateCompany

> CreateCompanyResponse CreateCompany(ctx).UpsertCompanyRequestBody(upsertCompanyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create company

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
    upsertCompanyRequestBody := *openapiclient.NewUpsertCompanyRequestBody(map[string]interface{}(123)) // UpsertCompanyRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.CreateCompany(context.Background()).UpsertCompanyRequestBody(upsertCompanyRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CreateCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompany`: CreateCompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertCompanyRequestBody** | [**UpsertCompanyRequestBody**](UpsertCompanyRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateCompanyResponse**](CreateCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyMembership

> CreateCompanyMembershipResponse CreateCompanyMembership(ctx).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create company membership

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
    getOrCreateCompanyMembershipRequestBody := *openapiclient.NewGetOrCreateCompanyMembershipRequestBody("CompanyId_example", "UserId_example") // GetOrCreateCompanyMembershipRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.CreateCompanyMembership(context.Background()).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CreateCompanyMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompanyMembership`: CreateCompanyMembershipResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CreateCompanyMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrCreateCompanyMembershipRequestBody** | [**GetOrCreateCompanyMembershipRequestBody**](GetOrCreateCompanyMembershipRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateCompanyMembershipResponse**](CreateCompanyMembershipResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUserResponse CreateUser(ctx).UpsertUserRequestBody(upsertUserRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create user

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
    upsertUserRequestBody := *openapiclient.NewUpsertUserRequestBody(map[string]interface{}(123), map[string]interface{}(123)) // UpsertUserRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.CreateUser(context.Background()).UpsertUserRequestBody(upsertUserRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: CreateUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertUserRequestBody** | [**UpsertUserRequestBody**](UpsertUserRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**CreateUserResponse**](CreateUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyMembership

> DeleteCompanyMembershipResponse DeleteCompanyMembership(ctx, companyMembershipId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Delete company membership

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
    companyMembershipId := "companyMembershipId_example" // string | company_membership_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.DeleteCompanyMembership(context.Background(), companyMembershipId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.DeleteCompanyMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCompanyMembership`: DeleteCompanyMembershipResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.DeleteCompanyMembership`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyMembershipId** | **string** | company_membership_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**DeleteCompanyMembershipResponse**](DeleteCompanyMembershipResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> GetCompanyResponse GetCompany(ctx, companyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get company

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
    companyId := "companyId_example" // string | company_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetCompany(context.Background(), companyId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: GetCompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | company_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetCompanyResponse**](GetCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUserResponse GetUser(ctx, userId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get user

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
    userId := "userId_example" // string | user_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.GetUser(context.Background(), userId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: GetUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetUserResponse**](GetUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanies

> ListCompaniesResponse ListCompanies(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Ids(ids).Limit(limit).Offset(offset).Execute()

List companies

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
    ids := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ListCompanies(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Ids(ids).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanies`: ListCompaniesResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **ids** | **[]string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCompaniesResponse**](ListCompaniesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyMemberships

> ListCompanyMembershipsResponse ListCompanyMemberships(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()

List company memberships

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
    userId := "userId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ListCompanyMemberships(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListCompanyMemberships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyMemberships`: ListCompanyMembershipsResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListCompanyMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **companyId** | **string** |  | 
 **userId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCompanyMembershipsResponse**](ListCompanyMembershipsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompanyPlans

> ListCompanyPlansResponse ListCompanyPlans(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).PlanId(planId).Active(active).Limit(limit).Offset(offset).Execute()

List company plans

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
    planId := "planId_example" // string |  (optional)
    active := true // bool |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ListCompanyPlans(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).PlanId(planId).Active(active).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListCompanyPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompanyPlans`: ListCompanyPlansResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListCompanyPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **companyId** | **string** |  | 
 **planId** | **string** |  | 
 **active** | **bool** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCompanyPlansResponse**](ListCompanyPlansResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsersResponse ListUsers(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Ids(ids).Limit(limit).Offset(offset).Execute()

List users

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
    ids := []string{"Inner_example"} // []string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.ListUsers(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Ids(ids).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: ListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **ids** | **[]string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityTraitDefinition

> UpdateEntityTraitDefinitionResponse UpdateEntityTraitDefinition(ctx, entityTraitDefinitionId).UpdateEntityTraitDefinitionRequestBody(updateEntityTraitDefinitionRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Update entity trait definition

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
    entityTraitDefinitionId := "entityTraitDefinitionId_example" // string | entity_trait_definition_id
    updateEntityTraitDefinitionRequestBody := *openapiclient.NewUpdateEntityTraitDefinitionRequestBody("TraitType_example") // UpdateEntityTraitDefinitionRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompaniesApi.UpdateEntityTraitDefinition(context.Background(), entityTraitDefinitionId).UpdateEntityTraitDefinitionRequestBody(updateEntityTraitDefinitionRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesApi.UpdateEntityTraitDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityTraitDefinition`: UpdateEntityTraitDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesApi.UpdateEntityTraitDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityTraitDefinitionId** | **string** | entity_trait_definition_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityTraitDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityTraitDefinitionRequestBody** | [**UpdateEntityTraitDefinitionRequestBody**](UpdateEntityTraitDefinitionRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**UpdateEntityTraitDefinitionResponse**](UpdateEntityTraitDefinitionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

