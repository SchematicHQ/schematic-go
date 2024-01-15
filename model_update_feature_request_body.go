/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
)

// checks if the UpdateFeatureRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFeatureRequestBody{}

// UpdateFeatureRequestBody struct for UpdateFeatureRequestBody
type UpdateFeatureRequestBody struct {
	Description NullableString `json:"description,omitempty"`
	EventSubtype NullableString `json:"event_subtype,omitempty"`
	FeatureType NullableString `json:"feature_type,omitempty"`
	Flag *CreateOrUpdateFlagRequestBody `json:"flag,omitempty"`
	LifecyclePhase NullableString `json:"lifecycle_phase,omitempty"`
	Name NullableString `json:"name,omitempty"`
	SkipWebhooks NullableBool `json:"skip_webhooks,omitempty"`
	TraitId NullableString `json:"trait_id,omitempty"`
}

// NewUpdateFeatureRequestBody instantiates a new UpdateFeatureRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFeatureRequestBody() *UpdateFeatureRequestBody {
	this := UpdateFeatureRequestBody{}
	return &this
}

// NewUpdateFeatureRequestBodyWithDefaults instantiates a new UpdateFeatureRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFeatureRequestBodyWithDefaults() *UpdateFeatureRequestBody {
	this := UpdateFeatureRequestBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateFeatureRequestBody) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateFeatureRequestBody) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetDescription() {
	o.Description.Unset()
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubtype.Get()
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubtype.Get(), o.EventSubtype.IsSet()
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasEventSubtype() bool {
	if o != nil && o.EventSubtype.IsSet() {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given NullableString and assigns it to the EventSubtype field.
func (o *UpdateFeatureRequestBody) SetEventSubtype(v string) {
	o.EventSubtype.Set(&v)
}
// SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil
func (o *UpdateFeatureRequestBody) SetEventSubtypeNil() {
	o.EventSubtype.Set(nil)
}

// UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetEventSubtype() {
	o.EventSubtype.Unset()
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetFeatureType() string {
	if o == nil || IsNil(o.FeatureType.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureType.Get()
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetFeatureTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureType.Get(), o.FeatureType.IsSet()
}

// HasFeatureType returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasFeatureType() bool {
	if o != nil && o.FeatureType.IsSet() {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given NullableString and assigns it to the FeatureType field.
func (o *UpdateFeatureRequestBody) SetFeatureType(v string) {
	o.FeatureType.Set(&v)
}
// SetFeatureTypeNil sets the value for FeatureType to be an explicit nil
func (o *UpdateFeatureRequestBody) SetFeatureTypeNil() {
	o.FeatureType.Set(nil)
}

// UnsetFeatureType ensures that no value is present for FeatureType, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetFeatureType() {
	o.FeatureType.Unset()
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *UpdateFeatureRequestBody) GetFlag() CreateOrUpdateFlagRequestBody {
	if o == nil || IsNil(o.Flag) {
		var ret CreateOrUpdateFlagRequestBody
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureRequestBody) GetFlagOk() (*CreateOrUpdateFlagRequestBody, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given CreateOrUpdateFlagRequestBody and assigns it to the Flag field.
func (o *UpdateFeatureRequestBody) SetFlag(v CreateOrUpdateFlagRequestBody) {
	o.Flag = &v
}

// GetLifecyclePhase returns the LifecyclePhase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetLifecyclePhase() string {
	if o == nil || IsNil(o.LifecyclePhase.Get()) {
		var ret string
		return ret
	}
	return *o.LifecyclePhase.Get()
}

// GetLifecyclePhaseOk returns a tuple with the LifecyclePhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetLifecyclePhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecyclePhase.Get(), o.LifecyclePhase.IsSet()
}

// HasLifecyclePhase returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasLifecyclePhase() bool {
	if o != nil && o.LifecyclePhase.IsSet() {
		return true
	}

	return false
}

// SetLifecyclePhase gets a reference to the given NullableString and assigns it to the LifecyclePhase field.
func (o *UpdateFeatureRequestBody) SetLifecyclePhase(v string) {
	o.LifecyclePhase.Set(&v)
}
// SetLifecyclePhaseNil sets the value for LifecyclePhase to be an explicit nil
func (o *UpdateFeatureRequestBody) SetLifecyclePhaseNil() {
	o.LifecyclePhase.Set(nil)
}

// UnsetLifecyclePhase ensures that no value is present for LifecyclePhase, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetLifecyclePhase() {
	o.LifecyclePhase.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateFeatureRequestBody) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateFeatureRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetName() {
	o.Name.Unset()
}

// GetSkipWebhooks returns the SkipWebhooks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetSkipWebhooks() bool {
	if o == nil || IsNil(o.SkipWebhooks.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipWebhooks.Get()
}

// GetSkipWebhooksOk returns a tuple with the SkipWebhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetSkipWebhooksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipWebhooks.Get(), o.SkipWebhooks.IsSet()
}

// HasSkipWebhooks returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasSkipWebhooks() bool {
	if o != nil && o.SkipWebhooks.IsSet() {
		return true
	}

	return false
}

