# \CompaniesAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompaniesAPI.md#CreateCompany) | **Post** /companies | Create company
[**CreateCompanyMembership**](CompaniesAPI.md#CreateCompanyMembership) | **Post** /company-memberships | Create company membership
[**CreateCompanyTrait**](CompaniesAPI.md#CreateCompanyTrait) | **Post** /company-traits | Create company trait
[**CreateEntityTraitDefinition**](CompaniesAPI.md#CreateEntityTraitDefinition) | **Post** /entity-trait-definitions | Create entity trait definition
[**CreateUser**](CompaniesAPI.md#CreateUser) | **Post** /users | Create user
[**CreateUserTrait**](CompaniesAPI.md#CreateUserTrait) | **Post** /user-traits | Create user trait
[**DeleteCompany**](CompaniesAPI.md#DeleteCompany) | **Delete** /companies/{company_id} | Delete company
[**DeleteCompanyMembership**](CompaniesAPI.md#DeleteCompanyMembership) | **Delete** /company-memberships/{company_membership_id} | Delete company membership
[**DeleteUser**](CompaniesAPI.md#DeleteUser) | **Delete** /users/{user_id} | Delete user
[**GetCompany**](CompaniesAPI.md#GetCompany) | **Get** /companies/{company_id} | Get company
[**GetUser**](CompaniesAPI.md#GetUser) | **Get** /users/{user_id} | Get user
[**ListCompanies**](CompaniesAPI.md#ListCompanies) | **Get** /companies | List companies
[**ListCompanyMemberships**](CompaniesAPI.md#ListCompanyMemberships) | **Get** /company-memberships | List company memberships
[**ListCompanyPlans**](CompaniesAPI.md#ListCompanyPlans) | **Get** /company-plans | List company plans
[**ListUsers**](CompaniesAPI.md#ListUsers) | **Get** /users | List users
[**LookupCompany**](CompaniesAPI.md#LookupCompany) | **Get** /companies/lookup | Lookup company
[**LookupUser**](CompaniesAPI.md#LookupUser) | **Get** /users/lookup | Lookup user
[**UpdateEntityTraitDefinition**](CompaniesAPI.md#UpdateEntityTraitDefinition) | **Put** /entity-trait-definitions/{entity_trait_definition_id} | Update entity trait definition



## CreateCompany

> CreateCompanyResponse CreateCompany(ctx).UpsertCompanyRequestBody(upsertCompanyRequestBody).Execute()

Create company

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertCompanyRequestBody := *schematicapi.NewUpsertCompanyRequestBody(map[string]interface{}(123)) // UpsertCompanyRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateCompany(context.Background()).UpsertCompanyRequestBody(upsertCompanyRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompany`: CreateCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertCompanyRequestBody** | [**UpsertCompanyRequestBody**](UpsertCompanyRequestBody.md) |  | 

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

> CreateCompanyMembershipResponse CreateCompanyMembership(ctx).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).Execute()

Create company membership

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	getOrCreateCompanyMembershipRequestBody := *schematicapi.NewGetOrCreateCompanyMembershipRequestBody("CompanyId_example", "UserId_example") // GetOrCreateCompanyMembershipRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateCompanyMembership(context.Background()).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateCompanyMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompanyMembership`: CreateCompanyMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateCompanyMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrCreateCompanyMembershipRequestBody** | [**GetOrCreateCompanyMembershipRequestBody**](GetOrCreateCompanyMembershipRequestBody.md) |  | 

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


## CreateCompanyTrait

> CreateCompanyTraitResponse CreateCompanyTrait(ctx).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()

Create company trait

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertTraitRequestBody := *schematicapi.NewUpsertTraitRequestBody(map[string]interface{}(123), "Trait_example") // UpsertTraitRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateCompanyTrait(context.Background()).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateCompanyTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompanyTrait`: CreateCompanyTraitResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateCompanyTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertTraitRequestBody** | [**UpsertTraitRequestBody**](UpsertTraitRequestBody.md) |  | 

### Return type

[**CreateCompanyTraitResponse**](CreateCompanyTraitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntityTraitDefinition

> CreateEntityTraitDefinitionResponse CreateEntityTraitDefinition(ctx).CreateEntityTraitDefinitionRequestBody(createEntityTraitDefinitionRequestBody).Execute()

Create entity trait definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	createEntityTraitDefinitionRequestBody := *schematicapi.NewCreateEntityTraitDefinitionRequestBody("EntityType_example", []string{"Hierarchy_example"}, "TraitType_example") // CreateEntityTraitDefinitionRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateEntityTraitDefinition(context.Background()).CreateEntityTraitDefinitionRequestBody(createEntityTraitDefinitionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateEntityTraitDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntityTraitDefinition`: CreateEntityTraitDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateEntityTraitDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityTraitDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntityTraitDefinitionRequestBody** | [**CreateEntityTraitDefinitionRequestBody**](CreateEntityTraitDefinitionRequestBody.md) |  | 

### Return type

[**CreateEntityTraitDefinitionResponse**](CreateEntityTraitDefinitionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> CreateUserResponse CreateUser(ctx).UpsertUserRequestBody(upsertUserRequestBody).Execute()

Create user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertUserRequestBody := *schematicapi.NewUpsertUserRequestBody(map[string]interface{}(123), map[string]interface{}(123)) // UpsertUserRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateUser(context.Background()).UpsertUserRequestBody(upsertUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUser`: CreateUserResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertUserRequestBody** | [**UpsertUserRequestBody**](UpsertUserRequestBody.md) |  | 

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


## CreateUserTrait

> CreateUserTraitResponse CreateUserTrait(ctx).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()

Create user trait

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertTraitRequestBody := *schematicapi.NewUpsertTraitRequestBody(map[string]interface{}(123), "Trait_example") // UpsertTraitRequestBody | 

	resp, r, err := client.API().CompaniesAPI.CreateUserTrait(context.Background()).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateUserTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserTrait`: CreateUserTraitResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateUserTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertTraitRequestBody** | [**UpsertTraitRequestBody**](UpsertTraitRequestBody.md) |  | 

### Return type

[**CreateUserTraitResponse**](CreateUserTraitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompany

> DeleteCompanyResponse DeleteCompany(ctx, companyId).Execute()

Delete company

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	companyId := "companyId_example" // string | company_id

	resp, r, err := client.API().CompaniesAPI.DeleteCompany(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.DeleteCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompany`: DeleteCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.DeleteCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | company_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCompanyResponse**](DeleteCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyMembership

> DeleteCompanyMembershipResponse DeleteCompanyMembership(ctx, companyMembershipId).Execute()

Delete company membership

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	companyMembershipId := "companyMembershipId_example" // string | company_membership_id

	resp, r, err := client.API().CompaniesAPI.DeleteCompanyMembership(context.Background(), companyMembershipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.DeleteCompanyMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyMembership`: DeleteCompanyMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.DeleteCompanyMembership`: %v\n", resp)
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


## DeleteUser

> DeleteUserResponse DeleteUser(ctx, userId).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	userId := "userId_example" // string | user_id

	resp, r, err := client.API().CompaniesAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: DeleteUserResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteUserResponse**](DeleteUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> GetCompanyResponse GetCompany(ctx, companyId).Execute()

Get company

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	companyId := "companyId_example" // string | company_id

	resp, r, err := client.API().CompaniesAPI.GetCompany(context.Background(), companyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompany`: GetCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetCompany`: %v\n", resp)
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

> GetUserResponse GetUser(ctx, userId).Execute()

Get user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	userId := "userId_example" // string | user_id

	resp, r, err := client.API().CompaniesAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUserResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetUser`: %v\n", resp)
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

> ListCompaniesResponse ListCompanies(ctx).Ids(ids).Limit(limit).Offset(offset).Execute()

List companies

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	ids := []string{"Inner_example"} // []string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListCompanies(context.Background()).Ids(ids).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanies`: ListCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> ListCompanyMembershipsResponse ListCompanyMemberships(ctx).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()

List company memberships

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	companyId := "companyId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListCompanyMemberships(context.Background()).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompanyMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanyMemberships`: ListCompanyMembershipsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompanyMemberships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

> ListCompanyPlansResponse ListCompanyPlans(ctx).CompanyId(companyId).PlanId(planId).Limit(limit).Offset(offset).Execute()

List company plans

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	companyId := "companyId_example" // string |  (optional)
	planId := "planId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListCompanyPlans(context.Background()).CompanyId(companyId).PlanId(planId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListCompanyPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanyPlans`: ListCompanyPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListCompanyPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **planId** | **string** |  | 
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

> ListUsersResponse ListUsers(ctx).Ids(ids).Limit(limit).Offset(offset).Execute()

List users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	ids := []string{"Inner_example"} // []string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListUsers(context.Background()).Ids(ids).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: ListUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## LookupCompany

> LookupCompanyResponse LookupCompany(ctx).Keys(keys).Execute()

Lookup company

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	keys := map[string]interface{}{ ... } // map[string]interface{} | Key/value pairs

	resp, r, err := client.API().CompaniesAPI.LookupCompany(context.Background()).Keys(keys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.LookupCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupCompany`: LookupCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.LookupCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keys** | [**map[string]interface{}**](map[string]interface{}.md) | Key/value pairs | 

### Return type

[**LookupCompanyResponse**](LookupCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupUser

> LookupUserResponse LookupUser(ctx).Keys(keys).Execute()

Lookup user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	keys := map[string]interface{}{ ... } // map[string]interface{} | Key/value pairs

	resp, r, err := client.API().CompaniesAPI.LookupUser(context.Background()).Keys(keys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.LookupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUser`: LookupUserResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.LookupUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keys** | [**map[string]interface{}**](map[string]interface{}.md) | Key/value pairs | 

### Return type

[**LookupUserResponse**](LookupUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityTraitDefinition

> UpdateEntityTraitDefinitionResponse UpdateEntityTraitDefinition(ctx, entityTraitDefinitionId).UpdateEntityTraitDefinitionRequestBody(updateEntityTraitDefinitionRequestBody).Execute()

Update entity trait definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	entityTraitDefinitionId := "entityTraitDefinitionId_example" // string | entity_trait_definition_id
	updateEntityTraitDefinitionRequestBody := *schematicapi.NewUpdateEntityTraitDefinitionRequestBody("TraitType_example") // UpdateEntityTraitDefinitionRequestBody | 

	resp, r, err := client.API().CompaniesAPI.UpdateEntityTraitDefinition(context.Background(), entityTraitDefinitionId).UpdateEntityTraitDefinitionRequestBody(updateEntityTraitDefinitionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpdateEntityTraitDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntityTraitDefinition`: UpdateEntityTraitDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpdateEntityTraitDefinition`: %v\n", resp)
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

