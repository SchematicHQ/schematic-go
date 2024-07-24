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

// checks if the CountPlanEntitlementsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountPlanEntitlementsParams{}

// CountPlanEntitlementsParams Input parameters
type CountPlanEntitlementsParams struct {
	FeatureId  *string  `json:"feature_id,omitempty"`
	FeatureIds []string `json:"feature_ids,omitempty"`
	Ids        []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32   `json:"offset,omitempty"`
	PlanId               *string  `json:"plan_id,omitempty"`
	PlanIds              []string `json:"plan_ids,omitempty"`
	Q                    *string  `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountPlanEntitlementsParams CountPlanEntitlementsParams

// NewCountPlanEntitlementsParams instantiates a new CountPlanEntitlementsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountPlanEntitlementsParams() *CountPlanEntitlementsParams {
	this := CountPlanEntitlementsParams{}
	return &this
}

// NewCountPlanEntitlementsParamsWithDefaults instantiates a new CountPlanEntitlementsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountPlanEntitlementsParamsWithDefaults() *CountPlanEntitlementsParams {
	this := CountPlanEntitlementsParams{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *CountPlanEntitlementsParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetFeatureIds() []string {
	if o == nil || IsNil(o.FeatureIds) {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureIds) {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasFeatureIds() bool {
	if o != nil && !IsNil(o.FeatureIds) {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *CountPlanEntitlementsParams) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CountPlanEntitlementsParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountPlanEntitlementsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountPlanEntitlementsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *CountPlanEntitlementsParams) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanIds returns the PlanIds field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetPlanIds() []string {
	if o == nil || IsNil(o.PlanIds) {
		var ret []string
		return ret
	}
	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetPlanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlanIds) {
		return nil, false
	}
	return o.PlanIds, true
}

// HasPlanIds returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasPlanIds() bool {
	if o != nil && !IsNil(o.PlanIds) {
		return true
	}

	return false
}

// SetPlanIds gets a reference to the given []string and assigns it to the PlanIds field.
func (o *CountPlanEntitlementsParams) SetPlanIds(v []string) {
	o.PlanIds = v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountPlanEntitlementsParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountPlanEntitlementsParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountPlanEntitlementsParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountPlanEntitlementsParams) SetQ(v string) {
	o.Q = &v
}

func (o CountPlanEntitlementsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountPlanEntitlementsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.FeatureIds) {
		toSerialize["feature_ids"] = o.FeatureIds
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
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.PlanIds) {
		toSerialize["plan_ids"] = o.PlanIds
	}
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountPlanEntitlementsParams) UnmarshalJSON(data []byte) (err error) {
	varCountPlanEntitlementsParams := _CountPlanEntitlementsParams{}

	err = json.Unmarshal(data, &varCountPlanEntitlementsParams)

	if err != nil {
		return err
	}

	*o = CountPlanEntitlementsParams(varCountPlanEntitlementsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature_id")
		delete(additionalProperties, "feature_ids")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "plan_ids")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountPlanEntitlementsParams struct {
	value *CountPlanEntitlementsParams
	isSet bool
}

func (v NullableCountPlanEntitlementsParams) Get() *CountPlanEntitlementsParams {
	return v.value
}

func (v *NullableCountPlanEntitlementsParams) Set(val *CountPlanEntitlementsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountPlanEntitlementsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountPlanEntitlementsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountPlanEntitlementsParams(val *CountPlanEntitlementsParams) *NullableCountPlanEntitlementsParams {
	return &NullableCountPlanEntitlementsParams{value: val, isSet: true}
}

func (v NullableCountPlanEntitlementsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountPlanEntitlementsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
