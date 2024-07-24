/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the CheckFlagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckFlagsResponse{}

// CheckFlagsResponse struct for CheckFlagsResponse
type CheckFlagsResponse struct {
	Data CheckFlagsResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _CheckFlagsResponse CheckFlagsResponse

// NewCheckFlagsResponse instantiates a new CheckFlagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckFlagsResponse(data CheckFlagsResponseData, params map[string]interface{}) *CheckFlagsResponse {
	this := CheckFlagsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCheckFlagsResponseWithDefaults instantiates a new CheckFlagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckFlagsResponseWithDefaults() *CheckFlagsResponse {
	this := CheckFlagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CheckFlagsResponse) GetData() CheckFlagsResponseData {
	if o == nil {
		var ret CheckFlagsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckFlagsResponse) GetDataOk() (*CheckFlagsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckFlagsResponse) SetData(v CheckFlagsResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CheckFlagsResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CheckFlagsResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *CheckFlagsResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o CheckFlagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckFlagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckFlagsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCheckFlagsResponse := _CheckFlagsResponse{}

	err = json.Unmarshal(data, &varCheckFlagsResponse)

	if err != nil {
		return err
	}

	*o = CheckFlagsResponse(varCheckFlagsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckFlagsResponse struct {
	value *CheckFlagsResponse
	isSet bool
}

func (v NullableCheckFlagsResponse) Get() *CheckFlagsResponse {
	return v.value
}

func (v *NullableCheckFlagsResponse) Set(val *CheckFlagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckFlagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckFlagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckFlagsResponse(val *CheckFlagsResponse) *NullableCheckFlagsResponse {
	return &NullableCheckFlagsResponse{value: val, isSet: true}
}

func (v NullableCheckFlagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckFlagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
