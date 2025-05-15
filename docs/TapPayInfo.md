# TapPayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMediumType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTapPayInfo

`func NewTapPayInfo() *TapPayInfo`

NewTapPayInfo instantiates a new TapPayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTapPayInfoWithDefaults

`func NewTapPayInfoWithDefaults() *TapPayInfo`

NewTapPayInfoWithDefaults instantiates a new TapPayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMediumType

`func (o *TapPayInfo) GetPaymentMediumType() string`

GetPaymentMediumType returns the PaymentMediumType field if non-nil, zero value otherwise.

### GetPaymentMediumTypeOk

`func (o *TapPayInfo) GetPaymentMediumTypeOk() (*string, bool)`

GetPaymentMediumTypeOk returns a tuple with the PaymentMediumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMediumType

`func (o *TapPayInfo) SetPaymentMediumType(v string)`

SetPaymentMediumType sets PaymentMediumType field to given value.

### HasPaymentMediumType

`func (o *TapPayInfo) HasPaymentMediumType() bool`

HasPaymentMediumType returns a boolean if a field has been set.

### SetPaymentMediumTypeNil

`func (o *TapPayInfo) SetPaymentMediumTypeNil(b bool)`

 SetPaymentMediumTypeNil sets the value for PaymentMediumType to be an explicit nil

### UnsetPaymentMediumType
`func (o *TapPayInfo) UnsetPaymentMediumType()`

UnsetPaymentMediumType ensures that no value is present for PaymentMediumType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


