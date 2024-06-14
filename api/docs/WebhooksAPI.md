# \WebhooksAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWebhookEvents**](WebhooksAPI.md#CountWebhookEvents) | **Get** /webhook-events/count | Count webhook events
[**CountWebhooks**](WebhooksAPI.md#CountWebhooks) | **Get** /webhooks/count | Count webhooks
[**CreateWebhook**](WebhooksAPI.md#CreateWebhook) | **Post** /webhooks | Create webhook
[**DeleteWebhook**](WebhooksAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete webhook
[**GetWebhook**](WebhooksAPI.md#GetWebhook) | **Get** /webhooks/{webhook_id} | Get webhook
[**GetWebhookEvent**](WebhooksAPI.md#GetWebhookEvent) | **Get** /webhook-events/{webhook_event_id} | Get webhook event
[**ListWebhookEvents**](WebhooksAPI.md#ListWebhookEvents) | **Get** /webhook-events | List webhook events
[**ListWebhooks**](WebhooksAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**UpdateWebhook**](WebhooksAPI.md#UpdateWebhook) | **Put** /webhooks/{webhook_id} | Update webhook



## CountWebhookEvents

> CountWebhookEventsResponse CountWebhookEvents(ctx).WebhookId(webhookId).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

Count webhook events

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

	webhookId := "webhookId_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().WebhooksAPI.CountWebhookEvents(context.Background()).WebhookId(webhookId).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CountWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebhookEvents`: CountWebhookEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CountWebhookEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountWebhookEventsResponse**](CountWebhookEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountWebhooks

> CountWebhooksResponse CountWebhooks(ctx).Q(q).Limit(limit).Offset(offset).Execute()

Count webhooks

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

	resp, r, err := client.API().WebhooksAPI.CountWebhooks(context.Background()).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CountWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebhooks`: CountWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CountWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountWebhooksResponse**](CountWebhooksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> CreateWebhookResponse CreateWebhook(ctx).CreateWebhookRequestBody(createWebhookRequestBody).Execute()

Create webhook

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

	createWebhookRequestBody := *schematicapi.NewCreateWebhookRequestBody("Name_example", []string{"RequestTypes_example"}, "Url_example") // CreateWebhookRequestBody | 

	resp, r, err := client.API().WebhooksAPI.CreateWebhook(context.Background()).CreateWebhookRequestBody(createWebhookRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: CreateWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequestBody** | [**CreateWebhookRequestBody**](CreateWebhookRequestBody.md) |  | 

### Return type

[**CreateWebhookResponse**](CreateWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhookResponse DeleteWebhook(ctx, webhookId).Execute()

Delete webhook

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

	webhookId := "webhookId_example" // string | webhook_id

	resp, r, err := client.API().WebhooksAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: DeleteWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | webhook_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteWebhookResponse**](DeleteWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> GetWebhookResponse GetWebhook(ctx, webhookId).Execute()

Get webhook

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

	webhookId := "webhookId_example" // string | webhook_id

	resp, r, err := client.API().WebhooksAPI.GetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: GetWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | webhook_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWebhookResponse**](GetWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookEvent

> GetWebhookEventResponse GetWebhookEvent(ctx, webhookEventId).Execute()

Get webhook event

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

	webhookEventId := "webhookEventId_example" // string | webhook_event_id

	resp, r, err := client.API().WebhooksAPI.GetWebhookEvent(context.Background(), webhookEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookEvent`: GetWebhookEventResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookEventId** | **string** | webhook_event_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWebhookEventResponse**](GetWebhookEventResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookEvents

> ListWebhookEventsResponse ListWebhookEvents(ctx).WebhookId(webhookId).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()

List webhook events

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

	webhookId := "webhookId_example" // string |  (optional)
	ids := []string{"Inner_example"} // []string |  (optional)
	q := "q_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().WebhooksAPI.ListWebhookEvents(context.Background()).WebhookId(webhookId).Ids(ids).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhookEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookEvents`: ListWebhookEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhookEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookId** | **string** |  | 
 **ids** | **[]string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListWebhookEventsResponse**](ListWebhookEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> ListWebhooksResponse ListWebhooks(ctx).Q(q).Limit(limit).Offset(offset).Execute()

List webhooks

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

	resp, r, err := client.API().WebhooksAPI.ListWebhooks(context.Background()).Q(q).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: ListWebhooksResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListWebhooksResponse**](ListWebhooksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> UpdateWebhookResponse UpdateWebhook(ctx, webhookId).UpdateWebhookRequestBody(updateWebhookRequestBody).Execute()

Update webhook

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

	webhookId := "webhookId_example" // string | webhook_id
	updateWebhookRequestBody := *schematicapi.NewUpdateWebhookRequestBody() // UpdateWebhookRequestBody | 

	resp, r, err := client.API().WebhooksAPI.UpdateWebhook(context.Background(), webhookId).UpdateWebhookRequestBody(updateWebhookRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: UpdateWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | webhook_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookRequestBody** | [**UpdateWebhookRequestBody**](UpdateWebhookRequestBody.md) |  | 

### Return type

[**UpdateWebhookResponse**](UpdateWebhookResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

