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

// checks if the GetEventSummariesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventSummariesParams{}

// GetEventSummariesParams Input parameters
type GetEventSummariesParams struct {
	EventSubtypes []string `json:"event_subtypes,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetEventSummariesParams GetEventSummariesParams

// NewGetEventSummariesParams instantiates a new GetEventSummariesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventSummariesParams() *GetEventSummariesParams {
	this := GetEventSummariesParams{}
	return &this
}

// NewGetEventSummariesParamsWithDefaults instantiates a new GetEventSummariesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventSummariesParamsWithDefaults() *GetEventSummariesParams {
	this := GetEventSummariesParams{}
	return &this
}

// GetEventSubtypes returns the EventSubtypes field value if set, zero value otherwise.
func (o *GetEventSummariesParams) GetEventSubtypes() []string {
	if o == nil || IsNil(o.EventSubtypes) {
		var ret []string
		return ret
	}
	return o.EventSubtypes
}

// GetEventSubtypesOk returns a tuple with the EventSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventSummariesParams) GetEventSubtypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EventSubtypes) {
		return nil, false
	}
	return o.EventSubtypes, true
}

// HasEventSubtypes returns a boolean if a field has been set.
func (o *GetEventSummariesParams) HasEventSubtypes() bool {
	if o != nil && !IsNil(o.EventSubtypes) {
		return true
	}

	return false
}

// SetEventSubtypes gets a reference to the given []string and assigns it to the EventSubtypes field.
func (o *GetEventSummariesParams) SetEventSubtypes(v []string) {
	o.EventSubtypes = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetEventSummariesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventSummariesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetEventSummariesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetEventSummariesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetEventSummariesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventSummariesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetEventSummariesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetEventSummariesParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *GetEventSummariesParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventSummariesParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *GetEventSummariesParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *GetEventSummariesParams) SetQ(v string) {
	o.Q = &v
}

func (o GetEventSummariesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventSummariesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventSubtypes) {
		toSerialize["event_subtypes"] = o.EventSubtypes
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetEventSummariesParams) UnmarshalJSON(data []byte) (err error) {
	varGetEventSummariesParams := _GetEventSummariesParams{}

	err = json.Unmarshal(data, &varGetEventSummariesParams)

	if err != nil {
		return err
	}

	*o = GetEventSummariesParams(varGetEventSummariesParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event_subtypes")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetEventSummariesParams struct {
	value *GetEventSummariesParams
	isSet bool
}

func (v NullableGetEventSummariesParams) Get() *GetEventSummariesParams {
	return v.value
}

func (v *NullableGetEventSummariesParams) Set(val *GetEventSummariesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventSummariesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventSummariesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventSummariesParams(val *GetEventSummariesParams) *NullableGetEventSummariesParams {
	return &NullableGetEventSummariesParams{value: val, isSet: true}
}

func (v NullableGetEventSummariesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventSummariesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
