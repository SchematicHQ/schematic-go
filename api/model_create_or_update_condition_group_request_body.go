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

// checks if the CreateOrUpdateConditionGroupRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateConditionGroupRequestBody{}

// CreateOrUpdateConditionGroupRequestBody struct for CreateOrUpdateConditionGroupRequestBody
type CreateOrUpdateConditionGroupRequestBody struct {
	Conditions []CreateOrUpdateConditionRequestBody `json:"conditions"`
	FlagId     NullableString                       `json:"flag_id,omitempty"`
	Id         NullableString                       `json:"id,omitempty"`
	PlanId     NullableString                       `json:"plan_id,omitempty"`
}

type _CreateOrUpdateConditionGroupRequestBody CreateOrUpdateConditionGroupRequestBody

// NewCreateOrUpdateConditionGroupRequestBody instantiates a new CreateOrUpdateConditionGroupRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateConditionGroupRequestBody(conditions []CreateOrUpdateConditionRequestBody) *CreateOrUpdateConditionGroupRequestBody {
	this := CreateOrUpdateConditionGroupRequestBody{}
	this.Conditions = conditions
	return &this
}

// NewCreateOrUpdateConditionGroupRequestBodyWithDefaults instantiates a new CreateOrUpdateConditionGroupRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateConditionGroupRequestBodyWithDefaults() *CreateOrUpdateConditionGroupRequestBody {
	this := CreateOrUpdateConditionGroupRequestBody{}
	return &this
}

// GetConditions returns the Conditions field value
func (o *CreateOrUpdateConditionGroupRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody {
	if o == nil {
		var ret []CreateOrUpdateConditionRequestBody
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConditionGroupRequestBody) GetConditionsOk() ([]CreateOrUpdateConditionRequestBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *CreateOrUpdateConditionGroupRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody) {
	o.Conditions = v
}

// GetFlagId returns the FlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateConditionGroupRequestBody) GetFlagId() string {
	if o == nil || IsNil(o.FlagId.Get()) {
		var ret string
		return ret
	}
	return *o.FlagId.Get()
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateConditionGroupRequestBody) GetFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagId.Get(), o.FlagId.IsSet()
}

// HasFlagId returns a boolean if a field has been set.
func (o *CreateOrUpdateConditionGroupRequestBody) HasFlagId() bool {
	if o != nil && o.FlagId.IsSet() {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given NullableString and assigns it to the FlagId field.
func (o *CreateOrUpdateConditionGroupRequestBody) SetFlagId(v string) {
	o.FlagId.Set(&v)
}

// SetFlagIdNil sets the value for FlagId to be an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) SetFlagIdNil() {
	o.FlagId.Set(nil)
}

// UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) UnsetFlagId() {
	o.FlagId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateConditionGroupRequestBody) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateConditionGroupRequestBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrUpdateConditionGroupRequestBody) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateOrUpdateConditionGroupRequestBody) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) UnsetId() {
	o.Id.Unset()
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateConditionGroupRequestBody) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateConditionGroupRequestBody) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *CreateOrUpdateConditionGroupRequestBody) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *CreateOrUpdateConditionGroupRequestBody) SetPlanId(v string) {
	o.PlanId.Set(&v)
}

// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *CreateOrUpdateConditionGroupRequestBody) UnsetPlanId() {
	o.PlanId.Unset()
}

func (o CreateOrUpdateConditionGroupRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateConditionGroupRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditions"] = o.Conditions
	if o.FlagId.IsSet() {
		toSerialize["flag_id"] = o.FlagId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.PlanId.IsSet() {
		toSerialize["plan_id"] = o.PlanId.Get()
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateConditionGroupRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conditions",
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

	varCreateOrUpdateConditionGroupRequestBody := _CreateOrUpdateConditionGroupRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateConditionGroupRequestBody)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateConditionGroupRequestBody(varCreateOrUpdateConditionGroupRequestBody)

	return err
}

type NullableCreateOrUpdateConditionGroupRequestBody struct {
	value *CreateOrUpdateConditionGroupRequestBody
	isSet bool
}

func (v NullableCreateOrUpdateConditionGroupRequestBody) Get() *CreateOrUpdateConditionGroupRequestBody {
	return v.value
}

func (v *NullableCreateOrUpdateConditionGroupRequestBody) Set(val *CreateOrUpdateConditionGroupRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateConditionGroupRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateConditionGroupRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateConditionGroupRequestBody(val *CreateOrUpdateConditionGroupRequestBody) *NullableCreateOrUpdateConditionGroupRequestBody {
	return &NullableCreateOrUpdateConditionGroupRequestBody{value: val, isSet: true}
}

func (v NullableCreateOrUpdateConditionGroupRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateConditionGroupRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
