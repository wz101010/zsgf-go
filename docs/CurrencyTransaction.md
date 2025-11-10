# CurrencyTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 货币交易记录的唯一标识符。 | [optional] 
**FromUserID** | Pointer to **int64** | 发起交易的发送方用户ID，若为转账交易时必填。 | [optional] 
**UserID** | Pointer to **int64** | 进行货币交易的用户ID。 | [optional] 
**TransactionType** | Pointer to **NullableString** | 货币交易的类型，例如 &#39;消费&#39;, &#39;奖励&#39;, &#39;兑换&#39;, &#39;转账&#39; 等。 | [optional] 
**CurrencyType** | Pointer to **NullableString** | 交易的货币类型，例如 &#39;USD&#39;, &#39;CNY&#39; 等。 | [optional] 
**CurrencyChange** | Pointer to **int64** | 货币的变动数量，正数表示增加，负数表示减少。 | [optional] 
**CurrencyBalance** | Pointer to **float64** | 交易完成后的货币余额。 | [optional] 
**Description** | Pointer to **NullableString** | 描述货币变动的具体原因或相关交易详情。 | [optional] 
**Status** | Pointer to **NullableString** | 货币交易的当前状态，例如 &#39;成功&#39;, &#39;失败&#39;, &#39;待审核&#39; 等。 | [optional] 
**Remark** | Pointer to **NullableString** | 交易的额外信息或管理员的备注。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记交易的标签。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 货币交易发生的时间，默认为当前时间。 | [optional] 

## Methods

### NewCurrencyTransaction

`func NewCurrencyTransaction() *CurrencyTransaction`

NewCurrencyTransaction instantiates a new CurrencyTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyTransactionWithDefaults

`func NewCurrencyTransactionWithDefaults() *CurrencyTransaction`

NewCurrencyTransactionWithDefaults instantiates a new CurrencyTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyTransaction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CurrencyTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFromUserID

`func (o *CurrencyTransaction) GetFromUserID() int64`

GetFromUserID returns the FromUserID field if non-nil, zero value otherwise.

### GetFromUserIDOk

`func (o *CurrencyTransaction) GetFromUserIDOk() (*int64, bool)`

GetFromUserIDOk returns a tuple with the FromUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserID

`func (o *CurrencyTransaction) SetFromUserID(v int64)`

SetFromUserID sets FromUserID field to given value.

### HasFromUserID

`func (o *CurrencyTransaction) HasFromUserID() bool`

HasFromUserID returns a boolean if a field has been set.

### GetUserID

`func (o *CurrencyTransaction) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CurrencyTransaction) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CurrencyTransaction) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *CurrencyTransaction) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetTransactionType

`func (o *CurrencyTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *CurrencyTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *CurrencyTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *CurrencyTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### SetTransactionTypeNil

`func (o *CurrencyTransaction) SetTransactionTypeNil(b bool)`

 SetTransactionTypeNil sets the value for TransactionType to be an explicit nil

### UnsetTransactionType
`func (o *CurrencyTransaction) UnsetTransactionType()`

UnsetTransactionType ensures that no value is present for TransactionType, not even an explicit nil
### GetCurrencyType

`func (o *CurrencyTransaction) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *CurrencyTransaction) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *CurrencyTransaction) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.

### HasCurrencyType

`func (o *CurrencyTransaction) HasCurrencyType() bool`

HasCurrencyType returns a boolean if a field has been set.

### SetCurrencyTypeNil

`func (o *CurrencyTransaction) SetCurrencyTypeNil(b bool)`

 SetCurrencyTypeNil sets the value for CurrencyType to be an explicit nil

### UnsetCurrencyType
`func (o *CurrencyTransaction) UnsetCurrencyType()`

UnsetCurrencyType ensures that no value is present for CurrencyType, not even an explicit nil
### GetCurrencyChange

`func (o *CurrencyTransaction) GetCurrencyChange() int64`

GetCurrencyChange returns the CurrencyChange field if non-nil, zero value otherwise.

### GetCurrencyChangeOk

`func (o *CurrencyTransaction) GetCurrencyChangeOk() (*int64, bool)`

GetCurrencyChangeOk returns a tuple with the CurrencyChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyChange

`func (o *CurrencyTransaction) SetCurrencyChange(v int64)`

SetCurrencyChange sets CurrencyChange field to given value.

### HasCurrencyChange

`func (o *CurrencyTransaction) HasCurrencyChange() bool`

HasCurrencyChange returns a boolean if a field has been set.

### GetCurrencyBalance

`func (o *CurrencyTransaction) GetCurrencyBalance() float64`

GetCurrencyBalance returns the CurrencyBalance field if non-nil, zero value otherwise.

### GetCurrencyBalanceOk

`func (o *CurrencyTransaction) GetCurrencyBalanceOk() (*float64, bool)`

GetCurrencyBalanceOk returns a tuple with the CurrencyBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyBalance

`func (o *CurrencyTransaction) SetCurrencyBalance(v float64)`

SetCurrencyBalance sets CurrencyBalance field to given value.

### HasCurrencyBalance

`func (o *CurrencyTransaction) HasCurrencyBalance() bool`

HasCurrencyBalance returns a boolean if a field has been set.

### GetDescription

`func (o *CurrencyTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurrencyTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurrencyTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurrencyTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurrencyTransaction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurrencyTransaction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *CurrencyTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrencyTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrencyTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CurrencyTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CurrencyTransaction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CurrencyTransaction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRemark

`func (o *CurrencyTransaction) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CurrencyTransaction) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CurrencyTransaction) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CurrencyTransaction) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *CurrencyTransaction) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *CurrencyTransaction) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetTags

`func (o *CurrencyTransaction) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CurrencyTransaction) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CurrencyTransaction) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CurrencyTransaction) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CurrencyTransaction) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CurrencyTransaction) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCreateDate

`func (o *CurrencyTransaction) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *CurrencyTransaction) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *CurrencyTransaction) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *CurrencyTransaction) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


