# ExchangeCurrencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromCurrency** | **string** | 源虚拟币代码 | 
**Currency** | **string** | 目标虚拟币代码 | 
**Balance** | **int32** | 兑换额 | 
**Remark** | Pointer to **NullableString** | 备注 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 

## Methods

### NewExchangeCurrencyRequest

`func NewExchangeCurrencyRequest(fromCurrency string, currency string, balance int32, ) *ExchangeCurrencyRequest`

NewExchangeCurrencyRequest instantiates a new ExchangeCurrencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeCurrencyRequestWithDefaults

`func NewExchangeCurrencyRequestWithDefaults() *ExchangeCurrencyRequest`

NewExchangeCurrencyRequestWithDefaults instantiates a new ExchangeCurrencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromCurrency

`func (o *ExchangeCurrencyRequest) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *ExchangeCurrencyRequest) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *ExchangeCurrencyRequest) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.


### GetCurrency

`func (o *ExchangeCurrencyRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExchangeCurrencyRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExchangeCurrencyRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *ExchangeCurrencyRequest) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ExchangeCurrencyRequest) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ExchangeCurrencyRequest) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetRemark

`func (o *ExchangeCurrencyRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *ExchangeCurrencyRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *ExchangeCurrencyRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *ExchangeCurrencyRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *ExchangeCurrencyRequest) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *ExchangeCurrencyRequest) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetDescription

`func (o *ExchangeCurrencyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeCurrencyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeCurrencyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeCurrencyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExchangeCurrencyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExchangeCurrencyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *ExchangeCurrencyRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExchangeCurrencyRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExchangeCurrencyRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExchangeCurrencyRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ExchangeCurrencyRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ExchangeCurrencyRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


