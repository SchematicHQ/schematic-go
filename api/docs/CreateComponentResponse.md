# CreateComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ComponentResponseData**](ComponentResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewCreateComponentResponse

`func NewCreateComponentResponse(data ComponentResponseData, params map[string]interface{}, ) *CreateComponentResponse`

NewCreateComponentResponse instantiates a new CreateComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComponentResponseWithDefaults

`func NewCreateComponentResponseWithDefaults() *CreateComponentResponse`

NewCreateComponentResponseWithDefaults instantiates a new CreateComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateComponentResponse) GetData() ComponentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateComponentResponse) GetDataOk() (*ComponentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateComponentResponse) SetData(v ComponentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *CreateComponentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateComponentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateComponentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


