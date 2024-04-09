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

// checks if the UpdateRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRuleRequestBody{}

// UpdateRuleRequestBody struct for UpdateRuleRequestBody
type UpdateRuleRequestBody struct {
	ConditionGroups []CreateOrUpdateConditionGroupRequestBody `json:"condition_groups"`
	Conditions      []CreateOrUpdateConditionRequestBody      `json:"conditions"`
	Name            string                                    `json:"name"`
	Priority        int32                                     `json:"priority"`
	Value           bool                                      `json:"value"`
}

type _UpdateRuleRequestBody UpdateRuleRequestBody

// NewUpdateRuleRequestBody instantiates a new UpdateRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRuleRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, name string, priority int32, value bool) *UpdateRuleRequestBody {
	this := UpdateRuleRequestBody{}
	this.ConditionGroups = conditionGroups
	this.Conditions = conditions
	this.Name = name
	this.Priority = priority
	this.Value = value
	return &this
}

// NewUpdateRuleRequestBodyWithDefaults instantiates a new UpdateRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRuleRequestBodyWithDefaults() *UpdateRuleRequestBody {
	this := UpdateRuleRequestBody{}
	return &this
}

// GetConditionGroups returns the ConditionGroups field value
func (o *UpdateRuleRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody {
	if o == nil {
		var ret []CreateOrUpdateConditionGroupRequestBody
		return ret
	}

	return o.ConditionGroups
}

// GetConditionGroupsOk returns a tuple with the ConditionGroups field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleRequestBody) GetConditionGroupsOk() ([]CreateOrUpdateConditionGroupRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConditionGroups, true
}

// SetConditionGroups sets field value
func (o *UpdateRuleRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody) {
	o.ConditionGroups = v
}

// GetConditions returns the Conditions field value
func (o *UpdateRuleRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody {
	if o == nil {
		var ret []CreateOrUpdateConditionRequestBody
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleRequestBody) GetConditionsOk() ([]CreateOrUpdateConditionRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *UpdateRuleRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody) {
	o.Conditions = v
}

// GetName returns the Name field value
func (o *UpdateRuleRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRuleRequestBody) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *UpdateRuleRequestBody) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleRequestBody) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *UpdateRuleRequestBody) SetPriority(v int32) {
	o.Priority = v
}

// GetValue returns the Value field value
func (o *UpdateRuleRequestBody) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UpdateRuleRequestBody) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UpdateRuleRequestBody) SetValue(v bool) {
	o.Value = v
}

func (o UpdateRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition_groups"] = o.ConditionGroups
	toSerialize["conditions"] = o.Conditions
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *UpdateRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"condition_groups",
		"conditions",
		"name",
		"priority",
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

	varUpdateRuleRequestBody := _UpdateRuleRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateRuleRequestBody)

	if err != nil {
		return err
	}

	*o = UpdateRuleRequestBody(varUpdateRuleRequestBody)

	return err
}

type NullableUpdateRuleRequestBody struct {
	value *UpdateRuleRequestBody
	isSet bool
}

func (v NullableUpdateRuleRequestBody) Get() *UpdateRuleRequestBody {
	return v.value
}

func (v *NullableUpdateRuleRequestBody) Set(val *UpdateRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRuleRequestBody(val *UpdateRuleRequestBody) *NullableUpdateRuleRequestBody {
	return &NullableUpdateRuleRequestBody{value: val, isSet: true}
}

func (v NullableUpdateRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
