# \CompaniesAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCompanies**](CompaniesAPI.md#CountCompanies) | **Get** /companies/count | Count companies
[**CountEntityKeyDefinitions**](CompaniesAPI.md#CountEntityKeyDefinitions) | **Get** /entity-key-definitions/count | Count entity key definitions
[**CountEntityTraitDefinitions**](CompaniesAPI.md#CountEntityTraitDefinitions) | **Get** /entity-trait-definitions/count | Count entity trait definitions
[**CountUsers**](CompaniesAPI.md#CountUsers) | **Get** /users/count | Count users
[**CreateCompany**](CompaniesAPI.md#CreateCompany) | **Post** /companies/create | Create company
[**CreateUser**](CompaniesAPI.md#CreateUser) | **Post** /users/create | Create user
[**DeleteCompany**](CompaniesAPI.md#DeleteCompany) | **Delete** /companies/{company_id} | Delete company
[**DeleteCompanyByKeys**](CompaniesAPI.md#DeleteCompanyByKeys) | **Post** /companies/delete | Delete company by keys
[**DeleteCompanyMembership**](CompaniesAPI.md#DeleteCompanyMembership) | **Delete** /company-memberships/{company_membership_id} | Delete company membership
[**DeleteUser**](CompaniesAPI.md#DeleteUser) | **Delete** /users/{user_id} | Delete user
[**DeleteUserByKeys**](CompaniesAPI.md#DeleteUserByKeys) | **Post** /users/delete | Delete user by keys
[**GetActiveCompanySubscription**](CompaniesAPI.md#GetActiveCompanySubscription) | **Get** /company-subscriptions | Get active company subscription
[**GetActiveDeals**](CompaniesAPI.md#GetActiveDeals) | **Get** /company-crm-deals | Get active deals
[**GetCompany**](CompaniesAPI.md#GetCompany) | **Get** /companies/{company_id} | Get company
[**GetEntityTraitDefinition**](CompaniesAPI.md#GetEntityTraitDefinition) | **Get** /entity-trait-definitions/{entity_trait_definition_id} | Get entity trait definition
[**GetEntityTraitValues**](CompaniesAPI.md#GetEntityTraitValues) | **Get** /entity-trait-values | Get entity trait values
[**GetOrCreateCompanyMembership**](CompaniesAPI.md#GetOrCreateCompanyMembership) | **Post** /company-memberships | Get or create company membership
[**GetOrCreateEntityTraitDefinition**](CompaniesAPI.md#GetOrCreateEntityTraitDefinition) | **Post** /entity-trait-definitions | Get or create entity trait definition
[**GetUser**](CompaniesAPI.md#GetUser) | **Get** /users/{user_id} | Get user
[**ListCompanies**](CompaniesAPI.md#ListCompanies) | **Get** /companies | List companies
[**ListCompanyMemberships**](CompaniesAPI.md#ListCompanyMemberships) | **Get** /company-memberships | List company memberships
[**ListCompanyPlans**](CompaniesAPI.md#ListCompanyPlans) | **Get** /company-plans | List company plans
[**ListEntityKeyDefinitions**](CompaniesAPI.md#ListEntityKeyDefinitions) | **Get** /entity-key-definitions | List entity key definitions
[**ListEntityTraitDefinitions**](CompaniesAPI.md#ListEntityTraitDefinitions) | **Get** /entity-trait-definitions | List entity trait definitions
[**ListUsers**](CompaniesAPI.md#ListUsers) | **Get** /users | List users
[**LookupCompany**](CompaniesAPI.md#LookupCompany) | **Get** /companies/lookup | Lookup company
[**LookupUser**](CompaniesAPI.md#LookupUser) | **Get** /users/lookup | Lookup user
[**UpdateEntityTraitDefinition**](CompaniesAPI.md#UpdateEntityTraitDefinition) | **Put** /entity-trait-definitions/{entity_trait_definition_id} | Update entity trait definition
[**UpsertCompany**](CompaniesAPI.md#UpsertCompany) | **Post** /companies | Upsert company
[**UpsertCompanyTrait**](CompaniesAPI.md#UpsertCompanyTrait) | **Post** /company-traits | Upsert company trait
[**UpsertUser**](CompaniesAPI.md#UpsertUser) | **Post** /users | Upsert user
[**UpsertUserTrait**](CompaniesAPI.md#UpsertUserTrait) | **Post** /user-traits | Upsert user trait



## CountCompanies

> CountCompaniesResponse CountCompanies(ctx).Ids(ids).PlanId(planId).Q(q).WithoutFeatureOverrideFor(withoutFeatureOverrideFor).Limit(limit).Offset(offset).Execute()

Count companies

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

	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	q := "q_example" // string | Search filter (optional)
	withoutFeatureOverrideFor := "withoutFeatureOverrideFor_example" // string | Filter out companies that already have a company override for the specified feature ID (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.CountCompanies(context.Background()).Ids(ids).PlanId(planId).Q(q).WithoutFeatureOverrideFor(withoutFeatureOverrideFor).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CountCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCompanies`: CountCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CountCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **planId** | **string** |  | 
 **q** | **string** | Search filter | 
 **withoutFeatureOverrideFor** | **string** | Filter out companies that already have a company override for the specified feature ID | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountCompaniesResponse**](CountCompaniesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountEntityKeyDefinitions

> CountEntityKeyDefinitionsResponse CountEntityKeyDefinitions(ctx).EntityType(entityType).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

Count entity key definitions

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

	entityType := "entityType_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.CountEntityKeyDefinitions(context.Background()).EntityType(entityType).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CountEntityKeyDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountEntityKeyDefinitions`: CountEntityKeyDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CountEntityKeyDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountEntityKeyDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountEntityKeyDefinitionsResponse**](CountEntityKeyDefinitionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountEntityTraitDefinitions

> CountEntityTraitDefinitionsResponse CountEntityTraitDefinitions(ctx).EntityType(entityType).Ids(ids).TraitType(traitType).Q(q).Limit(limit).Offset(offset).Execute()

Count entity trait definitions

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

	entityType := "entityType_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	traitType := "traitType_example" // string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.CountEntityTraitDefinitions(context.Background()).EntityType(entityType).Ids(ids).TraitType(traitType).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CountEntityTraitDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountEntityTraitDefinitions`: CountEntityTraitDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CountEntityTraitDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountEntityTraitDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** |  | 
 **ids** | **[]string** |  | 
 **traitType** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountEntityTraitDefinitionsResponse**](CountEntityTraitDefinitionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountUsers

> CountUsersResponse CountUsers(ctx).CompanyId(companyId).Ids(ids).PlanId(planId).Q(q).Limit(limit).Offset(offset).Execute()

Count users

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

	companyId := "companyId_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	q := "q_example" // string | Search filter (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.CountUsers(context.Background()).CompanyId(companyId).Ids(ids).PlanId(planId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CountUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountUsers`: CountUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CountUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **ids** | **[]string** |  | 
 **planId** | **string** |  | 
 **q** | **string** | Search filter | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountUsersResponse**](CountUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertCompanyRequestBody := *schematicapi.NewUpsertCompanyRequestBody(map[string]string{"key": "Inner_example"}) // UpsertCompanyRequestBody | 

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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	upsertUserRequestBody := *schematicapi.NewUpsertUserRequestBody(map[string]string{"key": "Inner_example"}, map[string]string{"key": "Inner_example"}) // UpsertUserRequestBody | 

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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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


## DeleteCompanyByKeys

> DeleteCompanyByKeysResponse DeleteCompanyByKeys(ctx).KeysRequestBody(keysRequestBody).Execute()

Delete company by keys

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

	keysRequestBody := *schematicapi.NewKeysRequestBody(map[string]string{"key": "Inner_example"}) // KeysRequestBody | 

	resp, r, err := client.API().CompaniesAPI.DeleteCompanyByKeys(context.Background()).KeysRequestBody(keysRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.DeleteCompanyByKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyByKeys`: DeleteCompanyByKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.DeleteCompanyByKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyByKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keysRequestBody** | [**KeysRequestBody**](KeysRequestBody.md) |  | 

### Return type

[**DeleteCompanyByKeysResponse**](DeleteCompanyByKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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


## DeleteUserByKeys

> DeleteUserByKeysResponse DeleteUserByKeys(ctx).KeysRequestBody(keysRequestBody).Execute()

Delete user by keys

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

	keysRequestBody := *schematicapi.NewKeysRequestBody(map[string]string{"key": "Inner_example"}) // KeysRequestBody | 

	resp, r, err := client.API().CompaniesAPI.DeleteUserByKeys(context.Background()).KeysRequestBody(keysRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.DeleteUserByKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserByKeys`: DeleteUserByKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.DeleteUserByKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keysRequestBody** | [**KeysRequestBody**](KeysRequestBody.md) |  | 

### Return type

[**DeleteUserByKeysResponse**](DeleteUserByKeysResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveCompanySubscription

> GetActiveCompanySubscriptionResponse GetActiveCompanySubscription(ctx).CompanyId(companyId).Limit(limit).Offset(offset).Execute()

Get active company subscription

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

	companyId := "companyId_example" // string | 
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.GetActiveCompanySubscription(context.Background()).CompanyId(companyId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetActiveCompanySubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveCompanySubscription`: GetActiveCompanySubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetActiveCompanySubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveCompanySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**GetActiveCompanySubscriptionResponse**](GetActiveCompanySubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveDeals

> GetActiveDealsResponse GetActiveDeals(ctx).CompanyId(companyId).DealStage(dealStage).Limit(limit).Offset(offset).Execute()

Get active deals

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

	companyId := "companyId_example" // string | 
	dealStage := "dealStage_example" // string | 
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.GetActiveDeals(context.Background()).CompanyId(companyId).DealStage(dealStage).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetActiveDeals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveDeals`: GetActiveDealsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetActiveDeals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveDealsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **dealStage** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**GetActiveDealsResponse**](GetActiveDealsResponse.md)

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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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


## GetEntityTraitDefinition

> GetEntityTraitDefinitionResponse GetEntityTraitDefinition(ctx, entityTraitDefinitionId).Execute()

Get entity trait definition

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

	entityTraitDefinitionId := "entityTraitDefinitionId_example" // string | entity_trait_definition_id

	resp, r, err := client.API().CompaniesAPI.GetEntityTraitDefinition(context.Background(), entityTraitDefinitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetEntityTraitDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntityTraitDefinition`: GetEntityTraitDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetEntityTraitDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityTraitDefinitionId** | **string** | entity_trait_definition_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTraitDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEntityTraitDefinitionResponse**](GetEntityTraitDefinitionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityTraitValues

> GetEntityTraitValuesResponse GetEntityTraitValues(ctx).DefinitionId(definitionId).Q(q).Limit(limit).Offset(offset).Execute()

Get entity trait values

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

	definitionId := "definitionId_example" // string | 
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.GetEntityTraitValues(context.Background()).DefinitionId(definitionId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetEntityTraitValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntityTraitValues`: GetEntityTraitValuesResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetEntityTraitValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTraitValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **definitionId** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**GetEntityTraitValuesResponse**](GetEntityTraitValuesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrCreateCompanyMembership

> GetOrCreateCompanyMembershipResponse GetOrCreateCompanyMembership(ctx).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).Execute()

Get or create company membership

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

	getOrCreateCompanyMembershipRequestBody := *schematicapi.NewGetOrCreateCompanyMembershipRequestBody("CompanyId_example", "UserId_example") // GetOrCreateCompanyMembershipRequestBody | 

	resp, r, err := client.API().CompaniesAPI.GetOrCreateCompanyMembership(context.Background()).GetOrCreateCompanyMembershipRequestBody(getOrCreateCompanyMembershipRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetOrCreateCompanyMembership``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrCreateCompanyMembership`: GetOrCreateCompanyMembershipResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetOrCreateCompanyMembership`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrCreateCompanyMembershipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getOrCreateCompanyMembershipRequestBody** | [**GetOrCreateCompanyMembershipRequestBody**](GetOrCreateCompanyMembershipRequestBody.md) |  | 

### Return type

[**GetOrCreateCompanyMembershipResponse**](GetOrCreateCompanyMembershipResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrCreateEntityTraitDefinition

> GetOrCreateEntityTraitDefinitionResponse GetOrCreateEntityTraitDefinition(ctx).CreateEntityTraitDefinitionRequestBody(createEntityTraitDefinitionRequestBody).Execute()

Get or create entity trait definition

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

	createEntityTraitDefinitionRequestBody := *schematicapi.NewCreateEntityTraitDefinitionRequestBody("EntityType_example", []string{"Hierarchy_example"}, "TraitType_example") // CreateEntityTraitDefinitionRequestBody | 

	resp, r, err := client.API().CompaniesAPI.GetOrCreateEntityTraitDefinition(context.Background()).CreateEntityTraitDefinitionRequestBody(createEntityTraitDefinitionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetOrCreateEntityTraitDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrCreateEntityTraitDefinition`: GetOrCreateEntityTraitDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetOrCreateEntityTraitDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrCreateEntityTraitDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntityTraitDefinitionRequestBody** | [**CreateEntityTraitDefinitionRequestBody**](CreateEntityTraitDefinitionRequestBody.md) |  | 

### Return type

[**GetOrCreateEntityTraitDefinitionResponse**](GetOrCreateEntityTraitDefinitionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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

> ListCompaniesResponse ListCompanies(ctx).Ids(ids).PlanId(planId).Q(q).WithoutFeatureOverrideFor(withoutFeatureOverrideFor).Limit(limit).Offset(offset).Execute()

List companies

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

	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	q := "q_example" // string | Search filter (optional)
	withoutFeatureOverrideFor := "withoutFeatureOverrideFor_example" // string | Filter out companies that already have a company override for the specified feature ID (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListCompanies(context.Background()).Ids(ids).PlanId(planId).Q(q).WithoutFeatureOverrideFor(withoutFeatureOverrideFor).Limit(limit).Offset(offset).Execute()
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
 **planId** | **string** |  | 
 **q** | **string** | Search filter | 
 **withoutFeatureOverrideFor** | **string** | Filter out companies that already have a company override for the specified feature ID | 
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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


## ListEntityKeyDefinitions

> ListEntityKeyDefinitionsResponse ListEntityKeyDefinitions(ctx).EntityType(entityType).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

List entity key definitions

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

	entityType := "entityType_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListEntityKeyDefinitions(context.Background()).EntityType(entityType).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListEntityKeyDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntityKeyDefinitions`: ListEntityKeyDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListEntityKeyDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntityKeyDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListEntityKeyDefinitionsResponse**](ListEntityKeyDefinitionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntityTraitDefinitions

> ListEntityTraitDefinitionsResponse ListEntityTraitDefinitions(ctx).EntityType(entityType).Ids(ids).TraitType(traitType).Q(q).Limit(limit).Offset(offset).Execute()

List entity trait definitions

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

	entityType := "entityType_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	traitType := "traitType_example" // string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListEntityTraitDefinitions(context.Background()).EntityType(entityType).Ids(ids).TraitType(traitType).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.ListEntityTraitDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntityTraitDefinitions`: ListEntityTraitDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.ListEntityTraitDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntityTraitDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** |  | 
 **ids** | **[]string** |  | 
 **traitType** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListEntityTraitDefinitionsResponse**](ListEntityTraitDefinitionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> ListUsersResponse ListUsers(ctx).CompanyId(companyId).Ids(ids).PlanId(planId).Q(q).Limit(limit).Offset(offset).Execute()

List users

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

	companyId := "companyId_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	q := "q_example" // string | Search filter (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().CompaniesAPI.ListUsers(context.Background()).CompanyId(companyId).Ids(ids).PlanId(planId).Q(q).Limit(limit).Offset(offset).Execute()
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
 **companyId** | **string** |  | 
 **ids** | **[]string** |  | 
 **planId** | **string** |  | 
 **q** | **string** | Search filter | 
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
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


## UpsertCompany

> UpsertCompanyResponse UpsertCompany(ctx).UpsertCompanyRequestBody(upsertCompanyRequestBody).Execute()

Upsert company

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

	upsertCompanyRequestBody := *schematicapi.NewUpsertCompanyRequestBody(map[string]string{"key": "Inner_example"}) // UpsertCompanyRequestBody | 

	resp, r, err := client.API().CompaniesAPI.UpsertCompany(context.Background()).UpsertCompanyRequestBody(upsertCompanyRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpsertCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCompany`: UpsertCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpsertCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertCompanyRequestBody** | [**UpsertCompanyRequestBody**](UpsertCompanyRequestBody.md) |  | 

### Return type

[**UpsertCompanyResponse**](UpsertCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCompanyTrait

> UpsertCompanyTraitResponse UpsertCompanyTrait(ctx).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()

Upsert company trait

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

	upsertTraitRequestBody := *schematicapi.NewUpsertTraitRequestBody(map[string]string{"key": "Inner_example"}, "Trait_example") // UpsertTraitRequestBody | 

	resp, r, err := client.API().CompaniesAPI.UpsertCompanyTrait(context.Background()).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpsertCompanyTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCompanyTrait`: UpsertCompanyTraitResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpsertCompanyTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCompanyTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertTraitRequestBody** | [**UpsertTraitRequestBody**](UpsertTraitRequestBody.md) |  | 

### Return type

[**UpsertCompanyTraitResponse**](UpsertCompanyTraitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUser

> UpsertUserResponse UpsertUser(ctx).UpsertUserRequestBody(upsertUserRequestBody).Execute()

Upsert user

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

	upsertUserRequestBody := *schematicapi.NewUpsertUserRequestBody(map[string]string{"key": "Inner_example"}, map[string]string{"key": "Inner_example"}) // UpsertUserRequestBody | 

	resp, r, err := client.API().CompaniesAPI.UpsertUser(context.Background()).UpsertUserRequestBody(upsertUserRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpsertUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertUser`: UpsertUserResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpsertUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertUserRequestBody** | [**UpsertUserRequestBody**](UpsertUserRequestBody.md) |  | 

### Return type

[**UpsertUserResponse**](UpsertUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUserTrait

> UpsertUserTraitResponse UpsertUserTrait(ctx).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()

Upsert user trait

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

	upsertTraitRequestBody := *schematicapi.NewUpsertTraitRequestBody(map[string]string{"key": "Inner_example"}, "Trait_example") // UpsertTraitRequestBody | 

	resp, r, err := client.API().CompaniesAPI.UpsertUserTrait(context.Background()).UpsertTraitRequestBody(upsertTraitRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpsertUserTrait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertUserTrait`: UpsertUserTraitResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpsertUserTrait`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUserTraitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertTraitRequestBody** | [**UpsertTraitRequestBody**](UpsertTraitRequestBody.md) |  | 

### Return type

[**UpsertUserTraitResponse**](UpsertUserTraitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

