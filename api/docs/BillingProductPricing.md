# BillingProductPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | Pointer to **NullableString** |  | [optional] 
**Price** | **int32** |  | 
**PriceExternalId** | Pointer to **NullableString** |  | [optional] 
**ProductExternalId** | **string** |  | 

## Methods

### NewBillingProductPricing

`func NewBillingProductPricing(price int32, productExternalId string, ) *BillingProductPricing`

NewBillingProductPricing instantiates a new BillingProductPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProductPricingWithDefaults

`func NewBillingProductPricingWithDefaults() *BillingProductPricing`

NewBillingProductPricingWithDefaults instantiates a new BillingProductPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *BillingProductPricing) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BillingProductPricing) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BillingProductPricing) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BillingProductPricing) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *BillingProductPricing) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *BillingProductPricing) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetPrice

`func (o *BillingProductPricing) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingProductPricing) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingProductPricing) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetPriceExternalId

`func (o *BillingProductPricing) GetPriceExternalId() string`

GetPriceExternalId returns the PriceExternalId field if non-nil, zero value otherwise.

### GetPriceExternalIdOk

`func (o *BillingProductPricing) GetPriceExternalIdOk() (*string, bool)`

GetPriceExternalIdOk returns a tuple with the PriceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExternalId

`func (o *BillingProductPricing) SetPriceExternalId(v string)`

SetPriceExternalId sets PriceExternalId field to given value.

### HasPriceExternalId

`func (o *BillingProductPricing) HasPriceExternalId() bool`

HasPriceExternalId returns a boolean if a field has been set.

### SetPriceExternalIdNil

`func (o *BillingProductPricing) SetPriceExternalIdNil(b bool)`

 SetPriceExternalIdNil sets the value for PriceExternalId to be an explicit nil

### UnsetPriceExternalId
`func (o *BillingProductPricing) UnsetPriceExternalId()`

UnsetPriceExternalId ensures that no value is present for PriceExternalId, not even an explicit nil
### GetProductExternalId

`func (o *BillingProductPricing) GetProductExternalId() string`

GetProductExternalId returns the ProductExternalId field if non-nil, zero value otherwise.

### GetProductExternalIdOk

`func (o *BillingProductPricing) GetProductExternalIdOk() (*string, bool)`

GetProductExternalIdOk returns a tuple with the ProductExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalId

`func (o *BillingProductPricing) SetProductExternalId(v string)`

SetProductExternalId sets ProductExternalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


