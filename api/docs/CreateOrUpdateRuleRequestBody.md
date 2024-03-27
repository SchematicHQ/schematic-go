# CreateOrUpdateRuleRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Priority** | **int32** |  | 
**RuleType** | Pointer to **NullableString** |  | [optional] 
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


### GetRuleType

`func (o *CreateOrUpdateRuleRequestBody) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *CreateOrUpdateRuleRequestBody) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *CreateOrUpdateRuleRequestBody) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *CreateOrUpdateRuleRequestBody) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### SetRuleTypeNil

`func (o *CreateOrUpdateRuleRequestBody) SetRuleTypeNil(b bool)`

 SetRuleTypeNil sets the value for RuleType to be an explicit nil

### UnsetRuleType
`func (o *CreateOrUpdateRuleRequestBody) UnsetRuleType()`

UnsetRuleType ensures that no value is present for RuleType, not even an explicit nil
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


