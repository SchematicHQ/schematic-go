# FlagDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**DefaultValue** | **bool** |  | 
**Description** | **string** |  | 
**Feature** | Pointer to [**FeatureResponseData**](FeatureResponseData.md) |  | [optional] 
**FeatureId** | Pointer to **NullableString** |  | [optional] 
**FlagType** | **string** |  | 
**Id** | **string** |  | 
**Key** | **string** |  | 
**LastCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**LatestCheck** | Pointer to [**FlagCheckLogResponseData**](FlagCheckLogResponseData.md) |  | [optional] 
**MaintainerId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Rules** | [**[]RuleDetailResponseData**](RuleDetailResponseData.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewFlagDetailResponseData

`func NewFlagDetailResponseData(createdAt time.Time, defaultValue bool, description string, flagType string, id string, key string, name string, rules []RuleDetailResponseData, updatedAt time.Time, ) *FlagDetailResponseData`

NewFlagDetailResponseData instantiates a new FlagDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagDetailResponseDataWithDefaults

`func NewFlagDetailResponseDataWithDefaults() *FlagDetailResponseData`

NewFlagDetailResponseDataWithDefaults instantiates a new FlagDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FlagDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlagDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlagDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultValue

`func (o *FlagDetailResponseData) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *FlagDetailResponseData) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *FlagDetailResponseData) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.


### GetDescription

`func (o *FlagDetailResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlagDetailResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlagDetailResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeature

`func (o *FlagDetailResponseData) GetFeature() FeatureResponseData`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *FlagDetailResponseData) GetFeatureOk() (*FeatureResponseData, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *FlagDetailResponseData) SetFeature(v FeatureResponseData)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *FlagDetailResponseData) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetFeatureId

`func (o *FlagDetailResponseData) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *FlagDetailResponseData) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *FlagDetailResponseData) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *FlagDetailResponseData) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### SetFeatureIdNil

`func (o *FlagDetailResponseData) SetFeatureIdNil(b bool)`

 SetFeatureIdNil sets the value for FeatureId to be an explicit nil

### UnsetFeatureId
`func (o *FlagDetailResponseData) UnsetFeatureId()`

UnsetFeatureId ensures that no value is present for FeatureId, not even an explicit nil
### GetFlagType

`func (o *FlagDetailResponseData) GetFlagType() string`

GetFlagType returns the FlagType field if non-nil, zero value otherwise.

### GetFlagTypeOk

`func (o *FlagDetailResponseData) GetFlagTypeOk() (*string, bool)`

GetFlagTypeOk returns a tuple with the FlagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagType

`func (o *FlagDetailResponseData) SetFlagType(v string)`

SetFlagType sets FlagType field to given value.


### GetId

`func (o *FlagDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *FlagDetailResponseData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FlagDetailResponseData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FlagDetailResponseData) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastCheckedAt

`func (o *FlagDetailResponseData) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *FlagDetailResponseData) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *FlagDetailResponseData) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *FlagDetailResponseData) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *FlagDetailResponseData) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *FlagDetailResponseData) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetLatestCheck

`func (o *FlagDetailResponseData) GetLatestCheck() FlagCheckLogResponseData`

GetLatestCheck returns the LatestCheck field if non-nil, zero value otherwise.

### GetLatestCheckOk

`func (o *FlagDetailResponseData) GetLatestCheckOk() (*FlagCheckLogResponseData, bool)`

GetLatestCheckOk returns a tuple with the LatestCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCheck

`func (o *FlagDetailResponseData) SetLatestCheck(v FlagCheckLogResponseData)`

SetLatestCheck sets LatestCheck field to given value.

### HasLatestCheck

`func (o *FlagDetailResponseData) HasLatestCheck() bool`

HasLatestCheck returns a boolean if a field has been set.

### GetMaintainerId

`func (o *FlagDetailResponseData) GetMaintainerId() string`

GetMaintainerId returns the MaintainerId field if non-nil, zero value otherwise.

### GetMaintainerIdOk

`func (o *FlagDetailResponseData) GetMaintainerIdOk() (*string, bool)`

GetMaintainerIdOk returns a tuple with the MaintainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerId

`func (o *FlagDetailResponseData) SetMaintainerId(v string)`

SetMaintainerId sets MaintainerId field to given value.

### HasMaintainerId

`func (o *FlagDetailResponseData) HasMaintainerId() bool`

HasMaintainerId returns a boolean if a field has been set.

### SetMaintainerIdNil

`func (o *FlagDetailResponseData) SetMaintainerIdNil(b bool)`

 SetMaintainerIdNil sets the value for MaintainerId to be an explicit nil

### UnsetMaintainerId
`func (o *FlagDetailResponseData) UnsetMaintainerId()`

UnsetMaintainerId ensures that no value is present for MaintainerId, not even an explicit nil
### GetName

`func (o *FlagDetailResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlagDetailResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlagDetailResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *FlagDetailResponseData) GetRules() []RuleDetailResponseData`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *FlagDetailResponseData) GetRulesOk() (*[]RuleDetailResponseData, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *FlagDetailResponseData) SetRules(v []RuleDetailResponseData)`

SetRules sets Rules field to given value.


### GetUpdatedAt

`func (o *FlagDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlagDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlagDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


