# ReturnPageNotifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **NullableString** | 应用ID | [optional] 
**AuthAppId** | Pointer to **NullableString** | 授权应用ID | [optional] 
**Charset** | Pointer to **NullableString** | 字符集 | [optional] 
**Method** | Pointer to **NullableString** | 接口名称 | [optional] 
**OutTradeNo** | Pointer to **NullableString** | 商户订单号 | [optional] 
**SellerId** | Pointer to **NullableString** | 卖家支付宝用户号 | [optional] 
**Sign** | Pointer to **NullableString** | 签名 | [optional] 
**SignType** | Pointer to **NullableString** | 签名类型 | [optional] 
**Timestamp** | Pointer to **NullableString** | 时间戳 | [optional] 
**TotalAmount** | Pointer to **NullableString** | 订单总金额 | [optional] 
**TradeNo** | Pointer to **NullableString** | 支付宝交易号 | [optional] 
**Version** | Pointer to **NullableString** | 接口版本 | [optional] 

## Methods

### NewReturnPageNotifyRequest

`func NewReturnPageNotifyRequest() *ReturnPageNotifyRequest`

NewReturnPageNotifyRequest instantiates a new ReturnPageNotifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnPageNotifyRequestWithDefaults

`func NewReturnPageNotifyRequestWithDefaults() *ReturnPageNotifyRequest`

NewReturnPageNotifyRequestWithDefaults instantiates a new ReturnPageNotifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *ReturnPageNotifyRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ReturnPageNotifyRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ReturnPageNotifyRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ReturnPageNotifyRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *ReturnPageNotifyRequest) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *ReturnPageNotifyRequest) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetAuthAppId

`func (o *ReturnPageNotifyRequest) GetAuthAppId() string`

GetAuthAppId returns the AuthAppId field if non-nil, zero value otherwise.

### GetAuthAppIdOk

`func (o *ReturnPageNotifyRequest) GetAuthAppIdOk() (*string, bool)`

GetAuthAppIdOk returns a tuple with the AuthAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthAppId

`func (o *ReturnPageNotifyRequest) SetAuthAppId(v string)`

SetAuthAppId sets AuthAppId field to given value.

### HasAuthAppId

`func (o *ReturnPageNotifyRequest) HasAuthAppId() bool`

HasAuthAppId returns a boolean if a field has been set.

### SetAuthAppIdNil

`func (o *ReturnPageNotifyRequest) SetAuthAppIdNil(b bool)`

 SetAuthAppIdNil sets the value for AuthAppId to be an explicit nil

### UnsetAuthAppId
`func (o *ReturnPageNotifyRequest) UnsetAuthAppId()`

UnsetAuthAppId ensures that no value is present for AuthAppId, not even an explicit nil
### GetCharset

`func (o *ReturnPageNotifyRequest) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *ReturnPageNotifyRequest) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *ReturnPageNotifyRequest) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *ReturnPageNotifyRequest) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### SetCharsetNil

`func (o *ReturnPageNotifyRequest) SetCharsetNil(b bool)`

 SetCharsetNil sets the value for Charset to be an explicit nil

### UnsetCharset
`func (o *ReturnPageNotifyRequest) UnsetCharset()`

UnsetCharset ensures that no value is present for Charset, not even an explicit nil
### GetMethod

`func (o *ReturnPageNotifyRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ReturnPageNotifyRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ReturnPageNotifyRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ReturnPageNotifyRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ReturnPageNotifyRequest) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ReturnPageNotifyRequest) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetOutTradeNo

`func (o *ReturnPageNotifyRequest) GetOutTradeNo() string`

GetOutTradeNo returns the OutTradeNo field if non-nil, zero value otherwise.

### GetOutTradeNoOk

`func (o *ReturnPageNotifyRequest) GetOutTradeNoOk() (*string, bool)`

GetOutTradeNoOk returns a tuple with the OutTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTradeNo

`func (o *ReturnPageNotifyRequest) SetOutTradeNo(v string)`

SetOutTradeNo sets OutTradeNo field to given value.

### HasOutTradeNo

`func (o *ReturnPageNotifyRequest) HasOutTradeNo() bool`

