# ListPlansResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PlanDetailResponseData**](PlanDetailResponseData.md) | The returned resources | 
**Params** | [**ListPlansParams**](ListPlansParams.md) |  | 

## Methods

### NewListPlansResponse

`func NewListPlansResponse(data []PlanDetailResponseData, params ListPlansParams, ) *ListPlansResponse`

NewListPlansResponse instantiates a new ListPlansResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlansResponseWithDefaults

`func NewListPlansResponseWithDefaults() *ListPlansResponse`

NewListPlansResponseWithDefaults instantiates a new ListPlansResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPlansResponse) GetData() []PlanDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPlansResponse) GetDataOk() (*[]PlanDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPlansResponse) SetData(v []PlanDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListPlansResponse) GetParams() ListPlansParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListPlansResponse) GetParamsOk() (*ListPlansParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListPlansResponse) SetParams(v ListPlansParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


