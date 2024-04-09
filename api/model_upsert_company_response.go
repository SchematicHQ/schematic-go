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

// checks if the UpsertCompanyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertCompanyResponse{}

// UpsertCompanyResponse struct for UpsertCompanyResponse
type UpsertCompanyResponse struct {
	Data CompanyDetailResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _UpsertCompanyResponse UpsertCompanyResponse

// NewUpsertCompanyResponse instantiates a new UpsertCompanyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertCompanyResponse(data CompanyDetailResponseData, params map[string]interface{}) *UpsertCompanyResponse {
	this := UpsertCompanyResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpsertCompanyResponseWithDefaults instantiates a new UpsertCompanyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertCompanyResponseWithDefaults() *UpsertCompanyResponse {
	this := UpsertCompanyResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpsertCompanyResponse) GetData() CompanyDetailResponseData {
	if o == nil {
		var ret CompanyDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpsertCompanyResponse) GetDataOk() (*CompanyDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpsertCompanyResponse) SetData(v CompanyDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpsertCompanyResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpsertCompanyResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpsertCompanyResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpsertCompanyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertCompanyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *UpsertCompanyResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpsertCompanyResponse := _UpsertCompanyResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpsertCompanyResponse)

	if err != nil {
		return err
	}

	*o = UpsertCompanyResponse(varUpsertCompanyResponse)

	return err
}

type NullableUpsertCompanyResponse struct {
	value *UpsertCompanyResponse
	isSet bool
}

func (v NullableUpsertCompanyResponse) Get() *UpsertCompanyResponse {
	return v.value
}

func (v *NullableUpsertCompanyResponse) Set(val *UpsertCompanyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertCompanyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertCompanyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertCompanyResponse(val *UpsertCompanyResponse) *NullableUpsertCompanyResponse {
	return &NullableUpsertCompanyResponse{value: val, isSet: true}
}

func (v NullableUpsertCompanyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertCompanyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
