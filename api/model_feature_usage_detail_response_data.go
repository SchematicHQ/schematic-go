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

// checks if the FeatureUsageDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureUsageDetailResponseData{}

// FeatureUsageDetailResponseData struct for FeatureUsageDetailResponseData
type FeatureUsageDetailResponseData struct {
	Features []FeatureUsageResponseData `json:"features"`
}

type _FeatureUsageDetailResponseData FeatureUsageDetailResponseData

// NewFeatureUsageDetailResponseData instantiates a new FeatureUsageDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureUsageDetailResponseData(features []FeatureUsageResponseData) *FeatureUsageDetailResponseData {
	this := FeatureUsageDetailResponseData{}
	this.Features = features
	return &this
}

// NewFeatureUsageDetailResponseDataWithDefaults instantiates a new FeatureUsageDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureUsageDetailResponseDataWithDefaults() *FeatureUsageDetailResponseData {
	this := FeatureUsageDetailResponseData{}
	return &this
}

// GetFeatures returns the Features field value
func (o *FeatureUsageDetailResponseData) GetFeatures() []FeatureUsageResponseData {
	if o == nil {
		var ret []FeatureUsageResponseData
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *FeatureUsageDetailResponseData) GetFeaturesOk() ([]FeatureUsageResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *FeatureUsageDetailResponseData) SetFeatures(v []FeatureUsageResponseData) {
	o.Features = v
}

func (o FeatureUsageDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureUsageDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *FeatureUsageDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"features",
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

	varFeatureUsageDetailResponseData := _FeatureUsageDetailResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeatureUsageDetailResponseData)

	if err != nil {
		return err
	}

	*o = FeatureUsageDetailResponseData(varFeatureUsageDetailResponseData)

	return err
}

type NullableFeatureUsageDetailResponseData struct {
	value *FeatureUsageDetailResponseData
	isSet bool
}

func (v NullableFeatureUsageDetailResponseData) Get() *FeatureUsageDetailResponseData {
	return v.value
}

func (v *NullableFeatureUsageDetailResponseData) Set(val *FeatureUsageDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureUsageDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureUsageDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureUsageDetailResponseData(val *FeatureUsageDetailResponseData) *NullableFeatureUsageDetailResponseData {
	return &NullableFeatureUsageDetailResponseData{value: val, isSet: true}
}

func (v NullableFeatureUsageDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureUsageDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
