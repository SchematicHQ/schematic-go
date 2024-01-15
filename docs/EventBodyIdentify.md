# EventBodyIdentify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**EventBodyIdentifyCompany**](EventBodyIdentifyCompany.md) |  | [optional] 
**Keys** | **map[string]interface{}** | Key-value pairs to identify the user | 
**Name** | Pointer to **string** | The display name of the user being identified; required only if it is a new user | [optional] 
**Traits** | Pointer to **map[string]interface{}** | A map of user trait names to trait values | [optional] 

## Methods

### NewEventBodyIdentify

`func NewEventBodyIdentify(keys map[string]interface{}, ) *EventBodyIdentify`

NewEventBodyIdentify instantiates a new EventBodyIdentify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyIdentifyWithDefaults

`func NewEventBodyIdentifyWithDefaults() *EventBodyIdentify`

NewEventBodyIdentifyWithDefaults instantiates a new EventBodyIdentify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *EventBodyIdentify) GetCompany() EventBodyIdentifyCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EventBodyIdentify) GetCompanyOk() (*EventBodyIdentifyCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EventBodyIdentify) SetCompany(v EventBodyIdentifyCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EventBodyIdentify) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetKeys

`func (o *EventBodyIdentify) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *EventBodyIdentify) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *EventBodyIdentify) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetName

`func (o *EventBodyIdentify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventBodyIdentify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventBodyIdentify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventBodyIdentify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTraits

`func (o *EventBodyIdentify) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *EventBodyIdentify) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *EventBodyIdentify) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *EventBodyIdentify) HasTraits() bool`

HasTraits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


