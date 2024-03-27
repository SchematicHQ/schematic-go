# UpdateFeatureRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**FeatureType** | Pointer to **NullableString** |  | [optional] 
**Flag** | Pointer to [**CreateOrUpdateFlagRequestBody**](CreateOrUpdateFlagRequestBody.md) |  | [optional] 
**LifecyclePhase** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TraitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateFeatureRequestBody

`func NewUpdateFeatureRequestBody() *UpdateFeatureRequestBody`

NewUpdateFeatureRequestBody instantiates a new UpdateFeatureRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureRequestBodyWithDefaults

`func NewUpdateFeatureRequestBodyWithDefaults() *UpdateFeatureRequestBody`

NewUpdateFeatureRequestBodyWithDefaults instantiates a new UpdateFeatureRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateFeatureRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFeatureRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFeatureRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFeatureRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateFeatureRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateFeatureRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventSubtype

`func (o *UpdateFeatureRequestBody) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *UpdateFeatureRequestBody) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *UpdateFeatureRequestBody) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *UpdateFeatureRequestBody) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *UpdateFeatureRequestBody) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *UpdateFeatureRequestBody) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetFeatureType

`func (o *UpdateFeatureRequestBody) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *UpdateFeatureRequestBody) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *UpdateFeatureRequestBody) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *UpdateFeatureRequestBody) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### SetFeatureTypeNil

`func (o *UpdateFeatureRequestBody) SetFeatureTypeNil(b bool)`

 SetFeatureTypeNil sets the value for FeatureType to be an explicit nil

### UnsetFeatureType
`func (o *UpdateFeatureRequestBody) UnsetFeatureType()`

UnsetFeatureType ensures that no value is present for FeatureType, not even an explicit nil
### GetFlag

`func (o *UpdateFeatureRequestBody) GetFlag() CreateOrUpdateFlagRequestBody`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *UpdateFeatureRequestBody) GetFlagOk() (*CreateOrUpdateFlagRequestBody, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *UpdateFeatureRequestBody) SetFlag(v CreateOrUpdateFlagRequestBody)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *UpdateFeatureRequestBody) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetLifecyclePhase

`func (o *UpdateFeatureRequestBody) GetLifecyclePhase() string`

GetLifecyclePhase returns the LifecyclePhase field if non-nil, zero value otherwise.

### GetLifecyclePhaseOk

`func (o *UpdateFeatureRequestBody) GetLifecyclePhaseOk() (*string, bool)`

GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecyclePhase

`func (o *UpdateFeatureRequestBody) SetLifecyclePhase(v string)`

SetLifecyclePhase sets LifecyclePhase field to given value.

### HasLifecyclePhase

`func (o *UpdateFeatureRequestBody) HasLifecyclePhase() bool`

HasLifecyclePhase returns a boolean if a field has been set.

### SetLifecyclePhaseNil

`func (o *UpdateFeatureRequestBody) SetLifecyclePhaseNil(b bool)`

 SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil

### UnsetLifecyclePhase
`func (o *UpdateFeatureRequestBody) UnsetLifecyclePhase()`

UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
### GetName

`func (o *UpdateFeatureRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFeatureRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFeatureRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFeatureRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateFeatureRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateFeatureRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTraitId

`func (o *UpdateFeatureRequestBody) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *UpdateFeatureRequestBody) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *UpdateFeatureRequestBody) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *UpdateFeatureRequestBody) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *UpdateFeatureRequestBody) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *UpdateFeatureRequestBody) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


