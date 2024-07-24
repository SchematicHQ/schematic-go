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

// checks if the ListProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProductsResponse{}

// ListProductsResponse struct for ListProductsResponse
type ListProductsResponse struct {
	// The returned resources
	Data                 []BillingProductResponseData `json:"data"`
	Params               ListProductsParams           `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListProductsResponse ListProductsResponse

// NewListProductsResponse instantiates a new ListProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductsResponse(data []BillingProductResponseData, params ListProductsParams) *ListProductsResponse {
	this := ListProductsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListProductsResponseWithDefaults instantiates a new ListProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductsResponseWithDefaults() *ListProductsResponse {
	this := ListProductsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListProductsResponse) GetData() []BillingProductResponseData {
	if o == nil {
		var ret []BillingProductResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListProductsResponse) GetDataOk() ([]BillingProductResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListProductsResponse) SetData(v []BillingProductResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListProductsResponse) GetParams() ListProductsParams {
	if o == nil {
		var ret ListProductsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListProductsResponse) GetParamsOk() (*ListProductsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListProductsResponse) SetParams(v ListProductsParams) {
	o.Params = v
}

func (o ListProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListProductsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListProductsResponse := _ListProductsResponse{}

	err = json.Unmarshal(data, &varListProductsResponse)

	if err != nil {
		return err
	}

	*o = ListProductsResponse(varListProductsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListProductsResponse struct {
	value *ListProductsResponse
	isSet bool
}

func (v NullableListProductsResponse) Get() *ListProductsResponse {
	return v.value
}

func (v *NullableListProductsResponse) Set(val *ListProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductsResponse(val *ListProductsResponse) *NullableListProductsResponse {
	return &NullableListProductsResponse{value: val, isSet: true}
}

func (v NullableListProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
