# UpsertUserSubRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Optionally specify company using Schematic company ID | [optional] 
**Id** | Pointer to **NullableString** | If you know the Schematic ID, you can use that here instead of keys | [optional] 
**Keys** | **map[string]string** |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdateOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpsertUserSubRequestBody

`func NewUpsertUserSubRequestBody(keys map[string]string, ) *UpsertUserSubRequestBody`

NewUpsertUserSubRequestBody instantiates a new UpsertUserSubRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertUserSubRequestBodyWithDefaults

`func NewUpsertUserSubRequestBodyWithDefaults() *UpsertUserSubRequestBody`

NewUpsertUserSubRequestBodyWithDefaults instantiates a new UpsertUserSubRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *UpsertUserSubRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UpsertUserSubRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UpsertUserSubRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UpsertUserSubRequestBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *UpsertUserSubRequestBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *UpsertUserSubRequestBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetId

`func (o *UpsertUserSubRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpsertUserSubRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpsertUserSubRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpsertUserSubRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpsertUserSubRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpsertUserSubRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKeys

`func (o *UpsertUserSubRequestBody) GetKeys() map[string]string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UpsertUserSubRequestBody) GetKeysOk() (*map[string]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UpsertUserSubRequestBody) SetKeys(v map[string]string)`

SetKeys sets Keys field to given value.


### GetLastSeenAt

`func (o *UpsertUserSubRequestBody) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *UpsertUserSubRequestBody) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *UpsertUserSubRequestBody) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *UpsertUserSubRequestBody) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *UpsertUserSubRequestBody) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *UpsertUserSubRequestBody) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *UpsertUserSubRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertUserSubRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertUserSubRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpsertUserSubRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpsertUserSubRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpsertUserSubRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTraits

`func (o *UpsertUserSubRequestBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *UpsertUserSubRequestBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *UpsertUserSubRequestBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *UpsertUserSubRequestBody) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUpdateOnly

`func (o *UpsertUserSubRequestBody) GetUpdateOnly() bool`

GetUpdateOnly returns the UpdateOnly field if non-nil, zero value otherwise.

### GetUpdateOnlyOk

`func (o *UpsertUserSubRequestBody) GetUpdateOnlyOk() (*bool, bool)`

GetUpdateOnlyOk returns a tuple with the UpdateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnly

`func (o *UpsertUserSubRequestBody) SetUpdateOnly(v bool)`

SetUpdateOnly sets UpdateOnly field to given value.

### HasUpdateOnly

`func (o *UpsertUserSubRequestBody) HasUpdateOnly() bool`

HasUpdateOnly returns a boolean if a field has been set.

### SetUpdateOnlyNil

`func (o *UpsertUserSubRequestBody) SetUpdateOnlyNil(b bool)`

 SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil

### UnsetUpdateOnly
`func (o *UpsertUserSubRequestBody) UnsetUpdateOnly()`

UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


