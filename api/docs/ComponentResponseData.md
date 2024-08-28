# ComponentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ast** | Pointer to **map[string]float32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**State** | **string** |  | 
**Type** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewComponentResponseData

`func NewComponentResponseData(createdAt time.Time, id string, name string, state string, type_ string, updatedAt time.Time, ) *ComponentResponseData`

NewComponentResponseData instantiates a new ComponentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentResponseDataWithDefaults

`func NewComponentResponseDataWithDefaults() *ComponentResponseData`

NewComponentResponseDataWithDefaults instantiates a new ComponentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAst

`func (o *ComponentResponseData) GetAst() map[string]float32`

GetAst returns the Ast field if non-nil, zero value otherwise.

### GetAstOk

`func (o *ComponentResponseData) GetAstOk() (*map[string]float32, bool)`

GetAstOk returns a tuple with the Ast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAst

`func (o *ComponentResponseData) SetAst(v map[string]float32)`

SetAst sets Ast field to given value.

### HasAst

`func (o *ComponentResponseData) HasAst() bool`

HasAst returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ComponentResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ComponentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ComponentResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *ComponentResponseData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ComponentResponseData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ComponentResponseData) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *ComponentResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ComponentResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


