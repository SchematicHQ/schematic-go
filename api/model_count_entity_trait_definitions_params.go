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

// checks if the CountEntityTraitDefinitionsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountEntityTraitDefinitionsParams{}

// CountEntityTraitDefinitionsParams Input parameters
type CountEntityTraitDefinitionsParams struct {
	EntityType *string  `json:"entity_type,omitempty"`
	Ids        []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	TraitType            *string `json:"trait_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountEntityTraitDefinitionsParams CountEntityTraitDefinitionsParams

// NewCountEntityTraitDefinitionsParams instantiates a new CountEntityTraitDefinitionsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountEntityTraitDefinitionsParams() *CountEntityTraitDefinitionsParams {
	this := CountEntityTraitDefinitionsParams{}
	return &this
}

// NewCountEntityTraitDefinitionsParamsWithDefaults instantiates a new CountEntityTraitDefinitionsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountEntityTraitDefinitionsParamsWithDefaults() *CountEntityTraitDefinitionsParams {
	this := CountEntityTraitDefinitionsParams{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CountEntityTraitDefinitionsParams) SetEntityType(v string) {
	o.EntityType = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CountEntityTraitDefinitionsParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountEntityTraitDefinitionsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountEntityTraitDefinitionsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountEntityTraitDefinitionsParams) SetQ(v string) {
	o.Q = &v
}

// GetTraitType returns the TraitType field value if set, zero value otherwise.
func (o *CountEntityTraitDefinitionsParams) GetTraitType() string {
	if o == nil || IsNil(o.TraitType) {
		var ret string
		return ret
	}
	return *o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEntityTraitDefinitionsParams) GetTraitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TraitType) {
		return nil, false
	}
	return o.TraitType, true
}

// HasTraitType returns a boolean if a field has been set.
func (o *CountEntityTraitDefinitionsParams) HasTraitType() bool {
	if o != nil && !IsNil(o.TraitType) {
		return true
	}

	return false
}

// SetTraitType gets a reference to the given string and assigns it to the TraitType field.
func (o *CountEntityTraitDefinitionsParams) SetTraitType(v string) {
	o.TraitType = &v
}

func (o CountEntityTraitDefinitionsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountEntityTraitDefinitionsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityType) {
		toSerialize["entity_type"] = o.EntityType
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
	if !IsNil(o.TraitType) {
		toSerialize["trait_type"] = o.TraitType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountEntityTraitDefinitionsParams) UnmarshalJSON(data []byte) (err error) {
	varCountEntityTraitDefinitionsParams := _CountEntityTraitDefinitionsParams{}

	err = json.Unmarshal(data, &varCountEntityTraitDefinitionsParams)

	if err != nil {
		return err
	}

	*o = CountEntityTraitDefinitionsParams(varCountEntityTraitDefinitionsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_type")
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		delete(additionalProperties, "trait_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountEntityTraitDefinitionsParams struct {
	value *CountEntityTraitDefinitionsParams
	isSet bool
}

func (v NullableCountEntityTraitDefinitionsParams) Get() *CountEntityTraitDefinitionsParams {
	return v.value
}

func (v *NullableCountEntityTraitDefinitionsParams) Set(val *CountEntityTraitDefinitionsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountEntityTraitDefinitionsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountEntityTraitDefinitionsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountEntityTraitDefinitionsParams(val *CountEntityTraitDefinitionsParams) *NullableCountEntityTraitDefinitionsParams {
	return &NullableCountEntityTraitDefinitionsParams{value: val, isSet: true}
}

func (v NullableCountEntityTraitDefinitionsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountEntityTraitDefinitionsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
