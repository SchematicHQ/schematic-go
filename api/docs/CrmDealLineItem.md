# CrmDealLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Currency** | **string** |  | 
**Description** | **string** |  | 
**DiscountPercentage** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Price** | **float32** |  | 
**Quantity** | **int32** |  | 
**TermMonth** | Pointer to **NullableInt32** |  | [optional] 
**TotalDiscount** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCrmDealLineItem

`func NewCrmDealLineItem(billingFrequency string, createdAt time.Time, currency string, description string, id string, name string, price float32, quantity int32, updatedAt time.Time, ) *CrmDealLineItem`

NewCrmDealLineItem instantiates a new CrmDealLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmDealLineItemWithDefaults

`func NewCrmDealLineItemWithDefaults() *CrmDealLineItem`

NewCrmDealLineItemWithDefaults instantiates a new CrmDealLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *CrmDealLineItem) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *CrmDealLineItem) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *CrmDealLineItem) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetCreatedAt

`func (o *CrmDealLineItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmDealLineItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmDealLineItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *CrmDealLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CrmDealLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CrmDealLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *CrmDealLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CrmDealLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CrmDealLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDiscountPercentage

`func (o *CrmDealLineItem) GetDiscountPercentage() map[string]interface{}`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *CrmDealLineItem) GetDiscountPercentageOk() (*map[string]interface{}, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *CrmDealLineItem) SetDiscountPercentage(v map[string]interface{})`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *CrmDealLineItem) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetId

`func (o *CrmDealLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrmDealLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrmDealLineItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CrmDealLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmDealLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmDealLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *CrmDealLineItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CrmDealLineItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CrmDealLineItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *CrmDealLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CrmDealLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CrmDealLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTermMonth

`func (o *CrmDealLineItem) GetTermMonth() int32`

GetTermMonth returns the TermMonth field if non-nil, zero value otherwise.

### GetTermMonthOk

`func (o *CrmDealLineItem) GetTermMonthOk() (*int32, bool)`

GetTermMonthOk returns a tuple with the TermMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermMonth

`func (o *CrmDealLineItem) SetTermMonth(v int32)`

SetTermMonth sets TermMonth field to given value.

### HasTermMonth

`func (o *CrmDealLineItem) HasTermMonth() bool`

HasTermMonth returns a boolean if a field has been set.

### SetTermMonthNil

`func (o *CrmDealLineItem) SetTermMonthNil(b bool)`

 SetTermMonthNil sets the value for TermMonth to be an explicit nil

### UnsetTermMonth
`func (o *CrmDealLineItem) UnsetTermMonth()`

UnsetTermMonth ensures that no value is present for TermMonth, not even an explicit nil
### GetTotalDiscount

`func (o *CrmDealLineItem) GetTotalDiscount() map[string]interface{}`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *CrmDealLineItem) GetTotalDiscountOk() (*map[string]interface{}, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *CrmDealLineItem) SetTotalDiscount(v map[string]interface{})`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *CrmDealLineItem) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CrmDealLineItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmDealLineItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmDealLineItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


