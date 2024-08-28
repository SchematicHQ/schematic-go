# CreateBillingPaymentMethodRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardBrand** | Pointer to **NullableString** |  | [optional] 
**CardExpMonth** | Pointer to **NullableInt32** |  | [optional] 
**CardExpYear** | Pointer to **NullableInt32** |  | [optional] 
**CardLast4** | Pointer to **NullableString** |  | [optional] 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CustomerExternalId** | **string** |  | 
**ExternalId** | **string** |  | 
**InvoiceExternalId** | Pointer to **NullableString** |  | [optional] 
**PaymentMethodType** | **string** |  | 
**SubscriptionExternalId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateBillingPaymentMethodRequestBody

`func NewCreateBillingPaymentMethodRequestBody(customerExternalId string, externalId string, paymentMethodType string, ) *CreateBillingPaymentMethodRequestBody`

NewCreateBillingPaymentMethodRequestBody instantiates a new CreateBillingPaymentMethodRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingPaymentMethodRequestBodyWithDefaults

`func NewCreateBillingPaymentMethodRequestBodyWithDefaults() *CreateBillingPaymentMethodRequestBody`

NewCreateBillingPaymentMethodRequestBodyWithDefaults instantiates a new CreateBillingPaymentMethodRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardBrand

`func (o *CreateBillingPaymentMethodRequestBody) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CreateBillingPaymentMethodRequestBody) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CreateBillingPaymentMethodRequestBody) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *CreateBillingPaymentMethodRequestBody) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *CreateBillingPaymentMethodRequestBody) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *CreateBillingPaymentMethodRequestBody) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardExpMonth

`func (o *CreateBillingPaymentMethodRequestBody) GetCardExpMonth() int32`

GetCardExpMonth returns the CardExpMonth field if non-nil, zero value otherwise.

### GetCardExpMonthOk

`func (o *CreateBillingPaymentMethodRequestBody) GetCardExpMonthOk() (*int32, bool)`

GetCardExpMonthOk returns a tuple with the CardExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpMonth

`func (o *CreateBillingPaymentMethodRequestBody) SetCardExpMonth(v int32)`

SetCardExpMonth sets CardExpMonth field to given value.

### HasCardExpMonth

`func (o *CreateBillingPaymentMethodRequestBody) HasCardExpMonth() bool`

HasCardExpMonth returns a boolean if a field has been set.

### SetCardExpMonthNil

`func (o *CreateBillingPaymentMethodRequestBody) SetCardExpMonthNil(b bool)`

 SetCardExpMonthNil sets the value for CardExpMonth to be an explicit nil

### UnsetCardExpMonth
`func (o *CreateBillingPaymentMethodRequestBody) UnsetCardExpMonth()`

UnsetCardExpMonth ensures that no value is present for CardExpMonth, not even an explicit nil
### GetCardExpYear

`func (o *CreateBillingPaymentMethodRequestBody) GetCardExpYear() int32`

GetCardExpYear returns the CardExpYear field if non-nil, zero value otherwise.

### GetCardExpYearOk

`func (o *CreateBillingPaymentMethodRequestBody) GetCardExpYearOk() (*int32, bool)`

GetCardExpYearOk returns a tuple with the CardExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardExpYear

`func (o *CreateBillingPaymentMethodRequestBody) SetCardExpYear(v int32)`

SetCardExpYear sets CardExpYear field to given value.

### HasCardExpYear

`func (o *CreateBillingPaymentMethodRequestBody) HasCardExpYear() bool`

HasCardExpYear returns a boolean if a field has been set.

### SetCardExpYearNil

`func (o *CreateBillingPaymentMethodRequestBody) SetCardExpYearNil(b bool)`

 SetCardExpYearNil sets the value for CardExpYear to be an explicit nil

### UnsetCardExpYear
`func (o *CreateBillingPaymentMethodRequestBody) UnsetCardExpYear()`

UnsetCardExpYear ensures that no value is present for CardExpYear, not even an explicit nil
### GetCardLast4

`func (o *CreateBillingPaymentMethodRequestBody) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *CreateBillingPaymentMethodRequestBody) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *CreateBillingPaymentMethodRequestBody) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.

### HasCardLast4

`func (o *CreateBillingPaymentMethodRequestBody) HasCardLast4() bool`

HasCardLast4 returns a boolean if a field has been set.

### SetCardLast4Nil

`func (o *CreateBillingPaymentMethodRequestBody) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *CreateBillingPaymentMethodRequestBody) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetCompanyId

`func (o *CreateBillingPaymentMethodRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateBillingPaymentMethodRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateBillingPaymentMethodRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CreateBillingPaymentMethodRequestBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CreateBillingPaymentMethodRequestBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CreateBillingPaymentMethodRequestBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCustomerExternalId

`func (o *CreateBillingPaymentMethodRequestBody) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *CreateBillingPaymentMethodRequestBody) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *CreateBillingPaymentMethodRequestBody) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetExternalId

`func (o *CreateBillingPaymentMethodRequestBody) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateBillingPaymentMethodRequestBody) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateBillingPaymentMethodRequestBody) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetInvoiceExternalId

`func (o *CreateBillingPaymentMethodRequestBody) GetInvoiceExternalId() string`

GetInvoiceExternalId returns the InvoiceExternalId field if non-nil, zero value otherwise.

### GetInvoiceExternalIdOk

`func (o *CreateBillingPaymentMethodRequestBody) GetInvoiceExternalIdOk() (*string, bool)`

GetInvoiceExternalIdOk returns a tuple with the InvoiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExternalId

`func (o *CreateBillingPaymentMethodRequestBody) SetInvoiceExternalId(v string)`

SetInvoiceExternalId sets InvoiceExternalId field to given value.

### HasInvoiceExternalId

`func (o *CreateBillingPaymentMethodRequestBody) HasInvoiceExternalId() bool`

HasInvoiceExternalId returns a boolean if a field has been set.

### SetInvoiceExternalIdNil

`func (o *CreateBillingPaymentMethodRequestBody) SetInvoiceExternalIdNil(b bool)`

 SetInvoiceExternalIdNil sets the value for InvoiceExternalId to be an explicit nil

### UnsetInvoiceExternalId
`func (o *CreateBillingPaymentMethodRequestBody) UnsetInvoiceExternalId()`

UnsetInvoiceExternalId ensures that no value is present for InvoiceExternalId, not even an explicit nil
### GetPaymentMethodType

`func (o *CreateBillingPaymentMethodRequestBody) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *CreateBillingPaymentMethodRequestBody) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *CreateBillingPaymentMethodRequestBody) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.


### GetSubscriptionExternalId

`func (o *CreateBillingPaymentMethodRequestBody) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *CreateBillingPaymentMethodRequestBody) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *CreateBillingPaymentMethodRequestBody) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *CreateBillingPaymentMethodRequestBody) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.

### SetSubscriptionExternalIdNil

`func (o *CreateBillingPaymentMethodRequestBody) SetSubscriptionExternalIdNil(b bool)`

 SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil

### UnsetSubscriptionExternalId
`func (o *CreateBillingPaymentMethodRequestBody) UnsetSubscriptionExternalId()`

UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


