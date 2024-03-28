# RuleDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]RuleConditionGroupDetailResponseData**](RuleConditionGroupDetailResponseData.md) |  | 
**Conditions** | [**[]RuleConditionResponseData**](RuleConditionResponseData.md) |  | 
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

### NewRuleDetailResponseData

`func NewRuleDetailResponseData(conditionGroups []RuleConditionGroupDetailResponseData, conditions []RuleConditionResponseData, createdAt time.Time, environmentId string, id string, name string, priority int32, ruleType string, updatedAt time.Time, value bool, ) *RuleDetailResponseData`

NewRuleDetailResponseData instantiates a new RuleDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleDetailResponseDataWithDefaults

`func NewRuleDetailResponseDataWithDefaults() *RuleDetailResponseData`

NewRuleDetailResponseDataWithDefaults instantiates a new RuleDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *RuleDetailResponseData) GetConditionGroups() []RuleConditionGroupDetailResponseData`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *RuleDetailResponseData) GetConditionGroupsOk() (*[]RuleConditionGroupDetailResponseData, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *RuleDetailResponseData) SetConditionGroups(v []RuleConditionGroupDetailResponseData)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *RuleDetailResponseData) GetConditions() []RuleConditionResponseData`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleDetailResponseData) GetConditionsOk() (*[]RuleConditionResponseData, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleDetailResponseData) SetConditions(v []RuleConditionResponseData)`

SetConditions sets Conditions field to given value.


### GetCreatedAt

`func (o *RuleDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *RuleDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RuleDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RuleDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFlagId

`func (o *RuleDetailResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *RuleDetailResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *RuleDetailResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *RuleDetailResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *RuleDetailResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *RuleDetailResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *RuleDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RuleDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *RuleDetailResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RuleDetailResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RuleDetailResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RuleDetailResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RuleDetailResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RuleDetailResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPriority

`func (o *RuleDetailResponseData) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RuleDetailResponseData) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RuleDetailResponseData) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetRuleType

`func (o *RuleDetailResponseData) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *RuleDetailResponseData) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *RuleDetailResponseData) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetUpdatedAt

`func (o *RuleDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *RuleDetailResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleDetailResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleDetailResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


