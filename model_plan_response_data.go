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

// checks if the PlanResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanResponseData{}

// PlanResponseData The updated resource
type PlanResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	Id string `json:"id"`
	Name string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPlanResponseData instantiates a new PlanResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanResponseData(createdAt time.Time, id string, name string, updatedAt time.Time) *PlanResponseData {
	this := PlanResponseData{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewPlanResponseDataWithDefaults instantiates a new PlanResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanResponseDataWithDefaults() *PlanResponseData {
	this := PlanResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *PlanResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PlanResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *PlanResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlanResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PlanResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanResponseData) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PlanResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PlanResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PlanResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullablePlanResponseData struct {
	value *PlanResponseData
	isSet bool
}

func (v NullablePlanResponseData) Get() *PlanResponseData {
	return v.value
}

func (v *NullablePlanResponseData) Set(val *PlanResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanResponseData(val *PlanResponseData) *NullablePlanResponseData {
	return &NullablePlanResponseData{value: val, isSet: true}
}

func (v NullablePlanResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


