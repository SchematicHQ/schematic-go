# CountFlagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CountResponse**](CountResponse.md) |  | 
**Params** | [**CountFlagsParams**](CountFlagsParams.md) |  | 

## Methods

### NewCountFlagsResponse

`func NewCountFlagsResponse(data CountResponse, params CountFlagsParams, ) *CountFlagsResponse`

NewCountFlagsResponse instantiates a new CountFlagsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountFlagsResponseWithDefaults

`func NewCountFlagsResponseWithDefaults() *CountFlagsResponse`

NewCountFlagsResponseWithDefaults instantiates a new CountFlagsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CountFlagsResponse) GetData() CountResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CountFlagsResponse) GetDataOk() (*CountResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CountFlagsResponse) SetData(v CountResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *CountFlagsResponse) GetParams() CountFlagsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CountFlagsResponse) GetParamsOk() (*CountFlagsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CountFlagsResponse) SetParams(v CountFlagsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


