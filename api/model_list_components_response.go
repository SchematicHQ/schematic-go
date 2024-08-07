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

// checks if the ListComponentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListComponentsResponse{}

// ListComponentsResponse struct for ListComponentsResponse
type ListComponentsResponse struct {
	// The returned resources
	Data                 []ComponentResponseData `json:"data"`
	Params               ListComponentsParams    `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListComponentsResponse ListComponentsResponse

// NewListComponentsResponse instantiates a new ListComponentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListComponentsResponse(data []ComponentResponseData, params ListComponentsParams) *ListComponentsResponse {
	this := ListComponentsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListComponentsResponseWithDefaults instantiates a new ListComponentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListComponentsResponseWithDefaults() *ListComponentsResponse {
	this := ListComponentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListComponentsResponse) GetData() []ComponentResponseData {
	if o == nil {
		var ret []ComponentResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListComponentsResponse) GetDataOk() ([]ComponentResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListComponentsResponse) SetData(v []ComponentResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListComponentsResponse) GetParams() ListComponentsParams {
	if o == nil {
		var ret ListComponentsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListComponentsResponse) GetParamsOk() (*ListComponentsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListComponentsResponse) SetParams(v ListComponentsParams) {
	o.Params = v
}

func (o ListComponentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListComponentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListComponentsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListComponentsResponse := _ListComponentsResponse{}

	err = json.Unmarshal(data, &varListComponentsResponse)

	if err != nil {
		return err
	}

	*o = ListComponentsResponse(varListComponentsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListComponentsResponse struct {
	value *ListComponentsResponse
	isSet bool
}

func (v NullableListComponentsResponse) Get() *ListComponentsResponse {
	return v.value
}

func (v *NullableListComponentsResponse) Set(val *ListComponentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListComponentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListComponentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListComponentsResponse(val *ListComponentsResponse) *NullableListComponentsResponse {
	return &NullableListComponentsResponse{value: val, isSet: true}
}

func (v NullableListComponentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListComponentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
