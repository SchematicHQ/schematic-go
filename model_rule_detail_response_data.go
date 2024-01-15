/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
	"time"
)

// checks if the RuleDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleDetailResponseData{}

// RuleDetailResponseData The updated resource
type RuleDetailResponseData struct {
	ConditionGroups []RuleConditionGroupDetailResponseData `json:"condition_groups"`
	Conditions []RuleConditionResponseData `json:"conditions"`
	CreatedAt time.Time `json:"created_at"`
	EnvironmentId string `json:"environment_id"`
	FlagId NullableString `json:"flag_id,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	PlanId NullableString `json:"plan_id,omitempty"`
	Priority int32 `json:"priority"`
	PriorityGroup NullableInt32 `json:"priority_group,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
	Value bool `json:"value"`
}

// NewRuleDetailResponseData instantiates a new RuleDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDetailResponseData(conditionGroups []RuleConditionGroupDetailResponseData, conditions []RuleConditionResponseData, createdAt time.Time, environmentId string, id string, name string, priority int32, updatedAt time.Time, value bool) *RuleDetailResponseData {
	this := RuleDetailResponseData{}
	this.ConditionGroups = conditionGroups
	this.Conditions = conditions
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.Id = id
	this.Name = name
	this.Priority = priority
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewRuleDetailResponseDataWithDefaults instantiates a new RuleDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDetailResponseDataWithDefaults() *RuleDetailResponseData {
	this := RuleDetailResponseData{}
	return &this
}

// GetConditionGroups returns the ConditionGroups field value
func (o *RuleDetailResponseData) GetConditionGroups() []RuleConditionGroupDetailResponseData {
	if o == nil {
		var ret []RuleConditionGroupDetailResponseData
		return ret
	}

	return o.ConditionGroups
}

// GetConditionGroupsOk returns a tuple with the ConditionGroups field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetConditionGroupsOk() ([]RuleConditionGroupDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConditionGroups, true
}

// SetConditionGroups sets field value
func (o *RuleDetailResponseData) SetConditionGroups(v []RuleConditionGroupDetailResponseData) {
	o.ConditionGroups = v
}

// GetConditions returns the Conditions field value
func (o *RuleDetailResponseData) GetConditions() []RuleConditionResponseData {
	if o == nil {
		var ret []RuleConditionResponseData
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetConditionsOk() ([]RuleConditionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *RuleDetailResponseData) SetConditions(v []RuleConditionResponseData) {
	o.Conditions = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RuleDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RuleDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *RuleDetailResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *RuleDetailResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetFlagId returns the FlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleDetailResponseData) GetFlagId() string {
	if o == nil || IsNil(o.FlagId.Get()) {
		var ret string
		return ret
	}
	return *o.FlagId.Get()
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleDetailResponseData) GetFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagId.Get(), o.FlagId.IsSet()
}

// HasFlagId returns a boolean if a field has been set.
func (o *RuleDetailResponseData) HasFlagId() bool {
	if o != nil && o.FlagId.IsSet() {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given NullableString and assigns it to the FlagId field.
func (o *RuleDetailResponseData) SetFlagId(v string) {
	o.FlagId.Set(&v)
}
// SetFlagIdNil sets the value for FlagId to be an explicit nil
func (o *RuleDetailResponseData) SetFlagIdNil() {
	o.FlagId.Set(nil)
}

// UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
func (o *RuleDetailResponseData) UnsetFlagId() {
	o.FlagId.Unset()
}

// GetId returns the Id field value
func (o *RuleDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RuleDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RuleDetailResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleDetailResponseData) SetName(v string) {
	o.Name = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleDetailResponseData) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleDetailResponseData) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *RuleDetailResponseData) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *RuleDetailResponseData) SetPlanId(v string) {
	o.PlanId.Set(&v)
}
// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *RuleDetailResponseData) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *RuleDetailResponseData) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetPriority returns the Priority field value
func (o *RuleDetailResponseData) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *RuleDetailResponseData) SetPriority(v int32) {
	o.Priority = v
}

// GetPriorityGroup returns the PriorityGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleDetailResponseData) GetPriorityGroup() int32 {
	if o == nil || IsNil(o.PriorityGroup.Get()) {
		var ret int32
		return ret
	}
	return *o.PriorityGroup.Get()
}

// GetPriorityGroupOk returns a tuple with the PriorityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleDetailResponseData) GetPriorityGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityGroup.Get(), o.PriorityGroup.IsSet()
}

// HasPriorityGroup returns a boolean if a field has been set.
func (o *RuleDetailResponseData) HasPriorityGroup() bool {
	if o != nil && o.PriorityGroup.IsSet() {
		return true
	}

	return false
}

// SetPriorityGroup gets a reference to the given NullableInt32 and assigns it to the PriorityGroup field.
func (o *RuleDetailResponseData) SetPriorityGroup(v int32) {
	o.PriorityGroup.Set(&v)
}
// SetPriorityGroupNil sets the value for PriorityGroup to be an explicit nil
func (o *RuleDetailResponseData) SetPriorityGroupNil() {
	o.PriorityGroup.Set(nil)
}

// UnsetPriorityGroup ensures that no value is present for PriorityGroup, not even an explicit nil
func (o *RuleDetailResponseData) UnsetPriorityGroup() {
	o.PriorityGroup.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RuleDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RuleDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
func (o *RuleDetailResponseData) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RuleDetailResponseData) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RuleDetailResponseData) SetValue(v bool) {
	o.Value = v
}

func (o RuleDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition_groups"] = o.ConditionGroups
	toSerialize["conditions"] = o.Conditions
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment_id"] = o.EnvironmentId
	if o.FlagId.IsSet() {
		toSerialize["flag_id"] = o.FlagId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.PlanId.IsSet() {
		toSerialize["plan_id"] = o.PlanId.Get()
	}
	toSerialize["priority"] = o.Priority
	if o.PriorityGroup.IsSet() {
		toSerialize["priority_group"] = o.PriorityGroup.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableRuleDetailResponseData struct {
	value *RuleDetailResponseData
	isSet bool
}

func (v NullableRuleDetailResponseData) Get() *RuleDetailResponseData {
	return v.value
}

func (v *NullableRuleDetailResponseData) Set(val *RuleDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDetailResponseData(val *RuleDetailResponseData) *NullableRuleDetailResponseData {
	return &NullableRuleDetailResponseData{value: val, isSet: true}
}

func (v NullableRuleDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


