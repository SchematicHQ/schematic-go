# RuleConditionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonTraitId** | Pointer to **NullableString** |  | [optional] 
**ConditionGroupId** | Pointer to **NullableString** |  | [optional] 
**ConditionType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**MetricValue** | Pointer to **NullableInt32** |  | [optional] 
**Operator** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**ResourceIds** | **[]string** |  | 
**RuleId** | **string** |  | 
**TraitEntityType** | Pointer to **NullableString** |  | [optional] 
**TraitId** | Pointer to **NullableString** |  | [optional] 
**TraitValue** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRuleConditionResponseData

`func NewRuleConditionResponseData(conditionType string, createdAt time.Time, environmentId string, id string, operator string, resourceIds []string, ruleId string, traitValue string, updatedAt time.Time, ) *RuleConditionResponseData`

NewRuleConditionResponseData instantiates a new RuleConditionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConditionResponseDataWithDefaults

`func NewRuleConditionResponseDataWithDefaults() *RuleConditionResponseData`

NewRuleConditionResponseDataWithDefaults instantiates a new RuleConditionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonTraitId

`func (o *RuleConditionResponseData) GetComparisonTraitId() string`

GetComparisonTraitId returns the ComparisonTraitId field if non-nil, zero value otherwise.

### GetComparisonTraitIdOk

`func (o *RuleConditionResponseData) GetComparisonTraitIdOk() (*string, bool)`

GetComparisonTraitIdOk returns a tuple with the ComparisonTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonTraitId

`func (o *RuleConditionResponseData) SetComparisonTraitId(v string)`

SetComparisonTraitId sets ComparisonTraitId field to given value.

### HasComparisonTraitId

`func (o *RuleConditionResponseData) HasComparisonTraitId() bool`

HasComparisonTraitId returns a boolean if a field has been set.

### SetComparisonTraitIdNil

`func (o *RuleConditionResponseData) SetComparisonTraitIdNil(b bool)`

 SetComparisonTraitIdNil sets the value for ComparisonTraitId to be an explicit nil

### UnsetComparisonTraitId
`func (o *RuleConditionResponseData) UnsetComparisonTraitId()`

UnsetComparisonTraitId ensures that no value is present for ComparisonTraitId, not even an explicit nil
### GetConditionGroupId

`func (o *RuleConditionResponseData) GetConditionGroupId() string`

GetConditionGroupId returns the ConditionGroupId field if non-nil, zero value otherwise.

### GetConditionGroupIdOk

`func (o *RuleConditionResponseData) GetConditionGroupIdOk() (*string, bool)`

GetConditionGroupIdOk returns a tuple with the ConditionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroupId

`func (o *RuleConditionResponseData) SetConditionGroupId(v string)`

SetConditionGroupId sets ConditionGroupId field to given value.

### HasConditionGroupId

`func (o *RuleConditionResponseData) HasConditionGroupId() bool`

HasConditionGroupId returns a boolean if a field has been set.

### SetConditionGroupIdNil

`func (o *RuleConditionResponseData) SetConditionGroupIdNil(b bool)`

 SetConditionGroupIdNil sets the value for ConditionGroupId to be an explicit nil

### UnsetConditionGroupId
`func (o *RuleConditionResponseData) UnsetConditionGroupId()`

UnsetConditionGroupId ensures that no value is present for ConditionGroupId, not even an explicit nil
### GetConditionType

`func (o *RuleConditionResponseData) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *RuleConditionResponseData) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *RuleConditionResponseData) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetCreatedAt

`func (o *RuleConditionResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleConditionResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleConditionResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *RuleConditionResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RuleConditionResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RuleConditionResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventSubtype

`func (o *RuleConditionResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *RuleConditionResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *RuleConditionResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *RuleConditionResponseData) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *RuleConditionResponseData) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *RuleConditionResponseData) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetFlagId

`func (o *RuleConditionResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *RuleConditionResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *RuleConditionResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *RuleConditionResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *RuleConditionResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *RuleConditionResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *RuleConditionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleConditionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleConditionResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMetricPeriod

`func (o *RuleConditionResponseData) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *RuleConditionResponseData) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *RuleConditionResponseData) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *RuleConditionResponseData) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *RuleConditionResponseData) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *RuleConditionResponseData) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetMetricValue

`func (o *RuleConditionResponseData) GetMetricValue() int32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *RuleConditionResponseData) GetMetricValueOk() (*int32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *RuleConditionResponseData) SetMetricValue(v int32)`

SetMetricValue sets MetricValue field to given value.

### HasMetricValue

`func (o *RuleConditionResponseData) HasMetricValue() bool`

HasMetricValue returns a boolean if a field has been set.

### SetMetricValueNil

`func (o *RuleConditionResponseData) SetMetricValueNil(b bool)`

 SetMetricValueNil sets the value for MetricValue to be an explicit nil

### UnsetMetricValue
`func (o *RuleConditionResponseData) UnsetMetricValue()`

UnsetMetricValue ensures that no value is present for MetricValue, not even an explicit nil
### GetOperator

`func (o *RuleConditionResponseData) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RuleConditionResponseData) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RuleConditionResponseData) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPlanId

`func (o *RuleConditionResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RuleConditionResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RuleConditionResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RuleConditionResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RuleConditionResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RuleConditionResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetResourceIds

`func (o *RuleConditionResponseData) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *RuleConditionResponseData) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *RuleConditionResponseData) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetRuleId

`func (o *RuleConditionResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleConditionResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleConditionResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetTraitEntityType

`func (o *RuleConditionResponseData) GetTraitEntityType() string`

GetTraitEntityType returns the TraitEntityType field if non-nil, zero value otherwise.

### GetTraitEntityTypeOk

`func (o *RuleConditionResponseData) GetTraitEntityTypeOk() (*string, bool)`

GetTraitEntityTypeOk returns a tuple with the TraitEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitEntityType

`func (o *RuleConditionResponseData) SetTraitEntityType(v string)`

SetTraitEntityType sets TraitEntityType field to given value.

### HasTraitEntityType

`func (o *RuleConditionResponseData) HasTraitEntityType() bool`

HasTraitEntityType returns a boolean if a field has been set.

### SetTraitEntityTypeNil

`func (o *RuleConditionResponseData) SetTraitEntityTypeNil(b bool)`

 SetTraitEntityTypeNil sets the value for TraitEntityType to be an explicit nil

### UnsetTraitEntityType
`func (o *RuleConditionResponseData) UnsetTraitEntityType()`

UnsetTraitEntityType ensures that no value is present for TraitEntityType, not even an explicit nil
### GetTraitId

`func (o *RuleConditionResponseData) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *RuleConditionResponseData) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *RuleConditionResponseData) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *RuleConditionResponseData) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *RuleConditionResponseData) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *RuleConditionResponseData) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
### GetTraitValue

`func (o *RuleConditionResponseData) GetTraitValue() string`

GetTraitValue returns the TraitValue field if non-nil, zero value otherwise.

### GetTraitValueOk

`func (o *RuleConditionResponseData) GetTraitValueOk() (*string, bool)`

GetTraitValueOk returns a tuple with the TraitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitValue

`func (o *RuleConditionResponseData) SetTraitValue(v string)`

SetTraitValue sets TraitValue field to given value.


### GetUpdatedAt

`func (o *RuleConditionResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleConditionResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleConditionResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


