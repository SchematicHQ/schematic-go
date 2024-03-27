# CreateFlagRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | **bool** |  | 
**Description** | **string** |  | 
**FeatureId** | Pointer to **NullableString** |  | [optional] 
**FlagType** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewCreateFlagRequestBody

`func NewCreateFlagRequestBody(defaultValue bool, description string, flagType string, key string, name string, ) *CreateFlagRequestBody`

NewCreateFlagRequestBody instantiates a new CreateFlagRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlagRequestBodyWithDefaults

`func NewCreateFlagRequestBodyWithDefaults() *CreateFlagRequestBody`

NewCreateFlagRequestBodyWithDefaults instantiates a new CreateFlagRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *CreateFlagRequestBody) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *CreateFlagRequestBody) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *CreateFlagRequestBody) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *CreateFlagRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFlagRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFlagRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatureId

`func (o *CreateFlagRequestBody) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *CreateFlagRequestBody) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *CreateFlagRequestBody) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *CreateFlagRequestBody) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### SetFeatureIdNil

`func (o *CreateFlagRequestBody) SetFeatureIdNil(b bool)`

 SetFeatureIdNil sets the value for FeatureId to be an explicit nil

### UnsetFeatureId
`func (o *CreateFlagRequestBody) UnsetFeatureId()`

UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
### GetFlagType

`func (o *CreateFlagRequestBody) GetFlagType() string`

GetFlagType returns the FlagType field if non-nil, zero value otherwise.

### GetFlagTypeOk

`func (o *CreateFlagRequestBody) GetFlagTypeOk() (*string, bool)`

GetFlagTypeOk returns a tuple with the FlagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagType

`func (o *CreateFlagRequestBody) SetFlagType(v string)`

SetFlagType sets FlagType field to given value.


### GetKey

`func (o *CreateFlagRequestBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateFlagRequestBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateFlagRequestBody) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateFlagRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFlagRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFlagRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


