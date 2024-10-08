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

// checks if the CountBillingProductsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountBillingProductsParams{}

// CountBillingProductsParams Input parameters
type CountBillingProductsParams struct {
	Ids []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32  `json:"limit,omitempty"`
	Name  *string `json:"name,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountBillingProductsParams CountBillingProductsParams

// NewCountBillingProductsParams instantiates a new CountBillingProductsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountBillingProductsParams() *CountBillingProductsParams {
	this := CountBillingProductsParams{}
	return &this
}

// NewCountBillingProductsParamsWithDefaults instantiates a new CountBillingProductsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountBillingProductsParamsWithDefaults() *CountBillingProductsParams {
	this := CountBillingProductsParams{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CountBillingProductsParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountBillingProductsParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CountBillingProductsParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CountBillingProductsParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountBillingProductsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountBillingProductsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountBillingProductsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountBillingProductsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CountBillingProductsParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountBillingProductsParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CountBillingProductsParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CountBillingProductsParams) SetName(v string) {
	o.Name = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountBillingProductsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountBillingProductsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountBillingProductsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountBillingProductsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountBillingProductsParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountBillingProductsParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountBillingProductsParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountBillingProductsParams) SetQ(v string) {
	o.Q = &v
}

func (o CountBillingProductsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountBillingProductsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

func (o *CountBillingProductsParams) UnmarshalJSON(data []byte) (err error) {
	varCountBillingProductsParams := _CountBillingProductsParams{}

	err = json.Unmarshal(data, &varCountBillingProductsParams)

	if err != nil {
		return err
	}

	*o = CountBillingProductsParams(varCountBillingProductsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "name")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountBillingProductsParams struct {
	value *CountBillingProductsParams
	isSet bool
}

func (v NullableCountBillingProductsParams) Get() *CountBillingProductsParams {
	return v.value
}

func (v *NullableCountBillingProductsParams) Set(val *CountBillingProductsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountBillingProductsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountBillingProductsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountBillingProductsParams(val *CountBillingProductsParams) *NullableCountBillingProductsParams {
	return &NullableCountBillingProductsParams{value: val, isSet: true}
}

func (v NullableCountBillingProductsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountBillingProductsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
