# RechargePointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | 虚拟币代码 | 
**Balance** | **int32** | 充值额 | 
**Remark** | Pointer to **NullableString** | 备注 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 

## Methods

### NewRechargePointRequest

`func NewRechargePointRequest(currency string, balance int32, ) *RechargePointRequest`

NewRechargePointRequest instantiates a new RechargePointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRechargePointRequestWithDefaults

`func NewRechargePointRequestWithDefaults() *RechargePointRequest`

NewRechargePointRequestWithDefaults instantiates a new RechargePointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *RechargePointRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RechargePointRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RechargePointRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBalance

`func (o *RechargePointRequest) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *RechargePointRequest) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *RechargePointRequest) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetRemark

`func (o *RechargePointRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *RechargePointRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *RechargePointRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *RechargePointRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *RechargePointRequest) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *RechargePointRequest) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetDescription

`func (o *RechargePointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RechargePointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RechargePointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RechargePointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RechargePointRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RechargePointRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *RechargePointRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RechargePointRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RechargePointRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RechargePointRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *RechargePointRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *RechargePointRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


