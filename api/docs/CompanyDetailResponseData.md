# CompanyDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOns** | [**[]PreviewObject**](PreviewObject.md) |  | 
**CreatedAt** | **time.Time** |  | 
**EntityTraits** | [**[]EntityTraitDetailResponseData**](EntityTraitDetailResponseData.md) |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**Keys** | [**[]EntityKeyDetailResponseData**](EntityKeyDetailResponseData.md) |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Plan** | Pointer to [**PreviewObject**](PreviewObject.md) |  | [optional] 
**Plans** | [**[]PreviewObject**](PreviewObject.md) |  | 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdatedAt** | **time.Time** |  | 
**UserCount** | **int32** |  | 

## Methods

### NewCompanyDetailResponseData

`func NewCompanyDetailResponseData(addOns []PreviewObject, createdAt time.Time, entityTraits []EntityTraitDetailResponseData, environmentId string, id string, keys []EntityKeyDetailResponseData, name string, plans []PreviewObject, updatedAt time.Time, userCount int32, ) *CompanyDetailResponseData`

NewCompanyDetailResponseData instantiates a new CompanyDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDetailResponseDataWithDefaults

`func NewCompanyDetailResponseDataWithDefaults() *CompanyDetailResponseData`

NewCompanyDetailResponseDataWithDefaults instantiates a new CompanyDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOns

`func (o *CompanyDetailResponseData) GetAddOns() []PreviewObject`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *CompanyDetailResponseData) GetAddOnsOk() (*[]PreviewObject, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *CompanyDetailResponseData) SetAddOns(v []PreviewObject)`

SetAddOns sets AddOns field to given value.


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


### GetEntityTraits

`func (o *CompanyDetailResponseData) GetEntityTraits() []EntityTraitDetailResponseData`

GetEntityTraits returns the EntityTraits field if non-nil, zero value otherwise.

### GetEntityTraitsOk

`func (o *CompanyDetailResponseData) GetEntityTraitsOk() (*[]EntityTraitDetailResponseData, bool)`

GetEntityTraitsOk returns a tuple with the EntityTraits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTraits

`func (o *CompanyDetailResponseData) SetEntityTraits(v []EntityTraitDetailResponseData)`

SetEntityTraits sets EntityTraits field to given value.


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

`func (o *CompanyDetailResponseData) GetKeys() []EntityKeyDetailResponseData`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CompanyDetailResponseData) GetKeysOk() (*[]EntityKeyDetailResponseData, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CompanyDetailResponseData) SetKeys(v []EntityKeyDetailResponseData)`

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
### GetLogoUrl

`func (o *CompanyDetailResponseData) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CompanyDetailResponseData) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CompanyDetailResponseData) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CompanyDetailResponseData) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *CompanyDetailResponseData) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *CompanyDetailResponseData) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
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


### GetPlan

`func (o *CompanyDetailResponseData) GetPlan() PreviewObject`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CompanyDetailResponseData) GetPlanOk() (*PreviewObject, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CompanyDetailResponseData) SetPlan(v PreviewObject)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CompanyDetailResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlans

`func (o *CompanyDetailResponseData) GetPlans() []PreviewObject`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CompanyDetailResponseData) GetPlansOk() (*[]PreviewObject, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CompanyDetailResponseData) SetPlans(v []PreviewObject)`

SetPlans sets Plans field to given value.


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


### GetUserCount

`func (o *CompanyDetailResponseData) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *CompanyDetailResponseData) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *CompanyDetailResponseData) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


