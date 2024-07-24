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

// checks if the EventBodyIdentify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventBodyIdentify{}

// EventBodyIdentify struct for EventBodyIdentify
type EventBodyIdentify struct {
	Company *EventBodyIdentifyCompany `json:"company,omitempty"`
	// Key-value pairs to identify the user
	Keys map[string]string `json:"keys"`
	// The display name of the user being identified; required only if it is a new user
	Name *string `json:"name,omitempty"`
	// A map of trait names to trait values
	Traits               map[string]interface{} `json:"traits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventBodyIdentify EventBodyIdentify

// NewEventBodyIdentify instantiates a new EventBodyIdentify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBodyIdentify(keys map[string]string) *EventBodyIdentify {
	this := EventBodyIdentify{}
	this.Keys = keys
	return &this
}

// NewEventBodyIdentifyWithDefaults instantiates a new EventBodyIdentify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBodyIdentifyWithDefaults() *EventBodyIdentify {
	this := EventBodyIdentify{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *EventBodyIdentify) GetCompany() EventBodyIdentifyCompany {
	if o == nil || IsNil(o.Company) {
		var ret EventBodyIdentifyCompany
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBodyIdentify) GetCompanyOk() (*EventBodyIdentifyCompany, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *EventBodyIdentify) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given EventBodyIdentifyCompany and assigns it to the Company field.
func (o *EventBodyIdentify) SetCompany(v EventBodyIdentifyCompany) {
	o.Company = &v
}

// GetKeys returns the Keys field value
func (o *EventBodyIdentify) GetKeys() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *EventBodyIdentify) GetKeysOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value
func (o *EventBodyIdentify) SetKeys(v map[string]string) {
	o.Keys = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventBodyIdentify) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBodyIdentify) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventBodyIdentify) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventBodyIdentify) SetName(v string) {
	o.Name = &v
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *EventBodyIdentify) GetTraits() map[string]interface{} {
	if o == nil || IsNil(o.Traits) {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventBodyIdentify) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Traits) {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *EventBodyIdentify) HasTraits() bool {
	if o != nil && !IsNil(o.Traits) {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *EventBodyIdentify) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

func (o EventBodyIdentify) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventBodyIdentify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["keys"] = o.Keys
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Traits) {
		toSerialize["traits"] = o.Traits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventBodyIdentify) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"keys",
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

	varEventBodyIdentify := _EventBodyIdentify{}

	err = json.Unmarshal(data, &varEventBodyIdentify)

	if err != nil {
		return err
	}

	*o = EventBodyIdentify(varEventBodyIdentify)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company")
		delete(additionalProperties, "keys")
		delete(additionalProperties, "name")
		delete(additionalProperties, "traits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventBodyIdentify struct {
	value *EventBodyIdentify
	isSet bool
}

func (v NullableEventBodyIdentify) Get() *EventBodyIdentify {
	return v.value
}

func (v *NullableEventBodyIdentify) Set(val *EventBodyIdentify) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBodyIdentify) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBodyIdentify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBodyIdentify(val *EventBodyIdentify) *NullableEventBodyIdentify {
	return &NullableEventBodyIdentify{value: val, isSet: true}
}

func (v NullableEventBodyIdentify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventBodyIdentify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
