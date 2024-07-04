# \EntitlementsAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCompanyOverrides**](EntitlementsAPI.md#CountCompanyOverrides) | **Get** /company-overrides/count | Count company overrides
[**CountFeatureCompanies**](EntitlementsAPI.md#CountFeatureCompanies) | **Get** /feature-companies/count | Count feature companies
[**CountFeatureUsage**](EntitlementsAPI.md#CountFeatureUsage) | **Get** /feature-usage/count | Count feature usage
[**CountFeatureUsers**](EntitlementsAPI.md#CountFeatureUsers) | **Get** /feature-users/count | Count feature users
[**CountPlanEntitlements**](EntitlementsAPI.md#CountPlanEntitlements) | **Get** /plan-entitlements/count | Count plan entitlements
[**CreateCompanyOverride**](EntitlementsAPI.md#CreateCompanyOverride) | **Post** /company-overrides | Create company override
[**CreatePlanEntitlement**](EntitlementsAPI.md#CreatePlanEntitlement) | **Post** /plan-entitlements | Create plan entitlement
[**DeleteCompanyOverride**](EntitlementsAPI.md#DeleteCompanyOverride) | **Delete** /company-overrides/{company_override_id} | Delete company override
[**DeletePlanEntitlement**](EntitlementsAPI.md#DeletePlanEntitlement) | **Delete** /plan-entitlements/{plan_entitlement_id} | Delete plan entitlement
[**GetCompanyOverride**](EntitlementsAPI.md#GetCompanyOverride) | **Get** /company-overrides/{company_override_id} | Get company override
[**GetFeatureUsageByCompany**](EntitlementsAPI.md#GetFeatureUsageByCompany) | **Get** /usage-by-company | Get feature usage by company
[**GetPlanEntitlement**](EntitlementsAPI.md#GetPlanEntitlement) | **Get** /plan-entitlements/{plan_entitlement_id} | Get plan entitlement
[**ListCompanyOverrides**](EntitlementsAPI.md#ListCompanyOverrides) | **Get** /company-overrides | List company overrides
[**ListFeatureCompanies**](EntitlementsAPI.md#ListFeatureCompanies) | **Get** /feature-companies | List feature companies
[**ListFeatureUsage**](EntitlementsAPI.md#ListFeatureUsage) | **Get** /feature-usage | List feature usage
[**ListFeatureUsers**](EntitlementsAPI.md#ListFeatureUsers) | **Get** /feature-users | List feature users
[**ListPlanEntitlements**](EntitlementsAPI.md#ListPlanEntitlements) | **Get** /plan-entitlements | List plan entitlements
[**UpdateCompanyOverride**](EntitlementsAPI.md#UpdateCompanyOverride) | **Put** /company-overrides/{company_override_id} | Update company override
[**UpdatePlanEntitlement**](EntitlementsAPI.md#UpdatePlanEntitlement) | **Put** /plan-entitlements/{plan_entitlement_id} | Update plan entitlement



## CountCompanyOverrides

> CountCompanyOverridesResponse CountCompanyOverrides(ctx).CompanyId(companyId).CompanyIds(companyIds).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

Count company overrides

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
	companyIds := []string{"Inner_example"} // []string |  (optional)
	featureId := "featureId_example" // string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.CountCompanyOverrides(context.Background()).CompanyId(companyId).CompanyIds(companyIds).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CountCompanyOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCompanyOverrides`: CountCompanyOverridesResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CountCompanyOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompanyOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **companyIds** | **[]string** |  | 
 **featureId** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountCompanyOverridesResponse**](CountCompanyOverridesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountFeatureCompanies

> CountFeatureCompaniesResponse CountFeatureCompanies(ctx).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()

Count feature companies

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

	featureId := "featureId_example" // string | 
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.CountFeatureCompanies(context.Background()).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CountFeatureCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFeatureCompanies`: CountFeatureCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CountFeatureCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountFeatureCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountFeatureCompaniesResponse**](CountFeatureCompaniesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountFeatureUsage

> CountFeatureUsageResponse CountFeatureUsage(ctx).CompanyId(companyId).CompanyKeys(companyKeys).FeatureIds(featureIds).Q(q).Limit(limit).Offset(offset).Execute()

Count feature usage

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
	companyKeys := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.CountFeatureUsage(context.Background()).CompanyId(companyId).CompanyKeys(companyKeys).FeatureIds(featureIds).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CountFeatureUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFeatureUsage`: CountFeatureUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CountFeatureUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountFeatureUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **companyKeys** | **map[string]string** |  | 
 **featureIds** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountFeatureUsageResponse**](CountFeatureUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountFeatureUsers

> CountFeatureUsersResponse CountFeatureUsers(ctx).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()

Count feature users

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

	featureId := "featureId_example" // string | 
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.CountFeatureUsers(context.Background()).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CountFeatureUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFeatureUsers`: CountFeatureUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CountFeatureUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountFeatureUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountFeatureUsersResponse**](CountFeatureUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountPlanEntitlements

> CountPlanEntitlementsResponse CountPlanEntitlements(ctx).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).PlanId(planId).PlanIds(planIds).Q(q).Limit(limit).Offset(offset).Execute()

Count plan entitlements

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

	featureId := "featureId_example" // string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.CountPlanEntitlements(context.Background()).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).PlanId(planId).PlanIds(planIds).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CountPlanEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountPlanEntitlements`: CountPlanEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CountPlanEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPlanEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **ids** | **[]string** |  | 
 **planId** | **string** |  | 
 **planIds** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountPlanEntitlementsResponse**](CountPlanEntitlementsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyOverride

> CreateCompanyOverrideResponse CreateCompanyOverride(ctx).CreateCompanyOverrideRequestBody(createCompanyOverrideRequestBody).Execute()

Create company override

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

	createCompanyOverrideRequestBody := *schematicapi.NewCreateCompanyOverrideRequestBody("CompanyId_example", "FeatureId_example", "ValueType_example") // CreateCompanyOverrideRequestBody | 

	resp, r, err := client.API().EntitlementsAPI.CreateCompanyOverride(context.Background()).CreateCompanyOverrideRequestBody(createCompanyOverrideRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateCompanyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompanyOverride`: CreateCompanyOverrideResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateCompanyOverride`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCompanyOverrideRequestBody** | [**CreateCompanyOverrideRequestBody**](CreateCompanyOverrideRequestBody.md) |  | 

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

> CreatePlanEntitlementResponse CreatePlanEntitlement(ctx).CreatePlanEntitlementRequestBody(createPlanEntitlementRequestBody).Execute()

Create plan entitlement

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

	createPlanEntitlementRequestBody := *schematicapi.NewCreatePlanEntitlementRequestBody("FeatureId_example", "PlanId_example", "ValueType_example") // CreatePlanEntitlementRequestBody | 

	resp, r, err := client.API().EntitlementsAPI.CreatePlanEntitlement(context.Background()).CreatePlanEntitlementRequestBody(createPlanEntitlementRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreatePlanEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanEntitlement`: CreatePlanEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreatePlanEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanEntitlementRequestBody** | [**CreatePlanEntitlementRequestBody**](CreatePlanEntitlementRequestBody.md) |  | 

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

> DeleteCompanyOverrideResponse DeleteCompanyOverride(ctx, companyOverrideId).Execute()

Delete company override

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

	companyOverrideId := "companyOverrideId_example" // string | company_override_id

	resp, r, err := client.API().EntitlementsAPI.DeleteCompanyOverride(context.Background(), companyOverrideId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.DeleteCompanyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyOverride`: DeleteCompanyOverrideResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.DeleteCompanyOverride`: %v\n", resp)
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

> DeletePlanEntitlementResponse DeletePlanEntitlement(ctx, planEntitlementId).Execute()

Delete plan entitlement

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

	planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id

	resp, r, err := client.API().EntitlementsAPI.DeletePlanEntitlement(context.Background(), planEntitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.DeletePlanEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlanEntitlement`: DeletePlanEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.DeletePlanEntitlement`: %v\n", resp)
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

> GetCompanyOverrideResponse GetCompanyOverride(ctx, companyOverrideId).Execute()

Get company override

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

	companyOverrideId := "companyOverrideId_example" // string | company_override_id

	resp, r, err := client.API().EntitlementsAPI.GetCompanyOverride(context.Background(), companyOverrideId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetCompanyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyOverride`: GetCompanyOverrideResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetCompanyOverride`: %v\n", resp)
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


## GetFeatureUsageByCompany

> GetFeatureUsageByCompanyResponse GetFeatureUsageByCompany(ctx).Keys(keys).Execute()

Get feature usage by company

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

	resp, r, err := client.API().EntitlementsAPI.GetFeatureUsageByCompany(context.Background()).Keys(keys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetFeatureUsageByCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureUsageByCompany`: GetFeatureUsageByCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetFeatureUsageByCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureUsageByCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keys** | [**map[string]interface{}**](map[string]interface{}.md) | Key/value pairs | 

### Return type

[**GetFeatureUsageByCompanyResponse**](GetFeatureUsageByCompanyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanEntitlement

> GetPlanEntitlementResponse GetPlanEntitlement(ctx, planEntitlementId).Execute()

Get plan entitlement

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

	planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id

	resp, r, err := client.API().EntitlementsAPI.GetPlanEntitlement(context.Background(), planEntitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetPlanEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanEntitlement`: GetPlanEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetPlanEntitlement`: %v\n", resp)
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

> ListCompanyOverridesResponse ListCompanyOverrides(ctx).CompanyId(companyId).CompanyIds(companyIds).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

List company overrides

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
	companyIds := []string{"Inner_example"} // []string |  (optional)
	featureId := "featureId_example" // string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListCompanyOverrides(context.Background()).CompanyId(companyId).CompanyIds(companyIds).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListCompanyOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanyOverrides`: ListCompanyOverridesResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListCompanyOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompanyOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **companyIds** | **[]string** |  | 
 **featureId** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
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


## ListFeatureCompanies

> ListFeatureCompaniesResponse ListFeatureCompanies(ctx).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()

List feature companies

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

	featureId := "featureId_example" // string | 
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListFeatureCompanies(context.Background()).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListFeatureCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureCompanies`: ListFeatureCompaniesResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListFeatureCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFeatureCompaniesResponse**](ListFeatureCompaniesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureUsage

> ListFeatureUsageResponse ListFeatureUsage(ctx).CompanyId(companyId).CompanyKeys(companyKeys).FeatureIds(featureIds).Q(q).Limit(limit).Offset(offset).Execute()

List feature usage

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
	companyKeys := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListFeatureUsage(context.Background()).CompanyId(companyId).CompanyKeys(companyKeys).FeatureIds(featureIds).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListFeatureUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureUsage`: ListFeatureUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListFeatureUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **companyKeys** | **map[string]string** |  | 
 **featureIds** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFeatureUsageResponse**](ListFeatureUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFeatureUsers

> ListFeatureUsersResponse ListFeatureUsers(ctx).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()

List feature users

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

	featureId := "featureId_example" // string | 
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListFeatureUsers(context.Background()).FeatureId(featureId).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListFeatureUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFeatureUsers`: ListFeatureUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListFeatureUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFeatureUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListFeatureUsersResponse**](ListFeatureUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlanEntitlements

> ListPlanEntitlementsResponse ListPlanEntitlements(ctx).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).PlanId(planId).PlanIds(planIds).Q(q).Limit(limit).Offset(offset).Execute()

List plan entitlements

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

	featureId := "featureId_example" // string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	planId := "planId_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListPlanEntitlements(context.Background()).FeatureId(featureId).FeatureIds(featureIds).Ids(ids).PlanId(planId).PlanIds(planIds).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListPlanEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlanEntitlements`: ListPlanEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListPlanEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlanEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featureId** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **ids** | **[]string** |  | 
 **planId** | **string** |  | 
 **planIds** | **[]string** |  | 
 **q** | **string** |  | 
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

> UpdateCompanyOverrideResponse UpdateCompanyOverride(ctx, companyOverrideId).UpdateCompanyOverrideRequestBody(updateCompanyOverrideRequestBody).Execute()

Update company override

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

	companyOverrideId := "companyOverrideId_example" // string | company_override_id
	updateCompanyOverrideRequestBody := *schematicapi.NewUpdateCompanyOverrideRequestBody("ValueType_example") // UpdateCompanyOverrideRequestBody | 

	resp, r, err := client.API().EntitlementsAPI.UpdateCompanyOverride(context.Background(), companyOverrideId).UpdateCompanyOverrideRequestBody(updateCompanyOverrideRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.UpdateCompanyOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCompanyOverride`: UpdateCompanyOverrideResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.UpdateCompanyOverride`: %v\n", resp)
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

> UpdatePlanEntitlementResponse UpdatePlanEntitlement(ctx, planEntitlementId).UpdatePlanEntitlementRequestBody(updatePlanEntitlementRequestBody).Execute()

Update plan entitlement

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

	planEntitlementId := "planEntitlementId_example" // string | plan_entitlement_id
	updatePlanEntitlementRequestBody := *schematicapi.NewUpdatePlanEntitlementRequestBody("ValueType_example") // UpdatePlanEntitlementRequestBody | 

	resp, r, err := client.API().EntitlementsAPI.UpdatePlanEntitlement(context.Background(), planEntitlementId).UpdatePlanEntitlementRequestBody(updatePlanEntitlementRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.UpdatePlanEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlanEntitlement`: UpdatePlanEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.UpdatePlanEntitlement`: %v\n", resp)
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

