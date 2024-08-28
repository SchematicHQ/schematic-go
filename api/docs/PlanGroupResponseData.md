# PlanGroupResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPlanId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**PlanIds** | **[]string** |  | 

## Methods

### NewPlanGroupResponseData

`func NewPlanGroupResponseData(id string, planIds []string, ) *PlanGroupResponseData`

NewPlanGroupResponseData instantiates a new PlanGroupResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanGroupResponseDataWithDefaults

`func NewPlanGroupResponseDataWithDefaults() *PlanGroupResponseData`

NewPlanGroupResponseDataWithDefaults instantiates a new PlanGroupResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPlanId

`func (o *PlanGroupResponseData) GetDefaultPlanId() string`

GetDefaultPlanId returns the DefaultPlanId field if non-nil, zero value otherwise.

### GetDefaultPlanIdOk

`func (o *PlanGroupResponseData) GetDefaultPlanIdOk() (*string, bool)`

GetDefaultPlanIdOk returns a tuple with the DefaultPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanId

`func (o *PlanGroupResponseData) SetDefaultPlanId(v string)`

SetDefaultPlanId sets DefaultPlanId field to given value.

### HasDefaultPlanId

`func (o *PlanGroupResponseData) HasDefaultPlanId() bool`

HasDefaultPlanId returns a boolean if a field has been set.

### SetDefaultPlanIdNil

`func (o *PlanGroupResponseData) SetDefaultPlanIdNil(b bool)`

 SetDefaultPlanIdNil sets the value for DefaultPlanId to be an explicit nil

### UnsetDefaultPlanId
`func (o *PlanGroupResponseData) UnsetDefaultPlanId()`

UnsetDefaultPlanId ensures that no value is present for DefaultPlanId, not even an explicit nil
### GetId

`func (o *PlanGroupResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlanGroupResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlanGroupResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetPlanIds

`func (o *PlanGroupResponseData) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *PlanGroupResponseData) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *PlanGroupResponseData) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


