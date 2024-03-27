# ListCompanyPlansResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanyPlanResponseData**](CompanyPlanResponseData.md) | The returned resources | 
**Params** | [**ListCompanyPlansParams**](ListCompanyPlansParams.md) |  | 

## Methods

### NewListCompanyPlansResponse

`func NewListCompanyPlansResponse(data []CompanyPlanResponseData, params ListCompanyPlansParams, ) *ListCompanyPlansResponse`

NewListCompanyPlansResponse instantiates a new ListCompanyPlansResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyPlansResponseWithDefaults

`func NewListCompanyPlansResponseWithDefaults() *ListCompanyPlansResponse`

NewListCompanyPlansResponseWithDefaults instantiates a new ListCompanyPlansResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCompanyPlansResponse) GetData() []CompanyPlanResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompanyPlansResponse) GetDataOk() (*[]CompanyPlanResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompanyPlansResponse) SetData(v []CompanyPlanResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCompanyPlansResponse) GetParams() ListCompanyPlansParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCompanyPlansResponse) GetParamsOk() (*ListCompanyPlansParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCompanyPlansResponse) SetParams(v ListCompanyPlansParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


