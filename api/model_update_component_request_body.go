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

// checks if the UpdateComponentRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComponentRequestBody{}

// UpdateComponentRequestBody struct for UpdateComponentRequestBody
type UpdateComponentRequestBody struct {
	Ast        []int32        `json:"ast,omitempty"`
	EntityType NullableString `json:"entity_type,omitempty"`
	Name       NullableString `json:"name,omitempty"`
	State      NullableString `json:"state,omitempty"`
}

// NewUpdateComponentRequestBody instantiates a new UpdateComponentRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComponentRequestBody() *UpdateComponentRequestBody {
	this := UpdateComponentRequestBody{}
	return &this
}

// NewUpdateComponentRequestBodyWithDefaults instantiates a new UpdateComponentRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComponentRequestBodyWithDefaults() *UpdateComponentRequestBody {
	this := UpdateComponentRequestBody{}
	return &this
}

// GetAst returns the Ast field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateComponentRequestBody) GetAst() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Ast
}

// GetAstOk returns a tuple with the Ast field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateComponentRequestBody) GetAstOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ast) {
		return nil, false
	}
	return o.Ast, true
}

// HasAst returns a boolean if a field has been set.
func (o *UpdateComponentRequestBody) HasAst() bool {
	if o != nil && !IsNil(o.Ast) {
		return true
	}

	return false
}

// SetAst gets a reference to the given []int32 and assigns it to the Ast field.
func (o *UpdateComponentRequestBody) SetAst(v []int32) {
	o.Ast = v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateComponentRequestBody) GetEntityType() string {
	if o == nil || IsNil(o.EntityType.Get()) {
		var ret string
		return ret
	}
	return *o.EntityType.Get()
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateComponentRequestBody) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntityType.Get(), o.EntityType.IsSet()
}

// HasEntityType returns a boolean if a field has been set.
func (o *UpdateComponentRequestBody) HasEntityType() bool {
	if o != nil && o.EntityType.IsSet() {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given NullableString and assigns it to the EntityType field.
func (o *UpdateComponentRequestBody) SetEntityType(v string) {
	o.EntityType.Set(&v)
}

// SetEntityTypeNil sets the value for EntityType to be an explicit nil
func (o *UpdateComponentRequestBody) SetEntityTypeNil() {
	o.EntityType.Set(nil)
}

// UnsetEntityType ensures that no value is present for EntityType, not even an explicit nil
func (o *UpdateComponentRequestBody) UnsetEntityType() {
	o.EntityType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateComponentRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateComponentRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateComponentRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateComponentRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateComponentRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateComponentRequestBody) UnsetName() {
	o.Name.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateComponentRequestBody) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateComponentRequestBody) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *UpdateComponentRequestBody) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *UpdateComponentRequestBody) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *UpdateComponentRequestBody) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *UpdateComponentRequestBody) UnsetState() {
	o.State.Unset()
}

func (o UpdateComponentRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComponentRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ast != nil {
		toSerialize["ast"] = o.Ast
	}
	if o.EntityType.IsSet() {
		toSerialize["entity_type"] = o.EntityType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	return toSerialize, nil
}

type NullableUpdateComponentRequestBody struct {
	value *UpdateComponentRequestBody
	isSet bool
}

func (v NullableUpdateComponentRequestBody) Get() *UpdateComponentRequestBody {
	return v.value
}

func (v *NullableUpdateComponentRequestBody) Set(val *UpdateComponentRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComponentRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComponentRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComponentRequestBody(val *UpdateComponentRequestBody) *NullableUpdateComponentRequestBody {
	return &NullableUpdateComponentRequestBody{value: val, isSet: true}
}

func (v NullableUpdateComponentRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComponentRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
