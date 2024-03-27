# UpdateApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ApiKeyResponseData**](ApiKeyResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewUpdateApiKeyResponse

`func NewUpdateApiKeyResponse(data ApiKeyResponseData, params map[string]interface{}, ) *UpdateApiKeyResponse`

NewUpdateApiKeyResponse instantiates a new UpdateApiKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApiKeyResponseWithDefaults

`func NewUpdateApiKeyResponseWithDefaults() *UpdateApiKeyResponse`

NewUpdateApiKeyResponseWithDefaults instantiates a new UpdateApiKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UpdateApiKeyResponse) GetData() ApiKeyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateApiKeyResponse) GetDataOk() (*ApiKeyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateApiKeyResponse) SetData(v ApiKeyResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *UpdateApiKeyResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateApiKeyResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UpdateApiKeyResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


