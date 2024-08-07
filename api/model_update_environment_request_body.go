/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateEnvironmentRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEnvironmentRequestBody{}

// UpdateEnvironmentRequestBody struct for UpdateEnvironmentRequestBody
type UpdateEnvironmentRequestBody struct {
	EnvironmentType      NullableString `json:"environment_type,omitempty"`
	Name                 NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateEnvironmentRequestBody UpdateEnvironmentRequestBody

// NewUpdateEnvironmentRequestBody instantiates a new UpdateEnvironmentRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentRequestBody() *UpdateEnvironmentRequestBody {
	this := UpdateEnvironmentRequestBody{}
	return &this
}

// NewUpdateEnvironmentRequestBodyWithDefaults instantiates a new UpdateEnvironmentRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentRequestBodyWithDefaults() *UpdateEnvironmentRequestBody {
	this := UpdateEnvironmentRequestBody{}
	return &this
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateEnvironmentRequestBody) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentType.Get()
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateEnvironmentRequestBody) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentType.Get(), o.EnvironmentType.IsSet()
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *UpdateEnvironmentRequestBody) HasEnvironmentType() bool {
	if o != nil && o.EnvironmentType.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given NullableString and assigns it to the EnvironmentType field.
func (o *UpdateEnvironmentRequestBody) SetEnvironmentType(v string) {
	o.EnvironmentType.Set(&v)
}

// SetEnvironmentTypeNil sets the value for EnvironmentType to be an explicit nil
func (o *UpdateEnvironmentRequestBody) SetEnvironmentTypeNil() {
	o.EnvironmentType.Set(nil)
}

// UnsetEnvironmentType ensures that no value is present for EnvironmentType, not even an explicit nil
func (o *UpdateEnvironmentRequestBody) UnsetEnvironmentType() {
	o.EnvironmentType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateEnvironmentRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateEnvironmentRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateEnvironmentRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateEnvironmentRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateEnvironmentRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateEnvironmentRequestBody) UnsetName() {
	o.Name.Unset()
}

func (o UpdateEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEnvironmentRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnvironmentType.IsSet() {
		toSerialize["environment_type"] = o.EnvironmentType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateEnvironmentRequestBody) UnmarshalJSON(data []byte) (err error) {
	varUpdateEnvironmentRequestBody := _UpdateEnvironmentRequestBody{}

	err = json.Unmarshal(data, &varUpdateEnvironmentRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateEnvironmentRequestBody(varUpdateEnvironmentRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment_type")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateEnvironmentRequestBody struct {
	value *UpdateEnvironmentRequestBody
	isSet bool
}

func (v NullableUpdateEnvironmentRequestBody) Get() *UpdateEnvironmentRequestBody {
	return v.value
}

func (v *NullableUpdateEnvironmentRequestBody) Set(val *UpdateEnvironmentRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentRequestBody(val *UpdateEnvironmentRequestBody) *NullableUpdateEnvironmentRequestBody {
	return &NullableUpdateEnvironmentRequestBody{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
