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

// checks if the FlagResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagResponseData{}

// FlagResponseData struct for FlagResponseData
type FlagResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	DefaultValue bool `json:"default_value"`
	Description string `json:"description"`
	FeatureId NullableString `json:"feature_id,omitempty"`
	FlagType string `json:"flag_type"`
	Id string `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewFlagResponseData instantiates a new FlagResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagResponseData(createdAt time.Time, defaultValue bool, description string, flagType string, id string, key string, name string, updatedAt time.Time) *FlagResponseData {
	this := FlagResponseData{}
	this.CreatedAt = createdAt
	this.DefaultValue = defaultValue
	this.Description = description
	this.FlagType = flagType
	this.Id = id
	this.Key = key
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewFlagResponseDataWithDefaults instantiates a new FlagResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagResponseDataWithDefaults() *FlagResponseData {
	this := FlagResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *FlagResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FlagResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefaultValue returns the DefaultValue field value
func (o *FlagResponseData) GetDefaultValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetDefaultValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *FlagResponseData) SetDefaultValue(v bool) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value
func (o *FlagResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FlagResponseData) SetDescription(v string) {
	o.Description = v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagResponseData) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureId.Get()
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagResponseData) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureId.Get(), o.FeatureId.IsSet()
}

// HasFeatureId returns a boolean if a field has been set.
func (o *FlagResponseData) HasFeatureId() bool {
	if o != nil && o.FeatureId.IsSet() {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given NullableString and assigns it to the FeatureId field.
func (o *FlagResponseData) SetFeatureId(v string) {
	o.FeatureId.Set(&v)
}
// SetFeatureIdNil sets the value for FeatureId to be an explicit nil
func (o *FlagResponseData) SetFeatureIdNil() {
	o.FeatureId.Set(nil)
}

// UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
func (o *FlagResponseData) UnsetFeatureId() {
	o.FeatureId.Unset()
}

// GetFlagType returns the FlagType field value
func (o *FlagResponseData) GetFlagType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlagType
}

// GetFlagTypeOk returns a tuple with the FlagType field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetFlagTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagType, true
}

// SetFlagType sets field value
func (o *FlagResponseData) SetFlagType(v string) {
	o.FlagType = v
}

// GetId returns the Id field value
func (o *FlagResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlagResponseData) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *FlagResponseData) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FlagResponseData) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *FlagResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlagResponseData) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FlagResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FlagResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FlagResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o FlagResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["description"] = o.Description
	if o.FeatureId.IsSet() {
		toSerialize["feature_id"] = o.FeatureId.Get()
	}
	toSerialize["flag_type"] = o.FlagType
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableFlagResponseData struct {
	value *FlagResponseData
	isSet bool
}

func (v NullableFlagResponseData) Get() *FlagResponseData {
	return v.value
}

func (v *NullableFlagResponseData) Set(val *FlagResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagResponseData(val *FlagResponseData) *NullableFlagResponseData {
	return &NullableFlagResponseData{value: val, isSet: true}
}

func (v NullableFlagResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


