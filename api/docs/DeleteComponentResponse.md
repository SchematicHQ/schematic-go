# DeleteComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**DeleteResponse**](DeleteResponse.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewDeleteComponentResponse

`func NewDeleteComponentResponse(data DeleteResponse, params map[string]interface{}, ) *DeleteComponentResponse`

NewDeleteComponentResponse instantiates a new DeleteComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteComponentResponseWithDefaults

`func NewDeleteComponentResponseWithDefaults() *DeleteComponentResponse`

NewDeleteComponentResponseWithDefaults instantiates a new DeleteComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeleteComponentResponse) GetData() DeleteResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteComponentResponse) GetDataOk() (*DeleteResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteComponentResponse) SetData(v DeleteResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *DeleteComponentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DeleteComponentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DeleteComponentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


