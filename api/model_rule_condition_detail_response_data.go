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

// checks if the RuleConditionDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleConditionDetailResponseData{}

// RuleConditionDetailResponseData struct for RuleConditionDetailResponseData
type RuleConditionDetailResponseData struct {
	ComparisonTrait   *EntityTraitDefinitionResponseData  `json:"comparison_trait,omitempty"`
	ComparisonTraitId NullableString                      `json:"comparison_trait_id,omitempty"`
	ConditionGroupId  NullableString                      `json:"condition_group_id,omitempty"`
	ConditionType     string                              `json:"condition_type"`
	CreatedAt         time.Time                           `json:"created_at"`
	EnvironmentId     string                              `json:"environment_id"`
	EventSubtype      NullableString                      `json:"event_subtype,omitempty"`
	FlagId            NullableString                      `json:"flag_id,omitempty"`
	Id                string                              `json:"id"`
	MetricPeriod      NullableString                      `json:"metric_period,omitempty"`
	MetricValue       NullableInt32                       `json:"metric_value,omitempty"`
	Operator          string                              `json:"operator"`
	PlanId            NullableString                      `json:"plan_id,omitempty"`
	ResourceIds       []string                            `json:"resource_ids"`
	Resources         []RuleConditionResourceResponseData `json:"resources"`
	RuleId            string                              `json:"rule_id"`
	Trait             *EntityTraitDefinitionResponseData  `json:"trait,omitempty"`
	TraitEntityType   NullableString                      `json:"trait_entity_type,omitempty"`
	TraitId           NullableString                      `json:"trait_id,omitempty"`
	TraitValue        string                              `json:"trait_value"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

type _RuleConditionDetailResponseData RuleConditionDetailResponseData

// NewRuleConditionDetailResponseData instantiates a new RuleConditionDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleConditionDetailResponseData(conditionType string, createdAt time.Time, environmentId string, id string, operator string, resourceIds []string, resources []RuleConditionResourceResponseData, ruleId string, traitValue string, updatedAt time.Time) *RuleConditionDetailResponseData {
	this := RuleConditionDetailResponseData{}
	this.ConditionType = conditionType
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.Id = id
	this.Operator = operator
	this.ResourceIds = resourceIds
	this.Resources = resources
	this.RuleId = ruleId
	this.TraitValue = traitValue
	this.UpdatedAt = updatedAt
	return &this
}

// NewRuleConditionDetailResponseDataWithDefaults instantiates a new RuleConditionDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleConditionDetailResponseDataWithDefaults() *RuleConditionDetailResponseData {
	this := RuleConditionDetailResponseData{}
	return &this
}

// GetComparisonTrait returns the ComparisonTrait field value if set, zero value otherwise.
func (o *RuleConditionDetailResponseData) GetComparisonTrait() EntityTraitDefinitionResponseData {
	if o == nil || IsNil(o.ComparisonTrait) {
		var ret EntityTraitDefinitionResponseData
		return ret
	}
	return *o.ComparisonTrait
}

// GetComparisonTraitOk returns a tuple with the ComparisonTrait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetComparisonTraitOk() (*EntityTraitDefinitionResponseData, bool) {
	if o == nil || IsNil(o.ComparisonTrait) {
		return nil, false
	}
	return o.ComparisonTrait, true
}

// HasComparisonTrait returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasComparisonTrait() bool {
	if o != nil && !IsNil(o.ComparisonTrait) {
		return true
	}

	return false
}

// SetComparisonTrait gets a reference to the given EntityTraitDefinitionResponseData and assigns it to the ComparisonTrait field.
func (o *RuleConditionDetailResponseData) SetComparisonTrait(v EntityTraitDefinitionResponseData) {
	o.ComparisonTrait = &v
}

// GetComparisonTraitId returns the ComparisonTraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetComparisonTraitId() string {
	if o == nil || IsNil(o.ComparisonTraitId.Get()) {
		var ret string
		return ret
	}
	return *o.ComparisonTraitId.Get()
}

// GetComparisonTraitIdOk returns a tuple with the ComparisonTraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetComparisonTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonTraitId.Get(), o.ComparisonTraitId.IsSet()
}

// HasComparisonTraitId returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasComparisonTraitId() bool {
	if o != nil && o.ComparisonTraitId.IsSet() {
		return true
	}

	return false
}

// SetComparisonTraitId gets a reference to the given NullableString and assigns it to the ComparisonTraitId field.
func (o *RuleConditionDetailResponseData) SetComparisonTraitId(v string) {
	o.ComparisonTraitId.Set(&v)
}

// SetComparisonTraitIdNil sets the value for ComparisonTraitId to be an explicit nil
func (o *RuleConditionDetailResponseData) SetComparisonTraitIdNil() {
	o.ComparisonTraitId.Set(nil)
}

// UnsetComparisonTraitId ensures that no value is present for ComparisonTraitId, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetComparisonTraitId() {
	o.ComparisonTraitId.Unset()
}

// GetConditionGroupId returns the ConditionGroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetConditionGroupId() string {
	if o == nil || IsNil(o.ConditionGroupId.Get()) {
		var ret string
		return ret
	}
	return *o.ConditionGroupId.Get()
}

// GetConditionGroupIdOk returns a tuple with the ConditionGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetConditionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConditionGroupId.Get(), o.ConditionGroupId.IsSet()
}

// HasConditionGroupId returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasConditionGroupId() bool {
	if o != nil && o.ConditionGroupId.IsSet() {
		return true
	}

	return false
}

// SetConditionGroupId gets a reference to the given NullableString and assigns it to the ConditionGroupId field.
func (o *RuleConditionDetailResponseData) SetConditionGroupId(v string) {
	o.ConditionGroupId.Set(&v)
}

// SetConditionGroupIdNil sets the value for ConditionGroupId to be an explicit nil
func (o *RuleConditionDetailResponseData) SetConditionGroupIdNil() {
	o.ConditionGroupId.Set(nil)
}

// UnsetConditionGroupId ensures that no value is present for ConditionGroupId, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetConditionGroupId() {
	o.ConditionGroupId.Unset()
}

// GetConditionType returns the ConditionType field value
func (o *RuleConditionDetailResponseData) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *RuleConditionDetailResponseData) SetConditionType(v string) {
	o.ConditionType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RuleConditionDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RuleConditionDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *RuleConditionDetailResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *RuleConditionDetailResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.EventSubtype.Get()
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventSubtype.Get(), o.EventSubtype.IsSet()
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasEventSubtype() bool {
	if o != nil && o.EventSubtype.IsSet() {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given NullableString and assigns it to the EventSubtype field.
func (o *RuleConditionDetailResponseData) SetEventSubtype(v string) {
	o.EventSubtype.Set(&v)
}

// SetEventSubtypeNil sets the value for EventSubtype to be an explicit nil
func (o *RuleConditionDetailResponseData) SetEventSubtypeNil() {
	o.EventSubtype.Set(nil)
}

// UnsetEventSubtype ensures that no value is present for EventSubtype, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetEventSubtype() {
	o.EventSubtype.Unset()
}

// GetFlagId returns the FlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetFlagId() string {
	if o == nil || IsNil(o.FlagId.Get()) {
		var ret string
		return ret
	}
	return *o.FlagId.Get()
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagId.Get(), o.FlagId.IsSet()
}

// HasFlagId returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasFlagId() bool {
	if o != nil && o.FlagId.IsSet() {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given NullableString and assigns it to the FlagId field.
func (o *RuleConditionDetailResponseData) SetFlagId(v string) {
	o.FlagId.Set(&v)
}

// SetFlagIdNil sets the value for FlagId to be an explicit nil
func (o *RuleConditionDetailResponseData) SetFlagIdNil() {
	o.FlagId.Set(nil)
}

// UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetFlagId() {
	o.FlagId.Unset()
}

// GetId returns the Id field value
func (o *RuleConditionDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RuleConditionDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetMetricPeriod returns the MetricPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetMetricPeriod() string {
	if o == nil || IsNil(o.MetricPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.MetricPeriod.Get()
}

// GetMetricPeriodOk returns a tuple with the MetricPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetMetricPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricPeriod.Get(), o.MetricPeriod.IsSet()
}

// HasMetricPeriod returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasMetricPeriod() bool {
	if o != nil && o.MetricPeriod.IsSet() {
		return true
	}

	return false
}

// SetMetricPeriod gets a reference to the given NullableString and assigns it to the MetricPeriod field.
func (o *RuleConditionDetailResponseData) SetMetricPeriod(v string) {
	o.MetricPeriod.Set(&v)
}

// SetMetricPeriodNil sets the value for MetricPeriod to be an explicit nil
func (o *RuleConditionDetailResponseData) SetMetricPeriodNil() {
	o.MetricPeriod.Set(nil)
}

// UnsetMetricPeriod ensures that no value is present for MetricPeriod, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetMetricPeriod() {
	o.MetricPeriod.Unset()
}

// GetMetricValue returns the MetricValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetMetricValue() int32 {
	if o == nil || IsNil(o.MetricValue.Get()) {
		var ret int32
		return ret
	}
	return *o.MetricValue.Get()
}

// GetMetricValueOk returns a tuple with the MetricValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetMetricValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricValue.Get(), o.MetricValue.IsSet()
}

// HasMetricValue returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasMetricValue() bool {
	if o != nil && o.MetricValue.IsSet() {
		return true
	}

	return false
}

// SetMetricValue gets a reference to the given NullableInt32 and assigns it to the MetricValue field.
func (o *RuleConditionDetailResponseData) SetMetricValue(v int32) {
	o.MetricValue.Set(&v)
}

// SetMetricValueNil sets the value for MetricValue to be an explicit nil
func (o *RuleConditionDetailResponseData) SetMetricValueNil() {
	o.MetricValue.Set(nil)
}

// UnsetMetricValue ensures that no value is present for MetricValue, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetMetricValue() {
	o.MetricValue.Unset()
}

// GetOperator returns the Operator field value
func (o *RuleConditionDetailResponseData) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *RuleConditionDetailResponseData) SetOperator(v string) {
	o.Operator = v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *RuleConditionDetailResponseData) SetPlanId(v string) {
	o.PlanId.Set(&v)
}

// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *RuleConditionDetailResponseData) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetResourceIds returns the ResourceIds field value
func (o *RuleConditionDetailResponseData) GetResourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetResourceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceIds, true
}

// SetResourceIds sets field value
func (o *RuleConditionDetailResponseData) SetResourceIds(v []string) {
	o.ResourceIds = v
}

// GetResources returns the Resources field value
func (o *RuleConditionDetailResponseData) GetResources() []RuleConditionResourceResponseData {
	if o == nil {
		var ret []RuleConditionResourceResponseData
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetResourcesOk() ([]RuleConditionResourceResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *RuleConditionDetailResponseData) SetResources(v []RuleConditionResourceResponseData) {
	o.Resources = v
}

// GetRuleId returns the RuleId field value
func (o *RuleConditionDetailResponseData) GetRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleId, true
}

// SetRuleId sets field value
func (o *RuleConditionDetailResponseData) SetRuleId(v string) {
	o.RuleId = v
}

// GetTrait returns the Trait field value if set, zero value otherwise.
func (o *RuleConditionDetailResponseData) GetTrait() EntityTraitDefinitionResponseData {
	if o == nil || IsNil(o.Trait) {
		var ret EntityTraitDefinitionResponseData
		return ret
	}
	return *o.Trait
}

// GetTraitOk returns a tuple with the Trait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetTraitOk() (*EntityTraitDefinitionResponseData, bool) {
	if o == nil || IsNil(o.Trait) {
		return nil, false
	}
	return o.Trait, true
}

// HasTrait returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasTrait() bool {
	if o != nil && !IsNil(o.Trait) {
		return true
	}

	return false
}

// SetTrait gets a reference to the given EntityTraitDefinitionResponseData and assigns it to the Trait field.
func (o *RuleConditionDetailResponseData) SetTrait(v EntityTraitDefinitionResponseData) {
	o.Trait = &v
}

// GetTraitEntityType returns the TraitEntityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetTraitEntityType() string {
	if o == nil || IsNil(o.TraitEntityType.Get()) {
		var ret string
		return ret
	}
	return *o.TraitEntityType.Get()
}

// GetTraitEntityTypeOk returns a tuple with the TraitEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetTraitEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitEntityType.Get(), o.TraitEntityType.IsSet()
}

// HasTraitEntityType returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasTraitEntityType() bool {
	if o != nil && o.TraitEntityType.IsSet() {
		return true
	}

	return false
}

// SetTraitEntityType gets a reference to the given NullableString and assigns it to the TraitEntityType field.
func (o *RuleConditionDetailResponseData) SetTraitEntityType(v string) {
	o.TraitEntityType.Set(&v)
}

// SetTraitEntityTypeNil sets the value for TraitEntityType to be an explicit nil
func (o *RuleConditionDetailResponseData) SetTraitEntityTypeNil() {
	o.TraitEntityType.Set(nil)
}

// UnsetTraitEntityType ensures that no value is present for TraitEntityType, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetTraitEntityType() {
	o.TraitEntityType.Unset()
}

// GetTraitId returns the TraitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleConditionDetailResponseData) GetTraitId() string {
	if o == nil || IsNil(o.TraitId.Get()) {
		var ret string
		return ret
	}
	return *o.TraitId.Get()
}

// GetTraitIdOk returns a tuple with the TraitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleConditionDetailResponseData) GetTraitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitId.Get(), o.TraitId.IsSet()
}

// HasTraitId returns a boolean if a field has been set.
func (o *RuleConditionDetailResponseData) HasTraitId() bool {
	if o != nil && o.TraitId.IsSet() {
		return true
	}

	return false
}

// SetTraitId gets a reference to the given NullableString and assigns it to the TraitId field.
func (o *RuleConditionDetailResponseData) SetTraitId(v string) {
	o.TraitId.Set(&v)
}

// SetTraitIdNil sets the value for TraitId to be an explicit nil
func (o *RuleConditionDetailResponseData) SetTraitIdNil() {
	o.TraitId.Set(nil)
}

// UnsetTraitId ensures that no value is present for TraitId, not even an explicit nil
func (o *RuleConditionDetailResponseData) UnsetTraitId() {
	o.TraitId.Unset()
}

// GetTraitValue returns the TraitValue field value
func (o *RuleConditionDetailResponseData) GetTraitValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraitValue
}

// GetTraitValueOk returns a tuple with the TraitValue field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetTraitValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraitValue, true
}

// SetTraitValue sets field value
func (o *RuleConditionDetailResponseData) SetTraitValue(v string) {
	o.TraitValue = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RuleConditionDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RuleConditionDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RuleConditionDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RuleConditionDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleConditionDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComparisonTrait) {
		toSerialize["comparison_trait"] = o.ComparisonTrait
	}
	if o.ComparisonTraitId.IsSet() {
		toSerialize["comparison_trait_id"] = o.ComparisonTraitId.Get()
	}
	if o.ConditionGroupId.IsSet() {
		toSerialize["condition_group_id"] = o.ConditionGroupId.Get()
	}
	toSerialize["condition_type"] = o.ConditionType
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment_id"] = o.EnvironmentId
	if o.EventSubtype.IsSet() {
		toSerialize["event_subtype"] = o.EventSubtype.Get()
	}
	if o.FlagId.IsSet() {
		toSerialize["flag_id"] = o.FlagId.Get()
	}
	toSerialize["id"] = o.Id
	if o.MetricPeriod.IsSet() {
		toSerialize["metric_period"] = o.MetricPeriod.Get()
	}
	if o.MetricValue.IsSet() {
		toSerialize["metric_value"] = o.MetricValue.Get()
	}
	toSerialize["operator"] = o.Operator
	if o.PlanId.IsSet() {
		toSerialize["plan_id"] = o.PlanId.Get()
	}
	toSerialize["resource_ids"] = o.ResourceIds
	toSerialize["resources"] = o.Resources
	toSerialize["rule_id"] = o.RuleId
	if !IsNil(o.Trait) {
		toSerialize["trait"] = o.Trait
	}
	if o.TraitEntityType.IsSet() {
		toSerialize["trait_entity_type"] = o.TraitEntityType.Get()
	}
	if o.TraitId.IsSet() {
		toSerialize["trait_id"] = o.TraitId.Get()
	}
	toSerialize["trait_value"] = o.TraitValue
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *RuleConditionDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"condition_type",
		"created_at",
		"environment_id",
		"id",
		"operator",
		"resource_ids",
		"resources",
		"rule_id",
		"trait_value",
		"updated_at",
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

	varRuleConditionDetailResponseData := _RuleConditionDetailResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRuleConditionDetailResponseData)

	if err != nil {
		return err
	}

	*o = RuleConditionDetailResponseData(varRuleConditionDetailResponseData)

	return err
}

type NullableRuleConditionDetailResponseData struct {
	value *RuleConditionDetailResponseData
	isSet bool
}

func (v NullableRuleConditionDetailResponseData) Get() *RuleConditionDetailResponseData {
	return v.value
}

func (v *NullableRuleConditionDetailResponseData) Set(val *RuleConditionDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleConditionDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleConditionDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleConditionDetailResponseData(val *RuleConditionDetailResponseData) *NullableRuleConditionDetailResponseData {
	return &NullableRuleConditionDetailResponseData{value: val, isSet: true}
}

func (v NullableRuleConditionDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleConditionDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
