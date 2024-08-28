# \PlansAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountPlans**](PlansAPI.md#CountPlans) | **Get** /plans/count | Count plans
[**CreatePlan**](PlansAPI.md#CreatePlan) | **Post** /plans | Create plan
[**DeleteAudience**](PlansAPI.md#DeleteAudience) | **Delete** /plan-audiences/{plan_audience_id} | Delete audience
[**DeletePlan**](PlansAPI.md#DeletePlan) | **Delete** /plans/{plan_id} | Delete plan
[**GetAudience**](PlansAPI.md#GetAudience) | **Get** /plan-audiences/{plan_audience_id} | Get audience
[**GetPlan**](PlansAPI.md#GetPlan) | **Get** /plans/{plan_id} | Get plan
[**ListPlans**](PlansAPI.md#ListPlans) | **Get** /plans | List plans
[**UpdateAudience**](PlansAPI.md#UpdateAudience) | **Put** /plan-audiences/{plan_audience_id} | Update audience
[**UpdatePlan**](PlansAPI.md#UpdatePlan) | **Put** /plans/{plan_id} | Update plan
[**UpsertBillingProductPlan**](PlansAPI.md#UpsertBillingProductPlan) | **Put** /plans/{plan_id}/billing_products | Upsert billing product plan



## CountPlans

> CountPlansResponse CountPlans(ctx).CompanyId(companyId).Ids(ids).Q(q).PlanType(planType).HasProductId(hasProductId).WithoutEntitlementFor(withoutEntitlementFor).Limit(limit).Offset(offset).Execute()

Count plans

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
	q := "q_example" // string |  (optional)
	planType := "planType_example" // string | Filter by plan type (optional)
	hasProductId := true // bool | Filter out plans that do not have a billing product ID (optional)
	withoutEntitlementFor := "withoutEntitlementFor_example" // string | Filter out plans that already have a plan entitlement for the specified feature ID (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().PlansAPI.CountPlans(context.Background()).CompanyId(companyId).Ids(ids).Q(q).PlanType(planType).HasProductId(hasProductId).WithoutEntitlementFor(withoutEntitlementFor).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CountPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountPlans`: CountPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CountPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **planType** | **string** | Filter by plan type | 
 **hasProductId** | **bool** | Filter out plans that do not have a billing product ID | 
 **withoutEntitlementFor** | **string** | Filter out plans that already have a plan entitlement for the specified feature ID | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountPlansResponse**](CountPlansResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlan

> CreatePlanResponse CreatePlan(ctx).CreatePlanRequestBody(createPlanRequestBody).Execute()

Create plan

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

	createPlanRequestBody := *schematicapi.NewCreatePlanRequestBody("Description_example", "Name_example", "PlanType_example") // CreatePlanRequestBody | 

	resp, r, err := client.API().PlansAPI.CreatePlan(context.Background()).CreatePlanRequestBody(createPlanRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CreatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlan`: CreatePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanRequestBody** | [**CreatePlanRequestBody**](CreatePlanRequestBody.md) |  | 

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


## DeleteAudience

> DeleteAudienceResponse DeleteAudience(ctx, planAudienceId).Execute()

Delete audience

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

	planAudienceId := "planAudienceId_example" // string | plan_audience_id

	resp, r, err := client.API().PlansAPI.DeleteAudience(context.Background(), planAudienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.DeleteAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAudience`: DeleteAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.DeleteAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planAudienceId** | **string** | plan_audience_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAudienceResponse**](DeleteAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlanResponse DeletePlan(ctx, planId).Execute()

Delete plan

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

	planId := "planId_example" // string | plan_id

	resp, r, err := client.API().PlansAPI.DeletePlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.DeletePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlan`: DeletePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.DeletePlan`: %v\n", resp)
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


## GetAudience

> GetAudienceResponse GetAudience(ctx, planAudienceId).Execute()

Get audience

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

	planAudienceId := "planAudienceId_example" // string | plan_audience_id

	resp, r, err := client.API().PlansAPI.GetAudience(context.Background(), planAudienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudience`: GetAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planAudienceId** | **string** | plan_audience_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAudienceResponse**](GetAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlan

> GetPlanResponse GetPlan(ctx, planId).Execute()

Get plan

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

	planId := "planId_example" // string | plan_id

	resp, r, err := client.API().PlansAPI.GetPlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: GetPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlan`: %v\n", resp)
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

> ListPlansResponse ListPlans(ctx).CompanyId(companyId).Ids(ids).Q(q).PlanType(planType).HasProductId(hasProductId).WithoutEntitlementFor(withoutEntitlementFor).Limit(limit).Offset(offset).Execute()

List plans

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
	q := "q_example" // string |  (optional)
	planType := "planType_example" // string | Filter by plan type (optional)
	hasProductId := true // bool | Filter out plans that do not have a billing product ID (optional)
	withoutEntitlementFor := "withoutEntitlementFor_example" // string | Filter out plans that already have a plan entitlement for the specified feature ID (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().PlansAPI.ListPlans(context.Background()).CompanyId(companyId).Ids(ids).Q(q).PlanType(planType).HasProductId(hasProductId).WithoutEntitlementFor(withoutEntitlementFor).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.ListPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPlans`: ListPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.ListPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **planType** | **string** | Filter by plan type | 
 **hasProductId** | **bool** | Filter out plans that do not have a billing product ID | 
 **withoutEntitlementFor** | **string** | Filter out plans that already have a plan entitlement for the specified feature ID | 
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


## UpdateAudience

> UpdateAudienceResponse UpdateAudience(ctx, planAudienceId).UpdateAudienceRequestBody(updateAudienceRequestBody).Execute()

Update audience

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

	planAudienceId := "planAudienceId_example" // string | plan_audience_id
	updateAudienceRequestBody := *schematicapi.NewUpdateAudienceRequestBody([]schematicapi.CreateOrUpdateConditionGroupRequestBody{*schematicapi.NewCreateOrUpdateConditionGroupRequestBody([]schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", "Operator_example", []string{"ResourceIds_example"})})}, []schematicapi.CreateOrUpdateConditionRequestBody{*schematicapi.NewCreateOrUpdateConditionRequestBody("ConditionType_example", "Operator_example", []string{"ResourceIds_example"})}) // UpdateAudienceRequestBody | 

	resp, r, err := client.API().PlansAPI.UpdateAudience(context.Background(), planAudienceId).UpdateAudienceRequestBody(updateAudienceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.UpdateAudience``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAudience`: UpdateAudienceResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.UpdateAudience`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planAudienceId** | **string** | plan_audience_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAudienceRequestBody** | [**UpdateAudienceRequestBody**](UpdateAudienceRequestBody.md) |  | 

### Return type

[**UpdateAudienceResponse**](UpdateAudienceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> UpdatePlanResponse UpdatePlan(ctx, planId).UpdatePlanRequestBody(updatePlanRequestBody).Execute()

Update plan

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

	planId := "planId_example" // string | plan_id
	updatePlanRequestBody := *schematicapi.NewUpdatePlanRequestBody("Name_example") // UpdatePlanRequestBody | 

	resp, r, err := client.API().PlansAPI.UpdatePlan(context.Background(), planId).UpdatePlanRequestBody(updatePlanRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.UpdatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlan`: UpdatePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.UpdatePlan`: %v\n", resp)
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


## UpsertBillingProductPlan

> UpsertBillingProductPlanResponse UpsertBillingProductPlan(ctx, planId).UpsertBillingProductRequestBody(upsertBillingProductRequestBody).Execute()

Upsert billing product plan

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

	planId := "planId_example" // string | plan_id
	upsertBillingProductRequestBody := *schematicapi.NewUpsertBillingProductRequestBody("BillingProductId_example") // UpsertBillingProductRequestBody | 

	resp, r, err := client.API().PlansAPI.UpsertBillingProductPlan(context.Background(), planId).UpsertBillingProductRequestBody(upsertBillingProductRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.UpsertBillingProductPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingProductPlan`: UpsertBillingProductPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.UpsertBillingProductPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | plan_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingProductPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertBillingProductRequestBody** | [**UpsertBillingProductRequestBody**](UpsertBillingProductRequestBody.md) |  | 

### Return type

[**UpsertBillingProductPlanResponse**](UpsertBillingProductPlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

