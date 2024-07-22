# GetComponentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ComponentResponseData**](ComponentResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewGetComponentResponse

`func NewGetComponentResponse(data ComponentResponseData, params map[string]interface{}, ) *GetComponentResponse`

NewGetComponentResponse instantiates a new GetComponentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetComponentResponseWithDefaults

`func NewGetComponentResponseWithDefaults() *GetComponentResponse`

NewGetComponentResponseWithDefaults instantiates a new GetComponentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetComponentResponse) GetData() ComponentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetComponentResponse) GetDataOk() (*ComponentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetComponentResponse) SetData(v ComponentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetComponentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetComponentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetComponentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


