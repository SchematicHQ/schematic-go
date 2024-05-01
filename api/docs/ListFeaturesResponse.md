# ListFeaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FeatureDetailResponseData**](FeatureDetailResponseData.md) | The returned resources | 
**Params** | [**ListFeaturesParams**](ListFeaturesParams.md) |  | 

## Methods

### NewListFeaturesResponse

`func NewListFeaturesResponse(data []FeatureDetailResponseData, params ListFeaturesParams, ) *ListFeaturesResponse`

NewListFeaturesResponse instantiates a new ListFeaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeaturesResponseWithDefaults

`func NewListFeaturesResponseWithDefaults() *ListFeaturesResponse`

NewListFeaturesResponseWithDefaults instantiates a new ListFeaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFeaturesResponse) GetData() []FeatureDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFeaturesResponse) GetDataOk() (*[]FeatureDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFeaturesResponse) SetData(v []FeatureDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListFeaturesResponse) GetParams() ListFeaturesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListFeaturesResponse) GetParamsOk() (*ListFeaturesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListFeaturesResponse) SetParams(v ListFeaturesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


