# CompanyOverrideResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompanyResponseData**](CompanyResponseData.md) |  | [optional] 
**CompanyId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**Feature** | Pointer to [**FeatureResponseData**](FeatureResponseData.md) |  | [optional] 
**FeatureId** | **string** |  | 
**Id** | **string** |  | 
**MetricPeriod** | Pointer to **NullableString** |  | [optional] 
**RuleId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**ValueBool** | Pointer to **NullableBool** |  | [optional] 
**ValueNumeric** | Pointer to **NullableInt32** |  | [optional] 
**ValueTrait** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**ValueTraitId** | Pointer to **NullableString** |  | [optional] 
**ValueType** | **string** |  | 

## Methods

### NewCompanyOverrideResponseData

`func NewCompanyOverrideResponseData(companyId string, createdAt time.Time, environmentId string, featureId string, id string, ruleId string, updatedAt time.Time, valueType string, ) *CompanyOverrideResponseData`

NewCompanyOverrideResponseData instantiates a new CompanyOverrideResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyOverrideResponseDataWithDefaults

`func NewCompanyOverrideResponseDataWithDefaults() *CompanyOverrideResponseData`

NewCompanyOverrideResponseDataWithDefaults instantiates a new CompanyOverrideResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *CompanyOverrideResponseData) GetCompany() CompanyResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyOverrideResponseData) GetCompanyOk() (*CompanyResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyOverrideResponseData) SetCompany(v CompanyResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyOverrideResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *CompanyOverrideResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyOverrideResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyOverrideResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedAt

`func (o *CompanyOverrideResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyOverrideResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyOverrideResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *CompanyOverrideResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CompanyOverrideResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CompanyOverrideResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetFeature

`func (o *CompanyOverrideResponseData) GetFeature() FeatureResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *CompanyOverrideResponseData) GetFeatureOk() (*FeatureResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *CompanyOverrideResponseData) SetFeature(v FeatureResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *CompanyOverrideResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureId

`func (o *CompanyOverrideResponseData) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CompanyOverrideResponseData) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CompanyOverrideResponseData) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetId

`func (o *CompanyOverrideResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyOverrideResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyOverrideResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMetricPeriod

`func (o *CompanyOverrideResponseData) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *CompanyOverrideResponseData) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *CompanyOverrideResponseData) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *CompanyOverrideResponseData) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *CompanyOverrideResponseData) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *CompanyOverrideResponseData) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetRuleId

`func (o *CompanyOverrideResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *CompanyOverrideResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *CompanyOverrideResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetUpdatedAt

`func (o *CompanyOverrideResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyOverrideResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyOverrideResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValueBool

`func (o *CompanyOverrideResponseData) GetValueBool() bool`

GetValueBool returns the ValueBool field if non-nil, zero value otherwise.

### GetValueBoolOk

`func (o *CompanyOverrideResponseData) GetValueBoolOk() (*bool, bool)`

GetValueBoolOk returns a tuple with the ValueBool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueBool

`func (o *CompanyOverrideResponseData) SetValueBool(v bool)`

SetValueBool sets ValueBool field to given value.

### HasValueBool

`func (o *CompanyOverrideResponseData) HasValueBool() bool`

HasValueBool returns a boolean if a field has been set.

### SetValueBoolNil

`func (o *CompanyOverrideResponseData) SetValueBoolNil(b bool)`

 SetValueBoolNil sets the value for ValueBool to be an explicit nil

### UnsetValueBool
`func (o *CompanyOverrideResponseData) UnsetValueBool()`

UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
### GetValueNumeric

`func (o *CompanyOverrideResponseData) GetValueNumeric() int32`

GetValueNumeric returns the ValueNumeric field if non-nil, zero value otherwise.

### GetValueNumericOk

`func (o *CompanyOverrideResponseData) GetValueNumericOk() (*int32, bool)`

GetValueNumericOk returns a tuple with the ValueNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueNumeric

`func (o *CompanyOverrideResponseData) SetValueNumeric(v int32)`

SetValueNumeric sets ValueNumeric field to given value.

### HasValueNumeric

`func (o *CompanyOverrideResponseData) HasValueNumeric() bool`

HasValueNumeric returns a boolean if a field has been set.

### SetValueNumericNil

`func (o *CompanyOverrideResponseData) SetValueNumericNil(b bool)`

 SetValueNumericNil sets the value for ValueNumeric to be an explicit nil

### UnsetValueNumeric
`func (o *CompanyOverrideResponseData) UnsetValueNumeric()`

UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
### GetValueTrait

`func (o *CompanyOverrideResponseData) GetValueTrait() EntityTraitDefinitionResponseData`

GetValueTrait returns the ValueTrait field if non-nil, zero value otherwise.

### GetValueTraitOk

`func (o *CompanyOverrideResponseData) GetValueTraitOk() (*EntityTraitDefinitionResponseData, bool)`

GetValueTraitOk returns a tuple with the ValueTrait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTrait

`func (o *CompanyOverrideResponseData) SetValueTrait(v EntityTraitDefinitionResponseData)`

SetValueTrait sets ValueTrait field to given value.

### HasValueTrait

`func (o *CompanyOverrideResponseData) HasValueTrait() bool`

HasValueTrait returns a boolean if a field has been set.

### GetValueTraitId

`func (o *CompanyOverrideResponseData) GetValueTraitId() string`

GetValueTraitId returns the ValueTraitId field if non-nil, zero value otherwise.

### GetValueTraitIdOk

`func (o *CompanyOverrideResponseData) GetValueTraitIdOk() (*string, bool)`

GetValueTraitIdOk returns a tuple with the ValueTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueTraitId

`func (o *CompanyOverrideResponseData) SetValueTraitId(v string)`

SetValueTraitId sets ValueTraitId field to given value.

### HasValueTraitId

`func (o *CompanyOverrideResponseData) HasValueTraitId() bool`

HasValueTraitId returns a boolean if a field has been set.

### SetValueTraitIdNil

`func (o *CompanyOverrideResponseData) SetValueTraitIdNil(b bool)`

 SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil

### UnsetValueTraitId
`func (o *CompanyOverrideResponseData) UnsetValueTraitId()`

UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
### GetValueType

`func (o *CompanyOverrideResponseData) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CompanyOverrideResponseData) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CompanyOverrideResponseData) SetValueType(v string)`

SetValueType sets ValueType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


