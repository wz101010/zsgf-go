# AlipayTradeRefundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 
**SubCode** | Pointer to **NullableString** |  | [optional] 
**SubMsg** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] [readonly] 
**BuyerLogonId** | Pointer to **NullableString** |  | [optional] 
**BuyerOpenId** | Pointer to **NullableString** |  | [optional] 
**BuyerUserId** | Pointer to **NullableString** |  | [optional] 
**FundChange** | Pointer to **NullableString** |  | [optional] 
**GmtRefundPay** | Pointer to **NullableString** |  | [optional] 
**HasDepositBack** | Pointer to **NullableString** |  | [optional] 
**OpenId** | Pointer to **NullableString** |  | [optional] 
**OutTradeNo** | Pointer to **NullableString** |  | [optional] 
**PreAuthCancelFee** | Pointer to **NullableString** |  | [optional] 
**PresentRefundBuyerAmount** | Pointer to **NullableString** |  | [optional] 
**PresentRefundDiscountAmount** | Pointer to **NullableString** |  | [optional] 
**PresentRefundMdiscountAmount** | Pointer to **NullableString** |  | [optional] 
**RefundChargeAmount** | Pointer to **NullableString** |  | [optional] 
**RefundChargeInfoList** | Pointer to [**[]RefundChargeInfo**](RefundChargeInfo.md) |  | [optional] 
**RefundCurrency** | Pointer to **NullableString** |  | [optional] 
**RefundDetailItemList** | Pointer to [**[]TradeFundBill**](TradeFundBill.md) |  | [optional] 
**RefundFee** | Pointer to **NullableString** |  | [optional] 
**RefundHybAmount** | Pointer to **NullableString** |  | [optional] 
**RefundPresetPaytoolList** | Pointer to [**PresetPayToolInfo**](PresetPayToolInfo.md) |  | [optional] 
**RefundSettlementId** | Pointer to **NullableString** |  | [optional] 
**RefundVoucherDetailList** | Pointer to [**[]VoucherDetail**](VoucherDetail.md) |  | [optional] 
**SendBackFee** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **NullableString** |  | [optional] 
**TradeNo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlipayTradeRefundResponse

`func NewAlipayTradeRefundResponse() *AlipayTradeRefundResponse`

NewAlipayTradeRefundResponse instantiates a new AlipayTradeRefundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlipayTradeRefundResponseWithDefaults

`func NewAlipayTradeRefundResponseWithDefaults() *AlipayTradeRefundResponse`

NewAlipayTradeRefundResponseWithDefaults instantiates a new AlipayTradeRefundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AlipayTradeRefundResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AlipayTradeRefundResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AlipayTradeRefundResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AlipayTradeRefundResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AlipayTradeRefundResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AlipayTradeRefundResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMsg

