# UpdateComponentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ast** | Pointer to **map[string]float32** |  | [optional] 
**EntityType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateComponentRequestBody

`func NewUpdateComponentRequestBody() *UpdateComponentRequestBody`

NewUpdateComponentRequestBody instantiates a new UpdateComponentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateComponentRequestBodyWithDefaults

`func NewUpdateComponentRequestBodyWithDefaults() *UpdateComponentRequestBody`

NewUpdateComponentRequestBodyWithDefaults instantiates a new UpdateComponentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAst

`func (o *UpdateComponentRequestBody) GetAst() map[string]float32`

GetAst returns the Ast field if non-nil, zero value otherwise.

### GetAstOk

`func (o *UpdateComponentRequestBody) GetAstOk() (*map[string]float32, bool)`

GetAstOk returns a tuple with the Ast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAst

`func (o *UpdateComponentRequestBody) SetAst(v map[string]float32)`

SetAst sets Ast field to given value.

### HasAst

`func (o *UpdateComponentRequestBody) HasAst() bool`

HasAst returns a boolean if a field has been set.

### GetEntityType

`func (o *UpdateComponentRequestBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UpdateComponentRequestBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UpdateComponentRequestBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *UpdateComponentRequestBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityTypeNil

`func (o *UpdateComponentRequestBody) SetEntityTypeNil(b bool)`

 SetEntityTypeNil sets the value for EntityType to be an explicit nil

### UnsetEntityType
`func (o *UpdateComponentRequestBody) UnsetEntityType()`

UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
### GetName

`func (o *UpdateComponentRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateComponentRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateComponentRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateComponentRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateComponentRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateComponentRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *UpdateComponentRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateComponentRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateComponentRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateComponentRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateComponentRequestBody) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateComponentRequestBody) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


