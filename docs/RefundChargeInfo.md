# RefundChargeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **NullableString** |  | [optional] 
**RefundChargeFee** | Pointer to **NullableString** |  | [optional] 
**RefundSubFeeDetailList** | Pointer to [**[]RefundSubFee**](RefundSubFee.md) |  | [optional] 
**SwitchFeeRate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRefundChargeInfo

`func NewRefundChargeInfo() *RefundChargeInfo`

NewRefundChargeInfo instantiates a new RefundChargeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundChargeInfoWithDefaults

`func NewRefundChargeInfoWithDefaults() *RefundChargeInfo`

NewRefundChargeInfoWithDefaults instantiates a new RefundChargeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *RefundChargeInfo) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *RefundChargeInfo) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *RefundChargeInfo) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *RefundChargeInfo) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### SetChargeTypeNil

`func (o *RefundChargeInfo) SetChargeTypeNil(b bool)`

 SetChargeTypeNil sets the value for ChargeType to be an explicit nil

### UnsetChargeType
`func (o *RefundChargeInfo) UnsetChargeType()`

UnsetChargeType ensures that no value is present for ChargeType, not even an explicit nil
### GetRefundChargeFee

`func (o *RefundChargeInfo) GetRefundChargeFee() string`

GetRefundChargeFee returns the RefundChargeFee field if non-nil, zero value otherwise.

### GetRefundChargeFeeOk

`func (o *RefundChargeInfo) GetRefundChargeFeeOk() (*string, bool)`

GetRefundChargeFeeOk returns a tuple with the RefundChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundChargeFee

`func (o *RefundChargeInfo) SetRefundChargeFee(v string)`

SetRefundChargeFee sets RefundChargeFee field to given value.

### HasRefundChargeFee

`func (o *RefundChargeInfo) HasRefundChargeFee() bool`

HasRefundChargeFee returns a boolean if a field has been set.

### SetRefundChargeFeeNil

`func (o *RefundChargeInfo) SetRefundChargeFeeNil(b bool)`

 SetRefundChargeFeeNil sets the value for RefundChargeFee to be an explicit nil

### UnsetRefundChargeFee
`func (o *RefundChargeInfo) UnsetRefundChargeFee()`

UnsetRefundChargeFee ensures that no value is present for RefundChargeFee, not even an explicit nil
### GetRefundSubFeeDetailList

`func (o *RefundChargeInfo) GetRefundSubFeeDetailList() []RefundSubFee`

GetRefundSubFeeDetailList returns the RefundSubFeeDetailList field if non-nil, zero value otherwise.

### GetRefundSubFeeDetailListOk

`func (o *RefundChargeInfo) GetRefundSubFeeDetailListOk() (*[]RefundSubFee, bool)`

GetRefundSubFeeDetailListOk returns a tuple with the RefundSubFeeDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSubFeeDetailList

`func (o *RefundChargeInfo) SetRefundSubFeeDetailList(v []RefundSubFee)`

SetRefundSubFeeDetailList sets RefundSubFeeDetailList field to given value.

### HasRefundSubFeeDetailList

`func (o *RefundChargeInfo) HasRefundSubFeeDetailList() bool`

HasRefundSubFeeDetailList returns a boolean if a field has been set.

### SetRefundSubFeeDetailListNil

`func (o *RefundChargeInfo) SetRefundSubFeeDetailListNil(b bool)`

 SetRefundSubFeeDetailListNil sets the value for RefundSubFeeDetailList to be an explicit nil

### UnsetRefundSubFeeDetailList
`func (o *RefundChargeInfo) UnsetRefundSubFeeDetailList()`

UnsetRefundSubFeeDetailList ensures that no value is present for RefundSubFeeDetailList, not even an explicit nil
### GetSwitchFeeRate

`func (o *RefundChargeInfo) GetSwitchFeeRate() string`

GetSwitchFeeRate returns the SwitchFeeRate field if non-nil, zero value otherwise.

### GetSwitchFeeRateOk

`func (o *RefundChargeInfo) GetSwitchFeeRateOk() (*string, bool)`

GetSwitchFeeRateOk returns a tuple with the SwitchFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchFeeRate

`func (o *RefundChargeInfo) SetSwitchFeeRate(v string)`

SetSwitchFeeRate sets SwitchFeeRate field to given value.

### HasSwitchFeeRate

`func (o *RefundChargeInfo) HasSwitchFeeRate() bool`

HasSwitchFeeRate returns a boolean if a field has been set.

### SetSwitchFeeRateNil

`func (o *RefundChargeInfo) SetSwitchFeeRateNil(b bool)`

 SetSwitchFeeRateNil sets the value for SwitchFeeRate to be an explicit nil

### UnsetSwitchFeeRate
`func (o *RefundChargeInfo) UnsetSwitchFeeRate()`

UnsetSwitchFeeRate ensures that no value is present for SwitchFeeRate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


