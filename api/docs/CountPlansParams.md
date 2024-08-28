# CountPlansParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**HasProductId** | Pointer to **bool** | Filter out plans that do not have a billing product ID | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**PlanType** | Pointer to **string** | Filter by plan type | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**WithoutEntitlementFor** | Pointer to **string** | Filter out plans that already have a plan entitlement for the specified feature ID | [optional] 

## Methods

### NewCountPlansParams

`func NewCountPlansParams() *CountPlansParams`

NewCountPlansParams instantiates a new CountPlansParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountPlansParamsWithDefaults

`func NewCountPlansParamsWithDefaults() *CountPlansParams`

NewCountPlansParamsWithDefaults instantiates a new CountPlansParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CountPlansParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CountPlansParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CountPlansParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CountPlansParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetHasProductId

`func (o *CountPlansParams) GetHasProductId() bool`

GetHasProductId returns the HasProductId field if non-nil, zero value otherwise.

### GetHasProductIdOk

`func (o *CountPlansParams) GetHasProductIdOk() (*bool, bool)`

GetHasProductIdOk returns a tuple with the HasProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProductId

`func (o *CountPlansParams) SetHasProductId(v bool)`

SetHasProductId sets HasProductId field to given value.

### HasHasProductId

`func (o *CountPlansParams) HasHasProductId() bool`

HasHasProductId returns a boolean if a field has been set.

### GetIds

`func (o *CountPlansParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CountPlansParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CountPlansParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CountPlansParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLimit

`func (o *CountPlansParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CountPlansParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CountPlansParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CountPlansParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *CountPlansParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CountPlansParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CountPlansParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CountPlansParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPlanType

`func (o *CountPlansParams) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CountPlansParams) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CountPlansParams) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *CountPlansParams) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetQ

`func (o *CountPlansParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *CountPlansParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *CountPlansParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *CountPlansParams) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetWithoutEntitlementFor

`func (o *CountPlansParams) GetWithoutEntitlementFor() string`

GetWithoutEntitlementFor returns the WithoutEntitlementFor field if non-nil, zero value otherwise.

### GetWithoutEntitlementForOk

`func (o *CountPlansParams) GetWithoutEntitlementForOk() (*string, bool)`

GetWithoutEntitlementForOk returns a tuple with the WithoutEntitlementFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutEntitlementFor

`func (o *CountPlansParams) SetWithoutEntitlementFor(v string)`

SetWithoutEntitlementFor sets WithoutEntitlementFor field to given value.

### HasWithoutEntitlementFor

`func (o *CountPlansParams) HasWithoutEntitlementFor() bool`

HasWithoutEntitlementFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


