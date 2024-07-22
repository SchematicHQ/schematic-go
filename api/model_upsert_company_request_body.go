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

// checks if the UpsertCompanyRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertCompanyRequestBody{}

// UpsertCompanyRequestBody struct for UpsertCompanyRequestBody
type UpsertCompanyRequestBody struct {
	// If you know the Schematic ID, you can use that here instead of keys
	Id         NullableString    `json:"id,omitempty"`
	Keys       map[string]string `json:"keys"`
	LastSeenAt NullableTime      `json:"last_seen_at,omitempty"`
	Name       NullableString    `json:"name,omitempty"`
	// A map of trait names to trait values
	Traits     map[string]interface{} `json:"traits,omitempty"`
	UpdateOnly NullableBool           `json:"update_only,omitempty"`
}

type _UpsertCompanyRequestBody UpsertCompanyRequestBody

// NewUpsertCompanyRequestBody instantiates a new UpsertCompanyRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertCompanyRequestBody(keys map[string]string) *UpsertCompanyRequestBody {
	this := UpsertCompanyRequestBody{}
	this.Keys = keys
	return &this
}

// NewUpsertCompanyRequestBodyWithDefaults instantiates a new UpsertCompanyRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertCompanyRequestBodyWithDefaults() *UpsertCompanyRequestBody {
	this := UpsertCompanyRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertCompanyRequestBody) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertCompanyRequestBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpsertCompanyRequestBody) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *UpsertCompanyRequestBody) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *UpsertCompanyRequestBody) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpsertCompanyRequestBody) UnsetId() {
	o.Id.Unset()
}

// GetKeys returns the Keys field value
func (o *UpsertCompanyRequestBody) GetKeys() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *UpsertCompanyRequestBody) GetKeysOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value
func (o *UpsertCompanyRequestBody) SetKeys(v map[string]string) {
	o.Keys = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertCompanyRequestBody) GetLastSeenAt() time.Time {
	if o == nil || IsNil(o.LastSeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt.Get()
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertCompanyRequestBody) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenAt.Get(), o.LastSeenAt.IsSet()
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *UpsertCompanyRequestBody) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt.IsSet() {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given NullableTime and assigns it to the LastSeenAt field.
func (o *UpsertCompanyRequestBody) SetLastSeenAt(v time.Time) {
	o.LastSeenAt.Set(&v)
}

// SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil
func (o *UpsertCompanyRequestBody) SetLastSeenAtNil() {
	o.LastSeenAt.Set(nil)
}

// UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
func (o *UpsertCompanyRequestBody) UnsetLastSeenAt() {
	o.LastSeenAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertCompanyRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertCompanyRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpsertCompanyRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpsertCompanyRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpsertCompanyRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpsertCompanyRequestBody) UnsetName() {
	o.Name.Unset()
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *UpsertCompanyRequestBody) GetTraits() map[string]interface{} {
	if o == nil || IsNil(o.Traits) {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertCompanyRequestBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Traits) {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *UpsertCompanyRequestBody) HasTraits() bool {
	if o != nil && !IsNil(o.Traits) {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *UpsertCompanyRequestBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetUpdateOnly returns the UpdateOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertCompanyRequestBody) GetUpdateOnly() bool {
	if o == nil || IsNil(o.UpdateOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.UpdateOnly.Get()
}

// GetUpdateOnlyOk returns a tuple with the UpdateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertCompanyRequestBody) GetUpdateOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateOnly.Get(), o.UpdateOnly.IsSet()
}

// HasUpdateOnly returns a boolean if a field has been set.
func (o *UpsertCompanyRequestBody) HasUpdateOnly() bool {
	if o != nil && o.UpdateOnly.IsSet() {
		return true
	}

	return false
}

// SetUpdateOnly gets a reference to the given NullableBool and assigns it to the UpdateOnly field.
func (o *UpsertCompanyRequestBody) SetUpdateOnly(v bool) {
	o.UpdateOnly.Set(&v)
}

// SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil
func (o *UpsertCompanyRequestBody) SetUpdateOnlyNil() {
	o.UpdateOnly.Set(nil)
}

// UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil
func (o *UpsertCompanyRequestBody) UnsetUpdateOnly() {
	o.UpdateOnly.Unset()
}

func (o UpsertCompanyRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertCompanyRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["keys"] = o.Keys
	if o.LastSeenAt.IsSet() {
		toSerialize["last_seen_at"] = o.LastSeenAt.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Traits) {
		toSerialize["traits"] = o.Traits
	}
	if o.UpdateOnly.IsSet() {
		toSerialize["update_only"] = o.UpdateOnly.Get()
	}
	return toSerialize, nil
}

func (o *UpsertCompanyRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varUpsertCompanyRequestBody := _UpsertCompanyRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpsertCompanyRequestBody)

	if err != nil {
		return err
	}

	*o = UpsertCompanyRequestBody(varUpsertCompanyRequestBody)

	return err
}

type NullableUpsertCompanyRequestBody struct {
	value *UpsertCompanyRequestBody
	isSet bool
}

func (v NullableUpsertCompanyRequestBody) Get() *UpsertCompanyRequestBody {
	return v.value
}

func (v *NullableUpsertCompanyRequestBody) Set(val *UpsertCompanyRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertCompanyRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertCompanyRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertCompanyRequestBody(val *UpsertCompanyRequestBody) *NullableUpsertCompanyRequestBody {
	return &NullableUpsertCompanyRequestBody{value: val, isSet: true}
}

func (v NullableUpsertCompanyRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertCompanyRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
