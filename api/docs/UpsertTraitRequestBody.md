# UpsertTraitRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incr** | Pointer to **NullableInt32** | Amount to increment the trait by (positive or negative) | [optional] 
**Keys** | **map[string]interface{}** | Key/value pairs too identify a company or user | 
**Set** | Pointer to **NullableString** | Value to set the trait to | [optional] 
**Trait** | **string** | DealName of the trait to update | 
**UpdateOnly** | Pointer to **NullableBool** | Unless this is set, the company or user will be created if it does not already exist | [optional] 

## Methods

### NewUpsertTraitRequestBody

`func NewUpsertTraitRequestBody(keys map[string]interface{}, trait string, ) *UpsertTraitRequestBody`

NewUpsertTraitRequestBody instantiates a new UpsertTraitRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertTraitRequestBodyWithDefaults

`func NewUpsertTraitRequestBodyWithDefaults() *UpsertTraitRequestBody`

NewUpsertTraitRequestBodyWithDefaults instantiates a new UpsertTraitRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncr

`func (o *UpsertTraitRequestBody) GetIncr() int32`

GetIncr returns the Incr field if non-nil, zero value otherwise.

### GetIncrOk

`func (o *UpsertTraitRequestBody) GetIncrOk() (*int32, bool)`

GetIncrOk returns a tuple with the Incr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncr

`func (o *UpsertTraitRequestBody) SetIncr(v int32)`

SetIncr sets Incr field to given value.

### HasIncr

`func (o *UpsertTraitRequestBody) HasIncr() bool`

HasIncr returns a boolean if a field has been set.

### SetIncrNil

`func (o *UpsertTraitRequestBody) SetIncrNil(b bool)`

 SetIncrNil sets the value for Incr to be an explicit nil

### UnsetIncr
`func (o *UpsertTraitRequestBody) UnsetIncr()`

UnsetIncr ensures that no value is present for Incr, not even an explicit nil
### GetKeys

`func (o *UpsertTraitRequestBody) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UpsertTraitRequestBody) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UpsertTraitRequestBody) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetSet

`func (o *UpsertTraitRequestBody) GetSet() string`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *UpsertTraitRequestBody) GetSetOk() (*string, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *UpsertTraitRequestBody) SetSet(v string)`

SetSet sets Set field to given value.

### HasSet

`func (o *UpsertTraitRequestBody) HasSet() bool`

HasSet returns a boolean if a field has been set.

### SetSetNil

`func (o *UpsertTraitRequestBody) SetSetNil(b bool)`

 SetSetNil sets the value for Set to be an explicit nil

### UnsetSet
`func (o *UpsertTraitRequestBody) UnsetSet()`

UnsetSet ensures that no value is present for Set, not even an explicit nil
### GetTrait

`func (o *UpsertTraitRequestBody) GetTrait() string`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *UpsertTraitRequestBody) GetTraitOk() (*string, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *UpsertTraitRequestBody) SetTrait(v string)`

SetTrait sets Trait field to given value.


### GetUpdateOnly

`func (o *UpsertTraitRequestBody) GetUpdateOnly() bool`

GetUpdateOnly returns the UpdateOnly field if non-nil, zero value otherwise.

### GetUpdateOnlyOk

`func (o *UpsertTraitRequestBody) GetUpdateOnlyOk() (*bool, bool)`

GetUpdateOnlyOk returns a tuple with the UpdateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnly

`func (o *UpsertTraitRequestBody) SetUpdateOnly(v bool)`

SetUpdateOnly sets UpdateOnly field to given value.

### HasUpdateOnly

`func (o *UpsertTraitRequestBody) HasUpdateOnly() bool`

HasUpdateOnly returns a boolean if a field has been set.

### SetUpdateOnlyNil

`func (o *UpsertTraitRequestBody) SetUpdateOnlyNil(b bool)`

 SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil

### UnsetUpdateOnly
`func (o *UpsertTraitRequestBody) UnsetUpdateOnly()`

UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


