# ListCompanyOverridesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanyOverrideResponseData**](CompanyOverrideResponseData.md) | The returned resources | 
**Params** | [**ListCompanyOverridesParams**](ListCompanyOverridesParams.md) |  | 

## Methods

### NewListCompanyOverridesResponse

`func NewListCompanyOverridesResponse(data []CompanyOverrideResponseData, params ListCompanyOverridesParams, ) *ListCompanyOverridesResponse`

NewListCompanyOverridesResponse instantiates a new ListCompanyOverridesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyOverridesResponseWithDefaults

`func NewListCompanyOverridesResponseWithDefaults() *ListCompanyOverridesResponse`

NewListCompanyOverridesResponseWithDefaults instantiates a new ListCompanyOverridesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCompanyOverridesResponse) GetData() []CompanyOverrideResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompanyOverridesResponse) GetDataOk() (*[]CompanyOverrideResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompanyOverridesResponse) SetData(v []CompanyOverrideResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCompanyOverridesResponse) GetParams() ListCompanyOverridesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCompanyOverridesResponse) GetParamsOk() (*ListCompanyOverridesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCompanyOverridesResponse) SetParams(v ListCompanyOverridesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


