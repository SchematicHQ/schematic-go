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

// checks if the CreateCompanyOverrideRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCompanyOverrideRequestBody{}

// CreateCompanyOverrideRequestBody struct for CreateCompanyOverrideRequestBody
type CreateCompanyOverrideRequestBody struct {
	CompanyId            string         `json:"company_id"`
	FeatureId            string         `json:"feature_id"`
	MetricPeriod         NullableString `json:"metric_period,omitempty"`
	ValueBool            NullableBool   `json:"value_bool,omitempty"`
	ValueNumeric         NullableInt32  `json:"value_numeric,omitempty"`
	ValueTraitId         NullableString `json:"value_trait_id,omitempty"`
	ValueType            string         `json:"value_type"`
	AdditionalProperties map[string]interface{}
}

type _CreateCompanyOverrideRequestBody CreateCompanyOverrideRequestBody

// NewCreateCompanyOverrideRequestBody instantiates a new CreateCompanyOverrideRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompanyOverrideRequestBody(companyId string, featureId string, valueType string) *CreateCompanyOverrideRequestBody {
	this := CreateCompanyOverrideRequestBody{}
	this.CompanyId = companyId
	this.FeatureId = featureId
	this.ValueType = valueType
	return &this
}

// NewCreateCompanyOverrideRequestBodyWithDefaults instantiates a new CreateCompanyOverrideRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompanyOverrideRequestBodyWithDefaults() *CreateCompanyOverrideRequestBody {
	this := CreateCompanyOverrideRequestBody{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *CreateCompanyOverrideRequestBody) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyOverrideRequestBody) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *CreateCompanyOverrideRequestBody) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetFeatureId returns the FeatureId field value
func (o *CreateCompanyOverrideRequestBody) GetFeatureId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyOverrideRequestBody) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureId, true
}

// SetFeatureId sets field value
func (o *CreateCompanyOverrideRequestBody) SetFeatureId(v string) {
	o.FeatureId = v
}

// GetMetricPeriod returns the MetricPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCompanyOverrideRequestBody) GetMetricPeriod() string {
	if o == nil || IsNil(o.MetricPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.MetricPeriod.Get()
}

// GetMetricPeriodOk returns a tuple with the MetricPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCompanyOverrideRequestBody) GetMetricPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricPeriod.Get(), o.MetricPeriod.IsSet()
}

// HasMetricPeriod returns a boolean if a field has been set.
func (o *CreateCompanyOverrideRequestBody) HasMetricPeriod() bool {
	if o != nil && o.MetricPeriod.IsSet() {
		return true
	}

	return false
}

// SetMetricPeriod gets a reference to the given NullableString and assigns it to the MetricPeriod field.
func (o *CreateCompanyOverrideRequestBody) SetMetricPeriod(v string) {
	o.MetricPeriod.Set(&v)
}

// SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil
func (o *CreateCompanyOverrideRequestBody) SetMetricPeriodNil() {
	o.MetricPeriod.Set(nil)
}

// UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
func (o *CreateCompanyOverrideRequestBody) UnsetMetricPeriod() {
	o.MetricPeriod.Unset()
}

// GetValueBool returns the ValueBool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCompanyOverrideRequestBody) GetValueBool() bool {
	if o == nil || IsNil(o.ValueBool.Get()) {
		var ret bool
		return ret
	}
	return *o.ValueBool.Get()
}

// GetValueBoolOk returns a tuple with the ValueBool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCompanyOverrideRequestBody) GetValueBoolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueBool.Get(), o.ValueBool.IsSet()
}

// HasValueBool returns a boolean if a field has been set.
func (o *CreateCompanyOverrideRequestBody) HasValueBool() bool {
	if o != nil && o.ValueBool.IsSet() {
		return true
	}

	return false
}

// SetValueBool gets a reference to the given NullableBool and assigns it to the ValueBool field.
func (o *CreateCompanyOverrideRequestBody) SetValueBool(v bool) {
	o.ValueBool.Set(&v)
}

// SetValueBoolNil sets the value for ValueBool to be an explicit nil
func (o *CreateCompanyOverrideRequestBody) SetValueBoolNil() {
	o.ValueBool.Set(nil)
}

// UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
func (o *CreateCompanyOverrideRequestBody) UnsetValueBool() {
	o.ValueBool.Unset()
}

// GetValueNumeric returns the ValueNumeric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCompanyOverrideRequestBody) GetValueNumeric() int32 {
	if o == nil || IsNil(o.ValueNumeric.Get()) {
		var ret int32
		return ret
	}
	return *o.ValueNumeric.Get()
}

