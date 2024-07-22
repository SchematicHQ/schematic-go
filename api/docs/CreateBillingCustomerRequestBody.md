# CreateBillingCustomerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**ExternalId** | **string** |  | 
**FailedToImport** | **bool** |  | 
**Meta** | **map[string]string** |  | 
**Name** | **string** |  | 

## Methods

### NewCreateBillingCustomerRequestBody

`func NewCreateBillingCustomerRequestBody(email string, externalId string, failedToImport bool, meta map[string]string, name string, ) *CreateBillingCustomerRequestBody`

NewCreateBillingCustomerRequestBody instantiates a new CreateBillingCustomerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingCustomerRequestBodyWithDefaults

`func NewCreateBillingCustomerRequestBodyWithDefaults() *CreateBillingCustomerRequestBody`

NewCreateBillingCustomerRequestBodyWithDefaults instantiates a new CreateBillingCustomerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CreateBillingCustomerRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CreateBillingCustomerRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CreateBillingCustomerRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CreateBillingCustomerRequestBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CreateBillingCustomerRequestBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CreateBillingCustomerRequestBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetEmail

`func (o *CreateBillingCustomerRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateBillingCustomerRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateBillingCustomerRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *CreateBillingCustomerRequestBody) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateBillingCustomerRequestBody) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateBillingCustomerRequestBody) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFailedToImport

`func (o *CreateBillingCustomerRequestBody) GetFailedToImport() bool`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *CreateBillingCustomerRequestBody) GetFailedToImportOk() (*bool, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *CreateBillingCustomerRequestBody) SetFailedToImport(v bool)`

SetFailedToImport sets FailedToImport field to given value.


### GetMeta

`func (o *CreateBillingCustomerRequestBody) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreateBillingCustomerRequestBody) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreateBillingCustomerRequestBody) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.


### GetName

`func (o *CreateBillingCustomerRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBillingCustomerRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBillingCustomerRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


