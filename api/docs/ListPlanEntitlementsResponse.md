# ListPlanEntitlementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PlanEntitlementResponseData**](PlanEntitlementResponseData.md) | The returned resources | 
**Params** | [**ListPlanEntitlementsParams**](ListPlanEntitlementsParams.md) |  | 

## Methods

### NewListPlanEntitlementsResponse

`func NewListPlanEntitlementsResponse(data []PlanEntitlementResponseData, params ListPlanEntitlementsParams, ) *ListPlanEntitlementsResponse`

NewListPlanEntitlementsResponse instantiates a new ListPlanEntitlementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlanEntitlementsResponseWithDefaults

`func NewListPlanEntitlementsResponseWithDefaults() *ListPlanEntitlementsResponse`

NewListPlanEntitlementsResponseWithDefaults instantiates a new ListPlanEntitlementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPlanEntitlementsResponse) GetData() []PlanEntitlementResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPlanEntitlementsResponse) GetDataOk() (*[]PlanEntitlementResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPlanEntitlementsResponse) SetData(v []PlanEntitlementResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListPlanEntitlementsResponse) GetParams() ListPlanEntitlementsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListPlanEntitlementsResponse) GetParamsOk() (*ListPlanEntitlementsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListPlanEntitlementsResponse) SetParams(v ListPlanEntitlementsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


