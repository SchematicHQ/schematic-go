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

// checks if the CountFeatureUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountFeatureUsersResponse{}

// CountFeatureUsersResponse struct for CountFeatureUsersResponse
type CountFeatureUsersResponse struct {
	Data                 CountResponse           `json:"data"`
	Params               CountFeatureUsersParams `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _CountFeatureUsersResponse CountFeatureUsersResponse

// NewCountFeatureUsersResponse instantiates a new CountFeatureUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountFeatureUsersResponse(data CountResponse, params CountFeatureUsersParams) *CountFeatureUsersResponse {
	this := CountFeatureUsersResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountFeatureUsersResponseWithDefaults instantiates a new CountFeatureUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountFeatureUsersResponseWithDefaults() *CountFeatureUsersResponse {
	this := CountFeatureUsersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountFeatureUsersResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountFeatureUsersResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountFeatureUsersResponse) GetParams() CountFeatureUsersParams {
	if o == nil {
		var ret CountFeatureUsersParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersResponse) GetParamsOk() (*CountFeatureUsersParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountFeatureUsersResponse) SetParams(v CountFeatureUsersParams) {
	o.Params = v
}

func (o CountFeatureUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountFeatureUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountFeatureUsersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCountFeatureUsersResponse := _CountFeatureUsersResponse{}

	err = json.Unmarshal(data, &varCountFeatureUsersResponse)

	if err != nil {
		return err
	}

	*o = CountFeatureUsersResponse(varCountFeatureUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountFeatureUsersResponse struct {
	value *CountFeatureUsersResponse
	isSet bool
}

func (v NullableCountFeatureUsersResponse) Get() *CountFeatureUsersResponse {
	return v.value
}

func (v *NullableCountFeatureUsersResponse) Set(val *CountFeatureUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountFeatureUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountFeatureUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountFeatureUsersResponse(val *CountFeatureUsersResponse) *NullableCountFeatureUsersResponse {
	return &NullableCountFeatureUsersResponse{value: val, isSet: true}
}

func (v NullableCountFeatureUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountFeatureUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
