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

// checks if the CreateOrUpdateRuleRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateRuleRequestBody{}

// CreateOrUpdateRuleRequestBody struct for CreateOrUpdateRuleRequestBody
type CreateOrUpdateRuleRequestBody struct {
	ConditionGroups      []CreateOrUpdateConditionGroupRequestBody `json:"condition_groups"`
	Conditions           []CreateOrUpdateConditionRequestBody      `json:"conditions"`
	Id                   NullableString                            `json:"id,omitempty"`
	Name                 string                                    `json:"name"`
	Priority             int32                                     `json:"priority"`
	RuleType             NullableString                            `json:"rule_type,omitempty"`
	Value                bool                                      `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CreateOrUpdateRuleRequestBody CreateOrUpdateRuleRequestBody

// NewCreateOrUpdateRuleRequestBody instantiates a new CreateOrUpdateRuleRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateRuleRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, name string, priority int32, value bool) *CreateOrUpdateRuleRequestBody {
	this := CreateOrUpdateRuleRequestBody{}
	this.ConditionGroups = conditionGroups
	this.Conditions = conditions
	this.Name = name
	this.Priority = priority
	this.Value = value
	return &this
}

// NewCreateOrUpdateRuleRequestBodyWithDefaults instantiates a new CreateOrUpdateRuleRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateRuleRequestBodyWithDefaults() *CreateOrUpdateRuleRequestBody {
	this := CreateOrUpdateRuleRequestBody{}
	return &this
}

// GetConditionGroups returns the ConditionGroups field value
func (o *CreateOrUpdateRuleRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody {
	if o == nil {
		var ret []CreateOrUpdateConditionGroupRequestBody
		return ret
	}

	return o.ConditionGroups
}

// GetConditionGroupsOk returns a tuple with the ConditionGroups field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRuleRequestBody) GetConditionGroupsOk() ([]CreateOrUpdateConditionGroupRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConditionGroups, true
}

// SetConditionGroups sets field value
func (o *CreateOrUpdateRuleRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody) {
	o.ConditionGroups = v
}

// GetConditions returns the Conditions field value
func (o *CreateOrUpdateRuleRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody {
	if o == nil {
		var ret []CreateOrUpdateConditionRequestBody
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRuleRequestBody) GetConditionsOk() ([]CreateOrUpdateConditionRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *CreateOrUpdateRuleRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody) {
	o.Conditions = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateRuleRequestBody) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateRuleRequestBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrUpdateRuleRequestBody) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateOrUpdateRuleRequestBody) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateOrUpdateRuleRequestBody) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateOrUpdateRuleRequestBody) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value
func (o *CreateOrUpdateRuleRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRuleRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateRuleRequestBody) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *CreateOrUpdateRuleRequestBody) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRuleRequestBody) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateOrUpdateRuleRequestBody) SetPriority(v int32) {
	o.Priority = v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateRuleRequestBody) GetRuleType() string {
	if o == nil || IsNil(o.RuleType.Get()) {
		var ret string
		return ret
	}
	return *o.RuleType.Get()
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateRuleRequestBody) GetRuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleType.Get(), o.RuleType.IsSet()
}

// HasRuleType returns a boolean if a field has been set.
func (o *CreateOrUpdateRuleRequestBody) HasRuleType() bool {
	if o != nil && o.RuleType.IsSet() {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given NullableString and assigns it to the RuleType field.
func (o *CreateOrUpdateRuleRequestBody) SetRuleType(v string) {
	o.RuleType.Set(&v)
}

// SetRuleTypeNil sets the value for RuleType to be an explicit nil
func (o *CreateOrUpdateRuleRequestBody) SetRuleTypeNil() {
	o.RuleType.Set(nil)
}

// UnsetRuleType ensures that no value is present for RuleType, not even an explicit nil
func (o *CreateOrUpdateRuleRequestBody) UnsetRuleType() {
	o.RuleType.Unset()
}

// GetValue returns the Value field value
func (o *CreateOrUpdateRuleRequestBody) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRuleRequestBody) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreateOrUpdateRuleRequestBody) SetValue(v bool) {
	o.Value = v
}

func (o CreateOrUpdateRuleRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateRuleRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition_groups"] = o.ConditionGroups
	toSerialize["conditions"] = o.Conditions
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	if o.RuleType.IsSet() {
		toSerialize["rule_type"] = o.RuleType.Get()
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOrUpdateRuleRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varCreateOrUpdateRuleRequestBody := _CreateOrUpdateRuleRequestBody{}

	err = json.Unmarshal(data, &varCreateOrUpdateRuleRequestBody)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateRuleRequestBody(varCreateOrUpdateRuleRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "condition_groups")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "rule_type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOrUpdateRuleRequestBody struct {
	value *CreateOrUpdateRuleRequestBody
	isSet bool
}

func (v NullableCreateOrUpdateRuleRequestBody) Get() *CreateOrUpdateRuleRequestBody {
	return v.value
}

func (v *NullableCreateOrUpdateRuleRequestBody) Set(val *CreateOrUpdateRuleRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateRuleRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateRuleRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateRuleRequestBody(val *CreateOrUpdateRuleRequestBody) *NullableCreateOrUpdateRuleRequestBody {
	return &NullableCreateOrUpdateRuleRequestBody{value: val, isSet: true}
}

func (v NullableCreateOrUpdateRuleRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateRuleRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
