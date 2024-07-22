# CreateComponentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ast** | **[]int32** |  | 
**EntityType** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewCreateComponentRequestBody

`func NewCreateComponentRequestBody(ast []int32, entityType string, name string, ) *CreateComponentRequestBody`

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

`func (o *CreateComponentRequestBody) GetAst() []int32`

GetAst returns the Ast field if non-nil, zero value otherwise.

### GetAstOk

`func (o *CreateComponentRequestBody) GetAstOk() (*[]int32, bool)`

GetAstOk returns a tuple with the Ast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAst

`func (o *CreateComponentRequestBody) SetAst(v []int32)`

SetAst sets Ast field to given value.


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


