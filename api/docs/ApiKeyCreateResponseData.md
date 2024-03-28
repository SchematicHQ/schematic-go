# ApiKeyCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**LastUsedAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**Scopes** | **[]string** |  | 
**Secret** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewApiKeyCreateResponseData

`func NewApiKeyCreateResponseData(createdAt time.Time, id string, name string, scopes []string, secret string, updatedAt time.Time, ) *ApiKeyCreateResponseData`

NewApiKeyCreateResponseData instantiates a new ApiKeyCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreateResponseDataWithDefaults

`func NewApiKeyCreateResponseDataWithDefaults() *ApiKeyCreateResponseData`

NewApiKeyCreateResponseDataWithDefaults instantiates a new ApiKeyCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiKeyCreateResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKeyCreateResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKeyCreateResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *ApiKeyCreateResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyCreateResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyCreateResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyCreateResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiKeyCreateResponseData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiKeyCreateResponseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvironmentId

`func (o *ApiKeyCreateResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApiKeyCreateResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApiKeyCreateResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApiKeyCreateResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *ApiKeyCreateResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *ApiKeyCreateResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetId

`func (o *ApiKeyCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLastUsedAt

`func (o *ApiKeyCreateResponseData) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ApiKeyCreateResponseData) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ApiKeyCreateResponseData) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ApiKeyCreateResponseData) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ApiKeyCreateResponseData) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ApiKeyCreateResponseData) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetName

`func (o *ApiKeyCreateResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyCreateResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyCreateResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetScopes

`func (o *ApiKeyCreateResponseData) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApiKeyCreateResponseData) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApiKeyCreateResponseData) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetSecret

`func (o *ApiKeyCreateResponseData) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApiKeyCreateResponseData) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApiKeyCreateResponseData) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetUpdatedAt

`func (o *ApiKeyCreateResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiKeyCreateResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiKeyCreateResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


