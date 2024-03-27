# RawEventResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapturedAt** | **time.Time** |  | 
**EventId** | Pointer to **NullableString** |  | [optional] 
**RemoteAddr** | **string** |  | 
**RemoteIp** | **string** |  | 
**UserAgent** | **string** |  | 

## Methods

### NewRawEventResponseData

`func NewRawEventResponseData(capturedAt time.Time, remoteAddr string, remoteIp string, userAgent string, ) *RawEventResponseData`

NewRawEventResponseData instantiates a new RawEventResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawEventResponseDataWithDefaults

`func NewRawEventResponseDataWithDefaults() *RawEventResponseData`

NewRawEventResponseDataWithDefaults instantiates a new RawEventResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapturedAt

`func (o *RawEventResponseData) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *RawEventResponseData) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *RawEventResponseData) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### GetEventId

`func (o *RawEventResponseData) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *RawEventResponseData) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *RawEventResponseData) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *RawEventResponseData) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventIdNil

`func (o *RawEventResponseData) SetEventIdNil(b bool)`

 SetEventIdNil sets the value for EventId to be an explicit nil

### UnsetEventId
`func (o *RawEventResponseData) UnsetEventId()`

UnsetEventId ensures that no value is present for EventId, not even an explicit nil
### GetRemoteAddr

`func (o *RawEventResponseData) GetRemoteAddr() string`

GetRemoteAddr returns the RemoteAddr field if non-nil, zero value otherwise.

### GetRemoteAddrOk

`func (o *RawEventResponseData) GetRemoteAddrOk() (*string, bool)`

GetRemoteAddrOk returns a tuple with the RemoteAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddr

`func (o *RawEventResponseData) SetRemoteAddr(v string)`

SetRemoteAddr sets RemoteAddr field to given value.


### GetRemoteIp

`func (o *RawEventResponseData) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *RawEventResponseData) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *RawEventResponseData) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetUserAgent

`func (o *RawEventResponseData) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RawEventResponseData) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RawEventResponseData) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


