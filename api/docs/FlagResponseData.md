# FlagResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**DefaultValue** | **bool** |  | 
**Description** | **string** |  | 
**FeatureId** | Pointer to **NullableString** |  | [optional] 
**FlagType** | **string** |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewFlagResponseData

`func NewFlagResponseData(createdAt time.Time, defaultValue bool, description string, flagType string, id string, key string, name string, updatedAt time.Time, ) *FlagResponseData`

NewFlagResponseData instantiates a new FlagResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagResponseDataWithDefaults

`func NewFlagResponseDataWithDefaults() *FlagResponseData`

NewFlagResponseDataWithDefaults instantiates a new FlagResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FlagResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlagResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlagResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultValue

`func (o *FlagResponseData) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FlagResponseData) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FlagResponseData) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *FlagResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatureId

`func (o *FlagResponseData) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *FlagResponseData) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *FlagResponseData) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *FlagResponseData) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### SetFeatureIdNil

`func (o *FlagResponseData) SetFeatureIdNil(b bool)`

 SetFeatureIdNil sets the value for FeatureId to be an explicit nil

### UnsetFeatureId
`func (o *FlagResponseData) UnsetFeatureId()`

UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
### GetFlagType

`func (o *FlagResponseData) GetFlagType() string`

GetFlagType returns the FlagType field if non-nil, zero value otherwise.

### GetFlagTypeOk

`func (o *FlagResponseData) GetFlagTypeOk() (*string, bool)`

GetFlagTypeOk returns a tuple with the FlagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagType

`func (o *FlagResponseData) SetFlagType(v string)`

SetFlagType sets FlagType field to given value.


### GetId

`func (o *FlagResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *FlagResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagResponseData) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *FlagResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *FlagResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlagResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlagResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


