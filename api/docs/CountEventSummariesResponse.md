# CountEventSummariesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CountResponse**](CountResponse.md) |  | 
**Params** | [**CountEventSummariesParams**](CountEventSummariesParams.md) |  | 

## Methods

### NewCountEventSummariesResponse

`func NewCountEventSummariesResponse(data CountResponse, params CountEventSummariesParams, ) *CountEventSummariesResponse`

NewCountEventSummariesResponse instantiates a new CountEventSummariesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountEventSummariesResponseWithDefaults

`func NewCountEventSummariesResponseWithDefaults() *CountEventSummariesResponse`

NewCountEventSummariesResponseWithDefaults instantiates a new CountEventSummariesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CountEventSummariesResponse) GetData() CountResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CountEventSummariesResponse) GetDataOk() (*CountResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CountEventSummariesResponse) SetData(v CountResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *CountEventSummariesResponse) GetParams() CountEventSummariesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CountEventSummariesResponse) GetParamsOk() (*CountEventSummariesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CountEventSummariesResponse) SetParams(v CountEventSummariesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