HasOutTradeNo returns a boolean if a field has been set.

### SetOutTradeNoNil

`func (o *ReturnPageNotifyRequest) SetOutTradeNoNil(b bool)`

 SetOutTradeNoNil sets the value for OutTradeNo to be an explicit nil

### UnsetOutTradeNo
`func (o *ReturnPageNotifyRequest) UnsetOutTradeNo()`

UnsetOutTradeNo ensures that no value is present for OutTradeNo, not even an explicit nil
### GetSellerId

`func (o *ReturnPageNotifyRequest) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *ReturnPageNotifyRequest) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *ReturnPageNotifyRequest) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.

### HasSellerId

`func (o *ReturnPageNotifyRequest) HasSellerId() bool`

HasSellerId returns a boolean if a field has been set.

### SetSellerIdNil

`func (o *ReturnPageNotifyRequest) SetSellerIdNil(b bool)`

 SetSellerIdNil sets the value for SellerId to be an explicit nil

### UnsetSellerId
`func (o *ReturnPageNotifyRequest) UnsetSellerId()`

UnsetSellerId ensures that no value is present for SellerId, not even an explicit nil
### GetSign

`func (o *ReturnPageNotifyRequest) GetSign() string`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *ReturnPageNotifyRequest) GetSignOk() (*string, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *ReturnPageNotifyRequest) SetSign(v string)`

SetSign sets Sign field to given value.

### HasSign

`func (o *ReturnPageNotifyRequest) HasSign() bool`

HasSign returns a boolean if a field has been set.

### SetSignNil

`func (o *ReturnPageNotifyRequest) SetSignNil(b bool)`

 SetSignNil sets the value for Sign to be an explicit nil

### UnsetSign
`func (o *ReturnPageNotifyRequest) UnsetSign()`

UnsetSign ensures that no value is present for Sign, not even an explicit nil
### GetSignType

`func (o *ReturnPageNotifyRequest) GetSignType() string`

GetSignType returns the SignType field if non-nil, zero value otherwise.

### GetSignTypeOk

`func (o *ReturnPageNotifyRequest) GetSignTypeOk() (*string, bool)`

GetSignTypeOk returns a tuple with the SignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignType

`func (o *ReturnPageNotifyRequest) SetSignType(v string)`

SetSignType sets SignType field to given value.

### HasSignType

`func (o *ReturnPageNotifyRequest) HasSignType() bool`

HasSignType returns a boolean if a field has been set.

### SetSignTypeNil

`func (o *ReturnPageNotifyRequest) SetSignTypeNil(b bool)`

 SetSignTypeNil sets the value for SignType to be an explicit nil

### UnsetSignType
`func (o *ReturnPageNotifyRequest) UnsetSignType()`

UnsetSignType ensures that no value is present for SignType, not even an explicit nil
### GetTimestamp

`func (o *ReturnPageNotifyRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReturnPageNotifyRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReturnPageNotifyRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ReturnPageNotifyRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ReturnPageNotifyRequest) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ReturnPageNotifyRequest) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTotalAmount

`func (o *ReturnPageNotifyRequest) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ReturnPageNotifyRequest) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ReturnPageNotifyRequest) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ReturnPageNotifyRequest) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *ReturnPageNotifyRequest) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *ReturnPageNotifyRequest) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetTradeNo

`func (o *ReturnPageNotifyRequest) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ReturnPageNotifyRequest) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ReturnPageNotifyRequest) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *ReturnPageNotifyRequest) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### SetTradeNoNil

`func (o *ReturnPageNotifyRequest) SetTradeNoNil(b bool)`

 SetTradeNoNil sets the value for TradeNo to be an explicit nil

### UnsetTradeNo
`func (o *ReturnPageNotifyRequest) UnsetTradeNo()`

UnsetTradeNo ensures that no value is present for TradeNo, not even an explicit nil
### GetVersion

`func (o *ReturnPageNotifyRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ReturnPageNotifyRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ReturnPageNotifyRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ReturnPageNotifyRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ReturnPageNotifyRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ReturnPageNotifyRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