`func (o *AlipayTradeRefundResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AlipayTradeRefundResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AlipayTradeRefundResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AlipayTradeRefundResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *AlipayTradeRefundResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *AlipayTradeRefundResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetSubCode

`func (o *AlipayTradeRefundResponse) GetSubCode() string`

GetSubCode returns the SubCode field if non-nil, zero value otherwise.

### GetSubCodeOk

`func (o *AlipayTradeRefundResponse) GetSubCodeOk() (*string, bool)`

GetSubCodeOk returns a tuple with the SubCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCode

`func (o *AlipayTradeRefundResponse) SetSubCode(v string)`

SetSubCode sets SubCode field to given value.

### HasSubCode

`func (o *AlipayTradeRefundResponse) HasSubCode() bool`

HasSubCode returns a boolean if a field has been set.

### SetSubCodeNil

`func (o *AlipayTradeRefundResponse) SetSubCodeNil(b bool)`

 SetSubCodeNil sets the value for SubCode to be an explicit nil

### UnsetSubCode
`func (o *AlipayTradeRefundResponse) UnsetSubCode()`

UnsetSubCode ensures that no value is present for SubCode, not even an explicit nil
### GetSubMsg

`func (o *AlipayTradeRefundResponse) GetSubMsg() string`

GetSubMsg returns the SubMsg field if non-nil, zero value otherwise.

### GetSubMsgOk

`func (o *AlipayTradeRefundResponse) GetSubMsgOk() (*string, bool)`

GetSubMsgOk returns a tuple with the SubMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMsg

`func (o *AlipayTradeRefundResponse) SetSubMsg(v string)`

SetSubMsg sets SubMsg field to given value.

### HasSubMsg

`func (o *AlipayTradeRefundResponse) HasSubMsg() bool`

HasSubMsg returns a boolean if a field has been set.

### SetSubMsgNil

`func (o *AlipayTradeRefundResponse) SetSubMsgNil(b bool)`

 SetSubMsgNil sets the value for SubMsg to be an explicit nil

### UnsetSubMsg
`func (o *AlipayTradeRefundResponse) UnsetSubMsg()`

UnsetSubMsg ensures that no value is present for SubMsg, not even an explicit nil
### GetBody

`func (o *AlipayTradeRefundResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AlipayTradeRefundResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AlipayTradeRefundResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AlipayTradeRefundResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *AlipayTradeRefundResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *AlipayTradeRefundResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetIsError

`func (o *AlipayTradeRefundResponse) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AlipayTradeRefundResponse) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AlipayTradeRefundResponse) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *AlipayTradeRefundResponse) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetBuyerLogonId

`func (o *AlipayTradeRefundResponse) GetBuyerLogonId() string`

GetBuyerLogonId returns the BuyerLogonId field if non-nil, zero value otherwise.

### GetBuyerLogonIdOk

`func (o *AlipayTradeRefundResponse) GetBuyerLogonIdOk() (*string, bool)`

GetBuyerLogonIdOk returns a tuple with the BuyerLogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerLogonId

`func (o *AlipayTradeRefundResponse) SetBuyerLogonId(v string)`

SetBuyerLogonId sets BuyerLogonId field to given value.

### HasBuyerLogonId

`func (o *AlipayTradeRefundResponse) HasBuyerLogonId() bool`

HasBuyerLogonId returns a boolean if a field has been set.

### SetBuyerLogonIdNil

`func (o *AlipayTradeRefundResponse) SetBuyerLogonIdNil(b bool)`

 SetBuyerLogonIdNil sets the value for BuyerLogonId to be an explicit nil

### UnsetBuyerLogonId
`func (o *AlipayTradeRefundResponse) UnsetBuyerLogonId()`

UnsetBuyerLogonId ensures that no value is present for BuyerLogonId, not even an explicit nil
### GetBuyerOpenId

`func (o *AlipayTradeRefundResponse) GetBuyerOpenId() string`

GetBuyerOpenId returns the BuyerOpenId field if non-nil, zero value otherwise.

### GetBuyerOpenIdOk

`func (o *AlipayTradeRefundResponse) GetBuyerOpenIdOk() (*string, bool)`

GetBuyerOpenIdOk returns a tuple with the BuyerOpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerOpenId

`func (o *AlipayTradeRefundResponse) SetBuyerOpenId(v string)`

SetBuyerOpenId sets BuyerOpenId field to given value.

### HasBuyerOpenId

`func (o *AlipayTradeRefundResponse) HasBuyerOpenId() bool`

HasBuyerOpenId returns a boolean if a field has been set.

### SetBuyerOpenIdNil

`func (o *AlipayTradeRefundResponse) SetBuyerOpenIdNil(b bool)`

 SetBuyerOpenIdNil sets the value for BuyerOpenId to be an explicit nil

### UnsetBuyerOpenId
`func (o *AlipayTradeRefundResponse) UnsetBuyerOpenId()`

UnsetBuyerOpenId ensures that no value is present for BuyerOpenId, not even an explicit nil
### GetBuyerUserId

`func (o *AlipayTradeRefundResponse) GetBuyerUserId() string`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *AlipayTradeRefundResponse) GetBuyerUserIdOk() (*string, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *AlipayTradeRefundResponse) SetBuyerUserId(v string)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *AlipayTradeRefundResponse) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### SetBuyerUserIdNil

`func (o *AlipayTradeRefundResponse) SetBuyerUserIdNil(b bool)`

 SetBuyerUserIdNil sets the value for BuyerUserId to be an explicit nil

### UnsetBuyerUserId
`func (o *AlipayTradeRefundResponse) UnsetBuyerUserId()`

UnsetBuyerUserId ensures that no value is present for BuyerUserId, not even an explicit nil
### GetFundChange

`func (o *AlipayTradeRefundResponse) GetFundChange() string`

GetFundChange returns the FundChange field if non-nil, zero value otherwise.

### GetFundChangeOk

`func (o *AlipayTradeRefundResponse) GetFundChangeOk() (*string, bool)`

GetFundChangeOk returns a tuple with the FundChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundChange

`func (o *AlipayTradeRefundResponse) SetFundChange(v string)`

SetFundChange sets FundChange field to given value.

### HasFundChange

`func (o *AlipayTradeRefundResponse) HasFundChange() bool`

HasFundChange returns a boolean if a field has been set.

### SetFundChangeNil

`func (o *AlipayTradeRefundResponse) SetFundChangeNil(b bool)`

 SetFundChangeNil sets the value for FundChange to be an explicit nil

### UnsetFundChange
`func (o *AlipayTradeRefundResponse) UnsetFundChange()`

UnsetFundChange ensures that no value is present for FundChange, not even an explicit nil
### GetGmtRefundPay

`func (o *AlipayTradeRefundResponse) GetGmtRefundPay() string`

GetGmtRefundPay returns the GmtRefundPay field if non-nil, zero value otherwise.

### GetGmtRefundPayOk

`func (o *AlipayTradeRefundResponse) GetGmtRefundPayOk() (*string, bool)`

GetGmtRefundPayOk returns a tuple with the GmtRefundPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtRefundPay

`func (o *AlipayTradeRefundResponse) SetGmtRefundPay(v string)`

SetGmtRefundPay sets GmtRefundPay field to given value.

### HasGmtRefundPay

`func (o *AlipayTradeRefundResponse) HasGmtRefundPay() bool`

HasGmtRefundPay returns a boolean if a field has been set.

### SetGmtRefundPayNil

`func (o *AlipayTradeRefundResponse) SetGmtRefundPayNil(b bool)`

 SetGmtRefundPayNil sets the value for GmtRefundPay to be an explicit nil

### UnsetGmtRefundPay
`func (o *AlipayTradeRefundResponse) UnsetGmtRefundPay()`

UnsetGmtRefundPay ensures that no value is present for GmtRefundPay, not even an explicit nil
### GetHasDepositBack

`func (o *AlipayTradeRefundResponse) GetHasDepositBack() string`

GetHasDepositBack returns the HasDepositBack field if non-nil, zero value otherwise.

### GetHasDepositBackOk

`func (o *AlipayTradeRefundResponse) GetHasDepositBackOk() (*string, bool)`

GetHasDepositBackOk returns a tuple with the HasDepositBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDepositBack

`func (o *AlipayTradeRefundResponse) SetHasDepositBack(v string)`

SetHasDepositBack sets HasDepositBack field to given value.

### HasHasDepositBack

`func (o *AlipayTradeRefundResponse) HasHasDepositBack() bool`

HasHasDepositBack returns a boolean if a field has been set.

### SetHasDepositBackNil

`func (o *AlipayTradeRefundResponse) SetHasDepositBackNil(b bool)`

 SetHasDepositBackNil sets the value for HasDepositBack to be an explicit nil

### UnsetHasDepositBack
`func (o *AlipayTradeRefundResponse) UnsetHasDepositBack()`

UnsetHasDepositBack ensures that no value is present for HasDepositBack, not even an explicit nil
### GetOpenId

`func (o *AlipayTradeRefundResponse) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *AlipayTradeRefundResponse) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *AlipayTradeRefundResponse) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *AlipayTradeRefundResponse) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### SetOpenIdNil