// SetSkipWebhooks gets a reference to the given NullableBool and assigns it to the SkipWebhooks field.
func (o *UpdateFeatureRequestBody) SetSkipWebhooks(v bool) {
	o.SkipWebhooks.Set(&v)
}
// SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil
func (o *UpdateFeatureRequestBody) SetSkipWebhooksNil() {
	o.SkipWebhooks.Set(nil)
}

// UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetSkipWebhooks() {
	o.SkipWebhooks.Unset()
}

// GetTraitId returns the TraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureRequestBody) GetTraitId() string {
	if o == nil || IsNil(o.TraitId.Get()) {
		var ret string
		return ret
	}
	return *o.TraitId.Get()
}

// GetTraitIdOk returns a tuple with the TraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureRequestBody) GetTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitId.Get(), o.TraitId.IsSet()
}

// HasTraitId returns a boolean if a field has been set.
func (o *UpdateFeatureRequestBody) HasTraitId() bool {
	if o != nil && o.TraitId.IsSet() {
		return true
	}

	return false
}

// SetTraitId gets a reference to the given NullableString and assigns it to the TraitId field.
func (o *UpdateFeatureRequestBody) SetTraitId(v string) {
	o.TraitId.Set(&v)
}
// SetTraitIdNil sets the value for TraitId to be an explicit nil
func (o *UpdateFeatureRequestBody) SetTraitIdNil() {
	o.TraitId.Set(nil)
}

// UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
func (o *UpdateFeatureRequestBody) UnsetTraitId() {
	o.TraitId.Unset()
}

func (o UpdateFeatureRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFeatureRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EventSubtype.IsSet() {
		toSerialize["event_subtype"] = o.EventSubtype.Get()
	}
	if o.FeatureType.IsSet() {
		toSerialize["feature_type"] = o.FeatureType.Get()
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if o.LifecyclePhase.IsSet() {
		toSerialize["lifecycle_phase"] = o.LifecyclePhase.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SkipWebhooks.IsSet() {
		toSerialize["skip_webhooks"] = o.SkipWebhooks.Get()
	}
	if o.TraitId.IsSet() {
		toSerialize["trait_id"] = o.TraitId.Get()
	}
	return toSerialize, nil
}

type NullableUpdateFeatureRequestBody struct {
	value *UpdateFeatureRequestBody
	isSet bool
}

func (v NullableUpdateFeatureRequestBody) Get() *UpdateFeatureRequestBody {
	return v.value
}

func (v *NullableUpdateFeatureRequestBody) Set(val *UpdateFeatureRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFeatureRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFeatureRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFeatureRequestBody(val *UpdateFeatureRequestBody) *NullableUpdateFeatureRequestBody {
	return &NullableUpdateFeatureRequestBody{value: val, isSet: true}
}

func (v NullableUpdateFeatureRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFeatureRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


