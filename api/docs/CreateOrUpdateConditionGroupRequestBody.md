# CreateOrUpdateConditionGroupRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**PlanId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateOrUpdateConditionGroupRequestBody

`func NewCreateOrUpdateConditionGroupRequestBody(conditions []CreateOrUpdateConditionRequestBody, ) *CreateOrUpdateConditionGroupRequestBody`

NewCreateOrUpdateConditionGroupRequestBody instantiates a new CreateOrUpdateConditionGroupRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConditionGroupRequestBodyWithDefaults

`func NewCreateOrUpdateConditionGroupRequestBodyWithDefaults() *CreateOrUpdateConditionGroupRequestBody`

NewCreateOrUpdateConditionGroupRequestBodyWithDefaults instantiates a new CreateOrUpdateConditionGroupRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CreateOrUpdateConditionGroupRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateOrUpdateConditionGroupRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateOrUpdateConditionGroupRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetFlagId

`func (o *CreateOrUpdateConditionGroupRequestBody) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *CreateOrUpdateConditionGroupRequestBody) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *CreateOrUpdateConditionGroupRequestBody) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *CreateOrUpdateConditionGroupRequestBody) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *CreateOrUpdateConditionGroupRequestBody) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *CreateOrUpdateConditionGroupRequestBody) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetId

`func (o *CreateOrUpdateConditionGroupRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrUpdateConditionGroupRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrUpdateConditionGroupRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrUpdateConditionGroupRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateOrUpdateConditionGroupRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateOrUpdateConditionGroupRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPlanId

`func (o *CreateOrUpdateConditionGroupRequestBody) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateOrUpdateConditionGroupRequestBody) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateOrUpdateConditionGroupRequestBody) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *CreateOrUpdateConditionGroupRequestBody) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *CreateOrUpdateConditionGroupRequestBody) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *CreateOrUpdateConditionGroupRequestBody) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


