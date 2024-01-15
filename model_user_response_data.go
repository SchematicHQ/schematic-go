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

// checks if the UserResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponseData{}

// UserResponseData struct for UserResponseData
type UserResponseData struct {
	CreatedAt time.Time `json:"created_at"`
	EnvironmentId string `json:"environment_id"`
	Id string `json:"id"`
	LastSeenAt NullableTime `json:"last_seen_at,omitempty"`
	Name string `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserResponseData instantiates a new UserResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseData(createdAt time.Time, environmentId string, id string, name string, updatedAt time.Time) *UserResponseData {
	this := UserResponseData{}
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewUserResponseDataWithDefaults instantiates a new UserResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseDataWithDefaults() *UserResponseData {
	this := UserResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *UserResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *UserResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetId returns the Id field value
func (o *UserResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserResponseData) SetId(v string) {
	o.Id = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponseData) GetLastSeenAt() time.Time {
	if o == nil || IsNil(o.LastSeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt.Get()
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponseData) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenAt.Get(), o.LastSeenAt.IsSet()
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *UserResponseData) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt.IsSet() {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given NullableTime and assigns it to the LastSeenAt field.
func (o *UserResponseData) SetLastSeenAt(v time.Time) {
	o.LastSeenAt.Set(&v)
}
// SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil
func (o *UserResponseData) SetLastSeenAtNil() {
	o.LastSeenAt.Set(nil)
}

// UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
func (o *UserResponseData) UnsetLastSeenAt() {
	o.LastSeenAt.Unset()
}

// GetName returns the Name field value
func (o *UserResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserResponseData) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o UserResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableUserResponseData struct {
	value *UserResponseData
	isSet bool
}

func (v NullableUserResponseData) Get() *UserResponseData {
	return v.value
}

func (v *NullableUserResponseData) Set(val *UserResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseData(val *UserResponseData) *NullableUserResponseData {
	return &NullableUserResponseData{value: val, isSet: true}
}

func (v NullableUserResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


