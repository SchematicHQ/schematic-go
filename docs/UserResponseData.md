# UserResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUserResponseData

`func NewUserResponseData(createdAt time.Time, environmentId string, id string, name string, updatedAt time.Time, ) *UserResponseData`

NewUserResponseData instantiates a new UserResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseDataWithDefaults

`func NewUserResponseDataWithDefaults() *UserResponseData`

NewUserResponseDataWithDefaults instantiates a new UserResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *UserResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UserResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UserResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *UserResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeenAt

`func (o *UserResponseData) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *UserResponseData) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *UserResponseData) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *UserResponseData) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *UserResponseData) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *UserResponseData) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *UserResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *UserResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


