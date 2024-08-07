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

// checks if the UpsertLineItemResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertLineItemResponse{}

// UpsertLineItemResponse struct for UpsertLineItemResponse
type UpsertLineItemResponse struct {
	Data CrmLineItemResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _UpsertLineItemResponse UpsertLineItemResponse

// NewUpsertLineItemResponse instantiates a new UpsertLineItemResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertLineItemResponse(data CrmLineItemResponseData, params map[string]interface{}) *UpsertLineItemResponse {
	this := UpsertLineItemResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpsertLineItemResponseWithDefaults instantiates a new UpsertLineItemResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertLineItemResponseWithDefaults() *UpsertLineItemResponse {
	this := UpsertLineItemResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpsertLineItemResponse) GetData() CrmLineItemResponseData {
	if o == nil {
		var ret CrmLineItemResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpsertLineItemResponse) GetDataOk() (*CrmLineItemResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpsertLineItemResponse) SetData(v CrmLineItemResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpsertLineItemResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpsertLineItemResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpsertLineItemResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpsertLineItemResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertLineItemResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpsertLineItemResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpsertLineItemResponse := _UpsertLineItemResponse{}

	err = json.Unmarshal(data, &varUpsertLineItemResponse)

	if err != nil {
		return err
	}

	*o = UpsertLineItemResponse(varUpsertLineItemResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpsertLineItemResponse struct {
	value *UpsertLineItemResponse
	isSet bool
}

func (v NullableUpsertLineItemResponse) Get() *UpsertLineItemResponse {
	return v.value
}

func (v *NullableUpsertLineItemResponse) Set(val *UpsertLineItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertLineItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertLineItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertLineItemResponse(val *UpsertLineItemResponse) *NullableUpsertLineItemResponse {
	return &NullableUpsertLineItemResponse{value: val, isSet: true}
}

func (v NullableUpsertLineItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertLineItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
