# ListEventsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**EventSubtype** | Pointer to **string** |  | [optional] 
**EventTypes** | Pointer to **[]string** |  | [optional] 
**FlagId** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewListEventsParams

`func NewListEventsParams() *ListEventsParams`

NewListEventsParams instantiates a new ListEventsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsParamsWithDefaults

`func NewListEventsParamsWithDefaults() *ListEventsParams`

NewListEventsParamsWithDefaults instantiates a new ListEventsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListEventsParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListEventsParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListEventsParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListEventsParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetEventSubtype

`func (o *ListEventsParams) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *ListEventsParams) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *ListEventsParams) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *ListEventsParams) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### GetEventTypes

`func (o *ListEventsParams) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *ListEventsParams) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *ListEventsParams) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *ListEventsParams) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetFlagId

`func (o *ListEventsParams) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *ListEventsParams) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *ListEventsParams) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *ListEventsParams) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### GetLimit

`func (o *ListEventsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListEventsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListEventsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListEventsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListEventsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListEventsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListEventsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListEventsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetUserId

`func (o *ListEventsParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListEventsParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListEventsParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListEventsParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


