# ListPlansParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 
**WithoutEntitlementFor** | Pointer to **string** | Filter out plans that already have a plan entitlement for the specified feature ID | [optional] 

## Methods

### NewListPlansParams

`func NewListPlansParams() *ListPlansParams`

NewListPlansParams instantiates a new ListPlansParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlansParamsWithDefaults

`func NewListPlansParamsWithDefaults() *ListPlansParams`

NewListPlansParamsWithDefaults instantiates a new ListPlansParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ListPlansParams) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ListPlansParams) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ListPlansParams) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ListPlansParams) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetIds

`func (o *ListPlansParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListPlansParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListPlansParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListPlansParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLimit

`func (o *ListPlansParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListPlansParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListPlansParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListPlansParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListPlansParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListPlansParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListPlansParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListPlansParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *ListPlansParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ListPlansParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ListPlansParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ListPlansParams) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetWithoutEntitlementFor

`func (o *ListPlansParams) GetWithoutEntitlementFor() string`

GetWithoutEntitlementFor returns the WithoutEntitlementFor field if non-nil, zero value otherwise.

### GetWithoutEntitlementForOk

`func (o *ListPlansParams) GetWithoutEntitlementForOk() (*string, bool)`

GetWithoutEntitlementForOk returns a tuple with the WithoutEntitlementFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithoutEntitlementFor

`func (o *ListPlansParams) SetWithoutEntitlementFor(v string)`

SetWithoutEntitlementFor sets WithoutEntitlementFor field to given value.

### HasWithoutEntitlementFor

`func (o *ListPlansParams) HasWithoutEntitlementFor() bool`

HasWithoutEntitlementFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


