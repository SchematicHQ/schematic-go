# PlanDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | **string** |  | 
**BillingProduct** | Pointer to [**BillingProductDetailResponseData**](BillingProductDetailResponseData.md) |  | [optional] 
**CompanyCount** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Features** | [**[]FeatureDetailResponseData**](FeatureDetailResponseData.md) |  | 
**Icon** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**PlanType** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPlanDetailResponseData

`func NewPlanDetailResponseData(audienceType string, companyCount int32, createdAt time.Time, description string, features []FeatureDetailResponseData, icon string, id string, name string, planType string, updatedAt time.Time, ) *PlanDetailResponseData`

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


### GetBillingProduct

`func (o *PlanDetailResponseData) GetBillingProduct() BillingProductDetailResponseData`

GetBillingProduct returns the BillingProduct field if non-nil, zero value otherwise.

### GetBillingProductOk

`func (o *PlanDetailResponseData) GetBillingProductOk() (*BillingProductDetailResponseData, bool)`

GetBillingProductOk returns a tuple with the BillingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProduct

`func (o *PlanDetailResponseData) SetBillingProduct(v BillingProductDetailResponseData)`

SetBillingProduct sets BillingProduct field to given value.

### HasBillingProduct

`func (o *PlanDetailResponseData) HasBillingProduct() bool`

HasBillingProduct returns a boolean if a field has been set.

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


### GetIcon

`func (o *PlanDetailResponseData) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PlanDetailResponseData) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PlanDetailResponseData) SetIcon(v string)`

SetIcon sets Icon field to given value.


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


