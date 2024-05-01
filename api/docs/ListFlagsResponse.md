# ListFlagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FlagDetailResponseData**](FlagDetailResponseData.md) | The returned resources | 
**Params** | [**ListFlagsParams**](ListFlagsParams.md) |  | 

## Methods

### NewListFlagsResponse

`func NewListFlagsResponse(data []FlagDetailResponseData, params ListFlagsParams, ) *ListFlagsResponse`

NewListFlagsResponse instantiates a new ListFlagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlagsResponseWithDefaults

`func NewListFlagsResponseWithDefaults() *ListFlagsResponse`

NewListFlagsResponseWithDefaults instantiates a new ListFlagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFlagsResponse) GetData() []FlagDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFlagsResponse) GetDataOk() (*[]FlagDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFlagsResponse) SetData(v []FlagDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListFlagsResponse) GetParams() ListFlagsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListFlagsResponse) GetParamsOk() (*ListFlagsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListFlagsResponse) SetParams(v ListFlagsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


