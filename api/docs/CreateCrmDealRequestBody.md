# CreateCrmDealRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arr** | Pointer to **NullableString** |  | [optional] 
**CrmCompanyId** | Pointer to **NullableString** |  | [optional] 
**CrmCompanyKey** | **string** |  | 
**CrmProductId** | Pointer to **NullableString** |  | [optional] 
**CrmType** | **string** |  | 
**DealExternalId** | **string** |  | 
**DealName** | Pointer to **NullableString** |  | [optional] 
**DealStage** | Pointer to **NullableString** |  | [optional] 
**Mrr** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCrmDealRequestBody

`func NewCreateCrmDealRequestBody(crmCompanyKey string, crmType string, dealExternalId string, ) *CreateCrmDealRequestBody`

NewCreateCrmDealRequestBody instantiates a new CreateCrmDealRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCrmDealRequestBodyWithDefaults

`func NewCreateCrmDealRequestBodyWithDefaults() *CreateCrmDealRequestBody`

NewCreateCrmDealRequestBodyWithDefaults instantiates a new CreateCrmDealRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArr

`func (o *CreateCrmDealRequestBody) GetArr() string`

GetArr returns the Arr field if non-nil, zero value otherwise.

### GetArrOk

`func (o *CreateCrmDealRequestBody) GetArrOk() (*string, bool)`

GetArrOk returns a tuple with the Arr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArr

`func (o *CreateCrmDealRequestBody) SetArr(v string)`

SetArr sets Arr field to given value.

### HasArr

`func (o *CreateCrmDealRequestBody) HasArr() bool`

HasArr returns a boolean if a field has been set.

### SetArrNil

`func (o *CreateCrmDealRequestBody) SetArrNil(b bool)`

 SetArrNil sets the value for Arr to be an explicit nil

### UnsetArr
`func (o *CreateCrmDealRequestBody) UnsetArr()`

UnsetArr ensures that no value is present for Arr, not even an explicit nil
### GetCrmCompanyId

`func (o *CreateCrmDealRequestBody) GetCrmCompanyId() string`

GetCrmCompanyId returns the CrmCompanyId field if non-nil, zero value otherwise.

### GetCrmCompanyIdOk

`func (o *CreateCrmDealRequestBody) GetCrmCompanyIdOk() (*string, bool)`

GetCrmCompanyIdOk returns a tuple with the CrmCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmCompanyId

`func (o *CreateCrmDealRequestBody) SetCrmCompanyId(v string)`

SetCrmCompanyId sets CrmCompanyId field to given value.

### HasCrmCompanyId

`func (o *CreateCrmDealRequestBody) HasCrmCompanyId() bool`

HasCrmCompanyId returns a boolean if a field has been set.

### SetCrmCompanyIdNil

`func (o *CreateCrmDealRequestBody) SetCrmCompanyIdNil(b bool)`

 SetCrmCompanyIdNil sets the value for CrmCompanyId to be an explicit nil

### UnsetCrmCompanyId
`func (o *CreateCrmDealRequestBody) UnsetCrmCompanyId()`

UnsetCrmCompanyId ensures that no value is present for CrmCompanyId, not even an explicit nil
### GetCrmCompanyKey

`func (o *CreateCrmDealRequestBody) GetCrmCompanyKey() string`

GetCrmCompanyKey returns the CrmCompanyKey field if non-nil, zero value otherwise.

### GetCrmCompanyKeyOk

`func (o *CreateCrmDealRequestBody) GetCrmCompanyKeyOk() (*string, bool)`

GetCrmCompanyKeyOk returns a tuple with the CrmCompanyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmCompanyKey

`func (o *CreateCrmDealRequestBody) SetCrmCompanyKey(v string)`

SetCrmCompanyKey sets CrmCompanyKey field to given value.


### GetCrmProductId

`func (o *CreateCrmDealRequestBody) GetCrmProductId() string`

GetCrmProductId returns the CrmProductId field if non-nil, zero value otherwise.

### GetCrmProductIdOk

`func (o *CreateCrmDealRequestBody) GetCrmProductIdOk() (*string, bool)`

GetCrmProductIdOk returns a tuple with the CrmProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmProductId

`func (o *CreateCrmDealRequestBody) SetCrmProductId(v string)`

SetCrmProductId sets CrmProductId field to given value.

### HasCrmProductId

