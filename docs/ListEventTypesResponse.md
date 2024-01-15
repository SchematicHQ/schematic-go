# ListEventTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EventSummaryResponseData**](EventSummaryResponseData.md) | The returned resources | 
**Params** | [**ListEventTypesParams**](ListEventTypesParams.md) |  | 

## Methods

### NewListEventTypesResponse

`func NewListEventTypesResponse(data []EventSummaryResponseData, params ListEventTypesParams, ) *ListEventTypesResponse`

NewListEventTypesResponse instantiates a new ListEventTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventTypesResponseWithDefaults

`func NewListEventTypesResponseWithDefaults() *ListEventTypesResponse`

NewListEventTypesResponseWithDefaults instantiates a new ListEventTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEventTypesResponse) GetData() []EventSummaryResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEventTypesResponse) GetDataOk() (*[]EventSummaryResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEventTypesResponse) SetData(v []EventSummaryResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListEventTypesResponse) GetParams() ListEventTypesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListEventTypesResponse) GetParamsOk() (*ListEventTypesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListEventTypesResponse) SetParams(v ListEventTypesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


