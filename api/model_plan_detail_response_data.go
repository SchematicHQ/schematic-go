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
	"time"
)

// checks if the PlanDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanDetailResponseData{}

// PlanDetailResponseData The updated resource
type PlanDetailResponseData struct {
	AudienceType         string                            `json:"audience_type"`
	BillingProduct       *BillingProductDetailResponseData `json:"billing_product,omitempty"`
	CompanyCount         int32                             `json:"company_count"`
	CreatedAt            time.Time                         `json:"created_at"`
	Description          string                            `json:"description"`
	Features             []FeatureDetailResponseData       `json:"features"`
	Icon                 string                            `json:"icon"`
	Id                   string                            `json:"id"`
	Name                 string                            `json:"name"`
	PlanType             string                            `json:"plan_type"`
	UpdatedAt            time.Time                         `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _PlanDetailResponseData PlanDetailResponseData

// NewPlanDetailResponseData instantiates a new PlanDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanDetailResponseData(audienceType string, companyCount int32, createdAt time.Time, description string, features []FeatureDetailResponseData, icon string, id string, name string, planType string, updatedAt time.Time) *PlanDetailResponseData {
	this := PlanDetailResponseData{}
	this.AudienceType = audienceType
	this.CompanyCount = companyCount
	this.CreatedAt = createdAt
	this.Description = description
	this.Features = features
	this.Icon = icon
	this.Id = id
	this.Name = name
	this.PlanType = planType
	this.UpdatedAt = updatedAt
	return &this
}

// NewPlanDetailResponseDataWithDefaults instantiates a new PlanDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanDetailResponseDataWithDefaults() *PlanDetailResponseData {
	this := PlanDetailResponseData{}
	return &this
}

// GetAudienceType returns the AudienceType field value
func (o *PlanDetailResponseData) GetAudienceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudienceType
}

// GetAudienceTypeOk returns a tuple with the AudienceType field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetAudienceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceType, true
}

// SetAudienceType sets field value
func (o *PlanDetailResponseData) SetAudienceType(v string) {
	o.AudienceType = v
}

// GetBillingProduct returns the BillingProduct field value if set, zero value otherwise.
func (o *PlanDetailResponseData) GetBillingProduct() BillingProductDetailResponseData {
	if o == nil || IsNil(o.BillingProduct) {
		var ret BillingProductDetailResponseData
		return ret
	}
	return *o.BillingProduct
}

// GetBillingProductOk returns a tuple with the BillingProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetBillingProductOk() (*BillingProductDetailResponseData, bool) {
	if o == nil || IsNil(o.BillingProduct) {
		return nil, false
	}
	return o.BillingProduct, true
}

// HasBillingProduct returns a boolean if a field has been set.
func (o *PlanDetailResponseData) HasBillingProduct() bool {
	if o != nil && !IsNil(o.BillingProduct) {
		return true
	}

	return false
}

// SetBillingProduct gets a reference to the given BillingProductDetailResponseData and assigns it to the BillingProduct field.
func (o *PlanDetailResponseData) SetBillingProduct(v BillingProductDetailResponseData) {
	o.BillingProduct = &v
}

// GetCompanyCount returns the CompanyCount field value
func (o *PlanDetailResponseData) GetCompanyCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompanyCount
}

// GetCompanyCountOk returns a tuple with the CompanyCount field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetCompanyCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyCount, true
}

// SetCompanyCount sets field value
func (o *PlanDetailResponseData) SetCompanyCount(v int32) {
	o.CompanyCount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PlanDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PlanDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *PlanDetailResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PlanDetailResponseData) SetDescription(v string) {
	o.Description = v
}

// GetFeatures returns the Features field value
func (o *PlanDetailResponseData) GetFeatures() []FeatureDetailResponseData {
	if o == nil {
		var ret []FeatureDetailResponseData
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetFeaturesOk() ([]FeatureDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *PlanDetailResponseData) SetFeatures(v []FeatureDetailResponseData) {
	o.Features = v
}

// GetIcon returns the Icon field value
func (o *PlanDetailResponseData) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *PlanDetailResponseData) SetIcon(v string) {
	o.Icon = v
}

// GetId returns the Id field value
func (o *PlanDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlanDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PlanDetailResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanDetailResponseData) SetName(v string) {
	o.Name = v
}

// GetPlanType returns the PlanType field value
func (o *PlanDetailResponseData) GetPlanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetPlanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanType, true
}

// SetPlanType sets field value
func (o *PlanDetailResponseData) SetPlanType(v string) {
	o.PlanType = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PlanDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PlanDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PlanDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PlanDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audience_type"] = o.AudienceType
	if !IsNil(o.BillingProduct) {
		toSerialize["billing_product"] = o.BillingProduct
	}
	toSerialize["company_count"] = o.CompanyCount
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["features"] = o.Features
	toSerialize["icon"] = o.Icon
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["plan_type"] = o.PlanType
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audience_type",
		"company_count",
		"created_at",
		"description",
		"features",
		"icon",
		"id",
		"name",
		"plan_type",
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

	varPlanDetailResponseData := _PlanDetailResponseData{}

	err = json.Unmarshal(data, &varPlanDetailResponseData)

	if err != nil {
		return err
	}

	*o = PlanDetailResponseData(varPlanDetailResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "audience_type")
		delete(additionalProperties, "billing_product")
		delete(additionalProperties, "company_count")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "description")
		delete(additionalProperties, "features")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "plan_type")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanDetailResponseData struct {
	value *PlanDetailResponseData
	isSet bool
}

func (v NullablePlanDetailResponseData) Get() *PlanDetailResponseData {
	return v.value
}

func (v *NullablePlanDetailResponseData) Set(val *PlanDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanDetailResponseData(val *PlanDetailResponseData) *NullablePlanDetailResponseData {
	return &NullablePlanDetailResponseData{value: val, isSet: true}
}

func (v NullablePlanDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
