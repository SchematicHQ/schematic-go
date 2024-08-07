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

// checks if the GetCompanyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCompanyResponse{}

// GetCompanyResponse struct for GetCompanyResponse
type GetCompanyResponse struct {
	Data CompanyDetailResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _GetCompanyResponse GetCompanyResponse

// NewGetCompanyResponse instantiates a new GetCompanyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCompanyResponse(data CompanyDetailResponseData, params map[string]interface{}) *GetCompanyResponse {
	this := GetCompanyResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetCompanyResponseWithDefaults instantiates a new GetCompanyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCompanyResponseWithDefaults() *GetCompanyResponse {
	this := GetCompanyResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetCompanyResponse) GetData() CompanyDetailResponseData {
	if o == nil {
		var ret CompanyDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCompanyResponse) GetDataOk() (*CompanyDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCompanyResponse) SetData(v CompanyDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetCompanyResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetCompanyResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetCompanyResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetCompanyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCompanyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCompanyResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetCompanyResponse := _GetCompanyResponse{}

	err = json.Unmarshal(data, &varGetCompanyResponse)

	if err != nil {
		return err
	}

	*o = GetCompanyResponse(varGetCompanyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCompanyResponse struct {
	value *GetCompanyResponse
	isSet bool
}

func (v NullableGetCompanyResponse) Get() *GetCompanyResponse {
	return v.value
}

func (v *NullableGetCompanyResponse) Set(val *GetCompanyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCompanyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCompanyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCompanyResponse(val *GetCompanyResponse) *NullableGetCompanyResponse {
	return &NullableGetCompanyResponse{value: val, isSet: true}
}

func (v NullableGetCompanyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCompanyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
