# ExchangeRatePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToCurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ExchangeRate** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExchangeRatePutRequest

`func NewExchangeRatePutRequest() *ExchangeRatePutRequest`

NewExchangeRatePutRequest instantiates a new ExchangeRatePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeRatePutRequestWithDefaults

`func NewExchangeRatePutRequestWithDefaults() *ExchangeRatePutRequest`

NewExchangeRatePutRequestWithDefaults instantiates a new ExchangeRatePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToCurrencyCode

`func (o *ExchangeRatePutRequest) GetToCurrencyCode() string`

GetToCurrencyCode returns the ToCurrencyCode field if non-nil, zero value otherwise.

### GetToCurrencyCodeOk

`func (o *ExchangeRatePutRequest) GetToCurrencyCodeOk() (*string, bool)`

GetToCurrencyCodeOk returns a tuple with the ToCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrencyCode

`func (o *ExchangeRatePutRequest) SetToCurrencyCode(v string)`

SetToCurrencyCode sets ToCurrencyCode field to given value.

### HasToCurrencyCode

`func (o *ExchangeRatePutRequest) HasToCurrencyCode() bool`

HasToCurrencyCode returns a boolean if a field has been set.

### SetToCurrencyCodeNil

`func (o *ExchangeRatePutRequest) SetToCurrencyCodeNil(b bool)`

 SetToCurrencyCodeNil sets the value for ToCurrencyCode to be an explicit nil

### UnsetToCurrencyCode
`func (o *ExchangeRatePutRequest) UnsetToCurrencyCode()`

UnsetToCurrencyCode ensures that no value is present for ToCurrencyCode, not even an explicit nil
### GetExchangeRate

`func (o *ExchangeRatePutRequest) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *ExchangeRatePutRequest) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *ExchangeRatePutRequest) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *ExchangeRatePutRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetDescription

`func (o *ExchangeRatePutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeRatePutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeRatePutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeRatePutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExchangeRatePutRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExchangeRatePutRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


