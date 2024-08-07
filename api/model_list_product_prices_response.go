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

// checks if the ListProductPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListProductPricesResponse{}

// ListProductPricesResponse struct for ListProductPricesResponse
type ListProductPricesResponse struct {
	// The returned resources
	Data                 []BillingPriceResponseData `json:"data"`
	Params               ListProductPricesParams    `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListProductPricesResponse ListProductPricesResponse

// NewListProductPricesResponse instantiates a new ListProductPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductPricesResponse(data []BillingPriceResponseData, params ListProductPricesParams) *ListProductPricesResponse {
	this := ListProductPricesResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListProductPricesResponseWithDefaults instantiates a new ListProductPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductPricesResponseWithDefaults() *ListProductPricesResponse {
	this := ListProductPricesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListProductPricesResponse) GetData() []BillingPriceResponseData {
	if o == nil {
		var ret []BillingPriceResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListProductPricesResponse) GetDataOk() ([]BillingPriceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListProductPricesResponse) SetData(v []BillingPriceResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListProductPricesResponse) GetParams() ListProductPricesParams {
	if o == nil {
		var ret ListProductPricesParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListProductPricesResponse) GetParamsOk() (*ListProductPricesParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListProductPricesResponse) SetParams(v ListProductPricesParams) {
	o.Params = v
}

func (o ListProductPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListProductPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListProductPricesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListProductPricesResponse := _ListProductPricesResponse{}

	err = json.Unmarshal(data, &varListProductPricesResponse)

	if err != nil {
		return err
	}

	*o = ListProductPricesResponse(varListProductPricesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListProductPricesResponse struct {
	value *ListProductPricesResponse
	isSet bool
}

func (v NullableListProductPricesResponse) Get() *ListProductPricesResponse {
	return v.value
}

func (v *NullableListProductPricesResponse) Set(val *ListProductPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductPricesResponse(val *ListProductPricesResponse) *NullableListProductPricesResponse {
	return &NullableListProductPricesResponse{value: val, isSet: true}
}

func (v NullableListProductPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
