# ListWebhookEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookEventDetailResponseData**](WebhookEventDetailResponseData.md) | The returned resources | 
**Params** | [**ListWebhookEventsParams**](ListWebhookEventsParams.md) |  | 

## Methods

### NewListWebhookEventsResponse

`func NewListWebhookEventsResponse(data []WebhookEventDetailResponseData, params ListWebhookEventsParams, ) *ListWebhookEventsResponse`

NewListWebhookEventsResponse instantiates a new ListWebhookEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhookEventsResponseWithDefaults

`func NewListWebhookEventsResponseWithDefaults() *ListWebhookEventsResponse`

NewListWebhookEventsResponseWithDefaults instantiates a new ListWebhookEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListWebhookEventsResponse) GetData() []WebhookEventDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhookEventsResponse) GetDataOk() (*[]WebhookEventDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhookEventsResponse) SetData(v []WebhookEventDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListWebhookEventsResponse) GetParams() ListWebhookEventsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListWebhookEventsResponse) GetParamsOk() (*ListWebhookEventsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListWebhookEventsResponse) SetParams(v ListWebhookEventsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


