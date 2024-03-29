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

// checks if the GetEventTypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventTypeResponse{}

// GetEventTypeResponse struct for GetEventTypeResponse
type GetEventTypeResponse struct {
	Data EventSummaryResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewGetEventTypeResponse instantiates a new GetEventTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventTypeResponse(data EventSummaryResponseData, params map[string]interface{}) *GetEventTypeResponse {
	this := GetEventTypeResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetEventTypeResponseWithDefaults instantiates a new GetEventTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventTypeResponseWithDefaults() *GetEventTypeResponse {
	this := GetEventTypeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventTypeResponse) GetData() EventSummaryResponseData {
	if o == nil {
		var ret EventSummaryResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventTypeResponse) GetDataOk() (*EventSummaryResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEventTypeResponse) SetData(v EventSummaryResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetEventTypeResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetEventTypeResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetEventTypeResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetEventTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventTypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableGetEventTypeResponse struct {
	value *GetEventTypeResponse
	isSet bool
}

func (v NullableGetEventTypeResponse) Get() *GetEventTypeResponse {
	return v.value
}

func (v *NullableGetEventTypeResponse) Set(val *GetEventTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventTypeResponse(val *GetEventTypeResponse) *NullableGetEventTypeResponse {
	return &NullableGetEventTypeResponse{value: val, isSet: true}
}

func (v NullableGetEventTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


