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

// checks if the ListPaymentMethodsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPaymentMethodsResponse{}

// ListPaymentMethodsResponse struct for ListPaymentMethodsResponse
type ListPaymentMethodsResponse struct {
	// The returned resources
	Data                 []PaymentMethodResponseData `json:"data"`
	Params               ListPaymentMethodsParams    `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListPaymentMethodsResponse ListPaymentMethodsResponse

// NewListPaymentMethodsResponse instantiates a new ListPaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentMethodsResponse(data []PaymentMethodResponseData, params ListPaymentMethodsParams) *ListPaymentMethodsResponse {
	this := ListPaymentMethodsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListPaymentMethodsResponseWithDefaults instantiates a new ListPaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentMethodsResponseWithDefaults() *ListPaymentMethodsResponse {
	this := ListPaymentMethodsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListPaymentMethodsResponse) GetData() []PaymentMethodResponseData {
	if o == nil {
		var ret []PaymentMethodResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsResponse) GetDataOk() ([]PaymentMethodResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListPaymentMethodsResponse) SetData(v []PaymentMethodResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListPaymentMethodsResponse) GetParams() ListPaymentMethodsParams {
	if o == nil {
		var ret ListPaymentMethodsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListPaymentMethodsResponse) GetParamsOk() (*ListPaymentMethodsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListPaymentMethodsResponse) SetParams(v ListPaymentMethodsParams) {
	o.Params = v
}

func (o ListPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPaymentMethodsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPaymentMethodsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListPaymentMethodsResponse := _ListPaymentMethodsResponse{}

	err = json.Unmarshal(data, &varListPaymentMethodsResponse)

	if err != nil {
		return err
	}

	*o = ListPaymentMethodsResponse(varListPaymentMethodsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPaymentMethodsResponse struct {
	value *ListPaymentMethodsResponse
	isSet bool
}

func (v NullableListPaymentMethodsResponse) Get() *ListPaymentMethodsResponse {
	return v.value
}

func (v *NullableListPaymentMethodsResponse) Set(val *ListPaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentMethodsResponse(val *ListPaymentMethodsResponse) *NullableListPaymentMethodsResponse {
	return &NullableListPaymentMethodsResponse{value: val, isSet: true}
}

func (v NullableListPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
