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

// checks if the CountApiRequestsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountApiRequestsResponse{}

// CountApiRequestsResponse struct for CountApiRequestsResponse
type CountApiRequestsResponse struct {
	Data CountResponse `json:"data"`
	Params CountApiRequestsParams `json:"params"`
}

// NewCountApiRequestsResponse instantiates a new CountApiRequestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountApiRequestsResponse(data CountResponse, params CountApiRequestsParams) *CountApiRequestsResponse {
	this := CountApiRequestsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCountApiRequestsResponseWithDefaults instantiates a new CountApiRequestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountApiRequestsResponseWithDefaults() *CountApiRequestsResponse {
	this := CountApiRequestsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CountApiRequestsResponse) GetData() CountResponse {
	if o == nil {
		var ret CountResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CountApiRequestsResponse) GetDataOk() (*CountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CountApiRequestsResponse) SetData(v CountResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CountApiRequestsResponse) GetParams() CountApiRequestsParams {
	if o == nil {
		var ret CountApiRequestsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CountApiRequestsResponse) GetParamsOk() (*CountApiRequestsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *CountApiRequestsResponse) SetParams(v CountApiRequestsParams) {
	o.Params = v
}

func (o CountApiRequestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountApiRequestsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableCountApiRequestsResponse struct {
	value *CountApiRequestsResponse
	isSet bool
}

func (v NullableCountApiRequestsResponse) Get() *CountApiRequestsResponse {
	return v.value
}

func (v *NullableCountApiRequestsResponse) Set(val *CountApiRequestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountApiRequestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountApiRequestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountApiRequestsResponse(val *CountApiRequestsResponse) *NullableCountApiRequestsResponse {
	return &NullableCountApiRequestsResponse{value: val, isSet: true}
}

func (v NullableCountApiRequestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountApiRequestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


