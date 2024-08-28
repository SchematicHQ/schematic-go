# ListPaymentMethodsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyID** | Pointer to **string** |  | [optional] 
**CustomerExternalID** | Pointer to **string** |  | [optional] 
**InvoiceExternalID** | Pointer to **string** |  | [optional] 
**SubscriptionExternalID** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 

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

### GetCompanyID

`func (o *ListPaymentMethodsParams) GetCompanyID() string`

GetCompanyID returns the CompanyID field if non-nil, zero value otherwise.

### GetCompanyIDOk

`func (o *ListPaymentMethodsParams) GetCompanyIDOk() (*string, bool)`

GetCompanyIDOk returns a tuple with the CompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyID

`func (o *ListPaymentMethodsParams) SetCompanyID(v string)`

SetCompanyID sets CompanyID field to given value.

### HasCompanyID

`func (o *ListPaymentMethodsParams) HasCompanyID() bool`

HasCompanyID returns a boolean if a field has been set.

### GetCustomerExternalID

`func (o *ListPaymentMethodsParams) GetCustomerExternalID() string`

GetCustomerExternalID returns the CustomerExternalID field if non-nil, zero value otherwise.

### GetCustomerExternalIDOk

`func (o *ListPaymentMethodsParams) GetCustomerExternalIDOk() (*string, bool)`

GetCustomerExternalIDOk returns a tuple with the CustomerExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalID

`func (o *ListPaymentMethodsParams) SetCustomerExternalID(v string)`

SetCustomerExternalID sets CustomerExternalID field to given value.

### HasCustomerExternalID

`func (o *ListPaymentMethodsParams) HasCustomerExternalID() bool`

HasCustomerExternalID returns a boolean if a field has been set.

### GetInvoiceExternalID

`func (o *ListPaymentMethodsParams) GetInvoiceExternalID() string`

GetInvoiceExternalID returns the InvoiceExternalID field if non-nil, zero value otherwise.

### GetInvoiceExternalIDOk

`func (o *ListPaymentMethodsParams) GetInvoiceExternalIDOk() (*string, bool)`

GetInvoiceExternalIDOk returns a tuple with the InvoiceExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceExternalID

`func (o *ListPaymentMethodsParams) SetInvoiceExternalID(v string)`

SetInvoiceExternalID sets InvoiceExternalID field to given value.

### HasInvoiceExternalID

`func (o *ListPaymentMethodsParams) HasInvoiceExternalID() bool`

HasInvoiceExternalID returns a boolean if a field has been set.

### GetSubscriptionExternalID

`func (o *ListPaymentMethodsParams) GetSubscriptionExternalID() string`

GetSubscriptionExternalID returns the SubscriptionExternalID field if non-nil, zero value otherwise.

### GetSubscriptionExternalIDOk

`func (o *ListPaymentMethodsParams) GetSubscriptionExternalIDOk() (*string, bool)`

GetSubscriptionExternalIDOk returns a tuple with the SubscriptionExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalID

`func (o *ListPaymentMethodsParams) SetSubscriptionExternalID(v string)`

SetSubscriptionExternalID sets SubscriptionExternalID field to given value.

### HasSubscriptionExternalID

`func (o *ListPaymentMethodsParams) HasSubscriptionExternalID() bool`

HasSubscriptionExternalID returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


