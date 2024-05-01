# EntityTraitDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Definition** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**DefinitionId** | **string** |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Value** | **string** |  | 

## Methods

### NewEntityTraitDetailResponseData

`func NewEntityTraitDetailResponseData(createdAt time.Time, definitionId string, environmentId string, id string, updatedAt time.Time, value string, ) *EntityTraitDetailResponseData`

NewEntityTraitDetailResponseData instantiates a new EntityTraitDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTraitDetailResponseDataWithDefaults

`func NewEntityTraitDetailResponseDataWithDefaults() *EntityTraitDetailResponseData`

NewEntityTraitDetailResponseDataWithDefaults instantiates a new EntityTraitDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EntityTraitDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityTraitDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityTraitDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefinition

`func (o *EntityTraitDetailResponseData) GetDefinition() EntityTraitDefinitionResponseData`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *EntityTraitDetailResponseData) GetDefinitionOk() (*EntityTraitDefinitionResponseData, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *EntityTraitDetailResponseData) SetDefinition(v EntityTraitDefinitionResponseData)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *EntityTraitDetailResponseData) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDefinitionId

`func (o *EntityTraitDetailResponseData) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *EntityTraitDetailResponseData) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *EntityTraitDetailResponseData) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.


### GetEnvironmentId

`func (o *EntityTraitDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EntityTraitDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EntityTraitDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *EntityTraitDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityTraitDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityTraitDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *EntityTraitDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityTraitDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityTraitDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *EntityTraitDetailResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityTraitDetailResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityTraitDetailResponseData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


