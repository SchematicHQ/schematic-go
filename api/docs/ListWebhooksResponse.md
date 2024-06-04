# ListWebhooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookResponseData**](WebhookResponseData.md) | The returned resources | 
**Params** | [**ListWebhooksParams**](ListWebhooksParams.md) |  | 

## Methods

### NewListWebhooksResponse

`func NewListWebhooksResponse(data []WebhookResponseData, params ListWebhooksParams, ) *ListWebhooksResponse`

NewListWebhooksResponse instantiates a new ListWebhooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksResponseWithDefaults

`func NewListWebhooksResponseWithDefaults() *ListWebhooksResponse`

NewListWebhooksResponseWithDefaults instantiates a new ListWebhooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListWebhooksResponse) GetData() []WebhookResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWebhooksResponse) GetDataOk() (*[]WebhookResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWebhooksResponse) SetData(v []WebhookResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListWebhooksResponse) GetParams() ListWebhooksParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListWebhooksResponse) GetParamsOk() (*ListWebhooksParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListWebhooksResponse) SetParams(v ListWebhooksParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


