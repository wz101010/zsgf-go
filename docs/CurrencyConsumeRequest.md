# CurrencyConsumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | 虚拟币代码 | 
**Balance** | **int32** | 消费额 | 
**Remark** | Pointer to **NullableString** | 备注 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 

## Methods

### NewCurrencyConsumeRequest

`func NewCurrencyConsumeRequest(currency string, balance int32, ) *CurrencyConsumeRequest`

NewCurrencyConsumeRequest instantiates a new CurrencyConsumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyConsumeRequestWithDefaults

`func NewCurrencyConsumeRequestWithDefaults() *CurrencyConsumeRequest`

NewCurrencyConsumeRequestWithDefaults instantiates a new CurrencyConsumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CurrencyConsumeRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyConsumeRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyConsumeRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *CurrencyConsumeRequest) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CurrencyConsumeRequest) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CurrencyConsumeRequest) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetRemark

`func (o *CurrencyConsumeRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CurrencyConsumeRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CurrencyConsumeRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CurrencyConsumeRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *CurrencyConsumeRequest) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *CurrencyConsumeRequest) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetDescription

`func (o *CurrencyConsumeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurrencyConsumeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurrencyConsumeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurrencyConsumeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurrencyConsumeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurrencyConsumeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *CurrencyConsumeRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CurrencyConsumeRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CurrencyConsumeRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CurrencyConsumeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CurrencyConsumeRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CurrencyConsumeRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


