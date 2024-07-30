# FeatureResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**FeatureType** | **string** |  | 
**Icon** | **string** |  | 
**Id** | **string** |  | 
**LifecyclePhase** | Pointer to **NullableString** |  | [optional] 
**MaintainerId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**TraitId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewFeatureResponseData

`func NewFeatureResponseData(createdAt time.Time, description string, featureType string, icon string, id string, name string, updatedAt time.Time, ) *FeatureResponseData`

NewFeatureResponseData instantiates a new FeatureResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureResponseDataWithDefaults

`func NewFeatureResponseDataWithDefaults() *FeatureResponseData`

NewFeatureResponseDataWithDefaults instantiates a new FeatureResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FeatureResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *FeatureResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventSubtype

`func (o *FeatureResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *FeatureResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *FeatureResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *FeatureResponseData) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *FeatureResponseData) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *FeatureResponseData) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetFeatureType

`func (o *FeatureResponseData) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *FeatureResponseData) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *FeatureResponseData) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.


### GetIcon

`func (o *FeatureResponseData) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FeatureResponseData) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FeatureResponseData) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetId

`func (o *FeatureResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLifecyclePhase

`func (o *FeatureResponseData) GetLifecyclePhase() string`

GetLifecyclePhase returns the LifecyclePhase field if non-nil, zero value otherwise.

### GetLifecyclePhaseOk

`func (o *FeatureResponseData) GetLifecyclePhaseOk() (*string, bool)`

GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecyclePhase

`func (o *FeatureResponseData) SetLifecyclePhase(v string)`

SetLifecyclePhase sets LifecyclePhase field to given value.

### HasLifecyclePhase

`func (o *FeatureResponseData) HasLifecyclePhase() bool`

HasLifecyclePhase returns a boolean if a field has been set.

### SetLifecyclePhaseNil

`func (o *FeatureResponseData) SetLifecyclePhaseNil(b bool)`

 SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil

### UnsetLifecyclePhase
`func (o *FeatureResponseData) UnsetLifecyclePhase()`

UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
### GetMaintainerId

`func (o *FeatureResponseData) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FeatureResponseData) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FeatureResponseData) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *FeatureResponseData) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### SetMaintainerIdNil

`func (o *FeatureResponseData) SetMaintainerIdNil(b bool)`

 SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil

### UnsetMaintainerId
`func (o *FeatureResponseData) UnsetMaintainerId()`

UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
### GetName

`func (o *FeatureResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetTraitId

`func (o *FeatureResponseData) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *FeatureResponseData) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *FeatureResponseData) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *FeatureResponseData) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *FeatureResponseData) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *FeatureResponseData) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
### GetUpdatedAt

`func (o *FeatureResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeatureResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeatureResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


