/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the ListMetricCountsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetricCountsParams{}

// ListMetricCountsParams Input parameters
type ListMetricCountsParams struct {
	CompanyId     *string    `json:"company_id,omitempty"`
	CompanyIds    []string   `json:"company_ids,omitempty"`
	EndTime       *time.Time `json:"end_time,omitempty"`
	EventSubtype  *string    `json:"event_subtype,omitempty"`
	EventSubtypes []string   `json:"event_subtypes,omitempty"`
	Grouping      *string    `json:"grouping,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32     `json:"offset,omitempty"`
	StartTime            *time.Time `json:"start_time,omitempty"`
	UserId               *string    `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListMetricCountsParams ListMetricCountsParams

// NewListMetricCountsParams instantiates a new ListMetricCountsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetricCountsParams() *ListMetricCountsParams {
	this := ListMetricCountsParams{}
	return &this
}

// NewListMetricCountsParamsWithDefaults instantiates a new ListMetricCountsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetricCountsParamsWithDefaults() *ListMetricCountsParams {
	this := ListMetricCountsParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *ListMetricCountsParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCompanyIds returns the CompanyIds field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetCompanyIds() []string {
	if o == nil || IsNil(o.CompanyIds) {
		var ret []string
		return ret
	}
	return o.CompanyIds
}

// GetCompanyIdsOk returns a tuple with the CompanyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetCompanyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CompanyIds) {
		return nil, false
	}
	return o.CompanyIds, true
}

// HasCompanyIds returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasCompanyIds() bool {
	if o != nil && !IsNil(o.CompanyIds) {
		return true
	}

	return false
}

// SetCompanyIds gets a reference to the given []string and assigns it to the CompanyIds field.
func (o *ListMetricCountsParams) SetCompanyIds(v []string) {
	o.CompanyIds = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ListMetricCountsParams) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetEventSubtype returns the EventSubtype field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetEventSubtype() string {
	if o == nil || IsNil(o.EventSubtype) {
		var ret string
		return ret
	}
	return *o.EventSubtype
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetEventSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventSubtype) {
		return nil, false
	}
	return o.EventSubtype, true
}

// HasEventSubtype returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasEventSubtype() bool {
	if o != nil && !IsNil(o.EventSubtype) {
		return true
	}

	return false
}

// SetEventSubtype gets a reference to the given string and assigns it to the EventSubtype field.
func (o *ListMetricCountsParams) SetEventSubtype(v string) {
	o.EventSubtype = &v
}

// GetEventSubtypes returns the EventSubtypes field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetEventSubtypes() []string {
	if o == nil || IsNil(o.EventSubtypes) {
		var ret []string
		return ret
	}
	return o.EventSubtypes
}

// GetEventSubtypesOk returns a tuple with the EventSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetEventSubtypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EventSubtypes) {
		return nil, false
	}
	return o.EventSubtypes, true
}

// HasEventSubtypes returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasEventSubtypes() bool {
	if o != nil && !IsNil(o.EventSubtypes) {
		return true
	}

	return false
}

// SetEventSubtypes gets a reference to the given []string and assigns it to the EventSubtypes field.
func (o *ListMetricCountsParams) SetEventSubtypes(v []string) {
	o.EventSubtypes = v
}

// GetGrouping returns the Grouping field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetGrouping() string {
	if o == nil || IsNil(o.Grouping) {
		var ret string
		return ret
	}
	return *o.Grouping
}

// GetGroupingOk returns a tuple with the Grouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetGroupingOk() (*string, bool) {
	if o == nil || IsNil(o.Grouping) {
		return nil, false
	}
	return o.Grouping, true
}

// HasGrouping returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasGrouping() bool {
	if o != nil && !IsNil(o.Grouping) {
		return true
	}

	return false
}

// SetGrouping gets a reference to the given string and assigns it to the Grouping field.
func (o *ListMetricCountsParams) SetGrouping(v string) {
	o.Grouping = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListMetricCountsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListMetricCountsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ListMetricCountsParams) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ListMetricCountsParams) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetricCountsParams) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ListMetricCountsParams) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ListMetricCountsParams) SetUserId(v string) {
	o.UserId = &v
}

func (o ListMetricCountsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetricCountsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.CompanyIds) {
		toSerialize["company_ids"] = o.CompanyIds
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.EventSubtype) {
		toSerialize["event_subtype"] = o.EventSubtype
	}
	if !IsNil(o.EventSubtypes) {
		toSerialize["event_subtypes"] = o.EventSubtypes
	}
	if !IsNil(o.Grouping) {
		toSerialize["grouping"] = o.Grouping
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListMetricCountsParams) UnmarshalJSON(data []byte) (err error) {
	varListMetricCountsParams := _ListMetricCountsParams{}

	err = json.Unmarshal(data, &varListMetricCountsParams)

	if err != nil {
		return err
	}

	*o = ListMetricCountsParams(varListMetricCountsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "company_ids")
		delete(additionalProperties, "end_time")
		delete(additionalProperties, "event_subtype")
		delete(additionalProperties, "event_subtypes")
		delete(additionalProperties, "grouping")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListMetricCountsParams struct {
	value *ListMetricCountsParams
	isSet bool
}

func (v NullableListMetricCountsParams) Get() *ListMetricCountsParams {
	return v.value
}

func (v *NullableListMetricCountsParams) Set(val *ListMetricCountsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetricCountsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetricCountsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetricCountsParams(val *ListMetricCountsParams) *NullableListMetricCountsParams {
	return &NullableListMetricCountsParams{value: val, isSet: true}
}

func (v NullableListMetricCountsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetricCountsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}