# FeatureCompanyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **bool** | Whether further usage is permitted. | 
**Allocation** | Pointer to **NullableInt32** | The maximum amount of usage that is permitted; a null value indicates that unlimited usage is permitted. | [optional] 
**AllocationType** | **string** | The type of allocation that is being used. | 
**Company** | Pointer to [**CompanyDetailResponseData**](CompanyDetailResponseData.md) |  | [optional] 
**EntitlementId** | **string** |  | 
**EntitlementType** | **string** |  | 
**Feature** | Pointer to [**FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | [optional] 
**Period** | Pointer to **NullableString** | The period over which usage is measured. | [optional] 
**Plan** | Pointer to [**PlanResponseData**](PlanResponseData.md) |  | [optional] 
**Usage** | Pointer to **NullableInt32** | The amount of usage that has been consumed; a null value indicates that usage is not being measured. | [optional] 

## Methods

### NewFeatureCompanyResponseData

`func NewFeatureCompanyResponseData(access bool, allocationType string, entitlementId string, entitlementType string, ) *FeatureCompanyResponseData`

NewFeatureCompanyResponseData instantiates a new FeatureCompanyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureCompanyResponseDataWithDefaults

`func NewFeatureCompanyResponseDataWithDefaults() *FeatureCompanyResponseData`

NewFeatureCompanyResponseDataWithDefaults instantiates a new FeatureCompanyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FeatureCompanyResponseData) GetAccess() bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FeatureCompanyResponseData) GetAccessOk() (*bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FeatureCompanyResponseData) SetAccess(v bool)`

SetAccess sets Access field to given value.


### GetAllocation

`func (o *FeatureCompanyResponseData) GetAllocation() int32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *FeatureCompanyResponseData) GetAllocationOk() (*int32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *FeatureCompanyResponseData) SetAllocation(v int32)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *FeatureCompanyResponseData) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### SetAllocationNil

`func (o *FeatureCompanyResponseData) SetAllocationNil(b bool)`

 SetAllocationNil sets the value for Allocation to be an explicit nil

### UnsetAllocation
`func (o *FeatureCompanyResponseData) UnsetAllocation()`

UnsetAllocation ensures that no value is present for Allocation, not even an explicit nil
### GetAllocationType

`func (o *FeatureCompanyResponseData) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *FeatureCompanyResponseData) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *FeatureCompanyResponseData) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.


### GetCompany

`func (o *FeatureCompanyResponseData) GetCompany() CompanyDetailResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FeatureCompanyResponseData) GetCompanyOk() (*CompanyDetailResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FeatureCompanyResponseData) SetCompany(v CompanyDetailResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FeatureCompanyResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEntitlementId

`func (o *FeatureCompanyResponseData) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *FeatureCompanyResponseData) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *FeatureCompanyResponseData) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetEntitlementType

`func (o *FeatureCompanyResponseData) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *FeatureCompanyResponseData) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *FeatureCompanyResponseData) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetFeature

`func (o *FeatureCompanyResponseData) GetFeature() FeatureDetailResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureCompanyResponseData) GetFeatureOk() (*FeatureDetailResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureCompanyResponseData) SetFeature(v FeatureDetailResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureCompanyResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetPeriod

`func (o *FeatureCompanyResponseData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *FeatureCompanyResponseData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *FeatureCompanyResponseData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *FeatureCompanyResponseData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *FeatureCompanyResponseData) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *FeatureCompanyResponseData) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPlan

`func (o *FeatureCompanyResponseData) GetPlan() PlanResponseData`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FeatureCompanyResponseData) GetPlanOk() (*PlanResponseData, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FeatureCompanyResponseData) SetPlan(v PlanResponseData)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *FeatureCompanyResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUsage

`func (o *FeatureCompanyResponseData) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FeatureCompanyResponseData) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FeatureCompanyResponseData) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FeatureCompanyResponseData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *FeatureCompanyResponseData) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *FeatureCompanyResponseData) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


