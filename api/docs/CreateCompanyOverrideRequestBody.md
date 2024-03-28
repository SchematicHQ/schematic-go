# CreateCompanyOverrideRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** |  | 
**FeatureId** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**ValueBool** | Pointer to **NullableBool** |  | [optional] 
**ValueNumeric** | Pointer to **NullableInt32** |  | [optional] 
**ValueTraitId** | Pointer to **NullableString** |  | [optional] 
**ValueType** | **string** |  | 

## Methods

### NewCreateCompanyOverrideRequestBody

`func NewCreateCompanyOverrideRequestBody(companyId string, featureId string, valueType string, ) *CreateCompanyOverrideRequestBody`

NewCreateCompanyOverrideRequestBody instantiates a new CreateCompanyOverrideRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompanyOverrideRequestBodyWithDefaults

`func NewCreateCompanyOverrideRequestBodyWithDefaults() *CreateCompanyOverrideRequestBody`

NewCreateCompanyOverrideRequestBodyWithDefaults instantiates a new CreateCompanyOverrideRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CreateCompanyOverrideRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateCompanyOverrideRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateCompanyOverrideRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetFeatureId

`func (o *CreateCompanyOverrideRequestBody) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CreateCompanyOverrideRequestBody) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CreateCompanyOverrideRequestBody) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetMetricPeriod

`func (o *CreateCompanyOverrideRequestBody) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *CreateCompanyOverrideRequestBody) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *CreateCompanyOverrideRequestBody) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *CreateCompanyOverrideRequestBody) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *CreateCompanyOverrideRequestBody) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *CreateCompanyOverrideRequestBody) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetValueBool

`func (o *CreateCompanyOverrideRequestBody) GetValueBool() bool`

GetValueBool returns the ValueBool field if non-nil, zero value otherwise.

### GetValueBoolOk

`func (o *CreateCompanyOverrideRequestBody) GetValueBoolOk() (*bool, bool)`

GetValueBoolOk returns a tuple with the ValueBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBool

`func (o *CreateCompanyOverrideRequestBody) SetValueBool(v bool)`

SetValueBool sets ValueBool field to given value.

### HasValueBool

`func (o *CreateCompanyOverrideRequestBody) HasValueBool() bool`

HasValueBool returns a boolean if a field has been set.

### SetValueBoolNil

`func (o *CreateCompanyOverrideRequestBody) SetValueBoolNil(b bool)`

 SetValueBoolNil sets the value for ValueBool to be an explicit nil

### UnsetValueBool
`func (o *CreateCompanyOverrideRequestBody) UnsetValueBool()`

UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
### GetValueNumeric

`func (o *CreateCompanyOverrideRequestBody) GetValueNumeric() int32`

GetValueNumeric returns the ValueNumeric field if non-nil, zero value otherwise.

### GetValueNumericOk

`func (o *CreateCompanyOverrideRequestBody) GetValueNumericOk() (*int32, bool)`

GetValueNumericOk returns a tuple with the ValueNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNumeric

`func (o *CreateCompanyOverrideRequestBody) SetValueNumeric(v int32)`

SetValueNumeric sets ValueNumeric field to given value.

### HasValueNumeric

`func (o *CreateCompanyOverrideRequestBody) HasValueNumeric() bool`

HasValueNumeric returns a boolean if a field has been set.

### SetValueNumericNil

`func (o *CreateCompanyOverrideRequestBody) SetValueNumericNil(b bool)`

 SetValueNumericNil sets the value for ValueNumeric to be an explicit nil

### UnsetValueNumeric
`func (o *CreateCompanyOverrideRequestBody) UnsetValueNumeric()`

UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
### GetValueTraitId

`func (o *CreateCompanyOverrideRequestBody) GetValueTraitId() string`

GetValueTraitId returns the ValueTraitId field if non-nil, zero value otherwise.

### GetValueTraitIdOk

`func (o *CreateCompanyOverrideRequestBody) GetValueTraitIdOk() (*string, bool)`

GetValueTraitIdOk returns a tuple with the ValueTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTraitId

`func (o *CreateCompanyOverrideRequestBody) SetValueTraitId(v string)`

SetValueTraitId sets ValueTraitId field to given value.

### HasValueTraitId

`func (o *CreateCompanyOverrideRequestBody) HasValueTraitId() bool`

HasValueTraitId returns a boolean if a field has been set.

### SetValueTraitIdNil

`func (o *CreateCompanyOverrideRequestBody) SetValueTraitIdNil(b bool)`

 SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil

### UnsetValueTraitId
`func (o *CreateCompanyOverrideRequestBody) UnsetValueTraitId()`

UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
### GetValueType

`func (o *CreateCompanyOverrideRequestBody) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateCompanyOverrideRequestBody) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateCompanyOverrideRequestBody) SetValueType(v string)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


