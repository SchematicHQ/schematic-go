# EntityKeyDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Definition** | Pointer to [**EntityKeyDefinitionResponseData**](EntityKeyDefinitionResponseData.md) |  | [optional] 
**DefinitionId** | **string** |  | 
**EntityId** | **string** |  | 
**EntityType** | **string** |  | 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Value** | **string** |  | 

## Methods

### NewEntityKeyDetailResponseData

`func NewEntityKeyDetailResponseData(createdAt time.Time, definitionId string, entityId string, entityType string, environmentId string, id string, key string, updatedAt time.Time, value string, ) *EntityKeyDetailResponseData`

NewEntityKeyDetailResponseData instantiates a new EntityKeyDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityKeyDetailResponseDataWithDefaults

`func NewEntityKeyDetailResponseDataWithDefaults() *EntityKeyDetailResponseData`

NewEntityKeyDetailResponseDataWithDefaults instantiates a new EntityKeyDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EntityKeyDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityKeyDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityKeyDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefinition

`func (o *EntityKeyDetailResponseData) GetDefinition() EntityKeyDefinitionResponseData`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *EntityKeyDetailResponseData) GetDefinitionOk() (*EntityKeyDefinitionResponseData, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *EntityKeyDetailResponseData) SetDefinition(v EntityKeyDefinitionResponseData)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *EntityKeyDetailResponseData) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetDefinitionId

`func (o *EntityKeyDetailResponseData) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *EntityKeyDetailResponseData) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *EntityKeyDetailResponseData) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.


### GetEntityId

`func (o *EntityKeyDetailResponseData) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityKeyDetailResponseData) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityKeyDetailResponseData) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *EntityKeyDetailResponseData) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *EntityKeyDetailResponseData) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *EntityKeyDetailResponseData) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetEnvironmentId

`func (o *EntityKeyDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EntityKeyDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EntityKeyDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *EntityKeyDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityKeyDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityKeyDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *EntityKeyDetailResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityKeyDetailResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityKeyDetailResponseData) SetKey(v string)`

SetKey sets Key field to given value.


### GetUpdatedAt

`func (o *EntityKeyDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityKeyDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityKeyDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValue

`func (o *EntityKeyDetailResponseData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityKeyDetailResponseData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityKeyDetailResponseData) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


