# ChargeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeFee** | Pointer to **NullableString** |  | [optional] 
**ChargeType** | Pointer to **NullableString** |  | [optional] 
**IsRatingOnSwitch** | Pointer to **NullableString** |  | [optional] 
**IsRatingOnTradeReceiver** | Pointer to **NullableString** |  | [optional] 
**OriginalChargeFee** | Pointer to **NullableString** |  | [optional] 
**SubFeeDetailList** | Pointer to [**[]SubFee**](SubFee.md) |  | [optional] 
**SwitchFeeRate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChargeInfo

`func NewChargeInfo() *ChargeInfo`

NewChargeInfo instantiates a new ChargeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeInfoWithDefaults

`func NewChargeInfoWithDefaults() *ChargeInfo`

NewChargeInfoWithDefaults instantiates a new ChargeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeFee

`func (o *ChargeInfo) GetChargeFee() string`

GetChargeFee returns the ChargeFee field if non-nil, zero value otherwise.

### GetChargeFeeOk

`func (o *ChargeInfo) GetChargeFeeOk() (*string, bool)`

GetChargeFeeOk returns a tuple with the ChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeFee

`func (o *ChargeInfo) SetChargeFee(v string)`

SetChargeFee sets ChargeFee field to given value.

### HasChargeFee

`func (o *ChargeInfo) HasChargeFee() bool`

HasChargeFee returns a boolean if a field has been set.

### SetChargeFeeNil

`func (o *ChargeInfo) SetChargeFeeNil(b bool)`

 SetChargeFeeNil sets the value for ChargeFee to be an explicit nil

### UnsetChargeFee
`func (o *ChargeInfo) UnsetChargeFee()`

UnsetChargeFee ensures that no value is present for ChargeFee, not even an explicit nil
### GetChargeType

