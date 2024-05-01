# PlanAudienceDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]RuleConditionGroupDetailResponseData**](RuleConditionGroupDetailResponseData.md) |  | 
**Conditions** | [**[]RuleConditionDetailResponseData**](RuleConditionDetailResponseData.md) |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**Priority** | **int32** |  | 
**RuleType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Value** | **bool** |  | 

## Methods

### NewPlanAudienceDetailResponseData

`func NewPlanAudienceDetailResponseData(conditionGroups []RuleConditionGroupDetailResponseData, conditions []RuleConditionDetailResponseData, createdAt time.Time, environmentId string, id string, name string, priority int32, ruleType string, updatedAt time.Time, value bool, ) *PlanAudienceDetailResponseData`

NewPlanAudienceDetailResponseData instantiates a new PlanAudienceDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanAudienceDetailResponseDataWithDefaults

`func NewPlanAudienceDetailResponseDataWithDefaults() *PlanAudienceDetailResponseData`

NewPlanAudienceDetailResponseDataWithDefaults instantiates a new PlanAudienceDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *PlanAudienceDetailResponseData) GetConditionGroups() []RuleConditionGroupDetailResponseData`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *PlanAudienceDetailResponseData) GetConditionGroupsOk() (*[]RuleConditionGroupDetailResponseData, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *PlanAudienceDetailResponseData) SetConditionGroups(v []RuleConditionGroupDetailResponseData)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *PlanAudienceDetailResponseData) GetConditions() []RuleConditionDetailResponseData`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PlanAudienceDetailResponseData) GetConditionsOk() (*[]RuleConditionDetailResponseData, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PlanAudienceDetailResponseData) SetConditions(v []RuleConditionDetailResponseData)`

SetConditions sets Conditions field to given value.


### GetCreatedAt

`func (o *PlanAudienceDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanAudienceDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanAudienceDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *PlanAudienceDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PlanAudienceDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PlanAudienceDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFlagId

`func (o *PlanAudienceDetailResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *PlanAudienceDetailResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *PlanAudienceDetailResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *PlanAudienceDetailResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *PlanAudienceDetailResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *PlanAudienceDetailResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *PlanAudienceDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanAudienceDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanAudienceDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PlanAudienceDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanAudienceDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanAudienceDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *PlanAudienceDetailResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanAudienceDetailResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanAudienceDetailResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *PlanAudienceDetailResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *PlanAudienceDetailResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *PlanAudienceDetailResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPriority

`func (o *PlanAudienceDetailResponseData) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PlanAudienceDetailResponseData) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PlanAudienceDetailResponseData) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetRuleType

`func (o *PlanAudienceDetailResponseData) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *PlanAudienceDetailResponseData) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *PlanAudienceDetailResponseData) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetUpdatedAt

`func (o *PlanAudienceDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanAudienceDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanAudienceDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *PlanAudienceDetailResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlanAudienceDetailResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlanAudienceDetailResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


