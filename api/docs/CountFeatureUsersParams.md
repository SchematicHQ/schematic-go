# CountFeatureUsersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewCountFeatureUsersParams

`func NewCountFeatureUsersParams() *CountFeatureUsersParams`

NewCountFeatureUsersParams instantiates a new CountFeatureUsersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountFeatureUsersParamsWithDefaults

`func NewCountFeatureUsersParamsWithDefaults() *CountFeatureUsersParams`

NewCountFeatureUsersParamsWithDefaults instantiates a new CountFeatureUsersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *CountFeatureUsersParams) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CountFeatureUsersParams) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CountFeatureUsersParams) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CountFeatureUsersParams) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetLimit

`func (o *CountFeatureUsersParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CountFeatureUsersParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CountFeatureUsersParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CountFeatureUsersParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CountFeatureUsersParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CountFeatureUsersParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CountFeatureUsersParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CountFeatureUsersParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *CountFeatureUsersParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *CountFeatureUsersParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *CountFeatureUsersParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *CountFeatureUsersParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


