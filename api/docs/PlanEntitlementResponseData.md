# PlanEntitlementResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**Feature** | Pointer to [**FeatureResponseData**](FeatureResponseData.md) |  | [optional] 
**FeatureId** | **string** |  | 
**Id** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**Plan** | Pointer to [**PlanResponseData**](PlanResponseData.md) |  | [optional] 
**PlanId** | **string** |  | 
**RuleId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**ValueBool** | Pointer to **NullableBool** |  | [optional] 
**ValueNumeric** | Pointer to **NullableInt32** |  | [optional] 
**ValueTrait** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**ValueTraitId** | Pointer to **NullableString** |  | [optional] 
**ValueType** | **string** |  | 

## Methods

### NewPlanEntitlementResponseData

`func NewPlanEntitlementResponseData(createdAt time.Time, environmentId string, featureId string, id string, planId string, ruleId string, updatedAt time.Time, valueType string, ) *PlanEntitlementResponseData`

NewPlanEntitlementResponseData instantiates a new PlanEntitlementResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanEntitlementResponseDataWithDefaults

`func NewPlanEntitlementResponseDataWithDefaults() *PlanEntitlementResponseData`

NewPlanEntitlementResponseDataWithDefaults instantiates a new PlanEntitlementResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PlanEntitlementResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanEntitlementResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanEntitlementResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *PlanEntitlementResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PlanEntitlementResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PlanEntitlementResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFeature

`func (o *PlanEntitlementResponseData) GetFeature() FeatureResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *PlanEntitlementResponseData) GetFeatureOk() (*FeatureResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *PlanEntitlementResponseData) SetFeature(v FeatureResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *PlanEntitlementResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureId

`func (o *PlanEntitlementResponseData) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *PlanEntitlementResponseData) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *PlanEntitlementResponseData) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetId

`func (o *PlanEntitlementResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanEntitlementResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanEntitlementResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMetricPeriod

`func (o *PlanEntitlementResponseData) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *PlanEntitlementResponseData) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *PlanEntitlementResponseData) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *PlanEntitlementResponseData) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *PlanEntitlementResponseData) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *PlanEntitlementResponseData) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetPlan

`func (o *PlanEntitlementResponseData) GetPlan() PlanResponseData`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PlanEntitlementResponseData) GetPlanOk() (*PlanResponseData, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PlanEntitlementResponseData) SetPlan(v PlanResponseData)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PlanEntitlementResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *PlanEntitlementResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *PlanEntitlementResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *PlanEntitlementResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetRuleId

`func (o *PlanEntitlementResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *PlanEntitlementResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *PlanEntitlementResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetUpdatedAt

`func (o *PlanEntitlementResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanEntitlementResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanEntitlementResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValueBool

`func (o *PlanEntitlementResponseData) GetValueBool() bool`

GetValueBool returns the ValueBool field if non-nil, zero value otherwise.

### GetValueBoolOk

`func (o *PlanEntitlementResponseData) GetValueBoolOk() (*bool, bool)`

GetValueBoolOk returns a tuple with the ValueBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBool

`func (o *PlanEntitlementResponseData) SetValueBool(v bool)`

SetValueBool sets ValueBool field to given value.

### HasValueBool

`func (o *PlanEntitlementResponseData) HasValueBool() bool`

HasValueBool returns a boolean if a field has been set.

### SetValueBoolNil

`func (o *PlanEntitlementResponseData) SetValueBoolNil(b bool)`

 SetValueBoolNil sets the value for ValueBool to be an explicit nil

### UnsetValueBool
`func (o *PlanEntitlementResponseData) UnsetValueBool()`

UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
### GetValueNumeric

`func (o *PlanEntitlementResponseData) GetValueNumeric() int32`

GetValueNumeric returns the ValueNumeric field if non-nil, zero value otherwise.

### GetValueNumericOk

`func (o *PlanEntitlementResponseData) GetValueNumericOk() (*int32, bool)`

GetValueNumericOk returns a tuple with the ValueNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNumeric

`func (o *PlanEntitlementResponseData) SetValueNumeric(v int32)`

SetValueNumeric sets ValueNumeric field to given value.

### HasValueNumeric

`func (o *PlanEntitlementResponseData) HasValueNumeric() bool`

HasValueNumeric returns a boolean if a field has been set.

### SetValueNumericNil

`func (o *PlanEntitlementResponseData) SetValueNumericNil(b bool)`

 SetValueNumericNil sets the value for ValueNumeric to be an explicit nil

### UnsetValueNumeric
`func (o *PlanEntitlementResponseData) UnsetValueNumeric()`

UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
### GetValueTrait

`func (o *PlanEntitlementResponseData) GetValueTrait() EntityTraitDefinitionResponseData`

GetValueTrait returns the ValueTrait field if non-nil, zero value otherwise.

### GetValueTraitOk

`func (o *PlanEntitlementResponseData) GetValueTraitOk() (*EntityTraitDefinitionResponseData, bool)`

GetValueTraitOk returns a tuple with the ValueTrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTrait

`func (o *PlanEntitlementResponseData) SetValueTrait(v EntityTraitDefinitionResponseData)`

SetValueTrait sets ValueTrait field to given value.

### HasValueTrait

`func (o *PlanEntitlementResponseData) HasValueTrait() bool`

HasValueTrait returns a boolean if a field has been set.

### GetValueTraitId

`func (o *PlanEntitlementResponseData) GetValueTraitId() string`

GetValueTraitId returns the ValueTraitId field if non-nil, zero value otherwise.

### GetValueTraitIdOk

`func (o *PlanEntitlementResponseData) GetValueTraitIdOk() (*string, bool)`

GetValueTraitIdOk returns a tuple with the ValueTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTraitId

`func (o *PlanEntitlementResponseData) SetValueTraitId(v string)`

SetValueTraitId sets ValueTraitId field to given value.

### HasValueTraitId

`func (o *PlanEntitlementResponseData) HasValueTraitId() bool`

HasValueTraitId returns a boolean if a field has been set.

### SetValueTraitIdNil

`func (o *PlanEntitlementResponseData) SetValueTraitIdNil(b bool)`

 SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil

### UnsetValueTraitId
`func (o *PlanEntitlementResponseData) UnsetValueTraitId()`

UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
### GetValueType

`func (o *PlanEntitlementResponseData) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *PlanEntitlementResponseData) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *PlanEntitlementResponseData) SetValueType(v string)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


