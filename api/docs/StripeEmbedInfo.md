# StripeEmbedInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerEkey** | Pointer to **NullableString** |  | [optional] 
**PublishableKey** | **string** |  | 

## Methods

### NewStripeEmbedInfo

`func NewStripeEmbedInfo(publishableKey string, ) *StripeEmbedInfo`

NewStripeEmbedInfo instantiates a new StripeEmbedInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeEmbedInfoWithDefaults

`func NewStripeEmbedInfoWithDefaults() *StripeEmbedInfo`

NewStripeEmbedInfoWithDefaults instantiates a new StripeEmbedInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerEkey

`func (o *StripeEmbedInfo) GetCustomerEkey() string`

GetCustomerEkey returns the CustomerEkey field if non-nil, zero value otherwise.

### GetCustomerEkeyOk

`func (o *StripeEmbedInfo) GetCustomerEkeyOk() (*string, bool)`

GetCustomerEkeyOk returns a tuple with the CustomerEkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEkey

`func (o *StripeEmbedInfo) SetCustomerEkey(v string)`

SetCustomerEkey sets CustomerEkey field to given value.

### HasCustomerEkey

`func (o *StripeEmbedInfo) HasCustomerEkey() bool`

HasCustomerEkey returns a boolean if a field has been set.

### SetCustomerEkeyNil

`func (o *StripeEmbedInfo) SetCustomerEkeyNil(b bool)`

 SetCustomerEkeyNil sets the value for CustomerEkey to be an explicit nil

### UnsetCustomerEkey
`func (o *StripeEmbedInfo) UnsetCustomerEkey()`

UnsetCustomerEkey ensures that no value is present for CustomerEkey, not even an explicit nil
### GetPublishableKey

`func (o *StripeEmbedInfo) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *StripeEmbedInfo) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *StripeEmbedInfo) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