`func (o *AlipayTradeRefundResponse) SetOpenIdNil(b bool)`

 SetOpenIdNil sets the value for OpenId to be an explicit nil

### UnsetOpenId
`func (o *AlipayTradeRefundResponse) UnsetOpenId()`

UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
### GetOutTradeNo

`func (o *AlipayTradeRefundResponse) GetOutTradeNo() string`

GetOutTradeNo returns the OutTradeNo field if non-nil, zero value otherwise.

### GetOutTradeNoOk

`func (o *AlipayTradeRefundResponse) GetOutTradeNoOk() (*string, bool)`

GetOutTradeNoOk returns a tuple with the OutTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTradeNo

`func (o *AlipayTradeRefundResponse) SetOutTradeNo(v string)`

SetOutTradeNo sets OutTradeNo field to given value.

### HasOutTradeNo

`func (o *AlipayTradeRefundResponse) HasOutTradeNo() bool`

HasOutTradeNo returns a boolean if a field has been set.

### SetOutTradeNoNil

`func (o *AlipayTradeRefundResponse) SetOutTradeNoNil(b bool)`

 SetOutTradeNoNil sets the value for OutTradeNo to be an explicit nil

### UnsetOutTradeNo
`func (o *AlipayTradeRefundResponse) UnsetOutTradeNo()`

