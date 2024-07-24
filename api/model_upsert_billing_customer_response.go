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

// checks if the UpsertBillingCustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertBillingCustomerResponse{}

// UpsertBillingCustomerResponse struct for UpsertBillingCustomerResponse
type UpsertBillingCustomerResponse struct {
	Data BillingCustomerResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _UpsertBillingCustomerResponse UpsertBillingCustomerResponse

// NewUpsertBillingCustomerResponse instantiates a new UpsertBillingCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertBillingCustomerResponse(data BillingCustomerResponseData, params map[string]interface{}) *UpsertBillingCustomerResponse {
	this := UpsertBillingCustomerResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpsertBillingCustomerResponseWithDefaults instantiates a new UpsertBillingCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertBillingCustomerResponseWithDefaults() *UpsertBillingCustomerResponse {
	this := UpsertBillingCustomerResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpsertBillingCustomerResponse) GetData() BillingCustomerResponseData {
	if o == nil {
		var ret BillingCustomerResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpsertBillingCustomerResponse) GetDataOk() (*BillingCustomerResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpsertBillingCustomerResponse) SetData(v BillingCustomerResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpsertBillingCustomerResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpsertBillingCustomerResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpsertBillingCustomerResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpsertBillingCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertBillingCustomerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpsertBillingCustomerResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpsertBillingCustomerResponse := _UpsertBillingCustomerResponse{}

	err = json.Unmarshal(data, &varUpsertBillingCustomerResponse)

	if err != nil {
		return err
	}

	*o = UpsertBillingCustomerResponse(varUpsertBillingCustomerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpsertBillingCustomerResponse struct {
	value *UpsertBillingCustomerResponse
	isSet bool
}

func (v NullableUpsertBillingCustomerResponse) Get() *UpsertBillingCustomerResponse {
	return v.value
}

func (v *NullableUpsertBillingCustomerResponse) Set(val *UpsertBillingCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertBillingCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertBillingCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertBillingCustomerResponse(val *UpsertBillingCustomerResponse) *NullableUpsertBillingCustomerResponse {
	return &NullableUpsertBillingCustomerResponse{value: val, isSet: true}
}

func (v NullableUpsertBillingCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertBillingCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}