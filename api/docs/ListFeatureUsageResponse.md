# ListFeatureUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FeatureUsageResponseData**](FeatureUsageResponseData.md) | The returned resources | 
**Params** | [**ListFeatureUsageParams**](ListFeatureUsageParams.md) |  | 

## Methods

### NewListFeatureUsageResponse

`func NewListFeatureUsageResponse(data []FeatureUsageResponseData, params ListFeatureUsageParams, ) *ListFeatureUsageResponse`

NewListFeatureUsageResponse instantiates a new ListFeatureUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeatureUsageResponseWithDefaults

`func NewListFeatureUsageResponseWithDefaults() *ListFeatureUsageResponse`

NewListFeatureUsageResponseWithDefaults instantiates a new ListFeatureUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFeatureUsageResponse) GetData() []FeatureUsageResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFeatureUsageResponse) GetDataOk() (*[]FeatureUsageResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFeatureUsageResponse) SetData(v []FeatureUsageResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListFeatureUsageResponse) GetParams() ListFeatureUsageParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListFeatureUsageResponse) GetParamsOk() (*ListFeatureUsageParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListFeatureUsageResponse) SetParams(v ListFeatureUsageParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


