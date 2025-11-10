# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 订单的唯一标识符。 | [optional] 
**UserID** | Pointer to **int64** | 创建订单的用户ID。 | [optional] 
**PayType** | Pointer to **NullableString** | 订单的支付类型，例如 &#39;信用卡&#39;, &#39;支付宝&#39;, &#39;微信支付&#39; 等。 | [optional] 
**Amount** | Pointer to **float64** | 订单的总金额。 | [optional] 
**OrderNo** | Pointer to **NullableString** | 订单的唯一编号，通常由系统生成。 | [optional] 
**TradeNo** | Pointer to **NullableString** | 与订单关联的交易编号，通常由支付平台提供。 | [optional] 
**Status** | Pointer to **NullableString** | 订单的当前状态，例如 &#39;待支付&#39;, &#39;已完成&#39;, &#39;已取消&#39; 等。 | [optional] 
**ProductType** | Pointer to **NullableString** | 订单中商品的类型分类。 | [optional] 
**ProductID** | Pointer to **NullableString** | 订单中商品的唯一标识符。 | [optional] 
**ProductName** | Pointer to **NullableString** | 订单中商品的名称。 | [optional] 
**AllowRefund** | Pointer to **bool** | 指示订单是否允许进行退款操作。 | [optional] 
**AllowRefundUntil** | Pointer to **time.Time** | 订单允许进行退款操作的截止时间。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记订单的标签。 | [optional] 
**Remark** | Pointer to **NullableString** | 订单的额外备注信息。 | [optional] 
**Description** | Pointer to **NullableString** | 订单的详细描述信息。 | [optional] 
**OrderPayTime** | Pointer to **time.Time** | 订单完成支付的时间。 | [optional] 
**ExpireTime** | Pointer to **time.Time** | 订单的过期时间，超过该时间订单将自动取消。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 订单的创建时间，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 订单的最后更新时间，默认为当前时间。 | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *Order) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Order) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Order) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *Order) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetPayType

`func (o *Order) GetPayType() string`

GetPayType returns the PayType field if non-nil, zero value otherwise.

### GetPayTypeOk

`func (o *Order) GetPayTypeOk() (*string, bool)`

GetPayTypeOk returns a tuple with the PayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayType

`func (o *Order) SetPayType(v string)`

SetPayType sets PayType field to given value.

### HasPayType

`func (o *Order) HasPayType() bool`

HasPayType returns a boolean if a field has been set.

### SetPayTypeNil

`func (o *Order) SetPayTypeNil(b bool)`

 SetPayTypeNil sets the value for PayType to be an explicit nil

### UnsetPayType
`func (o *Order) UnsetPayType()`

UnsetPayType ensures that no value is present for PayType, not even an explicit nil
### GetAmount

`func (o *Order) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Order) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Order) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Order) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetOrderNo

`func (o *Order) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *Order) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *Order) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.

### HasOrderNo

`func (o *Order) HasOrderNo() bool`

HasOrderNo returns a boolean if a field has been set.

### SetOrderNoNil

`func (o *Order) SetOrderNoNil(b bool)`

 SetOrderNoNil sets the value for OrderNo to be an explicit nil

### UnsetOrderNo
`func (o *Order) UnsetOrderNo()`

UnsetOrderNo ensures that no value is present for OrderNo, not even an explicit nil
### GetTradeNo

`func (o *Order) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *Order) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *Order) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *Order) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### SetTradeNoNil

`func (o *Order) SetTradeNoNil(b bool)`

 SetTradeNoNil sets the value for TradeNo to be an explicit nil

### UnsetTradeNo
`func (o *Order) UnsetTradeNo()`

UnsetTradeNo ensures that no value is present for TradeNo, not even an explicit nil
### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Order) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Order) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProductType

`func (o *Order) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Order) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Order) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *Order) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *Order) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *Order) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetProductID

`func (o *Order) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *Order) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *Order) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *Order) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### SetProductIDNil

`func (o *Order) SetProductIDNil(b bool)`

 SetProductIDNil sets the value for ProductID to be an explicit nil

### UnsetProductID
`func (o *Order) UnsetProductID()`

UnsetProductID ensures that no value is present for ProductID, not even an explicit nil
### GetProductName

`func (o *Order) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Order) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Order) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *Order) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *Order) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *Order) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetAllowRefund

`func (o *Order) GetAllowRefund() bool`

GetAllowRefund returns the AllowRefund field if non-nil, zero value otherwise.

### GetAllowRefundOk

`func (o *Order) GetAllowRefundOk() (*bool, bool)`

GetAllowRefundOk returns a tuple with the AllowRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRefund

`func (o *Order) SetAllowRefund(v bool)`

SetAllowRefund sets AllowRefund field to given value.

### HasAllowRefund

`func (o *Order) HasAllowRefund() bool`

HasAllowRefund returns a boolean if a field has been set.

### GetAllowRefundUntil

`func (o *Order) GetAllowRefundUntil() time.Time`

GetAllowRefundUntil returns the AllowRefundUntil field if non-nil, zero value otherwise.

### GetAllowRefundUntilOk

`func (o *Order) GetAllowRefundUntilOk() (*time.Time, bool)`

GetAllowRefundUntilOk returns a tuple with the AllowRefundUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRefundUntil

`func (o *Order) SetAllowRefundUntil(v time.Time)`

SetAllowRefundUntil sets AllowRefundUntil field to given value.

### HasAllowRefundUntil

`func (o *Order) HasAllowRefundUntil() bool`

HasAllowRefundUntil returns a boolean if a field has been set.

### GetTags

`func (o *Order) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Order) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Order) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Order) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Order) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Order) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetRemark

`func (o *Order) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *Order) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *Order) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *Order) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *Order) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *Order) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetDescription

`func (o *Order) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Order) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Order) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Order) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Order) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Order) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrderPayTime

`func (o *Order) GetOrderPayTime() time.Time`

GetOrderPayTime returns the OrderPayTime field if non-nil, zero value otherwise.

### GetOrderPayTimeOk

`func (o *Order) GetOrderPayTimeOk() (*time.Time, bool)`

GetOrderPayTimeOk returns a tuple with the OrderPayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderPayTime

`func (o *Order) SetOrderPayTime(v time.Time)`

SetOrderPayTime sets OrderPayTime field to given value.

### HasOrderPayTime

`func (o *Order) HasOrderPayTime() bool`

HasOrderPayTime returns a boolean if a field has been set.

### GetExpireTime

`func (o *Order) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *Order) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *Order) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *Order) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetCreateDate

`func (o *Order) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Order) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Order) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Order) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *Order) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Order) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Order) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Order) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


