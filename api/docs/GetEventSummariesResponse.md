# GetEventSummariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EventSummaryResponseData**](EventSummaryResponseData.md) | The returned resources | 
**Params** | [**GetEventSummariesParams**](GetEventSummariesParams.md) |  | 

## Methods

### NewGetEventSummariesResponse

`func NewGetEventSummariesResponse(data []EventSummaryResponseData, params GetEventSummariesParams, ) *GetEventSummariesResponse`

NewGetEventSummariesResponse instantiates a new GetEventSummariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventSummariesResponseWithDefaults

`func NewGetEventSummariesResponseWithDefaults() *GetEventSummariesResponse`

NewGetEventSummariesResponseWithDefaults instantiates a new GetEventSummariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEventSummariesResponse) GetData() []EventSummaryResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventSummariesResponse) GetDataOk() (*[]EventSummaryResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventSummariesResponse) SetData(v []EventSummaryResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetEventSummariesResponse) GetParams() GetEventSummariesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetEventSummariesResponse) GetParamsOk() (*GetEventSummariesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetEventSummariesResponse) SetParams(v GetEventSummariesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


