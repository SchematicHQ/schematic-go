# CRMDealResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CompanyExternalId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**DealExternalId** | **string** |  | 
**DealId** | **string** |  | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProductExternalId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCRMDealResponseData

`func NewCRMDealResponseData(accountId string, createdAt time.Time, dealExternalId string, dealId string, environmentId string, updatedAt time.Time, ) *CRMDealResponseData`

NewCRMDealResponseData instantiates a new CRMDealResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCRMDealResponseDataWithDefaults

`func NewCRMDealResponseDataWithDefaults() *CRMDealResponseData`

NewCRMDealResponseDataWithDefaults instantiates a new CRMDealResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CRMDealResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CRMDealResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CRMDealResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCompanyExternalId

`func (o *CRMDealResponseData) GetCompanyExternalId() string`

GetCompanyExternalId returns the CompanyExternalId field if non-nil, zero value otherwise.

### GetCompanyExternalIdOk

`func (o *CRMDealResponseData) GetCompanyExternalIdOk() (*string, bool)`

GetCompanyExternalIdOk returns a tuple with the CompanyExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyExternalId

`func (o *CRMDealResponseData) SetCompanyExternalId(v string)`

SetCompanyExternalId sets CompanyExternalId field to given value.

### HasCompanyExternalId

`func (o *CRMDealResponseData) HasCompanyExternalId() bool`

HasCompanyExternalId returns a boolean if a field has been set.

### SetCompanyExternalIdNil

`func (o *CRMDealResponseData) SetCompanyExternalIdNil(b bool)`

 SetCompanyExternalIdNil sets the value for CompanyExternalId to be an explicit nil

### UnsetCompanyExternalId
`func (o *CRMDealResponseData) UnsetCompanyExternalId()`

UnsetCompanyExternalId ensures that no value is present for CompanyExternalId, not even an explicit nil
### GetCreatedAt

`func (o *CRMDealResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CRMDealResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CRMDealResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDealExternalId

`func (o *CRMDealResponseData) GetDealExternalId() string`

GetDealExternalId returns the DealExternalId field if non-nil, zero value otherwise.

### GetDealExternalIdOk

`func (o *CRMDealResponseData) GetDealExternalIdOk() (*string, bool)`

GetDealExternalIdOk returns a tuple with the DealExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExternalId

`func (o *CRMDealResponseData) SetDealExternalId(v string)`

SetDealExternalId sets DealExternalId field to given value.


### GetDealId

`func (o *CRMDealResponseData) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *CRMDealResponseData) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *CRMDealResponseData) SetDealId(v string)`

SetDealId sets DealId field to given value.


### GetDeletedAt

`func (o *CRMDealResponseData) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CRMDealResponseData) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CRMDealResponseData) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CRMDealResponseData) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *CRMDealResponseData) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CRMDealResponseData) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetEnvironmentId

`func (o *CRMDealResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CRMDealResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CRMDealResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *CRMDealResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CRMDealResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CRMDealResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CRMDealResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CRMDealResponseData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CRMDealResponseData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProductExternalId

`func (o *CRMDealResponseData) GetProductExternalId() string`

GetProductExternalId returns the ProductExternalId field if non-nil, zero value otherwise.

### GetProductExternalIdOk

`func (o *CRMDealResponseData) GetProductExternalIdOk() (*string, bool)`

GetProductExternalIdOk returns a tuple with the ProductExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalId

`func (o *CRMDealResponseData) SetProductExternalId(v string)`

SetProductExternalId sets ProductExternalId field to given value.

### HasProductExternalId

`func (o *CRMDealResponseData) HasProductExternalId() bool`

HasProductExternalId returns a boolean if a field has been set.

### SetProductExternalIdNil

`func (o *CRMDealResponseData) SetProductExternalIdNil(b bool)`

 SetProductExternalIdNil sets the value for ProductExternalId to be an explicit nil

### UnsetProductExternalId
`func (o *CRMDealResponseData) UnsetProductExternalId()`

UnsetProductExternalId ensures that no value is present for ProductExternalId, not even an explicit nil
### GetUpdatedAt

`func (o *CRMDealResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CRMDealResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CRMDealResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


