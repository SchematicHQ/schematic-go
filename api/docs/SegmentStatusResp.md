# SegmentStatusResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** |  | 
**EnvironmentId** | **string** |  | 
**LastEventAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSegmentStatusResp

`func NewSegmentStatusResp(connected bool, environmentId string, ) *SegmentStatusResp`

NewSegmentStatusResp instantiates a new SegmentStatusResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentStatusRespWithDefaults

`func NewSegmentStatusRespWithDefaults() *SegmentStatusResp`

NewSegmentStatusRespWithDefaults instantiates a new SegmentStatusResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *SegmentStatusResp) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *SegmentStatusResp) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *SegmentStatusResp) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetEnvironmentId

`func (o *SegmentStatusResp) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SegmentStatusResp) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SegmentStatusResp) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetLastEventAt

`func (o *SegmentStatusResp) GetLastEventAt() time.Time`

GetLastEventAt returns the LastEventAt field if non-nil, zero value otherwise.

### GetLastEventAtOk

`func (o *SegmentStatusResp) GetLastEventAtOk() (*time.Time, bool)`

GetLastEventAtOk returns a tuple with the LastEventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventAt

`func (o *SegmentStatusResp) SetLastEventAt(v time.Time)`

SetLastEventAt sets LastEventAt field to given value.

### HasLastEventAt

`func (o *SegmentStatusResp) HasLastEventAt() bool`

HasLastEventAt returns a boolean if a field has been set.

### SetLastEventAtNil

`func (o *SegmentStatusResp) SetLastEventAtNil(b bool)`

 SetLastEventAtNil sets the value for LastEventAt to be an explicit nil

### UnsetLastEventAt
`func (o *SegmentStatusResp) UnsetLastEventAt()`

UnsetLastEventAt ensures that no value is present for LastEventAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