UnsetOutTradeNo ensures that no value is present for OutTradeNo, not even an explicit nil
### GetPreAuthCancelFee

`func (o *AlipayTradeRefundResponse) GetPreAuthCancelFee() string`

GetPreAuthCancelFee returns the PreAuthCancelFee field if non-nil, zero value otherwise.

### GetPreAuthCancelFeeOk

`func (o *AlipayTradeRefundResponse) GetPreAuthCancelFeeOk() (*string, bool)`

GetPreAuthCancelFeeOk returns a tuple with the PreAuthCancelFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthCancelFee

`func (o *AlipayTradeRefundResponse) SetPreAuthCancelFee(v string)`

SetPreAuthCancelFee sets PreAuthCancelFee field to given value.

### HasPreAuthCancelFee

`func (o *AlipayTradeRefundResponse) HasPreAuthCancelFee() bool`

HasPreAuthCancelFee returns a boolean if a field has been set.

### SetPreAuthCancelFeeNil

`func (o *AlipayTradeRefundResponse) SetPreAuthCancelFeeNil(b bool)`

 SetPreAuthCancelFeeNil sets the value for PreAuthCancelFee to be an explicit nil

### UnsetPreAuthCancelFee
`func (o *AlipayTradeRefundResponse) UnsetPreAuthCancelFee()`

UnsetPreAuthCancelFee ensures that no value is present for PreAuthCancelFee, not even an explicit nil
### GetPresentRefundBuyerAmount

`func (o *AlipayTradeRefundResponse) GetPresentRefundBuyerAmount() string`

GetPresentRefundBuyerAmount returns the PresentRefundBuyerAmount field if non-nil, zero value otherwise.

### GetPresentRefundBuyerAmountOk

`func (o *AlipayTradeRefundResponse) GetPresentRefundBuyerAmountOk() (*string, bool)`

GetPresentRefundBuyerAmountOk returns a tuple with the PresentRefundBuyerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentRefundBuyerAmount

`func (o *AlipayTradeRefundResponse) SetPresentRefundBuyerAmount(v string)`

SetPresentRefundBuyerAmount sets PresentRefundBuyerAmount field to given value.

### HasPresentRefundBuyerAmount

`func (o *AlipayTradeRefundResponse) HasPresentRefundBuyerAmount() bool`

HasPresentRefundBuyerAmount returns a boolean if a field has been set.

### SetPresentRefundBuyerAmountNil

`func (o *AlipayTradeRefundResponse) SetPresentRefundBuyerAmountNil(b bool)`

 SetPresentRefundBuyerAmountNil sets the value for PresentRefundBuyerAmount to be an explicit nil

### UnsetPresentRefundBuyerAmount
`func (o *AlipayTradeRefundResponse) UnsetPresentRefundBuyerAmount()`

