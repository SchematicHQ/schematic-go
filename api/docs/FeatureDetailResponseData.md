# FeatureDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**EventSubtype** | Pointer to **NullableString** |  | [optional] 
**EventSummary** | Pointer to [**EventSummaryResponseData**](EventSummaryResponseData.md) |  | [optional] 
**FeatureType** | **string** |  | 
**Flags** | [**[]FlagDetailResponseData**](FlagDetailResponseData.md) |  | 
**Id** | **string** |  | 
**LifecyclePhase** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Plans** | [**[]PreviewObject**](PreviewObject.md) |  | 
**Trait** | Pointer to [**EntityTraitDefinitionResponseData**](EntityTraitDefinitionResponseData.md) |  | [optional] 
**TraitId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewFeatureDetailResponseData

`func NewFeatureDetailResponseData(createdAt time.Time, description string, featureType string, flags []FlagDetailResponseData, id string, name string, plans []PreviewObject, updatedAt time.Time, ) *FeatureDetailResponseData`

NewFeatureDetailResponseData instantiates a new FeatureDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureDetailResponseDataWithDefaults

`func NewFeatureDetailResponseDataWithDefaults() *FeatureDetailResponseData`

NewFeatureDetailResponseDataWithDefaults instantiates a new FeatureDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FeatureDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FeatureDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FeatureDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *FeatureDetailResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FeatureDetailResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FeatureDetailResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEventSubtype

`func (o *FeatureDetailResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *FeatureDetailResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *FeatureDetailResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *FeatureDetailResponseData) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### SetEventSubtypeNil

`func (o *FeatureDetailResponseData) SetEventSubtypeNil(b bool)`

 SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil

### UnsetEventSubtype
`func (o *FeatureDetailResponseData) UnsetEventSubtype()`

UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
### GetEventSummary

`func (o *FeatureDetailResponseData) GetEventSummary() EventSummaryResponseData`

GetEventSummary returns the EventSummary field if non-nil, zero value otherwise.

### GetEventSummaryOk

`func (o *FeatureDetailResponseData) GetEventSummaryOk() (*EventSummaryResponseData, bool)`

GetEventSummaryOk returns a tuple with the EventSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSummary

`func (o *FeatureDetailResponseData) SetEventSummary(v EventSummaryResponseData)`

SetEventSummary sets EventSummary field to given value.

### HasEventSummary

`func (o *FeatureDetailResponseData) HasEventSummary() bool`

HasEventSummary returns a boolean if a field has been set.

### GetFeatureType

`func (o *FeatureDetailResponseData) GetFeatureType() string`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *FeatureDetailResponseData) GetFeatureTypeOk() (*string, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *FeatureDetailResponseData) SetFeatureType(v string)`

SetFeatureType sets FeatureType field to given value.


### GetFlags

`func (o *FeatureDetailResponseData) GetFlags() []FlagDetailResponseData`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *FeatureDetailResponseData) GetFlagsOk() (*[]FlagDetailResponseData, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *FeatureDetailResponseData) SetFlags(v []FlagDetailResponseData)`

SetFlags sets Flags field to given value.


### GetId

`func (o *FeatureDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLifecyclePhase

`func (o *FeatureDetailResponseData) GetLifecyclePhase() string`

GetLifecyclePhase returns the LifecyclePhase field if non-nil, zero value otherwise.

### GetLifecyclePhaseOk

`func (o *FeatureDetailResponseData) GetLifecyclePhaseOk() (*string, bool)`

GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecyclePhase

`func (o *FeatureDetailResponseData) SetLifecyclePhase(v string)`

SetLifecyclePhase sets LifecyclePhase field to given value.

### HasLifecyclePhase

`func (o *FeatureDetailResponseData) HasLifecyclePhase() bool`

HasLifecyclePhase returns a boolean if a field has been set.

### SetLifecyclePhaseNil

`func (o *FeatureDetailResponseData) SetLifecyclePhaseNil(b bool)`

 SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil

### UnsetLifecyclePhase
`func (o *FeatureDetailResponseData) UnsetLifecyclePhase()`

UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
### GetName

`func (o *FeatureDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetPlans

`func (o *FeatureDetailResponseData) GetPlans() []PreviewObject`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *FeatureDetailResponseData) GetPlansOk() (*[]PreviewObject, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *FeatureDetailResponseData) SetPlans(v []PreviewObject)`

SetPlans sets Plans field to given value.


### GetTrait

`func (o *FeatureDetailResponseData) GetTrait() EntityTraitDefinitionResponseData`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *FeatureDetailResponseData) GetTraitOk() (*EntityTraitDefinitionResponseData, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *FeatureDetailResponseData) SetTrait(v EntityTraitDefinitionResponseData)`

SetTrait sets Trait field to given value.

### HasTrait

`func (o *FeatureDetailResponseData) HasTrait() bool`

HasTrait returns a boolean if a field has been set.

### GetTraitId

`func (o *FeatureDetailResponseData) GetTraitId() string`

GetTraitId returns the TraitId field if non-nil, zero value otherwise.

### GetTraitIdOk

`func (o *FeatureDetailResponseData) GetTraitIdOk() (*string, bool)`

GetTraitIdOk returns a tuple with the TraitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitId

`func (o *FeatureDetailResponseData) SetTraitId(v string)`

SetTraitId sets TraitId field to given value.

### HasTraitId

`func (o *FeatureDetailResponseData) HasTraitId() bool`

HasTraitId returns a boolean if a field has been set.

### SetTraitIdNil

`func (o *FeatureDetailResponseData) SetTraitIdNil(b bool)`

 SetTraitIdNil sets the value for TraitId to be an explicit nil

### UnsetTraitId
`func (o *FeatureDetailResponseData) UnsetTraitId()`

UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
### GetUpdatedAt

`func (o *FeatureDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FeatureDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FeatureDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


