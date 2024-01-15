# UpdateRuleRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**Name** | **string** |  | 
**Priority** | **int32** |  | 
**PriorityGroup** | Pointer to **NullableInt32** |  | [optional] 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewUpdateRuleRequestBody

`func NewUpdateRuleRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, name string, priority int32, value bool, ) *UpdateRuleRequestBody`

NewUpdateRuleRequestBody instantiates a new UpdateRuleRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRuleRequestBodyWithDefaults

`func NewUpdateRuleRequestBodyWithDefaults() *UpdateRuleRequestBody`

NewUpdateRuleRequestBodyWithDefaults instantiates a new UpdateRuleRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *UpdateRuleRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *UpdateRuleRequestBody) GetConditionGroupsOk() (*[]CreateOrUpdateConditionGroupRequestBody, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *UpdateRuleRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *UpdateRuleRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateRuleRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateRuleRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetName

`func (o *UpdateRuleRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRuleRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRuleRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *UpdateRuleRequestBody) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateRuleRequestBody) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateRuleRequestBody) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPriorityGroup

`func (o *UpdateRuleRequestBody) GetPriorityGroup() int32`

GetPriorityGroup returns the PriorityGroup field if non-nil, zero value otherwise.

### GetPriorityGroupOk

`func (o *UpdateRuleRequestBody) GetPriorityGroupOk() (*int32, bool)`

GetPriorityGroupOk returns a tuple with the PriorityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityGroup

`func (o *UpdateRuleRequestBody) SetPriorityGroup(v int32)`

SetPriorityGroup sets PriorityGroup field to given value.

### HasPriorityGroup

`func (o *UpdateRuleRequestBody) HasPriorityGroup() bool`

HasPriorityGroup returns a boolean if a field has been set.

### SetPriorityGroupNil

`func (o *UpdateRuleRequestBody) SetPriorityGroupNil(b bool)`

 SetPriorityGroupNil sets the value for PriorityGroup to be an explicit nil

### UnsetPriorityGroup
`func (o *UpdateRuleRequestBody) UnsetPriorityGroup()`

UnsetPriorityGroup ensures that no value is present for PriorityGroup, not even an explicit nil
### GetSkipWebhooks

`func (o *UpdateRuleRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *UpdateRuleRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *UpdateRuleRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *UpdateRuleRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *UpdateRuleRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *UpdateRuleRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
### GetValue

`func (o *UpdateRuleRequestBody) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateRuleRequestBody) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateRuleRequestBody) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


