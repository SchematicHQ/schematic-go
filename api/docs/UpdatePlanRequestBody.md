# UpdatePlanRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**PlanType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdatePlanRequestBody

`func NewUpdatePlanRequestBody(name string, ) *UpdatePlanRequestBody`

NewUpdatePlanRequestBody instantiates a new UpdatePlanRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlanRequestBodyWithDefaults

`func NewUpdatePlanRequestBodyWithDefaults() *UpdatePlanRequestBody`

NewUpdatePlanRequestBodyWithDefaults instantiates a new UpdatePlanRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdatePlanRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePlanRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePlanRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePlanRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePlanRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePlanRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *UpdatePlanRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlanRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlanRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPlanType

`func (o *UpdatePlanRequestBody) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *UpdatePlanRequestBody) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *UpdatePlanRequestBody) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *UpdatePlanRequestBody) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### SetPlanTypeNil

`func (o *UpdatePlanRequestBody) SetPlanTypeNil(b bool)`

 SetPlanTypeNil sets the value for PlanType to be an explicit nil

### UnsetPlanType
`func (o *UpdatePlanRequestBody) UnsetPlanType()`

UnsetPlanType ensures that no value is present for PlanType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


