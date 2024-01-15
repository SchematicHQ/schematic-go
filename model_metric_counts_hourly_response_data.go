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

// checks if the MetricCountsHourlyResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricCountsHourlyResponseData{}

// MetricCountsHourlyResponseData struct for MetricCountsHourlyResponseData
type MetricCountsHourlyResponseData struct {
	CompanyId NullableString `json:"company_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	EnvironmentId string `json:"environment_id"`
	EventSubtype string `json:"event_subtype"`
	StartTime time.Time `json:"start_time"`
	UserId NullableString `json:"user_id,omitempty"`
	Value int32 `json:"value"`
}

// NewMetricCountsHourlyResponseData instantiates a new MetricCountsHourlyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricCountsHourlyResponseData(createdAt time.Time, environmentId string, eventSubtype string, startTime time.Time, value int32) *MetricCountsHourlyResponseData {
	this := MetricCountsHourlyResponseData{}
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.EventSubtype = eventSubtype
	this.StartTime = startTime
	this.Value = value
	return &this
}

// NewMetricCountsHourlyResponseDataWithDefaults instantiates a new MetricCountsHourlyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricCountsHourlyResponseDataWithDefaults() *MetricCountsHourlyResponseData {
	this := MetricCountsHourlyResponseData{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricCountsHourlyResponseData) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricCountsHourlyResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *MetricCountsHourlyResponseData) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *MetricCountsHourlyResponseData) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}
// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *MetricCountsHourlyResponseData) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *MetricCountsHourlyResponseData) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetricCountsHourlyResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetricCountsHourlyResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetricCountsHourlyResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *MetricCountsHourlyResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *MetricCountsHourlyResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *MetricCountsHourlyResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetEventSubtype returns the EventSubtype field value
func (o *MetricCountsHourlyResponseData) GetEventSubtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventSubtype
}

// GetEventSubtypeOk returns a tuple with the EventSubtype field value
// and a boolean to check if the value has been set.
func (o *MetricCountsHourlyResponseData) GetEventSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventSubtype, true
}

// SetEventSubtype sets field value
func (o *MetricCountsHourlyResponseData) SetEventSubtype(v string) {
	o.EventSubtype = v
}

// GetStartTime returns the StartTime field value
func (o *MetricCountsHourlyResponseData) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *MetricCountsHourlyResponseData) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *MetricCountsHourlyResponseData) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricCountsHourlyResponseData) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricCountsHourlyResponseData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MetricCountsHourlyResponseData) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *MetricCountsHourlyResponseData) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MetricCountsHourlyResponseData) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MetricCountsHourlyResponseData) UnsetUserId() {
	o.UserId.Unset()
}

// GetValue returns the Value field value
func (o *MetricCountsHourlyResponseData) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MetricCountsHourlyResponseData) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MetricCountsHourlyResponseData) SetValue(v int32) {
	o.Value = v
}

func (o MetricCountsHourlyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricCountsHourlyResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["event_subtype"] = o.EventSubtype
	toSerialize["start_time"] = o.StartTime
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableMetricCountsHourlyResponseData struct {
	value *MetricCountsHourlyResponseData
	isSet bool
}

func (v NullableMetricCountsHourlyResponseData) Get() *MetricCountsHourlyResponseData {
	return v.value
}

func (v *NullableMetricCountsHourlyResponseData) Set(val *MetricCountsHourlyResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricCountsHourlyResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricCountsHourlyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricCountsHourlyResponseData(val *MetricCountsHourlyResponseData) *NullableMetricCountsHourlyResponseData {
	return &NullableMetricCountsHourlyResponseData{value: val, isSet: true}
}

func (v NullableMetricCountsHourlyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricCountsHourlyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


