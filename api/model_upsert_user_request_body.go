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

// checks if the UpsertUserRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertUserRequestBody{}

// UpsertUserRequestBody struct for UpsertUserRequestBody
type UpsertUserRequestBody struct {
	// Optionally specify company using key/value pairs
	Company map[string]string `json:"company"`
	// Optionally specify company using Schematic company ID
	CompanyId NullableString `json:"company_id,omitempty"`
	// If you know the Schematic ID, you can use that here instead of keys
	Id         NullableString    `json:"id,omitempty"`
	Keys       map[string]string `json:"keys"`
	LastSeenAt NullableTime      `json:"last_seen_at,omitempty"`
	Name       NullableString    `json:"name,omitempty"`
	// A map of trait names to trait values
	Traits               map[string]interface{} `json:"traits,omitempty"`
	UpdateOnly           NullableBool           `json:"update_only,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpsertUserRequestBody UpsertUserRequestBody

// NewUpsertUserRequestBody instantiates a new UpsertUserRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertUserRequestBody(company map[string]string, keys map[string]string) *UpsertUserRequestBody {
	this := UpsertUserRequestBody{}
	this.Company = company
	this.Keys = keys
	return &this
}

// NewUpsertUserRequestBodyWithDefaults instantiates a new UpsertUserRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertUserRequestBodyWithDefaults() *UpsertUserRequestBody {
	this := UpsertUserRequestBody{}
	return &this
}

// GetCompany returns the Company field value
func (o *UpsertUserRequestBody) GetCompany() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *UpsertUserRequestBody) GetCompanyOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *UpsertUserRequestBody) SetCompany(v map[string]string) {
	o.Company = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertUserRequestBody) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertUserRequestBody) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *UpsertUserRequestBody) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *UpsertUserRequestBody) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *UpsertUserRequestBody) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertUserRequestBody) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertUserRequestBody) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *UpsertUserRequestBody) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *UpsertUserRequestBody) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpsertUserRequestBody) UnsetId() {
	o.Id.Unset()
}

// GetKeys returns the Keys field value
func (o *UpsertUserRequestBody) GetKeys() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value
// and a boolean to check if the value has been set.
func (o *UpsertUserRequestBody) GetKeysOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keys, true
}

// SetKeys sets field value
func (o *UpsertUserRequestBody) SetKeys(v map[string]string) {
	o.Keys = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertUserRequestBody) GetLastSeenAt() time.Time {
	if o == nil || IsNil(o.LastSeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt.Get()
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertUserRequestBody) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenAt.Get(), o.LastSeenAt.IsSet()
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt.IsSet() {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given NullableTime and assigns it to the LastSeenAt field.
func (o *UpsertUserRequestBody) SetLastSeenAt(v time.Time) {
	o.LastSeenAt.Set(&v)
}

// SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil
func (o *UpsertUserRequestBody) SetLastSeenAtNil() {
	o.LastSeenAt.Set(nil)
}

// UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
func (o *UpsertUserRequestBody) UnsetLastSeenAt() {
	o.LastSeenAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertUserRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertUserRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpsertUserRequestBody) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpsertUserRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpsertUserRequestBody) UnsetName() {
	o.Name.Unset()
}

// GetTraits returns the Traits field value if set, zero value otherwise.
func (o *UpsertUserRequestBody) GetTraits() map[string]interface{} {
	if o == nil || IsNil(o.Traits) {
		var ret map[string]interface{}
		return ret
	}
	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertUserRequestBody) GetTraitsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Traits) {
		return map[string]interface{}{}, false
	}
	return o.Traits, true
}

// HasTraits returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasTraits() bool {
	if o != nil && !IsNil(o.Traits) {
		return true
	}

	return false
}

// SetTraits gets a reference to the given map[string]interface{} and assigns it to the Traits field.
func (o *UpsertUserRequestBody) SetTraits(v map[string]interface{}) {
	o.Traits = v
}

// GetUpdateOnly returns the UpdateOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertUserRequestBody) GetUpdateOnly() bool {
	if o == nil || IsNil(o.UpdateOnly.Get()) {
		var ret bool
		return ret
	}
	return *o.UpdateOnly.Get()
}

// GetUpdateOnlyOk returns a tuple with the UpdateOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertUserRequestBody) GetUpdateOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateOnly.Get(), o.UpdateOnly.IsSet()
}

// HasUpdateOnly returns a boolean if a field has been set.
func (o *UpsertUserRequestBody) HasUpdateOnly() bool {
	if o != nil && o.UpdateOnly.IsSet() {
		return true
	}

	return false
}

// SetUpdateOnly gets a reference to the given NullableBool and assigns it to the UpdateOnly field.
func (o *UpsertUserRequestBody) SetUpdateOnly(v bool) {
	o.UpdateOnly.Set(&v)
}

// SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil
func (o *UpsertUserRequestBody) SetUpdateOnlyNil() {
	o.UpdateOnly.Set(nil)
}

// UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil
func (o *UpsertUserRequestBody) UnsetUpdateOnly() {
	o.UpdateOnly.Unset()
}

func (o UpsertUserRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertUserRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpsertUserRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
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

	varUpsertUserRequestBody := _UpsertUserRequestBody{}

	err = json.Unmarshal(data, &varUpsertUserRequestBody)

	if err != nil {
		return err
	}

	*o = UpsertUserRequestBody(varUpsertUserRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company")
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "keys")
		delete(additionalProperties, "last_seen_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "update_only")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpsertUserRequestBody struct {
	value *UpsertUserRequestBody
	isSet bool
}

func (v NullableUpsertUserRequestBody) Get() *UpsertUserRequestBody {
	return v.value
}

func (v *NullableUpsertUserRequestBody) Set(val *UpsertUserRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertUserRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertUserRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertUserRequestBody(val *UpsertUserRequestBody) *NullableUpsertUserRequestBody {
	return &NullableUpsertUserRequestBody{value: val, isSet: true}
}

func (v NullableUpsertUserRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertUserRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
