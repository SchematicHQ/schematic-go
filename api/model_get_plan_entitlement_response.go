/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetPlanEntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPlanEntitlementResponse{}

// GetPlanEntitlementResponse struct for GetPlanEntitlementResponse
type GetPlanEntitlementResponse struct {
	Data PlanEntitlementResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _GetPlanEntitlementResponse GetPlanEntitlementResponse

// NewGetPlanEntitlementResponse instantiates a new GetPlanEntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPlanEntitlementResponse(data PlanEntitlementResponseData, params map[string]interface{}) *GetPlanEntitlementResponse {
	this := GetPlanEntitlementResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetPlanEntitlementResponseWithDefaults instantiates a new GetPlanEntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPlanEntitlementResponseWithDefaults() *GetPlanEntitlementResponse {
	this := GetPlanEntitlementResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetPlanEntitlementResponse) GetData() PlanEntitlementResponseData {
	if o == nil {
		var ret PlanEntitlementResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPlanEntitlementResponse) GetDataOk() (*PlanEntitlementResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetPlanEntitlementResponse) SetData(v PlanEntitlementResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetPlanEntitlementResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetPlanEntitlementResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetPlanEntitlementResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetPlanEntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPlanEntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetPlanEntitlementResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetPlanEntitlementResponse := _GetPlanEntitlementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPlanEntitlementResponse)

	if err != nil {
		return err
	}

	*o = GetPlanEntitlementResponse(varGetPlanEntitlementResponse)

	return err
}

type NullableGetPlanEntitlementResponse struct {
	value *GetPlanEntitlementResponse
	isSet bool
}

func (v NullableGetPlanEntitlementResponse) Get() *GetPlanEntitlementResponse {
	return v.value
}

func (v *NullableGetPlanEntitlementResponse) Set(val *GetPlanEntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPlanEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPlanEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPlanEntitlementResponse(val *GetPlanEntitlementResponse) *NullableGetPlanEntitlementResponse {
	return &NullableGetPlanEntitlementResponse{value: val, isSet: true}
}

func (v NullableGetPlanEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPlanEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
