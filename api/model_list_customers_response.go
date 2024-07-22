/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ListCustomersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomersResponse{}

// ListCustomersResponse struct for ListCustomersResponse
type ListCustomersResponse struct {
	// The returned resources
	Data   []BillingCustomerWithSubscriptionsResponseData `json:"data"`
	Params ListCustomersParams                            `json:"params"`
}

type _ListCustomersResponse ListCustomersResponse

// NewListCustomersResponse instantiates a new ListCustomersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomersResponse(data []BillingCustomerWithSubscriptionsResponseData, params ListCustomersParams) *ListCustomersResponse {
	this := ListCustomersResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListCustomersResponseWithDefaults instantiates a new ListCustomersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomersResponseWithDefaults() *ListCustomersResponse {
	this := ListCustomersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListCustomersResponse) GetData() []BillingCustomerWithSubscriptionsResponseData {
	if o == nil {
		var ret []BillingCustomerWithSubscriptionsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetDataOk() ([]BillingCustomerWithSubscriptionsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListCustomersResponse) SetData(v []BillingCustomerWithSubscriptionsResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListCustomersResponse) GetParams() ListCustomersParams {
	if o == nil {
		var ret ListCustomersParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetParamsOk() (*ListCustomersParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListCustomersResponse) SetParams(v ListCustomersParams) {
	o.Params = v
}

func (o ListCustomersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *ListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"params",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListCustomersResponse := _ListCustomersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCustomersResponse)

	if err != nil {
		return err
	}

	*o = ListCustomersResponse(varListCustomersResponse)

	return err
}

type NullableListCustomersResponse struct {
	value *ListCustomersResponse
	isSet bool
}

func (v NullableListCustomersResponse) Get() *ListCustomersResponse {
	return v.value
}

func (v *NullableListCustomersResponse) Set(val *ListCustomersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomersResponse(val *ListCustomersResponse) *NullableListCustomersResponse {
	return &NullableListCustomersResponse{value: val, isSet: true}
}

func (v NullableListCustomersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
