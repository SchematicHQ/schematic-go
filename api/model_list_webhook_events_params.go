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

// checks if the ListWebhookEventsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookEventsParams{}

// ListWebhookEventsParams Input parameters
type ListWebhookEventsParams struct {
	Ids []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	WebhookId            *string `json:"webhook_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListWebhookEventsParams ListWebhookEventsParams

// NewListWebhookEventsParams instantiates a new ListWebhookEventsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookEventsParams() *ListWebhookEventsParams {
	this := ListWebhookEventsParams{}
	return &this
}

// NewListWebhookEventsParamsWithDefaults instantiates a new ListWebhookEventsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookEventsParamsWithDefaults() *ListWebhookEventsParams {
	this := ListWebhookEventsParams{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ListWebhookEventsParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ListWebhookEventsParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ListWebhookEventsParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListWebhookEventsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListWebhookEventsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListWebhookEventsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListWebhookEventsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListWebhookEventsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListWebhookEventsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ListWebhookEventsParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *ListWebhookEventsParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ListWebhookEventsParams) SetQ(v string) {
	o.Q = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *ListWebhookEventsParams) GetWebhookId() string {
	if o == nil || IsNil(o.WebhookId) {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsParams) GetWebhookIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookId) {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *ListWebhookEventsParams) HasWebhookId() bool {
	if o != nil && !IsNil(o.WebhookId) {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *ListWebhookEventsParams) SetWebhookId(v string) {
	o.WebhookId = &v
}

func (o ListWebhookEventsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookEventsParams) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}
	if !IsNil(o.WebhookId) {
		toSerialize["webhook_id"] = o.WebhookId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListWebhookEventsParams) UnmarshalJSON(data []byte) (err error) {
	varListWebhookEventsParams := _ListWebhookEventsParams{}

	err = json.Unmarshal(data, &varListWebhookEventsParams)

	if err != nil {
		return err
	}

	*o = ListWebhookEventsParams(varListWebhookEventsParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		delete(additionalProperties, "webhook_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListWebhookEventsParams struct {
	value *ListWebhookEventsParams
	isSet bool
}

func (v NullableListWebhookEventsParams) Get() *ListWebhookEventsParams {
	return v.value
}

func (v *NullableListWebhookEventsParams) Set(val *ListWebhookEventsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookEventsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookEventsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookEventsParams(val *ListWebhookEventsParams) *NullableListWebhookEventsParams {
	return &NullableListWebhookEventsParams{value: val, isSet: true}
}

func (v NullableListWebhookEventsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookEventsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}