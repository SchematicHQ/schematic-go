# UserDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyMemberships** | [**[]CompanyMembershipDetailResponseData**](CompanyMembershipDetailResponseData.md) |  | 
**CreatedAt** | **time.Time** |  | 
**EntityTraits** | [**[]EntityTraitDetailResponseData**](EntityTraitDetailResponseData.md) |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**Keys** | [**[]EntityKeyDetailResponseData**](EntityKeyDetailResponseData.md) |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUserDetailResponseData

`func NewUserDetailResponseData(companyMemberships []CompanyMembershipDetailResponseData, createdAt time.Time, entityTraits []EntityTraitDetailResponseData, environmentId string, id string, keys []EntityKeyDetailResponseData, name string, updatedAt time.Time, ) *UserDetailResponseData`

NewUserDetailResponseData instantiates a new UserDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailResponseDataWithDefaults

`func NewUserDetailResponseDataWithDefaults() *UserDetailResponseData`

NewUserDetailResponseDataWithDefaults instantiates a new UserDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyMemberships

`func (o *UserDetailResponseData) GetCompanyMemberships() []CompanyMembershipDetailResponseData`

GetCompanyMemberships returns the CompanyMemberships field if non-nil, zero value otherwise.

### GetCompanyMembershipsOk

`func (o *UserDetailResponseData) GetCompanyMembershipsOk() (*[]CompanyMembershipDetailResponseData, bool)`

GetCompanyMembershipsOk returns a tuple with the CompanyMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyMemberships

`func (o *UserDetailResponseData) SetCompanyMemberships(v []CompanyMembershipDetailResponseData)`

SetCompanyMemberships sets CompanyMemberships field to given value.


### GetCreatedAt

`func (o *UserDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEntityTraits

`func (o *UserDetailResponseData) GetEntityTraits() []EntityTraitDetailResponseData`

GetEntityTraits returns the EntityTraits field if non-nil, zero value otherwise.

### GetEntityTraitsOk

`func (o *UserDetailResponseData) GetEntityTraitsOk() (*[]EntityTraitDetailResponseData, bool)`

GetEntityTraitsOk returns a tuple with the EntityTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTraits

`func (o *UserDetailResponseData) SetEntityTraits(v []EntityTraitDetailResponseData)`

SetEntityTraits sets EntityTraits field to given value.


### GetEnvironmentId

`func (o *UserDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *UserDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetKeys

`func (o *UserDetailResponseData) GetKeys() []EntityKeyDetailResponseData`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UserDetailResponseData) GetKeysOk() (*[]EntityKeyDetailResponseData, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UserDetailResponseData) SetKeys(v []EntityKeyDetailResponseData)`

SetKeys sets Keys field to given value.


### GetLastSeenAt

`func (o *UserDetailResponseData) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *UserDetailResponseData) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *UserDetailResponseData) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *UserDetailResponseData) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *UserDetailResponseData) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *UserDetailResponseData) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *UserDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetTraits

`func (o *UserDetailResponseData) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *UserDetailResponseData) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *UserDetailResponseData) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *UserDetailResponseData) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


