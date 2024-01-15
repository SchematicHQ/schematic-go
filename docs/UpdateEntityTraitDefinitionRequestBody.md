# UpdateEntityTraitDefinitionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 
**TraitType** | **string** |  | 

## Methods

### NewUpdateEntityTraitDefinitionRequestBody

`func NewUpdateEntityTraitDefinitionRequestBody(traitType string, ) *UpdateEntityTraitDefinitionRequestBody`

NewUpdateEntityTraitDefinitionRequestBody instantiates a new UpdateEntityTraitDefinitionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntityTraitDefinitionRequestBodyWithDefaults

`func NewUpdateEntityTraitDefinitionRequestBodyWithDefaults() *UpdateEntityTraitDefinitionRequestBody`

NewUpdateEntityTraitDefinitionRequestBodyWithDefaults instantiates a new UpdateEntityTraitDefinitionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipWebhooks

`func (o *UpdateEntityTraitDefinitionRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *UpdateEntityTraitDefinitionRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *UpdateEntityTraitDefinitionRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *UpdateEntityTraitDefinitionRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *UpdateEntityTraitDefinitionRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *UpdateEntityTraitDefinitionRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
### GetTraitType

`func (o *UpdateEntityTraitDefinitionRequestBody) GetTraitType() string`

GetTraitType returns the TraitType field if non-nil, zero value otherwise.

### GetTraitTypeOk

`func (o *UpdateEntityTraitDefinitionRequestBody) GetTraitTypeOk() (*string, bool)`

GetTraitTypeOk returns a tuple with the TraitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitType

`func (o *UpdateEntityTraitDefinitionRequestBody) SetTraitType(v string)`

SetTraitType sets TraitType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


