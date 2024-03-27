# GetEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**EnvironmentResponseData**](EnvironmentResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewGetEnvironmentResponse

`func NewGetEnvironmentResponse(data EnvironmentResponseData, params map[string]interface{}, ) *GetEnvironmentResponse`

NewGetEnvironmentResponse instantiates a new GetEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentResponseWithDefaults

`func NewGetEnvironmentResponseWithDefaults() *GetEnvironmentResponse`

NewGetEnvironmentResponseWithDefaults instantiates a new GetEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEnvironmentResponse) GetData() EnvironmentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEnvironmentResponse) GetDataOk() (*EnvironmentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEnvironmentResponse) SetData(v EnvironmentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetEnvironmentResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetEnvironmentResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetEnvironmentResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


