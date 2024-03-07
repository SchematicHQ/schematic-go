# CreateOrUpdateConditionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonTraitId** | Pointer to **NullableString** | Optionally provide a trait ID to compare a metric or trait value against instead of a value | [optional] 
**ConditionType** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** | Name of track event type used to measure this condition | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**MetricPeriod** | Pointer to **NullableString** | Period of time over which to measure the track event metric | [optional] 
**MetricValue** | **int32** | Value to compare the track event metric against | 
**Operator** | **string** |  | 
**ResourceIds** | **[]string** | List of resource IDs (companies, users, or plans) targeted by this condition | 
**TraitId** | Pointer to **NullableString** | ID of trait to use to measure this condition | [optional] 
**TraitValue** | Pointer to **NullableString** | Value to compare the trait value against | [optional] 

## Methods

### NewCreateOrUpdateConditionRequestBody

`func NewCreateOrUpdateConditionRequestBody(conditionType string, metricValue int32, operator string, resourceIds []string, ) *CreateOrUpdateConditionRequestBody`

NewCreateOrUpdateConditionRequestBody instantiates a new CreateOrUpdateConditionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConditionRequestBodyWithDefaults

`func NewCreateOrUpdateConditionRequestBodyWithDefaults() *CreateOrUpdateConditionRequestBody`

NewCreateOrUpdateConditionRequestBodyWithDefaults instantiates a new CreateOrUpdateConditionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonTraitId

`func (o *CreateOrUpdateConditionRequestBody) GetComparisonTraitId() string`

GetComparisonTraitId returns the ComparisonTraitId field if non-nil, zero value otherwise.

### GetComparisonTraitIdOk

`func (o *CreateOrUpdateConditionRequestBody) GetComparisonTraitIdOk() (*string, bool)`

GetComparisonTraitIdOk returns a tuple with the ComparisonTraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonTraitId

`func (o *CreateOrUpdateConditionRequestBody) SetComparisonTraitId(v string)`

SetComparisonTraitId sets ComparisonTraitId field to given value.

### HasComparisonTraitId

`func (o *CreateOrUpdateConditionRequestBody) HasComparisonTraitId() bool`

HasComparisonTraitId returns a boolean if a field has been set.

### SetComparisonTraitIdNil

`func (o *CreateOrUpdateConditionRequestBody) SetComparisonTraitIdNil(b bool)`

 SetComparisonTraitIdNil sets the value for ComparisonTraitId to be an explicit nil

### UnsetComparisonTraitId
`func (o *CreateOrUpdateConditionRequestBody) UnsetComparisonTraitId()`

UnsetComparisonTraitId ensures that no value is present for ComparisonTraitId, not even an explicit nil
### GetConditionType

`func (o *CreateOrUpdateConditionRequestBody) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *CreateOrUpdateConditionRequestBody) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *CreateOrUpdateConditionRequestBody) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetEventSubtype

`func (o *CreateOrUpdateConditionRequestBody) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *CreateOrUpdateConditionRequestBody) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *CreateOrUpdateConditionRequestBody) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *CreateOrUpdateConditionRequestBody) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *CreateOrUpdateConditionRequestBody) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *CreateOrUpdateConditionRequestBody) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetId

`func (o *CreateOrUpdateConditionRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrUpdateConditionRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrUpdateConditionRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrUpdateConditionRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateOrUpdateConditionRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateOrUpdateConditionRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMetricPeriod

`func (o *CreateOrUpdateConditionRequestBody) GetMetricPeriod() string`

GetMetricPeriod returns the MetricPeriod field if non-nil, zero value otherwise.

### GetMetricPeriodOk

`func (o *CreateOrUpdateConditionRequestBody) GetMetricPeriodOk() (*string, bool)`

GetMetricPeriodOk returns a tuple with the MetricPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPeriod

`func (o *CreateOrUpdateConditionRequestBody) SetMetricPeriod(v string)`

SetMetricPeriod sets MetricPeriod field to given value.

### HasMetricPeriod

`func (o *CreateOrUpdateConditionRequestBody) HasMetricPeriod() bool`

HasMetricPeriod returns a boolean if a field has been set.

### SetMetricPeriodNil

`func (o *CreateOrUpdateConditionRequestBody) SetMetricPeriodNil(b bool)`

 SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil

### UnsetMetricPeriod
`func (o *CreateOrUpdateConditionRequestBody) UnsetMetricPeriod()`

UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
### GetMetricValue

`func (o *CreateOrUpdateConditionRequestBody) GetMetricValue() int32`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *CreateOrUpdateConditionRequestBody) GetMetricValueOk() (*int32, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *CreateOrUpdateConditionRequestBody) SetMetricValue(v int32)`

SetMetricValue sets MetricValue field to given value.


### GetOperator

`func (o *CreateOrUpdateConditionRequestBody) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CreateOrUpdateConditionRequestBody) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CreateOrUpdateConditionRequestBody) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetResourceIds

`func (o *CreateOrUpdateConditionRequestBody) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *CreateOrUpdateConditionRequestBody) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *CreateOrUpdateConditionRequestBody) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetTraitId

`func (o *CreateOrUpdateConditionRequestBody) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *CreateOrUpdateConditionRequestBody) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *CreateOrUpdateConditionRequestBody) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *CreateOrUpdateConditionRequestBody) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *CreateOrUpdateConditionRequestBody) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *CreateOrUpdateConditionRequestBody) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
### GetTraitValue

`func (o *CreateOrUpdateConditionRequestBody) GetTraitValue() string`

GetTraitValue returns the TraitValue field if non-nil, zero value otherwise.

### GetTraitValueOk

`func (o *CreateOrUpdateConditionRequestBody) GetTraitValueOk() (*string, bool)`

GetTraitValueOk returns a tuple with the TraitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitValue

`func (o *CreateOrUpdateConditionRequestBody) SetTraitValue(v string)`

SetTraitValue sets TraitValue field to given value.

### HasTraitValue

`func (o *CreateOrUpdateConditionRequestBody) HasTraitValue() bool`

HasTraitValue returns a boolean if a field has been set.

### SetTraitValueNil

`func (o *CreateOrUpdateConditionRequestBody) SetTraitValueNil(b bool)`

 SetTraitValueNil sets the value for TraitValue to be an explicit nil

### UnsetTraitValue
`func (o *CreateOrUpdateConditionRequestBody) UnsetTraitValue()`

UnsetTraitValue ensures that no value is present for TraitValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


