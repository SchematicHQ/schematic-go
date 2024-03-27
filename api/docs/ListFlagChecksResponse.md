# ListFlagChecksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FlagCheckLogDetailResponseData**](FlagCheckLogDetailResponseData.md) | The returned resources | 
**Params** | [**ListFlagChecksParams**](ListFlagChecksParams.md) |  | 

## Methods

### NewListFlagChecksResponse

`func NewListFlagChecksResponse(data []FlagCheckLogDetailResponseData, params ListFlagChecksParams, ) *ListFlagChecksResponse`

NewListFlagChecksResponse instantiates a new ListFlagChecksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlagChecksResponseWithDefaults

`func NewListFlagChecksResponseWithDefaults() *ListFlagChecksResponse`

NewListFlagChecksResponseWithDefaults instantiates a new ListFlagChecksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFlagChecksResponse) GetData() []FlagCheckLogDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFlagChecksResponse) GetDataOk() (*[]FlagCheckLogDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFlagChecksResponse) SetData(v []FlagCheckLogDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListFlagChecksResponse) GetParams() ListFlagChecksParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListFlagChecksResponse) GetParamsOk() (*ListFlagChecksParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListFlagChecksResponse) SetParams(v ListFlagChecksParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


