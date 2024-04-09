/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateEntityTraitDefinitionRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEntityTraitDefinitionRequestBody{}

// CreateEntityTraitDefinitionRequestBody struct for CreateEntityTraitDefinitionRequestBody
type CreateEntityTraitDefinitionRequestBody struct {
	EntityType string   `json:"entity_type"`
	Hierarchy  []string `json:"hierarchy"`
	TraitType  string   `json:"trait_type"`
}

type _CreateEntityTraitDefinitionRequestBody CreateEntityTraitDefinitionRequestBody

// NewCreateEntityTraitDefinitionRequestBody instantiates a new CreateEntityTraitDefinitionRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEntityTraitDefinitionRequestBody(entityType string, hierarchy []string, traitType string) *CreateEntityTraitDefinitionRequestBody {
	this := CreateEntityTraitDefinitionRequestBody{}
	this.EntityType = entityType
	this.Hierarchy = hierarchy
	this.TraitType = traitType
	return &this
}

// NewCreateEntityTraitDefinitionRequestBodyWithDefaults instantiates a new CreateEntityTraitDefinitionRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEntityTraitDefinitionRequestBodyWithDefaults() *CreateEntityTraitDefinitionRequestBody {
	this := CreateEntityTraitDefinitionRequestBody{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *CreateEntityTraitDefinitionRequestBody) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *CreateEntityTraitDefinitionRequestBody) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *CreateEntityTraitDefinitionRequestBody) SetEntityType(v string) {
	o.EntityType = v
}

// GetHierarchy returns the Hierarchy field value
func (o *CreateEntityTraitDefinitionRequestBody) GetHierarchy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hierarchy
}

// GetHierarchyOk returns a tuple with the Hierarchy field value
// and a boolean to check if the value has been set.
func (o *CreateEntityTraitDefinitionRequestBody) GetHierarchyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hierarchy, true
}

// SetHierarchy sets field value
func (o *CreateEntityTraitDefinitionRequestBody) SetHierarchy(v []string) {
	o.Hierarchy = v
}

// GetTraitType returns the TraitType field value
func (o *CreateEntityTraitDefinitionRequestBody) GetTraitType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value
// and a boolean to check if the value has been set.
func (o *CreateEntityTraitDefinitionRequestBody) GetTraitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraitType, true
}

// SetTraitType sets field value
func (o *CreateEntityTraitDefinitionRequestBody) SetTraitType(v string) {
	o.TraitType = v
}

func (o CreateEntityTraitDefinitionRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEntityTraitDefinitionRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	toSerialize["hierarchy"] = o.Hierarchy
	toSerialize["trait_type"] = o.TraitType
	return toSerialize, nil
}

func (o *CreateEntityTraitDefinitionRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_type",
		"hierarchy",
		"trait_type",
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

	varCreateEntityTraitDefinitionRequestBody := _CreateEntityTraitDefinitionRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEntityTraitDefinitionRequestBody)

	if err != nil {
		return err
	}

	*o = CreateEntityTraitDefinitionRequestBody(varCreateEntityTraitDefinitionRequestBody)

	return err
}

type NullableCreateEntityTraitDefinitionRequestBody struct {
	value *CreateEntityTraitDefinitionRequestBody
	isSet bool
}

func (v NullableCreateEntityTraitDefinitionRequestBody) Get() *CreateEntityTraitDefinitionRequestBody {
	return v.value
}

func (v *NullableCreateEntityTraitDefinitionRequestBody) Set(val *CreateEntityTraitDefinitionRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEntityTraitDefinitionRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEntityTraitDefinitionRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEntityTraitDefinitionRequestBody(val *CreateEntityTraitDefinitionRequestBody) *NullableCreateEntityTraitDefinitionRequestBody {
	return &NullableCreateEntityTraitDefinitionRequestBody{value: val, isSet: true}
}

func (v NullableCreateEntityTraitDefinitionRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEntityTraitDefinitionRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
