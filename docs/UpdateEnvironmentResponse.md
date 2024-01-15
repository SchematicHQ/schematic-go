# UpdateEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnvironmentResponseData**](EnvironmentResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewUpdateEnvironmentResponse

`func NewUpdateEnvironmentResponse(data EnvironmentResponseData, params map[string]interface{}, ) *UpdateEnvironmentResponse`

NewUpdateEnvironmentResponse instantiates a new UpdateEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentResponseWithDefaults

`func NewUpdateEnvironmentResponseWithDefaults() *UpdateEnvironmentResponse`

NewUpdateEnvironmentResponseWithDefaults instantiates a new UpdateEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateEnvironmentResponse) GetData() EnvironmentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateEnvironmentResponse) GetDataOk() (*EnvironmentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateEnvironmentResponse) SetData(v EnvironmentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *UpdateEnvironmentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateEnvironmentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UpdateEnvironmentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


