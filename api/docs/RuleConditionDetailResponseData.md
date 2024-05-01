# RuleConditionDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonTrait** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**ComparisonTraitId** | Pointer to **NullableString** |  | [optional] 
**ConditionGroupId** | Pointer to **NullableString** |  | [optional] 
**ConditionType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**MetricValue** | **int32** |  | 
**Operator** | **string** |  | 
**PlanId** | Pointer to **NullableString** |  | [optional] 
**ResourceIds** | **[]string** |  | 
**Resources** | [**[]RuleConditionResourceResponseData**](RuleConditionResourceResponseData.md) |  | 
**RuleId** | **string** |  | 
**Trait** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**TraitEntityType** | Pointer to **NullableString** |  | [optional] 
**TraitId** | Pointer to **NullableString** |  | [optional] 
**TraitValue** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewRuleConditionDetailResponseData

`func NewRuleConditionDetailResponseData(conditionType string, createdAt time.Time, environmentId string, id string, metricValue int32, operator string, resourceIds []string, resources []RuleConditionResourceResponseData, ruleId string, traitValue string, updatedAt time.Time, ) *RuleConditionDetailResponseData`

NewRuleConditionDetailResponseData instantiates a new RuleConditionDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConditionDetailResponseDataWithDefaults

`func NewRuleConditionDetailResponseDataWithDefaults() *RuleConditionDetailResponseData`

NewRuleConditionDetailResponseDataWithDefaults instantiates a new RuleConditionDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonTrait

`func (o *RuleConditionDetailResponseData) GetComparisonTrait() EntityTraitDefinitionResponseData`

GetComparisonTrait returns the ComparisonTrait field if non-nil, zero value otherwise.

### GetComparisonTraitOk

`func (o *RuleConditionDetailResponseData) GetComparisonTraitOk() (*EntityTraitDefinitionResponseData, bool)`

GetComparisonTraitOk returns a tuple with the ComparisonTrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonTrait

`func (o *RuleConditionDetailResponseData) SetComparisonTrait(v EntityTraitDefinitionResponseData)`

SetComparisonTrait sets ComparisonTrait field to given value.

### HasComparisonTrait

`func (o *RuleConditionDetailResponseData) HasComparisonTrait() bool`

HasComparisonTrait returns a boolean if a field has been set.

### GetComparisonTraitId

`func (o *RuleConditionDetailResponseData) GetComparisonTraitId() string`

GetComparisonTraitId returns the ComparisonTraitId field if non-nil, zero value otherwise.

### GetComparisonTraitIdOk

`func (o *RuleConditionDetailResponseData) GetComparisonTraitIdOk() (*string, bool)`

GetComparisonTraitIdOk returns a tuple with the ComparisonTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonTraitId

`func (o *RuleConditionDetailResponseData) SetComparisonTraitId(v string)`

SetComparisonTraitId sets ComparisonTraitId field to given value.

### HasComparisonTraitId

`func (o *RuleConditionDetailResponseData) HasComparisonTraitId() bool`

HasComparisonTraitId returns a boolean if a field has been set.

### SetComparisonTraitIdNil

`func (o *RuleConditionDetailResponseData) SetComparisonTraitIdNil(b bool)`

 SetComparisonTraitIdNil sets the value for ComparisonTraitId to be an explicit nil

### UnsetComparisonTraitId
`func (o *RuleConditionDetailResponseData) UnsetComparisonTraitId()`

UnsetComparisonTraitId ensures that no value is present for ComparisonTraitId, not even an explicit nil
### GetConditionGroupId

`func (o *RuleConditionDetailResponseData) GetConditionGroupId() string`

GetConditionGroupId returns the ConditionGroupId field if non-nil, zero value otherwise.

### GetConditionGroupIdOk

`func (o *RuleConditionDetailResponseData) GetConditionGroupIdOk() (*string, bool)`

GetConditionGroupIdOk returns a tuple with the ConditionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroupId

`func (o *RuleConditionDetailResponseData) SetConditionGroupId(v string)`

SetConditionGroupId sets ConditionGroupId field to given value.

### HasConditionGroupId

`func (o *RuleConditionDetailResponseData) HasConditionGroupId() bool`

HasConditionGroupId returns a boolean if a field has been set.

### SetConditionGroupIdNil

`func (o *RuleConditionDetailResponseData) SetConditionGroupIdNil(b bool)`

 SetConditionGroupIdNil sets the value for ConditionGroupId to be an explicit nil

### UnsetConditionGroupId
`func (o *RuleConditionDetailResponseData) UnsetConditionGroupId()`

UnsetConditionGroupId ensures that no value is present for ConditionGroupId, not even an explicit nil
### GetConditionType

