# \EventsAPI

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventsAPI.md#CreateEvent) | **Post** /events | Create event
[**CreateEventBatch**](EventsAPI.md#CreateEventBatch) | **Post** /event-batch | Create event batch
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /events/{event_id} | Get event
[**GetEventSummaries**](EventsAPI.md#GetEventSummaries) | **Get** /event-types | Get event summaries
[**GetEventSummaryBySubtype**](EventsAPI.md#GetEventSummaryBySubtype) | **Get** /event-types/{key} | Get event summary by subtype
[**GetSegmentIntegrationStatus**](EventsAPI.md#GetSegmentIntegrationStatus) | **Get** /segment-integration | Get segment integration status
[**ListEvents**](EventsAPI.md#ListEvents) | **Get** /events | List events
[**ListMetricCounts**](EventsAPI.md#ListMetricCounts) | **Get** /metric-counts | List metric counts



## CreateEvent

> CreateEventResponse CreateEvent(ctx).CreateEventRequestBody(createEventRequestBody).Execute()

Create event

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

	createEventRequestBody := *schematicapi.NewCreateEventRequestBody("EventType_example") // CreateEventRequestBody | 

	resp, r, err := client.API().EventsAPI.CreateEvent(context.Background()).CreateEventRequestBody(createEventRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEvent`: CreateEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.CreateEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventRequestBody** | [**CreateEventRequestBody**](CreateEventRequestBody.md) |  | 

### Return type

[**CreateEventResponse**](CreateEventResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEventBatch

> CreateEventBatchResponse CreateEventBatch(ctx).CreateEventBatchRequestBody(createEventBatchRequestBody).Execute()

Create event batch

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

	createEventBatchRequestBody := *schematicapi.NewCreateEventBatchRequestBody([]schematicapi.CreateEventRequestBody{*schematicapi.NewCreateEventRequestBody("EventType_example")}) // CreateEventBatchRequestBody | 

	resp, r, err := client.API().EventsAPI.CreateEventBatch(context.Background()).CreateEventBatchRequestBody(createEventBatchRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateEventBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEventBatch`: CreateEventBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.CreateEventBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventBatchRequestBody** | [**CreateEventBatchRequestBody**](CreateEventBatchRequestBody.md) |  | 

### Return type

[**CreateEventBatchResponse**](CreateEventBatchResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> GetEventResponse GetEvent(ctx, eventId).Execute()

Get event

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

	eventId := "eventId_example" // string | event_id

	resp, r, err := client.API().EventsAPI.GetEvent(context.Background(), eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: GetEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | event_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventResponse**](GetEventResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSummaries

> GetEventSummariesResponse GetEventSummaries(ctx).Q(q).EventSubtypes(eventSubtypes).Limit(limit).Offset(offset).Execute()

Get event summaries

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
	eventSubtypes := []string{"Inner_example"} // []string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EventsAPI.GetEventSummaries(context.Background()).Q(q).EventSubtypes(eventSubtypes).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSummaries`: GetEventSummariesResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **eventSubtypes** | **[]string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**GetEventSummariesResponse**](GetEventSummariesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSummaryBySubtype

> GetEventSummaryBySubtypeResponse GetEventSummaryBySubtype(ctx, key).Execute()

Get event summary by subtype

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

	key := "key_example" // string | key

	resp, r, err := client.API().EventsAPI.GetEventSummaryBySubtype(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventSummaryBySubtype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventSummaryBySubtype`: GetEventSummaryBySubtypeResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventSummaryBySubtype`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventSummaryBySubtypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventSummaryBySubtypeResponse**](GetEventSummaryBySubtypeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentIntegrationStatus

> GetSegmentIntegrationStatusResponse GetSegmentIntegrationStatus(ctx).Execute()

Get segment integration status

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


	resp, r, err := client.API().EventsAPI.GetSegmentIntegrationStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetSegmentIntegrationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSegmentIntegrationStatus`: GetSegmentIntegrationStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetSegmentIntegrationStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentIntegrationStatusRequest struct via the builder pattern


### Return type

[**GetSegmentIntegrationStatusResponse**](GetSegmentIntegrationStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListEventsResponse ListEvents(ctx).CompanyId(companyId).EventSubtype(eventSubtype).EventTypes(eventTypes).FlagId(flagId).UserId(userId).Limit(limit).Offset(offset).Execute()

List events

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
	eventSubtype := "eventSubtype_example" // string |  (optional)
	eventTypes := []string{"Inner_example"} // []string |  (optional)
	flagId := "flagId_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)

	resp, r, err := client.API().EventsAPI.ListEvents(context.Background()).CompanyId(companyId).EventSubtype(eventSubtype).EventTypes(eventTypes).FlagId(flagId).UserId(userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: ListEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** |  | 
 **eventSubtype** | **string** |  | 
 **eventTypes** | **[]string** |  | 
 **flagId** | **string** |  | 
 **userId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListEventsResponse**](ListEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricCounts

> ListMetricCountsResponse ListMetricCounts(ctx).StartTime(startTime).EndTime(endTime).EventSubtype(eventSubtype).EventSubtypes(eventSubtypes).CompanyId(companyId).CompanyIds(companyIds).UserId(userId).Limit(limit).Offset(offset).Grouping(grouping).Execute()

List metric counts

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

	startTime := time.Now() // time.Time |  (optional)
	endTime := time.Now() // time.Time |  (optional)
	eventSubtype := "eventSubtype_example" // string |  (optional)
	eventSubtypes := []string{"Inner_example"} // []string |  (optional)
	companyId := "companyId_example" // string |  (optional)
	companyIds := []string{"Inner_example"} // []string |  (optional)
	userId := "userId_example" // string |  (optional)
	limit := int32(100) // int32 | Page limit (default 100) (optional)
	offset := int32(0) // int32 | Page offset (default 0) (optional)
	grouping := "grouping_example" // string |  (optional)

	resp, r, err := client.API().EventsAPI.ListMetricCounts(context.Background()).StartTime(startTime).EndTime(endTime).EventSubtype(eventSubtype).EventSubtypes(eventSubtypes).CompanyId(companyId).CompanyIds(companyIds).UserId(userId).Limit(limit).Offset(offset).Grouping(grouping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ListMetricCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetricCounts`: ListMetricCountsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ListMetricCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **time.Time** |  | 
 **endTime** | **time.Time** |  | 
 **eventSubtype** | **string** |  | 
 **eventSubtypes** | **[]string** |  | 
 **companyId** | **string** |  | 
 **companyIds** | **[]string** |  | 
 **userId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 
 **grouping** | **string** |  | 

### Return type

[**ListMetricCountsResponse**](ListMetricCountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

