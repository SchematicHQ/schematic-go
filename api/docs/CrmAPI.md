# \CrmAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCrmProducts**](CrmAPI.md#ListCrmProducts) | **Get** /crm/products | List crm products
[**UpsertCrmDeal**](CrmAPI.md#UpsertCrmDeal) | **Post** /crm/deals/upsert | Upsert crm deal
[**UpsertCrmProduct**](CrmAPI.md#UpsertCrmProduct) | **Post** /crm/products/upsert | Upsert crm product
[**UpsertDealLineItemAssociation**](CrmAPI.md#UpsertDealLineItemAssociation) | **Post** /crm/associations/deal-line-item | Upsert deal line item association
[**UpsertLineItem**](CrmAPI.md#UpsertLineItem) | **Post** /crm/deal-line-item/upsert | Upsert line item



## ListCrmProducts

> ListCrmProductsResponse ListCrmProducts(ctx).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()

List crm products

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

	resp, r, err := client.API().CrmAPI.ListCrmProducts(context.Background()).Ids(ids).Name(name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.ListCrmProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCrmProducts`: ListCrmProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.ListCrmProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCrmProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCrmProductsResponse**](ListCrmProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCrmDeal

> UpsertCrmDealResponse UpsertCrmDeal(ctx).CreateCrmDealRequestBody(createCrmDealRequestBody).Execute()

Upsert crm deal

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

	createCrmDealRequestBody := *schematicapi.NewCreateCrmDealRequestBody("CrmCompanyKey_example", "CrmType_example", "DealExternalId_example") // CreateCrmDealRequestBody | 

	resp, r, err := client.API().CrmAPI.UpsertCrmDeal(context.Background()).CreateCrmDealRequestBody(createCrmDealRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.UpsertCrmDeal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCrmDeal`: UpsertCrmDealResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.UpsertCrmDeal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCrmDealRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCrmDealRequestBody** | [**CreateCrmDealRequestBody**](CreateCrmDealRequestBody.md) |  | 

### Return type

[**UpsertCrmDealResponse**](UpsertCrmDealResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCrmProduct

> UpsertCrmProductResponse UpsertCrmProduct(ctx).CreateCrmProductRequestBody(createCrmProductRequestBody).Execute()

Upsert crm product

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

	createCrmProductRequestBody := *schematicapi.NewCreateCrmProductRequestBody("Currency_example", "Description_example", "ExternalId_example", "Interval_example", "Name_example", float32(123), int32(123), "Sku_example") // CreateCrmProductRequestBody | 

	resp, r, err := client.API().CrmAPI.UpsertCrmProduct(context.Background()).CreateCrmProductRequestBody(createCrmProductRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.UpsertCrmProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCrmProduct`: UpsertCrmProductResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.UpsertCrmProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCrmProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCrmProductRequestBody** | [**CreateCrmProductRequestBody**](CreateCrmProductRequestBody.md) |  | 

### Return type

[**UpsertCrmProductResponse**](UpsertCrmProductResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertDealLineItemAssociation

> UpsertDealLineItemAssociationResponse UpsertDealLineItemAssociation(ctx).CreateCrmDealLineItemAssociationRequestBody(createCrmDealLineItemAssociationRequestBody).Execute()

Upsert deal line item association

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

	createCrmDealLineItemAssociationRequestBody := *schematicapi.NewCreateCrmDealLineItemAssociationRequestBody("DealExternalId_example", "LineItemExternalId_example") // CreateCrmDealLineItemAssociationRequestBody | 

	resp, r, err := client.API().CrmAPI.UpsertDealLineItemAssociation(context.Background()).CreateCrmDealLineItemAssociationRequestBody(createCrmDealLineItemAssociationRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.UpsertDealLineItemAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertDealLineItemAssociation`: UpsertDealLineItemAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.UpsertDealLineItemAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertDealLineItemAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCrmDealLineItemAssociationRequestBody** | [**CreateCrmDealLineItemAssociationRequestBody**](CreateCrmDealLineItemAssociationRequestBody.md) |  | 

### Return type

[**UpsertDealLineItemAssociationResponse**](UpsertDealLineItemAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertLineItem

> UpsertLineItemResponse UpsertLineItem(ctx).CreateCrmLineItemRequestBody(createCrmLineItemRequestBody).Execute()

Upsert line item

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

	createCrmLineItemRequestBody := *schematicapi.NewCreateCrmLineItemRequestBody(float32(123), "Interval_example", "LineItemExternalId_example", "ProductExternalId_example", int32(123)) // CreateCrmLineItemRequestBody | 

	resp, r, err := client.API().CrmAPI.UpsertLineItem(context.Background()).CreateCrmLineItemRequestBody(createCrmLineItemRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.UpsertLineItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertLineItem`: UpsertLineItemResponse
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.UpsertLineItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCrmLineItemRequestBody** | [**CreateCrmLineItemRequestBody**](CreateCrmLineItemRequestBody.md) |  | 

### Return type

[**UpsertLineItemResponse**](UpsertLineItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

