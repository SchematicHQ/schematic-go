# ListEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EventListResponseData**](EventListResponseData.md) | The returned resources | 
**Params** | [**ListEventsParams**](ListEventsParams.md) |  | 

## Methods

### NewListEventsResponse

`func NewListEventsResponse(data []EventListResponseData, params ListEventsParams, ) *ListEventsResponse`

NewListEventsResponse instantiates a new ListEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsResponseWithDefaults

`func NewListEventsResponseWithDefaults() *ListEventsResponse`

NewListEventsResponseWithDefaults instantiates a new ListEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEventsResponse) GetData() []EventListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEventsResponse) GetDataOk() (*[]EventListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEventsResponse) SetData(v []EventListResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListEventsResponse) GetParams() ListEventsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListEventsResponse) GetParamsOk() (*ListEventsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListEventsResponse) SetParams(v ListEventsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


