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

// checks if the PlanGroupResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanGroupResponseData{}

// PlanGroupResponseData The updated resource
type PlanGroupResponseData struct {
	DefaultPlanId        NullableString `json:"default_plan_id,omitempty"`
	Id                   string         `json:"id"`
	PlanIds              []string       `json:"plan_ids"`
	AdditionalProperties map[string]interface{}
}

type _PlanGroupResponseData PlanGroupResponseData

// NewPlanGroupResponseData instantiates a new PlanGroupResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanGroupResponseData(id string, planIds []string) *PlanGroupResponseData {
	this := PlanGroupResponseData{}
	this.Id = id
	this.PlanIds = planIds
	return &this
}

// NewPlanGroupResponseDataWithDefaults instantiates a new PlanGroupResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanGroupResponseDataWithDefaults() *PlanGroupResponseData {
	this := PlanGroupResponseData{}
	return &this
}

// GetDefaultPlanId returns the DefaultPlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlanGroupResponseData) GetDefaultPlanId() string {
	if o == nil || IsNil(o.DefaultPlanId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultPlanId.Get()
}

// GetDefaultPlanIdOk returns a tuple with the DefaultPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlanGroupResponseData) GetDefaultPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPlanId.Get(), o.DefaultPlanId.IsSet()
}

// HasDefaultPlanId returns a boolean if a field has been set.
func (o *PlanGroupResponseData) HasDefaultPlanId() bool {
	if o != nil && o.DefaultPlanId.IsSet() {
		return true
	}

	return false
}

// SetDefaultPlanId gets a reference to the given NullableString and assigns it to the DefaultPlanId field.
func (o *PlanGroupResponseData) SetDefaultPlanId(v string) {
	o.DefaultPlanId.Set(&v)
}

// SetDefaultPlanIdNil sets the value for DefaultPlanId to be an explicit nil
func (o *PlanGroupResponseData) SetDefaultPlanIdNil() {
	o.DefaultPlanId.Set(nil)
}

// UnsetDefaultPlanId ensures that no value is present for DefaultPlanId, not even an explicit nil
func (o *PlanGroupResponseData) UnsetDefaultPlanId() {
	o.DefaultPlanId.Unset()
}

// GetId returns the Id field value
func (o *PlanGroupResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanGroupResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlanGroupResponseData) SetId(v string) {
	o.Id = v
}

// GetPlanIds returns the PlanIds field value
func (o *PlanGroupResponseData) GetPlanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value
// and a boolean to check if the value has been set.
func (o *PlanGroupResponseData) GetPlanIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanIds, true
}

// SetPlanIds sets field value
func (o *PlanGroupResponseData) SetPlanIds(v []string) {
	o.PlanIds = v
}

func (o PlanGroupResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanGroupResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultPlanId.IsSet() {
		toSerialize["default_plan_id"] = o.DefaultPlanId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["plan_ids"] = o.PlanIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanGroupResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"plan_ids",
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

	varPlanGroupResponseData := _PlanGroupResponseData{}

	err = json.Unmarshal(data, &varPlanGroupResponseData)

	if err != nil {
		return err
	}

	*o = PlanGroupResponseData(varPlanGroupResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default_plan_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "plan_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanGroupResponseData struct {
	value *PlanGroupResponseData
	isSet bool
}

func (v NullablePlanGroupResponseData) Get() *PlanGroupResponseData {
	return v.value
}

func (v *NullablePlanGroupResponseData) Set(val *PlanGroupResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanGroupResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanGroupResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanGroupResponseData(val *PlanGroupResponseData) *NullablePlanGroupResponseData {
	return &NullablePlanGroupResponseData{value: val, isSet: true}
}

func (v NullablePlanGroupResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanGroupResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
