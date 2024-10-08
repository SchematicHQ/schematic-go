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

// checks if the GetPlanGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPlanGroupResponse{}

// GetPlanGroupResponse struct for GetPlanGroupResponse
type GetPlanGroupResponse struct {
	Data PlanGroupDetailResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _GetPlanGroupResponse GetPlanGroupResponse

// NewGetPlanGroupResponse instantiates a new GetPlanGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlanGroupResponse(data PlanGroupDetailResponseData, params map[string]interface{}) *GetPlanGroupResponse {
	this := GetPlanGroupResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetPlanGroupResponseWithDefaults instantiates a new GetPlanGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlanGroupResponseWithDefaults() *GetPlanGroupResponse {
	this := GetPlanGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetPlanGroupResponse) GetData() PlanGroupDetailResponseData {
	if o == nil {
		var ret PlanGroupDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPlanGroupResponse) GetDataOk() (*PlanGroupDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetPlanGroupResponse) SetData(v PlanGroupDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetPlanGroupResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetPlanGroupResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetPlanGroupResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetPlanGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPlanGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPlanGroupResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetPlanGroupResponse := _GetPlanGroupResponse{}

	err = json.Unmarshal(data, &varGetPlanGroupResponse)

	if err != nil {
		return err
	}

	*o = GetPlanGroupResponse(varGetPlanGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPlanGroupResponse struct {
	value *GetPlanGroupResponse
	isSet bool
}

func (v NullableGetPlanGroupResponse) Get() *GetPlanGroupResponse {
	return v.value
}

func (v *NullableGetPlanGroupResponse) Set(val *GetPlanGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlanGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlanGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlanGroupResponse(val *GetPlanGroupResponse) *NullableGetPlanGroupResponse {
	return &NullableGetPlanGroupResponse{value: val, isSet: true}
}

func (v NullableGetPlanGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlanGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
