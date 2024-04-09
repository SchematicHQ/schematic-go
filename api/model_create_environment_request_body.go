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

// checks if the CreateEnvironmentRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEnvironmentRequestBody{}

// CreateEnvironmentRequestBody struct for CreateEnvironmentRequestBody
type CreateEnvironmentRequestBody struct {
	EnvironmentType string `json:"environment_type"`
	Name            string `json:"name"`
}

type _CreateEnvironmentRequestBody CreateEnvironmentRequestBody

// NewCreateEnvironmentRequestBody instantiates a new CreateEnvironmentRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentRequestBody(environmentType string, name string) *CreateEnvironmentRequestBody {
	this := CreateEnvironmentRequestBody{}
	this.EnvironmentType = environmentType
	this.Name = name
	return &this
}

// NewCreateEnvironmentRequestBodyWithDefaults instantiates a new CreateEnvironmentRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentRequestBodyWithDefaults() *CreateEnvironmentRequestBody {
	this := CreateEnvironmentRequestBody{}
	return &this
}

// GetEnvironmentType returns the EnvironmentType field value
func (o *CreateEnvironmentRequestBody) GetEnvironmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentRequestBody) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentType, true
}

// SetEnvironmentType sets field value
func (o *CreateEnvironmentRequestBody) SetEnvironmentType(v string) {
	o.EnvironmentType = v
}

// GetName returns the Name field value
func (o *CreateEnvironmentRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateEnvironmentRequestBody) SetName(v string) {
	o.Name = v
}

func (o CreateEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEnvironmentRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_type"] = o.EnvironmentType
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CreateEnvironmentRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment_type",
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

	varCreateEnvironmentRequestBody := _CreateEnvironmentRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEnvironmentRequestBody)

	if err != nil {
		return err
	}

	*o = CreateEnvironmentRequestBody(varCreateEnvironmentRequestBody)

	return err
}

type NullableCreateEnvironmentRequestBody struct {
	value *CreateEnvironmentRequestBody
	isSet bool
}

func (v NullableCreateEnvironmentRequestBody) Get() *CreateEnvironmentRequestBody {
	return v.value
}

func (v *NullableCreateEnvironmentRequestBody) Set(val *CreateEnvironmentRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentRequestBody(val *CreateEnvironmentRequestBody) *NullableCreateEnvironmentRequestBody {
	return &NullableCreateEnvironmentRequestBody{value: val, isSet: true}
}

func (v NullableCreateEnvironmentRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
