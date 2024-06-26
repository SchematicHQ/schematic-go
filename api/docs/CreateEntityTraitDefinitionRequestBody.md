# CreateEntityTraitDefinitionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**EntityType** | **string** |  | 
**Hierarchy** | **[]string** |  | 
**TraitType** | **string** |  | 

## Methods

### NewCreateEntityTraitDefinitionRequestBody

`func NewCreateEntityTraitDefinitionRequestBody(entityType string, hierarchy []string, traitType string, ) *CreateEntityTraitDefinitionRequestBody`

NewCreateEntityTraitDefinitionRequestBody instantiates a new CreateEntityTraitDefinitionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntityTraitDefinitionRequestBodyWithDefaults

`func NewCreateEntityTraitDefinitionRequestBodyWithDefaults() *CreateEntityTraitDefinitionRequestBody`

NewCreateEntityTraitDefinitionRequestBodyWithDefaults instantiates a new CreateEntityTraitDefinitionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateEntityTraitDefinitionRequestBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateEntityTraitDefinitionRequestBody) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateEntityTraitDefinitionRequestBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateEntityTraitDefinitionRequestBody) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateEntityTraitDefinitionRequestBody) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateEntityTraitDefinitionRequestBody) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEntityType

`func (o *CreateEntityTraitDefinitionRequestBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateEntityTraitDefinitionRequestBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateEntityTraitDefinitionRequestBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetHierarchy

`func (o *CreateEntityTraitDefinitionRequestBody) GetHierarchy() []string`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *CreateEntityTraitDefinitionRequestBody) GetHierarchyOk() (*[]string, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *CreateEntityTraitDefinitionRequestBody) SetHierarchy(v []string)`

SetHierarchy sets Hierarchy field to given value.


### GetTraitType

`func (o *CreateEntityTraitDefinitionRequestBody) GetTraitType() string`

GetTraitType returns the TraitType field if non-nil, zero value otherwise.

### GetTraitTypeOk

`func (o *CreateEntityTraitDefinitionRequestBody) GetTraitTypeOk() (*string, bool)`

GetTraitTypeOk returns a tuple with the TraitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitType

`func (o *CreateEntityTraitDefinitionRequestBody) SetTraitType(v string)`

SetTraitType sets TraitType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


