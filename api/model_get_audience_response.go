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

// checks if the GetAudienceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAudienceResponse{}

// GetAudienceResponse struct for GetAudienceResponse
type GetAudienceResponse struct {
	Data PlanAudienceDetailResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _GetAudienceResponse GetAudienceResponse

// NewGetAudienceResponse instantiates a new GetAudienceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAudienceResponse(data PlanAudienceDetailResponseData, params map[string]interface{}) *GetAudienceResponse {
	this := GetAudienceResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetAudienceResponseWithDefaults instantiates a new GetAudienceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAudienceResponseWithDefaults() *GetAudienceResponse {
	this := GetAudienceResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetAudienceResponse) GetData() PlanAudienceDetailResponseData {
	if o == nil {
		var ret PlanAudienceDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetAudienceResponse) GetDataOk() (*PlanAudienceDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetAudienceResponse) SetData(v PlanAudienceDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetAudienceResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetAudienceResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetAudienceResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetAudienceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAudienceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetAudienceResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetAudienceResponse := _GetAudienceResponse{}

	err = json.Unmarshal(data, &varGetAudienceResponse)

	if err != nil {
		return err
	}

	*o = GetAudienceResponse(varGetAudienceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAudienceResponse struct {
	value *GetAudienceResponse
	isSet bool
}

func (v NullableGetAudienceResponse) Get() *GetAudienceResponse {
	return v.value
}

func (v *NullableGetAudienceResponse) Set(val *GetAudienceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAudienceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAudienceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAudienceResponse(val *GetAudienceResponse) *NullableGetAudienceResponse {
	return &NullableGetAudienceResponse{value: val, isSet: true}
}

func (v NullableGetAudienceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAudienceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
