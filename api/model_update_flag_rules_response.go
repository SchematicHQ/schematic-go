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

// checks if the UpdateFlagRulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFlagRulesResponse{}

// UpdateFlagRulesResponse struct for UpdateFlagRulesResponse
type UpdateFlagRulesResponse struct {
	Data RulesDetailResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFlagRulesResponse UpdateFlagRulesResponse

// NewUpdateFlagRulesResponse instantiates a new UpdateFlagRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFlagRulesResponse(data RulesDetailResponseData, params map[string]interface{}) *UpdateFlagRulesResponse {
	this := UpdateFlagRulesResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdateFlagRulesResponseWithDefaults instantiates a new UpdateFlagRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFlagRulesResponseWithDefaults() *UpdateFlagRulesResponse {
	this := UpdateFlagRulesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateFlagRulesResponse) GetData() RulesDetailResponseData {
	if o == nil {
		var ret RulesDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateFlagRulesResponse) GetDataOk() (*RulesDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateFlagRulesResponse) SetData(v RulesDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdateFlagRulesResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateFlagRulesResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdateFlagRulesResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdateFlagRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFlagRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFlagRulesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateFlagRulesResponse := _UpdateFlagRulesResponse{}

	err = json.Unmarshal(data, &varUpdateFlagRulesResponse)

	if err != nil {
		return err
	}

	*o = UpdateFlagRulesResponse(varUpdateFlagRulesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFlagRulesResponse struct {
	value *UpdateFlagRulesResponse
	isSet bool
}

func (v NullableUpdateFlagRulesResponse) Get() *UpdateFlagRulesResponse {
	return v.value
}

func (v *NullableUpdateFlagRulesResponse) Set(val *UpdateFlagRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFlagRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFlagRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFlagRulesResponse(val *UpdateFlagRulesResponse) *NullableUpdateFlagRulesResponse {
	return &NullableUpdateFlagRulesResponse{value: val, isSet: true}
}

func (v NullableUpdateFlagRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFlagRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
