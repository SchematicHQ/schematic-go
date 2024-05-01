# FeatureUsageResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **bool** | Whether further usage is permitted. | 
**Allocation** | Pointer to **NullableInt32** | The maximum amount of usage that is permitted; a null value indicates that unlimited usage is permitted. | [optional] 
**EntitlementId** | **string** |  | 
**EntitlementType** | **string** |  | 
**Feature** | Pointer to [**FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | [optional] 
**Period** | Pointer to **NullableString** | The period over which usage is measured. | [optional] 
**Plan** | Pointer to [**PlanResponseData**](PlanResponseData.md) |  | [optional] 
**Usage** | Pointer to **NullableInt32** | The amount of usage that has been consumed; a null value indicates that usage is not being measured. | [optional] 

## Methods

### NewFeatureUsageResponseData

`func NewFeatureUsageResponseData(access bool, entitlementId string, entitlementType string, ) *FeatureUsageResponseData`

NewFeatureUsageResponseData instantiates a new FeatureUsageResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUsageResponseDataWithDefaults

`func NewFeatureUsageResponseDataWithDefaults() *FeatureUsageResponseData`

NewFeatureUsageResponseDataWithDefaults instantiates a new FeatureUsageResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FeatureUsageResponseData) GetAccess() bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FeatureUsageResponseData) GetAccessOk() (*bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FeatureUsageResponseData) SetAccess(v bool)`

SetAccess sets Access field to given value.


### GetAllocation

`func (o *FeatureUsageResponseData) GetAllocation() int32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *FeatureUsageResponseData) GetAllocationOk() (*int32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *FeatureUsageResponseData) SetAllocation(v int32)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *FeatureUsageResponseData) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### SetAllocationNil

`func (o *FeatureUsageResponseData) SetAllocationNil(b bool)`

 SetAllocationNil sets the value for Allocation to be an explicit nil

### UnsetAllocation
`func (o *FeatureUsageResponseData) UnsetAllocation()`

UnsetAllocation ensures that no value is present for Allocation, not even an explicit nil
### GetEntitlementId

`func (o *FeatureUsageResponseData) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *FeatureUsageResponseData) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *FeatureUsageResponseData) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetEntitlementType

`func (o *FeatureUsageResponseData) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *FeatureUsageResponseData) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *FeatureUsageResponseData) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetFeature

`func (o *FeatureUsageResponseData) GetFeature() FeatureDetailResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureUsageResponseData) GetFeatureOk() (*FeatureDetailResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureUsageResponseData) SetFeature(v FeatureDetailResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureUsageResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetPeriod

`func (o *FeatureUsageResponseData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *FeatureUsageResponseData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *FeatureUsageResponseData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *FeatureUsageResponseData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *FeatureUsageResponseData) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *FeatureUsageResponseData) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPlan

`func (o *FeatureUsageResponseData) GetPlan() PlanResponseData`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FeatureUsageResponseData) GetPlanOk() (*PlanResponseData, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FeatureUsageResponseData) SetPlan(v PlanResponseData)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *FeatureUsageResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUsage

`func (o *FeatureUsageResponseData) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FeatureUsageResponseData) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FeatureUsageResponseData) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FeatureUsageResponseData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *FeatureUsageResponseData) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *FeatureUsageResponseData) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