UnsetPresentRefundBuyerAmount ensures that no value is present for PresentRefundBuyerAmount, not even an explicit nil
### GetPresentRefundDiscountAmount

`func (o *AlipayTradeRefundResponse) GetPresentRefundDiscountAmount() string`

GetPresentRefundDiscountAmount returns the PresentRefundDiscountAmount field if non-nil, zero value otherwise.

### GetPresentRefundDiscountAmountOk

`func (o *AlipayTradeRefundResponse) GetPresentRefundDiscountAmountOk() (*string, bool)`

GetPresentRefundDiscountAmountOk returns a tuple with the PresentRefundDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentRefundDiscountAmount

`func (o *AlipayTradeRefundResponse) SetPresentRefundDiscountAmount(v string)`

SetPresentRefundDiscountAmount sets PresentRefundDiscountAmount field to given value.

### HasPresentRefundDiscountAmount

`func (o *AlipayTradeRefundResponse) HasPresentRefundDiscountAmount() bool`

HasPresentRefundDiscountAmount returns a boolean if a field has been set.

### SetPresentRefundDiscountAmountNil

`func (o *AlipayTradeRefundResponse) SetPresentRefundDiscountAmountNil(b bool)`

 SetPresentRefundDiscountAmountNil sets the value for PresentRefundDiscountAmount to be an explicit nil

### UnsetPresentRefundDiscountAmount
`func (o *AlipayTradeRefundResponse) UnsetPresentRefundDiscountAmount()`

UnsetPresentRefundDiscountAmount ensures that no value is present for PresentRefundDiscountAmount, not even an explicit nil
### GetPresentRefundMdiscountAmount

`func (o *AlipayTradeRefundResponse) GetPresentRefundMdiscountAmount() string`

GetPresentRefundMdiscountAmount returns the PresentRefundMdiscountAmount field if non-nil, zero value otherwise.

### GetPresentRefundMdiscountAmountOk

`func (o *AlipayTradeRefundResponse) GetPresentRefundMdiscountAmountOk() (*string, bool)`

GetPresentRefundMdiscountAmountOk returns a tuple with the PresentRefundMdiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentRefundMdiscountAmount

`func (o *AlipayTradeRefundResponse) SetPresentRefundMdiscountAmount(v string)`

SetPresentRefundMdiscountAmount sets PresentRefundMdiscountAmount field to given value.

### HasPresentRefundMdiscountAmount

`func (o *AlipayTradeRefundResponse) HasPresentRefundMdiscountAmount() bool`

HasPresentRefundMdiscountAmount returns a boolean if a field has been set.

### SetPresentRefundMdiscountAmountNil

`func (o *AlipayTradeRefundResponse) SetPresentRefundMdiscountAmountNil(b bool)`

 SetPresentRefundMdiscountAmountNil sets the value for PresentRefundMdiscountAmount to be an explicit nil

### UnsetPresentRefundMdiscountAmount
`func (o *AlipayTradeRefundResponse) UnsetPresentRefundMdiscountAmount()`

UnsetPresentRefundMdiscountAmount ensures that no value is present for PresentRefundMdiscountAmount, not even an explicit nil
### GetRefundChargeAmount

`func (o *AlipayTradeRefundResponse) GetRefundChargeAmount() string`

GetRefundChargeAmount returns the RefundChargeAmount field if non-nil, zero value otherwise.

### GetRefundChargeAmountOk

`func (o *AlipayTradeRefundResponse) GetRefundChargeAmountOk() (*string, bool)`

GetRefundChargeAmountOk returns a tuple with the RefundChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundChargeAmount

`func (o *AlipayTradeRefundResponse) SetRefundChargeAmount(v string)`

SetRefundChargeAmount sets RefundChargeAmount field to given value.

### HasRefundChargeAmount

`func (o *AlipayTradeRefundResponse) HasRefundChargeAmount() bool`

HasRefundChargeAmount returns a boolean if a field has been set.

### SetRefundChargeAmountNil

`func (o *AlipayTradeRefundResponse) SetRefundChargeAmountNil(b bool)`

 SetRefundChargeAmountNil sets the value for RefundChargeAmount to be an explicit nil

