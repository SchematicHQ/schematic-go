/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
	"time"
)

// checks if the FeatureResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureResponseData{}

// FeatureResponseData struct for FeatureResponseData
type FeatureResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	EventSubtype NullableString `json:"event_subtype,omitempty"`
	FeatureType string `json:"feature_type"`
	Id string `json:"id"`
	LifecyclePhase NullableString `json:"lifecycle_phase,omitempty"`
	Name string `json:"name"`
	TraitId NullableString `json:"trait_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewFeatureResponseData instantiates a new FeatureResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureResponseData(createdAt time.Time, description string, featureType string, id string, name string, updatedAt time.Time) *FeatureResponseData {
	this := FeatureResponseData{}
	this.CreatedAt = createdAt
	this.Description = description
	this.FeatureType = featureType
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewFeatureResponseDataWithDefaults instantiates a new FeatureResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureResponseDataWithDefaults() *FeatureResponseData {
	this := FeatureResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FeatureResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FeatureResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *FeatureResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FeatureResponseData) SetDescription(v string) {
	o.Description = v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureResponseData) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubtype.Get()
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureResponseData) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubtype.Get(), o.EventSubtype.IsSet()
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *FeatureResponseData) HasEventSubtype() bool {
	if o != nil && o.EventSubtype.IsSet() {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given NullableString and assigns it to the EventSubtype field.
func (o *FeatureResponseData) SetEventSubtype(v string) {
	o.EventSubtype.Set(&v)
}
// SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil
func (o *FeatureResponseData) SetEventSubtypeNil() {
	o.EventSubtype.Set(nil)
}

// UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
func (o *FeatureResponseData) UnsetEventSubtype() {
	o.EventSubtype.Unset()
}

// GetFeatureType returns the FeatureType field value
func (o *FeatureResponseData) GetFeatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetFeatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureType, true
}

// SetFeatureType sets field value
func (o *FeatureResponseData) SetFeatureType(v string) {
	o.FeatureType = v
}

// GetId returns the Id field value
func (o *FeatureResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FeatureResponseData) SetId(v string) {
	o.Id = v
}

// GetLifecyclePhase returns the LifecyclePhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureResponseData) GetLifecyclePhase() string {
	if o == nil || IsNil(o.LifecyclePhase.Get()) {
		var ret string
		return ret
	}
	return *o.LifecyclePhase.Get()
}

// GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureResponseData) GetLifecyclePhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecyclePhase.Get(), o.LifecyclePhase.IsSet()
}

// HasLifecyclePhase returns a boolean if a field has been set.
func (o *FeatureResponseData) HasLifecyclePhase() bool {
	if o != nil && o.LifecyclePhase.IsSet() {
		return true
	}

	return false
}

// SetLifecyclePhase gets a reference to the given NullableString and assigns it to the LifecyclePhase field.
func (o *FeatureResponseData) SetLifecyclePhase(v string) {
	o.LifecyclePhase.Set(&v)
}
// SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil
func (o *FeatureResponseData) SetLifecyclePhaseNil() {
	o.LifecyclePhase.Set(nil)
}

// UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
func (o *FeatureResponseData) UnsetLifecyclePhase() {
	o.LifecyclePhase.Unset()
}

// GetName returns the Name field value
func (o *FeatureResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FeatureResponseData) SetName(v string) {
	o.Name = v
}

// GetTraitId returns the TraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeatureResponseData) GetTraitId() string {
	if o == nil || IsNil(o.TraitId.Get()) {
		var ret string
		return ret
	}
	return *o.TraitId.Get()
}

// GetTraitIdOk returns a tuple with the TraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeatureResponseData) GetTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitId.Get(), o.TraitId.IsSet()
}

// HasTraitId returns a boolean if a field has been set.
func (o *FeatureResponseData) HasTraitId() bool {
	if o != nil && o.TraitId.IsSet() {
		return true
	}

	return false
}

// SetTraitId gets a reference to the given NullableString and assigns it to the TraitId field.
func (o *FeatureResponseData) SetTraitId(v string) {
	o.TraitId.Set(&v)
}
// SetTraitIdNil sets the value for TraitId to be an explicit nil
func (o *FeatureResponseData) SetTraitIdNil() {
	o.TraitId.Set(nil)
}

// UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
func (o *FeatureResponseData) UnsetTraitId() {
	o.TraitId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FeatureResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FeatureResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FeatureResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o FeatureResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	if o.EventSubtype.IsSet() {
		toSerialize["event_subtype"] = o.EventSubtype.Get()
	}
	toSerialize["feature_type"] = o.FeatureType
	toSerialize["id"] = o.Id
	if o.LifecyclePhase.IsSet() {
		toSerialize["lifecycle_phase"] = o.LifecyclePhase.Get()
	}
	toSerialize["name"] = o.Name
	if o.TraitId.IsSet() {
		toSerialize["trait_id"] = o.TraitId.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableFeatureResponseData struct {
	value *FeatureResponseData
	isSet bool
}

func (v NullableFeatureResponseData) Get() *FeatureResponseData {
	return v.value
}

func (v *NullableFeatureResponseData) Set(val *FeatureResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureResponseData(val *FeatureResponseData) *NullableFeatureResponseData {
	return &NullableFeatureResponseData{value: val, isSet: true}
}

func (v NullableFeatureResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


