# ListCustomersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedToImport** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewListCustomersParams

`func NewListCustomersParams() *ListCustomersParams`

NewListCustomersParams instantiates a new ListCustomersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomersParamsWithDefaults

`func NewListCustomersParamsWithDefaults() *ListCustomersParams`

NewListCustomersParamsWithDefaults instantiates a new ListCustomersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedToImport

`func (o *ListCustomersParams) GetFailedToImport() bool`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *ListCustomersParams) GetFailedToImportOk() (*bool, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *ListCustomersParams) SetFailedToImport(v bool)`

SetFailedToImport sets FailedToImport field to given value.

### HasFailedToImport

`func (o *ListCustomersParams) HasFailedToImport() bool`

HasFailedToImport returns a boolean if a field has been set.

### GetLimit

`func (o *ListCustomersParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCustomersParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCustomersParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListCustomersParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *ListCustomersParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCustomersParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCustomersParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCustomersParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *ListCustomersParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCustomersParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCustomersParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListCustomersParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *ListCustomersParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ListCustomersParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ListCustomersParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ListCustomersParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


