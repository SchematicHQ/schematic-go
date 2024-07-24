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

// checks if the CheckFlagOutputWithFlagKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckFlagOutputWithFlagKey{}

// CheckFlagOutputWithFlagKey struct for CheckFlagOutputWithFlagKey
type CheckFlagOutputWithFlagKey struct {
	CompanyId            NullableString `json:"company_id,omitempty"`
	Error                NullableString `json:"error,omitempty"`
	Flag                 string         `json:"flag"`
	Reason               string         `json:"reason"`
	RuleId               NullableString `json:"rule_id,omitempty"`
	UserId               NullableString `json:"user_id,omitempty"`
	Value                bool           `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CheckFlagOutputWithFlagKey CheckFlagOutputWithFlagKey

// NewCheckFlagOutputWithFlagKey instantiates a new CheckFlagOutputWithFlagKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckFlagOutputWithFlagKey(flag string, reason string, value bool) *CheckFlagOutputWithFlagKey {
	this := CheckFlagOutputWithFlagKey{}
	this.Flag = flag
	this.Reason = reason
	this.Value = value
	return &this
}

// NewCheckFlagOutputWithFlagKeyWithDefaults instantiates a new CheckFlagOutputWithFlagKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckFlagOutputWithFlagKeyWithDefaults() *CheckFlagOutputWithFlagKey {
	this := CheckFlagOutputWithFlagKey{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckFlagOutputWithFlagKey) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckFlagOutputWithFlagKey) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CheckFlagOutputWithFlagKey) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *CheckFlagOutputWithFlagKey) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *CheckFlagOutputWithFlagKey) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *CheckFlagOutputWithFlagKey) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckFlagOutputWithFlagKey) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckFlagOutputWithFlagKey) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CheckFlagOutputWithFlagKey) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *CheckFlagOutputWithFlagKey) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *CheckFlagOutputWithFlagKey) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *CheckFlagOutputWithFlagKey) UnsetError() {
	o.Error.Unset()
}

// GetFlag returns the Flag field value
func (o *CheckFlagOutputWithFlagKey) GetFlag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Flag
}

// GetFlagOk returns a tuple with the Flag field value
// and a boolean to check if the value has been set.
func (o *CheckFlagOutputWithFlagKey) GetFlagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flag, true
}

// SetFlag sets field value
func (o *CheckFlagOutputWithFlagKey) SetFlag(v string) {
	o.Flag = v
}

// GetReason returns the Reason field value
func (o *CheckFlagOutputWithFlagKey) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *CheckFlagOutputWithFlagKey) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *CheckFlagOutputWithFlagKey) SetReason(v string) {
	o.Reason = v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckFlagOutputWithFlagKey) GetRuleId() string {
	if o == nil || IsNil(o.RuleId.Get()) {
		var ret string
		return ret
	}
	return *o.RuleId.Get()
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckFlagOutputWithFlagKey) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleId.Get(), o.RuleId.IsSet()
}

// HasRuleId returns a boolean if a field has been set.
func (o *CheckFlagOutputWithFlagKey) HasRuleId() bool {
	if o != nil && o.RuleId.IsSet() {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given NullableString and assigns it to the RuleId field.
func (o *CheckFlagOutputWithFlagKey) SetRuleId(v string) {
	o.RuleId.Set(&v)
}

// SetRuleIdNil sets the value for RuleId to be an explicit nil
func (o *CheckFlagOutputWithFlagKey) SetRuleIdNil() {
	o.RuleId.Set(nil)
}

// UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
func (o *CheckFlagOutputWithFlagKey) UnsetRuleId() {
	o.RuleId.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckFlagOutputWithFlagKey) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckFlagOutputWithFlagKey) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CheckFlagOutputWithFlagKey) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *CheckFlagOutputWithFlagKey) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CheckFlagOutputWithFlagKey) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CheckFlagOutputWithFlagKey) UnsetUserId() {
	o.UserId.Unset()
}

// GetValue returns the Value field value
func (o *CheckFlagOutputWithFlagKey) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CheckFlagOutputWithFlagKey) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CheckFlagOutputWithFlagKey) SetValue(v bool) {
	o.Value = v
}

func (o CheckFlagOutputWithFlagKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckFlagOutputWithFlagKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["flag"] = o.Flag
	toSerialize["reason"] = o.Reason
	if o.RuleId.IsSet() {
		toSerialize["rule_id"] = o.RuleId.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckFlagOutputWithFlagKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flag",
		"reason",
		"value",
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

	varCheckFlagOutputWithFlagKey := _CheckFlagOutputWithFlagKey{}

	err = json.Unmarshal(data, &varCheckFlagOutputWithFlagKey)

	if err != nil {
		return err
	}

	*o = CheckFlagOutputWithFlagKey(varCheckFlagOutputWithFlagKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "error")
		delete(additionalProperties, "flag")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckFlagOutputWithFlagKey struct {
	value *CheckFlagOutputWithFlagKey
	isSet bool
}

func (v NullableCheckFlagOutputWithFlagKey) Get() *CheckFlagOutputWithFlagKey {
	return v.value
}

func (v *NullableCheckFlagOutputWithFlagKey) Set(val *CheckFlagOutputWithFlagKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckFlagOutputWithFlagKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckFlagOutputWithFlagKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckFlagOutputWithFlagKey(val *CheckFlagOutputWithFlagKey) *NullableCheckFlagOutputWithFlagKey {
	return &NullableCheckFlagOutputWithFlagKey{value: val, isSet: true}
}

func (v NullableCheckFlagOutputWithFlagKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckFlagOutputWithFlagKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
