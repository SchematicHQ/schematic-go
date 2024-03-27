# \EntitlementsAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompanyOverride**](EntitlementsAPI.md#CreateCompanyOverride) | **Post** /company-overrides | Create company override
[**CreatePlanEntitlement**](EntitlementsAPI.md#CreatePlanEntitlement) | **Post** /plan-entitlements | Create plan entitlement
[**DeleteCompanyOverride**](EntitlementsAPI.md#DeleteCompanyOverride) | **Delete** /company-overrides/{company_override_id} | Delete company override
[**DeletePlanEntitlement**](EntitlementsAPI.md#DeletePlanEntitlement) | **Delete** /plan-entitlements/{plan_entitlement_id} | Delete plan entitlement
[**GetCompanyOverride**](EntitlementsAPI.md#GetCompanyOverride) | **Get** /company-overrides/{company_override_id} | Get company override
[**GetPlanEntitlement**](EntitlementsAPI.md#GetPlanEntitlement) | **Get** /plan-entitlements/{plan_entitlement_id} | Get plan entitlement
[**ListCompanyOverrides**](EntitlementsAPI.md#ListCompanyOverrides) | **Get** /company-overrides | List company overrides
[**ListPlanEntitlements**](EntitlementsAPI.md#ListPlanEntitlements) | **Get** /plan-entitlements | List plan entitlements
[**UpdateCompanyOverride**](EntitlementsAPI.md#UpdateCompanyOverride) | **Put** /company-overrides/{company_override_id} | Update company override
[**UpdatePlanEntitlement**](EntitlementsAPI.md#UpdatePlanEntitlement) | **Put** /plan-entitlements/{plan_entitlement_id} | Update plan entitlement



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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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

> ListCompanyOverridesResponse ListCompanyOverrides(ctx).CompanyId(companyId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()

List company overrides

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
	featureId := "featureId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListCompanyOverrides(context.Background()).CompanyId(companyId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()
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

> ListPlanEntitlementsResponse ListPlanEntitlements(ctx).PlanId(planId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()

List plan entitlements

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

	planId := "planId_example" // string |  (optional)
	featureId := "featureId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EntitlementsAPI.ListPlanEntitlements(context.Background()).PlanId(planId).FeatureId(featureId).Limit(limit).Offset(offset).Execute()
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

> UpdateCompanyOverrideResponse UpdateCompanyOverride(ctx, companyOverrideId).UpdateCompanyOverrideRequestBody(updateCompanyOverrideRequestBody).Execute()

Update company override

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
	schematicapi "github.com/schematichq/schematic-go/api"
	"github.com/schematichq/schematic-go"
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

