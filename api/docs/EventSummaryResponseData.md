# EventSummaryResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyCount** | **int32** |  | 
**EnvironmentId** | **string** |  | 
**EventCount** | **int32** |  | 
**EventSubtype** | **string** |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**UserCount** | **int32** |  | 

## Methods

### NewEventSummaryResponseData

`func NewEventSummaryResponseData(companyCount int32, environmentId string, eventCount int32, eventSubtype string, userCount int32, ) *EventSummaryResponseData`

NewEventSummaryResponseData instantiates a new EventSummaryResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSummaryResponseDataWithDefaults

`func NewEventSummaryResponseDataWithDefaults() *EventSummaryResponseData`

NewEventSummaryResponseDataWithDefaults instantiates a new EventSummaryResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyCount

`func (o *EventSummaryResponseData) GetCompanyCount() int32`

GetCompanyCount returns the CompanyCount field if non-nil, zero value otherwise.

### GetCompanyCountOk

`func (o *EventSummaryResponseData) GetCompanyCountOk() (*int32, bool)`

GetCompanyCountOk returns a tuple with the CompanyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyCount

`func (o *EventSummaryResponseData) SetCompanyCount(v int32)`

SetCompanyCount sets CompanyCount field to given value.


### GetEnvironmentId

`func (o *EventSummaryResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EventSummaryResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EventSummaryResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventCount

`func (o *EventSummaryResponseData) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *EventSummaryResponseData) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *EventSummaryResponseData) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetEventSubtype

`func (o *EventSummaryResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *EventSummaryResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *EventSummaryResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.


### GetLastSeenAt

`func (o *EventSummaryResponseData) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *EventSummaryResponseData) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *EventSummaryResponseData) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *EventSummaryResponseData) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *EventSummaryResponseData) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *EventSummaryResponseData) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetUserCount

`func (o *EventSummaryResponseData) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *EventSummaryResponseData) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *EventSummaryResponseData) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


