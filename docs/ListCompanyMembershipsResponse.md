# ListCompanyMembershipsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanyMembershipDetailResponseData**](CompanyMembershipDetailResponseData.md) | The returned resources | 
**Params** | [**ListCompanyMembershipsParams**](ListCompanyMembershipsParams.md) |  | 

## Methods

### NewListCompanyMembershipsResponse

`func NewListCompanyMembershipsResponse(data []CompanyMembershipDetailResponseData, params ListCompanyMembershipsParams, ) *ListCompanyMembershipsResponse`

NewListCompanyMembershipsResponse instantiates a new ListCompanyMembershipsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyMembershipsResponseWithDefaults

`func NewListCompanyMembershipsResponseWithDefaults() *ListCompanyMembershipsResponse`

NewListCompanyMembershipsResponseWithDefaults instantiates a new ListCompanyMembershipsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCompanyMembershipsResponse) GetData() []CompanyMembershipDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCompanyMembershipsResponse) GetDataOk() (*[]CompanyMembershipDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCompanyMembershipsResponse) SetData(v []CompanyMembershipDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCompanyMembershipsResponse) GetParams() ListCompanyMembershipsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCompanyMembershipsResponse) GetParamsOk() (*ListCompanyMembershipsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCompanyMembershipsResponse) SetParams(v ListCompanyMembershipsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


