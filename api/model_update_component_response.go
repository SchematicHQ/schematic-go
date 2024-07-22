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

// checks if the UpdateComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComponentResponse{}

// UpdateComponentResponse struct for UpdateComponentResponse
type UpdateComponentResponse struct {
	Data ComponentResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _UpdateComponentResponse UpdateComponentResponse

// NewUpdateComponentResponse instantiates a new UpdateComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComponentResponse(data ComponentResponseData, params map[string]interface{}) *UpdateComponentResponse {
	this := UpdateComponentResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdateComponentResponseWithDefaults instantiates a new UpdateComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComponentResponseWithDefaults() *UpdateComponentResponse {
	this := UpdateComponentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateComponentResponse) GetData() ComponentResponseData {
	if o == nil {
		var ret ComponentResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateComponentResponse) GetDataOk() (*ComponentResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateComponentResponse) SetData(v ComponentResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdateComponentResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateComponentResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdateComponentResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdateComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *UpdateComponentResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateComponentResponse := _UpdateComponentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateComponentResponse)

	if err != nil {
		return err
	}

	*o = UpdateComponentResponse(varUpdateComponentResponse)

	return err
}

type NullableUpdateComponentResponse struct {
	value *UpdateComponentResponse
	isSet bool
}

func (v NullableUpdateComponentResponse) Get() *UpdateComponentResponse {
	return v.value
}

func (v *NullableUpdateComponentResponse) Set(val *UpdateComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComponentResponse(val *UpdateComponentResponse) *NullableUpdateComponentResponse {
	return &NullableUpdateComponentResponse{value: val, isSet: true}
}

func (v NullableUpdateComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