`func (o *ChargeInfo) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *ChargeInfo) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *ChargeInfo) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *ChargeInfo) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### SetChargeTypeNil

`func (o *ChargeInfo) SetChargeTypeNil(b bool)`

 SetChargeTypeNil sets the value for ChargeType to be an explicit nil

### UnsetChargeType
`func (o *ChargeInfo) UnsetChargeType()`

UnsetChargeType ensures that no value is present for ChargeType, not even an explicit nil
### GetIsRatingOnSwitch

`func (o *ChargeInfo) GetIsRatingOnSwitch() string`

GetIsRatingOnSwitch returns the IsRatingOnSwitch field if non-nil, zero value otherwise.

### GetIsRatingOnSwitchOk

`func (o *ChargeInfo) GetIsRatingOnSwitchOk() (*string, bool)`

GetIsRatingOnSwitchOk returns a tuple with the IsRatingOnSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRatingOnSwitch

`func (o *ChargeInfo) SetIsRatingOnSwitch(v string)`

SetIsRatingOnSwitch sets IsRatingOnSwitch field to given value.

### HasIsRatingOnSwitch

`func (o *ChargeInfo) HasIsRatingOnSwitch() bool`

HasIsRatingOnSwitch returns a boolean if a field has been set.

### SetIsRatingOnSwitchNil

`func (o *ChargeInfo) SetIsRatingOnSwitchNil(b bool)`

 SetIsRatingOnSwitchNil sets the value for IsRatingOnSwitch to be an explicit nil

### UnsetIsRatingOnSwitch
`func (o *ChargeInfo) UnsetIsRatingOnSwitch()`

UnsetIsRatingOnSwitch ensures that no value is present for IsRatingOnSwitch, not even an explicit nil
### GetIsRatingOnTradeReceiver

`func (o *ChargeInfo) GetIsRatingOnTradeReceiver() string`

GetIsRatingOnTradeReceiver returns the IsRatingOnTradeReceiver field if non-nil, zero value otherwise.

### GetIsRatingOnTradeReceiverOk

`func (o *ChargeInfo) GetIsRatingOnTradeReceiverOk() (*string, bool)`

GetIsRatingOnTradeReceiverOk returns a tuple with the IsRatingOnTradeReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRatingOnTradeReceiver

`func (o *ChargeInfo) SetIsRatingOnTradeReceiver(v string)`

SetIsRatingOnTradeReceiver sets IsRatingOnTradeReceiver field to given value.

### HasIsRatingOnTradeReceiver

`func (o *ChargeInfo) HasIsRatingOnTradeReceiver() bool`

HasIsRatingOnTradeReceiver returns a boolean if a field has been set.

### SetIsRatingOnTradeReceiverNil

`func (o *ChargeInfo) SetIsRatingOnTradeReceiverNil(b bool)`

 SetIsRatingOnTradeReceiverNil sets the value for IsRatingOnTradeReceiver to be an explicit nil

### UnsetIsRatingOnTradeReceiver
`func (o *ChargeInfo) UnsetIsRatingOnTradeReceiver()`

UnsetIsRatingOnTradeReceiver ensures that no value is present for IsRatingOnTradeReceiver, not even an explicit nil
### GetOriginalChargeFee

`func (o *ChargeInfo) GetOriginalChargeFee() string`

GetOriginalChargeFee returns the OriginalChargeFee field if non-nil, zero value otherwise.

### GetOriginalChargeFeeOk

`func (o *ChargeInfo) GetOriginalChargeFeeOk() (*string, bool)`

GetOriginalChargeFeeOk returns a tuple with the OriginalChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalChargeFee

`func (o *ChargeInfo) SetOriginalChargeFee(v string)`

SetOriginalChargeFee sets OriginalChargeFee field to given value.

### HasOriginalChargeFee

`func (o *ChargeInfo) HasOriginalChargeFee() bool`

HasOriginalChargeFee returns a boolean if a field has been set.

### SetOriginalChargeFeeNil

`func (o *ChargeInfo) SetOriginalChargeFeeNil(b bool)`

 SetOriginalChargeFeeNil sets the value for OriginalChargeFee to be an explicit nil

### UnsetOriginalChargeFee
`func (o *ChargeInfo) UnsetOriginalChargeFee()`

UnsetOriginalChargeFee ensures that no value is present for OriginalChargeFee, not even an explicit nil
### GetSubFeeDetailList

`func (o *ChargeInfo) GetSubFeeDetailList() []SubFee`

GetSubFeeDetailList returns the SubFeeDetailList field if non-nil, zero value otherwise.

### GetSubFeeDetailListOk

`func (o *ChargeInfo) GetSubFeeDetailListOk() (*[]SubFee, bool)`

GetSubFeeDetailListOk returns a tuple with the SubFeeDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubFeeDetailList

`func (o *ChargeInfo) SetSubFeeDetailList(v []SubFee)`

SetSubFeeDetailList sets SubFeeDetailList field to given value.

### HasSubFeeDetailList

`func (o *ChargeInfo) HasSubFeeDetailList() bool`

HasSubFeeDetailList returns a boolean if a field has been set.

### SetSubFeeDetailListNil

`func (o *ChargeInfo) SetSubFeeDetailListNil(b bool)`

 SetSubFeeDetailListNil sets the value for SubFeeDetailList to be an explicit nil

### UnsetSubFeeDetailList
`func (o *ChargeInfo) UnsetSubFeeDetailList()`

UnsetSubFeeDetailList ensures that no value is present for SubFeeDetailList, not even an explicit nil
### GetSwitchFeeRate

`func (o *ChargeInfo) GetSwitchFeeRate() string`

GetSwitchFeeRate returns the SwitchFeeRate field if non-nil, zero value otherwise.

### GetSwitchFeeRateOk

`func (o *ChargeInfo) GetSwitchFeeRateOk() (*string, bool)`

GetSwitchFeeRateOk returns a tuple with the SwitchFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchFeeRate

`func (o *ChargeInfo) SetSwitchFeeRate(v string)`

SetSwitchFeeRate sets SwitchFeeRate field to given value.

### HasSwitchFeeRate

`func (o *ChargeInfo) HasSwitchFeeRate() bool`

HasSwitchFeeRate returns a boolean if a field has been set.

### SetSwitchFeeRateNil

`func (o *ChargeInfo) SetSwitchFeeRateNil(b bool)`

 SetSwitchFeeRateNil sets the value for SwitchFeeRate to be an explicit nil

### UnsetSwitchFeeRate
`func (o *ChargeInfo) UnsetSwitchFeeRate()`

UnsetSwitchFeeRate ensures that no value is present for SwitchFeeRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


