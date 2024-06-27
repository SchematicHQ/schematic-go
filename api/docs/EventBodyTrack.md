# EventBodyTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **map[string]string** | Key-value pairs to identify company associated with track event | [optional] 
**Event** | **string** | The name of the type of track event | 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**User** | Pointer to **map[string]string** | Key-value pairs to identify user associated with track event | [optional] 

## Methods

### NewEventBodyTrack

`func NewEventBodyTrack(event string, ) *EventBodyTrack`

NewEventBodyTrack instantiates a new EventBodyTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyTrackWithDefaults

`func NewEventBodyTrackWithDefaults() *EventBodyTrack`

NewEventBodyTrackWithDefaults instantiates a new EventBodyTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *EventBodyTrack) GetCompany() map[string]string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EventBodyTrack) GetCompanyOk() (*map[string]string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EventBodyTrack) SetCompany(v map[string]string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EventBodyTrack) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEvent

`func (o *EventBodyTrack) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventBodyTrack) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventBodyTrack) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTraits

`func (o *EventBodyTrack) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *EventBodyTrack) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *EventBodyTrack) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *EventBodyTrack) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUser

`func (o *EventBodyTrack) GetUser() map[string]string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventBodyTrack) GetUserOk() (*map[string]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventBodyTrack) SetUser(v map[string]string)`

SetUser sets User field to given value.

### HasUser

`func (o *EventBodyTrack) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


