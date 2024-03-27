# CreatePlanEntitlementRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**PlanId** | **string** |  | 
**ValueBool** | Pointer to **NullableBool** |  | [optional] 
**ValueNumeric** | Pointer to **NullableInt32** |  | [optional] 
**ValueTraitId** | Pointer to **NullableString** |  | [optional] 
**ValueType** | **string** |  | 

## Methods

### NewCreatePlanEntitlementRequestBody

`func NewCreatePlanEntitlementRequestBody(featureId string, planId string, valueType string, ) *CreatePlanEntitlementRequestBody`

NewCreatePlanEntitlementRequestBody instantiates a new CreatePlanEntitlementRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanEntitlementRequestBodyWithDefaults

`func NewCreatePlanEntitlementRequestBodyWithDefaults() *CreatePlanEntitlementRequestBody`

NewCreatePlanEntitlementRequestBodyWithDefaults instantiates a new CreatePlanEntitlementRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *CreatePlanEntitlementRequestBody) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CreatePlanEntitlementRequestBody) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CreatePlanEntitlementRequestBody) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetMetricPeriod

`func (o *CreatePlanEntitlementRequestBody) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *CreatePlanEntitlementRequestBody) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *CreatePlanEntitlementRequestBody) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *CreatePlanEntitlementRequestBody) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *CreatePlanEntitlementRequestBody) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *CreatePlanEntitlementRequestBody) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetPlanId

`func (o *CreatePlanEntitlementRequestBody) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreatePlanEntitlementRequestBody) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreatePlanEntitlementRequestBody) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetValueBool

`func (o *CreatePlanEntitlementRequestBody) GetValueBool() bool`

GetValueBool returns the ValueBool field if non-nil, zero value otherwise.

### GetValueBoolOk

`func (o *CreatePlanEntitlementRequestBody) GetValueBoolOk() (*bool, bool)`

GetValueBoolOk returns a tuple with the ValueBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBool

`func (o *CreatePlanEntitlementRequestBody) SetValueBool(v bool)`

SetValueBool sets ValueBool field to given value.

### HasValueBool

`func (o *CreatePlanEntitlementRequestBody) HasValueBool() bool`

HasValueBool returns a boolean if a field has been set.

### SetValueBoolNil

`func (o *CreatePlanEntitlementRequestBody) SetValueBoolNil(b bool)`

 SetValueBoolNil sets the value for ValueBool to be an explicit nil

### UnsetValueBool
`func (o *CreatePlanEntitlementRequestBody) UnsetValueBool()`

UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
### GetValueNumeric

`func (o *CreatePlanEntitlementRequestBody) GetValueNumeric() int32`

GetValueNumeric returns the ValueNumeric field if non-nil, zero value otherwise.

### GetValueNumericOk

`func (o *CreatePlanEntitlementRequestBody) GetValueNumericOk() (*int32, bool)`

GetValueNumericOk returns a tuple with the ValueNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNumeric

`func (o *CreatePlanEntitlementRequestBody) SetValueNumeric(v int32)`

SetValueNumeric sets ValueNumeric field to given value.

### HasValueNumeric

`func (o *CreatePlanEntitlementRequestBody) HasValueNumeric() bool`

HasValueNumeric returns a boolean if a field has been set.

### SetValueNumericNil

`func (o *CreatePlanEntitlementRequestBody) SetValueNumericNil(b bool)`

 SetValueNumericNil sets the value for ValueNumeric to be an explicit nil

### UnsetValueNumeric
`func (o *CreatePlanEntitlementRequestBody) UnsetValueNumeric()`

UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
### GetValueTraitId

`func (o *CreatePlanEntitlementRequestBody) GetValueTraitId() string`

GetValueTraitId returns the ValueTraitId field if non-nil, zero value otherwise.

### GetValueTraitIdOk

`func (o *CreatePlanEntitlementRequestBody) GetValueTraitIdOk() (*string, bool)`

GetValueTraitIdOk returns a tuple with the ValueTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTraitId

`func (o *CreatePlanEntitlementRequestBody) SetValueTraitId(v string)`

SetValueTraitId sets ValueTraitId field to given value.

### HasValueTraitId

`func (o *CreatePlanEntitlementRequestBody) HasValueTraitId() bool`

HasValueTraitId returns a boolean if a field has been set.

### SetValueTraitIdNil

`func (o *CreatePlanEntitlementRequestBody) SetValueTraitIdNil(b bool)`

 SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil

### UnsetValueTraitId
`func (o *CreatePlanEntitlementRequestBody) UnsetValueTraitId()`

UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
### GetValueType

`func (o *CreatePlanEntitlementRequestBody) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreatePlanEntitlementRequestBody) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreatePlanEntitlementRequestBody) SetValueType(v string)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


