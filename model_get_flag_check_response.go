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

// checks if the GetFlagCheckResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlagCheckResponse{}

// GetFlagCheckResponse struct for GetFlagCheckResponse
type GetFlagCheckResponse struct {
	Data FlagCheckLogDetailResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewGetFlagCheckResponse instantiates a new GetFlagCheckResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlagCheckResponse(data FlagCheckLogDetailResponseData, params map[string]interface{}) *GetFlagCheckResponse {
	this := GetFlagCheckResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetFlagCheckResponseWithDefaults instantiates a new GetFlagCheckResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlagCheckResponseWithDefaults() *GetFlagCheckResponse {
	this := GetFlagCheckResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlagCheckResponse) GetData() FlagCheckLogDetailResponseData {
	if o == nil {
		var ret FlagCheckLogDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlagCheckResponse) GetDataOk() (*FlagCheckLogDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFlagCheckResponse) SetData(v FlagCheckLogDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetFlagCheckResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetFlagCheckResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetFlagCheckResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetFlagCheckResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlagCheckResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableGetFlagCheckResponse struct {
	value *GetFlagCheckResponse
	isSet bool
}

func (v NullableGetFlagCheckResponse) Get() *GetFlagCheckResponse {
	return v.value
}

func (v *NullableGetFlagCheckResponse) Set(val *GetFlagCheckResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlagCheckResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlagCheckResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlagCheckResponse(val *GetFlagCheckResponse) *NullableGetFlagCheckResponse {
	return &NullableGetFlagCheckResponse{value: val, isSet: true}
}

func (v NullableGetFlagCheckResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlagCheckResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


