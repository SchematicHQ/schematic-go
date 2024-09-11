# \BillingAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountBillingProducts**](BillingAPI.md#CountBillingProducts) | **Get** /billing/products/count | Count billing products
[**CountCustomers**](BillingAPI.md#CountCustomers) | **Get** /billing/customers/count | Count customers
[**ListBillingProducts**](BillingAPI.md#ListBillingProducts) | **Get** /billing/products | List billing products
[**ListCustomers**](BillingAPI.md#ListCustomers) | **Get** /billing/customers | List customers
[**ListInvoices**](BillingAPI.md#ListInvoices) | **Get** /billing/invoices | List invoices
[**ListPaymentMethods**](BillingAPI.md#ListPaymentMethods) | **Get** /billing/payment-methods | List payment methods
[**ListProductPrices**](BillingAPI.md#ListProductPrices) | **Get** /billing/product/prices | List product prices
[**UpsertBillingCustomer**](BillingAPI.md#UpsertBillingCustomer) | **Post** /billing/customer/upsert | Upsert billing customer
[**UpsertBillingPrice**](BillingAPI.md#UpsertBillingPrice) | **Post** /billing/price/upsert | Upsert billing price
[**UpsertBillingProduct**](BillingAPI.md#UpsertBillingProduct) | **Post** /billing/product/upsert | Upsert billing product
[**UpsertBillingSubscription**](BillingAPI.md#UpsertBillingSubscription) | **Post** /billing/subscription/upsert | Upsert billing subscription
[**UpsertInvoice**](BillingAPI.md#UpsertInvoice) | **Post** /billing/invoices | Upsert invoice
[**UpsertPaymentMethod**](BillingAPI.md#UpsertPaymentMethod) | **Post** /billing/payment-methods | Upsert payment method



## CountBillingProducts

> CountBillingProductsResponse CountBillingProducts(ctx).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()

Count billing products

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
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.CountBillingProducts(context.Background()).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CountBillingProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountBillingProducts`: CountBillingProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CountBillingProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountBillingProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountBillingProductsResponse**](CountBillingProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCustomers

> CountCustomersResponse CountCustomers(ctx).Name(name).FailedToImport(failedToImport).Q(q).Limit(limit).Offset(offset).Execute()

Count customers

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

	name := "name_example" // string |  (optional)
	failedToImport := true // bool |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.CountCustomers(context.Background()).Name(name).FailedToImport(failedToImport).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.CountCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCustomers`: CountCustomersResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.CountCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **failedToImport** | **bool** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountCustomersResponse**](CountCustomersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillingProducts

> ListBillingProductsResponse ListBillingProducts(ctx).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()

List billing products

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
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListBillingProducts(context.Background()).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListBillingProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBillingProducts`: ListBillingProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListBillingProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBillingProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListBillingProductsResponse**](ListBillingProductsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> ListCustomersResponse ListCustomers(ctx).Name(name).FailedToImport(failedToImport).Q(q).Limit(limit).Offset(offset).Execute()

List customers

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

	name := "name_example" // string |  (optional)
	failedToImport := true // bool |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListCustomers(context.Background()).Name(name).FailedToImport(failedToImport).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomers`: ListCustomersResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **failedToImport** | **bool** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListCustomersResponse**](ListCustomersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> ListInvoicesResponse ListInvoices(ctx).CustomerExternalId(customerExternalId).CompanyId(companyId).SubscriptionExternalId(subscriptionExternalId).Limit(limit).Offset(offset).Execute()

List invoices

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

	customerExternalId := "customerExternalId_example" // string | 
	companyId := "companyId_example" // string |  (optional)
	subscriptionExternalId := "subscriptionExternalId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListInvoices(context.Background()).CustomerExternalId(customerExternalId).CompanyId(companyId).SubscriptionExternalId(subscriptionExternalId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvoices`: ListInvoicesResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerExternalId** | **string** |  | 
 **companyId** | **string** |  | 
 **subscriptionExternalId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentMethods

> ListPaymentMethodsResponse ListPaymentMethods(ctx).CustomerExternalId(customerExternalId).CompanyId(companyId).InvoiceExternalId(invoiceExternalId).SubscriptionExternalId(subscriptionExternalId).Limit(limit).Offset(offset).Execute()

List payment methods

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

	customerExternalId := "customerExternalId_example" // string | 
	companyId := "companyId_example" // string |  (optional)
	invoiceExternalId := "invoiceExternalId_example" // string |  (optional)
	subscriptionExternalId := "subscriptionExternalId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListPaymentMethods(context.Background()).CustomerExternalId(customerExternalId).CompanyId(companyId).InvoiceExternalId(invoiceExternalId).SubscriptionExternalId(subscriptionExternalId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentMethods`: ListPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerExternalId** | **string** |  | 
 **companyId** | **string** |  | 
 **invoiceExternalId** | **string** |  | 
 **subscriptionExternalId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListPaymentMethodsResponse**](ListPaymentMethodsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProductPrices

> ListProductPricesResponse ListProductPrices(ctx).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()

List product prices

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
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().BillingAPI.ListProductPrices(context.Background()).Ids(ids).Name(name).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.ListProductPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProductPrices`: ListProductPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.ListProductPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProductPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** |  | 
 **name** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListProductPricesResponse**](ListProductPricesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBillingCustomer

> UpsertBillingCustomerResponse UpsertBillingCustomer(ctx).CreateBillingCustomerRequestBody(createBillingCustomerRequestBody).Execute()

Upsert billing customer

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

	createBillingCustomerRequestBody := *schematicapi.NewCreateBillingCustomerRequestBody("Email_example", "ExternalId_example", false, map[string]string{"key": "Inner_example"}, "Name_example") // CreateBillingCustomerRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertBillingCustomer(context.Background()).CreateBillingCustomerRequestBody(createBillingCustomerRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertBillingCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingCustomer`: UpsertBillingCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertBillingCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingCustomerRequestBody** | [**CreateBillingCustomerRequestBody**](CreateBillingCustomerRequestBody.md) |  | 

### Return type

[**UpsertBillingCustomerResponse**](UpsertBillingCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertBillingPrice

> UpsertBillingPriceResponse UpsertBillingPrice(ctx).CreateBillingPriceRequestBody(createBillingPriceRequestBody).Execute()

Upsert billing price

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

	createBillingPriceRequestBody := *schematicapi.NewCreateBillingPriceRequestBody("Interval_example", int32(123), "PriceExternalId_example", "ProductExternalId_example") // CreateBillingPriceRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertBillingPrice(context.Background()).CreateBillingPriceRequestBody(createBillingPriceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertBillingPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertBillingPrice`: UpsertBillingPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertBillingPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertBillingPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingPriceRequestBody** | [**CreateBillingPriceRequestBody**](CreateBillingPriceRequestBody.md) |  | 

### Return type

[**UpsertBillingPriceResponse**](UpsertBillingPriceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

	createBillingSubscriptionsRequestBody := *schematicapi.NewCreateBillingSubscriptionsRequestBody("CustomerExternalId_example", time.Now(), []schematicapi.BillingProductPricing{*schematicapi.NewBillingProductPricing(int32(123), "ProductExternalId_example")}, "SubscriptionExternalId_example", int32(123)) // CreateBillingSubscriptionsRequestBody | 

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


## UpsertInvoice

> UpsertInvoiceResponse UpsertInvoice(ctx).CreateInvoiceRequestBody(createInvoiceRequestBody).Execute()

Upsert invoice

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

	createInvoiceRequestBody := *schematicapi.NewCreateInvoiceRequestBody(int32(123), int32(123), int32(123), "CollectionMethod_example", "Currency_example", "CustomerExternalId_example", "ExternalId_example", int32(123)) // CreateInvoiceRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertInvoice(context.Background()).CreateInvoiceRequestBody(createInvoiceRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertInvoice`: UpsertInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInvoiceRequestBody** | [**CreateInvoiceRequestBody**](CreateInvoiceRequestBody.md) |  | 

### Return type

[**UpsertInvoiceResponse**](UpsertInvoiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertPaymentMethod

> UpsertPaymentMethodResponse UpsertPaymentMethod(ctx).CreatePaymentMethodRequestBody(createPaymentMethodRequestBody).Execute()

Upsert payment method

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

	createPaymentMethodRequestBody := *schematicapi.NewCreatePaymentMethodRequestBody("CustomerExternalId_example", "ExternalId_example", "PaymentMethodType_example") // CreatePaymentMethodRequestBody | 

	resp, r, err := client.API().BillingAPI.UpsertPaymentMethod(context.Background()).CreatePaymentMethodRequestBody(createPaymentMethodRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpsertPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertPaymentMethod`: UpsertPaymentMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpsertPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentMethodRequestBody** | [**CreatePaymentMethodRequestBody**](CreatePaymentMethodRequestBody.md) |  | 

### Return type

[**UpsertPaymentMethodResponse**](UpsertPaymentMethodResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

