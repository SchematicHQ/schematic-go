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

// checks if the ListFeatureUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFeatureUsageResponse{}

// ListFeatureUsageResponse struct for ListFeatureUsageResponse
type ListFeatureUsageResponse struct {
	// The returned resources
	Data   []FeatureUsageResponseData `json:"data"`
	Params ListFeatureUsageParams     `json:"params"`
}

type _ListFeatureUsageResponse ListFeatureUsageResponse

// NewListFeatureUsageResponse instantiates a new ListFeatureUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFeatureUsageResponse(data []FeatureUsageResponseData, params ListFeatureUsageParams) *ListFeatureUsageResponse {
	this := ListFeatureUsageResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListFeatureUsageResponseWithDefaults instantiates a new ListFeatureUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFeatureUsageResponseWithDefaults() *ListFeatureUsageResponse {
	this := ListFeatureUsageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListFeatureUsageResponse) GetData() []FeatureUsageResponseData {
	if o == nil {
		var ret []FeatureUsageResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListFeatureUsageResponse) GetDataOk() ([]FeatureUsageResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListFeatureUsageResponse) SetData(v []FeatureUsageResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListFeatureUsageResponse) GetParams() ListFeatureUsageParams {
	if o == nil {
		var ret ListFeatureUsageParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListFeatureUsageResponse) GetParamsOk() (*ListFeatureUsageParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListFeatureUsageResponse) SetParams(v ListFeatureUsageParams) {
	o.Params = v
}

func (o ListFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFeatureUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *ListFeatureUsageResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListFeatureUsageResponse := _ListFeatureUsageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListFeatureUsageResponse)

	if err != nil {
		return err
	}

	*o = ListFeatureUsageResponse(varListFeatureUsageResponse)

	return err
}

type NullableListFeatureUsageResponse struct {
	value *ListFeatureUsageResponse
	isSet bool
}

func (v NullableListFeatureUsageResponse) Get() *ListFeatureUsageResponse {
	return v.value
}

func (v *NullableListFeatureUsageResponse) Set(val *ListFeatureUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFeatureUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFeatureUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFeatureUsageResponse(val *ListFeatureUsageResponse) *NullableListFeatureUsageResponse {
	return &NullableListFeatureUsageResponse{value: val, isSet: true}
}

func (v NullableListFeatureUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFeatureUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
