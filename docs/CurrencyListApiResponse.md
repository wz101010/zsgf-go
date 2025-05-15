# CurrencyListApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**[]Currency**](Currency.md) | 响应数据 | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewCurrencyListApiResponse

`func NewCurrencyListApiResponse() *CurrencyListApiResponse`

NewCurrencyListApiResponse instantiates a new CurrencyListApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyListApiResponseWithDefaults

`func NewCurrencyListApiResponseWithDefaults() *CurrencyListApiResponse`

NewCurrencyListApiResponseWithDefaults instantiates a new CurrencyListApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CurrencyListApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyListApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyListApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CurrencyListApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *CurrencyListApiResponse) GetData() []Currency`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CurrencyListApiResponse) GetDataOk() (*[]Currency, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CurrencyListApiResponse) SetData(v []Currency)`

SetData sets Data field to given value.

### HasData

`func (o *CurrencyListApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CurrencyListApiResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CurrencyListApiResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *CurrencyListApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CurrencyListApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CurrencyListApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CurrencyListApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CurrencyListApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CurrencyListApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


