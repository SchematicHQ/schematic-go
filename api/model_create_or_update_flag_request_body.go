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

// checks if the CreateOrUpdateFlagRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateFlagRequestBody{}

// CreateOrUpdateFlagRequestBody struct for CreateOrUpdateFlagRequestBody
type CreateOrUpdateFlagRequestBody struct {
	DefaultValue         bool           `json:"default_value"`
	Description          string         `json:"description"`
	FeatureId            NullableString `json:"feature_id,omitempty"`
	FlagType             string         `json:"flag_type"`
	Id                   NullableString `json:"id,omitempty"`
	Key                  string         `json:"key"`
	MaintainerId         NullableString `json:"maintainer_id,omitempty"`
	Name                 string         `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _CreateOrUpdateFlagRequestBody CreateOrUpdateFlagRequestBody

// NewCreateOrUpdateFlagRequestBody instantiates a new CreateOrUpdateFlagRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateFlagRequestBody(defaultValue bool, description string, flagType string, key string, name string) *CreateOrUpdateFlagRequestBody {
	this := CreateOrUpdateFlagRequestBody{}
	this.DefaultValue = defaultValue
	this.Description = description
	this.FlagType = flagType
	this.Key = key
	this.Name = name
	return &this
}

// NewCreateOrUpdateFlagRequestBodyWithDefaults instantiates a new CreateOrUpdateFlagRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateFlagRequestBodyWithDefaults() *CreateOrUpdateFlagRequestBody {
	this := CreateOrUpdateFlagRequestBody{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value
func (o *CreateOrUpdateFlagRequestBody) GetDefaultValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFlagRequestBody) GetDefaultValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *CreateOrUpdateFlagRequestBody) SetDefaultValue(v bool) {
	o.DefaultValue = v
}

// GetDescription returns the Description field value
func (o *CreateOrUpdateFlagRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFlagRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateOrUpdateFlagRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateFlagRequestBody) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureId.Get()
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateFlagRequestBody) GetFeatureIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureId.Get(), o.FeatureId.IsSet()
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CreateOrUpdateFlagRequestBody) HasFeatureId() bool {
	if o != nil && o.FeatureId.IsSet() {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given NullableString and assigns it to the FeatureId field.
func (o *CreateOrUpdateFlagRequestBody) SetFeatureId(v string) {
	o.FeatureId.Set(&v)
}

// SetFeatureIdNil sets the value for FeatureId to be an explicit nil
func (o *CreateOrUpdateFlagRequestBody) SetFeatureIdNil() {
	o.FeatureId.Set(nil)
}

// UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
func (o *CreateOrUpdateFlagRequestBody) UnsetFeatureId() {
	o.FeatureId.Unset()
}

// GetFlagType returns the FlagType field value
func (o *CreateOrUpdateFlagRequestBody) GetFlagType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlagType
}

// GetFlagTypeOk returns a tuple with the FlagType field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFlagRequestBody) GetFlagTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagType, true
}

// SetFlagType sets field value
func (o *CreateOrUpdateFlagRequestBody) SetFlagType(v string) {
	o.FlagType = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateFlagRequestBody) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateFlagRequestBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrUpdateFlagRequestBody) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateOrUpdateFlagRequestBody) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateOrUpdateFlagRequestBody) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateOrUpdateFlagRequestBody) UnsetId() {
	o.Id.Unset()
}

// GetKey returns the Key field value
func (o *CreateOrUpdateFlagRequestBody) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFlagRequestBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CreateOrUpdateFlagRequestBody) SetKey(v string) {
	o.Key = v
}

// GetMaintainerId returns the MaintainerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateFlagRequestBody) GetMaintainerId() string {
	if o == nil || IsNil(o.MaintainerId.Get()) {
		var ret string
		return ret
	}
	return *o.MaintainerId.Get()
}

// GetMaintainerIdOk returns a tuple with the MaintainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateFlagRequestBody) GetMaintainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintainerId.Get(), o.MaintainerId.IsSet()
}

// HasMaintainerId returns a boolean if a field has been set.
func (o *CreateOrUpdateFlagRequestBody) HasMaintainerId() bool {
	if o != nil && o.MaintainerId.IsSet() {
		return true
	}

	return false
}

// SetMaintainerId gets a reference to the given NullableString and assigns it to the MaintainerId field.
func (o *CreateOrUpdateFlagRequestBody) SetMaintainerId(v string) {
	o.MaintainerId.Set(&v)
}

// SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil
func (o *CreateOrUpdateFlagRequestBody) SetMaintainerIdNil() {
	o.MaintainerId.Set(nil)
}

// UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
func (o *CreateOrUpdateFlagRequestBody) UnsetMaintainerId() {
	o.MaintainerId.Unset()
}

// GetName returns the Name field value
func (o *CreateOrUpdateFlagRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFlagRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateFlagRequestBody) SetName(v string) {
	o.Name = v
}

func (o CreateOrUpdateFlagRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateFlagRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default_value"] = o.DefaultValue
	toSerialize["description"] = o.Description
	if o.FeatureId.IsSet() {
		toSerialize["feature_id"] = o.FeatureId.Get()
	}
	toSerialize["flag_type"] = o.FlagType
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["key"] = o.Key
	if o.MaintainerId.IsSet() {
		toSerialize["maintainer_id"] = o.MaintainerId.Get()
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOrUpdateFlagRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default_value",
		"description",
		"flag_type",
		"key",
		"name",
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

	varCreateOrUpdateFlagRequestBody := _CreateOrUpdateFlagRequestBody{}

	err = json.Unmarshal(data, &varCreateOrUpdateFlagRequestBody)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateFlagRequestBody(varCreateOrUpdateFlagRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default_value")
		delete(additionalProperties, "description")
		delete(additionalProperties, "feature_id")
		delete(additionalProperties, "flag_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "maintainer_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOrUpdateFlagRequestBody struct {
	value *CreateOrUpdateFlagRequestBody
	isSet bool
}

func (v NullableCreateOrUpdateFlagRequestBody) Get() *CreateOrUpdateFlagRequestBody {
	return v.value
}

func (v *NullableCreateOrUpdateFlagRequestBody) Set(val *CreateOrUpdateFlagRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateFlagRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateFlagRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateFlagRequestBody(val *CreateOrUpdateFlagRequestBody) *NullableCreateOrUpdateFlagRequestBody {
	return &NullableCreateOrUpdateFlagRequestBody{value: val, isSet: true}
}

func (v NullableCreateOrUpdateFlagRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateFlagRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}