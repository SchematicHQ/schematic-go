# ListMetricCountsHourlyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyIds** | Pointer to **[]string** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**EventSubtype** | Pointer to **string** |  | [optional] 
**EventSubtypes** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewListMetricCountsHourlyParams

`func NewListMetricCountsHourlyParams() *ListMetricCountsHourlyParams`

NewListMetricCountsHourlyParams instantiates a new ListMetricCountsHourlyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricCountsHourlyParamsWithDefaults

`func NewListMetricCountsHourlyParamsWithDefaults() *ListMetricCountsHourlyParams`

NewListMetricCountsHourlyParamsWithDefaults instantiates a new ListMetricCountsHourlyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListMetricCountsHourlyParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListMetricCountsHourlyParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListMetricCountsHourlyParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListMetricCountsHourlyParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyIds

`func (o *ListMetricCountsHourlyParams) GetCompanyIds() []string`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ListMetricCountsHourlyParams) GetCompanyIdsOk() (*[]string, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ListMetricCountsHourlyParams) SetCompanyIds(v []string)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *ListMetricCountsHourlyParams) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### GetEndTime

`func (o *ListMetricCountsHourlyParams) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ListMetricCountsHourlyParams) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ListMetricCountsHourlyParams) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ListMetricCountsHourlyParams) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEventSubtype

`func (o *ListMetricCountsHourlyParams) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *ListMetricCountsHourlyParams) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *ListMetricCountsHourlyParams) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *ListMetricCountsHourlyParams) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### GetEventSubtypes

`func (o *ListMetricCountsHourlyParams) GetEventSubtypes() []string`

GetEventSubtypes returns the EventSubtypes field if non-nil, zero value otherwise.

### GetEventSubtypesOk

`func (o *ListMetricCountsHourlyParams) GetEventSubtypesOk() (*[]string, bool)`

GetEventSubtypesOk returns a tuple with the EventSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtypes

`func (o *ListMetricCountsHourlyParams) SetEventSubtypes(v []string)`

SetEventSubtypes sets EventSubtypes field to given value.

### HasEventSubtypes

`func (o *ListMetricCountsHourlyParams) HasEventSubtypes() bool`

HasEventSubtypes returns a boolean if a field has been set.

### GetLimit

`func (o *ListMetricCountsHourlyParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListMetricCountsHourlyParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListMetricCountsHourlyParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListMetricCountsHourlyParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListMetricCountsHourlyParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListMetricCountsHourlyParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListMetricCountsHourlyParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListMetricCountsHourlyParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStartTime

`func (o *ListMetricCountsHourlyParams) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ListMetricCountsHourlyParams) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ListMetricCountsHourlyParams) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ListMetricCountsHourlyParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUserId

`func (o *ListMetricCountsHourlyParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListMetricCountsHourlyParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListMetricCountsHourlyParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListMetricCountsHourlyParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


