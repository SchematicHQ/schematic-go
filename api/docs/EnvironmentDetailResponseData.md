# EnvironmentDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | [**[]ApiKeyResponseData**](ApiKeyResponseData.md) |  | 
**CreatedAt** | **time.Time** |  | 
**EnvironmentType** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewEnvironmentDetailResponseData

`func NewEnvironmentDetailResponseData(apiKeys []ApiKeyResponseData, createdAt time.Time, environmentType string, id string, name string, updatedAt time.Time, ) *EnvironmentDetailResponseData`

NewEnvironmentDetailResponseData instantiates a new EnvironmentDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDetailResponseDataWithDefaults

`func NewEnvironmentDetailResponseDataWithDefaults() *EnvironmentDetailResponseData`

NewEnvironmentDetailResponseDataWithDefaults instantiates a new EnvironmentDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *EnvironmentDetailResponseData) GetApiKeys() []ApiKeyResponseData`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *EnvironmentDetailResponseData) GetApiKeysOk() (*[]ApiKeyResponseData, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *EnvironmentDetailResponseData) SetApiKeys(v []ApiKeyResponseData)`

SetApiKeys sets ApiKeys field to given value.


### GetCreatedAt

`func (o *EnvironmentDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentType

`func (o *EnvironmentDetailResponseData) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *EnvironmentDetailResponseData) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *EnvironmentDetailResponseData) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetId

`func (o *EnvironmentDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EnvironmentDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *EnvironmentDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


