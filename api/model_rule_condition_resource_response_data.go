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

// checks if the RuleConditionResourceResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleConditionResourceResponseData{}

// RuleConditionResourceResponseData struct for RuleConditionResourceResponseData
type RuleConditionResourceResponseData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type _RuleConditionResourceResponseData RuleConditionResourceResponseData

// NewRuleConditionResourceResponseData instantiates a new RuleConditionResourceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConditionResourceResponseData(id string, name string) *RuleConditionResourceResponseData {
	this := RuleConditionResourceResponseData{}
	this.Id = id
	this.Name = name
	return &this
}

// NewRuleConditionResourceResponseDataWithDefaults instantiates a new RuleConditionResourceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConditionResourceResponseDataWithDefaults() *RuleConditionResourceResponseData {
	this := RuleConditionResourceResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *RuleConditionResourceResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleConditionResourceResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RuleConditionResourceResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RuleConditionResourceResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleConditionResourceResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleConditionResourceResponseData) SetName(v string) {
	o.Name = v
}

func (o RuleConditionResourceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleConditionResourceResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *RuleConditionResourceResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varRuleConditionResourceResponseData := _RuleConditionResourceResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleConditionResourceResponseData)

	if err != nil {
		return err
	}

	*o = RuleConditionResourceResponseData(varRuleConditionResourceResponseData)

	return err
}

type NullableRuleConditionResourceResponseData struct {
	value *RuleConditionResourceResponseData
	isSet bool
}

func (v NullableRuleConditionResourceResponseData) Get() *RuleConditionResourceResponseData {
	return v.value
}

func (v *NullableRuleConditionResourceResponseData) Set(val *RuleConditionResourceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConditionResourceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConditionResourceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConditionResourceResponseData(val *RuleConditionResourceResponseData) *NullableRuleConditionResourceResponseData {
	return &NullableRuleConditionResourceResponseData{value: val, isSet: true}
}

func (v NullableRuleConditionResourceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConditionResourceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}