// GetValueNumericOk returns a tuple with the ValueNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCompanyOverrideRequestBody) GetValueNumericOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueNumeric.Get(), o.ValueNumeric.IsSet()
}

// HasValueNumeric returns a boolean if a field has been set.
func (o *CreateCompanyOverrideRequestBody) HasValueNumeric() bool {
	if o != nil && o.ValueNumeric.IsSet() {
		return true
	}

	return false
}

// SetValueNumeric gets a reference to the given NullableInt32 and assigns it to the ValueNumeric field.
func (o *CreateCompanyOverrideRequestBody) SetValueNumeric(v int32) {
	o.ValueNumeric.Set(&v)
}

// SetValueNumericNil sets the value for ValueNumeric to be an explicit nil
func (o *CreateCompanyOverrideRequestBody) SetValueNumericNil() {
	o.ValueNumeric.Set(nil)
}

// UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
func (o *CreateCompanyOverrideRequestBody) UnsetValueNumeric() {
	o.ValueNumeric.Unset()
}

// GetValueTraitId returns the ValueTraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCompanyOverrideRequestBody) GetValueTraitId() string {
	if o == nil || IsNil(o.ValueTraitId.Get()) {
		var ret string
		return ret
	}
	return *o.ValueTraitId.Get()
}

// GetValueTraitIdOk returns a tuple with the ValueTraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCompanyOverrideRequestBody) GetValueTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueTraitId.Get(), o.ValueTraitId.IsSet()
}

// HasValueTraitId returns a boolean if a field has been set.
func (o *CreateCompanyOverrideRequestBody) HasValueTraitId() bool {
	if o != nil && o.ValueTraitId.IsSet() {
		return true
	}

	return false
}

// SetValueTraitId gets a reference to the given NullableString and assigns it to the ValueTraitId field.
func (o *CreateCompanyOverrideRequestBody) SetValueTraitId(v string) {
	o.ValueTraitId.Set(&v)
}

// SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil
func (o *CreateCompanyOverrideRequestBody) SetValueTraitIdNil() {
	o.ValueTraitId.Set(nil)
}

// UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
func (o *CreateCompanyOverrideRequestBody) UnsetValueTraitId() {
	o.ValueTraitId.Unset()
}

// GetValueType returns the ValueType field value
func (o *CreateCompanyOverrideRequestBody) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyOverrideRequestBody) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *CreateCompanyOverrideRequestBody) SetValueType(v string) {
	o.ValueType = v
}

func (o CreateCompanyOverrideRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompanyOverrideRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company_id"] = o.CompanyId
	toSerialize["feature_id"] = o.FeatureId
	if o.MetricPeriod.IsSet() {
		toSerialize["metric_period"] = o.MetricPeriod.Get()
	}
	if o.ValueBool.IsSet() {
		toSerialize["value_bool"] = o.ValueBool.Get()
	}
	if o.ValueNumeric.IsSet() {
		toSerialize["value_numeric"] = o.ValueNumeric.Get()
	}
	if o.ValueTraitId.IsSet() {
		toSerialize["value_trait_id"] = o.ValueTraitId.Get()
	}
	toSerialize["value_type"] = o.ValueType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCompanyOverrideRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company_id",
		"feature_id",
		"value_type",
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

	varCreateCompanyOverrideRequestBody := _CreateCompanyOverrideRequestBody{}

	err = json.Unmarshal(data, &varCreateCompanyOverrideRequestBody)

	if err != nil {
		return err
	}

	*o = CreateCompanyOverrideRequestBody(varCreateCompanyOverrideRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "feature_id")
		delete(additionalProperties, "metric_period")
		delete(additionalProperties, "value_bool")
		delete(additionalProperties, "value_numeric")
		delete(additionalProperties, "value_trait_id")
		delete(additionalProperties, "value_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCompanyOverrideRequestBody struct {
	value *CreateCompanyOverrideRequestBody
	isSet bool
}

func (v NullableCreateCompanyOverrideRequestBody) Get() *CreateCompanyOverrideRequestBody {
	return v.value
}

func (v *NullableCreateCompanyOverrideRequestBody) Set(val *CreateCompanyOverrideRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompanyOverrideRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompanyOverrideRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompanyOverrideRequestBody(val *CreateCompanyOverrideRequestBody) *NullableCreateCompanyOverrideRequestBody {
	return &NullableCreateCompanyOverrideRequestBody{value: val, isSet: true}
}

func (v NullableCreateCompanyOverrideRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompanyOverrideRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
