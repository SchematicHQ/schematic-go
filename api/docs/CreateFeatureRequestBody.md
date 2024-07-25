# CreateFeatureRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**FeatureType** | **string** |  | 
**Flag** | Pointer to [**CreateOrUpdateFlagRequestBody**](CreateOrUpdateFlagRequestBody.md) |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**LifecyclePhase** | Pointer to **NullableString** |  | [optional] 
**MaintainerId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**TraitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateFeatureRequestBody

`func NewCreateFeatureRequestBody(description string, featureType string, name string, ) *CreateFeatureRequestBody`

NewCreateFeatureRequestBody instantiates a new CreateFeatureRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeatureRequestBodyWithDefaults

`func NewCreateFeatureRequestBodyWithDefaults() *CreateFeatureRequestBody`

NewCreateFeatureRequestBodyWithDefaults instantiates a new CreateFeatureRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateFeatureRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFeatureRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFeatureRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventSubtype

`func (o *CreateFeatureRequestBody) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *CreateFeatureRequestBody) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *CreateFeatureRequestBody) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *CreateFeatureRequestBody) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *CreateFeatureRequestBody) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *CreateFeatureRequestBody) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetFeatureType

`func (o *CreateFeatureRequestBody) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *CreateFeatureRequestBody) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *CreateFeatureRequestBody) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.


### GetFlag

`func (o *CreateFeatureRequestBody) GetFlag() CreateOrUpdateFlagRequestBody`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *CreateFeatureRequestBody) GetFlagOk() (*CreateOrUpdateFlagRequestBody, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *CreateFeatureRequestBody) SetFlag(v CreateOrUpdateFlagRequestBody)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *CreateFeatureRequestBody) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetIcon

`func (o *CreateFeatureRequestBody) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateFeatureRequestBody) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateFeatureRequestBody) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateFeatureRequestBody) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateFeatureRequestBody) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateFeatureRequestBody) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetLifecyclePhase

`func (o *CreateFeatureRequestBody) GetLifecyclePhase() string`

GetLifecyclePhase returns the LifecyclePhase field if non-nil, zero value otherwise.

### GetLifecyclePhaseOk

`func (o *CreateFeatureRequestBody) GetLifecyclePhaseOk() (*string, bool)`

GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecyclePhase

`func (o *CreateFeatureRequestBody) SetLifecyclePhase(v string)`

SetLifecyclePhase sets LifecyclePhase field to given value.

### HasLifecyclePhase

`func (o *CreateFeatureRequestBody) HasLifecyclePhase() bool`

HasLifecyclePhase returns a boolean if a field has been set.

### SetLifecyclePhaseNil

`func (o *CreateFeatureRequestBody) SetLifecyclePhaseNil(b bool)`

 SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil

### UnsetLifecyclePhase
`func (o *CreateFeatureRequestBody) UnsetLifecyclePhase()`

UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
### GetMaintainerId

`func (o *CreateFeatureRequestBody) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *CreateFeatureRequestBody) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *CreateFeatureRequestBody) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *CreateFeatureRequestBody) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### SetMaintainerIdNil

`func (o *CreateFeatureRequestBody) SetMaintainerIdNil(b bool)`

 SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil

### UnsetMaintainerId
`func (o *CreateFeatureRequestBody) UnsetMaintainerId()`

UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
### GetName

`func (o *CreateFeatureRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFeatureRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFeatureRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetTraitId

`func (o *CreateFeatureRequestBody) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *CreateFeatureRequestBody) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *CreateFeatureRequestBody) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *CreateFeatureRequestBody) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *CreateFeatureRequestBody) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *CreateFeatureRequestBody) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


