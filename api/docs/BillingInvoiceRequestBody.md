# BillingInvoiceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **int32** |  | 
**AmountPaid** | **int32** |  | 
**AmountRemaining** | **int32** |  | 
**CollectionMethod** | **string** |  | 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**Currency** | **string** |  | 
**CustomerExternalId** | **string** |  | 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 
**Subtotal** | **int32** |  | 

## Methods

### NewBillingInvoiceRequestBody

`func NewBillingInvoiceRequestBody(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, currency string, customerExternalId string, subtotal int32, ) *BillingInvoiceRequestBody`

NewBillingInvoiceRequestBody instantiates a new BillingInvoiceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceRequestBodyWithDefaults

`func NewBillingInvoiceRequestBodyWithDefaults() *BillingInvoiceRequestBody`

NewBillingInvoiceRequestBodyWithDefaults instantiates a new BillingInvoiceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *BillingInvoiceRequestBody) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *BillingInvoiceRequestBody) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *BillingInvoiceRequestBody) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *BillingInvoiceRequestBody) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *BillingInvoiceRequestBody) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *BillingInvoiceRequestBody) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetAmountRemaining

`func (o *BillingInvoiceRequestBody) GetAmountRemaining() int32`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *BillingInvoiceRequestBody) GetAmountRemainingOk() (*int32, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *BillingInvoiceRequestBody) SetAmountRemaining(v int32)`

SetAmountRemaining sets AmountRemaining field to given value.


### GetCollectionMethod

`func (o *BillingInvoiceRequestBody) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *BillingInvoiceRequestBody) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *BillingInvoiceRequestBody) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetCompanyId

`func (o *BillingInvoiceRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingInvoiceRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingInvoiceRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BillingInvoiceRequestBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *BillingInvoiceRequestBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *BillingInvoiceRequestBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCurrency

`func (o *BillingInvoiceRequestBody) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoiceRequestBody) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoiceRequestBody) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerExternalId

`func (o *BillingInvoiceRequestBody) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *BillingInvoiceRequestBody) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *BillingInvoiceRequestBody) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetDueDate

`func (o *BillingInvoiceRequestBody) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingInvoiceRequestBody) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingInvoiceRequestBody) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingInvoiceRequestBody) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *BillingInvoiceRequestBody) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *BillingInvoiceRequestBody) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetSubscriptionExternalId

`func (o *BillingInvoiceRequestBody) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *BillingInvoiceRequestBody) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *BillingInvoiceRequestBody) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *BillingInvoiceRequestBody) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *BillingInvoiceRequestBody) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *BillingInvoiceRequestBody) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
### GetSubtotal

`func (o *BillingInvoiceRequestBody) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *BillingInvoiceRequestBody) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *BillingInvoiceRequestBody) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


