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
	"time"
)

// checks if the CompanyOverrideResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyOverrideResponseData{}

// CompanyOverrideResponseData The updated resource
type CompanyOverrideResponseData struct {
	Company       *CompanyDetailResponseData         `json:"company,omitempty"`
	CompanyId     string                             `json:"company_id"`
	CreatedAt     time.Time                          `json:"created_at"`
	EnvironmentId string                             `json:"environment_id"`
	Feature       *FeatureResponseData               `json:"feature,omitempty"`
	FeatureId     string                             `json:"feature_id"`
	Id            string                             `json:"id"`
	MetricPeriod  NullableString                     `json:"metric_period,omitempty"`
	RuleId        string                             `json:"rule_id"`
	UpdatedAt     time.Time                          `json:"updated_at"`
	ValueBool     NullableBool                       `json:"value_bool,omitempty"`
	ValueNumeric  NullableInt32                      `json:"value_numeric,omitempty"`
	ValueTrait    *EntityTraitDefinitionResponseData `json:"value_trait,omitempty"`
	ValueTraitId  NullableString                     `json:"value_trait_id,omitempty"`
	ValueType     string                             `json:"value_type"`
}

type _CompanyOverrideResponseData CompanyOverrideResponseData

// NewCompanyOverrideResponseData instantiates a new CompanyOverrideResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyOverrideResponseData(companyId string, createdAt time.Time, environmentId string, featureId string, id string, ruleId string, updatedAt time.Time, valueType string) *CompanyOverrideResponseData {
	this := CompanyOverrideResponseData{}
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.FeatureId = featureId
	this.Id = id
	this.RuleId = ruleId
	this.UpdatedAt = updatedAt
	this.ValueType = valueType
	return &this
}

// NewCompanyOverrideResponseDataWithDefaults instantiates a new CompanyOverrideResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyOverrideResponseDataWithDefaults() *CompanyOverrideResponseData {
	this := CompanyOverrideResponseData{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CompanyOverrideResponseData) GetCompany() CompanyDetailResponseData {
	if o == nil || IsNil(o.Company) {
		var ret CompanyDetailResponseData
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetCompanyOk() (*CompanyDetailResponseData, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyDetailResponseData and assigns it to the Company field.
func (o *CompanyOverrideResponseData) SetCompany(v CompanyDetailResponseData) {
	o.Company = &v
}

// GetCompanyId returns the CompanyId field value
func (o *CompanyOverrideResponseData) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *CompanyOverrideResponseData) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CompanyOverrideResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CompanyOverrideResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *CompanyOverrideResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *CompanyOverrideResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *CompanyOverrideResponseData) GetFeature() FeatureResponseData {
	if o == nil || IsNil(o.Feature) {
		var ret FeatureResponseData
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetFeatureOk() (*FeatureResponseData, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given FeatureResponseData and assigns it to the Feature field.
func (o *CompanyOverrideResponseData) SetFeature(v FeatureResponseData) {
	o.Feature = &v
}

// GetFeatureId returns the FeatureId field value
func (o *CompanyOverrideResponseData) GetFeatureId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureId, true
}

// SetFeatureId sets field value
func (o *CompanyOverrideResponseData) SetFeatureId(v string) {
	o.FeatureId = v
}

// GetId returns the Id field value
func (o *CompanyOverrideResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CompanyOverrideResponseData) SetId(v string) {
	o.Id = v
}

// GetMetricPeriod returns the MetricPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyOverrideResponseData) GetMetricPeriod() string {
	if o == nil || IsNil(o.MetricPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.MetricPeriod.Get()
}

// GetMetricPeriodOk returns a tuple with the MetricPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyOverrideResponseData) GetMetricPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricPeriod.Get(), o.MetricPeriod.IsSet()
}

// HasMetricPeriod returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasMetricPeriod() bool {
	if o != nil && o.MetricPeriod.IsSet() {
		return true
	}

	return false
}

// SetMetricPeriod gets a reference to the given NullableString and assigns it to the MetricPeriod field.
func (o *CompanyOverrideResponseData) SetMetricPeriod(v string) {
	o.MetricPeriod.Set(&v)
}

// SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil
func (o *CompanyOverrideResponseData) SetMetricPeriodNil() {
	o.MetricPeriod.Set(nil)
}

// UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
func (o *CompanyOverrideResponseData) UnsetMetricPeriod() {
	o.MetricPeriod.Unset()
}

// GetRuleId returns the RuleId field value
func (o *CompanyOverrideResponseData) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *CompanyOverrideResponseData) SetRuleId(v string) {
	o.RuleId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CompanyOverrideResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CompanyOverrideResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValueBool returns the ValueBool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyOverrideResponseData) GetValueBool() bool {
	if o == nil || IsNil(o.ValueBool.Get()) {
		var ret bool
		return ret
	}
	return *o.ValueBool.Get()
}

