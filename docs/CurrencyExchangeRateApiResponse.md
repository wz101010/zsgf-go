# CurrencyExchangeRateApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**CurrencyExchangeRate**](CurrencyExchangeRate.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewCurrencyExchangeRateApiResponse

`func NewCurrencyExchangeRateApiResponse() *CurrencyExchangeRateApiResponse`

NewCurrencyExchangeRateApiResponse instantiates a new CurrencyExchangeRateApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyExchangeRateApiResponseWithDefaults

`func NewCurrencyExchangeRateApiResponseWithDefaults() *CurrencyExchangeRateApiResponse`

NewCurrencyExchangeRateApiResponseWithDefaults instantiates a new CurrencyExchangeRateApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CurrencyExchangeRateApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyExchangeRateApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyExchangeRateApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CurrencyExchangeRateApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CurrencyExchangeRateApiResponse) GetData() CurrencyExchangeRate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CurrencyExchangeRateApiResponse) GetDataOk() (*CurrencyExchangeRate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CurrencyExchangeRateApiResponse) SetData(v CurrencyExchangeRate)`

SetData sets Data field to given value.

### HasData

`func (o *CurrencyExchangeRateApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *CurrencyExchangeRateApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CurrencyExchangeRateApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CurrencyExchangeRateApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CurrencyExchangeRateApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CurrencyExchangeRateApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CurrencyExchangeRateApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


