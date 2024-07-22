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

// checks if the HydrateComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HydrateComponentResponse{}

// HydrateComponentResponse struct for HydrateComponentResponse
type HydrateComponentResponse struct {
	Data ComponentHydrateResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _HydrateComponentResponse HydrateComponentResponse

// NewHydrateComponentResponse instantiates a new HydrateComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHydrateComponentResponse(data ComponentHydrateResponseData, params map[string]interface{}) *HydrateComponentResponse {
	this := HydrateComponentResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewHydrateComponentResponseWithDefaults instantiates a new HydrateComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHydrateComponentResponseWithDefaults() *HydrateComponentResponse {
	this := HydrateComponentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *HydrateComponentResponse) GetData() ComponentHydrateResponseData {
	if o == nil {
		var ret ComponentHydrateResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *HydrateComponentResponse) GetDataOk() (*ComponentHydrateResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *HydrateComponentResponse) SetData(v ComponentHydrateResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *HydrateComponentResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *HydrateComponentResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *HydrateComponentResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o HydrateComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HydrateComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *HydrateComponentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"params",
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

	varHydrateComponentResponse := _HydrateComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHydrateComponentResponse)

	if err != nil {
		return err
	}

	*o = HydrateComponentResponse(varHydrateComponentResponse)

	return err
}

type NullableHydrateComponentResponse struct {
	value *HydrateComponentResponse
	isSet bool
}

func (v NullableHydrateComponentResponse) Get() *HydrateComponentResponse {
	return v.value
}

func (v *NullableHydrateComponentResponse) Set(val *HydrateComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHydrateComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHydrateComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHydrateComponentResponse(val *HydrateComponentResponse) *NullableHydrateComponentResponse {
	return &NullableHydrateComponentResponse{value: val, isSet: true}
}

func (v NullableHydrateComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHydrateComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
