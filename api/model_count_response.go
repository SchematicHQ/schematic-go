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

// checks if the CountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountResponse{}

// CountResponse The created resource
type CountResponse struct {
	Count                int32 `json:"count"`
	AdditionalProperties map[string]interface{}
}

type _CountResponse CountResponse

// NewCountResponse instantiates a new CountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountResponse(count int32) *CountResponse {
	this := CountResponse{}
	this.Count = count
	return &this
}

// NewCountResponseWithDefaults instantiates a new CountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountResponseWithDefaults() *CountResponse {
	this := CountResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *CountResponse) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CountResponse) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *CountResponse) SetCount(v int32) {
	o.Count = v
}

func (o CountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varCountResponse := _CountResponse{}

	err = json.Unmarshal(data, &varCountResponse)

	if err != nil {
		return err
	}

	*o = CountResponse(varCountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountResponse struct {
	value *CountResponse
	isSet bool
}

func (v NullableCountResponse) Get() *CountResponse {
	return v.value
}

func (v *NullableCountResponse) Set(val *CountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountResponse(val *CountResponse) *NullableCountResponse {
	return &NullableCountResponse{value: val, isSet: true}
}

func (v NullableCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}