/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ComponentHydrateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentHydrateResponseData{}

// ComponentHydrateResponseData The returned resource
type ComponentHydrateResponseData struct {
	Company              *CompanyDetailResponseData      `json:"company,omitempty"`
	Component            *ComponentResponseData          `json:"component,omitempty"`
	FeatureUsage         *FeatureUsageDetailResponseData `json:"feature_usage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComponentHydrateResponseData ComponentHydrateResponseData

// NewComponentHydrateResponseData instantiates a new ComponentHydrateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentHydrateResponseData() *ComponentHydrateResponseData {
	this := ComponentHydrateResponseData{}
	return &this
}

// NewComponentHydrateResponseDataWithDefaults instantiates a new ComponentHydrateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentHydrateResponseDataWithDefaults() *ComponentHydrateResponseData {
	this := ComponentHydrateResponseData{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ComponentHydrateResponseData) GetCompany() CompanyDetailResponseData {
	if o == nil || IsNil(o.Company) {
		var ret CompanyDetailResponseData
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentHydrateResponseData) GetCompanyOk() (*CompanyDetailResponseData, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ComponentHydrateResponseData) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyDetailResponseData and assigns it to the Company field.
func (o *ComponentHydrateResponseData) SetCompany(v CompanyDetailResponseData) {
	o.Company = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ComponentHydrateResponseData) GetComponent() ComponentResponseData {
	if o == nil || IsNil(o.Component) {
		var ret ComponentResponseData
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentHydrateResponseData) GetComponentOk() (*ComponentResponseData, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentHydrateResponseData) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ComponentResponseData and assigns it to the Component field.
func (o *ComponentHydrateResponseData) SetComponent(v ComponentResponseData) {
	o.Component = &v
}

// GetFeatureUsage returns the FeatureUsage field value if set, zero value otherwise.
func (o *ComponentHydrateResponseData) GetFeatureUsage() FeatureUsageDetailResponseData {
	if o == nil || IsNil(o.FeatureUsage) {
		var ret FeatureUsageDetailResponseData
		return ret
	}
	return *o.FeatureUsage
}

// GetFeatureUsageOk returns a tuple with the FeatureUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentHydrateResponseData) GetFeatureUsageOk() (*FeatureUsageDetailResponseData, bool) {
	if o == nil || IsNil(o.FeatureUsage) {
		return nil, false
	}
	return o.FeatureUsage, true
}

// HasFeatureUsage returns a boolean if a field has been set.
func (o *ComponentHydrateResponseData) HasFeatureUsage() bool {
	if o != nil && !IsNil(o.FeatureUsage) {
		return true
	}

	return false
}

// SetFeatureUsage gets a reference to the given FeatureUsageDetailResponseData and assigns it to the FeatureUsage field.
func (o *ComponentHydrateResponseData) SetFeatureUsage(v FeatureUsageDetailResponseData) {
	o.FeatureUsage = &v
}

func (o ComponentHydrateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentHydrateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.FeatureUsage) {
		toSerialize["feature_usage"] = o.FeatureUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ComponentHydrateResponseData) UnmarshalJSON(data []byte) (err error) {
	varComponentHydrateResponseData := _ComponentHydrateResponseData{}

	err = json.Unmarshal(data, &varComponentHydrateResponseData)

	if err != nil {
		return err
	}

	*o = ComponentHydrateResponseData(varComponentHydrateResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company")
		delete(additionalProperties, "component")
		delete(additionalProperties, "feature_usage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComponentHydrateResponseData struct {
	value *ComponentHydrateResponseData
	isSet bool
}

func (v NullableComponentHydrateResponseData) Get() *ComponentHydrateResponseData {
	return v.value
}

func (v *NullableComponentHydrateResponseData) Set(val *ComponentHydrateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentHydrateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentHydrateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentHydrateResponseData(val *ComponentHydrateResponseData) *NullableComponentHydrateResponseData {
	return &NullableComponentHydrateResponseData{value: val, isSet: true}
}

func (v NullableComponentHydrateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentHydrateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
