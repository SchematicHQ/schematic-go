# ListComponentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ComponentResponseData**](ComponentResponseData.md) | The returned resources | 
**Params** | [**ListComponentsParams**](ListComponentsParams.md) |  | 

## Methods

### NewListComponentsResponse

`func NewListComponentsResponse(data []ComponentResponseData, params ListComponentsParams, ) *ListComponentsResponse`

NewListComponentsResponse instantiates a new ListComponentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListComponentsResponseWithDefaults

`func NewListComponentsResponseWithDefaults() *ListComponentsResponse`

NewListComponentsResponseWithDefaults instantiates a new ListComponentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListComponentsResponse) GetData() []ComponentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListComponentsResponse) GetDataOk() (*[]ComponentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListComponentsResponse) SetData(v []ComponentResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListComponentsResponse) GetParams() ListComponentsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListComponentsResponse) GetParamsOk() (*ListComponentsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListComponentsResponse) SetParams(v ListComponentsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


