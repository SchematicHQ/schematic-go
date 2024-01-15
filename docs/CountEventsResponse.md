# CountEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CountResponse**](CountResponse.md) |  | 
**Params** | [**CountEventsParams**](CountEventsParams.md) |  | 

## Methods

### NewCountEventsResponse

`func NewCountEventsResponse(data CountResponse, params CountEventsParams, ) *CountEventsResponse`

NewCountEventsResponse instantiates a new CountEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountEventsResponseWithDefaults

`func NewCountEventsResponseWithDefaults() *CountEventsResponse`

NewCountEventsResponseWithDefaults instantiates a new CountEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CountEventsResponse) GetData() CountResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CountEventsResponse) GetDataOk() (*CountResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CountEventsResponse) SetData(v CountResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *CountEventsResponse) GetParams() CountEventsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CountEventsResponse) GetParamsOk() (*CountEventsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CountEventsResponse) SetParams(v CountEventsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


