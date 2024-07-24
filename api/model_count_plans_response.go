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

// checks if the CountPlansResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountPlansResponse{}

// CountPlansResponse struct for CountPlansResponse
type CountPlansResponse struct {
	Data                 CountResponse    `json:"data"`
	Params               CountPlansParams `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _CountPlansResponse CountPlansResponse

// NewCountPlansResponse instantiates a new CountPlansResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountPlansResponse(data CountResponse, params CountPlansParams) *CountPlansResponse {
	this := CountPlansResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountPlansResponseWithDefaults instantiates a new CountPlansResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountPlansResponseWithDefaults() *CountPlansResponse {
	this := CountPlansResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountPlansResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountPlansResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountPlansResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountPlansResponse) GetParams() CountPlansParams {
	if o == nil {
		var ret CountPlansParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountPlansResponse) GetParamsOk() (*CountPlansParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountPlansResponse) SetParams(v CountPlansParams) {
	o.Params = v
}

func (o CountPlansResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountPlansResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountPlansResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCountPlansResponse := _CountPlansResponse{}

	err = json.Unmarshal(data, &varCountPlansResponse)

	if err != nil {
		return err
	}

	*o = CountPlansResponse(varCountPlansResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountPlansResponse struct {
	value *CountPlansResponse
	isSet bool
}

func (v NullableCountPlansResponse) Get() *CountPlansResponse {
	return v.value
}

func (v *NullableCountPlansResponse) Set(val *CountPlansResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountPlansResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountPlansResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountPlansResponse(val *CountPlansResponse) *NullableCountPlansResponse {
	return &NullableCountPlansResponse{value: val, isSet: true}
}

func (v NullableCountPlansResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountPlansResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
