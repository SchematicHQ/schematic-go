# ListEnvironmentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EnvironmentResponseData**](EnvironmentResponseData.md) | The returned resources | 
**Params** | [**ListEnvironmentsParams**](ListEnvironmentsParams.md) |  | 

## Methods

### NewListEnvironmentsResponse

`func NewListEnvironmentsResponse(data []EnvironmentResponseData, params ListEnvironmentsParams, ) *ListEnvironmentsResponse`

NewListEnvironmentsResponse instantiates a new ListEnvironmentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironmentsResponseWithDefaults

`func NewListEnvironmentsResponseWithDefaults() *ListEnvironmentsResponse`

NewListEnvironmentsResponseWithDefaults instantiates a new ListEnvironmentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEnvironmentsResponse) GetData() []EnvironmentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEnvironmentsResponse) GetDataOk() (*[]EnvironmentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEnvironmentsResponse) SetData(v []EnvironmentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListEnvironmentsResponse) GetParams() ListEnvironmentsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListEnvironmentsResponse) GetParamsOk() (*ListEnvironmentsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListEnvironmentsResponse) SetParams(v ListEnvironmentsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


