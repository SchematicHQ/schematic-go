# ListPaymentMethodsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CustomerExternalId** | Pointer to **string** |  | [optional] 
**InvoiceExternalId** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**SubscriptionExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewListPaymentMethodsParams

`func NewListPaymentMethodsParams() *ListPaymentMethodsParams`

NewListPaymentMethodsParams instantiates a new ListPaymentMethodsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentMethodsParamsWithDefaults

`func NewListPaymentMethodsParamsWithDefaults() *ListPaymentMethodsParams`

NewListPaymentMethodsParamsWithDefaults instantiates a new ListPaymentMethodsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListPaymentMethodsParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListPaymentMethodsParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListPaymentMethodsParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListPaymentMethodsParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCustomerExternalId

`func (o *ListPaymentMethodsParams) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *ListPaymentMethodsParams) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *ListPaymentMethodsParams) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.

### HasCustomerExternalId

`func (o *ListPaymentMethodsParams) HasCustomerExternalId() bool`

HasCustomerExternalId returns a boolean if a field has been set.

### GetInvoiceExternalId

`func (o *ListPaymentMethodsParams) GetInvoiceExternalId() string`

GetInvoiceExternalId returns the InvoiceExternalId field if non-nil, zero value otherwise.

### GetInvoiceExternalIdOk

`func (o *ListPaymentMethodsParams) GetInvoiceExternalIdOk() (*string, bool)`

GetInvoiceExternalIdOk returns a tuple with the InvoiceExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExternalId

`func (o *ListPaymentMethodsParams) SetInvoiceExternalId(v string)`

SetInvoiceExternalId sets InvoiceExternalId field to given value.

### HasInvoiceExternalId

`func (o *ListPaymentMethodsParams) HasInvoiceExternalId() bool`

HasInvoiceExternalId returns a boolean if a field has been set.

### GetLimit

`func (o *ListPaymentMethodsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPaymentMethodsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPaymentMethodsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPaymentMethodsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListPaymentMethodsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPaymentMethodsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPaymentMethodsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPaymentMethodsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSubscriptionExternalId

`func (o *ListPaymentMethodsParams) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *ListPaymentMethodsParams) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *ListPaymentMethodsParams) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.

### HasSubscriptionExternalId

`func (o *ListPaymentMethodsParams) HasSubscriptionExternalId() bool`

HasSubscriptionExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


