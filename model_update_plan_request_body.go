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

// checks if the UpdatePlanRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePlanRequestBody{}

// UpdatePlanRequestBody struct for UpdatePlanRequestBody
type UpdatePlanRequestBody struct {
	Name string `json:"name"`
}

// NewUpdatePlanRequestBody instantiates a new UpdatePlanRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlanRequestBody(name string) *UpdatePlanRequestBody {
	this := UpdatePlanRequestBody{}
	this.Name = name
	return &this
}

// NewUpdatePlanRequestBodyWithDefaults instantiates a new UpdatePlanRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlanRequestBodyWithDefaults() *UpdatePlanRequestBody {
	this := UpdatePlanRequestBody{}
	return &this
}

// GetName returns the Name field value
func (o *UpdatePlanRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdatePlanRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdatePlanRequestBody) SetName(v string) {
	o.Name = v
}

func (o UpdatePlanRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePlanRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableUpdatePlanRequestBody struct {
	value *UpdatePlanRequestBody
	isSet bool
}

func (v NullableUpdatePlanRequestBody) Get() *UpdatePlanRequestBody {
	return v.value
}

func (v *NullableUpdatePlanRequestBody) Set(val *UpdatePlanRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlanRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlanRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlanRequestBody(val *UpdatePlanRequestBody) *NullableUpdatePlanRequestBody {
	return &NullableUpdatePlanRequestBody{value: val, isSet: true}
}

func (v NullableUpdatePlanRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePlanRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


