# CurrencyExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 货币兑换比率的唯一标识符。 | [optional] 
**FromCurrencyCode** | Pointer to **NullableString** | 兑换的源货币代码，例如 &#39;USD&#39;。 | [optional] 
**ToCurrencyCode** | Pointer to **NullableString** | 兑换的目标货币代码，例如 &#39;CNY&#39;。 | [optional] 
**ExchangeRate** | Pointer to **int64** | 从源货币到目标货币的兑换比率。例如，1 USD &#x3D; 6.5 CNY。 | [optional] 
**Description** | Pointer to **NullableString** | 兑换比率的详细描述信息。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 货币兑换比率的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 货币兑换比率的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewCurrencyExchangeRate

`func NewCurrencyExchangeRate() *CurrencyExchangeRate`

NewCurrencyExchangeRate instantiates a new CurrencyExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyExchangeRateWithDefaults

`func NewCurrencyExchangeRateWithDefaults() *CurrencyExchangeRate`

NewCurrencyExchangeRateWithDefaults instantiates a new CurrencyExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyExchangeRate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyExchangeRate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyExchangeRate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CurrencyExchangeRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFromCurrencyCode

`func (o *CurrencyExchangeRate) GetFromCurrencyCode() string`

GetFromCurrencyCode returns the FromCurrencyCode field if non-nil, zero value otherwise.

### GetFromCurrencyCodeOk

`func (o *CurrencyExchangeRate) GetFromCurrencyCodeOk() (*string, bool)`

GetFromCurrencyCodeOk returns a tuple with the FromCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrencyCode

`func (o *CurrencyExchangeRate) SetFromCurrencyCode(v string)`

SetFromCurrencyCode sets FromCurrencyCode field to given value.

### HasFromCurrencyCode

`func (o *CurrencyExchangeRate) HasFromCurrencyCode() bool`

HasFromCurrencyCode returns a boolean if a field has been set.

### SetFromCurrencyCodeNil

`func (o *CurrencyExchangeRate) SetFromCurrencyCodeNil(b bool)`

 SetFromCurrencyCodeNil sets the value for FromCurrencyCode to be an explicit nil

### UnsetFromCurrencyCode
`func (o *CurrencyExchangeRate) UnsetFromCurrencyCode()`

UnsetFromCurrencyCode ensures that no value is present for FromCurrencyCode, not even an explicit nil
### GetToCurrencyCode

`func (o *CurrencyExchangeRate) GetToCurrencyCode() string`

GetToCurrencyCode returns the ToCurrencyCode field if non-nil, zero value otherwise.

### GetToCurrencyCodeOk

`func (o *CurrencyExchangeRate) GetToCurrencyCodeOk() (*string, bool)`

GetToCurrencyCodeOk returns a tuple with the ToCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyCode

`func (o *CurrencyExchangeRate) SetToCurrencyCode(v string)`

SetToCurrencyCode sets ToCurrencyCode field to given value.

### HasToCurrencyCode

`func (o *CurrencyExchangeRate) HasToCurrencyCode() bool`

HasToCurrencyCode returns a boolean if a field has been set.

### SetToCurrencyCodeNil

`func (o *CurrencyExchangeRate) SetToCurrencyCodeNil(b bool)`

 SetToCurrencyCodeNil sets the value for ToCurrencyCode to be an explicit nil

### UnsetToCurrencyCode
`func (o *CurrencyExchangeRate) UnsetToCurrencyCode()`

UnsetToCurrencyCode ensures that no value is present for ToCurrencyCode, not even an explicit nil
### GetExchangeRate

`func (o *CurrencyExchangeRate) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *CurrencyExchangeRate) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *CurrencyExchangeRate) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *CurrencyExchangeRate) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetDescription

`func (o *CurrencyExchangeRate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurrencyExchangeRate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurrencyExchangeRate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurrencyExchangeRate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CurrencyExchangeRate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CurrencyExchangeRate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreateDate

`func (o *CurrencyExchangeRate) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *CurrencyExchangeRate) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *CurrencyExchangeRate) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *CurrencyExchangeRate) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *CurrencyExchangeRate) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *CurrencyExchangeRate) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *CurrencyExchangeRate) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *CurrencyExchangeRate) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


