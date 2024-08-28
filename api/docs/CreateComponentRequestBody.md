# CreateComponentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ast** | Pointer to **map[string]float32** |  | [optional] 
**EntityType** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewCreateComponentRequestBody

`func NewCreateComponentRequestBody(entityType string, name string, ) *CreateComponentRequestBody`

NewCreateComponentRequestBody instantiates a new CreateComponentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComponentRequestBodyWithDefaults

`func NewCreateComponentRequestBodyWithDefaults() *CreateComponentRequestBody`

NewCreateComponentRequestBodyWithDefaults instantiates a new CreateComponentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAst

`func (o *CreateComponentRequestBody) GetAst() map[string]float32`

GetAst returns the Ast field if non-nil, zero value otherwise.

### GetAstOk

`func (o *CreateComponentRequestBody) GetAstOk() (*map[string]float32, bool)`

GetAstOk returns a tuple with the Ast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAst

`func (o *CreateComponentRequestBody) SetAst(v map[string]float32)`

SetAst sets Ast field to given value.

### HasAst

`func (o *CreateComponentRequestBody) HasAst() bool`

HasAst returns a boolean if a field has been set.

### GetEntityType

`func (o *CreateComponentRequestBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateComponentRequestBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateComponentRequestBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetName

`func (o *CreateComponentRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateComponentRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateComponentRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


