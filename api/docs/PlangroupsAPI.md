# \PlangroupsAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlanGroup**](PlangroupsAPI.md#CreatePlanGroup) | **Post** /plan-groups | Create plan group
[**GetPlanGroup**](PlangroupsAPI.md#GetPlanGroup) | **Get** /plan-groups | Get plan group
[**UpdatePlanGroup**](PlangroupsAPI.md#UpdatePlanGroup) | **Put** /plan-groups/{plan_group_id} | Update plan group



## CreatePlanGroup

> CreatePlanGroupResponse CreatePlanGroup(ctx).CreatePlanGroupRequestBody(createPlanGroupRequestBody).Execute()

Create plan group

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

	createPlanGroupRequestBody := *schematicapi.NewCreatePlanGroupRequestBody([]string{"PlanIds_example"}) // CreatePlanGroupRequestBody | 

	resp, r, err := client.API().PlangroupsAPI.CreatePlanGroup(context.Background()).CreatePlanGroupRequestBody(createPlanGroupRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlangroupsAPI.CreatePlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanGroup`: CreatePlanGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `PlangroupsAPI.CreatePlanGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanGroupRequestBody** | [**CreatePlanGroupRequestBody**](CreatePlanGroupRequestBody.md) |  | 

### Return type

[**CreatePlanGroupResponse**](CreatePlanGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanGroup

> GetPlanGroupResponse GetPlanGroup(ctx).Execute()

Get plan group

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


	resp, r, err := client.API().PlangroupsAPI.GetPlanGroup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlangroupsAPI.GetPlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanGroup`: GetPlanGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `PlangroupsAPI.GetPlanGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanGroupRequest struct via the builder pattern


### Return type

[**GetPlanGroupResponse**](GetPlanGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanGroup

> UpdatePlanGroupResponse UpdatePlanGroup(ctx, planGroupId).UpdatePlanGroupRequestBody(updatePlanGroupRequestBody).Execute()

Update plan group

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

	planGroupId := "planGroupId_example" // string | plan_group_id
	updatePlanGroupRequestBody := *schematicapi.NewUpdatePlanGroupRequestBody([]string{"PlanIds_example"}) // UpdatePlanGroupRequestBody | 

	resp, r, err := client.API().PlangroupsAPI.UpdatePlanGroup(context.Background(), planGroupId).UpdatePlanGroupRequestBody(updatePlanGroupRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlangroupsAPI.UpdatePlanGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlanGroup`: UpdatePlanGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `PlangroupsAPI.UpdatePlanGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planGroupId** | **string** | plan_group_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlanGroupRequestBody** | [**UpdatePlanGroupRequestBody**](UpdatePlanGroupRequestBody.md) |  | 

### Return type

[**UpdatePlanGroupResponse**](UpdatePlanGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

