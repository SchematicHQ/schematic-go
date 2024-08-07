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

// checks if the UpdatePlanEntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePlanEntitlementResponse{}

// UpdatePlanEntitlementResponse struct for UpdatePlanEntitlementResponse
type UpdatePlanEntitlementResponse struct {
	Data PlanEntitlementResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _UpdatePlanEntitlementResponse UpdatePlanEntitlementResponse

// NewUpdatePlanEntitlementResponse instantiates a new UpdatePlanEntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlanEntitlementResponse(data PlanEntitlementResponseData, params map[string]interface{}) *UpdatePlanEntitlementResponse {
	this := UpdatePlanEntitlementResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdatePlanEntitlementResponseWithDefaults instantiates a new UpdatePlanEntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlanEntitlementResponseWithDefaults() *UpdatePlanEntitlementResponse {
	this := UpdatePlanEntitlementResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdatePlanEntitlementResponse) GetData() PlanEntitlementResponseData {
	if o == nil {
		var ret PlanEntitlementResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdatePlanEntitlementResponse) GetDataOk() (*PlanEntitlementResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdatePlanEntitlementResponse) SetData(v PlanEntitlementResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdatePlanEntitlementResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdatePlanEntitlementResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdatePlanEntitlementResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdatePlanEntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePlanEntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdatePlanEntitlementResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdatePlanEntitlementResponse := _UpdatePlanEntitlementResponse{}

	err = json.Unmarshal(data, &varUpdatePlanEntitlementResponse)

	if err != nil {
		return err
	}

	*o = UpdatePlanEntitlementResponse(varUpdatePlanEntitlementResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdatePlanEntitlementResponse struct {
	value *UpdatePlanEntitlementResponse
	isSet bool
}

func (v NullableUpdatePlanEntitlementResponse) Get() *UpdatePlanEntitlementResponse {
	return v.value
}

func (v *NullableUpdatePlanEntitlementResponse) Set(val *UpdatePlanEntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlanEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlanEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlanEntitlementResponse(val *UpdatePlanEntitlementResponse) *NullableUpdatePlanEntitlementResponse {
	return &NullableUpdatePlanEntitlementResponse{value: val, isSet: true}
}

func (v NullableUpdatePlanEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePlanEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
