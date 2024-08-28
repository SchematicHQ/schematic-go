# BillingInvoiceResponseData

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
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 
**Subtotal** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingInvoiceResponseData

`func NewBillingInvoiceResponseData(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, createdAt time.Time, currency string, customerExternalId string, environmentId string, externalId string, id string, subtotal int32, updatedAt time.Time, ) *BillingInvoiceResponseData`

NewBillingInvoiceResponseData instantiates a new BillingInvoiceResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceResponseDataWithDefaults

`func NewBillingInvoiceResponseDataWithDefaults() *BillingInvoiceResponseData`

NewBillingInvoiceResponseDataWithDefaults instantiates a new BillingInvoiceResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *BillingInvoiceResponseData) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *BillingInvoiceResponseData) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *BillingInvoiceResponseData) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *BillingInvoiceResponseData) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *BillingInvoiceResponseData) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *BillingInvoiceResponseData) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetAmountRemaining

`func (o *BillingInvoiceResponseData) GetAmountRemaining() int32`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *BillingInvoiceResponseData) GetAmountRemainingOk() (*int32, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *BillingInvoiceResponseData) SetAmountRemaining(v int32)`

SetAmountRemaining sets AmountRemaining field to given value.


### GetCollectionMethod

`func (o *BillingInvoiceResponseData) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *BillingInvoiceResponseData) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *BillingInvoiceResponseData) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetCompanyId

`func (o *BillingInvoiceResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingInvoiceResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingInvoiceResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BillingInvoiceResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *BillingInvoiceResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *BillingInvoiceResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *BillingInvoiceResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingInvoiceResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingInvoiceResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *BillingInvoiceResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoiceResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoiceResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerExternalId

`func (o *BillingInvoiceResponseData) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *BillingInvoiceResponseData) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *BillingInvoiceResponseData) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetDueDate

`func (o *BillingInvoiceResponseData) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *BillingInvoiceResponseData) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *BillingInvoiceResponseData) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *BillingInvoiceResponseData) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *BillingInvoiceResponseData) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *BillingInvoiceResponseData) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetEnvironmentId

`func (o *BillingInvoiceResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BillingInvoiceResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BillingInvoiceResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalId

`func (o *BillingInvoiceResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingInvoiceResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingInvoiceResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetId

`func (o *BillingInvoiceResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoiceResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoiceResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetSubscriptionExternalId

`func (o *BillingInvoiceResponseData) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *BillingInvoiceResponseData) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *BillingInvoiceResponseData) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *BillingInvoiceResponseData) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *BillingInvoiceResponseData) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *BillingInvoiceResponseData) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
### GetSubtotal

`func (o *BillingInvoiceResponseData) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *BillingInvoiceResponseData) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *BillingInvoiceResponseData) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetUpdatedAt

`func (o *BillingInvoiceResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingInvoiceResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingInvoiceResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


