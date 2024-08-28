# CreatePlanGroupRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPlanId** | Pointer to **NullableString** |  | [optional] 
**PlanIds** | **[]string** |  | 

## Methods

### NewCreatePlanGroupRequestBody

`func NewCreatePlanGroupRequestBody(planIds []string, ) *CreatePlanGroupRequestBody`

NewCreatePlanGroupRequestBody instantiates a new CreatePlanGroupRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanGroupRequestBodyWithDefaults

`func NewCreatePlanGroupRequestBodyWithDefaults() *CreatePlanGroupRequestBody`

NewCreatePlanGroupRequestBodyWithDefaults instantiates a new CreatePlanGroupRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPlanId

`func (o *CreatePlanGroupRequestBody) GetDefaultPlanId() string`

GetDefaultPlanId returns the DefaultPlanId field if non-nil, zero value otherwise.

### GetDefaultPlanIdOk

`func (o *CreatePlanGroupRequestBody) GetDefaultPlanIdOk() (*string, bool)`

GetDefaultPlanIdOk returns a tuple with the DefaultPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanId

`func (o *CreatePlanGroupRequestBody) SetDefaultPlanId(v string)`

SetDefaultPlanId sets DefaultPlanId field to given value.

### HasDefaultPlanId

`func (o *CreatePlanGroupRequestBody) HasDefaultPlanId() bool`

HasDefaultPlanId returns a boolean if a field has been set.

### SetDefaultPlanIdNil

`func (o *CreatePlanGroupRequestBody) SetDefaultPlanIdNil(b bool)`

 SetDefaultPlanIdNil sets the value for DefaultPlanId to be an explicit nil

### UnsetDefaultPlanId
`func (o *CreatePlanGroupRequestBody) UnsetDefaultPlanId()`

UnsetDefaultPlanId ensures that no value is present for DefaultPlanId, not even an explicit nil
### GetPlanIds

`func (o *CreatePlanGroupRequestBody) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *CreatePlanGroupRequestBody) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *CreatePlanGroupRequestBody) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


