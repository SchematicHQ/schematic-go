# MetricCountsHourlyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**EventSubtype** | **string** |  | 
**StartTime** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Value** | **int32** |  | 

## Methods

### NewMetricCountsHourlyResponseData

`func NewMetricCountsHourlyResponseData(createdAt time.Time, environmentId string, eventSubtype string, startTime time.Time, value int32, ) *MetricCountsHourlyResponseData`

NewMetricCountsHourlyResponseData instantiates a new MetricCountsHourlyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCountsHourlyResponseDataWithDefaults

`func NewMetricCountsHourlyResponseDataWithDefaults() *MetricCountsHourlyResponseData`

NewMetricCountsHourlyResponseDataWithDefaults instantiates a new MetricCountsHourlyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *MetricCountsHourlyResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *MetricCountsHourlyResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *MetricCountsHourlyResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *MetricCountsHourlyResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *MetricCountsHourlyResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *MetricCountsHourlyResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *MetricCountsHourlyResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricCountsHourlyResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricCountsHourlyResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *MetricCountsHourlyResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *MetricCountsHourlyResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *MetricCountsHourlyResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetEventSubtype

`func (o *MetricCountsHourlyResponseData) GetEventSubtype() string`

GetEventSubtype returns the EventSubtype field if non-nil, zero value otherwise.

### GetEventSubtypeOk

`func (o *MetricCountsHourlyResponseData) GetEventSubtypeOk() (*string, bool)`

GetEventSubtypeOk returns a tuple with the EventSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtype

`func (o *MetricCountsHourlyResponseData) SetEventSubtype(v string)`

SetEventSubtype sets EventSubtype field to given value.


### GetStartTime

`func (o *MetricCountsHourlyResponseData) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MetricCountsHourlyResponseData) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MetricCountsHourlyResponseData) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetUserId

`func (o *MetricCountsHourlyResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MetricCountsHourlyResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MetricCountsHourlyResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MetricCountsHourlyResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MetricCountsHourlyResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MetricCountsHourlyResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *MetricCountsHourlyResponseData) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricCountsHourlyResponseData) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricCountsHourlyResponseData) SetValue(v int32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


