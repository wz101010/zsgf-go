# Currency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 货币的唯一标识符。 | [optional] 
**Name** | Pointer to **NullableString** | 货币的名称，例如 &#39;人民币&#39;, &#39;美元&#39; 等。 | [optional] 
**Code** | Pointer to **NullableString** | 货币的ISO标准代码，例如 &#39;CNY&#39;, &#39;USD&#39; 等。 | [optional] 
**Symbol** | Pointer to **NullableString** | 货币的符号，例如 &#39;$&#39;, &#39;€&#39;, &#39;¥&#39; 等。 | [optional] 
**Issuer** | Pointer to **NullableString** | 发行该货币的机构或国家名称。 | [optional] 
**CurrencyType** | Pointer to **NullableString** | 货币的类型，例如 &#39;法定货币&#39;, &#39;加密货币&#39; 等。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记货币的标签。 | [optional] 
**Status** | Pointer to **bool** | 指示货币当前是否可用。 | [optional] 
**EnableVirtualRecharge** | Pointer to **bool** | 指示该货币是否允许进行虚拟充值。 | [optional] 
**EnableVirtualConsume** | Pointer to **bool** | 指示该货币是否允许进行虚拟消费。 | [optional] 
**Description** | Pointer to **NullableString** | 货币的详细描述信息。 | [optional] 
**FiatExchangeRate** | Pointer to **int64** | 该货币与法定货币的兑换比率。 | [optional] 
**FiatDailyRechargeLimit** | Pointer to **int64** | 每日通过法定货币充值的最大限额。 | [optional] 
**TotalSupply** | Pointer to **int64** | 货币的总供应量，0 表示无限制。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 货币记录的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 货币记录的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewCurrency

`func NewCurrency() *Currency`

NewCurrency instantiates a new Currency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithDefaults

`func NewCurrencyWithDefaults() *Currency`

NewCurrencyWithDefaults instantiates a new Currency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Currency) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Currency) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Currency) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Currency) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Currency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Currency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Currency) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Currency) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Currency) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Currency) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *Currency) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Currency) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Currency) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Currency) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Currency) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Currency) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSymbol

`func (o *Currency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Currency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Currency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Currency) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### SetSymbolNil

`func (o *Currency) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *Currency) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil
### GetIssuer

`func (o *Currency) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Currency) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Currency) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Currency) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *Currency) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *Currency) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetCurrencyType

`func (o *Currency) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *Currency) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *Currency) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.

### HasCurrencyType

`func (o *Currency) HasCurrencyType() bool`

HasCurrencyType returns a boolean if a field has been set.

### SetCurrencyTypeNil

`func (o *Currency) SetCurrencyTypeNil(b bool)`

 SetCurrencyTypeNil sets the value for CurrencyType to be an explicit nil

### UnsetCurrencyType
`func (o *Currency) UnsetCurrencyType()`

UnsetCurrencyType ensures that no value is present for CurrencyType, not even an explicit nil
### GetTags

`func (o *Currency) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Currency) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Currency) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Currency) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Currency) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Currency) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetStatus

`func (o *Currency) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Currency) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Currency) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Currency) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnableVirtualRecharge

`func (o *Currency) GetEnableVirtualRecharge() bool`

GetEnableVirtualRecharge returns the EnableVirtualRecharge field if non-nil, zero value otherwise.

### GetEnableVirtualRechargeOk

`func (o *Currency) GetEnableVirtualRechargeOk() (*bool, bool)`

GetEnableVirtualRechargeOk returns a tuple with the EnableVirtualRecharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVirtualRecharge

`func (o *Currency) SetEnableVirtualRecharge(v bool)`

SetEnableVirtualRecharge sets EnableVirtualRecharge field to given value.

### HasEnableVirtualRecharge

`func (o *Currency) HasEnableVirtualRecharge() bool`

HasEnableVirtualRecharge returns a boolean if a field has been set.

### GetEnableVirtualConsume

`func (o *Currency) GetEnableVirtualConsume() bool`

GetEnableVirtualConsume returns the EnableVirtualConsume field if non-nil, zero value otherwise.

### GetEnableVirtualConsumeOk

`func (o *Currency) GetEnableVirtualConsumeOk() (*bool, bool)`

GetEnableVirtualConsumeOk returns a tuple with the EnableVirtualConsume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVirtualConsume

`func (o *Currency) SetEnableVirtualConsume(v bool)`

SetEnableVirtualConsume sets EnableVirtualConsume field to given value.

### HasEnableVirtualConsume

`func (o *Currency) HasEnableVirtualConsume() bool`

HasEnableVirtualConsume returns a boolean if a field has been set.

### GetDescription

`func (o *Currency) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Currency) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Currency) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Currency) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Currency) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Currency) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFiatExchangeRate

`func (o *Currency) GetFiatExchangeRate() int64`

GetFiatExchangeRate returns the FiatExchangeRate field if non-nil, zero value otherwise.

### GetFiatExchangeRateOk

`func (o *Currency) GetFiatExchangeRateOk() (*int64, bool)`

GetFiatExchangeRateOk returns a tuple with the FiatExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatExchangeRate

`func (o *Currency) SetFiatExchangeRate(v int64)`

SetFiatExchangeRate sets FiatExchangeRate field to given value.

### HasFiatExchangeRate

`func (o *Currency) HasFiatExchangeRate() bool`

HasFiatExchangeRate returns a boolean if a field has been set.

### GetFiatDailyRechargeLimit

`func (o *Currency) GetFiatDailyRechargeLimit() int64`

GetFiatDailyRechargeLimit returns the FiatDailyRechargeLimit field if non-nil, zero value otherwise.

### GetFiatDailyRechargeLimitOk

`func (o *Currency) GetFiatDailyRechargeLimitOk() (*int64, bool)`

GetFiatDailyRechargeLimitOk returns a tuple with the FiatDailyRechargeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatDailyRechargeLimit

`func (o *Currency) SetFiatDailyRechargeLimit(v int64)`

SetFiatDailyRechargeLimit sets FiatDailyRechargeLimit field to given value.

### HasFiatDailyRechargeLimit

`func (o *Currency) HasFiatDailyRechargeLimit() bool`

HasFiatDailyRechargeLimit returns a boolean if a field has been set.

### GetTotalSupply

`func (o *Currency) GetTotalSupply() int64`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *Currency) GetTotalSupplyOk() (*int64, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *Currency) SetTotalSupply(v int64)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *Currency) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetCreateDate

`func (o *Currency) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Currency) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Currency) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Currency) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *Currency) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Currency) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Currency) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Currency) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


