# RulesDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | Pointer to [**FlagResponseData**](FlagResponseData.md) |  | [optional] 
**Rules** | [**[]RuleDetailResponseData**](RuleDetailResponseData.md) |  | 

## Methods

### NewRulesDetailResponseData

`func NewRulesDetailResponseData(rules []RuleDetailResponseData, ) *RulesDetailResponseData`

NewRulesDetailResponseData instantiates a new RulesDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesDetailResponseDataWithDefaults

`func NewRulesDetailResponseDataWithDefaults() *RulesDetailResponseData`

NewRulesDetailResponseDataWithDefaults instantiates a new RulesDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *RulesDetailResponseData) GetFlag() FlagResponseData`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *RulesDetailResponseData) GetFlagOk() (*FlagResponseData, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *RulesDetailResponseData) SetFlag(v FlagResponseData)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *RulesDetailResponseData) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetRules

`func (o *RulesDetailResponseData) GetRules() []RuleDetailResponseData`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RulesDetailResponseData) GetRulesOk() (*[]RuleDetailResponseData, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RulesDetailResponseData) SetRules(v []RuleDetailResponseData)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


