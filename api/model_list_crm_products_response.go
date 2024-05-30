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

// checks if the ListCRMProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCRMProductsResponse{}

// ListCRMProductsResponse struct for ListCRMProductsResponse
type ListCRMProductsResponse struct {
	// The returned resources
	Data   []CRMProductResponseData `json:"data"`
	Params ListCRMProductsParams    `json:"params"`
}

type _ListCRMProductsResponse ListCRMProductsResponse

// NewListCRMProductsResponse instantiates a new ListCRMProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCRMProductsResponse(data []CRMProductResponseData, params ListCRMProductsParams) *ListCRMProductsResponse {
	this := ListCRMProductsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListCRMProductsResponseWithDefaults instantiates a new ListCRMProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCRMProductsResponseWithDefaults() *ListCRMProductsResponse {
	this := ListCRMProductsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListCRMProductsResponse) GetData() []CRMProductResponseData {
	if o == nil {
		var ret []CRMProductResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCRMProductsResponse) GetDataOk() ([]CRMProductResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListCRMProductsResponse) SetData(v []CRMProductResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListCRMProductsResponse) GetParams() ListCRMProductsParams {
	if o == nil {
		var ret ListCRMProductsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListCRMProductsResponse) GetParamsOk() (*ListCRMProductsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListCRMProductsResponse) SetParams(v ListCRMProductsParams) {
	o.Params = v
}

func (o ListCRMProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCRMProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *ListCRMProductsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListCRMProductsResponse := _ListCRMProductsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCRMProductsResponse)

	if err != nil {
		return err
	}

	*o = ListCRMProductsResponse(varListCRMProductsResponse)

	return err
}

type NullableListCRMProductsResponse struct {
	value *ListCRMProductsResponse
	isSet bool
}

func (v NullableListCRMProductsResponse) Get() *ListCRMProductsResponse {
	return v.value
}

func (v *NullableListCRMProductsResponse) Set(val *ListCRMProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCRMProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCRMProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCRMProductsResponse(val *ListCRMProductsResponse) *NullableListCRMProductsResponse {
	return &NullableListCRMProductsResponse{value: val, isSet: true}
}

func (v NullableListCRMProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCRMProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}