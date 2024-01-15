# CreateOrUpdateRuleRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Priority** | **int32** |  | 
**PriorityGroup** | Pointer to **NullableInt32** |  | [optional] 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewCreateOrUpdateRuleRequestBody

`func NewCreateOrUpdateRuleRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, name string, priority int32, value bool, ) *CreateOrUpdateRuleRequestBody`

NewCreateOrUpdateRuleRequestBody instantiates a new CreateOrUpdateRuleRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateRuleRequestBodyWithDefaults

`func NewCreateOrUpdateRuleRequestBodyWithDefaults() *CreateOrUpdateRuleRequestBody`

NewCreateOrUpdateRuleRequestBodyWithDefaults instantiates a new CreateOrUpdateRuleRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *CreateOrUpdateRuleRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *CreateOrUpdateRuleRequestBody) GetConditionGroupsOk() (*[]CreateOrUpdateConditionGroupRequestBody, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *CreateOrUpdateRuleRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *CreateOrUpdateRuleRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateOrUpdateRuleRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateOrUpdateRuleRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetId

`func (o *CreateOrUpdateRuleRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrUpdateRuleRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrUpdateRuleRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrUpdateRuleRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateOrUpdateRuleRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateOrUpdateRuleRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *CreateOrUpdateRuleRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateRuleRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateRuleRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *CreateOrUpdateRuleRequestBody) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateOrUpdateRuleRequestBody) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateOrUpdateRuleRequestBody) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPriorityGroup

`func (o *CreateOrUpdateRuleRequestBody) GetPriorityGroup() int32`

GetPriorityGroup returns the PriorityGroup field if non-nil, zero value otherwise.

### GetPriorityGroupOk

`func (o *CreateOrUpdateRuleRequestBody) GetPriorityGroupOk() (*int32, bool)`

GetPriorityGroupOk returns a tuple with the PriorityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityGroup

`func (o *CreateOrUpdateRuleRequestBody) SetPriorityGroup(v int32)`

SetPriorityGroup sets PriorityGroup field to given value.

### HasPriorityGroup

`func (o *CreateOrUpdateRuleRequestBody) HasPriorityGroup() bool`

HasPriorityGroup returns a boolean if a field has been set.

### SetPriorityGroupNil

`func (o *CreateOrUpdateRuleRequestBody) SetPriorityGroupNil(b bool)`

 SetPriorityGroupNil sets the value for PriorityGroup to be an explicit nil

### UnsetPriorityGroup
`func (o *CreateOrUpdateRuleRequestBody) UnsetPriorityGroup()`

UnsetPriorityGroup ensures that no value is present for PriorityGroup, not even an explicit nil
### GetSkipWebhooks

`func (o *CreateOrUpdateRuleRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *CreateOrUpdateRuleRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *CreateOrUpdateRuleRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *CreateOrUpdateRuleRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *CreateOrUpdateRuleRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *CreateOrUpdateRuleRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
### GetValue

`func (o *CreateOrUpdateRuleRequestBody) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateOrUpdateRuleRequestBody) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateOrUpdateRuleRequestBody) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


