# \EventsApi

All URIs are relative to *https://api.schematichq.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountEventTypes**](EventsApi.md#CountEventTypes) | **Get** /event-types/count | Count event types
[**CountEvents**](EventsApi.md#CountEvents) | **Get** /events/count | Count events
[**CreateEvent**](EventsApi.md#CreateEvent) | **Post** /events | Create event
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /events/{event_id} | Get event
[**GetEventType**](EventsApi.md#GetEventType) | **Get** /event-types/{key} | Get event type
[**ListEventTypes**](EventsApi.md#ListEventTypes) | **Get** /event-types | List event types
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /events | List events
[**ListMetricCounts**](EventsApi.md#ListMetricCounts) | **Get** /metric-counts | List metric counts
[**ListMetricCountsHourly**](EventsApi.md#ListMetricCountsHourly) | **Get** /metric-counts-hourly | List metric counts hourly



## CountEventTypes

> CountEventTypesResponse CountEventTypes(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()

Count event types

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CountEventTypes(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CountEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountEventTypes`: CountEventTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.CountEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountEventTypesResponse**](CountEventTypesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountEvents

> CountEventsResponse CountEvents(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).FeatureId(featureId).EventSubtype(eventSubtype).Limit(limit).Offset(offset).Execute()

Count events

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    companyId := "companyId_example" // string |  (optional)
    userId := "userId_example" // string |  (optional)
    featureId := "featureId_example" // string |  (optional)
    eventSubtype := "eventSubtype_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CountEvents(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).FeatureId(featureId).EventSubtype(eventSubtype).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CountEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountEvents`: CountEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.CountEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **companyId** | **string** |  | 
 **userId** | **string** |  | 
 **featureId** | **string** |  | 
 **eventSubtype** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**CountEventsResponse**](CountEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvent

> CreateEventResponse CreateEvent(ctx).CreateEventRequestBody(createEventRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Create event

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    createEventRequestBody := *openapiclient.NewCreateEventRequestBody("EventType_example") // CreateEventRequestBody | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.CreateEvent(context.Background()).CreateEventRequestBody(createEventRequestBody).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.CreateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvent`: CreateEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.CreateEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventRequestBody** | [**CreateEventRequestBody**](CreateEventRequestBody.md) |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

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


## GetEvent

> GetEventResponse GetEvent(ctx, eventId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get event

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    eventId := "eventId_example" // string | event_id
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEvent(context.Background(), eventId).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: GetEventResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEvent`: %v\n", resp)
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

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

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


## GetEventType

> GetEventTypeResponse GetEventType(ctx, key).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()

Get event type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    key := "key_example" // string | key
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetEventType(context.Background(), key).XSchematicEnvironmentId(xSchematicEnvironmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetEventType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventType`: GetEventTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 

### Return type

[**GetEventTypeResponse**](GetEventTypeResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventTypes

> ListEventTypesResponse ListEventTypes(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()

List event types

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListEventTypes(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).Q(q).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventTypes`: ListEventTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **q** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListEventTypesResponse**](ListEventTypesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListEventsResponse ListEvents(ctx).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).FeatureId(featureId).EventSubtype(eventSubtype).Limit(limit).Offset(offset).Execute()

List events

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    companyId := "companyId_example" // string |  (optional)
    userId := "userId_example" // string |  (optional)
    featureId := "featureId_example" // string |  (optional)
    eventSubtype := "eventSubtype_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListEvents(context.Background()).XSchematicEnvironmentId(xSchematicEnvironmentId).CompanyId(companyId).UserId(userId).FeatureId(featureId).EventSubtype(eventSubtype).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: ListEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **companyId** | **string** |  | 
 **userId** | **string** |  | 
 **featureId** | **string** |  | 
 **eventSubtype** | **string** |  | 
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

> ListMetricCountsResponse ListMetricCounts(ctx).EventSubtype(eventSubtype).XSchematicEnvironmentId(xSchematicEnvironmentId).StartTime(startTime).EndTime(endTime).EventSubtypes(eventSubtypes).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Grouping(grouping).Execute()

List metric counts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    eventSubtype := "eventSubtype_example" // string | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    startTime := time.Now() // time.Time |  (optional)
    endTime := time.Now() // time.Time |  (optional)
    eventSubtypes := []string{"Inner_example"} // []string |  (optional)
    companyId := "companyId_example" // string |  (optional)
    userId := "userId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)
    grouping := "grouping_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListMetricCounts(context.Background()).EventSubtype(eventSubtype).XSchematicEnvironmentId(xSchematicEnvironmentId).StartTime(startTime).EndTime(endTime).EventSubtypes(eventSubtypes).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Grouping(grouping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListMetricCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetricCounts`: ListMetricCountsResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListMetricCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSubtype** | **string** |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **startTime** | **time.Time** |  | 
 **endTime** | **time.Time** |  | 
 **eventSubtypes** | **[]string** |  | 
 **companyId** | **string** |  | 
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


## ListMetricCountsHourly

> ListMetricCountsHourlyResponse ListMetricCountsHourly(ctx).EventSubtype(eventSubtype).XSchematicEnvironmentId(xSchematicEnvironmentId).StartTime(startTime).EndTime(endTime).EventSubtypes(eventSubtypes).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()

List metric counts hourly

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/SchematicHQ/schematic-go"
)

func main() {
    eventSubtype := "eventSubtype_example" // string | 
    xSchematicEnvironmentId := "xSchematicEnvironmentId_example" // string | If the request is made using an API key that is not environment-scoped, specify the environment using this header (optional)
    startTime := time.Now() // time.Time |  (optional)
    endTime := time.Now() // time.Time |  (optional)
    eventSubtypes := []string{"Inner_example"} // []string |  (optional)
    companyId := "companyId_example" // string |  (optional)
    userId := "userId_example" // string |  (optional)
    limit := int32(56) // int32 | Page limit (default 100) (optional)
    offset := int32(56) // int32 | Page offset (default 0) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListMetricCountsHourly(context.Background()).EventSubtype(eventSubtype).XSchematicEnvironmentId(xSchematicEnvironmentId).StartTime(startTime).EndTime(endTime).EventSubtypes(eventSubtypes).CompanyId(companyId).UserId(userId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListMetricCountsHourly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetricCountsHourly`: ListMetricCountsHourlyResponse
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListMetricCountsHourly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricCountsHourlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventSubtype** | **string** |  | 
 **xSchematicEnvironmentId** | **string** | If the request is made using an API key that is not environment-scoped, specify the environment using this header | 
 **startTime** | **time.Time** |  | 
 **endTime** | **time.Time** |  | 
 **eventSubtypes** | **[]string** |  | 
 **companyId** | **string** |  | 
 **userId** | **string** |  | 
 **limit** | **int32** | Page limit (default 100) | 
 **offset** | **int32** | Page offset (default 0) | 

### Return type

[**ListMetricCountsHourlyResponse**](ListMetricCountsHourlyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