// GetValueBoolOk returns a tuple with the ValueBool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyOverrideResponseData) GetValueBoolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueBool.Get(), o.ValueBool.IsSet()
}

// HasValueBool returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasValueBool() bool {
	if o != nil && o.ValueBool.IsSet() {
		return true
	}

	return false
}

// SetValueBool gets a reference to the given NullableBool and assigns it to the ValueBool field.
func (o *CompanyOverrideResponseData) SetValueBool(v bool) {
	o.ValueBool.Set(&v)
}

// SetValueBoolNil sets the value for ValueBool to be an explicit nil
func (o *CompanyOverrideResponseData) SetValueBoolNil() {
	o.ValueBool.Set(nil)
}

// UnsetValueBool ensures that no value is present for ValueBool, not even an explicit nil
func (o *CompanyOverrideResponseData) UnsetValueBool() {
	o.ValueBool.Unset()
}

// GetValueNumeric returns the ValueNumeric field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyOverrideResponseData) GetValueNumeric() int32 {
	if o == nil || IsNil(o.ValueNumeric.Get()) {
		var ret int32
		return ret
	}
	return *o.ValueNumeric.Get()
}

// GetValueNumericOk returns a tuple with the ValueNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyOverrideResponseData) GetValueNumericOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueNumeric.Get(), o.ValueNumeric.IsSet()
}

// HasValueNumeric returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasValueNumeric() bool {
	if o != nil && o.ValueNumeric.IsSet() {
		return true
	}

	return false
}

// SetValueNumeric gets a reference to the given NullableInt32 and assigns it to the ValueNumeric field.
func (o *CompanyOverrideResponseData) SetValueNumeric(v int32) {
	o.ValueNumeric.Set(&v)
}

// SetValueNumericNil sets the value for ValueNumeric to be an explicit nil
func (o *CompanyOverrideResponseData) SetValueNumericNil() {
	o.ValueNumeric.Set(nil)
}

// UnsetValueNumeric ensures that no value is present for ValueNumeric, not even an explicit nil
func (o *CompanyOverrideResponseData) UnsetValueNumeric() {
	o.ValueNumeric.Unset()
}

// GetValueTrait returns the ValueTrait field value if set, zero value otherwise.
func (o *CompanyOverrideResponseData) GetValueTrait() EntityTraitDefinitionResponseData {
	if o == nil || IsNil(o.ValueTrait) {
		var ret EntityTraitDefinitionResponseData
		return ret
	}
	return *o.ValueTrait
}

// GetValueTraitOk returns a tuple with the ValueTrait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetValueTraitOk() (*EntityTraitDefinitionResponseData, bool) {
	if o == nil || IsNil(o.ValueTrait) {
		return nil, false
	}
	return o.ValueTrait, true
}

// HasValueTrait returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasValueTrait() bool {
	if o != nil && !IsNil(o.ValueTrait) {
		return true
	}

	return false
}

