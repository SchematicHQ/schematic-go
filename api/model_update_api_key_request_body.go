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

// checks if the UpdateApiKeyRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApiKeyRequestBody{}

// UpdateApiKeyRequestBody struct for UpdateApiKeyRequestBody
type UpdateApiKeyRequestBody struct {
	Description NullableString `json:"description,omitempty"`
	Name        NullableString `json:"name,omitempty"`
}

// NewUpdateApiKeyRequestBody instantiates a new UpdateApiKeyRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiKeyRequestBody() *UpdateApiKeyRequestBody {
	this := UpdateApiKeyRequestBody{}
	return &this
}

// NewUpdateApiKeyRequestBodyWithDefaults instantiates a new UpdateApiKeyRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiKeyRequestBodyWithDefaults() *UpdateApiKeyRequestBody {
	this := UpdateApiKeyRequestBody{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApiKeyRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApiKeyRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateApiKeyRequestBody) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateApiKeyRequestBody) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateApiKeyRequestBody) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateApiKeyRequestBody) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateApiKeyRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateApiKeyRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateApiKeyRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateApiKeyRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateApiKeyRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateApiKeyRequestBody) UnsetName() {
	o.Name.Unset()
}

func (o UpdateApiKeyRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApiKeyRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableUpdateApiKeyRequestBody struct {
	value *UpdateApiKeyRequestBody
	isSet bool
}

func (v NullableUpdateApiKeyRequestBody) Get() *UpdateApiKeyRequestBody {
	return v.value
}

func (v *NullableUpdateApiKeyRequestBody) Set(val *UpdateApiKeyRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiKeyRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiKeyRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiKeyRequestBody(val *UpdateApiKeyRequestBody) *NullableUpdateApiKeyRequestBody {
	return &NullableUpdateApiKeyRequestBody{value: val, isSet: true}
}

func (v NullableUpdateApiKeyRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiKeyRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
