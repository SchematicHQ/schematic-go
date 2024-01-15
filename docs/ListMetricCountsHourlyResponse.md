# ListMetricCountsHourlyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MetricCountsHourlyResponseData**](MetricCountsHourlyResponseData.md) | The returned resources | 
**Params** | [**ListMetricCountsHourlyParams**](ListMetricCountsHourlyParams.md) |  | 

## Methods

### NewListMetricCountsHourlyResponse

`func NewListMetricCountsHourlyResponse(data []MetricCountsHourlyResponseData, params ListMetricCountsHourlyParams, ) *ListMetricCountsHourlyResponse`

NewListMetricCountsHourlyResponse instantiates a new ListMetricCountsHourlyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricCountsHourlyResponseWithDefaults

`func NewListMetricCountsHourlyResponseWithDefaults() *ListMetricCountsHourlyResponse`

NewListMetricCountsHourlyResponseWithDefaults instantiates a new ListMetricCountsHourlyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListMetricCountsHourlyResponse) GetData() []MetricCountsHourlyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListMetricCountsHourlyResponse) GetDataOk() (*[]MetricCountsHourlyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListMetricCountsHourlyResponse) SetData(v []MetricCountsHourlyResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListMetricCountsHourlyResponse) GetParams() ListMetricCountsHourlyParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListMetricCountsHourlyResponse) GetParamsOk() (*ListMetricCountsHourlyParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListMetricCountsHourlyResponse) SetParams(v ListMetricCountsHourlyParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


