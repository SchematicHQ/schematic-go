# BillingProductDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prices** | [**[]BillingPriceResponseData**](BillingPriceResponseData.md) |  | 
**AccountId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Currency** | **string** |  | 
**EnvironmentId** | **string** |  | 
**ExternalId** | **string** |  | 
**Name** | **string** |  | 
**Price** | **float32** |  | 
**ProductId** | **string** |  | 
**Quantity** | **float32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingProductDetailResponseData

`func NewBillingProductDetailResponseData(prices []BillingPriceResponseData, accountId string, createdAt time.Time, currency string, environmentId string, externalId string, name string, price float32, productId string, quantity float32, updatedAt time.Time, ) *BillingProductDetailResponseData`

NewBillingProductDetailResponseData instantiates a new BillingProductDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProductDetailResponseDataWithDefaults

`func NewBillingProductDetailResponseDataWithDefaults() *BillingProductDetailResponseData`

NewBillingProductDetailResponseDataWithDefaults instantiates a new BillingProductDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrices

`func (o *BillingProductDetailResponseData) GetPrices() []BillingPriceResponseData`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *BillingProductDetailResponseData) GetPricesOk() (*[]BillingPriceResponseData, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *BillingProductDetailResponseData) SetPrices(v []BillingPriceResponseData)`

SetPrices sets Prices field to given value.


### GetAccountId

`func (o *BillingProductDetailResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingProductDetailResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingProductDetailResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *BillingProductDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingProductDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingProductDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *BillingProductDetailResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingProductDetailResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingProductDetailResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEnvironmentId

`func (o *BillingProductDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BillingProductDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BillingProductDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalId

`func (o *BillingProductDetailResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingProductDetailResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingProductDetailResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *BillingProductDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProductDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProductDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *BillingProductDetailResponseData) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingProductDetailResponseData) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingProductDetailResponseData) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetProductId

`func (o *BillingProductDetailResponseData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *BillingProductDetailResponseData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *BillingProductDetailResponseData) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *BillingProductDetailResponseData) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingProductDetailResponseData) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingProductDetailResponseData) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUpdatedAt

`func (o *BillingProductDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingProductDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingProductDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


