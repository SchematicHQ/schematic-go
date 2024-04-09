/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CountEventsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountEventsParams{}

// CountEventsParams Input parameters
type CountEventsParams struct {
	CompanyId    *string `json:"company_id,omitempty"`
	EventSubtype *string `json:"event_subtype,omitempty"`
	Limit        *int32  `json:"limit,omitempty"`
	Offset       *int32  `json:"offset,omitempty"`
	UserId       *string `json:"user_id,omitempty"`
}

// NewCountEventsParams instantiates a new CountEventsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountEventsParams() *CountEventsParams {
	this := CountEventsParams{}
	return &this
}

// NewCountEventsParamsWithDefaults instantiates a new CountEventsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountEventsParamsWithDefaults() *CountEventsParams {
	this := CountEventsParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *CountEventsParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventsParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CountEventsParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *CountEventsParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise.
func (o *CountEventsParams) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype) {
		var ret string
		return ret
	}
	return *o.EventSubtype
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventsParams) GetEventSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventSubtype) {
		return nil, false
	}
	return o.EventSubtype, true
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *CountEventsParams) HasEventSubtype() bool {
	if o != nil && !IsNil(o.EventSubtype) {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given string and assigns it to the EventSubtype field.
func (o *CountEventsParams) SetEventSubtype(v string) {
	o.EventSubtype = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountEventsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountEventsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountEventsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountEventsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountEventsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountEventsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CountEventsParams) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventsParams) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CountEventsParams) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CountEventsParams) SetUserId(v string) {
	o.UserId = &v
}

func (o CountEventsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountEventsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.EventSubtype) {
		toSerialize["event_subtype"] = o.EventSubtype
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableCountEventsParams struct {
	value *CountEventsParams
	isSet bool
}

func (v NullableCountEventsParams) Get() *CountEventsParams {
	return v.value
}

func (v *NullableCountEventsParams) Set(val *CountEventsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountEventsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountEventsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountEventsParams(val *CountEventsParams) *NullableCountEventsParams {
	return &NullableCountEventsParams{value: val, isSet: true}
}

func (v NullableCountEventsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountEventsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
