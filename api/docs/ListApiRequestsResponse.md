# ListApiRequestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ApiKeyRequestListResponseData**](ApiKeyRequestListResponseData.md) | The returned resources | 
**Params** | [**ListApiRequestsParams**](ListApiRequestsParams.md) |  | 

## Methods

### NewListApiRequestsResponse

`func NewListApiRequestsResponse(data []ApiKeyRequestListResponseData, params ListApiRequestsParams, ) *ListApiRequestsResponse`

NewListApiRequestsResponse instantiates a new ListApiRequestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApiRequestsResponseWithDefaults

`func NewListApiRequestsResponseWithDefaults() *ListApiRequestsResponse`

NewListApiRequestsResponseWithDefaults instantiates a new ListApiRequestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListApiRequestsResponse) GetData() []ApiKeyRequestListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListApiRequestsResponse) GetDataOk() (*[]ApiKeyRequestListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListApiRequestsResponse) SetData(v []ApiKeyRequestListResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListApiRequestsResponse) GetParams() ListApiRequestsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListApiRequestsResponse) GetParamsOk() (*ListApiRequestsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListApiRequestsResponse) SetParams(v ListApiRequestsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


