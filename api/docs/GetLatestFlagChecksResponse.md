# GetLatestFlagChecksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FlagCheckLogResponseData**](FlagCheckLogResponseData.md) | The returned resources | 
**Params** | [**GetLatestFlagChecksParams**](GetLatestFlagChecksParams.md) |  | 

## Methods

### NewGetLatestFlagChecksResponse

`func NewGetLatestFlagChecksResponse(data []FlagCheckLogResponseData, params GetLatestFlagChecksParams, ) *GetLatestFlagChecksResponse`

NewGetLatestFlagChecksResponse instantiates a new GetLatestFlagChecksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLatestFlagChecksResponseWithDefaults

`func NewGetLatestFlagChecksResponseWithDefaults() *GetLatestFlagChecksResponse`

NewGetLatestFlagChecksResponseWithDefaults instantiates a new GetLatestFlagChecksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetLatestFlagChecksResponse) GetData() []FlagCheckLogResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLatestFlagChecksResponse) GetDataOk() (*[]FlagCheckLogResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLatestFlagChecksResponse) SetData(v []FlagCheckLogResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetLatestFlagChecksResponse) GetParams() GetLatestFlagChecksParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetLatestFlagChecksResponse) GetParamsOk() (*GetLatestFlagChecksParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetLatestFlagChecksResponse) SetParams(v GetLatestFlagChecksParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


