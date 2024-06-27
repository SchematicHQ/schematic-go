# ListEntityKeyDefinitionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 

## Methods

### NewListEntityKeyDefinitionsParams

`func NewListEntityKeyDefinitionsParams() *ListEntityKeyDefinitionsParams`

NewListEntityKeyDefinitionsParams instantiates a new ListEntityKeyDefinitionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntityKeyDefinitionsParamsWithDefaults

`func NewListEntityKeyDefinitionsParamsWithDefaults() *ListEntityKeyDefinitionsParams`

NewListEntityKeyDefinitionsParamsWithDefaults instantiates a new ListEntityKeyDefinitionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *ListEntityKeyDefinitionsParams) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ListEntityKeyDefinitionsParams) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ListEntityKeyDefinitionsParams) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ListEntityKeyDefinitionsParams) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIds

`func (o *ListEntityKeyDefinitionsParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListEntityKeyDefinitionsParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListEntityKeyDefinitionsParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListEntityKeyDefinitionsParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetKey

`func (o *ListEntityKeyDefinitionsParams) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ListEntityKeyDefinitionsParams) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ListEntityKeyDefinitionsParams) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ListEntityKeyDefinitionsParams) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *ListEntityKeyDefinitionsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListEntityKeyDefinitionsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListEntityKeyDefinitionsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListEntityKeyDefinitionsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListEntityKeyDefinitionsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListEntityKeyDefinitionsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListEntityKeyDefinitionsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListEntityKeyDefinitionsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


