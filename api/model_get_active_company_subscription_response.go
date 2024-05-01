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

// checks if the GetActiveCompanySubscriptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActiveCompanySubscriptionResponse{}

// GetActiveCompanySubscriptionResponse struct for GetActiveCompanySubscriptionResponse
type GetActiveCompanySubscriptionResponse struct {
	// The returned resources
	Data   []CompanySubscriptionResponseData  `json:"data"`
	Params GetActiveCompanySubscriptionParams `json:"params"`
}

type _GetActiveCompanySubscriptionResponse GetActiveCompanySubscriptionResponse

// NewGetActiveCompanySubscriptionResponse instantiates a new GetActiveCompanySubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActiveCompanySubscriptionResponse(data []CompanySubscriptionResponseData, params GetActiveCompanySubscriptionParams) *GetActiveCompanySubscriptionResponse {
	this := GetActiveCompanySubscriptionResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetActiveCompanySubscriptionResponseWithDefaults instantiates a new GetActiveCompanySubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActiveCompanySubscriptionResponseWithDefaults() *GetActiveCompanySubscriptionResponse {
	this := GetActiveCompanySubscriptionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetActiveCompanySubscriptionResponse) GetData() []CompanySubscriptionResponseData {
	if o == nil {
		var ret []CompanySubscriptionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetActiveCompanySubscriptionResponse) GetDataOk() ([]CompanySubscriptionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetActiveCompanySubscriptionResponse) SetData(v []CompanySubscriptionResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetActiveCompanySubscriptionResponse) GetParams() GetActiveCompanySubscriptionParams {
	if o == nil {
		var ret GetActiveCompanySubscriptionParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetActiveCompanySubscriptionResponse) GetParamsOk() (*GetActiveCompanySubscriptionParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *GetActiveCompanySubscriptionResponse) SetParams(v GetActiveCompanySubscriptionParams) {
	o.Params = v
}

func (o GetActiveCompanySubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActiveCompanySubscriptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetActiveCompanySubscriptionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetActiveCompanySubscriptionResponse := _GetActiveCompanySubscriptionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetActiveCompanySubscriptionResponse)

	if err != nil {
		return err
	}

	*o = GetActiveCompanySubscriptionResponse(varGetActiveCompanySubscriptionResponse)

	return err
}

type NullableGetActiveCompanySubscriptionResponse struct {
	value *GetActiveCompanySubscriptionResponse
	isSet bool
}

func (v NullableGetActiveCompanySubscriptionResponse) Get() *GetActiveCompanySubscriptionResponse {
	return v.value
}

func (v *NullableGetActiveCompanySubscriptionResponse) Set(val *GetActiveCompanySubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActiveCompanySubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActiveCompanySubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActiveCompanySubscriptionResponse(val *GetActiveCompanySubscriptionResponse) *NullableGetActiveCompanySubscriptionResponse {
	return &NullableGetActiveCompanySubscriptionResponse{value: val, isSet: true}
}

func (v NullableGetActiveCompanySubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActiveCompanySubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
