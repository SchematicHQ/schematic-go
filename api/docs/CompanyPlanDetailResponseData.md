# CompanyPlanDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | Pointer to **NullableString** |  | [optional] 
**BillingProduct** | Pointer to [**BillingProductDetailResponseData**](BillingProductDetailResponseData.md) |  | [optional] 
**CompanyCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Current** | **bool** |  | 
**Description** | **string** |  | 
**Entitlements** | [**[]PlanEntitlementResponseData**](PlanEntitlementResponseData.md) |  | 
**Features** | [**[]FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | 
**Icon** | **string** |  | 
**Id** | **string** |  | 
**MonthlyPrice** | Pointer to [**BillingPriceResponseData**](BillingPriceResponseData.md) |  | [optional] 
**Name** | **string** |  | 
**PlanType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Valid** | **bool** |  | 
**YearlyPrice** | Pointer to [**BillingPriceResponseData**](BillingPriceResponseData.md) |  | [optional] 

## Methods

### NewCompanyPlanDetailResponseData

`func NewCompanyPlanDetailResponseData(companyCount int32, createdAt time.Time, current bool, description string, entitlements []PlanEntitlementResponseData, features []FeatureDetailResponseData, icon string, id string, name string, planType string, updatedAt time.Time, valid bool, ) *CompanyPlanDetailResponseData`

NewCompanyPlanDetailResponseData instantiates a new CompanyPlanDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyPlanDetailResponseDataWithDefaults

`func NewCompanyPlanDetailResponseDataWithDefaults() *CompanyPlanDetailResponseData`

NewCompanyPlanDetailResponseDataWithDefaults instantiates a new CompanyPlanDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceType

`func (o *CompanyPlanDetailResponseData) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *CompanyPlanDetailResponseData) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *CompanyPlanDetailResponseData) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.

### HasAudienceType

`func (o *CompanyPlanDetailResponseData) HasAudienceType() bool`

HasAudienceType returns a boolean if a field has been set.

### SetAudienceTypeNil

`func (o *CompanyPlanDetailResponseData) SetAudienceTypeNil(b bool)`

 SetAudienceTypeNil sets the value for AudienceType to be an explicit nil

### UnsetAudienceType
`func (o *CompanyPlanDetailResponseData) UnsetAudienceType()`

UnsetAudienceType ensures that no value is present for AudienceType, not even an explicit nil
### GetBillingProduct

`func (o *CompanyPlanDetailResponseData) GetBillingProduct() BillingProductDetailResponseData`

GetBillingProduct returns the BillingProduct field if non-nil, zero value otherwise.

### GetBillingProductOk

`func (o *CompanyPlanDetailResponseData) GetBillingProductOk() (*BillingProductDetailResponseData, bool)`

GetBillingProductOk returns a tuple with the BillingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProduct

`func (o *CompanyPlanDetailResponseData) SetBillingProduct(v BillingProductDetailResponseData)`

SetBillingProduct sets BillingProduct field to given value.

### HasBillingProduct

`func (o *CompanyPlanDetailResponseData) HasBillingProduct() bool`

HasBillingProduct returns a boolean if a field has been set.

### GetCompanyCount

`func (o *CompanyPlanDetailResponseData) GetCompanyCount() int32`

GetCompanyCount returns the CompanyCount field if non-nil, zero value otherwise.

### GetCompanyCountOk

`func (o *CompanyPlanDetailResponseData) GetCompanyCountOk() (*int32, bool)`

GetCompanyCountOk returns a tuple with the CompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCount

`func (o *CompanyPlanDetailResponseData) SetCompanyCount(v int32)`

SetCompanyCount sets CompanyCount field to given value.


### GetCreatedAt

`func (o *CompanyPlanDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyPlanDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyPlanDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrent

`func (o *CompanyPlanDetailResponseData) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CompanyPlanDetailResponseData) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CompanyPlanDetailResponseData) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetDescription

`func (o *CompanyPlanDetailResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompanyPlanDetailResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompanyPlanDetailResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEntitlements

`func (o *CompanyPlanDetailResponseData) GetEntitlements() []PlanEntitlementResponseData`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CompanyPlanDetailResponseData) GetEntitlementsOk() (*[]PlanEntitlementResponseData, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CompanyPlanDetailResponseData) SetEntitlements(v []PlanEntitlementResponseData)`

SetEntitlements sets Entitlements field to given value.


### GetFeatures

`func (o *CompanyPlanDetailResponseData) GetFeatures() []FeatureDetailResponseData`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CompanyPlanDetailResponseData) GetFeaturesOk() (*[]FeatureDetailResponseData, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CompanyPlanDetailResponseData) SetFeatures(v []FeatureDetailResponseData)`

SetFeatures sets Features field to given value.


### GetIcon

`func (o *CompanyPlanDetailResponseData) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CompanyPlanDetailResponseData) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CompanyPlanDetailResponseData) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetId

`func (o *CompanyPlanDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyPlanDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyPlanDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMonthlyPrice

`func (o *CompanyPlanDetailResponseData) GetMonthlyPrice() BillingPriceResponseData`

GetMonthlyPrice returns the MonthlyPrice field if non-nil, zero value otherwise.

### GetMonthlyPriceOk

`func (o *CompanyPlanDetailResponseData) GetMonthlyPriceOk() (*BillingPriceResponseData, bool)`

GetMonthlyPriceOk returns a tuple with the MonthlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPrice

`func (o *CompanyPlanDetailResponseData) SetMonthlyPrice(v BillingPriceResponseData)`

SetMonthlyPrice sets MonthlyPrice field to given value.

### HasMonthlyPrice

`func (o *CompanyPlanDetailResponseData) HasMonthlyPrice() bool`

HasMonthlyPrice returns a boolean if a field has been set.

### GetName

`func (o *CompanyPlanDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyPlanDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyPlanDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanType

`func (o *CompanyPlanDetailResponseData) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CompanyPlanDetailResponseData) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CompanyPlanDetailResponseData) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetUpdatedAt

`func (o *CompanyPlanDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyPlanDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyPlanDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValid

`func (o *CompanyPlanDetailResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CompanyPlanDetailResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CompanyPlanDetailResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetYearlyPrice

`func (o *CompanyPlanDetailResponseData) GetYearlyPrice() BillingPriceResponseData`

GetYearlyPrice returns the YearlyPrice field if non-nil, zero value otherwise.

### GetYearlyPriceOk

`func (o *CompanyPlanDetailResponseData) GetYearlyPriceOk() (*BillingPriceResponseData, bool)`

GetYearlyPriceOk returns a tuple with the YearlyPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPrice

`func (o *CompanyPlanDetailResponseData) SetYearlyPrice(v BillingPriceResponseData)`

SetYearlyPrice sets YearlyPrice field to given value.

### HasYearlyPrice

`func (o *CompanyPlanDetailResponseData) HasYearlyPrice() bool`

HasYearlyPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


