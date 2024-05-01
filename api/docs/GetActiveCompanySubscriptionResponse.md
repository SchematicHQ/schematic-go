# GetActiveCompanySubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompanySubscriptionResponseData**](CompanySubscriptionResponseData.md) | The returned resources | 
**Params** | [**GetActiveCompanySubscriptionParams**](GetActiveCompanySubscriptionParams.md) |  | 

## Methods

### NewGetActiveCompanySubscriptionResponse

`func NewGetActiveCompanySubscriptionResponse(data []CompanySubscriptionResponseData, params GetActiveCompanySubscriptionParams, ) *GetActiveCompanySubscriptionResponse`

NewGetActiveCompanySubscriptionResponse instantiates a new GetActiveCompanySubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActiveCompanySubscriptionResponseWithDefaults

`func NewGetActiveCompanySubscriptionResponseWithDefaults() *GetActiveCompanySubscriptionResponse`

NewGetActiveCompanySubscriptionResponseWithDefaults instantiates a new GetActiveCompanySubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetActiveCompanySubscriptionResponse) GetData() []CompanySubscriptionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetActiveCompanySubscriptionResponse) GetDataOk() (*[]CompanySubscriptionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetActiveCompanySubscriptionResponse) SetData(v []CompanySubscriptionResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetActiveCompanySubscriptionResponse) GetParams() GetActiveCompanySubscriptionParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetActiveCompanySubscriptionResponse) GetParamsOk() (*GetActiveCompanySubscriptionParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetActiveCompanySubscriptionResponse) SetParams(v GetActiveCompanySubscriptionParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


