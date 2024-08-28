# PlanGroupDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPlan** | Pointer to [**PlanGroupPlanDetailResponseData**](PlanGroupPlanDetailResponseData.md) |  | [optional] 
**DefaultPlanId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Plans** | [**[]PlanGroupPlanDetailResponseData**](PlanGroupPlanDetailResponseData.md) |  | 

## Methods

### NewPlanGroupDetailResponseData

`func NewPlanGroupDetailResponseData(id string, plans []PlanGroupPlanDetailResponseData, ) *PlanGroupDetailResponseData`

NewPlanGroupDetailResponseData instantiates a new PlanGroupDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanGroupDetailResponseDataWithDefaults

`func NewPlanGroupDetailResponseDataWithDefaults() *PlanGroupDetailResponseData`

NewPlanGroupDetailResponseDataWithDefaults instantiates a new PlanGroupDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPlan

`func (o *PlanGroupDetailResponseData) GetDefaultPlan() PlanGroupPlanDetailResponseData`

GetDefaultPlan returns the DefaultPlan field if non-nil, zero value otherwise.

### GetDefaultPlanOk

`func (o *PlanGroupDetailResponseData) GetDefaultPlanOk() (*PlanGroupPlanDetailResponseData, bool)`

GetDefaultPlanOk returns a tuple with the DefaultPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlan

`func (o *PlanGroupDetailResponseData) SetDefaultPlan(v PlanGroupPlanDetailResponseData)`

SetDefaultPlan sets DefaultPlan field to given value.

### HasDefaultPlan

`func (o *PlanGroupDetailResponseData) HasDefaultPlan() bool`

HasDefaultPlan returns a boolean if a field has been set.

### GetDefaultPlanId

`func (o *PlanGroupDetailResponseData) GetDefaultPlanId() string`

GetDefaultPlanId returns the DefaultPlanId field if non-nil, zero value otherwise.

### GetDefaultPlanIdOk

`func (o *PlanGroupDetailResponseData) GetDefaultPlanIdOk() (*string, bool)`

GetDefaultPlanIdOk returns a tuple with the DefaultPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanId

`func (o *PlanGroupDetailResponseData) SetDefaultPlanId(v string)`

SetDefaultPlanId sets DefaultPlanId field to given value.

### HasDefaultPlanId

`func (o *PlanGroupDetailResponseData) HasDefaultPlanId() bool`

HasDefaultPlanId returns a boolean if a field has been set.

### SetDefaultPlanIdNil

`func (o *PlanGroupDetailResponseData) SetDefaultPlanIdNil(b bool)`

 SetDefaultPlanIdNil sets the value for DefaultPlanId to be an explicit nil

### UnsetDefaultPlanId
`func (o *PlanGroupDetailResponseData) UnsetDefaultPlanId()`

UnsetDefaultPlanId ensures that no value is present for DefaultPlanId, not even an explicit nil
### GetId

`func (o *PlanGroupDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanGroupDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanGroupDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetPlans

`func (o *PlanGroupDetailResponseData) GetPlans() []PlanGroupPlanDetailResponseData`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *PlanGroupDetailResponseData) GetPlansOk() (*[]PlanGroupPlanDetailResponseData, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *PlanGroupDetailResponseData) SetPlans(v []PlanGroupPlanDetailResponseData)`

SetPlans sets Plans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


