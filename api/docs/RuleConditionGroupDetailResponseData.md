# RuleConditionGroupDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]RuleConditionDetailResponseData**](RuleConditionDetailResponseData.md) |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**RuleId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRuleConditionGroupDetailResponseData

`func NewRuleConditionGroupDetailResponseData(conditions []RuleConditionDetailResponseData, createdAt time.Time, environmentId string, id string, ruleId string, updatedAt time.Time, ) *RuleConditionGroupDetailResponseData`

NewRuleConditionGroupDetailResponseData instantiates a new RuleConditionGroupDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConditionGroupDetailResponseDataWithDefaults

`func NewRuleConditionGroupDetailResponseDataWithDefaults() *RuleConditionGroupDetailResponseData`

NewRuleConditionGroupDetailResponseDataWithDefaults instantiates a new RuleConditionGroupDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *RuleConditionGroupDetailResponseData) GetConditions() []RuleConditionDetailResponseData`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RuleConditionGroupDetailResponseData) GetConditionsOk() (*[]RuleConditionDetailResponseData, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RuleConditionGroupDetailResponseData) SetConditions(v []RuleConditionDetailResponseData)`

SetConditions sets Conditions field to given value.


### GetCreatedAt

`func (o *RuleConditionGroupDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleConditionGroupDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleConditionGroupDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *RuleConditionGroupDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RuleConditionGroupDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RuleConditionGroupDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFlagId

`func (o *RuleConditionGroupDetailResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *RuleConditionGroupDetailResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *RuleConditionGroupDetailResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *RuleConditionGroupDetailResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *RuleConditionGroupDetailResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *RuleConditionGroupDetailResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *RuleConditionGroupDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleConditionGroupDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleConditionGroupDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetPlanId

`func (o *RuleConditionGroupDetailResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RuleConditionGroupDetailResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RuleConditionGroupDetailResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RuleConditionGroupDetailResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RuleConditionGroupDetailResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RuleConditionGroupDetailResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetRuleId

`func (o *RuleConditionGroupDetailResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleConditionGroupDetailResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleConditionGroupDetailResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetUpdatedAt

`func (o *RuleConditionGroupDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleConditionGroupDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleConditionGroupDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


