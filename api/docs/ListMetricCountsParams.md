# ListMetricCountsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**CompanyIds** | Pointer to **[]string** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**EventSubtype** | Pointer to **string** |  | [optional] 
**EventSubtypes** | Pointer to **[]string** |  | [optional] 
**Grouping** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewListMetricCountsParams

`func NewListMetricCountsParams() *ListMetricCountsParams`

NewListMetricCountsParams instantiates a new ListMetricCountsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricCountsParamsWithDefaults

`func NewListMetricCountsParamsWithDefaults() *ListMetricCountsParams`

NewListMetricCountsParamsWithDefaults instantiates a new ListMetricCountsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListMetricCountsParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListMetricCountsParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListMetricCountsParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListMetricCountsParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCompanyIds

`func (o *ListMetricCountsParams) GetCompanyIds() []string`

GetCompanyIds returns the CompanyIds field if non-nil, zero value otherwise.

### GetCompanyIdsOk

`func (o *ListMetricCountsParams) GetCompanyIdsOk() (*[]string, bool)`

GetCompanyIdsOk returns a tuple with the CompanyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIds

`func (o *ListMetricCountsParams) SetCompanyIds(v []string)`

SetCompanyIds sets CompanyIds field to given value.

### HasCompanyIds

`func (o *ListMetricCountsParams) HasCompanyIds() bool`

HasCompanyIds returns a boolean if a field has been set.

### GetEndTime

`func (o *ListMetricCountsParams) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ListMetricCountsParams) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ListMetricCountsParams) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ListMetricCountsParams) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEventSubtype

`func (o *ListMetricCountsParams) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *ListMetricCountsParams) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *ListMetricCountsParams) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.

### HasEventSubtype

`func (o *ListMetricCountsParams) HasEventSubtype() bool`

HasEventSubtype returns a boolean if a field has been set.

### GetEventSubtypes

`func (o *ListMetricCountsParams) GetEventSubtypes() []string`

GetEventSubtypes returns the EventSubtypes field if non-nil, zero value otherwise.

### GetEventSubtypesOk

`func (o *ListMetricCountsParams) GetEventSubtypesOk() (*[]string, bool)`

GetEventSubtypesOk returns a tuple with the EventSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtypes

`func (o *ListMetricCountsParams) SetEventSubtypes(v []string)`

SetEventSubtypes sets EventSubtypes field to given value.

### HasEventSubtypes

`func (o *ListMetricCountsParams) HasEventSubtypes() bool`

HasEventSubtypes returns a boolean if a field has been set.

### GetGrouping

`func (o *ListMetricCountsParams) GetGrouping() string`

GetGrouping returns the Grouping field if non-nil, zero value otherwise.

### GetGroupingOk

`func (o *ListMetricCountsParams) GetGroupingOk() (*string, bool)`

GetGroupingOk returns a tuple with the Grouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrouping

`func (o *ListMetricCountsParams) SetGrouping(v string)`

SetGrouping sets Grouping field to given value.

### HasGrouping

`func (o *ListMetricCountsParams) HasGrouping() bool`

HasGrouping returns a boolean if a field has been set.

### GetLimit

`func (o *ListMetricCountsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListMetricCountsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListMetricCountsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListMetricCountsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListMetricCountsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListMetricCountsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListMetricCountsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListMetricCountsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStartTime

`func (o *ListMetricCountsParams) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ListMetricCountsParams) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ListMetricCountsParams) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ListMetricCountsParams) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUserId

`func (o *ListMetricCountsParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListMetricCountsParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListMetricCountsParams) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListMetricCountsParams) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


