# ListCompaniesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Q** | Pointer to **string** | Search filter | [optional] 
**WithoutFeatureOverrideFor** | Pointer to **string** | Filter out companies that already have a company override for the specified feature ID | [optional] 
**WithoutPlan** | Pointer to **bool** | Filter out companies that have a plan | [optional] 

## Methods

### NewListCompaniesParams

`func NewListCompaniesParams() *ListCompaniesParams`

NewListCompaniesParams instantiates a new ListCompaniesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompaniesParamsWithDefaults

`func NewListCompaniesParamsWithDefaults() *ListCompaniesParams`

NewListCompaniesParamsWithDefaults instantiates a new ListCompaniesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListCompaniesParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListCompaniesParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListCompaniesParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListCompaniesParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLimit

`func (o *ListCompaniesParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCompaniesParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCompaniesParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCompaniesParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListCompaniesParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCompaniesParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCompaniesParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCompaniesParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPlanId

`func (o *ListCompaniesParams) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ListCompaniesParams) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ListCompaniesParams) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ListCompaniesParams) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQ

`func (o *ListCompaniesParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ListCompaniesParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ListCompaniesParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ListCompaniesParams) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetWithoutFeatureOverrideFor

`func (o *ListCompaniesParams) GetWithoutFeatureOverrideFor() string`

GetWithoutFeatureOverrideFor returns the WithoutFeatureOverrideFor field if non-nil, zero value otherwise.

### GetWithoutFeatureOverrideForOk

`func (o *ListCompaniesParams) GetWithoutFeatureOverrideForOk() (*string, bool)`

GetWithoutFeatureOverrideForOk returns a tuple with the WithoutFeatureOverrideFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutFeatureOverrideFor

`func (o *ListCompaniesParams) SetWithoutFeatureOverrideFor(v string)`

SetWithoutFeatureOverrideFor sets WithoutFeatureOverrideFor field to given value.

### HasWithoutFeatureOverrideFor

`func (o *ListCompaniesParams) HasWithoutFeatureOverrideFor() bool`

HasWithoutFeatureOverrideFor returns a boolean if a field has been set.

### GetWithoutPlan

`func (o *ListCompaniesParams) GetWithoutPlan() bool`

GetWithoutPlan returns the WithoutPlan field if non-nil, zero value otherwise.

### GetWithoutPlanOk

`func (o *ListCompaniesParams) GetWithoutPlanOk() (*bool, bool)`

GetWithoutPlanOk returns a tuple with the WithoutPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutPlan

`func (o *ListCompaniesParams) SetWithoutPlan(v bool)`

SetWithoutPlan sets WithoutPlan field to given value.

### HasWithoutPlan

`func (o *ListCompaniesParams) HasWithoutPlan() bool`

HasWithoutPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