### UnsetRefundChargeAmount
`func (o *AlipayTradeRefundResponse) UnsetRefundChargeAmount()`

UnsetRefundChargeAmount ensures that no value is present for RefundChargeAmount, not even an explicit nil
### GetRefundChargeInfoList

`func (o *AlipayTradeRefundResponse) GetRefundChargeInfoList() []RefundChargeInfo`

GetRefundChargeInfoList returns the RefundChargeInfoList field if non-nil, zero value otherwise.

### GetRefundChargeInfoListOk

`func (o *AlipayTradeRefundResponse) GetRefundChargeInfoListOk() (*[]RefundChargeInfo, bool)`

GetRefundChargeInfoListOk returns a tuple with the RefundChargeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundChargeInfoList

`func (o *AlipayTradeRefundResponse) SetRefundChargeInfoList(v []RefundChargeInfo)`

SetRefundChargeInfoList sets RefundChargeInfoList field to given value.

### HasRefundChargeInfoList

`func (o *AlipayTradeRefundResponse) HasRefundChargeInfoList() bool`

HasRefundChargeInfoList returns a boolean if a field has been set.

### SetRefundChargeInfoListNil

`func (o *AlipayTradeRefundResponse) SetRefundChargeInfoListNil(b bool)`

 SetRefundChargeInfoListNil sets the value for RefundChargeInfoList to be an explicit nil

### UnsetRefundChargeInfoList
`func (o *AlipayTradeRefundResponse) UnsetRefundChargeInfoList()`

UnsetRefundChargeInfoList ensures that no value is present for RefundChargeInfoList, not even an explicit nil
### GetRefundCurrency

`func (o *AlipayTradeRefundResponse) GetRefundCurrency() string`

GetRefundCurrency returns the RefundCurrency field if non-nil, zero value otherwise.

### GetRefundCurrencyOk

`func (o *AlipayTradeRefundResponse) GetRefundCurrencyOk() (*string, bool)`

GetRefundCurrencyOk returns a tuple with the RefundCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCurrency

`func (o *AlipayTradeRefundResponse) SetRefundCurrency(v string)`

SetRefundCurrency sets RefundCurrency field to given value.

### HasRefundCurrency

`func (o *AlipayTradeRefundResponse) HasRefundCurrency() bool`

HasRefundCurrency returns a boolean if a field has been set.

### SetRefundCurrencyNil

`func (o *AlipayTradeRefundResponse) SetRefundCurrencyNil(b bool)`

 SetRefundCurrencyNil sets the value for RefundCurrency to be an explicit nil

### UnsetRefundCurrency
`func (o *AlipayTradeRefundResponse) UnsetRefundCurrency()`

UnsetRefundCurrency ensures that no value is present for RefundCurrency, not even an explicit nil
### GetRefundDetailItemList

`func (o *AlipayTradeRefundResponse) GetRefundDetailItemList() []TradeFundBill`

GetRefundDetailItemList returns the RefundDetailItemList field if non-nil, zero value otherwise.

### GetRefundDetailItemListOk

`func (o *AlipayTradeRefundResponse) GetRefundDetailItemListOk() (*[]TradeFundBill, bool)`

GetRefundDetailItemListOk returns a tuple with the RefundDetailItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetailItemList

`func (o *AlipayTradeRefundResponse) SetRefundDetailItemList(v []TradeFundBill)`

SetRefundDetailItemList sets RefundDetailItemList field to given value.

### HasRefundDetailItemList

`func (o *AlipayTradeRefundResponse) HasRefundDetailItemList() bool`

HasRefundDetailItemList returns a boolean if a field has been set.

### SetRefundDetailItemListNil

`func (o *AlipayTradeRefundResponse) SetRefundDetailItemListNil(b bool)`

 SetRefundDetailItemListNil sets the value for RefundDetailItemList to be an explicit nil

### UnsetRefundDetailItemList
`func (o *AlipayTradeRefundResponse) UnsetRefundDetailItemList()`