`func (o *CreateCrmDealRequestBody) HasCrmProductId() bool`

HasCrmProductId returns a boolean if a field has been set.

### SetCrmProductIdNil

`func (o *CreateCrmDealRequestBody) SetCrmProductIdNil(b bool)`

 SetCrmProductIdNil sets the value for CrmProductId to be an explicit nil

### UnsetCrmProductId
`func (o *CreateCrmDealRequestBody) UnsetCrmProductId()`

UnsetCrmProductId ensures that no value is present for CrmProductId, not even an explicit nil
### GetCrmType

`func (o *CreateCrmDealRequestBody) GetCrmType() string`

GetCrmType returns the CrmType field if non-nil, zero value otherwise.

### GetCrmTypeOk

`func (o *CreateCrmDealRequestBody) GetCrmTypeOk() (*string, bool)`

GetCrmTypeOk returns a tuple with the CrmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmType

`func (o *CreateCrmDealRequestBody) SetCrmType(v string)`

SetCrmType sets CrmType field to given value.


### GetDealExternalId

`func (o *CreateCrmDealRequestBody) GetDealExternalId() string`

GetDealExternalId returns the DealExternalId field if non-nil, zero value otherwise.

### GetDealExternalIdOk

`func (o *CreateCrmDealRequestBody) GetDealExternalIdOk() (*string, bool)`

GetDealExternalIdOk returns a tuple with the DealExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExternalId

`func (o *CreateCrmDealRequestBody) SetDealExternalId(v string)`

SetDealExternalId sets DealExternalId field to given value.


### GetDealName

`func (o *CreateCrmDealRequestBody) GetDealName() string`

GetDealName returns the DealName field if non-nil, zero value otherwise.

### GetDealNameOk

`func (o *CreateCrmDealRequestBody) GetDealNameOk() (*string, bool)`

GetDealNameOk returns a tuple with the DealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealName

`func (o *CreateCrmDealRequestBody) SetDealName(v string)`

SetDealName sets DealName field to given value.

### HasDealName

`func (o *CreateCrmDealRequestBody) HasDealName() bool`

HasDealName returns a boolean if a field has been set.

### SetDealNameNil

`func (o *CreateCrmDealRequestBody) SetDealNameNil(b bool)`

 SetDealNameNil sets the value for DealName to be an explicit nil

### UnsetDealName
`func (o *CreateCrmDealRequestBody) UnsetDealName()`

UnsetDealName ensures that no value is present for DealName, not even an explicit nil
### GetDealStage

`func (o *CreateCrmDealRequestBody) GetDealStage() string`

GetDealStage returns the DealStage field if non-nil, zero value otherwise.

### GetDealStageOk

`func (o *CreateCrmDealRequestBody) GetDealStageOk() (*string, bool)`

GetDealStageOk returns a tuple with the DealStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStage

`func (o *CreateCrmDealRequestBody) SetDealStage(v string)`

SetDealStage sets DealStage field to given value.

### HasDealStage

`func (o *CreateCrmDealRequestBody) HasDealStage() bool`

HasDealStage returns a boolean if a field has been set.

### SetDealStageNil

`func (o *CreateCrmDealRequestBody) SetDealStageNil(b bool)`

 SetDealStageNil sets the value for DealStage to be an explicit nil

### UnsetDealStage
`func (o *CreateCrmDealRequestBody) UnsetDealStage()`

UnsetDealStage ensures that no value is present for DealStage, not even an explicit nil
### GetMrr

`func (o *CreateCrmDealRequestBody) GetMrr() string`

GetMrr returns the Mrr field if non-nil, zero value otherwise.

### GetMrrOk

`func (o *CreateCrmDealRequestBody) GetMrrOk() (*string, bool)`

GetMrrOk returns a tuple with the Mrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrr

`func (o *CreateCrmDealRequestBody) SetMrr(v string)`

SetMrr sets Mrr field to given value.

### HasMrr

`func (o *CreateCrmDealRequestBody) HasMrr() bool`

HasMrr returns a boolean if a field has been set.

### SetMrrNil

`func (o *CreateCrmDealRequestBody) SetMrrNil(b bool)`

 SetMrrNil sets the value for Mrr to be an explicit nil

### UnsetMrr
`func (o *CreateCrmDealRequestBody) UnsetMrr()`

UnsetMrr ensures that no value is present for Mrr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


