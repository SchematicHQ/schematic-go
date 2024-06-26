# \BillingAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProducts**](BillingAPI.md#ListProducts) | **Get** /billing/products | List products
[**UpsertBillingProduct**](BillingAPI.md#UpsertBillingProduct) | **Post** /billing/product/upsert | Upsert billing product
[**UpsertBillingSubscription**](BillingAPI.md#UpsertBillingSubscription) | **Post** /billing/subscription/upsert | Upsert billing subscription



## ListProducts

> ListProductsResponse ListProducts(ctx).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()

List products

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
	name := "name_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListProducts(context.Background()).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProducts`: ListProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListProductsResponse**](ListProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBillingProduct

> UpsertBillingProductResponse UpsertBillingProduct(ctx).CreateBillingProductRequestBody(createBillingProductRequestBody).Execute()

Upsert billing product

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

	createBillingProductRequestBody := *schematicapi.NewCreateBillingProductRequestBody("Currency_example", "ExternalId_example", "Name_example", float32(123), int32(123)) // CreateBillingProductRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertBillingProduct(context.Background()).CreateBillingProductRequestBody(createBillingProductRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertBillingProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingProduct`: UpsertBillingProductResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertBillingProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingProductRequestBody** | [**CreateBillingProductRequestBody**](CreateBillingProductRequestBody.md) |  | 

### Return type

[**UpsertBillingProductResponse**](UpsertBillingProductResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBillingSubscription

> UpsertBillingSubscriptionResponse UpsertBillingSubscription(ctx).CreateBillingSubscriptionsRequestBody(createBillingSubscriptionsRequestBody).Execute()

Upsert billing subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	schematicapi "github.com/SchematicHQ/schematic-go/api"
	"github.com/SchematicHQ/schematic-go"
)

func main() {
	apiKey := os.Getenv("SCHEMATIC_API_KEY")
	client := schematic.NewClient(apiKey)
	defer client.Close()

	createBillingSubscriptionsRequestBody := *schematicapi.NewCreateBillingSubscriptionsRequestBody("CustomerExternalId_example", time.Now(), []string{"ProductExternalIds_example"}, "SubscriptionExternalId_example") // CreateBillingSubscriptionsRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertBillingSubscription(context.Background()).CreateBillingSubscriptionsRequestBody(createBillingSubscriptionsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertBillingSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingSubscription`: UpsertBillingSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertBillingSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingSubscriptionsRequestBody** | [**CreateBillingSubscriptionsRequestBody**](CreateBillingSubscriptionsRequestBody.md) |  | 

### Return type

[**UpsertBillingSubscriptionResponse**](UpsertBillingSubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

