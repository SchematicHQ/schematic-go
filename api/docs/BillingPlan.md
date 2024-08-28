# BillingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**PlanPeriod** | Pointer to **NullableString** |  | [optional] 
**PlanPrice** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBillingPlan

`func NewBillingPlan(id string, name string, ) *BillingPlan`

NewBillingPlan instantiates a new BillingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPlanWithDefaults

`func NewBillingPlanWithDefaults() *BillingPlan`

NewBillingPlanWithDefaults instantiates a new BillingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BillingPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillingPlan) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillingPlan) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *BillingPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPlan) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *BillingPlan) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BillingPlan) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BillingPlan) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BillingPlan) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *BillingPlan) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *BillingPlan) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetName

`func (o *BillingPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingPlan) SetName(v string)`

SetName sets Name field to given value.


### GetPlanPeriod

`func (o *BillingPlan) GetPlanPeriod() string`

GetPlanPeriod returns the PlanPeriod field if non-nil, zero value otherwise.

### GetPlanPeriodOk

`func (o *BillingPlan) GetPlanPeriodOk() (*string, bool)`

GetPlanPeriodOk returns a tuple with the PlanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPeriod

`func (o *BillingPlan) SetPlanPeriod(v string)`

SetPlanPeriod sets PlanPeriod field to given value.

### HasPlanPeriod

`func (o *BillingPlan) HasPlanPeriod() bool`

HasPlanPeriod returns a boolean if a field has been set.

### SetPlanPeriodNil

`func (o *BillingPlan) SetPlanPeriodNil(b bool)`

 SetPlanPeriodNil sets the value for PlanPeriod to be an explicit nil

### UnsetPlanPeriod
`func (o *BillingPlan) UnsetPlanPeriod()`

UnsetPlanPeriod ensures that no value is present for PlanPeriod, not even an explicit nil
### GetPlanPrice

`func (o *BillingPlan) GetPlanPrice() int32`

GetPlanPrice returns the PlanPrice field if non-nil, zero value otherwise.

### GetPlanPriceOk

`func (o *BillingPlan) GetPlanPriceOk() (*int32, bool)`

GetPlanPriceOk returns a tuple with the PlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPrice

`func (o *BillingPlan) SetPlanPrice(v int32)`

SetPlanPrice sets PlanPrice field to given value.

### HasPlanPrice

`func (o *BillingPlan) HasPlanPrice() bool`

HasPlanPrice returns a boolean if a field has been set.

### SetPlanPriceNil

`func (o *BillingPlan) SetPlanPriceNil(b bool)`

 SetPlanPriceNil sets the value for PlanPrice to be an explicit nil

### UnsetPlanPrice
`func (o *BillingPlan) UnsetPlanPrice()`

UnsetPlanPrice ensures that no value is present for PlanPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


