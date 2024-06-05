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

// checks if the UpsertCRMDealResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertCRMDealResponse{}

// UpsertCRMDealResponse struct for UpsertCRMDealResponse
type UpsertCRMDealResponse struct {
	Data CRMDealResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _UpsertCRMDealResponse UpsertCRMDealResponse

// NewUpsertCRMDealResponse instantiates a new UpsertCRMDealResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertCRMDealResponse(data CRMDealResponseData, params map[string]interface{}) *UpsertCRMDealResponse {
	this := UpsertCRMDealResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpsertCRMDealResponseWithDefaults instantiates a new UpsertCRMDealResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertCRMDealResponseWithDefaults() *UpsertCRMDealResponse {
	this := UpsertCRMDealResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpsertCRMDealResponse) GetData() CRMDealResponseData {
	if o == nil {
		var ret CRMDealResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpsertCRMDealResponse) GetDataOk() (*CRMDealResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpsertCRMDealResponse) SetData(v CRMDealResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpsertCRMDealResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpsertCRMDealResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpsertCRMDealResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpsertCRMDealResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertCRMDealResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *UpsertCRMDealResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpsertCRMDealResponse := _UpsertCRMDealResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpsertCRMDealResponse)

	if err != nil {
		return err
	}

	*o = UpsertCRMDealResponse(varUpsertCRMDealResponse)

	return err
}

type NullableUpsertCRMDealResponse struct {
	value *UpsertCRMDealResponse
	isSet bool
}

func (v NullableUpsertCRMDealResponse) Get() *UpsertCRMDealResponse {
	return v.value
}

func (v *NullableUpsertCRMDealResponse) Set(val *UpsertCRMDealResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertCRMDealResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertCRMDealResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertCRMDealResponse(val *UpsertCRMDealResponse) *NullableUpsertCRMDealResponse {
	return &NullableUpsertCRMDealResponse{value: val, isSet: true}
}

func (v NullableUpsertCRMDealResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertCRMDealResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
