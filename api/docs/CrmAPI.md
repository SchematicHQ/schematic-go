# \CrmAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCRMProducts**](CrmAPI.md#ListCRMProducts) | **Get** /crm/products | List c r m products
[**UpsertCRMProduct**](CrmAPI.md#UpsertCRMProduct) | **Post** /crm/products/upsert | Upsert c r m product



## ListCRMProducts

> ListCRMProductsResponse ListCRMProducts(ctx).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()

List c r m products

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

	resp, r, err := client.API().CrmAPI.ListCRMProducts(context.Background()).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.ListCRMProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCRMProducts`: ListCRMProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.ListCRMProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCRMProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCRMProductsResponse**](ListCRMProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCRMProduct

> UpsertCRMProductResponse UpsertCRMProduct(ctx).CreateBillingProductRequestBody(createBillingProductRequestBody).Execute()

Upsert c r m product

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

	createBillingProductRequestBody := *schematicapi.NewCreateBillingProductRequestBody("Currency_example", "ExternalId_example", "Interval_example", "Name_example", float32(123), int32(123)) // CreateBillingProductRequestBody | 

	resp, r, err := client.API().CrmAPI.UpsertCRMProduct(context.Background()).CreateBillingProductRequestBody(createBillingProductRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.UpsertCRMProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCRMProduct`: UpsertCRMProductResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.UpsertCRMProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCRMProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingProductRequestBody** | [**CreateBillingProductRequestBody**](CreateBillingProductRequestBody.md) |  | 

### Return type

[**UpsertCRMProductResponse**](UpsertCRMProductResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
