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

// checks if the GetCompanyOverrideResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanyOverrideResponse{}

// GetCompanyOverrideResponse struct for GetCompanyOverrideResponse
type GetCompanyOverrideResponse struct {
	Data CompanyOverrideResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _GetCompanyOverrideResponse GetCompanyOverrideResponse

// NewGetCompanyOverrideResponse instantiates a new GetCompanyOverrideResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanyOverrideResponse(data CompanyOverrideResponseData, params map[string]interface{}) *GetCompanyOverrideResponse {
	this := GetCompanyOverrideResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetCompanyOverrideResponseWithDefaults instantiates a new GetCompanyOverrideResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanyOverrideResponseWithDefaults() *GetCompanyOverrideResponse {
	this := GetCompanyOverrideResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCompanyOverrideResponse) GetData() CompanyOverrideResponseData {
	if o == nil {
		var ret CompanyOverrideResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCompanyOverrideResponse) GetDataOk() (*CompanyOverrideResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCompanyOverrideResponse) SetData(v CompanyOverrideResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetCompanyOverrideResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetCompanyOverrideResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetCompanyOverrideResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetCompanyOverrideResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanyOverrideResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetCompanyOverrideResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCompanyOverrideResponse := _GetCompanyOverrideResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCompanyOverrideResponse)

	if err != nil {
		return err
	}

	*o = GetCompanyOverrideResponse(varGetCompanyOverrideResponse)

	return err
}

type NullableGetCompanyOverrideResponse struct {
	value *GetCompanyOverrideResponse
	isSet bool
}

func (v NullableGetCompanyOverrideResponse) Get() *GetCompanyOverrideResponse {
	return v.value
}

func (v *NullableGetCompanyOverrideResponse) Set(val *GetCompanyOverrideResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanyOverrideResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanyOverrideResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanyOverrideResponse(val *GetCompanyOverrideResponse) *NullableGetCompanyOverrideResponse {
	return &NullableGetCompanyOverrideResponse{value: val, isSet: true}
}

func (v NullableGetCompanyOverrideResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanyOverrideResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
