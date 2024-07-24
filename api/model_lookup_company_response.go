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

// checks if the LookupCompanyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupCompanyResponse{}

// LookupCompanyResponse struct for LookupCompanyResponse
type LookupCompanyResponse struct {
	Data                 CompanyDetailResponseData `json:"data"`
	Params               LookupCompanyParams       `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _LookupCompanyResponse LookupCompanyResponse

// NewLookupCompanyResponse instantiates a new LookupCompanyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupCompanyResponse(data CompanyDetailResponseData, params LookupCompanyParams) *LookupCompanyResponse {
	this := LookupCompanyResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewLookupCompanyResponseWithDefaults instantiates a new LookupCompanyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupCompanyResponseWithDefaults() *LookupCompanyResponse {
	this := LookupCompanyResponse{}
	return &this
}

// GetData returns the Data field value
func (o *LookupCompanyResponse) GetData() CompanyDetailResponseData {
	if o == nil {
		var ret CompanyDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LookupCompanyResponse) GetDataOk() (*CompanyDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LookupCompanyResponse) SetData(v CompanyDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *LookupCompanyResponse) GetParams() LookupCompanyParams {
	if o == nil {
		var ret LookupCompanyParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *LookupCompanyResponse) GetParamsOk() (*LookupCompanyParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *LookupCompanyResponse) SetParams(v LookupCompanyParams) {
	o.Params = v
}

func (o LookupCompanyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupCompanyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LookupCompanyResponse) UnmarshalJSON(data []byte) (err error) {
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

	varLookupCompanyResponse := _LookupCompanyResponse{}

	err = json.Unmarshal(data, &varLookupCompanyResponse)

	if err != nil {
		return err
	}

	*o = LookupCompanyResponse(varLookupCompanyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLookupCompanyResponse struct {
	value *LookupCompanyResponse
	isSet bool
}

func (v NullableLookupCompanyResponse) Get() *LookupCompanyResponse {
	return v.value
}

func (v *NullableLookupCompanyResponse) Set(val *LookupCompanyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupCompanyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupCompanyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupCompanyResponse(val *LookupCompanyResponse) *NullableLookupCompanyResponse {
	return &NullableLookupCompanyResponse{value: val, isSet: true}
}

func (v NullableLookupCompanyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupCompanyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
