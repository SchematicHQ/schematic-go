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

// checks if the CountComponentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountComponentsResponse{}

// CountComponentsResponse struct for CountComponentsResponse
type CountComponentsResponse struct {
	Data   CountResponse         `json:"data"`
	Params CountComponentsParams `json:"params"`
}

type _CountComponentsResponse CountComponentsResponse

// NewCountComponentsResponse instantiates a new CountComponentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountComponentsResponse(data CountResponse, params CountComponentsParams) *CountComponentsResponse {
	this := CountComponentsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountComponentsResponseWithDefaults instantiates a new CountComponentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountComponentsResponseWithDefaults() *CountComponentsResponse {
	this := CountComponentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountComponentsResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountComponentsResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountComponentsResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountComponentsResponse) GetParams() CountComponentsParams {
	if o == nil {
		var ret CountComponentsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountComponentsResponse) GetParamsOk() (*CountComponentsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountComponentsResponse) SetParams(v CountComponentsParams) {
	o.Params = v
}

func (o CountComponentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountComponentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *CountComponentsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCountComponentsResponse := _CountComponentsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountComponentsResponse)

	if err != nil {
		return err
	}

	*o = CountComponentsResponse(varCountComponentsResponse)

	return err
}

type NullableCountComponentsResponse struct {
	value *CountComponentsResponse
	isSet bool
}

func (v NullableCountComponentsResponse) Get() *CountComponentsResponse {
	return v.value
}

func (v *NullableCountComponentsResponse) Set(val *CountComponentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountComponentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountComponentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountComponentsResponse(val *CountComponentsResponse) *NullableCountComponentsResponse {
	return &NullableCountComponentsResponse{value: val, isSet: true}
}

func (v NullableCountComponentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountComponentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
