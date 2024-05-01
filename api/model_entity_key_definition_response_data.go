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

// checks if the EntityKeyDefinitionResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityKeyDefinitionResponseData{}

// EntityKeyDefinitionResponseData struct for EntityKeyDefinitionResponseData
type EntityKeyDefinitionResponseData struct {
	CreatedAt  time.Time `json:"created_at"`
	EntityType string    `json:"entity_type"`
	Id         string    `json:"id"`
	Key        string    `json:"key"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type _EntityKeyDefinitionResponseData EntityKeyDefinitionResponseData

// NewEntityKeyDefinitionResponseData instantiates a new EntityKeyDefinitionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityKeyDefinitionResponseData(createdAt time.Time, entityType string, id string, key string, updatedAt time.Time) *EntityKeyDefinitionResponseData {
	this := EntityKeyDefinitionResponseData{}
	this.CreatedAt = createdAt
	this.EntityType = entityType
	this.Id = id
	this.Key = key
	this.UpdatedAt = updatedAt
	return &this
}

// NewEntityKeyDefinitionResponseDataWithDefaults instantiates a new EntityKeyDefinitionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityKeyDefinitionResponseDataWithDefaults() *EntityKeyDefinitionResponseData {
	this := EntityKeyDefinitionResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *EntityKeyDefinitionResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDefinitionResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EntityKeyDefinitionResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEntityType returns the EntityType field value
func (o *EntityKeyDefinitionResponseData) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDefinitionResponseData) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *EntityKeyDefinitionResponseData) SetEntityType(v string) {
	o.EntityType = v
}

// GetId returns the Id field value
func (o *EntityKeyDefinitionResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDefinitionResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityKeyDefinitionResponseData) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *EntityKeyDefinitionResponseData) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDefinitionResponseData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EntityKeyDefinitionResponseData) SetKey(v string) {
	o.Key = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EntityKeyDefinitionResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDefinitionResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EntityKeyDefinitionResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o EntityKeyDefinitionResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityKeyDefinitionResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["entity_type"] = o.EntityType
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *EntityKeyDefinitionResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"entity_type",
		"id",
		"key",
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

	varEntityKeyDefinitionResponseData := _EntityKeyDefinitionResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityKeyDefinitionResponseData)

	if err != nil {
		return err
	}

	*o = EntityKeyDefinitionResponseData(varEntityKeyDefinitionResponseData)

	return err
}

type NullableEntityKeyDefinitionResponseData struct {
	value *EntityKeyDefinitionResponseData
	isSet bool
}

func (v NullableEntityKeyDefinitionResponseData) Get() *EntityKeyDefinitionResponseData {
	return v.value
}

func (v *NullableEntityKeyDefinitionResponseData) Set(val *EntityKeyDefinitionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityKeyDefinitionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityKeyDefinitionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityKeyDefinitionResponseData(val *EntityKeyDefinitionResponseData) *NullableEntityKeyDefinitionResponseData {
	return &NullableEntityKeyDefinitionResponseData{value: val, isSet: true}
}

func (v NullableEntityKeyDefinitionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityKeyDefinitionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}