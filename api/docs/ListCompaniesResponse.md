# ListCompaniesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanyResponseData**](CompanyResponseData.md) | The returned resources | 
**Params** | [**ListCompaniesParams**](ListCompaniesParams.md) |  | 

## Methods

### NewListCompaniesResponse

`func NewListCompaniesResponse(data []CompanyResponseData, params ListCompaniesParams, ) *ListCompaniesResponse`

NewListCompaniesResponse instantiates a new ListCompaniesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompaniesResponseWithDefaults

`func NewListCompaniesResponseWithDefaults() *ListCompaniesResponse`

NewListCompaniesResponseWithDefaults instantiates a new ListCompaniesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCompaniesResponse) GetData() []CompanyResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompaniesResponse) GetDataOk() (*[]CompanyResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompaniesResponse) SetData(v []CompanyResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCompaniesResponse) GetParams() ListCompaniesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCompaniesResponse) GetParamsOk() (*ListCompaniesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCompaniesResponse) SetParams(v ListCompaniesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


