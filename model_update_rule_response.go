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

// checks if the UpdateRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRuleResponse{}

// UpdateRuleResponse struct for UpdateRuleResponse
type UpdateRuleResponse struct {
	Data RuleDetailResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewUpdateRuleResponse instantiates a new UpdateRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRuleResponse(data RuleDetailResponseData, params map[string]interface{}) *UpdateRuleResponse {
	this := UpdateRuleResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdateRuleResponseWithDefaults instantiates a new UpdateRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRuleResponseWithDefaults() *UpdateRuleResponse {
	this := UpdateRuleResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateRuleResponse) GetData() RuleDetailResponseData {
	if o == nil {
		var ret RuleDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponse) GetDataOk() (*RuleDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateRuleResponse) SetData(v RuleDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdateRuleResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdateRuleResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdateRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableUpdateRuleResponse struct {
	value *UpdateRuleResponse
	isSet bool
}

func (v NullableUpdateRuleResponse) Get() *UpdateRuleResponse {
	return v.value
}

func (v *NullableUpdateRuleResponse) Set(val *UpdateRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRuleResponse(val *UpdateRuleResponse) *NullableUpdateRuleResponse {
	return &NullableUpdateRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


