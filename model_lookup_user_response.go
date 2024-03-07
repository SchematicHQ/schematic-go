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

// checks if the LookupUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupUserResponse{}

// LookupUserResponse struct for LookupUserResponse
type LookupUserResponse struct {
	Data UserDetailResponseData `json:"data"`
	Params LookupUserParams `json:"params"`
}

// NewLookupUserResponse instantiates a new LookupUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupUserResponse(data UserDetailResponseData, params LookupUserParams) *LookupUserResponse {
	this := LookupUserResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewLookupUserResponseWithDefaults instantiates a new LookupUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupUserResponseWithDefaults() *LookupUserResponse {
	this := LookupUserResponse{}
	return &this
}

// GetData returns the Data field value
func (o *LookupUserResponse) GetData() UserDetailResponseData {
	if o == nil {
		var ret UserDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LookupUserResponse) GetDataOk() (*UserDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LookupUserResponse) SetData(v UserDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *LookupUserResponse) GetParams() LookupUserParams {
	if o == nil {
		var ret LookupUserParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *LookupUserResponse) GetParamsOk() (*LookupUserParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *LookupUserResponse) SetParams(v LookupUserParams) {
	o.Params = v
}

func (o LookupUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableLookupUserResponse struct {
	value *LookupUserResponse
	isSet bool
}

func (v NullableLookupUserResponse) Get() *LookupUserResponse {
	return v.value
}

func (v *NullableLookupUserResponse) Set(val *LookupUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupUserResponse(val *LookupUserResponse) *NullableLookupUserResponse {
	return &NullableLookupUserResponse{value: val, isSet: true}
}

func (v NullableLookupUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