UnsetRefundDetailItemList ensures that no value is present for RefundDetailItemList, not even an explicit nil
### GetRefundFee

`func (o *AlipayTradeRefundResponse) GetRefundFee() string`

GetRefundFee returns the RefundFee field if non-nil, zero value otherwise.

### GetRefundFeeOk

`func (o *AlipayTradeRefundResponse) GetRefundFeeOk() (*string, bool)`

GetRefundFeeOk returns a tuple with the RefundFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundFee

`func (o *AlipayTradeRefundResponse) SetRefundFee(v string)`

SetRefundFee sets RefundFee field to given value.

### HasRefundFee

`func (o *AlipayTradeRefundResponse) HasRefundFee() bool`

HasRefundFee returns a boolean if a field has been set.

### SetRefundFeeNil

`func (o *AlipayTradeRefundResponse) SetRefundFeeNil(b bool)`

 SetRefundFeeNil sets the value for RefundFee to be an explicit nil

### UnsetRefundFee
`func (o *AlipayTradeRefundResponse) UnsetRefundFee()`

UnsetRefundFee ensures that no value is present for RefundFee, not even an explicit nil
### GetRefundHybAmount

`func (o *AlipayTradeRefundResponse) GetRefundHybAmount() string`

GetRefundHybAmount returns the RefundHybAmount field if non-nil, zero value otherwise.

### GetRefundHybAmountOk

`func (o *AlipayTradeRefundResponse) GetRefundHybAmountOk() (*string, bool)`

GetRefundHybAmountOk returns a tuple with the RefundHybAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundHybAmount

`func (o *AlipayTradeRefundResponse) SetRefundHybAmount(v string)`

SetRefundHybAmount sets RefundHybAmount field to given value.

### HasRefundHybAmount

`func (o *AlipayTradeRefundResponse) HasRefundHybAmount() bool`

HasRefundHybAmount returns a boolean if a field has been set.

### SetRefundHybAmountNil

`func (o *AlipayTradeRefundResponse) SetRefundHybAmountNil(b bool)`

 SetRefundHybAmountNil sets the value for RefundHybAmount to be an explicit nil

### UnsetRefundHybAmount
`func (o *AlipayTradeRefundResponse) UnsetRefundHybAmount()`

UnsetRefundHybAmount ensures that no value is present for RefundHybAmount, not even an explicit nil
### GetRefundPresetPaytoolList

`func (o *AlipayTradeRefundResponse) GetRefundPresetPaytoolList() PresetPayToolInfo`

GetRefundPresetPaytoolList returns the RefundPresetPaytoolList field if non-nil, zero value otherwise.

### GetRefundPresetPaytoolListOk

`func (o *AlipayTradeRefundResponse) GetRefundPresetPaytoolListOk() (*PresetPayToolInfo, bool)`

GetRefundPresetPaytoolListOk returns a tuple with the RefundPresetPaytoolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPresetPaytoolList

`func (o *AlipayTradeRefundResponse) SetRefundPresetPaytoolList(v PresetPayToolInfo)`

SetRefundPresetPaytoolList sets RefundPresetPaytoolList field to given value.

### HasRefundPresetPaytoolList

`func (o *AlipayTradeRefundResponse) HasRefundPresetPaytoolList() bool`

HasRefundPresetPaytoolList returns a boolean if a field has been set.

### GetRefundSettlementId

`func (o *AlipayTradeRefundResponse) GetRefundSettlementId() string`

GetRefundSettlementId returns the RefundSettlementId field if non-nil, zero value otherwise.

### GetRefundSettlementIdOk

`func (o *AlipayTradeRefundResponse) GetRefundSettlementIdOk() (*string, bool)`

GetRefundSettlementIdOk returns a tuple with the RefundSettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundSettlementId

`func (o *AlipayTradeRefundResponse) SetRefundSettlementId(v string)`

SetRefundSettlementId sets RefundSettlementId field to given value.

### HasRefundSettlementId

`func (o *AlipayTradeRefundResponse) HasRefundSettlementId() bool`

HasRefundSettlementId returns a boolean if a field has been set.

