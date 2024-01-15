/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
)

// checks if the LatestFlagChecksParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LatestFlagChecksParams{}

// LatestFlagChecksParams Input parameters
type LatestFlagChecksParams struct {
	FlagId *string `json:"flag_id,omitempty"`
	FlagIds []string `json:"flag_ids,omitempty"`
	Id *string `json:"id,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
}

// NewLatestFlagChecksParams instantiates a new LatestFlagChecksParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLatestFlagChecksParams() *LatestFlagChecksParams {
	this := LatestFlagChecksParams{}
	return &this
}

// NewLatestFlagChecksParamsWithDefaults instantiates a new LatestFlagChecksParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLatestFlagChecksParamsWithDefaults() *LatestFlagChecksParams {
	this := LatestFlagChecksParams{}
	return &this
}

// GetFlagId returns the FlagId field value if set, zero value otherwise.
func (o *LatestFlagChecksParams) GetFlagId() string {
	if o == nil || IsNil(o.FlagId) {
		var ret string
		return ret
	}
	return *o.FlagId
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestFlagChecksParams) GetFlagIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlagId) {
		return nil, false
	}
	return o.FlagId, true
}

// HasFlagId returns a boolean if a field has been set.
func (o *LatestFlagChecksParams) HasFlagId() bool {
	if o != nil && !IsNil(o.FlagId) {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given string and assigns it to the FlagId field.
func (o *LatestFlagChecksParams) SetFlagId(v string) {
	o.FlagId = &v
}

// GetFlagIds returns the FlagIds field value if set, zero value otherwise.
func (o *LatestFlagChecksParams) GetFlagIds() []string {
	if o == nil || IsNil(o.FlagIds) {
		var ret []string
		return ret
	}
	return o.FlagIds
}

// GetFlagIdsOk returns a tuple with the FlagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestFlagChecksParams) GetFlagIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FlagIds) {
		return nil, false
	}
	return o.FlagIds, true
}

// HasFlagIds returns a boolean if a field has been set.
func (o *LatestFlagChecksParams) HasFlagIds() bool {
	if o != nil && !IsNil(o.FlagIds) {
		return true
	}

	return false
}

// SetFlagIds gets a reference to the given []string and assigns it to the FlagIds field.
func (o *LatestFlagChecksParams) SetFlagIds(v []string) {
	o.FlagIds = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LatestFlagChecksParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestFlagChecksParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LatestFlagChecksParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LatestFlagChecksParams) SetId(v string) {
	o.Id = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *LatestFlagChecksParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestFlagChecksParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *LatestFlagChecksParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *LatestFlagChecksParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *LatestFlagChecksParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LatestFlagChecksParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *LatestFlagChecksParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *LatestFlagChecksParams) SetOffset(v int32) {
	o.Offset = &v
}

func (o LatestFlagChecksParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LatestFlagChecksParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlagId) {
		toSerialize["flag_id"] = o.FlagId
	}
	if !IsNil(o.FlagIds) {
		toSerialize["flag_ids"] = o.FlagIds
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableLatestFlagChecksParams struct {
	value *LatestFlagChecksParams
	isSet bool
}

func (v NullableLatestFlagChecksParams) Get() *LatestFlagChecksParams {
	return v.value
}

func (v *NullableLatestFlagChecksParams) Set(val *LatestFlagChecksParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLatestFlagChecksParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLatestFlagChecksParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatestFlagChecksParams(val *LatestFlagChecksParams) *NullableLatestFlagChecksParams {
	return &NullableLatestFlagChecksParams{value: val, isSet: true}
}

func (v NullableLatestFlagChecksParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatestFlagChecksParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


