# AlipayTradeQueryResponseApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**AlipayTradeQueryResponse**](AlipayTradeQueryResponse.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewAlipayTradeQueryResponseApiResponse

`func NewAlipayTradeQueryResponseApiResponse() *AlipayTradeQueryResponseApiResponse`

NewAlipayTradeQueryResponseApiResponse instantiates a new AlipayTradeQueryResponseApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlipayTradeQueryResponseApiResponseWithDefaults

`func NewAlipayTradeQueryResponseApiResponseWithDefaults() *AlipayTradeQueryResponseApiResponse`

NewAlipayTradeQueryResponseApiResponseWithDefaults instantiates a new AlipayTradeQueryResponseApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AlipayTradeQueryResponseApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AlipayTradeQueryResponseApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AlipayTradeQueryResponseApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AlipayTradeQueryResponseApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *AlipayTradeQueryResponseApiResponse) GetData() AlipayTradeQueryResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlipayTradeQueryResponseApiResponse) GetDataOk() (*AlipayTradeQueryResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlipayTradeQueryResponseApiResponse) SetData(v AlipayTradeQueryResponse)`

SetData sets Data field to given value.

### HasData

`func (o *AlipayTradeQueryResponseApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *AlipayTradeQueryResponseApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AlipayTradeQueryResponseApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AlipayTradeQueryResponseApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AlipayTradeQueryResponseApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AlipayTradeQueryResponseApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AlipayTradeQueryResponseApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


