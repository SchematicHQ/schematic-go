# CreatePlanRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceType** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Name** | **string** |  | 
**PlanType** | **string** |  | 

## Methods

### NewCreatePlanRequestBody

`func NewCreatePlanRequestBody(description string, name string, planType string, ) *CreatePlanRequestBody`

NewCreatePlanRequestBody instantiates a new CreatePlanRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlanRequestBodyWithDefaults

`func NewCreatePlanRequestBodyWithDefaults() *CreatePlanRequestBody`

NewCreatePlanRequestBodyWithDefaults instantiates a new CreatePlanRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceType

`func (o *CreatePlanRequestBody) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *CreatePlanRequestBody) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *CreatePlanRequestBody) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.

### HasAudienceType

`func (o *CreatePlanRequestBody) HasAudienceType() bool`

HasAudienceType returns a boolean if a field has been set.

### SetAudienceTypeNil

`func (o *CreatePlanRequestBody) SetAudienceTypeNil(b bool)`

 SetAudienceTypeNil sets the value for AudienceType to be an explicit nil

### UnsetAudienceType
`func (o *CreatePlanRequestBody) UnsetAudienceType()`

UnsetAudienceType ensures that no value is present for AudienceType, not even an explicit nil
### GetDescription

`func (o *CreatePlanRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePlanRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePlanRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreatePlanRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlanRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlanRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPlanType

`func (o *CreatePlanRequestBody) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CreatePlanRequestBody) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CreatePlanRequestBody) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


