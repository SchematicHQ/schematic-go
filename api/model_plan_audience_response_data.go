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
	"time"
)

// checks if the PlanAudienceResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanAudienceResponseData{}

// PlanAudienceResponseData struct for PlanAudienceResponseData
type PlanAudienceResponseData struct {
	CreatedAt            time.Time      `json:"created_at"`
	EnvironmentId        string         `json:"environment_id"`
	FlagId               NullableString `json:"flag_id,omitempty"`
	Id                   string         `json:"id"`
	Name                 string         `json:"name"`
	PlanId               NullableString `json:"plan_id,omitempty"`
	Priority             int32          `json:"priority"`
	RuleType             string         `json:"rule_type"`
	UpdatedAt            time.Time      `json:"updated_at"`
	Value                bool           `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _PlanAudienceResponseData PlanAudienceResponseData

// NewPlanAudienceResponseData instantiates a new PlanAudienceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanAudienceResponseData(createdAt time.Time, environmentId string, id string, name string, priority int32, ruleType string, updatedAt time.Time, value bool) *PlanAudienceResponseData {
	this := PlanAudienceResponseData{}
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.Id = id
	this.Name = name
	this.Priority = priority
	this.RuleType = ruleType
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewPlanAudienceResponseDataWithDefaults instantiates a new PlanAudienceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanAudienceResponseDataWithDefaults() *PlanAudienceResponseData {
	this := PlanAudienceResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *PlanAudienceResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PlanAudienceResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *PlanAudienceResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *PlanAudienceResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetFlagId returns the FlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanAudienceResponseData) GetFlagId() string {
	if o == nil || IsNil(o.FlagId.Get()) {
		var ret string
		return ret
	}
	return *o.FlagId.Get()
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanAudienceResponseData) GetFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagId.Get(), o.FlagId.IsSet()
}

// HasFlagId returns a boolean if a field has been set.
func (o *PlanAudienceResponseData) HasFlagId() bool {
	if o != nil && o.FlagId.IsSet() {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given NullableString and assigns it to the FlagId field.
func (o *PlanAudienceResponseData) SetFlagId(v string) {
	o.FlagId.Set(&v)
}

// SetFlagIdNil sets the value for FlagId to be an explicit nil
func (o *PlanAudienceResponseData) SetFlagIdNil() {
	o.FlagId.Set(nil)
}

// UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
func (o *PlanAudienceResponseData) UnsetFlagId() {
	o.FlagId.Unset()
}

// GetId returns the Id field value
func (o *PlanAudienceResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlanAudienceResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PlanAudienceResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanAudienceResponseData) SetName(v string) {
	o.Name = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanAudienceResponseData) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanAudienceResponseData) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *PlanAudienceResponseData) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *PlanAudienceResponseData) SetPlanId(v string) {
	o.PlanId.Set(&v)
}

// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *PlanAudienceResponseData) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *PlanAudienceResponseData) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetPriority returns the Priority field value
func (o *PlanAudienceResponseData) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *PlanAudienceResponseData) SetPriority(v int32) {
	o.Priority = v
}

// GetRuleType returns the RuleType field value
func (o *PlanAudienceResponseData) GetRuleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetRuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *PlanAudienceResponseData) SetRuleType(v string) {
	o.RuleType = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PlanAudienceResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PlanAudienceResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
func (o *PlanAudienceResponseData) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PlanAudienceResponseData) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PlanAudienceResponseData) SetValue(v bool) {
	o.Value = v
}

func (o PlanAudienceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanAudienceResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["rule_type"] = o.RuleType
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanAudienceResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"environment_id",
		"id",
		"name",
		"priority",
		"rule_type",
		"updated_at",
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

	varPlanAudienceResponseData := _PlanAudienceResponseData{}

	err = json.Unmarshal(data, &varPlanAudienceResponseData)

	if err != nil {
		return err
	}

	*o = PlanAudienceResponseData(varPlanAudienceResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "flag_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "rule_type")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanAudienceResponseData struct {
	value *PlanAudienceResponseData
	isSet bool
}

func (v NullablePlanAudienceResponseData) Get() *PlanAudienceResponseData {
	return v.value
}

func (v *NullablePlanAudienceResponseData) Set(val *PlanAudienceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanAudienceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanAudienceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanAudienceResponseData(val *PlanAudienceResponseData) *NullablePlanAudienceResponseData {
	return &NullablePlanAudienceResponseData{value: val, isSet: true}
}

func (v NullablePlanAudienceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanAudienceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
