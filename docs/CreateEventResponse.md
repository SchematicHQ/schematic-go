# CreateEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EventResponseData**](EventResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewCreateEventResponse

`func NewCreateEventResponse(data EventResponseData, params map[string]interface{}, ) *CreateEventResponse`

NewCreateEventResponse instantiates a new CreateEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventResponseWithDefaults

`func NewCreateEventResponseWithDefaults() *CreateEventResponse`

NewCreateEventResponseWithDefaults instantiates a new CreateEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateEventResponse) GetData() EventResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateEventResponse) GetDataOk() (*EventResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateEventResponse) SetData(v EventResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *CreateEventResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateEventResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateEventResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


