/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KeysRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeysRequestBody{}

// KeysRequestBody struct for KeysRequestBody
type KeysRequestBody struct {
	Keys map[string]interface{} `json:"keys"`
}

type _KeysRequestBody KeysRequestBody

// NewKeysRequestBody instantiates a new KeysRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeysRequestBody(keys map[string]interface{}) *KeysRequestBody {
	this := KeysRequestBody{}
	this.Keys = keys
	return &this
}

// NewKeysRequestBodyWithDefaults instantiates a new KeysRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeysRequestBodyWithDefaults() *KeysRequestBody {
	this := KeysRequestBody{}
	return &this
}

// GetKeys returns the Keys field value
func (o *KeysRequestBody) GetKeys() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *KeysRequestBody) GetKeysOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Keys, true
}

// SetKeys sets field value
func (o *KeysRequestBody) SetKeys(v map[string]interface{}) {
	o.Keys = v
}

func (o KeysRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeysRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keys"] = o.Keys
	return toSerialize, nil
}

func (o *KeysRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keys",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKeysRequestBody := _KeysRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeysRequestBody)

	if err != nil {
		return err
	}

	*o = KeysRequestBody(varKeysRequestBody)

	return err
}

type NullableKeysRequestBody struct {
	value *KeysRequestBody
	isSet bool
}

func (v NullableKeysRequestBody) Get() *KeysRequestBody {
	return v.value
}

func (v *NullableKeysRequestBody) Set(val *KeysRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableKeysRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableKeysRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeysRequestBody(val *KeysRequestBody) *NullableKeysRequestBody {
	return &NullableKeysRequestBody{value: val, isSet: true}
}

func (v NullableKeysRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeysRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


