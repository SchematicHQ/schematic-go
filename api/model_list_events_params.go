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

// checks if the ListEventsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEventsParams{}

// ListEventsParams Input parameters
type ListEventsParams struct {
	CompanyId    *string  `json:"company_id,omitempty"`
	EventSubtype *string  `json:"event_subtype,omitempty"`
	EventTypes   []string `json:"event_types,omitempty"`
	FlagId       *string  `json:"flag_id,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int32  `json:"offset,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// NewListEventsParams instantiates a new ListEventsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEventsParams() *ListEventsParams {
	this := ListEventsParams{}
	return &this
}

// NewListEventsParamsWithDefaults instantiates a new ListEventsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEventsParamsWithDefaults() *ListEventsParams {
	this := ListEventsParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ListEventsParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ListEventsParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *ListEventsParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise.
func (o *ListEventsParams) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype) {
		var ret string
		return ret
	}
	return *o.EventSubtype
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetEventSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventSubtype) {
		return nil, false
	}
	return o.EventSubtype, true
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *ListEventsParams) HasEventSubtype() bool {
	if o != nil && !IsNil(o.EventSubtype) {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given string and assigns it to the EventSubtype field.
func (o *ListEventsParams) SetEventSubtype(v string) {
	o.EventSubtype = &v
}

// GetEventTypes returns the EventTypes field value if set, zero value otherwise.
func (o *ListEventsParams) GetEventTypes() []string {
	if o == nil || IsNil(o.EventTypes) {
		var ret []string
		return ret
	}
	return o.EventTypes
}

// GetEventTypesOk returns a tuple with the EventTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetEventTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EventTypes) {
		return nil, false
	}
	return o.EventTypes, true
}

// HasEventTypes returns a boolean if a field has been set.
func (o *ListEventsParams) HasEventTypes() bool {
	if o != nil && !IsNil(o.EventTypes) {
		return true
	}

	return false
}

// SetEventTypes gets a reference to the given []string and assigns it to the EventTypes field.
func (o *ListEventsParams) SetEventTypes(v []string) {
	o.EventTypes = v
}

// GetFlagId returns the FlagId field value if set, zero value otherwise.
func (o *ListEventsParams) GetFlagId() string {
	if o == nil || IsNil(o.FlagId) {
		var ret string
		return ret
	}
	return *o.FlagId
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetFlagIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlagId) {
		return nil, false
	}
	return o.FlagId, true
}

// HasFlagId returns a boolean if a field has been set.
func (o *ListEventsParams) HasFlagId() bool {
	if o != nil && !IsNil(o.FlagId) {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given string and assigns it to the FlagId field.
func (o *ListEventsParams) SetFlagId(v string) {
	o.FlagId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListEventsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListEventsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListEventsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListEventsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListEventsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListEventsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ListEventsParams) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventsParams) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ListEventsParams) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ListEventsParams) SetUserId(v string) {
	o.UserId = &v
}

func (o ListEventsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEventsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.EventSubtype) {
		toSerialize["event_subtype"] = o.EventSubtype
	}
	if !IsNil(o.EventTypes) {
		toSerialize["event_types"] = o.EventTypes
	}
	if !IsNil(o.FlagId) {
		toSerialize["flag_id"] = o.FlagId
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

type NullableListEventsParams struct {
	value *ListEventsParams
	isSet bool
}

func (v NullableListEventsParams) Get() *ListEventsParams {
	return v.value
}

func (v *NullableListEventsParams) Set(val *ListEventsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListEventsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListEventsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEventsParams(val *ListEventsParams) *NullableListEventsParams {
	return &NullableListEventsParams{value: val, isSet: true}
}

func (v NullableListEventsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEventsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
