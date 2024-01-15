# AudienceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**Limit** | Pointer to **NullableInt32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **NullableInt32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAudienceRequestBody

`func NewAudienceRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, ) *AudienceRequestBody`

NewAudienceRequestBody instantiates a new AudienceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceRequestBodyWithDefaults

`func NewAudienceRequestBodyWithDefaults() *AudienceRequestBody`

NewAudienceRequestBodyWithDefaults instantiates a new AudienceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *AudienceRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *AudienceRequestBody) GetConditionGroupsOk() (*[]CreateOrUpdateConditionGroupRequestBody, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *AudienceRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *AudienceRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AudienceRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AudienceRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetLimit

`func (o *AudienceRequestBody) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AudienceRequestBody) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AudienceRequestBody) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AudienceRequestBody) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *AudienceRequestBody) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AudienceRequestBody) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *AudienceRequestBody) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AudienceRequestBody) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AudienceRequestBody) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *AudienceRequestBody) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *AudienceRequestBody) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *AudienceRequestBody) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetQ

`func (o *AudienceRequestBody) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *AudienceRequestBody) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *AudienceRequestBody) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *AudienceRequestBody) HasQ() bool`

HasQ returns a boolean if a field has been set.

### SetQNil

`func (o *AudienceRequestBody) SetQNil(b bool)`

 SetQNil sets the value for Q to be an explicit nil

### UnsetQ
`func (o *AudienceRequestBody) UnsetQ()`

UnsetQ ensures that no value is present for Q, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


