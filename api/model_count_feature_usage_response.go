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

// checks if the CountFeatureUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountFeatureUsageResponse{}

// CountFeatureUsageResponse struct for CountFeatureUsageResponse
type CountFeatureUsageResponse struct {
	Data   CountResponse           `json:"data"`
	Params CountFeatureUsageParams `json:"params"`
}

type _CountFeatureUsageResponse CountFeatureUsageResponse

// NewCountFeatureUsageResponse instantiates a new CountFeatureUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountFeatureUsageResponse(data CountResponse, params CountFeatureUsageParams) *CountFeatureUsageResponse {
	this := CountFeatureUsageResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountFeatureUsageResponseWithDefaults instantiates a new CountFeatureUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountFeatureUsageResponseWithDefaults() *CountFeatureUsageResponse {
	this := CountFeatureUsageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountFeatureUsageResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountFeatureUsageResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountFeatureUsageResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountFeatureUsageResponse) GetParams() CountFeatureUsageParams {
	if o == nil {
		var ret CountFeatureUsageParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountFeatureUsageResponse) GetParamsOk() (*CountFeatureUsageParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountFeatureUsageResponse) SetParams(v CountFeatureUsageParams) {
	o.Params = v
}

func (o CountFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountFeatureUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *CountFeatureUsageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCountFeatureUsageResponse := _CountFeatureUsageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCountFeatureUsageResponse)

	if err != nil {
		return err
	}

	*o = CountFeatureUsageResponse(varCountFeatureUsageResponse)

	return err
}

type NullableCountFeatureUsageResponse struct {
	value *CountFeatureUsageResponse
	isSet bool
}

func (v NullableCountFeatureUsageResponse) Get() *CountFeatureUsageResponse {
	return v.value
}

func (v *NullableCountFeatureUsageResponse) Set(val *CountFeatureUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountFeatureUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountFeatureUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountFeatureUsageResponse(val *CountFeatureUsageResponse) *NullableCountFeatureUsageResponse {
	return &NullableCountFeatureUsageResponse{value: val, isSet: true}
}

func (v NullableCountFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountFeatureUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
