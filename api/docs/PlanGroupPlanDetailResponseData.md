# PlanGroupPlanDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | Pointer to **NullableString** |  | [optional] 
**BillingProduct** | Pointer to [**BillingProductDetailResponseData**](BillingProductDetailResponseData.md) |  | [optional] 
**CompanyCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Entitlements** | [**[]PlanEntitlementResponseData**](PlanEntitlementResponseData.md) |  | 
**Features** | [**[]FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | 
**Icon** | **string** |  | 
**Id** | **string** |  | 
**IsDefault** | **bool** |  | 
**MonthlyPrice** | Pointer to [**BillingPriceResponseData**](BillingPriceResponseData.md) |  | [optional] 
**Name** | **string** |  | 
**PlanType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**YearlyPrice** | Pointer to [**BillingPriceResponseData**](BillingPriceResponseData.md) |  | [optional] 

## Methods

### NewPlanGroupPlanDetailResponseData

`func NewPlanGroupPlanDetailResponseData(companyCount int32, createdAt time.Time, description string, entitlements []PlanEntitlementResponseData, features []FeatureDetailResponseData, icon string, id string, isDefault bool, name string, planType string, updatedAt time.Time, ) *PlanGroupPlanDetailResponseData`

NewPlanGroupPlanDetailResponseData instantiates a new PlanGroupPlanDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanGroupPlanDetailResponseDataWithDefaults

`func NewPlanGroupPlanDetailResponseDataWithDefaults() *PlanGroupPlanDetailResponseData`

NewPlanGroupPlanDetailResponseDataWithDefaults instantiates a new PlanGroupPlanDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceType

`func (o *PlanGroupPlanDetailResponseData) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *PlanGroupPlanDetailResponseData) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *PlanGroupPlanDetailResponseData) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.

### HasAudienceType

`func (o *PlanGroupPlanDetailResponseData) HasAudienceType() bool`

HasAudienceType returns a boolean if a field has been set.

### SetAudienceTypeNil

`func (o *PlanGroupPlanDetailResponseData) SetAudienceTypeNil(b bool)`

 SetAudienceTypeNil sets the value for AudienceType to be an explicit nil

### UnsetAudienceType
`func (o *PlanGroupPlanDetailResponseData) UnsetAudienceType()`

UnsetAudienceType ensures that no value is present for AudienceType, not even an explicit nil
### GetBillingProduct

`func (o *PlanGroupPlanDetailResponseData) GetBillingProduct() BillingProductDetailResponseData`

GetBillingProduct returns the BillingProduct field if non-nil, zero value otherwise.

### GetBillingProductOk

`func (o *PlanGroupPlanDetailResponseData) GetBillingProductOk() (*BillingProductDetailResponseData, bool)`

GetBillingProductOk returns a tuple with the BillingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProduct

`func (o *PlanGroupPlanDetailResponseData) SetBillingProduct(v BillingProductDetailResponseData)`

SetBillingProduct sets BillingProduct field to given value.

### HasBillingProduct

`func (o *PlanGroupPlanDetailResponseData) HasBillingProduct() bool`

HasBillingProduct returns a boolean if a field has been set.

### GetCompanyCount

`func (o *PlanGroupPlanDetailResponseData) GetCompanyCount() int32`

GetCompanyCount returns the CompanyCount field if non-nil, zero value otherwise.

### GetCompanyCountOk

`func (o *PlanGroupPlanDetailResponseData) GetCompanyCountOk() (*int32, bool)`

GetCompanyCountOk returns a tuple with the CompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCount

`func (o *PlanGroupPlanDetailResponseData) SetCompanyCount(v int32)`

SetCompanyCount sets CompanyCount field to given value.


### GetCreatedAt

`func (o *PlanGroupPlanDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanGroupPlanDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanGroupPlanDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *PlanGroupPlanDetailResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanGroupPlanDetailResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanGroupPlanDetailResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntitlements

`func (o *PlanGroupPlanDetailResponseData) GetEntitlements() []PlanEntitlementResponseData`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PlanGroupPlanDetailResponseData) GetEntitlementsOk() (*[]PlanEntitlementResponseData, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PlanGroupPlanDetailResponseData) SetEntitlements(v []PlanEntitlementResponseData)`

SetEntitlements sets Entitlements field to given value.


### GetFeatures

`func (o *PlanGroupPlanDetailResponseData) GetFeatures() []FeatureDetailResponseData`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PlanGroupPlanDetailResponseData) GetFeaturesOk() (*[]FeatureDetailResponseData, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PlanGroupPlanDetailResponseData) SetFeatures(v []FeatureDetailResponseData)`

SetFeatures sets Features field to given value.


### GetIcon

`func (o *PlanGroupPlanDetailResponseData) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PlanGroupPlanDetailResponseData) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PlanGroupPlanDetailResponseData) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetId

`func (o *PlanGroupPlanDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanGroupPlanDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanGroupPlanDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *PlanGroupPlanDetailResponseData) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PlanGroupPlanDetailResponseData) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PlanGroupPlanDetailResponseData) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetMonthlyPrice

`func (o *PlanGroupPlanDetailResponseData) GetMonthlyPrice() BillingPriceResponseData`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *PlanGroupPlanDetailResponseData) GetMonthlyPriceOk() (*BillingPriceResponseData, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *PlanGroupPlanDetailResponseData) SetMonthlyPrice(v BillingPriceResponseData)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *PlanGroupPlanDetailResponseData) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetName

`func (o *PlanGroupPlanDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanGroupPlanDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanGroupPlanDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanType

`func (o *PlanGroupPlanDetailResponseData) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *PlanGroupPlanDetailResponseData) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *PlanGroupPlanDetailResponseData) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetUpdatedAt

`func (o *PlanGroupPlanDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanGroupPlanDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanGroupPlanDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetYearlyPrice

`func (o *PlanGroupPlanDetailResponseData) GetYearlyPrice() BillingPriceResponseData`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *PlanGroupPlanDetailResponseData) GetYearlyPriceOk() (*BillingPriceResponseData, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *PlanGroupPlanDetailResponseData) SetYearlyPrice(v BillingPriceResponseData)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *PlanGroupPlanDetailResponseData) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


