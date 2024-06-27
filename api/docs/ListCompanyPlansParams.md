# ListCompanyPlansParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 

## Methods

### NewListCompanyPlansParams

`func NewListCompanyPlansParams() *ListCompanyPlansParams`

NewListCompanyPlansParams instantiates a new ListCompanyPlansParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyPlansParamsWithDefaults

`func NewListCompanyPlansParamsWithDefaults() *ListCompanyPlansParams`

NewListCompanyPlansParamsWithDefaults instantiates a new ListCompanyPlansParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListCompanyPlansParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListCompanyPlansParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListCompanyPlansParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListCompanyPlansParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetLimit

`func (o *ListCompanyPlansParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCompanyPlansParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCompanyPlansParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCompanyPlansParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCompanyPlansParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCompanyPlansParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCompanyPlansParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCompanyPlansParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPlanId

`func (o *ListCompanyPlansParams) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ListCompanyPlansParams) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ListCompanyPlansParams) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ListCompanyPlansParams) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


