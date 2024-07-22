# ListProductsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewListProductsParams

`func NewListProductsParams() *ListProductsParams`

NewListProductsParams instantiates a new ListProductsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductsParamsWithDefaults

`func NewListProductsParamsWithDefaults() *ListProductsParams`

NewListProductsParamsWithDefaults instantiates a new ListProductsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ListProductsParams) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ListProductsParams) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ListProductsParams) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ListProductsParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetLimit

`func (o *ListProductsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListProductsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListProductsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListProductsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *ListProductsParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListProductsParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListProductsParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListProductsParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *ListProductsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListProductsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListProductsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListProductsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *ListProductsParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ListProductsParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ListProductsParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ListProductsParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


