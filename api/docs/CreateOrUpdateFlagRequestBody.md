# CreateOrUpdateFlagRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | **bool** |  | 
**Description** | **string** |  | 
**FeatureId** | Pointer to **NullableString** |  | [optional] 
**FlagType** | **string** |  | 
**Id** | Pointer to **NullableString** |  | [optional] 
**Key** | **string** |  | 
**MaintainerId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewCreateOrUpdateFlagRequestBody

`func NewCreateOrUpdateFlagRequestBody(defaultValue bool, description string, flagType string, key string, name string, ) *CreateOrUpdateFlagRequestBody`

NewCreateOrUpdateFlagRequestBody instantiates a new CreateOrUpdateFlagRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateFlagRequestBodyWithDefaults

`func NewCreateOrUpdateFlagRequestBodyWithDefaults() *CreateOrUpdateFlagRequestBody`

NewCreateOrUpdateFlagRequestBodyWithDefaults instantiates a new CreateOrUpdateFlagRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *CreateOrUpdateFlagRequestBody) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateOrUpdateFlagRequestBody) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateOrUpdateFlagRequestBody) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *CreateOrUpdateFlagRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateFlagRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateFlagRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatureId

`func (o *CreateOrUpdateFlagRequestBody) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CreateOrUpdateFlagRequestBody) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CreateOrUpdateFlagRequestBody) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CreateOrUpdateFlagRequestBody) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### SetFeatureIdNil

`func (o *CreateOrUpdateFlagRequestBody) SetFeatureIdNil(b bool)`

 SetFeatureIdNil sets the value for FeatureId to be an explicit nil

### UnsetFeatureId
`func (o *CreateOrUpdateFlagRequestBody) UnsetFeatureId()`

UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
### GetFlagType

`func (o *CreateOrUpdateFlagRequestBody) GetFlagType() string`

GetFlagType returns the FlagType field if non-nil, zero value otherwise.

### GetFlagTypeOk

`func (o *CreateOrUpdateFlagRequestBody) GetFlagTypeOk() (*string, bool)`

GetFlagTypeOk returns a tuple with the FlagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagType

`func (o *CreateOrUpdateFlagRequestBody) SetFlagType(v string)`

SetFlagType sets FlagType field to given value.


### GetId

`func (o *CreateOrUpdateFlagRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrUpdateFlagRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrUpdateFlagRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrUpdateFlagRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateOrUpdateFlagRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateOrUpdateFlagRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKey

`func (o *CreateOrUpdateFlagRequestBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateOrUpdateFlagRequestBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateOrUpdateFlagRequestBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaintainerId

`func (o *CreateOrUpdateFlagRequestBody) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CreateOrUpdateFlagRequestBody) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CreateOrUpdateFlagRequestBody) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *CreateOrUpdateFlagRequestBody) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### SetMaintainerIdNil

`func (o *CreateOrUpdateFlagRequestBody) SetMaintainerIdNil(b bool)`

 SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil

### UnsetMaintainerId
`func (o *CreateOrUpdateFlagRequestBody) UnsetMaintainerId()`

UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
### GetName

`func (o *CreateOrUpdateFlagRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateFlagRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateFlagRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


