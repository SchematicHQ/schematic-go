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

// checks if the CompanyMembershipDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyMembershipDetailResponseData{}

// CompanyMembershipDetailResponseData struct for CompanyMembershipDetailResponseData
type CompanyMembershipDetailResponseData struct {
	Company   *CompanyResponseData `json:"company,omitempty"`
	CompanyId string               `json:"company_id"`
	CreatedAt time.Time            `json:"created_at"`
	Id        string               `json:"id"`
	UpdatedAt time.Time            `json:"updated_at"`
	UserId    string               `json:"user_id"`
}

type _CompanyMembershipDetailResponseData CompanyMembershipDetailResponseData

// NewCompanyMembershipDetailResponseData instantiates a new CompanyMembershipDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyMembershipDetailResponseData(companyId string, createdAt time.Time, id string, updatedAt time.Time, userId string) *CompanyMembershipDetailResponseData {
	this := CompanyMembershipDetailResponseData{}
	this.CompanyId = companyId
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	this.UserId = userId
	return &this
}

// NewCompanyMembershipDetailResponseDataWithDefaults instantiates a new CompanyMembershipDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyMembershipDetailResponseDataWithDefaults() *CompanyMembershipDetailResponseData {
	this := CompanyMembershipDetailResponseData{}
	return &this
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CompanyMembershipDetailResponseData) GetCompany() CompanyResponseData {
	if o == nil || IsNil(o.Company) {
		var ret CompanyResponseData
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetCompanyOk() (*CompanyResponseData, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CompanyMembershipDetailResponseData) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyResponseData and assigns it to the Company field.
func (o *CompanyMembershipDetailResponseData) SetCompany(v CompanyResponseData) {
	o.Company = &v
}

// GetCompanyId returns the CompanyId field value
func (o *CompanyMembershipDetailResponseData) GetCompanyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *CompanyMembershipDetailResponseData) SetCompanyId(v string) {
	o.CompanyId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CompanyMembershipDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CompanyMembershipDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *CompanyMembershipDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CompanyMembershipDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CompanyMembershipDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CompanyMembershipDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value
func (o *CompanyMembershipDetailResponseData) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CompanyMembershipDetailResponseData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CompanyMembershipDetailResponseData) SetUserId(v string) {
	o.UserId = v
}

func (o CompanyMembershipDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyMembershipDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["company_id"] = o.CompanyId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *CompanyMembershipDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company_id",
		"created_at",
		"id",
		"updated_at",
		"user_id",
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

	varCompanyMembershipDetailResponseData := _CompanyMembershipDetailResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyMembershipDetailResponseData)

	if err != nil {
		return err
	}

	*o = CompanyMembershipDetailResponseData(varCompanyMembershipDetailResponseData)

	return err
}

type NullableCompanyMembershipDetailResponseData struct {
	value *CompanyMembershipDetailResponseData
	isSet bool
}

func (v NullableCompanyMembershipDetailResponseData) Get() *CompanyMembershipDetailResponseData {
	return v.value
}

func (v *NullableCompanyMembershipDetailResponseData) Set(val *CompanyMembershipDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyMembershipDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyMembershipDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyMembershipDetailResponseData(val *CompanyMembershipDetailResponseData) *NullableCompanyMembershipDetailResponseData {
	return &NullableCompanyMembershipDetailResponseData{value: val, isSet: true}
}

func (v NullableCompanyMembershipDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyMembershipDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
