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

// checks if the CompanyResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyResponseData{}

// CompanyResponseData struct for CompanyResponseData
type CompanyResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	EnvironmentId string `json:"environment_id"`
	Id string `json:"id"`
	LastSeenAt NullableTime `json:"last_seen_at,omitempty"`
	Name string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewCompanyResponseData instantiates a new CompanyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyResponseData(createdAt time.Time, environmentId string, id string, name string, updatedAt time.Time) *CompanyResponseData {
	this := CompanyResponseData{}
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewCompanyResponseDataWithDefaults instantiates a new CompanyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyResponseDataWithDefaults() *CompanyResponseData {
	this := CompanyResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *CompanyResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CompanyResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *CompanyResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *CompanyResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *CompanyResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CompanyResponseData) SetId(v string) {
	o.Id = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompanyResponseData) GetLastSeenAt() time.Time {
	if o == nil || IsNil(o.LastSeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt.Get()
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompanyResponseData) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenAt.Get(), o.LastSeenAt.IsSet()
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *CompanyResponseData) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt.IsSet() {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given NullableTime and assigns it to the LastSeenAt field.
func (o *CompanyResponseData) SetLastSeenAt(v time.Time) {
	o.LastSeenAt.Set(&v)
}
// SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil
func (o *CompanyResponseData) SetLastSeenAtNil() {
	o.LastSeenAt.Set(nil)
}

// UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
func (o *CompanyResponseData) UnsetLastSeenAt() {
	o.LastSeenAt.Unset()
}

// GetName returns the Name field value
func (o *CompanyResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompanyResponseData) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CompanyResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CompanyResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CompanyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["id"] = o.Id
	if o.LastSeenAt.IsSet() {
		toSerialize["last_seen_at"] = o.LastSeenAt.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableCompanyResponseData struct {
	value *CompanyResponseData
	isSet bool
}

func (v NullableCompanyResponseData) Get() *CompanyResponseData {
	return v.value
}

func (v *NullableCompanyResponseData) Set(val *CompanyResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyResponseData(val *CompanyResponseData) *NullableCompanyResponseData {
	return &NullableCompanyResponseData{value: val, isSet: true}
}

func (v NullableCompanyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


