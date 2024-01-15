# EventBodyIdentifyCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | **map[string]interface{}** | Key-value pairs to identify the company | 
**Name** | Pointer to **string** | The display name of the company; required only if it is a new company | [optional] 
**Traits** | Pointer to **map[string]interface{}** | A map of company trait names to trait values | [optional] 

## Methods

### NewEventBodyIdentifyCompany

`func NewEventBodyIdentifyCompany(keys map[string]interface{}, ) *EventBodyIdentifyCompany`

NewEventBodyIdentifyCompany instantiates a new EventBodyIdentifyCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyIdentifyCompanyWithDefaults

`func NewEventBodyIdentifyCompanyWithDefaults() *EventBodyIdentifyCompany`

NewEventBodyIdentifyCompanyWithDefaults instantiates a new EventBodyIdentifyCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *EventBodyIdentifyCompany) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *EventBodyIdentifyCompany) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *EventBodyIdentifyCompany) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetName

`func (o *EventBodyIdentifyCompany) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventBodyIdentifyCompany) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventBodyIdentifyCompany) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventBodyIdentifyCompany) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTraits

`func (o *EventBodyIdentifyCompany) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *EventBodyIdentifyCompany) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *EventBodyIdentifyCompany) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *EventBodyIdentifyCompany) HasTraits() bool`

HasTraits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


