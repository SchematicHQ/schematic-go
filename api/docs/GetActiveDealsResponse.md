# GetActiveDealsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanyCrmDealsResponseData**](CompanyCrmDealsResponseData.md) | The returned resources | 
**Params** | [**GetActiveDealsParams**](GetActiveDealsParams.md) |  | 

## Methods

### NewGetActiveDealsResponse

`func NewGetActiveDealsResponse(data []CompanyCrmDealsResponseData, params GetActiveDealsParams, ) *GetActiveDealsResponse`

NewGetActiveDealsResponse instantiates a new GetActiveDealsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActiveDealsResponseWithDefaults

`func NewGetActiveDealsResponseWithDefaults() *GetActiveDealsResponse`

NewGetActiveDealsResponseWithDefaults instantiates a new GetActiveDealsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActiveDealsResponse) GetData() []CompanyCrmDealsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActiveDealsResponse) GetDataOk() (*[]CompanyCrmDealsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActiveDealsResponse) SetData(v []CompanyCrmDealsResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetActiveDealsResponse) GetParams() GetActiveDealsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetActiveDealsResponse) GetParamsOk() (*GetActiveDealsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetActiveDealsResponse) SetParams(v GetActiveDealsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


