# ComponentHydrateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePlans** | [**[]CompanyPlanDetailResponseData**](CompanyPlanDetailResponseData.md) |  | 
**Company** | Pointer to [**CompanyDetailResponseData**](CompanyDetailResponseData.md) |  | [optional] 
**Component** | Pointer to [**ComponentResponseData**](ComponentResponseData.md) |  | [optional] 
**FeatureUsage** | Pointer to [**FeatureUsageDetailResponseData**](FeatureUsageDetailResponseData.md) |  | [optional] 
**StripeEmbed** | Pointer to [**StripeEmbedInfo**](StripeEmbedInfo.md) |  | [optional] 
**Subscription** | Pointer to [**CompanySubscriptionResponseData**](CompanySubscriptionResponseData.md) |  | [optional] 

## Methods

### NewComponentHydrateResponseData

`func NewComponentHydrateResponseData(activePlans []CompanyPlanDetailResponseData, ) *ComponentHydrateResponseData`

NewComponentHydrateResponseData instantiates a new ComponentHydrateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentHydrateResponseDataWithDefaults

`func NewComponentHydrateResponseDataWithDefaults() *ComponentHydrateResponseData`

NewComponentHydrateResponseDataWithDefaults instantiates a new ComponentHydrateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePlans

`func (o *ComponentHydrateResponseData) GetActivePlans() []CompanyPlanDetailResponseData`

GetActivePlans returns the ActivePlans field if non-nil, zero value otherwise.

### GetActivePlansOk

`func (o *ComponentHydrateResponseData) GetActivePlansOk() (*[]CompanyPlanDetailResponseData, bool)`

GetActivePlansOk returns a tuple with the ActivePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePlans

`func (o *ComponentHydrateResponseData) SetActivePlans(v []CompanyPlanDetailResponseData)`

SetActivePlans sets ActivePlans field to given value.


### GetCompany

`func (o *ComponentHydrateResponseData) GetCompany() CompanyDetailResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ComponentHydrateResponseData) GetCompanyOk() (*CompanyDetailResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ComponentHydrateResponseData) SetCompany(v CompanyDetailResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ComponentHydrateResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetComponent

`func (o *ComponentHydrateResponseData) GetComponent() ComponentResponseData`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentHydrateResponseData) GetComponentOk() (*ComponentResponseData, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentHydrateResponseData) SetComponent(v ComponentResponseData)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ComponentHydrateResponseData) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFeatureUsage

`func (o *ComponentHydrateResponseData) GetFeatureUsage() FeatureUsageDetailResponseData`

GetFeatureUsage returns the FeatureUsage field if non-nil, zero value otherwise.

### GetFeatureUsageOk

`func (o *ComponentHydrateResponseData) GetFeatureUsageOk() (*FeatureUsageDetailResponseData, bool)`

GetFeatureUsageOk returns a tuple with the FeatureUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsage

`func (o *ComponentHydrateResponseData) SetFeatureUsage(v FeatureUsageDetailResponseData)`

SetFeatureUsage sets FeatureUsage field to given value.

### HasFeatureUsage

`func (o *ComponentHydrateResponseData) HasFeatureUsage() bool`

HasFeatureUsage returns a boolean if a field has been set.

### GetStripeEmbed

`func (o *ComponentHydrateResponseData) GetStripeEmbed() StripeEmbedInfo`

GetStripeEmbed returns the StripeEmbed field if non-nil, zero value otherwise.

### GetStripeEmbedOk

`func (o *ComponentHydrateResponseData) GetStripeEmbedOk() (*StripeEmbedInfo, bool)`

GetStripeEmbedOk returns a tuple with the StripeEmbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeEmbed

`func (o *ComponentHydrateResponseData) SetStripeEmbed(v StripeEmbedInfo)`

SetStripeEmbed sets StripeEmbed field to given value.

### HasStripeEmbed

`func (o *ComponentHydrateResponseData) HasStripeEmbed() bool`

HasStripeEmbed returns a boolean if a field has been set.

### GetSubscription

`func (o *ComponentHydrateResponseData) GetSubscription() CompanySubscriptionResponseData`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *ComponentHydrateResponseData) GetSubscriptionOk() (*CompanySubscriptionResponseData, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *ComponentHydrateResponseData) SetSubscription(v CompanySubscriptionResponseData)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *ComponentHydrateResponseData) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


