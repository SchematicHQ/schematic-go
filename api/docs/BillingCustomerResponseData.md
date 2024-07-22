# BillingCustomerResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**Email** | **string** |  | 
**ExternalId** | **string** |  | 
**FailedToImport** | **bool** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingCustomerResponseData

`func NewBillingCustomerResponseData(email string, externalId string, failedToImport bool, id string, name string, updatedAt time.Time, ) *BillingCustomerResponseData`

NewBillingCustomerResponseData instantiates a new BillingCustomerResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCustomerResponseDataWithDefaults

`func NewBillingCustomerResponseDataWithDefaults() *BillingCustomerResponseData`

NewBillingCustomerResponseDataWithDefaults instantiates a new BillingCustomerResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BillingCustomerResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingCustomerResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingCustomerResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BillingCustomerResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *BillingCustomerResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *BillingCustomerResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetDeletedAt

`func (o *BillingCustomerResponseData) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BillingCustomerResponseData) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BillingCustomerResponseData) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BillingCustomerResponseData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BillingCustomerResponseData) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BillingCustomerResponseData) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEmail

`func (o *BillingCustomerResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingCustomerResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingCustomerResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *BillingCustomerResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingCustomerResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingCustomerResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFailedToImport

`func (o *BillingCustomerResponseData) GetFailedToImport() bool`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *BillingCustomerResponseData) GetFailedToImportOk() (*bool, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *BillingCustomerResponseData) SetFailedToImport(v bool)`

SetFailedToImport sets FailedToImport field to given value.


### GetId

`func (o *BillingCustomerResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCustomerResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCustomerResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BillingCustomerResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingCustomerResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingCustomerResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *BillingCustomerResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingCustomerResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingCustomerResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


