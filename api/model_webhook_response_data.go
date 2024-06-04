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
	"time"
)

// checks if the WebhookResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponseData{}

// WebhookResponseData The updated resource
type WebhookResponseData struct {
	CreatedAt    time.Time `json:"created_at"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	RequestTypes []string  `json:"request_types"`
	Status       string    `json:"status"`
	UpdatedAt    time.Time `json:"updated_at"`
	Url          string    `json:"url"`
}

type _WebhookResponseData WebhookResponseData

// NewWebhookResponseData instantiates a new WebhookResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponseData(createdAt time.Time, id string, name string, requestTypes []string, status string, updatedAt time.Time, url string) *WebhookResponseData {
	this := WebhookResponseData{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.RequestTypes = requestTypes
	this.Status = status
	this.UpdatedAt = updatedAt
	this.Url = url
	return &this
}

// NewWebhookResponseDataWithDefaults instantiates a new WebhookResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseDataWithDefaults() *WebhookResponseData {
	this := WebhookResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *WebhookResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WebhookResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookResponseData) SetName(v string) {
	o.Name = v
}

// GetRequestTypes returns the RequestTypes field value
func (o *WebhookResponseData) GetRequestTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequestTypes
}

// GetRequestTypesOk returns a tuple with the RequestTypes field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetRequestTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestTypes, true
}

// SetRequestTypes sets field value
func (o *WebhookResponseData) SetRequestTypes(v []string) {
	o.RequestTypes = v
}

// GetStatus returns the Status field value
func (o *WebhookResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebhookResponseData) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WebhookResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WebhookResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUrl returns the Url field value
func (o *WebhookResponseData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookResponseData) SetUrl(v string) {
	o.Url = v
}

func (o WebhookResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["request_types"] = o.RequestTypes
	toSerialize["status"] = o.Status
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *WebhookResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"id",
		"name",
		"request_types",
		"status",
		"updated_at",
		"url",
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

	varWebhookResponseData := _WebhookResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResponseData)

	if err != nil {
		return err
	}

	*o = WebhookResponseData(varWebhookResponseData)

	return err
}

type NullableWebhookResponseData struct {
	value *WebhookResponseData
	isSet bool
}

func (v NullableWebhookResponseData) Get() *WebhookResponseData {
	return v.value
}

func (v *NullableWebhookResponseData) Set(val *WebhookResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponseData(val *WebhookResponseData) *NullableWebhookResponseData {
	return &NullableWebhookResponseData{value: val, isSet: true}
}

func (v NullableWebhookResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