`func (o *RuleConditionDetailResponseData) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *RuleConditionDetailResponseData) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *RuleConditionDetailResponseData) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetCreatedAt

`func (o *RuleConditionDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RuleConditionDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RuleConditionDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *RuleConditionDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *RuleConditionDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *RuleConditionDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventSubtype

`func (o *RuleConditionDetailResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *RuleConditionDetailResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *RuleConditionDetailResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *RuleConditionDetailResponseData) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *RuleConditionDetailResponseData) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *RuleConditionDetailResponseData) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetFlagId

`func (o *RuleConditionDetailResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *RuleConditionDetailResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *RuleConditionDetailResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *RuleConditionDetailResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *RuleConditionDetailResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *RuleConditionDetailResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *RuleConditionDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleConditionDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleConditionDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMetricPeriod

`func (o *RuleConditionDetailResponseData) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *RuleConditionDetailResponseData) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *RuleConditionDetailResponseData) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *RuleConditionDetailResponseData) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *RuleConditionDetailResponseData) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *RuleConditionDetailResponseData) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetMetricValue

`func (o *RuleConditionDetailResponseData) GetMetricValue() int32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *RuleConditionDetailResponseData) GetMetricValueOk() (*int32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *RuleConditionDetailResponseData) SetMetricValue(v int32)`

SetMetricValue sets MetricValue field to given value.


### GetOperator

`func (o *RuleConditionDetailResponseData) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RuleConditionDetailResponseData) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RuleConditionDetailResponseData) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPlanId

`func (o *RuleConditionDetailResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *RuleConditionDetailResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *RuleConditionDetailResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *RuleConditionDetailResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *RuleConditionDetailResponseData) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *RuleConditionDetailResponseData) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
### GetResourceIds

`func (o *RuleConditionDetailResponseData) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *RuleConditionDetailResponseData) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *RuleConditionDetailResponseData) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetResources

`func (o *RuleConditionDetailResponseData) GetResources() []RuleConditionResourceResponseData`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RuleConditionDetailResponseData) GetResourcesOk() (*[]RuleConditionResourceResponseData, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RuleConditionDetailResponseData) SetResources(v []RuleConditionResourceResponseData)`

SetResources sets Resources field to given value.


### GetRuleId

`func (o *RuleConditionDetailResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleConditionDetailResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleConditionDetailResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetTrait

`func (o *RuleConditionDetailResponseData) GetTrait() EntityTraitDefinitionResponseData`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *RuleConditionDetailResponseData) GetTraitOk() (*EntityTraitDefinitionResponseData, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *RuleConditionDetailResponseData) SetTrait(v EntityTraitDefinitionResponseData)`

SetTrait sets Trait field to given value.

### HasTrait

`func (o *RuleConditionDetailResponseData) HasTrait() bool`

HasTrait returns a boolean if a field has been set.

### GetTraitEntityType

`func (o *RuleConditionDetailResponseData) GetTraitEntityType() string`

GetTraitEntityType returns the TraitEntityType field if non-nil, zero value otherwise.

### GetTraitEntityTypeOk

`func (o *RuleConditionDetailResponseData) GetTraitEntityTypeOk() (*string, bool)`

GetTraitEntityTypeOk returns a tuple with the TraitEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitEntityType

`func (o *RuleConditionDetailResponseData) SetTraitEntityType(v string)`

SetTraitEntityType sets TraitEntityType field to given value.

### HasTraitEntityType

`func (o *RuleConditionDetailResponseData) HasTraitEntityType() bool`

HasTraitEntityType returns a boolean if a field has been set.

### SetTraitEntityTypeNil

`func (o *RuleConditionDetailResponseData) SetTraitEntityTypeNil(b bool)`

 SetTraitEntityTypeNil sets the value for TraitEntityType to be an explicit nil

### UnsetTraitEntityType
`func (o *RuleConditionDetailResponseData) UnsetTraitEntityType()`

UnsetTraitEntityType ensures that no value is present for TraitEntityType, not even an explicit nil
### GetTraitId

`func (o *RuleConditionDetailResponseData) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *RuleConditionDetailResponseData) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *RuleConditionDetailResponseData) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *RuleConditionDetailResponseData) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *RuleConditionDetailResponseData) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *RuleConditionDetailResponseData) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
### GetTraitValue

`func (o *RuleConditionDetailResponseData) GetTraitValue() string`

GetTraitValue returns the TraitValue field if non-nil, zero value otherwise.

### GetTraitValueOk

`func (o *RuleConditionDetailResponseData) GetTraitValueOk() (*string, bool)`

GetTraitValueOk returns a tuple with the TraitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitValue

`func (o *RuleConditionDetailResponseData) SetTraitValue(v string)`

SetTraitValue sets TraitValue field to given value.


### GetUpdatedAt

`func (o *RuleConditionDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RuleConditionDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RuleConditionDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


