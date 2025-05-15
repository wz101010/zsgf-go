# AlipayCreateOrderWapPayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderNo** | **string** | 自定义订单号 | 
**Amount** | **float64** | 订单金额。例如：0.01 | 
**Subject** | **string** | 商品名称 | 
**ReturnUrl** | Pointer to **NullableString** | 支付完成后返回的URL地址 | [optional] 

## Methods

### NewAlipayCreateOrderWapPayRequest

`func NewAlipayCreateOrderWapPayRequest(orderNo string, amount float64, subject string, ) *AlipayCreateOrderWapPayRequest`

NewAlipayCreateOrderWapPayRequest instantiates a new AlipayCreateOrderWapPayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlipayCreateOrderWapPayRequestWithDefaults

`func NewAlipayCreateOrderWapPayRequestWithDefaults() *AlipayCreateOrderWapPayRequest`

NewAlipayCreateOrderWapPayRequestWithDefaults instantiates a new AlipayCreateOrderWapPayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNo

`func (o *AlipayCreateOrderWapPayRequest) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *AlipayCreateOrderWapPayRequest) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *AlipayCreateOrderWapPayRequest) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.


### GetAmount

`func (o *AlipayCreateOrderWapPayRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AlipayCreateOrderWapPayRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AlipayCreateOrderWapPayRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetSubject

`func (o *AlipayCreateOrderWapPayRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlipayCreateOrderWapPayRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlipayCreateOrderWapPayRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetReturnUrl

`func (o *AlipayCreateOrderWapPayRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *AlipayCreateOrderWapPayRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *AlipayCreateOrderWapPayRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *AlipayCreateOrderWapPayRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *AlipayCreateOrderWapPayRequest) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *AlipayCreateOrderWapPayRequest) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


