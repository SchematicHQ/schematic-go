# PlanDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | **string** |  | 
**BillingProducts** | [**[]BillingProductResponseData**](BillingProductResponseData.md) |  | 
**CompanyCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Features** | [**[]FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**PlanType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPlanDetailResponseData

`func NewPlanDetailResponseData(audienceType string, billingProducts []BillingProductResponseData, companyCount int32, createdAt time.Time, description string, features []FeatureDetailResponseData, id string, name string, planType string, updatedAt time.Time, ) *PlanDetailResponseData`

NewPlanDetailResponseData instantiates a new PlanDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanDetailResponseDataWithDefaults

`func NewPlanDetailResponseDataWithDefaults() *PlanDetailResponseData`

NewPlanDetailResponseDataWithDefaults instantiates a new PlanDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceType

`func (o *PlanDetailResponseData) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *PlanDetailResponseData) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *PlanDetailResponseData) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetBillingProducts

`func (o *PlanDetailResponseData) GetBillingProducts() []BillingProductResponseData`

GetBillingProducts returns the BillingProducts field if non-nil, zero value otherwise.

### GetBillingProductsOk

`func (o *PlanDetailResponseData) GetBillingProductsOk() (*[]BillingProductResponseData, bool)`

GetBillingProductsOk returns a tuple with the BillingProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProducts

`func (o *PlanDetailResponseData) SetBillingProducts(v []BillingProductResponseData)`

SetBillingProducts sets BillingProducts field to given value.


### GetCompanyCount

`func (o *PlanDetailResponseData) GetCompanyCount() int32`

GetCompanyCount returns the CompanyCount field if non-nil, zero value otherwise.

### GetCompanyCountOk

`func (o *PlanDetailResponseData) GetCompanyCountOk() (*int32, bool)`

GetCompanyCountOk returns a tuple with the CompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCount

`func (o *PlanDetailResponseData) SetCompanyCount(v int32)`

SetCompanyCount sets CompanyCount field to given value.


### GetCreatedAt

`func (o *PlanDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlanDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlanDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *PlanDetailResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlanDetailResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlanDetailResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatures

`func (o *PlanDetailResponseData) GetFeatures() []FeatureDetailResponseData`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PlanDetailResponseData) GetFeaturesOk() (*[]FeatureDetailResponseData, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PlanDetailResponseData) SetFeatures(v []FeatureDetailResponseData)`

SetFeatures sets Features field to given value.


### GetId

`func (o *PlanDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PlanDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlanType

`func (o *PlanDetailResponseData) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *PlanDetailResponseData) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *PlanDetailResponseData) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetUpdatedAt

`func (o *PlanDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PlanDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PlanDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


