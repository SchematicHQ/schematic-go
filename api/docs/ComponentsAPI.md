# \ComponentsAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountComponents**](ComponentsAPI.md#CountComponents) | **Get** /components/count | Count components
[**CreateComponent**](ComponentsAPI.md#CreateComponent) | **Post** /components | Create component
[**DeleteComponent**](ComponentsAPI.md#DeleteComponent) | **Delete** /components/{component_id} | Delete component
[**GetComponent**](ComponentsAPI.md#GetComponent) | **Get** /components/{component_id} | Get component
[**HydrateComponent**](ComponentsAPI.md#HydrateComponent) | **Get** /components/{component_id}/hydrate | Hydrate component
[**ListComponents**](ComponentsAPI.md#ListComponents) | **Get** /components | List components
[**UpdateComponent**](ComponentsAPI.md#UpdateComponent) | **Put** /components/{component_id} | Update component



## CountComponents

> CountComponentsResponse CountComponents(ctx).Q(q).Limit(limit).Offset(offset).Execute()

Count components

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

	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().ComponentsAPI.CountComponents(context.Background()).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.CountComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountComponents`: CountComponentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.CountComponents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountComponentsResponse**](CountComponentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComponent

> CreateComponentResponse CreateComponent(ctx).CreateComponentRequestBody(createComponentRequestBody).Execute()

Create component

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

	createComponentRequestBody := *schematicapi.NewCreateComponentRequestBody([]int32{int32(123)}, "EntityType_example", "Name_example") // CreateComponentRequestBody | 

	resp, r, err := client.API().ComponentsAPI.CreateComponent(context.Background()).CreateComponentRequestBody(createComponentRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.CreateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComponent`: CreateComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.CreateComponent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createComponentRequestBody** | [**CreateComponentRequestBody**](CreateComponentRequestBody.md) |  | 

### Return type

[**CreateComponentResponse**](CreateComponentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComponent

> DeleteComponentResponse DeleteComponent(ctx, componentId).Execute()

Delete component

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

	componentId := "componentId_example" // string | component_id

	resp, r, err := client.API().ComponentsAPI.DeleteComponent(context.Background(), componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeleteComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteComponent`: DeleteComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.DeleteComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | component_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteComponentResponse**](DeleteComponentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponent

> GetComponentResponse GetComponent(ctx, componentId).Execute()

Get component

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

	componentId := "componentId_example" // string | component_id

	resp, r, err := client.API().ComponentsAPI.GetComponent(context.Background(), componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponent`: GetComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | component_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetComponentResponse**](GetComponentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HydrateComponent

> HydrateComponentResponse HydrateComponent(ctx, componentId).Execute()

Hydrate component

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

	componentId := "componentId_example" // string | component_id

	resp, r, err := client.API().ComponentsAPI.HydrateComponent(context.Background(), componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.HydrateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HydrateComponent`: HydrateComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.HydrateComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | component_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHydrateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HydrateComponentResponse**](HydrateComponentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListComponents

> ListComponentsResponse ListComponents(ctx).Q(q).Limit(limit).Offset(offset).Execute()

List components

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

	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().ComponentsAPI.ListComponents(context.Background()).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.ListComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComponents`: ListComponentsResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.ListComponents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListComponentsResponse**](ListComponentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponent

> UpdateComponentResponse UpdateComponent(ctx, componentId).UpdateComponentRequestBody(updateComponentRequestBody).Execute()

Update component

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

	componentId := "componentId_example" // string | component_id
	updateComponentRequestBody := *schematicapi.NewUpdateComponentRequestBody() // UpdateComponentRequestBody | 

	resp, r, err := client.API().ComponentsAPI.UpdateComponent(context.Background(), componentId).UpdateComponentRequestBody(updateComponentRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.UpdateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComponent`: UpdateComponentResponse
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.UpdateComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | component_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateComponentRequestBody** | [**UpdateComponentRequestBody**](UpdateComponentRequestBody.md) |  | 

### Return type

[**UpdateComponentResponse**](UpdateComponentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

