# ListApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ApiKeyResponseData**](ApiKeyResponseData.md) | The returned resources | 
**Params** | [**ListApiKeysParams**](ListApiKeysParams.md) |  | 

## Methods

### NewListApiKeysResponse

`func NewListApiKeysResponse(data []ApiKeyResponseData, params ListApiKeysParams, ) *ListApiKeysResponse`

NewListApiKeysResponse instantiates a new ListApiKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysResponseWithDefaults

`func NewListApiKeysResponseWithDefaults() *ListApiKeysResponse`

NewListApiKeysResponseWithDefaults instantiates a new ListApiKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListApiKeysResponse) GetData() []ApiKeyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApiKeysResponse) GetDataOk() (*[]ApiKeyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApiKeysResponse) SetData(v []ApiKeyResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListApiKeysResponse) GetParams() ListApiKeysParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListApiKeysResponse) GetParamsOk() (*ListApiKeysParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListApiKeysResponse) SetParams(v ListApiKeysParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


