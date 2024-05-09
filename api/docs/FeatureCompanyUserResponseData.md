# FeatureCompanyUserResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **bool** | Whether further usage is permitted. | 
**Allocation** | Pointer to **NullableInt32** | The maximum amount of usage that is permitted; a null value indicates that unlimited usage is permitted. | [optional] 
**Company** | Pointer to [**CompanyDetailResponseData**](CompanyDetailResponseData.md) |  | [optional] 
**EntitlementId** | **string** |  | 
**EntitlementType** | **string** |  | 
**Feature** | Pointer to [**FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | [optional] 
**Period** | Pointer to **NullableString** | The period over which usage is measured. | [optional] 
**Plan** | Pointer to [**PlanResponseData**](PlanResponseData.md) |  | [optional] 
**Usage** | Pointer to **NullableInt32** | The amount of usage that has been consumed; a null value indicates that usage is not being measured. | [optional] 
**User** | Pointer to [**UserResponseData**](UserResponseData.md) |  | [optional] 

## Methods

### NewFeatureCompanyUserResponseData

`func NewFeatureCompanyUserResponseData(access bool, entitlementId string, entitlementType string, ) *FeatureCompanyUserResponseData`

NewFeatureCompanyUserResponseData instantiates a new FeatureCompanyUserResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureCompanyUserResponseDataWithDefaults

`func NewFeatureCompanyUserResponseDataWithDefaults() *FeatureCompanyUserResponseData`

NewFeatureCompanyUserResponseDataWithDefaults instantiates a new FeatureCompanyUserResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FeatureCompanyUserResponseData) GetAccess() bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FeatureCompanyUserResponseData) GetAccessOk() (*bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FeatureCompanyUserResponseData) SetAccess(v bool)`

SetAccess sets Access field to given value.


### GetAllocation

`func (o *FeatureCompanyUserResponseData) GetAllocation() int32`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *FeatureCompanyUserResponseData) GetAllocationOk() (*int32, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *FeatureCompanyUserResponseData) SetAllocation(v int32)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *FeatureCompanyUserResponseData) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### SetAllocationNil

`func (o *FeatureCompanyUserResponseData) SetAllocationNil(b bool)`

 SetAllocationNil sets the value for Allocation to be an explicit nil

### UnsetAllocation
`func (o *FeatureCompanyUserResponseData) UnsetAllocation()`

UnsetAllocation ensures that no value is present for Allocation, not even an explicit nil
### GetCompany

`func (o *FeatureCompanyUserResponseData) GetCompany() CompanyDetailResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FeatureCompanyUserResponseData) GetCompanyOk() (*CompanyDetailResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FeatureCompanyUserResponseData) SetCompany(v CompanyDetailResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FeatureCompanyUserResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEntitlementId

`func (o *FeatureCompanyUserResponseData) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *FeatureCompanyUserResponseData) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *FeatureCompanyUserResponseData) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetEntitlementType

`func (o *FeatureCompanyUserResponseData) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *FeatureCompanyUserResponseData) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *FeatureCompanyUserResponseData) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetFeature

`func (o *FeatureCompanyUserResponseData) GetFeature() FeatureDetailResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FeatureCompanyUserResponseData) GetFeatureOk() (*FeatureDetailResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FeatureCompanyUserResponseData) SetFeature(v FeatureDetailResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FeatureCompanyUserResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetPeriod

`func (o *FeatureCompanyUserResponseData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *FeatureCompanyUserResponseData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *FeatureCompanyUserResponseData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *FeatureCompanyUserResponseData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *FeatureCompanyUserResponseData) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *FeatureCompanyUserResponseData) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetPlan

`func (o *FeatureCompanyUserResponseData) GetPlan() PlanResponseData`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FeatureCompanyUserResponseData) GetPlanOk() (*PlanResponseData, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FeatureCompanyUserResponseData) SetPlan(v PlanResponseData)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *FeatureCompanyUserResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetUsage

`func (o *FeatureCompanyUserResponseData) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FeatureCompanyUserResponseData) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FeatureCompanyUserResponseData) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FeatureCompanyUserResponseData) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *FeatureCompanyUserResponseData) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *FeatureCompanyUserResponseData) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUser

`func (o *FeatureCompanyUserResponseData) GetUser() UserResponseData`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FeatureCompanyUserResponseData) GetUserOk() (*UserResponseData, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FeatureCompanyUserResponseData) SetUser(v UserResponseData)`

SetUser sets User field to given value.

### HasUser

`func (o *FeatureCompanyUserResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


