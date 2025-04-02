# TradeSettleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeSettleDetailList** | Pointer to [**[]TradeSettleDetail**](TradeSettleDetail.md) |  | [optional] 
**TradeUnsettledAmount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTradeSettleInfo

`func NewTradeSettleInfo() *TradeSettleInfo`

NewTradeSettleInfo instantiates a new TradeSettleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeSettleInfoWithDefaults

`func NewTradeSettleInfoWithDefaults() *TradeSettleInfo`

NewTradeSettleInfoWithDefaults instantiates a new TradeSettleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeSettleDetailList

`func (o *TradeSettleInfo) GetTradeSettleDetailList() []TradeSettleDetail`

GetTradeSettleDetailList returns the TradeSettleDetailList field if non-nil, zero value otherwise.

### GetTradeSettleDetailListOk

`func (o *TradeSettleInfo) GetTradeSettleDetailListOk() (*[]TradeSettleDetail, bool)`

GetTradeSettleDetailListOk returns a tuple with the TradeSettleDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSettleDetailList

`func (o *TradeSettleInfo) SetTradeSettleDetailList(v []TradeSettleDetail)`

SetTradeSettleDetailList sets TradeSettleDetailList field to given value.

### HasTradeSettleDetailList

`func (o *TradeSettleInfo) HasTradeSettleDetailList() bool`

HasTradeSettleDetailList returns a boolean if a field has been set.

### SetTradeSettleDetailListNil

`func (o *TradeSettleInfo) SetTradeSettleDetailListNil(b bool)`

 SetTradeSettleDetailListNil sets the value for TradeSettleDetailList to be an explicit nil

### UnsetTradeSettleDetailList
`func (o *TradeSettleInfo) UnsetTradeSettleDetailList()`

UnsetTradeSettleDetailList ensures that no value is present for TradeSettleDetailList, not even an explicit nil
### GetTradeUnsettledAmount

`func (o *TradeSettleInfo) GetTradeUnsettledAmount() string`

GetTradeUnsettledAmount returns the TradeUnsettledAmount field if non-nil, zero value otherwise.

### GetTradeUnsettledAmountOk

`func (o *TradeSettleInfo) GetTradeUnsettledAmountOk() (*string, bool)`

GetTradeUnsettledAmountOk returns a tuple with the TradeUnsettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeUnsettledAmount

`func (o *TradeSettleInfo) SetTradeUnsettledAmount(v string)`

SetTradeUnsettledAmount sets TradeUnsettledAmount field to given value.

### HasTradeUnsettledAmount

`func (o *TradeSettleInfo) HasTradeUnsettledAmount() bool`

HasTradeUnsettledAmount returns a boolean if a field has been set.

### SetTradeUnsettledAmountNil

`func (o *TradeSettleInfo) SetTradeUnsettledAmountNil(b bool)`

 SetTradeUnsettledAmountNil sets the value for TradeUnsettledAmount to be an explicit nil

### UnsetTradeUnsettledAmount
`func (o *TradeSettleInfo) UnsetTradeUnsettledAmount()`

UnsetTradeUnsettledAmount ensures that no value is present for TradeUnsettledAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