### SetRefundSettlementIdNil

`func (o *AlipayTradeRefundResponse) SetRefundSettlementIdNil(b bool)`

 SetRefundSettlementIdNil sets the value for RefundSettlementId to be an explicit nil

### UnsetRefundSettlementId
`func (o *AlipayTradeRefundResponse) UnsetRefundSettlementId()`

UnsetRefundSettlementId ensures that no value is present for RefundSettlementId, not even an explicit nil
### GetRefundVoucherDetailList

`func (o *AlipayTradeRefundResponse) GetRefundVoucherDetailList() []VoucherDetail`

GetRefundVoucherDetailList returns the RefundVoucherDetailList field if non-nil, zero value otherwise.

### GetRefundVoucherDetailListOk

`func (o *AlipayTradeRefundResponse) GetRefundVoucherDetailListOk() (*[]VoucherDetail, bool)`

GetRefundVoucherDetailListOk returns a tuple with the RefundVoucherDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundVoucherDetailList

`func (o *AlipayTradeRefundResponse) SetRefundVoucherDetailList(v []VoucherDetail)`

SetRefundVoucherDetailList sets RefundVoucherDetailList field to given value.

### HasRefundVoucherDetailList

`func (o *AlipayTradeRefundResponse) HasRefundVoucherDetailList() bool`

HasRefundVoucherDetailList returns a boolean if a field has been set.

### SetRefundVoucherDetailListNil

`func (o *AlipayTradeRefundResponse) SetRefundVoucherDetailListNil(b bool)`

 SetRefundVoucherDetailListNil sets the value for RefundVoucherDetailList to be an explicit nil

### UnsetRefundVoucherDetailList
`func (o *AlipayTradeRefundResponse) UnsetRefundVoucherDetailList()`

UnsetRefundVoucherDetailList ensures that no value is present for RefundVoucherDetailList, not even an explicit nil
### GetSendBackFee

`func (o *AlipayTradeRefundResponse) GetSendBackFee() string`

GetSendBackFee returns the SendBackFee field if non-nil, zero value otherwise.

### GetSendBackFeeOk

`func (o *AlipayTradeRefundResponse) GetSendBackFeeOk() (*string, bool)`

GetSendBackFeeOk returns a tuple with the SendBackFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBackFee

`func (o *AlipayTradeRefundResponse) SetSendBackFee(v string)`

SetSendBackFee sets SendBackFee field to given value.

### HasSendBackFee

`func (o *AlipayTradeRefundResponse) HasSendBackFee() bool`

HasSendBackFee returns a boolean if a field has been set.

### SetSendBackFeeNil

`func (o *AlipayTradeRefundResponse) SetSendBackFeeNil(b bool)`

 SetSendBackFeeNil sets the value for SendBackFee to be an explicit nil

### UnsetSendBackFee
`func (o *AlipayTradeRefundResponse) UnsetSendBackFee()`

UnsetSendBackFee ensures that no value is present for SendBackFee, not even an explicit nil
### GetStoreName

`func (o *AlipayTradeRefundResponse) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *AlipayTradeRefundResponse) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *AlipayTradeRefundResponse) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *AlipayTradeRefundResponse) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### SetStoreNameNil

`func (o *AlipayTradeRefundResponse) SetStoreNameNil(b bool)`

 SetStoreNameNil sets the value for StoreName to be an explicit nil

### UnsetStoreName
`func (o *AlipayTradeRefundResponse) UnsetStoreName()`

UnsetStoreName ensures that no value is present for StoreName, not even an explicit nil
### GetTradeNo

`func (o *AlipayTradeRefundResponse) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *AlipayTradeRefundResponse) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *AlipayTradeRefundResponse) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *AlipayTradeRefundResponse) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### SetTradeNoNil

`func (o *AlipayTradeRefundResponse) SetTradeNoNil(b bool)`

 SetTradeNoNil sets the value for TradeNo to be an explicit nil

### UnsetTradeNo
`func (o *AlipayTradeRefundResponse) UnsetTradeNo()`

UnsetTradeNo ensures that no value is present for TradeNo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


