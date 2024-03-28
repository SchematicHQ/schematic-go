# EventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**EventBodyIdentifyCompany**](EventBodyIdentifyCompany.md) |  | [optional] 
**Event** | **string** | The name of the type of track event | 
**Traits** | Pointer to **map[string]interface{}** | A map of user trait names to trait values | [optional] 
**User** | Pointer to **map[string]interface{}** | Key-value pairs to identify user associated with track event | [optional] 
**Keys** | **map[string]interface{}** | Key-value pairs to identify the user | 
**Name** | Pointer to **string** | The display name of the user being identified; required only if it is a new user | [optional] 

## Methods

### NewEventBody

`func NewEventBody(event string, keys map[string]interface{}, ) *EventBody`

NewEventBody instantiates a new EventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyWithDefaults

`func NewEventBodyWithDefaults() *EventBody`

NewEventBodyWithDefaults instantiates a new EventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *EventBody) GetCompany() EventBodyIdentifyCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EventBody) GetCompanyOk() (*EventBodyIdentifyCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EventBody) SetCompany(v EventBodyIdentifyCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EventBody) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEvent

`func (o *EventBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventBody) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTraits

`func (o *EventBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *EventBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *EventBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *EventBody) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUser

`func (o *EventBody) GetUser() map[string]interface{}`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventBody) GetUserOk() (*map[string]interface{}, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventBody) SetUser(v map[string]interface{})`

SetUser sets User field to given value.

### HasUser

`func (o *EventBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetKeys

`func (o *EventBody) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *EventBody) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *EventBody) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetName

`func (o *EventBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


