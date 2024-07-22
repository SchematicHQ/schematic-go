# CrmLineItemResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**DealId** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | **string** |  | 
**ProductExternalId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCrmLineItemResponseData

`func NewCrmLineItemResponseData(accountId string, createdAt time.Time, environmentId string, updatedAt time.Time, ) *CrmLineItemResponseData`

NewCrmLineItemResponseData instantiates a new CrmLineItemResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmLineItemResponseDataWithDefaults

`func NewCrmLineItemResponseDataWithDefaults() *CrmLineItemResponseData`

NewCrmLineItemResponseDataWithDefaults instantiates a new CrmLineItemResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CrmLineItemResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CrmLineItemResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CrmLineItemResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *CrmLineItemResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CrmLineItemResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CrmLineItemResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDealId

`func (o *CrmLineItemResponseData) GetDealId() string`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *CrmLineItemResponseData) GetDealIdOk() (*string, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *CrmLineItemResponseData) SetDealId(v string)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *CrmLineItemResponseData) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### SetDealIdNil

`func (o *CrmLineItemResponseData) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *CrmLineItemResponseData) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetEnvironmentId

`func (o *CrmLineItemResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CrmLineItemResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CrmLineItemResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetProductExternalId

`func (o *CrmLineItemResponseData) GetProductExternalId() string`

GetProductExternalId returns the ProductExternalId field if non-nil, zero value otherwise.

### GetProductExternalIdOk

`func (o *CrmLineItemResponseData) GetProductExternalIdOk() (*string, bool)`

GetProductExternalIdOk returns a tuple with the ProductExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalId

`func (o *CrmLineItemResponseData) SetProductExternalId(v string)`

SetProductExternalId sets ProductExternalId field to given value.

### HasProductExternalId

`func (o *CrmLineItemResponseData) HasProductExternalId() bool`

HasProductExternalId returns a boolean if a field has been set.

### SetProductExternalIdNil

`func (o *CrmLineItemResponseData) SetProductExternalIdNil(b bool)`

 SetProductExternalIdNil sets the value for ProductExternalId to be an explicit nil

### UnsetProductExternalId
`func (o *CrmLineItemResponseData) UnsetProductExternalId()`

UnsetProductExternalId ensures that no value is present for ProductExternalId, not even an explicit nil
### GetUpdatedAt

`func (o *CrmLineItemResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CrmLineItemResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CrmLineItemResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


