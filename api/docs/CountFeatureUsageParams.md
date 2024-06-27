# CountFeatureUsageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyKeys** | Pointer to **map[string]string** |  | [optional] 
**FeatureIds** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewCountFeatureUsageParams

`func NewCountFeatureUsageParams() *CountFeatureUsageParams`

NewCountFeatureUsageParams instantiates a new CountFeatureUsageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountFeatureUsageParamsWithDefaults

`func NewCountFeatureUsageParamsWithDefaults() *CountFeatureUsageParams`

NewCountFeatureUsageParamsWithDefaults instantiates a new CountFeatureUsageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CountFeatureUsageParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CountFeatureUsageParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CountFeatureUsageParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CountFeatureUsageParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyKeys

`func (o *CountFeatureUsageParams) GetCompanyKeys() map[string]string`

GetCompanyKeys returns the CompanyKeys field if non-nil, zero value otherwise.

### GetCompanyKeysOk

`func (o *CountFeatureUsageParams) GetCompanyKeysOk() (*map[string]string, bool)`

GetCompanyKeysOk returns a tuple with the CompanyKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyKeys

`func (o *CountFeatureUsageParams) SetCompanyKeys(v map[string]string)`

SetCompanyKeys sets CompanyKeys field to given value.

### HasCompanyKeys

`func (o *CountFeatureUsageParams) HasCompanyKeys() bool`

HasCompanyKeys returns a boolean if a field has been set.

### GetFeatureIds

`func (o *CountFeatureUsageParams) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *CountFeatureUsageParams) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *CountFeatureUsageParams) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *CountFeatureUsageParams) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetLimit

`func (o *CountFeatureUsageParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CountFeatureUsageParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CountFeatureUsageParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CountFeatureUsageParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CountFeatureUsageParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CountFeatureUsageParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CountFeatureUsageParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CountFeatureUsageParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *CountFeatureUsageParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *CountFeatureUsageParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *CountFeatureUsageParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *CountFeatureUsageParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


