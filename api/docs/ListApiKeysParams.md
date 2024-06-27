# ListApiKeysParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**RequireEnvironment** | Pointer to **bool** |  | [optional] 

## Methods

### NewListApiKeysParams

`func NewListApiKeysParams() *ListApiKeysParams`

NewListApiKeysParams instantiates a new ListApiKeysParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiKeysParamsWithDefaults

`func NewListApiKeysParamsWithDefaults() *ListApiKeysParams`

NewListApiKeysParamsWithDefaults instantiates a new ListApiKeysParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ListApiKeysParams) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListApiKeysParams) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListApiKeysParams) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ListApiKeysParams) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetLimit

`func (o *ListApiKeysParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListApiKeysParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListApiKeysParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListApiKeysParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListApiKeysParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListApiKeysParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListApiKeysParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListApiKeysParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRequireEnvironment

`func (o *ListApiKeysParams) GetRequireEnvironment() bool`

GetRequireEnvironment returns the RequireEnvironment field if non-nil, zero value otherwise.

### GetRequireEnvironmentOk

`func (o *ListApiKeysParams) GetRequireEnvironmentOk() (*bool, bool)`

GetRequireEnvironmentOk returns a tuple with the RequireEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEnvironment

`func (o *ListApiKeysParams) SetRequireEnvironment(v bool)`

SetRequireEnvironment sets RequireEnvironment field to given value.

### HasRequireEnvironment

`func (o *ListApiKeysParams) HasRequireEnvironment() bool`

HasRequireEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


