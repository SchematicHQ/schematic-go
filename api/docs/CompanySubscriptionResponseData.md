# CompanySubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerExternalId** | **string** |  | 
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**Interval** | **string** |  | 
**LatestInvoice** | Pointer to [**InvoiceResponseData**](InvoiceResponseData.md) |  | [optional] 
**PaymentMethod** | Pointer to [**PaymentMethodResponseData**](PaymentMethodResponseData.md) |  | [optional] 
**Products** | [**[]BillingProductForSubscriptionResponseData**](BillingProductForSubscriptionResponseData.md) |  | 
**SubscriptionExternalId** | **string** |  | 
**TotalPrice** | **int32** |  | 

## Methods

### NewCompanySubscriptionResponseData

`func NewCompanySubscriptionResponseData(customerExternalId string, interval string, products []BillingProductForSubscriptionResponseData, subscriptionExternalId string, totalPrice int32, ) *CompanySubscriptionResponseData`

NewCompanySubscriptionResponseData instantiates a new CompanySubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySubscriptionResponseDataWithDefaults

`func NewCompanySubscriptionResponseDataWithDefaults() *CompanySubscriptionResponseData`

NewCompanySubscriptionResponseDataWithDefaults instantiates a new CompanySubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerExternalId

`func (o *CompanySubscriptionResponseData) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *CompanySubscriptionResponseData) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *CompanySubscriptionResponseData) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetExpiredAt

`func (o *CompanySubscriptionResponseData) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CompanySubscriptionResponseData) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CompanySubscriptionResponseData) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CompanySubscriptionResponseData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *CompanySubscriptionResponseData) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *CompanySubscriptionResponseData) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetInterval

`func (o *CompanySubscriptionResponseData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CompanySubscriptionResponseData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CompanySubscriptionResponseData) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetLatestInvoice

`func (o *CompanySubscriptionResponseData) GetLatestInvoice() InvoiceResponseData`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *CompanySubscriptionResponseData) GetLatestInvoiceOk() (*InvoiceResponseData, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *CompanySubscriptionResponseData) SetLatestInvoice(v InvoiceResponseData)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *CompanySubscriptionResponseData) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CompanySubscriptionResponseData) GetPaymentMethod() PaymentMethodResponseData`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CompanySubscriptionResponseData) GetPaymentMethodOk() (*PaymentMethodResponseData, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CompanySubscriptionResponseData) SetPaymentMethod(v PaymentMethodResponseData)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CompanySubscriptionResponseData) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetProducts

`func (o *CompanySubscriptionResponseData) GetProducts() []BillingProductForSubscriptionResponseData`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CompanySubscriptionResponseData) GetProductsOk() (*[]BillingProductForSubscriptionResponseData, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CompanySubscriptionResponseData) SetProducts(v []BillingProductForSubscriptionResponseData)`

SetProducts sets Products field to given value.


### GetSubscriptionExternalId

`func (o *CompanySubscriptionResponseData) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *CompanySubscriptionResponseData) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *CompanySubscriptionResponseData) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.


### GetTotalPrice

`func (o *CompanySubscriptionResponseData) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *CompanySubscriptionResponseData) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *CompanySubscriptionResponseData) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


