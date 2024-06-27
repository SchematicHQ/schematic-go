# CountEntityKeyDefinitionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 

## Methods

### NewCountEntityKeyDefinitionsParams

`func NewCountEntityKeyDefinitionsParams() *CountEntityKeyDefinitionsParams`

NewCountEntityKeyDefinitionsParams instantiates a new CountEntityKeyDefinitionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountEntityKeyDefinitionsParamsWithDefaults

`func NewCountEntityKeyDefinitionsParamsWithDefaults() *CountEntityKeyDefinitionsParams`

NewCountEntityKeyDefinitionsParamsWithDefaults instantiates a new CountEntityKeyDefinitionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *CountEntityKeyDefinitionsParams) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CountEntityKeyDefinitionsParams) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CountEntityKeyDefinitionsParams) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *CountEntityKeyDefinitionsParams) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIds

`func (o *CountEntityKeyDefinitionsParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CountEntityKeyDefinitionsParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CountEntityKeyDefinitionsParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CountEntityKeyDefinitionsParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetKey

`func (o *CountEntityKeyDefinitionsParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CountEntityKeyDefinitionsParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CountEntityKeyDefinitionsParams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CountEntityKeyDefinitionsParams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *CountEntityKeyDefinitionsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CountEntityKeyDefinitionsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CountEntityKeyDefinitionsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CountEntityKeyDefinitionsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CountEntityKeyDefinitionsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CountEntityKeyDefinitionsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CountEntityKeyDefinitionsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CountEntityKeyDefinitionsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


