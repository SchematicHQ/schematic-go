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

// checks if the ListFlagsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFlagsParams{}

// ListFlagsParams Input parameters
type ListFlagsParams struct {
	FeatureId *string  `json:"feature_id,omitempty"`
	Ids       []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListFlagsParams ListFlagsParams

// NewListFlagsParams instantiates a new ListFlagsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFlagsParams() *ListFlagsParams {
	this := ListFlagsParams{}
	return &this
}

// NewListFlagsParamsWithDefaults instantiates a new ListFlagsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFlagsParamsWithDefaults() *ListFlagsParams {
	this := ListFlagsParams{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ListFlagsParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlagsParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ListFlagsParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *ListFlagsParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListFlagsParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlagsParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListFlagsParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ListFlagsParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListFlagsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlagsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListFlagsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListFlagsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListFlagsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlagsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListFlagsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListFlagsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ListFlagsParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlagsParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *ListFlagsParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ListFlagsParams) SetQ(v string) {
	o.Q = &v
}

func (o ListFlagsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFlagsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
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

func (o *ListFlagsParams) UnmarshalJSON(data []byte) (err error) {
	varListFlagsParams := _ListFlagsParams{}

	err = json.Unmarshal(data, &varListFlagsParams)

	if err != nil {
		return err
	}

	*o = ListFlagsParams(varListFlagsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature_id")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFlagsParams struct {
	value *ListFlagsParams
	isSet bool
}

func (v NullableListFlagsParams) Get() *ListFlagsParams {
	return v.value
}

func (v *NullableListFlagsParams) Set(val *ListFlagsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListFlagsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListFlagsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFlagsParams(val *ListFlagsParams) *NullableListFlagsParams {
	return &NullableListFlagsParams{value: val, isSet: true}
}

func (v NullableListFlagsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFlagsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
