# BillingPaymentMethodResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**CardExpMonth** | Pointer to **NullableInt32** |  | [optional] 
**CardExpYear** | Pointer to **NullableInt32** |  | [optional] 
**CardLast4** | Pointer to **NullableString** |  | [optional] 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CustomerExternalId** | **string** |  | 
**EnvironmentId** | **string** |  | 
**ExternalId** | **string** |  | 
**Id** | **string** |  | 
**InvoiceExternalId** | Pointer to **NullableString** |  | [optional] 
**PaymentMethodType** | **string** |  | 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingPaymentMethodResponseData

`func NewBillingPaymentMethodResponseData(createdAt time.Time, customerExternalId string, environmentId string, externalId string, id string, paymentMethodType string, updatedAt time.Time, ) *BillingPaymentMethodResponseData`

NewBillingPaymentMethodResponseData instantiates a new BillingPaymentMethodResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPaymentMethodResponseDataWithDefaults

`func NewBillingPaymentMethodResponseDataWithDefaults() *BillingPaymentMethodResponseData`

NewBillingPaymentMethodResponseDataWithDefaults instantiates a new BillingPaymentMethodResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBrand

`func (o *BillingPaymentMethodResponseData) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *BillingPaymentMethodResponseData) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *BillingPaymentMethodResponseData) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *BillingPaymentMethodResponseData) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *BillingPaymentMethodResponseData) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *BillingPaymentMethodResponseData) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardExpMonth

`func (o *BillingPaymentMethodResponseData) GetCardExpMonth() int32`

GetCardExpMonth returns the CardExpMonth field if non-nil, zero value otherwise.

### GetCardExpMonthOk

`func (o *BillingPaymentMethodResponseData) GetCardExpMonthOk() (*int32, bool)`

GetCardExpMonthOk returns a tuple with the CardExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpMonth

`func (o *BillingPaymentMethodResponseData) SetCardExpMonth(v int32)`

SetCardExpMonth sets CardExpMonth field to given value.

### HasCardExpMonth

`func (o *BillingPaymentMethodResponseData) HasCardExpMonth() bool`

HasCardExpMonth returns a boolean if a field has been set.

### SetCardExpMonthNil

`func (o *BillingPaymentMethodResponseData) SetCardExpMonthNil(b bool)`

 SetCardExpMonthNil sets the value for CardExpMonth to be an explicit nil

### UnsetCardExpMonth
`func (o *BillingPaymentMethodResponseData) UnsetCardExpMonth()`

UnsetCardExpMonth ensures that no value is present for CardExpMonth, not even an explicit nil
### GetCardExpYear

`func (o *BillingPaymentMethodResponseData) GetCardExpYear() int32`

GetCardExpYear returns the CardExpYear field if non-nil, zero value otherwise.

### GetCardExpYearOk

`func (o *BillingPaymentMethodResponseData) GetCardExpYearOk() (*int32, bool)`

GetCardExpYearOk returns a tuple with the CardExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpYear

`func (o *BillingPaymentMethodResponseData) SetCardExpYear(v int32)`

SetCardExpYear sets CardExpYear field to given value.

### HasCardExpYear

`func (o *BillingPaymentMethodResponseData) HasCardExpYear() bool`

HasCardExpYear returns a boolean if a field has been set.

### SetCardExpYearNil

`func (o *BillingPaymentMethodResponseData) SetCardExpYearNil(b bool)`

 SetCardExpYearNil sets the value for CardExpYear to be an explicit nil

### UnsetCardExpYear
`func (o *BillingPaymentMethodResponseData) UnsetCardExpYear()`

UnsetCardExpYear ensures that no value is present for CardExpYear, not even an explicit nil
### GetCardLast4

`func (o *BillingPaymentMethodResponseData) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *BillingPaymentMethodResponseData) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *BillingPaymentMethodResponseData) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.

### HasCardLast4

`func (o *BillingPaymentMethodResponseData) HasCardLast4() bool`

HasCardLast4 returns a boolean if a field has been set.

### SetCardLast4Nil

`func (o *BillingPaymentMethodResponseData) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *BillingPaymentMethodResponseData) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetCompanyId

`func (o *BillingPaymentMethodResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingPaymentMethodResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingPaymentMethodResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BillingPaymentMethodResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *BillingPaymentMethodResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *BillingPaymentMethodResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *BillingPaymentMethodResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingPaymentMethodResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingPaymentMethodResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCustomerExternalId

`func (o *BillingPaymentMethodResponseData) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *BillingPaymentMethodResponseData) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *BillingPaymentMethodResponseData) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetEnvironmentId

`func (o *BillingPaymentMethodResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BillingPaymentMethodResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BillingPaymentMethodResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetExternalId

`func (o *BillingPaymentMethodResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingPaymentMethodResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingPaymentMethodResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetId

`func (o *BillingPaymentMethodResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPaymentMethodResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPaymentMethodResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceExternalId

`func (o *BillingPaymentMethodResponseData) GetInvoiceExternalId() string`

GetInvoiceExternalId returns the InvoiceExternalId field if non-nil, zero value otherwise.

### GetInvoiceExternalIdOk

`func (o *BillingPaymentMethodResponseData) GetInvoiceExternalIdOk() (*string, bool)`

GetInvoiceExternalIdOk returns a tuple with the InvoiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExternalId

`func (o *BillingPaymentMethodResponseData) SetInvoiceExternalId(v string)`

SetInvoiceExternalId sets InvoiceExternalId field to given value.

### HasInvoiceExternalId

`func (o *BillingPaymentMethodResponseData) HasInvoiceExternalId() bool`

HasInvoiceExternalId returns a boolean if a field has been set.

### SetInvoiceExternalIdNil

`func (o *BillingPaymentMethodResponseData) SetInvoiceExternalIdNil(b bool)`

 SetInvoiceExternalIdNil sets the value for InvoiceExternalId to be an explicit nil

### UnsetInvoiceExternalId
`func (o *BillingPaymentMethodResponseData) UnsetInvoiceExternalId()`

UnsetInvoiceExternalId ensures that no value is present for InvoiceExternalId, not even an explicit nil
### GetPaymentMethodType

`func (o *BillingPaymentMethodResponseData) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *BillingPaymentMethodResponseData) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *BillingPaymentMethodResponseData) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.


### GetSubscriptionExternalId

`func (o *BillingPaymentMethodResponseData) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *BillingPaymentMethodResponseData) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *BillingPaymentMethodResponseData) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *BillingPaymentMethodResponseData) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *BillingPaymentMethodResponseData) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *BillingPaymentMethodResponseData) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
### GetUpdatedAt

`func (o *BillingPaymentMethodResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingPaymentMethodResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingPaymentMethodResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


