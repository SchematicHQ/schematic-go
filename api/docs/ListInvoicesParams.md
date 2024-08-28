# ListInvoicesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyID** | Pointer to **string** |  | [optional] 
**CustomerExternalID** | Pointer to **string** |  | [optional] 
**SubscriptionExternalID** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 

## Methods

### NewListInvoicesParams

`func NewListInvoicesParams() *ListInvoicesParams`

NewListInvoicesParams instantiates a new ListInvoicesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesParamsWithDefaults

`func NewListInvoicesParamsWithDefaults() *ListInvoicesParams`

NewListInvoicesParamsWithDefaults instantiates a new ListInvoicesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyID

`func (o *ListInvoicesParams) GetCompanyID() string`

GetCompanyID returns the CompanyID field if non-nil, zero value otherwise.

### GetCompanyIDOk

`func (o *ListInvoicesParams) GetCompanyIDOk() (*string, bool)`

GetCompanyIDOk returns a tuple with the CompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyID

`func (o *ListInvoicesParams) SetCompanyID(v string)`

SetCompanyID sets CompanyID field to given value.

### HasCompanyID

`func (o *ListInvoicesParams) HasCompanyID() bool`

HasCompanyID returns a boolean if a field has been set.

### GetCustomerExternalID

`func (o *ListInvoicesParams) GetCustomerExternalID() string`

GetCustomerExternalID returns the CustomerExternalID field if non-nil, zero value otherwise.

### GetCustomerExternalIDOk

`func (o *ListInvoicesParams) GetCustomerExternalIDOk() (*string, bool)`

GetCustomerExternalIDOk returns a tuple with the CustomerExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalID

`func (o *ListInvoicesParams) SetCustomerExternalID(v string)`

SetCustomerExternalID sets CustomerExternalID field to given value.

### HasCustomerExternalID

`func (o *ListInvoicesParams) HasCustomerExternalID() bool`

HasCustomerExternalID returns a boolean if a field has been set.

### GetSubscriptionExternalID

`func (o *ListInvoicesParams) GetSubscriptionExternalID() string`

GetSubscriptionExternalID returns the SubscriptionExternalID field if non-nil, zero value otherwise.

### GetSubscriptionExternalIDOk

`func (o *ListInvoicesParams) GetSubscriptionExternalIDOk() (*string, bool)`

GetSubscriptionExternalIDOk returns a tuple with the SubscriptionExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalID

`func (o *ListInvoicesParams) SetSubscriptionExternalID(v string)`

SetSubscriptionExternalID sets SubscriptionExternalID field to given value.

### HasSubscriptionExternalID

`func (o *ListInvoicesParams) HasSubscriptionExternalID() bool`

HasSubscriptionExternalID returns a boolean if a field has been set.

### GetLimit

`func (o *ListInvoicesParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListInvoicesParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListInvoicesParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListInvoicesParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListInvoicesParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListInvoicesParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListInvoicesParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListInvoicesParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


