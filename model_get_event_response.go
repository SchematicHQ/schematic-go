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

// checks if the GetEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventResponse{}

// GetEventResponse struct for GetEventResponse
type GetEventResponse struct {
	Data EventResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewGetEventResponse instantiates a new GetEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventResponse(data EventResponseData, params map[string]interface{}) *GetEventResponse {
	this := GetEventResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetEventResponseWithDefaults instantiates a new GetEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventResponseWithDefaults() *GetEventResponse {
	this := GetEventResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventResponse) GetData() EventResponseData {
	if o == nil {
		var ret EventResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventResponse) GetDataOk() (*EventResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEventResponse) SetData(v EventResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetEventResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetEventResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetEventResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableGetEventResponse struct {
	value *GetEventResponse
	isSet bool
}

func (v NullableGetEventResponse) Get() *GetEventResponse {
	return v.value
}

func (v *NullableGetEventResponse) Set(val *GetEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventResponse(val *GetEventResponse) *NullableGetEventResponse {
	return &NullableGetEventResponse{value: val, isSet: true}
}

func (v NullableGetEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


