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

// checks if the UpdateAudienceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAudienceResponse{}

// UpdateAudienceResponse struct for UpdateAudienceResponse
type UpdateAudienceResponse struct {
	Data PlanAudienceDetailResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _UpdateAudienceResponse UpdateAudienceResponse

// NewUpdateAudienceResponse instantiates a new UpdateAudienceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAudienceResponse(data PlanAudienceDetailResponseData, params map[string]interface{}) *UpdateAudienceResponse {
	this := UpdateAudienceResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdateAudienceResponseWithDefaults instantiates a new UpdateAudienceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAudienceResponseWithDefaults() *UpdateAudienceResponse {
	this := UpdateAudienceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateAudienceResponse) GetData() PlanAudienceDetailResponseData {
	if o == nil {
		var ret PlanAudienceDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateAudienceResponse) GetDataOk() (*PlanAudienceDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateAudienceResponse) SetData(v PlanAudienceDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdateAudienceResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateAudienceResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdateAudienceResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdateAudienceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAudienceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *UpdateAudienceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateAudienceResponse := _UpdateAudienceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateAudienceResponse)

	if err != nil {
		return err
	}

	*o = UpdateAudienceResponse(varUpdateAudienceResponse)

	return err
}

type NullableUpdateAudienceResponse struct {
	value *UpdateAudienceResponse
	isSet bool
}

func (v NullableUpdateAudienceResponse) Get() *UpdateAudienceResponse {
	return v.value
}

func (v *NullableUpdateAudienceResponse) Set(val *UpdateAudienceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAudienceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAudienceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAudienceResponse(val *UpdateAudienceResponse) *NullableUpdateAudienceResponse {
	return &NullableUpdateAudienceResponse{value: val, isSet: true}
}

func (v NullableUpdateAudienceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAudienceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
