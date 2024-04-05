# UpsertUserRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | **map[string]interface{}** | Optionally specify company using key/value pairs | 
**CompanyId** | Pointer to **NullableString** | Optionally specify company using Schematic company ID | [optional] 
**Keys** | **map[string]interface{}** |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdateOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpsertUserRequestBody

`func NewUpsertUserRequestBody(company map[string]interface{}, keys map[string]interface{}, ) *UpsertUserRequestBody`

NewUpsertUserRequestBody instantiates a new UpsertUserRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertUserRequestBodyWithDefaults

`func NewUpsertUserRequestBodyWithDefaults() *UpsertUserRequestBody`

NewUpsertUserRequestBodyWithDefaults instantiates a new UpsertUserRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *UpsertUserRequestBody) GetCompany() map[string]interface{}`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpsertUserRequestBody) GetCompanyOk() (*map[string]interface{}, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpsertUserRequestBody) SetCompany(v map[string]interface{})`

SetCompany sets Company field to given value.


### GetCompanyId

`func (o *UpsertUserRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UpsertUserRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UpsertUserRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UpsertUserRequestBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *UpsertUserRequestBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *UpsertUserRequestBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetKeys

`func (o *UpsertUserRequestBody) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UpsertUserRequestBody) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UpsertUserRequestBody) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetLastSeenAt

`func (o *UpsertUserRequestBody) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *UpsertUserRequestBody) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *UpsertUserRequestBody) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *UpsertUserRequestBody) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *UpsertUserRequestBody) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *UpsertUserRequestBody) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *UpsertUserRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertUserRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertUserRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpsertUserRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpsertUserRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpsertUserRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTraits

`func (o *UpsertUserRequestBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *UpsertUserRequestBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *UpsertUserRequestBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *UpsertUserRequestBody) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUpdateOnly

`func (o *UpsertUserRequestBody) GetUpdateOnly() bool`

GetUpdateOnly returns the UpdateOnly field if non-nil, zero value otherwise.

### GetUpdateOnlyOk

`func (o *UpsertUserRequestBody) GetUpdateOnlyOk() (*bool, bool)`

GetUpdateOnlyOk returns a tuple with the UpdateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnly

`func (o *UpsertUserRequestBody) SetUpdateOnly(v bool)`

SetUpdateOnly sets UpdateOnly field to given value.

### HasUpdateOnly

`func (o *UpsertUserRequestBody) HasUpdateOnly() bool`

HasUpdateOnly returns a boolean if a field has been set.

### SetUpdateOnlyNil

`func (o *UpsertUserRequestBody) SetUpdateOnlyNil(b bool)`

 SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil

### UnsetUpdateOnly
`func (o *UpsertUserRequestBody) UnsetUpdateOnly()`

UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


