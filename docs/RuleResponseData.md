# RuleResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**Priority** | **int32** |  | 
**PriorityGroup** | Pointer to **NullableInt32** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**Value** | **bool** |  | 

## Methods

### NewRuleResponseData

`func NewRuleResponseData(createdAt time.Time, environmentId string, id string, name string, priority int32, updatedAt time.Time, value bool, ) *RuleResponseData`

NewRuleResponseData instantiates a new RuleResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleResponseDataWithDefaults

`func NewRuleResponseDataWithDefaults() *RuleResponseData`

NewRuleResponseDataWithDefaults instantiates a new RuleResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RuleResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *RuleResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RuleResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RuleResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFlagId

`func (o *RuleResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *RuleResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *RuleResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *RuleResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *RuleResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *RuleResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *RuleResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RuleResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanId

`func (o *RuleResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RuleResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RuleResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RuleResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RuleResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RuleResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetPriority

`func (o *RuleResponseData) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RuleResponseData) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RuleResponseData) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetPriorityGroup

`func (o *RuleResponseData) GetPriorityGroup() int32`

GetPriorityGroup returns the PriorityGroup field if non-nil, zero value otherwise.

### GetPriorityGroupOk

`func (o *RuleResponseData) GetPriorityGroupOk() (*int32, bool)`

GetPriorityGroupOk returns a tuple with the PriorityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityGroup

`func (o *RuleResponseData) SetPriorityGroup(v int32)`

SetPriorityGroup sets PriorityGroup field to given value.

### HasPriorityGroup

`func (o *RuleResponseData) HasPriorityGroup() bool`

HasPriorityGroup returns a boolean if a field has been set.

### SetPriorityGroupNil

`func (o *RuleResponseData) SetPriorityGroupNil(b bool)`

 SetPriorityGroupNil sets the value for PriorityGroup to be an explicit nil

### UnsetPriorityGroup
`func (o *RuleResponseData) UnsetPriorityGroup()`

UnsetPriorityGroup ensures that no value is present for PriorityGroup, not even an explicit nil
### GetUpdatedAt

`func (o *RuleResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *RuleResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


