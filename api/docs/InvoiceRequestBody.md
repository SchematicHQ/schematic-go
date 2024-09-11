# InvoiceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **int32** |  | 
**AmountPaid** | **int32** |  | 
**AmountRemaining** | **int32** |  | 
**CollectionMethod** | **string** |  | 
**Currency** | **string** |  | 
**CustomerExternalId** | **string** |  | 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**PaymentMethodExternalId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 
**Subtotal** | **int32** |  | 

## Methods

### NewInvoiceRequestBody

`func NewInvoiceRequestBody(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, currency string, customerExternalId string, subtotal int32, ) *InvoiceRequestBody`

NewInvoiceRequestBody instantiates a new InvoiceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRequestBodyWithDefaults

`func NewInvoiceRequestBodyWithDefaults() *InvoiceRequestBody`

NewInvoiceRequestBodyWithDefaults instantiates a new InvoiceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *InvoiceRequestBody) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *InvoiceRequestBody) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *InvoiceRequestBody) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *InvoiceRequestBody) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *InvoiceRequestBody) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *InvoiceRequestBody) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetAmountRemaining

`func (o *InvoiceRequestBody) GetAmountRemaining() int32`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *InvoiceRequestBody) GetAmountRemainingOk() (*int32, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *InvoiceRequestBody) SetAmountRemaining(v int32)`

SetAmountRemaining sets AmountRemaining field to given value.


### GetCollectionMethod

`func (o *InvoiceRequestBody) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *InvoiceRequestBody) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *InvoiceRequestBody) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetCurrency

`func (o *InvoiceRequestBody) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceRequestBody) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceRequestBody) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerExternalId

`func (o *InvoiceRequestBody) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *InvoiceRequestBody) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *InvoiceRequestBody) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetDueDate

`func (o *InvoiceRequestBody) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceRequestBody) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceRequestBody) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceRequestBody) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceRequestBody) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceRequestBody) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaymentMethodExternalId

`func (o *InvoiceRequestBody) GetPaymentMethodExternalId() string`

GetPaymentMethodExternalId returns the PaymentMethodExternalId field if non-nil, zero value otherwise.

### GetPaymentMethodExternalIdOk

`func (o *InvoiceRequestBody) GetPaymentMethodExternalIdOk() (*string, bool)`

GetPaymentMethodExternalIdOk returns a tuple with the PaymentMethodExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodExternalId

`func (o *InvoiceRequestBody) SetPaymentMethodExternalId(v string)`

SetPaymentMethodExternalId sets PaymentMethodExternalId field to given value.

### HasPaymentMethodExternalId

`func (o *InvoiceRequestBody) HasPaymentMethodExternalId() bool`

HasPaymentMethodExternalId returns a boolean if a field has been set.

### SetPaymentMethodExternalIdNil

`func (o *InvoiceRequestBody) SetPaymentMethodExternalIdNil(b bool)`

 SetPaymentMethodExternalIdNil sets the value for PaymentMethodExternalId to be an explicit nil

### UnsetPaymentMethodExternalId
`func (o *InvoiceRequestBody) UnsetPaymentMethodExternalId()`

UnsetPaymentMethodExternalId ensures that no value is present for PaymentMethodExternalId, not even an explicit nil
### GetSubscriptionExternalId

`func (o *InvoiceRequestBody) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *InvoiceRequestBody) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *InvoiceRequestBody) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *InvoiceRequestBody) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *InvoiceRequestBody) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *InvoiceRequestBody) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
### GetSubtotal

`func (o *InvoiceRequestBody) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceRequestBody) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceRequestBody) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


