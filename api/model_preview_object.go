/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the PreviewObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewObject{}

// PreviewObject struct for PreviewObject
type PreviewObject struct {
	Id                   string         `json:"id"`
	ImageUrl             NullableString `json:"image_url,omitempty"`
	Name                 string         `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _PreviewObject PreviewObject

// NewPreviewObject instantiates a new PreviewObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewObject(id string, name string) *PreviewObject {
	this := PreviewObject{}
	this.Id = id
	this.Name = name
	return &this
}

// NewPreviewObjectWithDefaults instantiates a new PreviewObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewObjectWithDefaults() *PreviewObject {
	this := PreviewObject{}
	return &this
}

// GetId returns the Id field value
func (o *PreviewObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PreviewObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PreviewObject) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreviewObject) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewObject) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PreviewObject) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *PreviewObject) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *PreviewObject) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *PreviewObject) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetName returns the Name field value
func (o *PreviewObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PreviewObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PreviewObject) SetName(v string) {
	o.Name = v
}

func (o PreviewObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreviewObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varPreviewObject := _PreviewObject{}

	err = json.Unmarshal(data, &varPreviewObject)

	if err != nil {
		return err
	}

	*o = PreviewObject(varPreviewObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreviewObject struct {
	value *PreviewObject
	isSet bool
}

func (v NullablePreviewObject) Get() *PreviewObject {
	return v.value
}

func (v *NullablePreviewObject) Set(val *PreviewObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewObject(val *PreviewObject) *NullablePreviewObject {
	return &NullablePreviewObject{value: val, isSet: true}
}

func (v NullablePreviewObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
