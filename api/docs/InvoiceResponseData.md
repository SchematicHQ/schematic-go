# InvoiceResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **int32** |  | 
**AmountPaid** | **int32** |  | 
**AmountRemaining** | **int32** |  | 
**CollectionMethod** | **string** |  | 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Currency** | **string** |  | 
**CustomerExternalId** | **string** |  | 
**DueDate** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | **string** |  | 
**ExternalId** | **string** |  | 
**Id** | **string** |  | 
**PaymentMethodExternalId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 
**Subtotal** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewInvoiceResponseData

`func NewInvoiceResponseData(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, createdAt time.Time, currency string, customerExternalId string, environmentId string, externalId string, id string, subtotal int32, updatedAt time.Time, ) *InvoiceResponseData`

NewInvoiceResponseData instantiates a new InvoiceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseDataWithDefaults

`func NewInvoiceResponseDataWithDefaults() *InvoiceResponseData`

NewInvoiceResponseDataWithDefaults instantiates a new InvoiceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *InvoiceResponseData) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *InvoiceResponseData) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *InvoiceResponseData) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *InvoiceResponseData) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *InvoiceResponseData) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *InvoiceResponseData) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetAmountRemaining

`func (o *InvoiceResponseData) GetAmountRemaining() int32`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *InvoiceResponseData) GetAmountRemainingOk() (*int32, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *InvoiceResponseData) SetAmountRemaining(v int32)`

SetAmountRemaining sets AmountRemaining field to given value.


### GetCollectionMethod

`func (o *InvoiceResponseData) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *InvoiceResponseData) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *InvoiceResponseData) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetCompanyId

`func (o *InvoiceResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *InvoiceResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *InvoiceResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *InvoiceResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *InvoiceResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *InvoiceResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerExternalId

`func (o *InvoiceResponseData) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *InvoiceResponseData) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *InvoiceResponseData) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetDueDate

`func (o *InvoiceResponseData) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceResponseData) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceResponseData) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceResponseData) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceResponseData) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceResponseData) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetEnvironmentId

`func (o *InvoiceResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *InvoiceResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *InvoiceResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalId

`func (o *InvoiceResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InvoiceResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InvoiceResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetId

`func (o *InvoiceResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentMethodExternalId

`func (o *InvoiceResponseData) GetPaymentMethodExternalId() string`

GetPaymentMethodExternalId returns the PaymentMethodExternalId field if non-nil, zero value otherwise.

### GetPaymentMethodExternalIdOk

`func (o *InvoiceResponseData) GetPaymentMethodExternalIdOk() (*string, bool)`

GetPaymentMethodExternalIdOk returns a tuple with the PaymentMethodExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodExternalId

`func (o *InvoiceResponseData) SetPaymentMethodExternalId(v string)`

SetPaymentMethodExternalId sets PaymentMethodExternalId field to given value.

### HasPaymentMethodExternalId

`func (o *InvoiceResponseData) HasPaymentMethodExternalId() bool`

HasPaymentMethodExternalId returns a boolean if a field has been set.

### SetPaymentMethodExternalIdNil

`func (o *InvoiceResponseData) SetPaymentMethodExternalIdNil(b bool)`

 SetPaymentMethodExternalIdNil sets the value for PaymentMethodExternalId to be an explicit nil

### UnsetPaymentMethodExternalId
`func (o *InvoiceResponseData) UnsetPaymentMethodExternalId()`

UnsetPaymentMethodExternalId ensures that no value is present for PaymentMethodExternalId, not even an explicit nil
### GetSubscriptionExternalId

`func (o *InvoiceResponseData) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *InvoiceResponseData) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *InvoiceResponseData) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *InvoiceResponseData) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *InvoiceResponseData) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *InvoiceResponseData) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
### GetSubtotal

`func (o *InvoiceResponseData) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *InvoiceResponseData) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *InvoiceResponseData) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetUpdatedAt

`func (o *InvoiceResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


