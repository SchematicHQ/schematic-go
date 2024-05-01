# ListFeatureCompaniesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FeatureCompanyResponseData**](FeatureCompanyResponseData.md) | The returned resources | 
**Params** | [**ListFeatureCompaniesParams**](ListFeatureCompaniesParams.md) |  | 

## Methods

### NewListFeatureCompaniesResponse

`func NewListFeatureCompaniesResponse(data []FeatureCompanyResponseData, params ListFeatureCompaniesParams, ) *ListFeatureCompaniesResponse`

NewListFeatureCompaniesResponse instantiates a new ListFeatureCompaniesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeatureCompaniesResponseWithDefaults

`func NewListFeatureCompaniesResponseWithDefaults() *ListFeatureCompaniesResponse`

NewListFeatureCompaniesResponseWithDefaults instantiates a new ListFeatureCompaniesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFeatureCompaniesResponse) GetData() []FeatureCompanyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFeatureCompaniesResponse) GetDataOk() (*[]FeatureCompanyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFeatureCompaniesResponse) SetData(v []FeatureCompanyResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListFeatureCompaniesResponse) GetParams() ListFeatureCompaniesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListFeatureCompaniesResponse) GetParamsOk() (*ListFeatureCompaniesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListFeatureCompaniesResponse) SetParams(v ListFeatureCompaniesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