// SetValueTrait gets a reference to the given EntityTraitDefinitionResponseData and assigns it to the ValueTrait field.
func (o *CompanyOverrideResponseData) SetValueTrait(v EntityTraitDefinitionResponseData) {
	o.ValueTrait = &v
}

// GetValueTraitId returns the ValueTraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyOverrideResponseData) GetValueTraitId() string {
	if o == nil || IsNil(o.ValueTraitId.Get()) {
		var ret string
		return ret
	}
	return *o.ValueTraitId.Get()
}

// GetValueTraitIdOk returns a tuple with the ValueTraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyOverrideResponseData) GetValueTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueTraitId.Get(), o.ValueTraitId.IsSet()
}

// HasValueTraitId returns a boolean if a field has been set.
func (o *CompanyOverrideResponseData) HasValueTraitId() bool {
	if o != nil && o.ValueTraitId.IsSet() {
		return true
	}

	return false
}

// SetValueTraitId gets a reference to the given NullableString and assigns it to the ValueTraitId field.
func (o *CompanyOverrideResponseData) SetValueTraitId(v string) {
	o.ValueTraitId.Set(&v)
}

// SetValueTraitIdNil sets the value for ValueTraitId to be an explicit nil
func (o *CompanyOverrideResponseData) SetValueTraitIdNil() {
	o.ValueTraitId.Set(nil)
}

// UnsetValueTraitId ensures that no value is present for ValueTraitId, not even an explicit nil
func (o *CompanyOverrideResponseData) UnsetValueTraitId() {
	o.ValueTraitId.Unset()
}

// GetValueType returns the ValueType field value
func (o *CompanyOverrideResponseData) GetValueType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *CompanyOverrideResponseData) GetValueTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *CompanyOverrideResponseData) SetValueType(v string) {
	o.ValueType = v
}

func (o CompanyOverrideResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyOverrideResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["company_id"] = o.CompanyId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment_id"] = o.EnvironmentId
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	toSerialize["feature_id"] = o.FeatureId
	toSerialize["id"] = o.Id
	if o.MetricPeriod.IsSet() {
		toSerialize["metric_period"] = o.MetricPeriod.Get()
	}
	toSerialize["rule_id"] = o.RuleId
	toSerialize["updated_at"] = o.UpdatedAt
	if o.ValueBool.IsSet() {
		toSerialize["value_bool"] = o.ValueBool.Get()
	}
	if o.ValueNumeric.IsSet() {
		toSerialize["value_numeric"] = o.ValueNumeric.Get()
	}
	if !IsNil(o.ValueTrait) {
		toSerialize["value_trait"] = o.ValueTrait
	}
	if o.ValueTraitId.IsSet() {
		toSerialize["value_trait_id"] = o.ValueTraitId.Get()
	}
	toSerialize["value_type"] = o.ValueType
	return toSerialize, nil
}

func (o *CompanyOverrideResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company_id",
		"created_at",
		"environment_id",
		"feature_id",
		"id",
		"rule_id",
		"updated_at",
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

	varCompanyOverrideResponseData := _CompanyOverrideResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyOverrideResponseData)

	if err != nil {
		return err
	}

	*o = CompanyOverrideResponseData(varCompanyOverrideResponseData)

	return err
}

type NullableCompanyOverrideResponseData struct {
	value *CompanyOverrideResponseData
	isSet bool
}

func (v NullableCompanyOverrideResponseData) Get() *CompanyOverrideResponseData {
	return v.value
}

func (v *NullableCompanyOverrideResponseData) Set(val *CompanyOverrideResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyOverrideResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyOverrideResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyOverrideResponseData(val *CompanyOverrideResponseData) *NullableCompanyOverrideResponseData {
	return &NullableCompanyOverrideResponseData{value: val, isSet: true}
}

func (v NullableCompanyOverrideResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyOverrideResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
