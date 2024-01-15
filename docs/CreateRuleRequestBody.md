# CreateRuleRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**Priority** | **int32** |  | 
**PriorityGroup** | Pointer to **NullableInt32** |  | [optional] 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewCreateRuleRequestBody

`func NewCreateRuleRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, name string, priority int32, value bool, ) *CreateRuleRequestBody`

NewCreateRuleRequestBody instantiates a new CreateRuleRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleRequestBodyWithDefaults

`func NewCreateRuleRequestBodyWithDefaults() *CreateRuleRequestBody`

NewCreateRuleRequestBodyWithDefaults instantiates a new CreateRuleRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *CreateRuleRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *CreateRuleRequestBody) GetConditionGroupsOk() (*[]CreateOrUpdateConditionGroupRequestBody, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *CreateRuleRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *CreateRuleRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateRuleRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateRuleRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetFlagId

`func (o *CreateRuleRequestBody) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *CreateRuleRequestBody) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *CreateRuleRequestBody) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *CreateRuleRequestBody) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *CreateRuleRequestBody) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *CreateRuleRequestBody) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetName

`func (o *CreateRuleRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRuleRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRuleRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *CreateRuleRequestBody) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateRuleRequestBody) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateRuleRequestBody) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CreateRuleRequestBody) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *CreateRuleRequestBody) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *CreateRuleRequestBody) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPriority

`func (o *CreateRuleRequestBody) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateRuleRequestBody) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateRuleRequestBody) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPriorityGroup

`func (o *CreateRuleRequestBody) GetPriorityGroup() int32`

GetPriorityGroup returns the PriorityGroup field if non-nil, zero value otherwise.

### GetPriorityGroupOk

`func (o *CreateRuleRequestBody) GetPriorityGroupOk() (*int32, bool)`

GetPriorityGroupOk returns a tuple with the PriorityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityGroup

`func (o *CreateRuleRequestBody) SetPriorityGroup(v int32)`

SetPriorityGroup sets PriorityGroup field to given value.

### HasPriorityGroup

`func (o *CreateRuleRequestBody) HasPriorityGroup() bool`

HasPriorityGroup returns a boolean if a field has been set.

### SetPriorityGroupNil

`func (o *CreateRuleRequestBody) SetPriorityGroupNil(b bool)`

 SetPriorityGroupNil sets the value for PriorityGroup to be an explicit nil

### UnsetPriorityGroup
`func (o *CreateRuleRequestBody) UnsetPriorityGroup()`

UnsetPriorityGroup ensures that no value is present for PriorityGroup, not even an explicit nil
### GetSkipWebhooks

`func (o *CreateRuleRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *CreateRuleRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *CreateRuleRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *CreateRuleRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *CreateRuleRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *CreateRuleRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
### GetValue

`func (o *CreateRuleRequestBody) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateRuleRequestBody) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateRuleRequestBody) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


