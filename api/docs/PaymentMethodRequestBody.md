# PaymentMethodRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**CardExpMonth** | Pointer to **NullableInt32** |  | [optional] 
**CardExpYear** | Pointer to **NullableInt32** |  | [optional] 
**CardLast4** | Pointer to **NullableString** |  | [optional] 
**CustomerExternalId** | **string** |  | 
**InvoiceExternalId** | Pointer to **NullableString** |  | [optional] 
**PaymentMethodType** | **string** |  | 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentMethodRequestBody

`func NewPaymentMethodRequestBody(customerExternalId string, paymentMethodType string, ) *PaymentMethodRequestBody`

NewPaymentMethodRequestBody instantiates a new PaymentMethodRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodRequestBodyWithDefaults

`func NewPaymentMethodRequestBodyWithDefaults() *PaymentMethodRequestBody`

NewPaymentMethodRequestBodyWithDefaults instantiates a new PaymentMethodRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBrand

`func (o *PaymentMethodRequestBody) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *PaymentMethodRequestBody) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *PaymentMethodRequestBody) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *PaymentMethodRequestBody) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *PaymentMethodRequestBody) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *PaymentMethodRequestBody) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardExpMonth

`func (o *PaymentMethodRequestBody) GetCardExpMonth() int32`

GetCardExpMonth returns the CardExpMonth field if non-nil, zero value otherwise.

### GetCardExpMonthOk

`func (o *PaymentMethodRequestBody) GetCardExpMonthOk() (*int32, bool)`

GetCardExpMonthOk returns a tuple with the CardExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpMonth

`func (o *PaymentMethodRequestBody) SetCardExpMonth(v int32)`

SetCardExpMonth sets CardExpMonth field to given value.

### HasCardExpMonth

`func (o *PaymentMethodRequestBody) HasCardExpMonth() bool`

HasCardExpMonth returns a boolean if a field has been set.

### SetCardExpMonthNil

`func (o *PaymentMethodRequestBody) SetCardExpMonthNil(b bool)`

 SetCardExpMonthNil sets the value for CardExpMonth to be an explicit nil

### UnsetCardExpMonth
`func (o *PaymentMethodRequestBody) UnsetCardExpMonth()`

UnsetCardExpMonth ensures that no value is present for CardExpMonth, not even an explicit nil
### GetCardExpYear

`func (o *PaymentMethodRequestBody) GetCardExpYear() int32`

GetCardExpYear returns the CardExpYear field if non-nil, zero value otherwise.

### GetCardExpYearOk

`func (o *PaymentMethodRequestBody) GetCardExpYearOk() (*int32, bool)`

GetCardExpYearOk returns a tuple with the CardExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpYear

`func (o *PaymentMethodRequestBody) SetCardExpYear(v int32)`

SetCardExpYear sets CardExpYear field to given value.

### HasCardExpYear

`func (o *PaymentMethodRequestBody) HasCardExpYear() bool`

HasCardExpYear returns a boolean if a field has been set.

### SetCardExpYearNil

`func (o *PaymentMethodRequestBody) SetCardExpYearNil(b bool)`

 SetCardExpYearNil sets the value for CardExpYear to be an explicit nil

### UnsetCardExpYear
`func (o *PaymentMethodRequestBody) UnsetCardExpYear()`

UnsetCardExpYear ensures that no value is present for CardExpYear, not even an explicit nil
### GetCardLast4

`func (o *PaymentMethodRequestBody) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *PaymentMethodRequestBody) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *PaymentMethodRequestBody) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.

### HasCardLast4

`func (o *PaymentMethodRequestBody) HasCardLast4() bool`

HasCardLast4 returns a boolean if a field has been set.

### SetCardLast4Nil

`func (o *PaymentMethodRequestBody) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *PaymentMethodRequestBody) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetCustomerExternalId

`func (o *PaymentMethodRequestBody) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *PaymentMethodRequestBody) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *PaymentMethodRequestBody) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetInvoiceExternalId

`func (o *PaymentMethodRequestBody) GetInvoiceExternalId() string`

GetInvoiceExternalId returns the InvoiceExternalId field if non-nil, zero value otherwise.

### GetInvoiceExternalIdOk

`func (o *PaymentMethodRequestBody) GetInvoiceExternalIdOk() (*string, bool)`

GetInvoiceExternalIdOk returns a tuple with the InvoiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExternalId

`func (o *PaymentMethodRequestBody) SetInvoiceExternalId(v string)`

SetInvoiceExternalId sets InvoiceExternalId field to given value.

### HasInvoiceExternalId

`func (o *PaymentMethodRequestBody) HasInvoiceExternalId() bool`

HasInvoiceExternalId returns a boolean if a field has been set.

### SetInvoiceExternalIdNil

`func (o *PaymentMethodRequestBody) SetInvoiceExternalIdNil(b bool)`

 SetInvoiceExternalIdNil sets the value for InvoiceExternalId to be an explicit nil

### UnsetInvoiceExternalId
`func (o *PaymentMethodRequestBody) UnsetInvoiceExternalId()`

UnsetInvoiceExternalId ensures that no value is present for InvoiceExternalId, not even an explicit nil
### GetPaymentMethodType

`func (o *PaymentMethodRequestBody) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *PaymentMethodRequestBody) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *PaymentMethodRequestBody) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.


### GetSubscriptionExternalId

`func (o *PaymentMethodRequestBody) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *PaymentMethodRequestBody) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *PaymentMethodRequestBody) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *PaymentMethodRequestBody) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *PaymentMethodRequestBody) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *PaymentMethodRequestBody) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


