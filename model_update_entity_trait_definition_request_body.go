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

// checks if the UpdateEntityTraitDefinitionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEntityTraitDefinitionRequestBody{}

// UpdateEntityTraitDefinitionRequestBody struct for UpdateEntityTraitDefinitionRequestBody
type UpdateEntityTraitDefinitionRequestBody struct {
	TraitType string `json:"trait_type"`
}

// NewUpdateEntityTraitDefinitionRequestBody instantiates a new UpdateEntityTraitDefinitionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEntityTraitDefinitionRequestBody(traitType string) *UpdateEntityTraitDefinitionRequestBody {
	this := UpdateEntityTraitDefinitionRequestBody{}
	this.TraitType = traitType
	return &this
}

// NewUpdateEntityTraitDefinitionRequestBodyWithDefaults instantiates a new UpdateEntityTraitDefinitionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEntityTraitDefinitionRequestBodyWithDefaults() *UpdateEntityTraitDefinitionRequestBody {
	this := UpdateEntityTraitDefinitionRequestBody{}
	return &this
}

// GetTraitType returns the TraitType field value
func (o *UpdateEntityTraitDefinitionRequestBody) GetTraitType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value
// and a boolean to check if the value has been set.
func (o *UpdateEntityTraitDefinitionRequestBody) GetTraitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraitType, true
}

// SetTraitType sets field value
func (o *UpdateEntityTraitDefinitionRequestBody) SetTraitType(v string) {
	o.TraitType = v
}

func (o UpdateEntityTraitDefinitionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEntityTraitDefinitionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trait_type"] = o.TraitType
	return toSerialize, nil
}

type NullableUpdateEntityTraitDefinitionRequestBody struct {
	value *UpdateEntityTraitDefinitionRequestBody
	isSet bool
}

func (v NullableUpdateEntityTraitDefinitionRequestBody) Get() *UpdateEntityTraitDefinitionRequestBody {
	return v.value
}

func (v *NullableUpdateEntityTraitDefinitionRequestBody) Set(val *UpdateEntityTraitDefinitionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEntityTraitDefinitionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEntityTraitDefinitionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEntityTraitDefinitionRequestBody(val *UpdateEntityTraitDefinitionRequestBody) *NullableUpdateEntityTraitDefinitionRequestBody {
	return &NullableUpdateEntityTraitDefinitionRequestBody{value: val, isSet: true}
}

func (v NullableUpdateEntityTraitDefinitionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEntityTraitDefinitionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


