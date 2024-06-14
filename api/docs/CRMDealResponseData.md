# CrmDealResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Arr** | **string** |  | 
**CompanyExternalId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**DealExternalId** | **string** |  | 
**DealId** | **string** |  | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | **string** |  | 
**Mrr** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProductExternalId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCrmDealResponseData

`func NewCrmDealResponseData(accountId string, arr string, createdAt time.Time, dealExternalId string, dealId string, environmentId string, mrr string, updatedAt time.Time, ) *CrmDealResponseData`

NewCrmDealResponseData instantiates a new CrmDealResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmDealResponseDataWithDefaults

`func NewCrmDealResponseDataWithDefaults() *CrmDealResponseData`

NewCrmDealResponseDataWithDefaults instantiates a new CrmDealResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CrmDealResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CrmDealResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CrmDealResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetArr

`func (o *CrmDealResponseData) GetArr() string`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CrmDealResponseData) GetArrOk() (*string, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CrmDealResponseData) SetArr(v string)`

SetArr sets Arr field to given value.


### GetCompanyExternalId

`func (o *CrmDealResponseData) GetCompanyExternalId() string`

GetCompanyExternalId returns the CompanyExternalId field if non-nil, zero value otherwise.

### GetCompanyExternalIdOk

`func (o *CrmDealResponseData) GetCompanyExternalIdOk() (*string, bool)`

GetCompanyExternalIdOk returns a tuple with the CompanyExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyExternalId

`func (o *CrmDealResponseData) SetCompanyExternalId(v string)`

SetCompanyExternalId sets CompanyExternalId field to given value.

### HasCompanyExternalId

`func (o *CrmDealResponseData) HasCompanyExternalId() bool`

HasCompanyExternalId returns a boolean if a field has been set.

### SetCompanyExternalIdNil

`func (o *CrmDealResponseData) SetCompanyExternalIdNil(b bool)`

 SetCompanyExternalIdNil sets the value for CompanyExternalId to be an explicit nil

### UnsetCompanyExternalId
`func (o *CrmDealResponseData) UnsetCompanyExternalId()`

UnsetCompanyExternalId ensures that no value is present for CompanyExternalId, not even an explicit nil
### GetCreatedAt

`func (o *CrmDealResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmDealResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmDealResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDealExternalId

`func (o *CrmDealResponseData) GetDealExternalId() string`

GetDealExternalId returns the DealExternalId field if non-nil, zero value otherwise.

### GetDealExternalIdOk

`func (o *CrmDealResponseData) GetDealExternalIdOk() (*string, bool)`

GetDealExternalIdOk returns a tuple with the DealExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExternalId

`func (o *CrmDealResponseData) SetDealExternalId(v string)`

SetDealExternalId sets DealExternalId field to given value.


### GetDealId

`func (o *CrmDealResponseData) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *CrmDealResponseData) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *CrmDealResponseData) SetDealId(v string)`

SetDealId sets DealId field to given value.


### GetDeletedAt

`func (o *CrmDealResponseData) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CrmDealResponseData) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CrmDealResponseData) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CrmDealResponseData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *CrmDealResponseData) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CrmDealResponseData) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEnvironmentId

`func (o *CrmDealResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CrmDealResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CrmDealResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetMrr

`func (o *CrmDealResponseData) GetMrr() string`

GetMrr returns the Mrr field if non-nil, zero value otherwise.

### GetMrrOk

`func (o *CrmDealResponseData) GetMrrOk() (*string, bool)`

GetMrrOk returns a tuple with the Mrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrr

`func (o *CrmDealResponseData) SetMrr(v string)`

SetMrr sets Mrr field to given value.


### GetName

`func (o *CrmDealResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CrmDealResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CrmDealResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CrmDealResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CrmDealResponseData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CrmDealResponseData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProductExternalId

`func (o *CrmDealResponseData) GetProductExternalId() string`

GetProductExternalId returns the ProductExternalId field if non-nil, zero value otherwise.

### GetProductExternalIdOk

`func (o *CrmDealResponseData) GetProductExternalIdOk() (*string, bool)`

GetProductExternalIdOk returns a tuple with the ProductExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalId

`func (o *CrmDealResponseData) SetProductExternalId(v string)`

SetProductExternalId sets ProductExternalId field to given value.

### HasProductExternalId

`func (o *CrmDealResponseData) HasProductExternalId() bool`

HasProductExternalId returns a boolean if a field has been set.

### SetProductExternalIdNil

`func (o *CrmDealResponseData) SetProductExternalIdNil(b bool)`

 SetProductExternalIdNil sets the value for ProductExternalId to be an explicit nil

### UnsetProductExternalId
`func (o *CrmDealResponseData) UnsetProductExternalId()`

UnsetProductExternalId ensures that no value is present for ProductExternalId, not even an explicit nil
### GetUpdatedAt

`func (o *CrmDealResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmDealResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmDealResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


