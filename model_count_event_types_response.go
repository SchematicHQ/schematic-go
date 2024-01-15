/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
)

// checks if the CountEventTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountEventTypesResponse{}

// CountEventTypesResponse struct for CountEventTypesResponse
type CountEventTypesResponse struct {
	Data CountResponse `json:"data"`
	Params CountEventTypesParams `json:"params"`
}

// NewCountEventTypesResponse instantiates a new CountEventTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountEventTypesResponse(data CountResponse, params CountEventTypesParams) *CountEventTypesResponse {
	this := CountEventTypesResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountEventTypesResponseWithDefaults instantiates a new CountEventTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountEventTypesResponseWithDefaults() *CountEventTypesResponse {
	this := CountEventTypesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountEventTypesResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountEventTypesResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountEventTypesResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountEventTypesResponse) GetParams() CountEventTypesParams {
	if o == nil {
		var ret CountEventTypesParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountEventTypesResponse) GetParamsOk() (*CountEventTypesParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountEventTypesResponse) SetParams(v CountEventTypesParams) {
	o.Params = v
}

func (o CountEventTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountEventTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableCountEventTypesResponse struct {
	value *CountEventTypesResponse
	isSet bool
}

func (v NullableCountEventTypesResponse) Get() *CountEventTypesResponse {
	return v.value
}

func (v *NullableCountEventTypesResponse) Set(val *CountEventTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountEventTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountEventTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountEventTypesResponse(val *CountEventTypesResponse) *NullableCountEventTypesResponse {
	return &NullableCountEventTypesResponse{value: val, isSet: true}
}

func (v NullableCountEventTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountEventTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


