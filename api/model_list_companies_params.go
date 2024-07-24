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

// checks if the ListCompaniesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCompaniesParams{}

// ListCompaniesParams Input parameters
type ListCompaniesParams struct {
	Ids []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int32  `json:"offset,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	// Search filter
	Q *string `json:"q,omitempty"`
	// Filter out companies that already have a company override for the specified feature ID
	WithoutFeatureOverrideFor *string `json:"without_feature_override_for,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _ListCompaniesParams ListCompaniesParams

// NewListCompaniesParams instantiates a new ListCompaniesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCompaniesParams() *ListCompaniesParams {
	this := ListCompaniesParams{}
	return &this
}

// NewListCompaniesParamsWithDefaults instantiates a new ListCompaniesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCompaniesParamsWithDefaults() *ListCompaniesParams {
	this := ListCompaniesParams{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ListCompaniesParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListCompaniesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListCompaniesParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ListCompaniesParams) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ListCompaniesParams) SetQ(v string) {
	o.Q = &v
}

// GetWithoutFeatureOverrideFor returns the WithoutFeatureOverrideFor field value if set, zero value otherwise.
func (o *ListCompaniesParams) GetWithoutFeatureOverrideFor() string {
	if o == nil || IsNil(o.WithoutFeatureOverrideFor) {
		var ret string
		return ret
	}
	return *o.WithoutFeatureOverrideFor
}

// GetWithoutFeatureOverrideForOk returns a tuple with the WithoutFeatureOverrideFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompaniesParams) GetWithoutFeatureOverrideForOk() (*string, bool) {
	if o == nil || IsNil(o.WithoutFeatureOverrideFor) {
		return nil, false
	}
	return o.WithoutFeatureOverrideFor, true
}

// HasWithoutFeatureOverrideFor returns a boolean if a field has been set.
func (o *ListCompaniesParams) HasWithoutFeatureOverrideFor() bool {
	if o != nil && !IsNil(o.WithoutFeatureOverrideFor) {
		return true
	}

	return false
}

// SetWithoutFeatureOverrideFor gets a reference to the given string and assigns it to the WithoutFeatureOverrideFor field.
func (o *ListCompaniesParams) SetWithoutFeatureOverrideFor(v string) {
	o.WithoutFeatureOverrideFor = &v
}

func (o ListCompaniesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCompaniesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}
	if !IsNil(o.WithoutFeatureOverrideFor) {
		toSerialize["without_feature_override_for"] = o.WithoutFeatureOverrideFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCompaniesParams) UnmarshalJSON(data []byte) (err error) {
	varListCompaniesParams := _ListCompaniesParams{}

	err = json.Unmarshal(data, &varListCompaniesParams)

	if err != nil {
		return err
	}

	*o = ListCompaniesParams(varListCompaniesParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "q")
		delete(additionalProperties, "without_feature_override_for")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCompaniesParams struct {
	value *ListCompaniesParams
	isSet bool
}

func (v NullableListCompaniesParams) Get() *ListCompaniesParams {
	return v.value
}

func (v *NullableListCompaniesParams) Set(val *ListCompaniesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompaniesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompaniesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompaniesParams(val *ListCompaniesParams) *NullableListCompaniesParams {
	return &NullableListCompaniesParams{value: val, isSet: true}
}

func (v NullableListCompaniesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompaniesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
