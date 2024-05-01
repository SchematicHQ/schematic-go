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

// checks if the CountFeaturesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountFeaturesResponse{}

// CountFeaturesResponse struct for CountFeaturesResponse
type CountFeaturesResponse struct {
	Data   CountResponse       `json:"data"`
	Params CountFeaturesParams `json:"params"`
}

type _CountFeaturesResponse CountFeaturesResponse

// NewCountFeaturesResponse instantiates a new CountFeaturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountFeaturesResponse(data CountResponse, params CountFeaturesParams) *CountFeaturesResponse {
	this := CountFeaturesResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountFeaturesResponseWithDefaults instantiates a new CountFeaturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountFeaturesResponseWithDefaults() *CountFeaturesResponse {
	this := CountFeaturesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountFeaturesResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountFeaturesResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountFeaturesResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountFeaturesResponse) GetParams() CountFeaturesParams {
	if o == nil {
		var ret CountFeaturesParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountFeaturesResponse) GetParamsOk() (*CountFeaturesParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountFeaturesResponse) SetParams(v CountFeaturesParams) {
	o.Params = v
}

func (o CountFeaturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountFeaturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *CountFeaturesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCountFeaturesResponse := _CountFeaturesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountFeaturesResponse)

	if err != nil {
		return err
	}

	*o = CountFeaturesResponse(varCountFeaturesResponse)

	return err
}

type NullableCountFeaturesResponse struct {
	value *CountFeaturesResponse
	isSet bool
}

func (v NullableCountFeaturesResponse) Get() *CountFeaturesResponse {
	return v.value
}

func (v *NullableCountFeaturesResponse) Set(val *CountFeaturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountFeaturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountFeaturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountFeaturesResponse(val *CountFeaturesResponse) *NullableCountFeaturesResponse {
	return &NullableCountFeaturesResponse{value: val, isSet: true}
}

func (v NullableCountFeaturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountFeaturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}