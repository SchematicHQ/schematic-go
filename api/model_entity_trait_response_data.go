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

// checks if the EntityTraitResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityTraitResponseData{}

// EntityTraitResponseData struct for EntityTraitResponseData
type EntityTraitResponseData struct {
	CreatedAt     time.Time `json:"created_at"`
	DefinitionId  string    `json:"definition_id"`
	EnvironmentId string    `json:"environment_id"`
	Id            string    `json:"id"`
	UpdatedAt     time.Time `json:"updated_at"`
	Value         string    `json:"value"`
}

type _EntityTraitResponseData EntityTraitResponseData

// NewEntityTraitResponseData instantiates a new EntityTraitResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityTraitResponseData(createdAt time.Time, definitionId string, environmentId string, id string, updatedAt time.Time, value string) *EntityTraitResponseData {
	this := EntityTraitResponseData{}
	this.CreatedAt = createdAt
	this.DefinitionId = definitionId
	this.EnvironmentId = environmentId
	this.Id = id
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewEntityTraitResponseDataWithDefaults instantiates a new EntityTraitResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityTraitResponseDataWithDefaults() *EntityTraitResponseData {
	this := EntityTraitResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *EntityTraitResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EntityTraitResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefinitionId returns the DefinitionId field value
func (o *EntityTraitResponseData) GetDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefinitionId
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionId, true
}

// SetDefinitionId sets field value
func (o *EntityTraitResponseData) SetDefinitionId(v string) {
	o.DefinitionId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *EntityTraitResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *EntityTraitResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *EntityTraitResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityTraitResponseData) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EntityTraitResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EntityTraitResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
func (o *EntityTraitResponseData) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EntityTraitResponseData) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EntityTraitResponseData) SetValue(v string) {
	o.Value = v
}

func (o EntityTraitResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityTraitResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["definition_id"] = o.DefinitionId
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["id"] = o.Id
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *EntityTraitResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"definition_id",
		"environment_id",
		"id",
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

	varEntityTraitResponseData := _EntityTraitResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityTraitResponseData)

	if err != nil {
		return err
	}

	*o = EntityTraitResponseData(varEntityTraitResponseData)

	return err
}

type NullableEntityTraitResponseData struct {
	value *EntityTraitResponseData
	isSet bool
}

func (v NullableEntityTraitResponseData) Get() *EntityTraitResponseData {
	return v.value
}

func (v *NullableEntityTraitResponseData) Set(val *EntityTraitResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityTraitResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityTraitResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityTraitResponseData(val *EntityTraitResponseData) *NullableEntityTraitResponseData {
	return &NullableEntityTraitResponseData{value: val, isSet: true}
}

func (v NullableEntityTraitResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityTraitResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
