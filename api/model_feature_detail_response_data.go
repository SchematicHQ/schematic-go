/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the FeatureDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureDetailResponseData{}

// FeatureDetailResponseData struct for FeatureDetailResponseData
type FeatureDetailResponseData struct {
	CreatedAt            time.Time                          `json:"created_at"`
	Description          string                             `json:"description"`
	EventSubtype         NullableString                     `json:"event_subtype,omitempty"`
	EventSummary         *EventSummaryResponseData          `json:"event_summary,omitempty"`
	FeatureType          string                             `json:"feature_type"`
	Flags                []FlagDetailResponseData           `json:"flags"`
	Icon                 string                             `json:"icon"`
	Id                   string                             `json:"id"`
	LifecyclePhase       NullableString                     `json:"lifecycle_phase,omitempty"`
	MaintainerId         NullableString                     `json:"maintainer_id,omitempty"`
	Name                 string                             `json:"name"`
	Plans                []PreviewObject                    `json:"plans"`
	Trait                *EntityTraitDefinitionResponseData `json:"trait,omitempty"`
	TraitId              NullableString                     `json:"trait_id,omitempty"`
	UpdatedAt            time.Time                          `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _FeatureDetailResponseData FeatureDetailResponseData

// NewFeatureDetailResponseData instantiates a new FeatureDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureDetailResponseData(createdAt time.Time, description string, featureType string, flags []FlagDetailResponseData, icon string, id string, name string, plans []PreviewObject, updatedAt time.Time) *FeatureDetailResponseData {
	this := FeatureDetailResponseData{}
	this.CreatedAt = createdAt
	this.Description = description
	this.FeatureType = featureType
	this.Flags = flags
	this.Icon = icon
	this.Id = id
	this.Name = name
	this.Plans = plans
	this.UpdatedAt = updatedAt
	return &this
}

// NewFeatureDetailResponseDataWithDefaults instantiates a new FeatureDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureDetailResponseDataWithDefaults() *FeatureDetailResponseData {
	this := FeatureDetailResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FeatureDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FeatureDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *FeatureDetailResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FeatureDetailResponseData) SetDescription(v string) {
	o.Description = v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureDetailResponseData) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubtype.Get()
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDetailResponseData) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubtype.Get(), o.EventSubtype.IsSet()
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasEventSubtype() bool {
	if o != nil && o.EventSubtype.IsSet() {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given NullableString and assigns it to the EventSubtype field.
func (o *FeatureDetailResponseData) SetEventSubtype(v string) {
	o.EventSubtype.Set(&v)
}

// SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil
func (o *FeatureDetailResponseData) SetEventSubtypeNil() {
	o.EventSubtype.Set(nil)
}

// UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
func (o *FeatureDetailResponseData) UnsetEventSubtype() {
	o.EventSubtype.Unset()
}

// GetEventSummary returns the EventSummary field value if set, zero value otherwise.
func (o *FeatureDetailResponseData) GetEventSummary() EventSummaryResponseData {
	if o == nil || IsNil(o.EventSummary) {
		var ret EventSummaryResponseData
		return ret
	}
	return *o.EventSummary
}

// GetEventSummaryOk returns a tuple with the EventSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetEventSummaryOk() (*EventSummaryResponseData, bool) {
	if o == nil || IsNil(o.EventSummary) {
		return nil, false
	}
	return o.EventSummary, true
}

// HasEventSummary returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasEventSummary() bool {
	if o != nil && !IsNil(o.EventSummary) {
		return true
	}

	return false
}

// SetEventSummary gets a reference to the given EventSummaryResponseData and assigns it to the EventSummary field.
func (o *FeatureDetailResponseData) SetEventSummary(v EventSummaryResponseData) {
	o.EventSummary = &v
}

// GetFeatureType returns the FeatureType field value
func (o *FeatureDetailResponseData) GetFeatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetFeatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureType, true
}

// SetFeatureType sets field value
func (o *FeatureDetailResponseData) SetFeatureType(v string) {
	o.FeatureType = v
}

// GetFlags returns the Flags field value
func (o *FeatureDetailResponseData) GetFlags() []FlagDetailResponseData {
	if o == nil {
		var ret []FlagDetailResponseData
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetFlagsOk() ([]FlagDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *FeatureDetailResponseData) SetFlags(v []FlagDetailResponseData) {
	o.Flags = v
}

// GetIcon returns the Icon field value
func (o *FeatureDetailResponseData) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *FeatureDetailResponseData) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value
func (o *FeatureDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FeatureDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetLifecyclePhase returns the LifecyclePhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureDetailResponseData) GetLifecyclePhase() string {
	if o == nil || IsNil(o.LifecyclePhase.Get()) {
		var ret string
		return ret
	}
	return *o.LifecyclePhase.Get()
}

// GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDetailResponseData) GetLifecyclePhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecyclePhase.Get(), o.LifecyclePhase.IsSet()
}

// HasLifecyclePhase returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasLifecyclePhase() bool {
	if o != nil && o.LifecyclePhase.IsSet() {
		return true
	}

	return false
}

// SetLifecyclePhase gets a reference to the given NullableString and assigns it to the LifecyclePhase field.
func (o *FeatureDetailResponseData) SetLifecyclePhase(v string) {
	o.LifecyclePhase.Set(&v)
}

// SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil
func (o *FeatureDetailResponseData) SetLifecyclePhaseNil() {
	o.LifecyclePhase.Set(nil)
}

// UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
func (o *FeatureDetailResponseData) UnsetLifecyclePhase() {
	o.LifecyclePhase.Unset()
}

// GetMaintainerId returns the MaintainerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureDetailResponseData) GetMaintainerId() string {
	if o == nil || IsNil(o.MaintainerId.Get()) {
		var ret string
		return ret
	}
	return *o.MaintainerId.Get()
}

// GetMaintainerIdOk returns a tuple with the MaintainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDetailResponseData) GetMaintainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintainerId.Get(), o.MaintainerId.IsSet()
}

// HasMaintainerId returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasMaintainerId() bool {
	if o != nil && o.MaintainerId.IsSet() {
		return true
	}

	return false
}

// SetMaintainerId gets a reference to the given NullableString and assigns it to the MaintainerId field.
func (o *FeatureDetailResponseData) SetMaintainerId(v string) {
	o.MaintainerId.Set(&v)
}

// SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil
func (o *FeatureDetailResponseData) SetMaintainerIdNil() {
	o.MaintainerId.Set(nil)
}

// UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
func (o *FeatureDetailResponseData) UnsetMaintainerId() {
	o.MaintainerId.Unset()
}

// GetName returns the Name field value
func (o *FeatureDetailResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FeatureDetailResponseData) SetName(v string) {
	o.Name = v
}

// GetPlans returns the Plans field value
func (o *FeatureDetailResponseData) GetPlans() []PreviewObject {
	if o == nil {
		var ret []PreviewObject
		return ret
	}

	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetPlansOk() ([]PreviewObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plans, true
}

// SetPlans sets field value
func (o *FeatureDetailResponseData) SetPlans(v []PreviewObject) {
	o.Plans = v
}

// GetTrait returns the Trait field value if set, zero value otherwise.
func (o *FeatureDetailResponseData) GetTrait() EntityTraitDefinitionResponseData {
	if o == nil || IsNil(o.Trait) {
		var ret EntityTraitDefinitionResponseData
		return ret
	}
	return *o.Trait
}

// GetTraitOk returns a tuple with the Trait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetTraitOk() (*EntityTraitDefinitionResponseData, bool) {
	if o == nil || IsNil(o.Trait) {
		return nil, false
	}
	return o.Trait, true
}

// HasTrait returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasTrait() bool {
	if o != nil && !IsNil(o.Trait) {
		return true
	}

	return false
}

// SetTrait gets a reference to the given EntityTraitDefinitionResponseData and assigns it to the Trait field.
func (o *FeatureDetailResponseData) SetTrait(v EntityTraitDefinitionResponseData) {
	o.Trait = &v
}

// GetTraitId returns the TraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureDetailResponseData) GetTraitId() string {
	if o == nil || IsNil(o.TraitId.Get()) {
		var ret string
		return ret
	}
	return *o.TraitId.Get()
}

// GetTraitIdOk returns a tuple with the TraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureDetailResponseData) GetTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitId.Get(), o.TraitId.IsSet()
}

// HasTraitId returns a boolean if a field has been set.
func (o *FeatureDetailResponseData) HasTraitId() bool {
	if o != nil && o.TraitId.IsSet() {
		return true
	}

	return false
}

// SetTraitId gets a reference to the given NullableString and assigns it to the TraitId field.
func (o *FeatureDetailResponseData) SetTraitId(v string) {
	o.TraitId.Set(&v)
}

// SetTraitIdNil sets the value for TraitId to be an explicit nil
func (o *FeatureDetailResponseData) SetTraitIdNil() {
	o.TraitId.Set(nil)
}

// UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
func (o *FeatureDetailResponseData) UnsetTraitId() {
	o.TraitId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FeatureDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FeatureDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FeatureDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o FeatureDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	if o.EventSubtype.IsSet() {
		toSerialize["event_subtype"] = o.EventSubtype.Get()
	}
	if !IsNil(o.EventSummary) {
		toSerialize["event_summary"] = o.EventSummary
	}
	toSerialize["feature_type"] = o.FeatureType
	toSerialize["flags"] = o.Flags
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	if o.LifecyclePhase.IsSet() {
		toSerialize["lifecycle_phase"] = o.LifecyclePhase.Get()
	}
	if o.MaintainerId.IsSet() {
		toSerialize["maintainer_id"] = o.MaintainerId.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["plans"] = o.Plans
	if !IsNil(o.Trait) {
		toSerialize["trait"] = o.Trait
	}
	if o.TraitId.IsSet() {
		toSerialize["trait_id"] = o.TraitId.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeatureDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"description",
		"feature_type",
		"flags",
		"icon",
		"id",
		"name",
		"plans",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFeatureDetailResponseData := _FeatureDetailResponseData{}

	err = json.Unmarshal(data, &varFeatureDetailResponseData)

	if err != nil {
		return err
	}

	*o = FeatureDetailResponseData(varFeatureDetailResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "event_subtype")
		delete(additionalProperties, "event_summary")
		delete(additionalProperties, "feature_type")
		delete(additionalProperties, "flags")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lifecycle_phase")
		delete(additionalProperties, "maintainer_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "plans")
		delete(additionalProperties, "trait")
		delete(additionalProperties, "trait_id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeatureDetailResponseData struct {
	value *FeatureDetailResponseData
	isSet bool
}

func (v NullableFeatureDetailResponseData) Get() *FeatureDetailResponseData {
	return v.value
}

func (v *NullableFeatureDetailResponseData) Set(val *FeatureDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureDetailResponseData(val *FeatureDetailResponseData) *NullableFeatureDetailResponseData {
	return &NullableFeatureDetailResponseData{value: val, isSet: true}
}

func (v NullableFeatureDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
