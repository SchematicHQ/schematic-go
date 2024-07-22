# BillingCustomerWithSubscriptionsResponseData

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
**Subscriptions** | [**[]BillingCustomerSubscription**](BillingCustomerSubscription.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingCustomerWithSubscriptionsResponseData

`func NewBillingCustomerWithSubscriptionsResponseData(email string, externalId string, failedToImport bool, id string, name string, subscriptions []BillingCustomerSubscription, updatedAt time.Time, ) *BillingCustomerWithSubscriptionsResponseData`

NewBillingCustomerWithSubscriptionsResponseData instantiates a new BillingCustomerWithSubscriptionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCustomerWithSubscriptionsResponseDataWithDefaults

`func NewBillingCustomerWithSubscriptionsResponseDataWithDefaults() *BillingCustomerWithSubscriptionsResponseData`

NewBillingCustomerWithSubscriptionsResponseDataWithDefaults instantiates a new BillingCustomerWithSubscriptionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BillingCustomerWithSubscriptionsResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingCustomerWithSubscriptionsResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BillingCustomerWithSubscriptionsResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *BillingCustomerWithSubscriptionsResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *BillingCustomerWithSubscriptionsResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetDeletedAt

`func (o *BillingCustomerWithSubscriptionsResponseData) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BillingCustomerWithSubscriptionsResponseData) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BillingCustomerWithSubscriptionsResponseData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BillingCustomerWithSubscriptionsResponseData) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BillingCustomerWithSubscriptionsResponseData) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEmail

`func (o *BillingCustomerWithSubscriptionsResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingCustomerWithSubscriptionsResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalId

`func (o *BillingCustomerWithSubscriptionsResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingCustomerWithSubscriptionsResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetFailedToImport

`func (o *BillingCustomerWithSubscriptionsResponseData) GetFailedToImport() bool`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetFailedToImportOk() (*bool, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *BillingCustomerWithSubscriptionsResponseData) SetFailedToImport(v bool)`

SetFailedToImport sets FailedToImport field to given value.


### GetId

`func (o *BillingCustomerWithSubscriptionsResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCustomerWithSubscriptionsResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BillingCustomerWithSubscriptionsResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingCustomerWithSubscriptionsResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetSubscriptions

`func (o *BillingCustomerWithSubscriptionsResponseData) GetSubscriptions() []BillingCustomerSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetSubscriptionsOk() (*[]BillingCustomerSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *BillingCustomerWithSubscriptionsResponseData) SetSubscriptions(v []BillingCustomerSubscription)`

SetSubscriptions sets Subscriptions field to given value.


### GetUpdatedAt

`func (o *BillingCustomerWithSubscriptionsResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingCustomerWithSubscriptionsResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingCustomerWithSubscriptionsResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


