# CompanyDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**Keys** | [**[]EntityKeyResponseData**](EntityKeyResponseData.md) |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCompanyDetailResponseData

`func NewCompanyDetailResponseData(createdAt time.Time, environmentId string, id string, keys []EntityKeyResponseData, name string, updatedAt time.Time, ) *CompanyDetailResponseData`

NewCompanyDetailResponseData instantiates a new CompanyDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDetailResponseDataWithDefaults

`func NewCompanyDetailResponseDataWithDefaults() *CompanyDetailResponseData`

NewCompanyDetailResponseDataWithDefaults instantiates a new CompanyDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CompanyDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *CompanyDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CompanyDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CompanyDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *CompanyDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetKeys

`func (o *CompanyDetailResponseData) GetKeys() []EntityKeyResponseData`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CompanyDetailResponseData) GetKeysOk() (*[]EntityKeyResponseData, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CompanyDetailResponseData) SetKeys(v []EntityKeyResponseData)`

SetKeys sets Keys field to given value.


### GetLastSeenAt

`func (o *CompanyDetailResponseData) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *CompanyDetailResponseData) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *CompanyDetailResponseData) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *CompanyDetailResponseData) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *CompanyDetailResponseData) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *CompanyDetailResponseData) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *CompanyDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetTraits

`func (o *CompanyDetailResponseData) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *CompanyDetailResponseData) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *CompanyDetailResponseData) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *CompanyDetailResponseData) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompanyDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


