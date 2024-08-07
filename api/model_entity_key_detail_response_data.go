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

// checks if the EntityKeyDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityKeyDetailResponseData{}

// EntityKeyDetailResponseData struct for EntityKeyDetailResponseData
type EntityKeyDetailResponseData struct {
	CreatedAt            time.Time                        `json:"created_at"`
	Definition           *EntityKeyDefinitionResponseData `json:"definition,omitempty"`
	DefinitionId         string                           `json:"definition_id"`
	EntityId             string                           `json:"entity_id"`
	EntityType           string                           `json:"entity_type"`
	EnvironmentId        string                           `json:"environment_id"`
	Id                   string                           `json:"id"`
	Key                  string                           `json:"key"`
	UpdatedAt            time.Time                        `json:"updated_at"`
	Value                string                           `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _EntityKeyDetailResponseData EntityKeyDetailResponseData

// NewEntityKeyDetailResponseData instantiates a new EntityKeyDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityKeyDetailResponseData(createdAt time.Time, definitionId string, entityId string, entityType string, environmentId string, id string, key string, updatedAt time.Time, value string) *EntityKeyDetailResponseData {
	this := EntityKeyDetailResponseData{}
	this.CreatedAt = createdAt
	this.DefinitionId = definitionId
	this.EntityId = entityId
	this.EntityType = entityType
	this.EnvironmentId = environmentId
	this.Id = id
	this.Key = key
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewEntityKeyDetailResponseDataWithDefaults instantiates a new EntityKeyDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityKeyDetailResponseDataWithDefaults() *EntityKeyDetailResponseData {
	this := EntityKeyDetailResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *EntityKeyDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EntityKeyDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *EntityKeyDetailResponseData) GetDefinition() EntityKeyDefinitionResponseData {
	if o == nil || IsNil(o.Definition) {
		var ret EntityKeyDefinitionResponseData
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetDefinitionOk() (*EntityKeyDefinitionResponseData, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *EntityKeyDetailResponseData) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given EntityKeyDefinitionResponseData and assigns it to the Definition field.
func (o *EntityKeyDetailResponseData) SetDefinition(v EntityKeyDefinitionResponseData) {
	o.Definition = &v
}

// GetDefinitionId returns the DefinitionId field value
func (o *EntityKeyDetailResponseData) GetDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefinitionId
}

// GetDefinitionIdOk returns a tuple with the DefinitionId field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionId, true
}

// SetDefinitionId sets field value
func (o *EntityKeyDetailResponseData) SetDefinitionId(v string) {
	o.DefinitionId = v
}

// GetEntityId returns the EntityId field value
func (o *EntityKeyDetailResponseData) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *EntityKeyDetailResponseData) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value
func (o *EntityKeyDetailResponseData) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *EntityKeyDetailResponseData) SetEntityType(v string) {
	o.EntityType = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *EntityKeyDetailResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *EntityKeyDetailResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *EntityKeyDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntityKeyDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *EntityKeyDetailResponseData) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *EntityKeyDetailResponseData) SetKey(v string) {
	o.Key = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EntityKeyDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EntityKeyDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValue returns the Value field value
func (o *EntityKeyDetailResponseData) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EntityKeyDetailResponseData) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EntityKeyDetailResponseData) SetValue(v string) {
	o.Value = v
}

func (o EntityKeyDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityKeyDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	toSerialize["definition_id"] = o.DefinitionId
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntityKeyDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"definition_id",
		"entity_id",
		"entity_type",
		"environment_id",
		"id",
		"key",
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

	varEntityKeyDetailResponseData := _EntityKeyDetailResponseData{}

	err = json.Unmarshal(data, &varEntityKeyDetailResponseData)

	if err != nil {
		return err
	}

	*o = EntityKeyDetailResponseData(varEntityKeyDetailResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "definition")
		delete(additionalProperties, "definition_id")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "entity_type")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityKeyDetailResponseData struct {
	value *EntityKeyDetailResponseData
	isSet bool
}

func (v NullableEntityKeyDetailResponseData) Get() *EntityKeyDetailResponseData {
	return v.value
}

func (v *NullableEntityKeyDetailResponseData) Set(val *EntityKeyDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityKeyDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityKeyDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityKeyDetailResponseData(val *EntityKeyDetailResponseData) *NullableEntityKeyDetailResponseData {
	return &NullableEntityKeyDetailResponseData{value: val, isSet: true}
}

func (v NullableEntityKeyDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityKeyDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
