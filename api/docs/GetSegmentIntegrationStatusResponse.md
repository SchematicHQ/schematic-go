# GetSegmentIntegrationStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**SegmentStatusResp**](SegmentStatusResp.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewGetSegmentIntegrationStatusResponse

`func NewGetSegmentIntegrationStatusResponse(data SegmentStatusResp, params map[string]interface{}, ) *GetSegmentIntegrationStatusResponse`

NewGetSegmentIntegrationStatusResponse instantiates a new GetSegmentIntegrationStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentIntegrationStatusResponseWithDefaults

`func NewGetSegmentIntegrationStatusResponseWithDefaults() *GetSegmentIntegrationStatusResponse`

NewGetSegmentIntegrationStatusResponseWithDefaults instantiates a new GetSegmentIntegrationStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSegmentIntegrationStatusResponse) GetData() SegmentStatusResp`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSegmentIntegrationStatusResponse) GetDataOk() (*SegmentStatusResp, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSegmentIntegrationStatusResponse) SetData(v SegmentStatusResp)`

SetData sets Data field to given value.


### GetParams

`func (o *GetSegmentIntegrationStatusResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetSegmentIntegrationStatusResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetSegmentIntegrationStatusResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


