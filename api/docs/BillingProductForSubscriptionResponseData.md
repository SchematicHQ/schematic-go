# BillingProductForSubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Currency** | **string** |  | 
**EnvironmentId** | **string** |  | 
**ExternalId** | **string** |  | 
**Id** | **string** |  | 
**Interval** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Price** | **float32** |  | 
**Quantity** | **float32** |  | 
**SubscriptionId** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingProductForSubscriptionResponseData

`func NewBillingProductForSubscriptionResponseData(accountId string, createdAt time.Time, currency string, environmentId string, externalId string, id string, name string, price float32, quantity float32, subscriptionId string, updatedAt time.Time, ) *BillingProductForSubscriptionResponseData`

NewBillingProductForSubscriptionResponseData instantiates a new BillingProductForSubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingProductForSubscriptionResponseDataWithDefaults

`func NewBillingProductForSubscriptionResponseDataWithDefaults() *BillingProductForSubscriptionResponseData`

NewBillingProductForSubscriptionResponseDataWithDefaults instantiates a new BillingProductForSubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BillingProductForSubscriptionResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BillingProductForSubscriptionResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BillingProductForSubscriptionResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *BillingProductForSubscriptionResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingProductForSubscriptionResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingProductForSubscriptionResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *BillingProductForSubscriptionResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingProductForSubscriptionResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingProductForSubscriptionResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEnvironmentId

`func (o *BillingProductForSubscriptionResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BillingProductForSubscriptionResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BillingProductForSubscriptionResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalId

`func (o *BillingProductForSubscriptionResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingProductForSubscriptionResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingProductForSubscriptionResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetId

`func (o *BillingProductForSubscriptionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingProductForSubscriptionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingProductForSubscriptionResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetInterval

`func (o *BillingProductForSubscriptionResponseData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BillingProductForSubscriptionResponseData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BillingProductForSubscriptionResponseData) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BillingProductForSubscriptionResponseData) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *BillingProductForSubscriptionResponseData) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *BillingProductForSubscriptionResponseData) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetName

`func (o *BillingProductForSubscriptionResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingProductForSubscriptionResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingProductForSubscriptionResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *BillingProductForSubscriptionResponseData) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingProductForSubscriptionResponseData) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingProductForSubscriptionResponseData) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *BillingProductForSubscriptionResponseData) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingProductForSubscriptionResponseData) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingProductForSubscriptionResponseData) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetSubscriptionId

`func (o *BillingProductForSubscriptionResponseData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *BillingProductForSubscriptionResponseData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *BillingProductForSubscriptionResponseData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetUpdatedAt

`func (o *BillingProductForSubscriptionResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingProductForSubscriptionResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingProductForSubscriptionResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


