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

// checks if the CompanyMembershipResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyMembershipResponseData{}

// CompanyMembershipResponseData struct for CompanyMembershipResponseData
type CompanyMembershipResponseData struct {
	CompanyId string `json:"company_id"`
	CreatedAt time.Time `json:"created_at"`
	Id string `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId string `json:"user_id"`
}

// NewCompanyMembershipResponseData instantiates a new CompanyMembershipResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyMembershipResponseData(companyId string, createdAt time.Time, id string, updatedAt time.Time, userId string) *CompanyMembershipResponseData {
	this := CompanyMembershipResponseData{}
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	this.UserId = userId
	return &this
}

// NewCompanyMembershipResponseDataWithDefaults instantiates a new CompanyMembershipResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyMembershipResponseDataWithDefaults() *CompanyMembershipResponseData {
	this := CompanyMembershipResponseData{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *CompanyMembershipResponseData) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *CompanyMembershipResponseData) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CompanyMembershipResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CompanyMembershipResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *CompanyMembershipResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CompanyMembershipResponseData) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CompanyMembershipResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CompanyMembershipResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value
func (o *CompanyMembershipResponseData) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipResponseData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CompanyMembershipResponseData) SetUserId(v string) {
	o.UserId = v
}

func (o CompanyMembershipResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyMembershipResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company_id"] = o.CompanyId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

type NullableCompanyMembershipResponseData struct {
	value *CompanyMembershipResponseData
	isSet bool
}

func (v NullableCompanyMembershipResponseData) Get() *CompanyMembershipResponseData {
	return v.value
}

func (v *NullableCompanyMembershipResponseData) Set(val *CompanyMembershipResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyMembershipResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyMembershipResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyMembershipResponseData(val *CompanyMembershipResponseData) *NullableCompanyMembershipResponseData {
	return &NullableCompanyMembershipResponseData{value: val, isSet: true}
}

func (v NullableCompanyMembershipResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyMembershipResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


