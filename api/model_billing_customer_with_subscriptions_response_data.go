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

// checks if the BillingCustomerWithSubscriptionsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingCustomerWithSubscriptionsResponseData{}

// BillingCustomerWithSubscriptionsResponseData struct for BillingCustomerWithSubscriptionsResponseData
type BillingCustomerWithSubscriptionsResponseData struct {
	CompanyId      NullableString                `json:"company_id,omitempty"`
	DeletedAt      NullableTime                  `json:"deleted_at,omitempty"`
	Email          string                        `json:"email"`
	ExternalId     string                        `json:"external_id"`
	FailedToImport bool                          `json:"failed_to_import"`
	Id             string                        `json:"id"`
	Name           string                        `json:"name"`
	Subscriptions  []BillingCustomerSubscription `json:"subscriptions"`
	UpdatedAt      time.Time                     `json:"updated_at"`
}

type _BillingCustomerWithSubscriptionsResponseData BillingCustomerWithSubscriptionsResponseData

// NewBillingCustomerWithSubscriptionsResponseData instantiates a new BillingCustomerWithSubscriptionsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingCustomerWithSubscriptionsResponseData(email string, externalId string, failedToImport bool, id string, name string, subscriptions []BillingCustomerSubscription, updatedAt time.Time) *BillingCustomerWithSubscriptionsResponseData {
	this := BillingCustomerWithSubscriptionsResponseData{}
	this.Email = email
	this.ExternalId = externalId
	this.FailedToImport = failedToImport
	this.Id = id
	this.Name = name
	this.Subscriptions = subscriptions
	this.UpdatedAt = updatedAt
	return &this
}

// NewBillingCustomerWithSubscriptionsResponseDataWithDefaults instantiates a new BillingCustomerWithSubscriptionsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingCustomerWithSubscriptionsResponseDataWithDefaults() *BillingCustomerWithSubscriptionsResponseData {
	this := BillingCustomerWithSubscriptionsResponseData{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingCustomerWithSubscriptionsResponseData) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingCustomerWithSubscriptionsResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *BillingCustomerWithSubscriptionsResponseData) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *BillingCustomerWithSubscriptionsResponseData) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *BillingCustomerWithSubscriptionsResponseData) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingCustomerWithSubscriptionsResponseData) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingCustomerWithSubscriptionsResponseData) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *BillingCustomerWithSubscriptionsResponseData) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *BillingCustomerWithSubscriptionsResponseData) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *BillingCustomerWithSubscriptionsResponseData) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetEmail returns the Email field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetEmail(v string) {
	o.Email = v
}

// GetExternalId returns the ExternalId field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetExternalId(v string) {
	o.ExternalId = v
}

// GetFailedToImport returns the FailedToImport field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetFailedToImport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FailedToImport
}

// GetFailedToImportOk returns a tuple with the FailedToImport field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetFailedToImportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedToImport, true
}

// SetFailedToImport sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetFailedToImport(v bool) {
	o.FailedToImport = v
}

// GetId returns the Id field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetName(v string) {
	o.Name = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetSubscriptions() []BillingCustomerSubscription {
	if o == nil {
		var ret []BillingCustomerSubscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetSubscriptionsOk() ([]BillingCustomerSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetSubscriptions(v []BillingCustomerSubscription) {
	o.Subscriptions = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BillingCustomerWithSubscriptionsResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingCustomerWithSubscriptionsResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BillingCustomerWithSubscriptionsResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o BillingCustomerWithSubscriptionsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingCustomerWithSubscriptionsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	toSerialize["email"] = o.Email
	toSerialize["external_id"] = o.ExternalId
	toSerialize["failed_to_import"] = o.FailedToImport
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *BillingCustomerWithSubscriptionsResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"external_id",
		"failed_to_import",
		"id",
		"name",
		"subscriptions",
		"updated_at",
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

	varBillingCustomerWithSubscriptionsResponseData := _BillingCustomerWithSubscriptionsResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingCustomerWithSubscriptionsResponseData)

	if err != nil {
		return err
	}

	*o = BillingCustomerWithSubscriptionsResponseData(varBillingCustomerWithSubscriptionsResponseData)

	return err
}

type NullableBillingCustomerWithSubscriptionsResponseData struct {
	value *BillingCustomerWithSubscriptionsResponseData
	isSet bool
}

func (v NullableBillingCustomerWithSubscriptionsResponseData) Get() *BillingCustomerWithSubscriptionsResponseData {
	return v.value
}

func (v *NullableBillingCustomerWithSubscriptionsResponseData) Set(val *BillingCustomerWithSubscriptionsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingCustomerWithSubscriptionsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingCustomerWithSubscriptionsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingCustomerWithSubscriptionsResponseData(val *BillingCustomerWithSubscriptionsResponseData) *NullableBillingCustomerWithSubscriptionsResponseData {
	return &NullableBillingCustomerWithSubscriptionsResponseData{value: val, isSet: true}
}

func (v NullableBillingCustomerWithSubscriptionsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingCustomerWithSubscriptionsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
