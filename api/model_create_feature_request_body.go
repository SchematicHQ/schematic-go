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
)

// checks if the CreateFeatureRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeatureRequestBody{}

// CreateFeatureRequestBody struct for CreateFeatureRequestBody
type CreateFeatureRequestBody struct {
	Description          string                         `json:"description"`
	EventSubtype         NullableString                 `json:"event_subtype,omitempty"`
	FeatureType          string                         `json:"feature_type"`
	Flag                 *CreateOrUpdateFlagRequestBody `json:"flag,omitempty"`
	LifecyclePhase       NullableString                 `json:"lifecycle_phase,omitempty"`
	Name                 string                         `json:"name"`
	TraitId              NullableString                 `json:"trait_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFeatureRequestBody CreateFeatureRequestBody

// NewCreateFeatureRequestBody instantiates a new CreateFeatureRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeatureRequestBody(description string, featureType string, name string) *CreateFeatureRequestBody {
	this := CreateFeatureRequestBody{}
	this.Description = description
	this.FeatureType = featureType
	this.Name = name
	return &this
}

// NewCreateFeatureRequestBodyWithDefaults instantiates a new CreateFeatureRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeatureRequestBodyWithDefaults() *CreateFeatureRequestBody {
	this := CreateFeatureRequestBody{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateFeatureRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateFeatureRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureRequestBody) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubtype.Get()
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureRequestBody) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubtype.Get(), o.EventSubtype.IsSet()
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *CreateFeatureRequestBody) HasEventSubtype() bool {
	if o != nil && o.EventSubtype.IsSet() {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given NullableString and assigns it to the EventSubtype field.
func (o *CreateFeatureRequestBody) SetEventSubtype(v string) {
	o.EventSubtype.Set(&v)
}

// SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil
func (o *CreateFeatureRequestBody) SetEventSubtypeNil() {
	o.EventSubtype.Set(nil)
}

// UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
func (o *CreateFeatureRequestBody) UnsetEventSubtype() {
	o.EventSubtype.Unset()
}

// GetFeatureType returns the FeatureType field value
func (o *CreateFeatureRequestBody) GetFeatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureRequestBody) GetFeatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureType, true
}

// SetFeatureType sets field value
func (o *CreateFeatureRequestBody) SetFeatureType(v string) {
	o.FeatureType = v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *CreateFeatureRequestBody) GetFlag() CreateOrUpdateFlagRequestBody {
	if o == nil || IsNil(o.Flag) {
		var ret CreateOrUpdateFlagRequestBody
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureRequestBody) GetFlagOk() (*CreateOrUpdateFlagRequestBody, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *CreateFeatureRequestBody) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given CreateOrUpdateFlagRequestBody and assigns it to the Flag field.
func (o *CreateFeatureRequestBody) SetFlag(v CreateOrUpdateFlagRequestBody) {
	o.Flag = &v
}

// GetLifecyclePhase returns the LifecyclePhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureRequestBody) GetLifecyclePhase() string {
	if o == nil || IsNil(o.LifecyclePhase.Get()) {
		var ret string
		return ret
	}
	return *o.LifecyclePhase.Get()
}

// GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureRequestBody) GetLifecyclePhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecyclePhase.Get(), o.LifecyclePhase.IsSet()
}

// HasLifecyclePhase returns a boolean if a field has been set.
func (o *CreateFeatureRequestBody) HasLifecyclePhase() bool {
	if o != nil && o.LifecyclePhase.IsSet() {
		return true
	}

	return false
}

// SetLifecyclePhase gets a reference to the given NullableString and assigns it to the LifecyclePhase field.
func (o *CreateFeatureRequestBody) SetLifecyclePhase(v string) {
	o.LifecyclePhase.Set(&v)
}

// SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil
func (o *CreateFeatureRequestBody) SetLifecyclePhaseNil() {
	o.LifecyclePhase.Set(nil)
}

// UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
func (o *CreateFeatureRequestBody) UnsetLifecyclePhase() {
	o.LifecyclePhase.Unset()
}

// GetName returns the Name field value
func (o *CreateFeatureRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFeatureRequestBody) SetName(v string) {
	o.Name = v
}

// GetTraitId returns the TraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureRequestBody) GetTraitId() string {
	if o == nil || IsNil(o.TraitId.Get()) {
		var ret string
		return ret
	}
	return *o.TraitId.Get()
}

// GetTraitIdOk returns a tuple with the TraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureRequestBody) GetTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitId.Get(), o.TraitId.IsSet()
}

// HasTraitId returns a boolean if a field has been set.
func (o *CreateFeatureRequestBody) HasTraitId() bool {
	if o != nil && o.TraitId.IsSet() {
		return true
	}

	return false
}

// SetTraitId gets a reference to the given NullableString and assigns it to the TraitId field.
func (o *CreateFeatureRequestBody) SetTraitId(v string) {
	o.TraitId.Set(&v)
}

// SetTraitIdNil sets the value for TraitId to be an explicit nil
func (o *CreateFeatureRequestBody) SetTraitIdNil() {
	o.TraitId.Set(nil)
}

// UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
func (o *CreateFeatureRequestBody) UnsetTraitId() {
	o.TraitId.Unset()
}

func (o CreateFeatureRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFeatureRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if o.EventSubtype.IsSet() {
		toSerialize["event_subtype"] = o.EventSubtype.Get()
	}
	toSerialize["feature_type"] = o.FeatureType
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if o.LifecyclePhase.IsSet() {
		toSerialize["lifecycle_phase"] = o.LifecyclePhase.Get()
	}
	toSerialize["name"] = o.Name
	if o.TraitId.IsSet() {
		toSerialize["trait_id"] = o.TraitId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFeatureRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"feature_type",
		"name",
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

	varCreateFeatureRequestBody := _CreateFeatureRequestBody{}

	err = json.Unmarshal(data, &varCreateFeatureRequestBody)

	if err != nil {
		return err
	}

	*o = CreateFeatureRequestBody(varCreateFeatureRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "event_subtype")
		delete(additionalProperties, "feature_type")
		delete(additionalProperties, "flag")
		delete(additionalProperties, "lifecycle_phase")
		delete(additionalProperties, "name")
		delete(additionalProperties, "trait_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFeatureRequestBody struct {
	value *CreateFeatureRequestBody
	isSet bool
}

func (v NullableCreateFeatureRequestBody) Get() *CreateFeatureRequestBody {
	return v.value
}

func (v *NullableCreateFeatureRequestBody) Set(val *CreateFeatureRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeatureRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeatureRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeatureRequestBody(val *CreateFeatureRequestBody) *NullableCreateFeatureRequestBody {
	return &NullableCreateFeatureRequestBody{value: val, isSet: true}
}

func (v NullableCreateFeatureRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFeatureRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
