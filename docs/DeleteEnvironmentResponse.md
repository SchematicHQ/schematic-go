# DeleteEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DeleteResponse**](DeleteResponse.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewDeleteEnvironmentResponse

`func NewDeleteEnvironmentResponse(data DeleteResponse, params map[string]interface{}, ) *DeleteEnvironmentResponse`

NewDeleteEnvironmentResponse instantiates a new DeleteEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEnvironmentResponseWithDefaults

`func NewDeleteEnvironmentResponseWithDefaults() *DeleteEnvironmentResponse`

NewDeleteEnvironmentResponseWithDefaults instantiates a new DeleteEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeleteEnvironmentResponse) GetData() DeleteResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteEnvironmentResponse) GetDataOk() (*DeleteResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteEnvironmentResponse) SetData(v DeleteResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *DeleteEnvironmentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DeleteEnvironmentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DeleteEnvironmentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


