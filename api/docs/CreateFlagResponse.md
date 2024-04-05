# CreateFlagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**FlagDetailResponseData**](FlagDetailResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewCreateFlagResponse

`func NewCreateFlagResponse(data FlagDetailResponseData, params map[string]interface{}, ) *CreateFlagResponse`

NewCreateFlagResponse instantiates a new CreateFlagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlagResponseWithDefaults

`func NewCreateFlagResponseWithDefaults() *CreateFlagResponse`

NewCreateFlagResponseWithDefaults instantiates a new CreateFlagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateFlagResponse) GetData() FlagDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateFlagResponse) GetDataOk() (*FlagDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateFlagResponse) SetData(v FlagDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *CreateFlagResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CreateFlagResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CreateFlagResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


