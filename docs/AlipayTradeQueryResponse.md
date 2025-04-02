# AlipayTradeQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 
**SubCode** | Pointer to **NullableString** |  | [optional] 
**SubMsg** | Pointer to **NullableString** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] [readonly] 
**AdditionalStatus** | Pointer to **NullableString** |  | [optional] 
**AlipayStoreId** | Pointer to **NullableString** |  | [optional] 
**AlipaySubMerchantId** | Pointer to **NullableString** |  | [optional] 
**AsyncPayApplyStatus** | Pointer to **NullableString** |  | [optional] 
**AuthTradePayMode** | Pointer to **NullableString** |  | [optional] 
**BizSettleMode** | Pointer to **NullableString** |  | [optional] 
**BkagentRespInfo** | Pointer to [**BkAgentRespInfo**](BkAgentRespInfo.md) |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**BuyerLogonId** | Pointer to **NullableString** |  | [optional] 
**BuyerOpenId** | Pointer to **NullableString** |  | [optional] 
**BuyerPayAmount** | Pointer to **NullableString** |  | [optional] 
**BuyerUserId** | Pointer to **NullableString** |  | [optional] 
**BuyerUserName** | Pointer to **NullableString** |  | [optional] 
**BuyerUserType** | Pointer to **NullableString** |  | [optional] 
**CashierType** | Pointer to **NullableString** |  | [optional] 
**ChargeAmount** | Pointer to **NullableString** |  | [optional] 
**ChargeFlags** | Pointer to **NullableString** |  | [optional] 
**ChargeInfoList** | Pointer to [**[]ChargeInfo**](ChargeInfo.md) |  | [optional] 
**CreditBizOrderId** | Pointer to **NullableString** |  | [optional] 
**CreditPayMode** | Pointer to **NullableString** |  | [optional] 
**DiscountAmount** | Pointer to **NullableString** |  | [optional] 
**DiscountGoodsDetail** | Pointer to **NullableString** |  | [optional] 
**EnterprisePayInfo** | Pointer to [**EnterprisePayInfo**](EnterprisePayInfo.md) |  | [optional] 
**ExtInfos** | Pointer to **NullableString** |  | [optional] 
**FulfillmentDetailList** | Pointer to [**[]FulfillmentDetail**](FulfillmentDetail.md) |  | [optional] 
**FundBillList** | Pointer to [**[]TradeFundBill**](TradeFundBill.md) |  | [optional] 
**HbFqPayInfo** | Pointer to [**HbFqPayInfo**](HbFqPayInfo.md) |  | [optional] 
**HybAmount** | Pointer to **NullableString** |  | [optional] 
**IndustrySepcDetail** | Pointer to **NullableString** |  | [optional] 
**IndustrySepcDetailAcc** | Pointer to **NullableString** |  | [optional] 
**IndustrySepcDetailGov** | Pointer to **NullableString** |  | [optional] 
**IntactChargeInfoList** | Pointer to [**[]IntactChargeInfo**](IntactChargeInfo.md) |  | [optional] 
**InvoiceAmount** | Pointer to **NullableString** |  | [optional] 
**MdiscountAmount** | Pointer to **NullableString** |  | [optional] 
**MedicalInsuranceInfo** | Pointer to **NullableString** |  | [optional] 
**OpenId** | Pointer to **NullableString** |  | [optional] 
**OutTradeNo** | Pointer to **NullableString** |  | [optional] 
**PassbackParams** | Pointer to **NullableString** |  | [optional] 
**PayAmount** | Pointer to **NullableString** |  | [optional] 
**PayCurrency** | Pointer to **NullableString** |  | [optional] 
**PaymentInfoWithIdList** | Pointer to [**[]PaymentInfoWithId**](PaymentInfoWithId.md) |  | [optional] 
**PeriodScene** | Pointer to **NullableString** |  | [optional] 
**PointAmount** | Pointer to **NullableString** |  | [optional] 
**PreAuthPayAmount** | Pointer to **NullableString** |  | [optional] 
**ReceiptAmount** | Pointer to **NullableString** |  | [optional] 
**ReceiptCurrencyType** | Pointer to **NullableString** |  | [optional] 
**ReqGoodsDetail** | Pointer to [**[]GoodsDetail**](GoodsDetail.md) |  | [optional] 
**SendPayDate** | Pointer to **NullableString** |  | [optional] 
**SettleAmount** | Pointer to **NullableString** |  | [optional] 
**SettleCurrency** | Pointer to **NullableString** |  | [optional] 
**SettleTransRate** | Pointer to **NullableString** |  | [optional] 
**SettlementId** | Pointer to **NullableString** |  | [optional] 
**StoreId** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**TapPayInfo** | Pointer to [**TapPayInfo**](TapPayInfo.md) |  | [optional] 
**TerminalId** | Pointer to **NullableString** |  | [optional] 
**TotalAmount** | Pointer to **NullableString** |  | [optional] 
**TradeNo** | Pointer to **NullableString** |  | [optional] 
**TradeSettleInfo** | Pointer to [**TradeSettleInfo**](TradeSettleInfo.md) |  | [optional] 
**TradeStatus** | Pointer to **NullableString** |  | [optional] 
**TransCurrency** | Pointer to **NullableString** |  | [optional] 
**TransPayRate** | Pointer to **NullableString** |  | [optional] 
**VoucherDetailList** | Pointer to [**[]VoucherDetail**](VoucherDetail.md) |  | [optional] 

## Methods

### NewAlipayTradeQueryResponse

`func NewAlipayTradeQueryResponse() *AlipayTradeQueryResponse`

NewAlipayTradeQueryResponse instantiates a new AlipayTradeQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlipayTradeQueryResponseWithDefaults

`func NewAlipayTradeQueryResponseWithDefaults() *AlipayTradeQueryResponse`

NewAlipayTradeQueryResponseWithDefaults instantiates a new AlipayTradeQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AlipayTradeQueryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AlipayTradeQueryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AlipayTradeQueryResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AlipayTradeQueryResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AlipayTradeQueryResponse) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AlipayTradeQueryResponse) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMsg

`func (o *AlipayTradeQueryResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AlipayTradeQueryResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AlipayTradeQueryResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AlipayTradeQueryResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *AlipayTradeQueryResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *AlipayTradeQueryResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetSubCode

`func (o *AlipayTradeQueryResponse) GetSubCode() string`

GetSubCode returns the SubCode field if non-nil, zero value otherwise.

### GetSubCodeOk

`func (o *AlipayTradeQueryResponse) GetSubCodeOk() (*string, bool)`

GetSubCodeOk returns a tuple with the SubCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCode

`func (o *AlipayTradeQueryResponse) SetSubCode(v string)`

SetSubCode sets SubCode field to given value.

### HasSubCode

`func (o *AlipayTradeQueryResponse) HasSubCode() bool`

HasSubCode returns a boolean if a field has been set.

### SetSubCodeNil

`func (o *AlipayTradeQueryResponse) SetSubCodeNil(b bool)`

 SetSubCodeNil sets the value for SubCode to be an explicit nil

### UnsetSubCode
`func (o *AlipayTradeQueryResponse) UnsetSubCode()`

UnsetSubCode ensures that no value is present for SubCode, not even an explicit nil
### GetSubMsg

`func (o *AlipayTradeQueryResponse) GetSubMsg() string`

GetSubMsg returns the SubMsg field if non-nil, zero value otherwise.

### GetSubMsgOk

`func (o *AlipayTradeQueryResponse) GetSubMsgOk() (*string, bool)`

GetSubMsgOk returns a tuple with the SubMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMsg

`func (o *AlipayTradeQueryResponse) SetSubMsg(v string)`

SetSubMsg sets SubMsg field to given value.

### HasSubMsg

`func (o *AlipayTradeQueryResponse) HasSubMsg() bool`

HasSubMsg returns a boolean if a field has been set.

### SetSubMsgNil

`func (o *AlipayTradeQueryResponse) SetSubMsgNil(b bool)`

 SetSubMsgNil sets the value for SubMsg to be an explicit nil

### UnsetSubMsg
`func (o *AlipayTradeQueryResponse) UnsetSubMsg()`

UnsetSubMsg ensures that no value is present for SubMsg, not even an explicit nil
### GetIsError

`func (o *AlipayTradeQueryResponse) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AlipayTradeQueryResponse) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AlipayTradeQueryResponse) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *AlipayTradeQueryResponse) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetAdditionalStatus

`func (o *AlipayTradeQueryResponse) GetAdditionalStatus() string`

GetAdditionalStatus returns the AdditionalStatus field if non-nil, zero value otherwise.

### GetAdditionalStatusOk

`func (o *AlipayTradeQueryResponse) GetAdditionalStatusOk() (*string, bool)`

GetAdditionalStatusOk returns a tuple with the AdditionalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStatus

`func (o *AlipayTradeQueryResponse) SetAdditionalStatus(v string)`

SetAdditionalStatus sets AdditionalStatus field to given value.

### HasAdditionalStatus

`func (o *AlipayTradeQueryResponse) HasAdditionalStatus() bool`

HasAdditionalStatus returns a boolean if a field has been set.

### SetAdditionalStatusNil

`func (o *AlipayTradeQueryResponse) SetAdditionalStatusNil(b bool)`

 SetAdditionalStatusNil sets the value for AdditionalStatus to be an explicit nil

### UnsetAdditionalStatus
`func (o *AlipayTradeQueryResponse) UnsetAdditionalStatus()`

UnsetAdditionalStatus ensures that no value is present for AdditionalStatus, not even an explicit nil
### GetAlipayStoreId

`func (o *AlipayTradeQueryResponse) GetAlipayStoreId() string`

GetAlipayStoreId returns the AlipayStoreId field if non-nil, zero value otherwise.

### GetAlipayStoreIdOk

`func (o *AlipayTradeQueryResponse) GetAlipayStoreIdOk() (*string, bool)`

GetAlipayStoreIdOk returns a tuple with the AlipayStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlipayStoreId

`func (o *AlipayTradeQueryResponse) SetAlipayStoreId(v string)`

SetAlipayStoreId sets AlipayStoreId field to given value.

### HasAlipayStoreId

`func (o *AlipayTradeQueryResponse) HasAlipayStoreId() bool`

HasAlipayStoreId returns a boolean if a field has been set.

### SetAlipayStoreIdNil

`func (o *AlipayTradeQueryResponse) SetAlipayStoreIdNil(b bool)`

 SetAlipayStoreIdNil sets the value for AlipayStoreId to be an explicit nil

### UnsetAlipayStoreId
`func (o *AlipayTradeQueryResponse) UnsetAlipayStoreId()`

UnsetAlipayStoreId ensures that no value is present for AlipayStoreId, not even an explicit nil
### GetAlipaySubMerchantId

`func (o *AlipayTradeQueryResponse) GetAlipaySubMerchantId() string`

GetAlipaySubMerchantId returns the AlipaySubMerchantId field if non-nil, zero value otherwise.

### GetAlipaySubMerchantIdOk

`func (o *AlipayTradeQueryResponse) GetAlipaySubMerchantIdOk() (*string, bool)`

GetAlipaySubMerchantIdOk returns a tuple with the AlipaySubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlipaySubMerchantId

`func (o *AlipayTradeQueryResponse) SetAlipaySubMerchantId(v string)`

SetAlipaySubMerchantId sets AlipaySubMerchantId field to given value.

### HasAlipaySubMerchantId

`func (o *AlipayTradeQueryResponse) HasAlipaySubMerchantId() bool`

HasAlipaySubMerchantId returns a boolean if a field has been set.

### SetAlipaySubMerchantIdNil

`func (o *AlipayTradeQueryResponse) SetAlipaySubMerchantIdNil(b bool)`

 SetAlipaySubMerchantIdNil sets the value for AlipaySubMerchantId to be an explicit nil

### UnsetAlipaySubMerchantId
`func (o *AlipayTradeQueryResponse) UnsetAlipaySubMerchantId()`

UnsetAlipaySubMerchantId ensures that no value is present for AlipaySubMerchantId, not even an explicit nil
### GetAsyncPayApplyStatus

`func (o *AlipayTradeQueryResponse) GetAsyncPayApplyStatus() string`

GetAsyncPayApplyStatus returns the AsyncPayApplyStatus field if non-nil, zero value otherwise.

### GetAsyncPayApplyStatusOk

`func (o *AlipayTradeQueryResponse) GetAsyncPayApplyStatusOk() (*string, bool)`

GetAsyncPayApplyStatusOk returns a tuple with the AsyncPayApplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncPayApplyStatus

`func (o *AlipayTradeQueryResponse) SetAsyncPayApplyStatus(v string)`

SetAsyncPayApplyStatus sets AsyncPayApplyStatus field to given value.

### HasAsyncPayApplyStatus

`func (o *AlipayTradeQueryResponse) HasAsyncPayApplyStatus() bool`

HasAsyncPayApplyStatus returns a boolean if a field has been set.

### SetAsyncPayApplyStatusNil

`func (o *AlipayTradeQueryResponse) SetAsyncPayApplyStatusNil(b bool)`

 SetAsyncPayApplyStatusNil sets the value for AsyncPayApplyStatus to be an explicit nil

### UnsetAsyncPayApplyStatus
`func (o *AlipayTradeQueryResponse) UnsetAsyncPayApplyStatus()`

UnsetAsyncPayApplyStatus ensures that no value is present for AsyncPayApplyStatus, not even an explicit nil
### GetAuthTradePayMode

`func (o *AlipayTradeQueryResponse) GetAuthTradePayMode() string`

GetAuthTradePayMode returns the AuthTradePayMode field if non-nil, zero value otherwise.

### GetAuthTradePayModeOk

`func (o *AlipayTradeQueryResponse) GetAuthTradePayModeOk() (*string, bool)`

GetAuthTradePayModeOk returns a tuple with the AuthTradePayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTradePayMode

`func (o *AlipayTradeQueryResponse) SetAuthTradePayMode(v string)`

SetAuthTradePayMode sets AuthTradePayMode field to given value.

### HasAuthTradePayMode

`func (o *AlipayTradeQueryResponse) HasAuthTradePayMode() bool`

HasAuthTradePayMode returns a boolean if a field has been set.

### SetAuthTradePayModeNil

`func (o *AlipayTradeQueryResponse) SetAuthTradePayModeNil(b bool)`

 SetAuthTradePayModeNil sets the value for AuthTradePayMode to be an explicit nil

### UnsetAuthTradePayMode
`func (o *AlipayTradeQueryResponse) UnsetAuthTradePayMode()`

UnsetAuthTradePayMode ensures that no value is present for AuthTradePayMode, not even an explicit nil
### GetBizSettleMode

`func (o *AlipayTradeQueryResponse) GetBizSettleMode() string`

GetBizSettleMode returns the BizSettleMode field if non-nil, zero value otherwise.

### GetBizSettleModeOk

`func (o *AlipayTradeQueryResponse) GetBizSettleModeOk() (*string, bool)`

GetBizSettleModeOk returns a tuple with the BizSettleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizSettleMode

`func (o *AlipayTradeQueryResponse) SetBizSettleMode(v string)`

SetBizSettleMode sets BizSettleMode field to given value.

### HasBizSettleMode

`func (o *AlipayTradeQueryResponse) HasBizSettleMode() bool`

HasBizSettleMode returns a boolean if a field has been set.

### SetBizSettleModeNil

`func (o *AlipayTradeQueryResponse) SetBizSettleModeNil(b bool)`

 SetBizSettleModeNil sets the value for BizSettleMode to be an explicit nil

### UnsetBizSettleMode
`func (o *AlipayTradeQueryResponse) UnsetBizSettleMode()`

UnsetBizSettleMode ensures that no value is present for BizSettleMode, not even an explicit nil
### GetBkagentRespInfo

`func (o *AlipayTradeQueryResponse) GetBkagentRespInfo() BkAgentRespInfo`

GetBkagentRespInfo returns the BkagentRespInfo field if non-nil, zero value otherwise.

### GetBkagentRespInfoOk

`func (o *AlipayTradeQueryResponse) GetBkagentRespInfoOk() (*BkAgentRespInfo, bool)`

GetBkagentRespInfoOk returns a tuple with the BkagentRespInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBkagentRespInfo

`func (o *AlipayTradeQueryResponse) SetBkagentRespInfo(v BkAgentRespInfo)`

SetBkagentRespInfo sets BkagentRespInfo field to given value.

### HasBkagentRespInfo

`func (o *AlipayTradeQueryResponse) HasBkagentRespInfo() bool`

HasBkagentRespInfo returns a boolean if a field has been set.

### GetBody

`func (o *AlipayTradeQueryResponse) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AlipayTradeQueryResponse) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AlipayTradeQueryResponse) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AlipayTradeQueryResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *AlipayTradeQueryResponse) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *AlipayTradeQueryResponse) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetBuyerLogonId

`func (o *AlipayTradeQueryResponse) GetBuyerLogonId() string`

GetBuyerLogonId returns the BuyerLogonId field if non-nil, zero value otherwise.

### GetBuyerLogonIdOk

`func (o *AlipayTradeQueryResponse) GetBuyerLogonIdOk() (*string, bool)`

GetBuyerLogonIdOk returns a tuple with the BuyerLogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerLogonId

`func (o *AlipayTradeQueryResponse) SetBuyerLogonId(v string)`

SetBuyerLogonId sets BuyerLogonId field to given value.

### HasBuyerLogonId

`func (o *AlipayTradeQueryResponse) HasBuyerLogonId() bool`

HasBuyerLogonId returns a boolean if a field has been set.

### SetBuyerLogonIdNil

`func (o *AlipayTradeQueryResponse) SetBuyerLogonIdNil(b bool)`

 SetBuyerLogonIdNil sets the value for BuyerLogonId to be an explicit nil

### UnsetBuyerLogonId
`func (o *AlipayTradeQueryResponse) UnsetBuyerLogonId()`

UnsetBuyerLogonId ensures that no value is present for BuyerLogonId, not even an explicit nil
### GetBuyerOpenId

`func (o *AlipayTradeQueryResponse) GetBuyerOpenId() string`

GetBuyerOpenId returns the BuyerOpenId field if non-nil, zero value otherwise.

### GetBuyerOpenIdOk

`func (o *AlipayTradeQueryResponse) GetBuyerOpenIdOk() (*string, bool)`

GetBuyerOpenIdOk returns a tuple with the BuyerOpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerOpenId

`func (o *AlipayTradeQueryResponse) SetBuyerOpenId(v string)`

SetBuyerOpenId sets BuyerOpenId field to given value.

### HasBuyerOpenId

`func (o *AlipayTradeQueryResponse) HasBuyerOpenId() bool`

HasBuyerOpenId returns a boolean if a field has been set.

### SetBuyerOpenIdNil

`func (o *AlipayTradeQueryResponse) SetBuyerOpenIdNil(b bool)`

 SetBuyerOpenIdNil sets the value for BuyerOpenId to be an explicit nil

### UnsetBuyerOpenId
`func (o *AlipayTradeQueryResponse) UnsetBuyerOpenId()`

UnsetBuyerOpenId ensures that no value is present for BuyerOpenId, not even an explicit nil
### GetBuyerPayAmount

`func (o *AlipayTradeQueryResponse) GetBuyerPayAmount() string`

GetBuyerPayAmount returns the BuyerPayAmount field if non-nil, zero value otherwise.

### GetBuyerPayAmountOk

`func (o *AlipayTradeQueryResponse) GetBuyerPayAmountOk() (*string, bool)`

GetBuyerPayAmountOk returns a tuple with the BuyerPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPayAmount

`func (o *AlipayTradeQueryResponse) SetBuyerPayAmount(v string)`

SetBuyerPayAmount sets BuyerPayAmount field to given value.

### HasBuyerPayAmount

`func (o *AlipayTradeQueryResponse) HasBuyerPayAmount() bool`

HasBuyerPayAmount returns a boolean if a field has been set.

### SetBuyerPayAmountNil

`func (o *AlipayTradeQueryResponse) SetBuyerPayAmountNil(b bool)`

 SetBuyerPayAmountNil sets the value for BuyerPayAmount to be an explicit nil

### UnsetBuyerPayAmount
`func (o *AlipayTradeQueryResponse) UnsetBuyerPayAmount()`

UnsetBuyerPayAmount ensures that no value is present for BuyerPayAmount, not even an explicit nil
### GetBuyerUserId

`func (o *AlipayTradeQueryResponse) GetBuyerUserId() string`

GetBuyerUserId returns the BuyerUserId field if non-nil, zero value otherwise.

### GetBuyerUserIdOk

`func (o *AlipayTradeQueryResponse) GetBuyerUserIdOk() (*string, bool)`

GetBuyerUserIdOk returns a tuple with the BuyerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserId

`func (o *AlipayTradeQueryResponse) SetBuyerUserId(v string)`

SetBuyerUserId sets BuyerUserId field to given value.

### HasBuyerUserId

`func (o *AlipayTradeQueryResponse) HasBuyerUserId() bool`

HasBuyerUserId returns a boolean if a field has been set.

### SetBuyerUserIdNil

`func (o *AlipayTradeQueryResponse) SetBuyerUserIdNil(b bool)`

 SetBuyerUserIdNil sets the value for BuyerUserId to be an explicit nil

### UnsetBuyerUserId
`func (o *AlipayTradeQueryResponse) UnsetBuyerUserId()`

UnsetBuyerUserId ensures that no value is present for BuyerUserId, not even an explicit nil
### GetBuyerUserName

`func (o *AlipayTradeQueryResponse) GetBuyerUserName() string`

GetBuyerUserName returns the BuyerUserName field if non-nil, zero value otherwise.

### GetBuyerUserNameOk

`func (o *AlipayTradeQueryResponse) GetBuyerUserNameOk() (*string, bool)`

GetBuyerUserNameOk returns a tuple with the BuyerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserName

`func (o *AlipayTradeQueryResponse) SetBuyerUserName(v string)`

SetBuyerUserName sets BuyerUserName field to given value.

### HasBuyerUserName

`func (o *AlipayTradeQueryResponse) HasBuyerUserName() bool`

HasBuyerUserName returns a boolean if a field has been set.

### SetBuyerUserNameNil

`func (o *AlipayTradeQueryResponse) SetBuyerUserNameNil(b bool)`

 SetBuyerUserNameNil sets the value for BuyerUserName to be an explicit nil

### UnsetBuyerUserName
`func (o *AlipayTradeQueryResponse) UnsetBuyerUserName()`

UnsetBuyerUserName ensures that no value is present for BuyerUserName, not even an explicit nil
### GetBuyerUserType

`func (o *AlipayTradeQueryResponse) GetBuyerUserType() string`

GetBuyerUserType returns the BuyerUserType field if non-nil, zero value otherwise.

### GetBuyerUserTypeOk

`func (o *AlipayTradeQueryResponse) GetBuyerUserTypeOk() (*string, bool)`

GetBuyerUserTypeOk returns a tuple with the BuyerUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUserType

`func (o *AlipayTradeQueryResponse) SetBuyerUserType(v string)`

SetBuyerUserType sets BuyerUserType field to given value.

### HasBuyerUserType

`func (o *AlipayTradeQueryResponse) HasBuyerUserType() bool`

HasBuyerUserType returns a boolean if a field has been set.

### SetBuyerUserTypeNil

`func (o *AlipayTradeQueryResponse) SetBuyerUserTypeNil(b bool)`

 SetBuyerUserTypeNil sets the value for BuyerUserType to be an explicit nil

### UnsetBuyerUserType
`func (o *AlipayTradeQueryResponse) UnsetBuyerUserType()`

UnsetBuyerUserType ensures that no value is present for BuyerUserType, not even an explicit nil
### GetCashierType

`func (o *AlipayTradeQueryResponse) GetCashierType() string`

GetCashierType returns the CashierType field if non-nil, zero value otherwise.

### GetCashierTypeOk

`func (o *AlipayTradeQueryResponse) GetCashierTypeOk() (*string, bool)`

GetCashierTypeOk returns a tuple with the CashierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierType

`func (o *AlipayTradeQueryResponse) SetCashierType(v string)`

SetCashierType sets CashierType field to given value.

### HasCashierType

`func (o *AlipayTradeQueryResponse) HasCashierType() bool`

HasCashierType returns a boolean if a field has been set.

### SetCashierTypeNil

`func (o *AlipayTradeQueryResponse) SetCashierTypeNil(b bool)`

 SetCashierTypeNil sets the value for CashierType to be an explicit nil

### UnsetCashierType
`func (o *AlipayTradeQueryResponse) UnsetCashierType()`

UnsetCashierType ensures that no value is present for CashierType, not even an explicit nil
### GetChargeAmount

`func (o *AlipayTradeQueryResponse) GetChargeAmount() string`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *AlipayTradeQueryResponse) GetChargeAmountOk() (*string, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *AlipayTradeQueryResponse) SetChargeAmount(v string)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *AlipayTradeQueryResponse) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### SetChargeAmountNil

`func (o *AlipayTradeQueryResponse) SetChargeAmountNil(b bool)`

 SetChargeAmountNil sets the value for ChargeAmount to be an explicit nil

### UnsetChargeAmount
`func (o *AlipayTradeQueryResponse) UnsetChargeAmount()`

UnsetChargeAmount ensures that no value is present for ChargeAmount, not even an explicit nil
### GetChargeFlags

`func (o *AlipayTradeQueryResponse) GetChargeFlags() string`

GetChargeFlags returns the ChargeFlags field if non-nil, zero value otherwise.

### GetChargeFlagsOk

`func (o *AlipayTradeQueryResponse) GetChargeFlagsOk() (*string, bool)`

GetChargeFlagsOk returns a tuple with the ChargeFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeFlags

`func (o *AlipayTradeQueryResponse) SetChargeFlags(v string)`

SetChargeFlags sets ChargeFlags field to given value.

### HasChargeFlags

`func (o *AlipayTradeQueryResponse) HasChargeFlags() bool`

HasChargeFlags returns a boolean if a field has been set.

### SetChargeFlagsNil

`func (o *AlipayTradeQueryResponse) SetChargeFlagsNil(b bool)`

 SetChargeFlagsNil sets the value for ChargeFlags to be an explicit nil

### UnsetChargeFlags
`func (o *AlipayTradeQueryResponse) UnsetChargeFlags()`

UnsetChargeFlags ensures that no value is present for ChargeFlags, not even an explicit nil
### GetChargeInfoList

`func (o *AlipayTradeQueryResponse) GetChargeInfoList() []ChargeInfo`

GetChargeInfoList returns the ChargeInfoList field if non-nil, zero value otherwise.

### GetChargeInfoListOk

`func (o *AlipayTradeQueryResponse) GetChargeInfoListOk() (*[]ChargeInfo, bool)`

GetChargeInfoListOk returns a tuple with the ChargeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeInfoList

`func (o *AlipayTradeQueryResponse) SetChargeInfoList(v []ChargeInfo)`

SetChargeInfoList sets ChargeInfoList field to given value.

### HasChargeInfoList

`func (o *AlipayTradeQueryResponse) HasChargeInfoList() bool`

HasChargeInfoList returns a boolean if a field has been set.

### SetChargeInfoListNil

`func (o *AlipayTradeQueryResponse) SetChargeInfoListNil(b bool)`

 SetChargeInfoListNil sets the value for ChargeInfoList to be an explicit nil

### UnsetChargeInfoList
`func (o *AlipayTradeQueryResponse) UnsetChargeInfoList()`

UnsetChargeInfoList ensures that no value is present for ChargeInfoList, not even an explicit nil
### GetCreditBizOrderId

`func (o *AlipayTradeQueryResponse) GetCreditBizOrderId() string`

GetCreditBizOrderId returns the CreditBizOrderId field if non-nil, zero value otherwise.

### GetCreditBizOrderIdOk

`func (o *AlipayTradeQueryResponse) GetCreditBizOrderIdOk() (*string, bool)`

GetCreditBizOrderIdOk returns a tuple with the CreditBizOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBizOrderId

`func (o *AlipayTradeQueryResponse) SetCreditBizOrderId(v string)`

SetCreditBizOrderId sets CreditBizOrderId field to given value.

### HasCreditBizOrderId

`func (o *AlipayTradeQueryResponse) HasCreditBizOrderId() bool`

HasCreditBizOrderId returns a boolean if a field has been set.

### SetCreditBizOrderIdNil

`func (o *AlipayTradeQueryResponse) SetCreditBizOrderIdNil(b bool)`

 SetCreditBizOrderIdNil sets the value for CreditBizOrderId to be an explicit nil

### UnsetCreditBizOrderId
`func (o *AlipayTradeQueryResponse) UnsetCreditBizOrderId()`

UnsetCreditBizOrderId ensures that no value is present for CreditBizOrderId, not even an explicit nil
### GetCreditPayMode

`func (o *AlipayTradeQueryResponse) GetCreditPayMode() string`

GetCreditPayMode returns the CreditPayMode field if non-nil, zero value otherwise.

### GetCreditPayModeOk

`func (o *AlipayTradeQueryResponse) GetCreditPayModeOk() (*string, bool)`

GetCreditPayModeOk returns a tuple with the CreditPayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPayMode

`func (o *AlipayTradeQueryResponse) SetCreditPayMode(v string)`

SetCreditPayMode sets CreditPayMode field to given value.

### HasCreditPayMode

`func (o *AlipayTradeQueryResponse) HasCreditPayMode() bool`

HasCreditPayMode returns a boolean if a field has been set.

### SetCreditPayModeNil

`func (o *AlipayTradeQueryResponse) SetCreditPayModeNil(b bool)`

 SetCreditPayModeNil sets the value for CreditPayMode to be an explicit nil

### UnsetCreditPayMode
`func (o *AlipayTradeQueryResponse) UnsetCreditPayMode()`

UnsetCreditPayMode ensures that no value is present for CreditPayMode, not even an explicit nil
### GetDiscountAmount

`func (o *AlipayTradeQueryResponse) GetDiscountAmount() string`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *AlipayTradeQueryResponse) GetDiscountAmountOk() (*string, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *AlipayTradeQueryResponse) SetDiscountAmount(v string)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *AlipayTradeQueryResponse) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### SetDiscountAmountNil

`func (o *AlipayTradeQueryResponse) SetDiscountAmountNil(b bool)`

 SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil

### UnsetDiscountAmount
`func (o *AlipayTradeQueryResponse) UnsetDiscountAmount()`

UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
### GetDiscountGoodsDetail

`func (o *AlipayTradeQueryResponse) GetDiscountGoodsDetail() string`

GetDiscountGoodsDetail returns the DiscountGoodsDetail field if non-nil, zero value otherwise.

### GetDiscountGoodsDetailOk

`func (o *AlipayTradeQueryResponse) GetDiscountGoodsDetailOk() (*string, bool)`

GetDiscountGoodsDetailOk returns a tuple with the DiscountGoodsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountGoodsDetail

`func (o *AlipayTradeQueryResponse) SetDiscountGoodsDetail(v string)`

SetDiscountGoodsDetail sets DiscountGoodsDetail field to given value.

### HasDiscountGoodsDetail

`func (o *AlipayTradeQueryResponse) HasDiscountGoodsDetail() bool`

HasDiscountGoodsDetail returns a boolean if a field has been set.

### SetDiscountGoodsDetailNil

`func (o *AlipayTradeQueryResponse) SetDiscountGoodsDetailNil(b bool)`

 SetDiscountGoodsDetailNil sets the value for DiscountGoodsDetail to be an explicit nil

### UnsetDiscountGoodsDetail
`func (o *AlipayTradeQueryResponse) UnsetDiscountGoodsDetail()`

UnsetDiscountGoodsDetail ensures that no value is present for DiscountGoodsDetail, not even an explicit nil
### GetEnterprisePayInfo

`func (o *AlipayTradeQueryResponse) GetEnterprisePayInfo() EnterprisePayInfo`

GetEnterprisePayInfo returns the EnterprisePayInfo field if non-nil, zero value otherwise.

### GetEnterprisePayInfoOk

`func (o *AlipayTradeQueryResponse) GetEnterprisePayInfoOk() (*EnterprisePayInfo, bool)`

GetEnterprisePayInfoOk returns a tuple with the EnterprisePayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePayInfo

`func (o *AlipayTradeQueryResponse) SetEnterprisePayInfo(v EnterprisePayInfo)`

SetEnterprisePayInfo sets EnterprisePayInfo field to given value.

### HasEnterprisePayInfo

`func (o *AlipayTradeQueryResponse) HasEnterprisePayInfo() bool`

HasEnterprisePayInfo returns a boolean if a field has been set.

### GetExtInfos

`func (o *AlipayTradeQueryResponse) GetExtInfos() string`

GetExtInfos returns the ExtInfos field if non-nil, zero value otherwise.

### GetExtInfosOk

`func (o *AlipayTradeQueryResponse) GetExtInfosOk() (*string, bool)`

GetExtInfosOk returns a tuple with the ExtInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfos

`func (o *AlipayTradeQueryResponse) SetExtInfos(v string)`

SetExtInfos sets ExtInfos field to given value.

### HasExtInfos

`func (o *AlipayTradeQueryResponse) HasExtInfos() bool`

HasExtInfos returns a boolean if a field has been set.

### SetExtInfosNil

`func (o *AlipayTradeQueryResponse) SetExtInfosNil(b bool)`

 SetExtInfosNil sets the value for ExtInfos to be an explicit nil

### UnsetExtInfos
`func (o *AlipayTradeQueryResponse) UnsetExtInfos()`

UnsetExtInfos ensures that no value is present for ExtInfos, not even an explicit nil
### GetFulfillmentDetailList

`func (o *AlipayTradeQueryResponse) GetFulfillmentDetailList() []FulfillmentDetail`

GetFulfillmentDetailList returns the FulfillmentDetailList field if non-nil, zero value otherwise.

### GetFulfillmentDetailListOk

`func (o *AlipayTradeQueryResponse) GetFulfillmentDetailListOk() (*[]FulfillmentDetail, bool)`

GetFulfillmentDetailListOk returns a tuple with the FulfillmentDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentDetailList

`func (o *AlipayTradeQueryResponse) SetFulfillmentDetailList(v []FulfillmentDetail)`

SetFulfillmentDetailList sets FulfillmentDetailList field to given value.

### HasFulfillmentDetailList

`func (o *AlipayTradeQueryResponse) HasFulfillmentDetailList() bool`

HasFulfillmentDetailList returns a boolean if a field has been set.

### SetFulfillmentDetailListNil

`func (o *AlipayTradeQueryResponse) SetFulfillmentDetailListNil(b bool)`

 SetFulfillmentDetailListNil sets the value for FulfillmentDetailList to be an explicit nil

### UnsetFulfillmentDetailList
`func (o *AlipayTradeQueryResponse) UnsetFulfillmentDetailList()`

UnsetFulfillmentDetailList ensures that no value is present for FulfillmentDetailList, not even an explicit nil
### GetFundBillList

`func (o *AlipayTradeQueryResponse) GetFundBillList() []TradeFundBill`

GetFundBillList returns the FundBillList field if non-nil, zero value otherwise.

### GetFundBillListOk

`func (o *AlipayTradeQueryResponse) GetFundBillListOk() (*[]TradeFundBill, bool)`

GetFundBillListOk returns a tuple with the FundBillList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundBillList

`func (o *AlipayTradeQueryResponse) SetFundBillList(v []TradeFundBill)`

SetFundBillList sets FundBillList field to given value.

### HasFundBillList

`func (o *AlipayTradeQueryResponse) HasFundBillList() bool`

HasFundBillList returns a boolean if a field has been set.

### SetFundBillListNil

`func (o *AlipayTradeQueryResponse) SetFundBillListNil(b bool)`

 SetFundBillListNil sets the value for FundBillList to be an explicit nil

### UnsetFundBillList
`func (o *AlipayTradeQueryResponse) UnsetFundBillList()`

UnsetFundBillList ensures that no value is present for FundBillList, not even an explicit nil
### GetHbFqPayInfo

`func (o *AlipayTradeQueryResponse) GetHbFqPayInfo() HbFqPayInfo`

GetHbFqPayInfo returns the HbFqPayInfo field if non-nil, zero value otherwise.

### GetHbFqPayInfoOk

`func (o *AlipayTradeQueryResponse) GetHbFqPayInfoOk() (*HbFqPayInfo, bool)`

GetHbFqPayInfoOk returns a tuple with the HbFqPayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbFqPayInfo

`func (o *AlipayTradeQueryResponse) SetHbFqPayInfo(v HbFqPayInfo)`

SetHbFqPayInfo sets HbFqPayInfo field to given value.

### HasHbFqPayInfo

`func (o *AlipayTradeQueryResponse) HasHbFqPayInfo() bool`

HasHbFqPayInfo returns a boolean if a field has been set.

### GetHybAmount

`func (o *AlipayTradeQueryResponse) GetHybAmount() string`

GetHybAmount returns the HybAmount field if non-nil, zero value otherwise.

### GetHybAmountOk

`func (o *AlipayTradeQueryResponse) GetHybAmountOk() (*string, bool)`

GetHybAmountOk returns a tuple with the HybAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybAmount

`func (o *AlipayTradeQueryResponse) SetHybAmount(v string)`

SetHybAmount sets HybAmount field to given value.

### HasHybAmount

`func (o *AlipayTradeQueryResponse) HasHybAmount() bool`

HasHybAmount returns a boolean if a field has been set.

### SetHybAmountNil

`func (o *AlipayTradeQueryResponse) SetHybAmountNil(b bool)`

 SetHybAmountNil sets the value for HybAmount to be an explicit nil

### UnsetHybAmount
`func (o *AlipayTradeQueryResponse) UnsetHybAmount()`

UnsetHybAmount ensures that no value is present for HybAmount, not even an explicit nil
### GetIndustrySepcDetail

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetail() string`

GetIndustrySepcDetail returns the IndustrySepcDetail field if non-nil, zero value otherwise.

### GetIndustrySepcDetailOk

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetailOk() (*string, bool)`

GetIndustrySepcDetailOk returns a tuple with the IndustrySepcDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustrySepcDetail

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetail(v string)`

SetIndustrySepcDetail sets IndustrySepcDetail field to given value.

### HasIndustrySepcDetail

`func (o *AlipayTradeQueryResponse) HasIndustrySepcDetail() bool`

HasIndustrySepcDetail returns a boolean if a field has been set.

### SetIndustrySepcDetailNil

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetailNil(b bool)`

 SetIndustrySepcDetailNil sets the value for IndustrySepcDetail to be an explicit nil

### UnsetIndustrySepcDetail
`func (o *AlipayTradeQueryResponse) UnsetIndustrySepcDetail()`

UnsetIndustrySepcDetail ensures that no value is present for IndustrySepcDetail, not even an explicit nil
### GetIndustrySepcDetailAcc

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetailAcc() string`

GetIndustrySepcDetailAcc returns the IndustrySepcDetailAcc field if non-nil, zero value otherwise.

### GetIndustrySepcDetailAccOk

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetailAccOk() (*string, bool)`

GetIndustrySepcDetailAccOk returns a tuple with the IndustrySepcDetailAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustrySepcDetailAcc

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetailAcc(v string)`

SetIndustrySepcDetailAcc sets IndustrySepcDetailAcc field to given value.

### HasIndustrySepcDetailAcc

`func (o *AlipayTradeQueryResponse) HasIndustrySepcDetailAcc() bool`

HasIndustrySepcDetailAcc returns a boolean if a field has been set.

### SetIndustrySepcDetailAccNil

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetailAccNil(b bool)`

 SetIndustrySepcDetailAccNil sets the value for IndustrySepcDetailAcc to be an explicit nil

### UnsetIndustrySepcDetailAcc
`func (o *AlipayTradeQueryResponse) UnsetIndustrySepcDetailAcc()`

UnsetIndustrySepcDetailAcc ensures that no value is present for IndustrySepcDetailAcc, not even an explicit nil
### GetIndustrySepcDetailGov

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetailGov() string`

GetIndustrySepcDetailGov returns the IndustrySepcDetailGov field if non-nil, zero value otherwise.

### GetIndustrySepcDetailGovOk

`func (o *AlipayTradeQueryResponse) GetIndustrySepcDetailGovOk() (*string, bool)`

GetIndustrySepcDetailGovOk returns a tuple with the IndustrySepcDetailGov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustrySepcDetailGov

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetailGov(v string)`

SetIndustrySepcDetailGov sets IndustrySepcDetailGov field to given value.

### HasIndustrySepcDetailGov

`func (o *AlipayTradeQueryResponse) HasIndustrySepcDetailGov() bool`

HasIndustrySepcDetailGov returns a boolean if a field has been set.

### SetIndustrySepcDetailGovNil

`func (o *AlipayTradeQueryResponse) SetIndustrySepcDetailGovNil(b bool)`

 SetIndustrySepcDetailGovNil sets the value for IndustrySepcDetailGov to be an explicit nil

### UnsetIndustrySepcDetailGov
`func (o *AlipayTradeQueryResponse) UnsetIndustrySepcDetailGov()`

UnsetIndustrySepcDetailGov ensures that no value is present for IndustrySepcDetailGov, not even an explicit nil
### GetIntactChargeInfoList

`func (o *AlipayTradeQueryResponse) GetIntactChargeInfoList() []IntactChargeInfo`

GetIntactChargeInfoList returns the IntactChargeInfoList field if non-nil, zero value otherwise.

### GetIntactChargeInfoListOk

`func (o *AlipayTradeQueryResponse) GetIntactChargeInfoListOk() (*[]IntactChargeInfo, bool)`

GetIntactChargeInfoListOk returns a tuple with the IntactChargeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntactChargeInfoList

`func (o *AlipayTradeQueryResponse) SetIntactChargeInfoList(v []IntactChargeInfo)`

SetIntactChargeInfoList sets IntactChargeInfoList field to given value.

### HasIntactChargeInfoList

`func (o *AlipayTradeQueryResponse) HasIntactChargeInfoList() bool`

HasIntactChargeInfoList returns a boolean if a field has been set.

### SetIntactChargeInfoListNil

`func (o *AlipayTradeQueryResponse) SetIntactChargeInfoListNil(b bool)`

 SetIntactChargeInfoListNil sets the value for IntactChargeInfoList to be an explicit nil

### UnsetIntactChargeInfoList
`func (o *AlipayTradeQueryResponse) UnsetIntactChargeInfoList()`

UnsetIntactChargeInfoList ensures that no value is present for IntactChargeInfoList, not even an explicit nil
### GetInvoiceAmount

`func (o *AlipayTradeQueryResponse) GetInvoiceAmount() string`

GetInvoiceAmount returns the InvoiceAmount field if non-nil, zero value otherwise.

### GetInvoiceAmountOk

`func (o *AlipayTradeQueryResponse) GetInvoiceAmountOk() (*string, bool)`

GetInvoiceAmountOk returns a tuple with the InvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAmount

`func (o *AlipayTradeQueryResponse) SetInvoiceAmount(v string)`

SetInvoiceAmount sets InvoiceAmount field to given value.

### HasInvoiceAmount

`func (o *AlipayTradeQueryResponse) HasInvoiceAmount() bool`

HasInvoiceAmount returns a boolean if a field has been set.

### SetInvoiceAmountNil

`func (o *AlipayTradeQueryResponse) SetInvoiceAmountNil(b bool)`

 SetInvoiceAmountNil sets the value for InvoiceAmount to be an explicit nil

### UnsetInvoiceAmount
`func (o *AlipayTradeQueryResponse) UnsetInvoiceAmount()`

UnsetInvoiceAmount ensures that no value is present for InvoiceAmount, not even an explicit nil
### GetMdiscountAmount

`func (o *AlipayTradeQueryResponse) GetMdiscountAmount() string`

GetMdiscountAmount returns the MdiscountAmount field if non-nil, zero value otherwise.

### GetMdiscountAmountOk

`func (o *AlipayTradeQueryResponse) GetMdiscountAmountOk() (*string, bool)`

GetMdiscountAmountOk returns a tuple with the MdiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdiscountAmount

`func (o *AlipayTradeQueryResponse) SetMdiscountAmount(v string)`

SetMdiscountAmount sets MdiscountAmount field to given value.

### HasMdiscountAmount

`func (o *AlipayTradeQueryResponse) HasMdiscountAmount() bool`

HasMdiscountAmount returns a boolean if a field has been set.

### SetMdiscountAmountNil

`func (o *AlipayTradeQueryResponse) SetMdiscountAmountNil(b bool)`

 SetMdiscountAmountNil sets the value for MdiscountAmount to be an explicit nil

### UnsetMdiscountAmount
`func (o *AlipayTradeQueryResponse) UnsetMdiscountAmount()`

UnsetMdiscountAmount ensures that no value is present for MdiscountAmount, not even an explicit nil
### GetMedicalInsuranceInfo

`func (o *AlipayTradeQueryResponse) GetMedicalInsuranceInfo() string`

GetMedicalInsuranceInfo returns the MedicalInsuranceInfo field if non-nil, zero value otherwise.

### GetMedicalInsuranceInfoOk

`func (o *AlipayTradeQueryResponse) GetMedicalInsuranceInfoOk() (*string, bool)`

GetMedicalInsuranceInfoOk returns a tuple with the MedicalInsuranceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedicalInsuranceInfo

`func (o *AlipayTradeQueryResponse) SetMedicalInsuranceInfo(v string)`

SetMedicalInsuranceInfo sets MedicalInsuranceInfo field to given value.

### HasMedicalInsuranceInfo

`func (o *AlipayTradeQueryResponse) HasMedicalInsuranceInfo() bool`

HasMedicalInsuranceInfo returns a boolean if a field has been set.

### SetMedicalInsuranceInfoNil

`func (o *AlipayTradeQueryResponse) SetMedicalInsuranceInfoNil(b bool)`

 SetMedicalInsuranceInfoNil sets the value for MedicalInsuranceInfo to be an explicit nil

### UnsetMedicalInsuranceInfo
`func (o *AlipayTradeQueryResponse) UnsetMedicalInsuranceInfo()`

UnsetMedicalInsuranceInfo ensures that no value is present for MedicalInsuranceInfo, not even an explicit nil
### GetOpenId

`func (o *AlipayTradeQueryResponse) GetOpenId() string`

GetOpenId returns the OpenId field if non-nil, zero value otherwise.

### GetOpenIdOk

`func (o *AlipayTradeQueryResponse) GetOpenIdOk() (*string, bool)`

GetOpenIdOk returns a tuple with the OpenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenId

`func (o *AlipayTradeQueryResponse) SetOpenId(v string)`

SetOpenId sets OpenId field to given value.

### HasOpenId

`func (o *AlipayTradeQueryResponse) HasOpenId() bool`

HasOpenId returns a boolean if a field has been set.

### SetOpenIdNil

`func (o *AlipayTradeQueryResponse) SetOpenIdNil(b bool)`

 SetOpenIdNil sets the value for OpenId to be an explicit nil

### UnsetOpenId
`func (o *AlipayTradeQueryResponse) UnsetOpenId()`

UnsetOpenId ensures that no value is present for OpenId, not even an explicit nil
### GetOutTradeNo

`func (o *AlipayTradeQueryResponse) GetOutTradeNo() string`

GetOutTradeNo returns the OutTradeNo field if non-nil, zero value otherwise.

### GetOutTradeNoOk

`func (o *AlipayTradeQueryResponse) GetOutTradeNoOk() (*string, bool)`

GetOutTradeNoOk returns a tuple with the OutTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTradeNo

`func (o *AlipayTradeQueryResponse) SetOutTradeNo(v string)`

SetOutTradeNo sets OutTradeNo field to given value.

### HasOutTradeNo

`func (o *AlipayTradeQueryResponse) HasOutTradeNo() bool`

HasOutTradeNo returns a boolean if a field has been set.

### SetOutTradeNoNil

`func (o *AlipayTradeQueryResponse) SetOutTradeNoNil(b bool)`

 SetOutTradeNoNil sets the value for OutTradeNo to be an explicit nil

### UnsetOutTradeNo
`func (o *AlipayTradeQueryResponse) UnsetOutTradeNo()`

UnsetOutTradeNo ensures that no value is present for OutTradeNo, not even an explicit nil
### GetPassbackParams

`func (o *AlipayTradeQueryResponse) GetPassbackParams() string`

GetPassbackParams returns the PassbackParams field if non-nil, zero value otherwise.

### GetPassbackParamsOk

`func (o *AlipayTradeQueryResponse) GetPassbackParamsOk() (*string, bool)`

GetPassbackParamsOk returns a tuple with the PassbackParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassbackParams

`func (o *AlipayTradeQueryResponse) SetPassbackParams(v string)`

SetPassbackParams sets PassbackParams field to given value.

### HasPassbackParams

`func (o *AlipayTradeQueryResponse) HasPassbackParams() bool`

HasPassbackParams returns a boolean if a field has been set.

### SetPassbackParamsNil

`func (o *AlipayTradeQueryResponse) SetPassbackParamsNil(b bool)`

 SetPassbackParamsNil sets the value for PassbackParams to be an explicit nil

### UnsetPassbackParams
`func (o *AlipayTradeQueryResponse) UnsetPassbackParams()`

UnsetPassbackParams ensures that no value is present for PassbackParams, not even an explicit nil
### GetPayAmount

`func (o *AlipayTradeQueryResponse) GetPayAmount() string`

GetPayAmount returns the PayAmount field if non-nil, zero value otherwise.

### GetPayAmountOk

`func (o *AlipayTradeQueryResponse) GetPayAmountOk() (*string, bool)`

GetPayAmountOk returns a tuple with the PayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayAmount

`func (o *AlipayTradeQueryResponse) SetPayAmount(v string)`

SetPayAmount sets PayAmount field to given value.

### HasPayAmount

`func (o *AlipayTradeQueryResponse) HasPayAmount() bool`

HasPayAmount returns a boolean if a field has been set.

### SetPayAmountNil

`func (o *AlipayTradeQueryResponse) SetPayAmountNil(b bool)`

 SetPayAmountNil sets the value for PayAmount to be an explicit nil

### UnsetPayAmount
`func (o *AlipayTradeQueryResponse) UnsetPayAmount()`

UnsetPayAmount ensures that no value is present for PayAmount, not even an explicit nil
### GetPayCurrency

`func (o *AlipayTradeQueryResponse) GetPayCurrency() string`

GetPayCurrency returns the PayCurrency field if non-nil, zero value otherwise.

### GetPayCurrencyOk

`func (o *AlipayTradeQueryResponse) GetPayCurrencyOk() (*string, bool)`

GetPayCurrencyOk returns a tuple with the PayCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayCurrency

`func (o *AlipayTradeQueryResponse) SetPayCurrency(v string)`

SetPayCurrency sets PayCurrency field to given value.

### HasPayCurrency

`func (o *AlipayTradeQueryResponse) HasPayCurrency() bool`

HasPayCurrency returns a boolean if a field has been set.

### SetPayCurrencyNil

`func (o *AlipayTradeQueryResponse) SetPayCurrencyNil(b bool)`

 SetPayCurrencyNil sets the value for PayCurrency to be an explicit nil

### UnsetPayCurrency
`func (o *AlipayTradeQueryResponse) UnsetPayCurrency()`

UnsetPayCurrency ensures that no value is present for PayCurrency, not even an explicit nil
### GetPaymentInfoWithIdList

`func (o *AlipayTradeQueryResponse) GetPaymentInfoWithIdList() []PaymentInfoWithId`

GetPaymentInfoWithIdList returns the PaymentInfoWithIdList field if non-nil, zero value otherwise.

### GetPaymentInfoWithIdListOk

`func (o *AlipayTradeQueryResponse) GetPaymentInfoWithIdListOk() (*[]PaymentInfoWithId, bool)`

GetPaymentInfoWithIdListOk returns a tuple with the PaymentInfoWithIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoWithIdList

`func (o *AlipayTradeQueryResponse) SetPaymentInfoWithIdList(v []PaymentInfoWithId)`

SetPaymentInfoWithIdList sets PaymentInfoWithIdList field to given value.

### HasPaymentInfoWithIdList

`func (o *AlipayTradeQueryResponse) HasPaymentInfoWithIdList() bool`

HasPaymentInfoWithIdList returns a boolean if a field has been set.

### SetPaymentInfoWithIdListNil

`func (o *AlipayTradeQueryResponse) SetPaymentInfoWithIdListNil(b bool)`

 SetPaymentInfoWithIdListNil sets the value for PaymentInfoWithIdList to be an explicit nil

### UnsetPaymentInfoWithIdList
`func (o *AlipayTradeQueryResponse) UnsetPaymentInfoWithIdList()`

UnsetPaymentInfoWithIdList ensures that no value is present for PaymentInfoWithIdList, not even an explicit nil
### GetPeriodScene

`func (o *AlipayTradeQueryResponse) GetPeriodScene() string`

GetPeriodScene returns the PeriodScene field if non-nil, zero value otherwise.

### GetPeriodSceneOk

`func (o *AlipayTradeQueryResponse) GetPeriodSceneOk() (*string, bool)`

GetPeriodSceneOk returns a tuple with the PeriodScene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodScene

`func (o *AlipayTradeQueryResponse) SetPeriodScene(v string)`

SetPeriodScene sets PeriodScene field to given value.

### HasPeriodScene

`func (o *AlipayTradeQueryResponse) HasPeriodScene() bool`

HasPeriodScene returns a boolean if a field has been set.

### SetPeriodSceneNil

`func (o *AlipayTradeQueryResponse) SetPeriodSceneNil(b bool)`

 SetPeriodSceneNil sets the value for PeriodScene to be an explicit nil

### UnsetPeriodScene
`func (o *AlipayTradeQueryResponse) UnsetPeriodScene()`

UnsetPeriodScene ensures that no value is present for PeriodScene, not even an explicit nil
### GetPointAmount

`func (o *AlipayTradeQueryResponse) GetPointAmount() string`

GetPointAmount returns the PointAmount field if non-nil, zero value otherwise.

### GetPointAmountOk

`func (o *AlipayTradeQueryResponse) GetPointAmountOk() (*string, bool)`

GetPointAmountOk returns a tuple with the PointAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointAmount

`func (o *AlipayTradeQueryResponse) SetPointAmount(v string)`

SetPointAmount sets PointAmount field to given value.

### HasPointAmount

`func (o *AlipayTradeQueryResponse) HasPointAmount() bool`

HasPointAmount returns a boolean if a field has been set.

### SetPointAmountNil

`func (o *AlipayTradeQueryResponse) SetPointAmountNil(b bool)`

 SetPointAmountNil sets the value for PointAmount to be an explicit nil

### UnsetPointAmount
`func (o *AlipayTradeQueryResponse) UnsetPointAmount()`

UnsetPointAmount ensures that no value is present for PointAmount, not even an explicit nil
### GetPreAuthPayAmount

`func (o *AlipayTradeQueryResponse) GetPreAuthPayAmount() string`

GetPreAuthPayAmount returns the PreAuthPayAmount field if non-nil, zero value otherwise.

### GetPreAuthPayAmountOk

`func (o *AlipayTradeQueryResponse) GetPreAuthPayAmountOk() (*string, bool)`

GetPreAuthPayAmountOk returns a tuple with the PreAuthPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthPayAmount

`func (o *AlipayTradeQueryResponse) SetPreAuthPayAmount(v string)`

SetPreAuthPayAmount sets PreAuthPayAmount field to given value.

### HasPreAuthPayAmount

`func (o *AlipayTradeQueryResponse) HasPreAuthPayAmount() bool`

HasPreAuthPayAmount returns a boolean if a field has been set.

### SetPreAuthPayAmountNil

`func (o *AlipayTradeQueryResponse) SetPreAuthPayAmountNil(b bool)`

 SetPreAuthPayAmountNil sets the value for PreAuthPayAmount to be an explicit nil

### UnsetPreAuthPayAmount
`func (o *AlipayTradeQueryResponse) UnsetPreAuthPayAmount()`

UnsetPreAuthPayAmount ensures that no value is present for PreAuthPayAmount, not even an explicit nil
### GetReceiptAmount

`func (o *AlipayTradeQueryResponse) GetReceiptAmount() string`

GetReceiptAmount returns the ReceiptAmount field if non-nil, zero value otherwise.

### GetReceiptAmountOk

`func (o *AlipayTradeQueryResponse) GetReceiptAmountOk() (*string, bool)`

GetReceiptAmountOk returns a tuple with the ReceiptAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptAmount

`func (o *AlipayTradeQueryResponse) SetReceiptAmount(v string)`

SetReceiptAmount sets ReceiptAmount field to given value.

### HasReceiptAmount

`func (o *AlipayTradeQueryResponse) HasReceiptAmount() bool`

HasReceiptAmount returns a boolean if a field has been set.

### SetReceiptAmountNil

`func (o *AlipayTradeQueryResponse) SetReceiptAmountNil(b bool)`

 SetReceiptAmountNil sets the value for ReceiptAmount to be an explicit nil

### UnsetReceiptAmount
`func (o *AlipayTradeQueryResponse) UnsetReceiptAmount()`

UnsetReceiptAmount ensures that no value is present for ReceiptAmount, not even an explicit nil
### GetReceiptCurrencyType

`func (o *AlipayTradeQueryResponse) GetReceiptCurrencyType() string`

GetReceiptCurrencyType returns the ReceiptCurrencyType field if non-nil, zero value otherwise.

### GetReceiptCurrencyTypeOk

`func (o *AlipayTradeQueryResponse) GetReceiptCurrencyTypeOk() (*string, bool)`

GetReceiptCurrencyTypeOk returns a tuple with the ReceiptCurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptCurrencyType

`func (o *AlipayTradeQueryResponse) SetReceiptCurrencyType(v string)`

SetReceiptCurrencyType sets ReceiptCurrencyType field to given value.

### HasReceiptCurrencyType

`func (o *AlipayTradeQueryResponse) HasReceiptCurrencyType() bool`

HasReceiptCurrencyType returns a boolean if a field has been set.

### SetReceiptCurrencyTypeNil

`func (o *AlipayTradeQueryResponse) SetReceiptCurrencyTypeNil(b bool)`

 SetReceiptCurrencyTypeNil sets the value for ReceiptCurrencyType to be an explicit nil

### UnsetReceiptCurrencyType
`func (o *AlipayTradeQueryResponse) UnsetReceiptCurrencyType()`

UnsetReceiptCurrencyType ensures that no value is present for ReceiptCurrencyType, not even an explicit nil
### GetReqGoodsDetail

`func (o *AlipayTradeQueryResponse) GetReqGoodsDetail() []GoodsDetail`

GetReqGoodsDetail returns the ReqGoodsDetail field if non-nil, zero value otherwise.

### GetReqGoodsDetailOk

`func (o *AlipayTradeQueryResponse) GetReqGoodsDetailOk() (*[]GoodsDetail, bool)`

GetReqGoodsDetailOk returns a tuple with the ReqGoodsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGoodsDetail

`func (o *AlipayTradeQueryResponse) SetReqGoodsDetail(v []GoodsDetail)`

SetReqGoodsDetail sets ReqGoodsDetail field to given value.

### HasReqGoodsDetail

`func (o *AlipayTradeQueryResponse) HasReqGoodsDetail() bool`

HasReqGoodsDetail returns a boolean if a field has been set.

### SetReqGoodsDetailNil

`func (o *AlipayTradeQueryResponse) SetReqGoodsDetailNil(b bool)`

 SetReqGoodsDetailNil sets the value for ReqGoodsDetail to be an explicit nil

### UnsetReqGoodsDetail
`func (o *AlipayTradeQueryResponse) UnsetReqGoodsDetail()`

UnsetReqGoodsDetail ensures that no value is present for ReqGoodsDetail, not even an explicit nil
### GetSendPayDate

`func (o *AlipayTradeQueryResponse) GetSendPayDate() string`

GetSendPayDate returns the SendPayDate field if non-nil, zero value otherwise.

### GetSendPayDateOk

`func (o *AlipayTradeQueryResponse) GetSendPayDateOk() (*string, bool)`

GetSendPayDateOk returns a tuple with the SendPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPayDate

`func (o *AlipayTradeQueryResponse) SetSendPayDate(v string)`

SetSendPayDate sets SendPayDate field to given value.

### HasSendPayDate

`func (o *AlipayTradeQueryResponse) HasSendPayDate() bool`

HasSendPayDate returns a boolean if a field has been set.

### SetSendPayDateNil

`func (o *AlipayTradeQueryResponse) SetSendPayDateNil(b bool)`

 SetSendPayDateNil sets the value for SendPayDate to be an explicit nil

### UnsetSendPayDate
`func (o *AlipayTradeQueryResponse) UnsetSendPayDate()`

UnsetSendPayDate ensures that no value is present for SendPayDate, not even an explicit nil
### GetSettleAmount

`func (o *AlipayTradeQueryResponse) GetSettleAmount() string`

GetSettleAmount returns the SettleAmount field if non-nil, zero value otherwise.

### GetSettleAmountOk

`func (o *AlipayTradeQueryResponse) GetSettleAmountOk() (*string, bool)`

GetSettleAmountOk returns a tuple with the SettleAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleAmount

`func (o *AlipayTradeQueryResponse) SetSettleAmount(v string)`

SetSettleAmount sets SettleAmount field to given value.

### HasSettleAmount

`func (o *AlipayTradeQueryResponse) HasSettleAmount() bool`

HasSettleAmount returns a boolean if a field has been set.

### SetSettleAmountNil

`func (o *AlipayTradeQueryResponse) SetSettleAmountNil(b bool)`

 SetSettleAmountNil sets the value for SettleAmount to be an explicit nil

### UnsetSettleAmount
`func (o *AlipayTradeQueryResponse) UnsetSettleAmount()`

UnsetSettleAmount ensures that no value is present for SettleAmount, not even an explicit nil
### GetSettleCurrency

`func (o *AlipayTradeQueryResponse) GetSettleCurrency() string`

GetSettleCurrency returns the SettleCurrency field if non-nil, zero value otherwise.

### GetSettleCurrencyOk

`func (o *AlipayTradeQueryResponse) GetSettleCurrencyOk() (*string, bool)`

GetSettleCurrencyOk returns a tuple with the SettleCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleCurrency

`func (o *AlipayTradeQueryResponse) SetSettleCurrency(v string)`

SetSettleCurrency sets SettleCurrency field to given value.

### HasSettleCurrency

`func (o *AlipayTradeQueryResponse) HasSettleCurrency() bool`

HasSettleCurrency returns a boolean if a field has been set.

### SetSettleCurrencyNil

`func (o *AlipayTradeQueryResponse) SetSettleCurrencyNil(b bool)`

 SetSettleCurrencyNil sets the value for SettleCurrency to be an explicit nil

### UnsetSettleCurrency
`func (o *AlipayTradeQueryResponse) UnsetSettleCurrency()`

UnsetSettleCurrency ensures that no value is present for SettleCurrency, not even an explicit nil
### GetSettleTransRate

`func (o *AlipayTradeQueryResponse) GetSettleTransRate() string`

GetSettleTransRate returns the SettleTransRate field if non-nil, zero value otherwise.

### GetSettleTransRateOk

`func (o *AlipayTradeQueryResponse) GetSettleTransRateOk() (*string, bool)`

GetSettleTransRateOk returns a tuple with the SettleTransRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleTransRate

`func (o *AlipayTradeQueryResponse) SetSettleTransRate(v string)`

SetSettleTransRate sets SettleTransRate field to given value.

### HasSettleTransRate

`func (o *AlipayTradeQueryResponse) HasSettleTransRate() bool`

HasSettleTransRate returns a boolean if a field has been set.

### SetSettleTransRateNil

`func (o *AlipayTradeQueryResponse) SetSettleTransRateNil(b bool)`

 SetSettleTransRateNil sets the value for SettleTransRate to be an explicit nil

### UnsetSettleTransRate
`func (o *AlipayTradeQueryResponse) UnsetSettleTransRate()`

UnsetSettleTransRate ensures that no value is present for SettleTransRate, not even an explicit nil
### GetSettlementId

`func (o *AlipayTradeQueryResponse) GetSettlementId() string`

GetSettlementId returns the SettlementId field if non-nil, zero value otherwise.

### GetSettlementIdOk

`func (o *AlipayTradeQueryResponse) GetSettlementIdOk() (*string, bool)`

GetSettlementIdOk returns a tuple with the SettlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementId

`func (o *AlipayTradeQueryResponse) SetSettlementId(v string)`

SetSettlementId sets SettlementId field to given value.

### HasSettlementId

`func (o *AlipayTradeQueryResponse) HasSettlementId() bool`

HasSettlementId returns a boolean if a field has been set.

### SetSettlementIdNil

`func (o *AlipayTradeQueryResponse) SetSettlementIdNil(b bool)`

 SetSettlementIdNil sets the value for SettlementId to be an explicit nil

### UnsetSettlementId
`func (o *AlipayTradeQueryResponse) UnsetSettlementId()`

UnsetSettlementId ensures that no value is present for SettlementId, not even an explicit nil
### GetStoreId

`func (o *AlipayTradeQueryResponse) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *AlipayTradeQueryResponse) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *AlipayTradeQueryResponse) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *AlipayTradeQueryResponse) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### SetStoreIdNil

`func (o *AlipayTradeQueryResponse) SetStoreIdNil(b bool)`

 SetStoreIdNil sets the value for StoreId to be an explicit nil

### UnsetStoreId
`func (o *AlipayTradeQueryResponse) UnsetStoreId()`

UnsetStoreId ensures that no value is present for StoreId, not even an explicit nil
### GetStoreName

`func (o *AlipayTradeQueryResponse) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *AlipayTradeQueryResponse) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *AlipayTradeQueryResponse) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *AlipayTradeQueryResponse) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### SetStoreNameNil

`func (o *AlipayTradeQueryResponse) SetStoreNameNil(b bool)`

 SetStoreNameNil sets the value for StoreName to be an explicit nil

### UnsetStoreName
`func (o *AlipayTradeQueryResponse) UnsetStoreName()`

UnsetStoreName ensures that no value is present for StoreName, not even an explicit nil
### GetSubject

`func (o *AlipayTradeQueryResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlipayTradeQueryResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlipayTradeQueryResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlipayTradeQueryResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlipayTradeQueryResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlipayTradeQueryResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTapPayInfo

`func (o *AlipayTradeQueryResponse) GetTapPayInfo() TapPayInfo`

GetTapPayInfo returns the TapPayInfo field if non-nil, zero value otherwise.

### GetTapPayInfoOk

`func (o *AlipayTradeQueryResponse) GetTapPayInfoOk() (*TapPayInfo, bool)`

GetTapPayInfoOk returns a tuple with the TapPayInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTapPayInfo

`func (o *AlipayTradeQueryResponse) SetTapPayInfo(v TapPayInfo)`

SetTapPayInfo sets TapPayInfo field to given value.

### HasTapPayInfo

`func (o *AlipayTradeQueryResponse) HasTapPayInfo() bool`

HasTapPayInfo returns a boolean if a field has been set.

### GetTerminalId

`func (o *AlipayTradeQueryResponse) GetTerminalId() string`

GetTerminalId returns the TerminalId field if non-nil, zero value otherwise.

### GetTerminalIdOk

`func (o *AlipayTradeQueryResponse) GetTerminalIdOk() (*string, bool)`

GetTerminalIdOk returns a tuple with the TerminalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalId

`func (o *AlipayTradeQueryResponse) SetTerminalId(v string)`

SetTerminalId sets TerminalId field to given value.

### HasTerminalId

`func (o *AlipayTradeQueryResponse) HasTerminalId() bool`

HasTerminalId returns a boolean if a field has been set.

### SetTerminalIdNil

`func (o *AlipayTradeQueryResponse) SetTerminalIdNil(b bool)`

 SetTerminalIdNil sets the value for TerminalId to be an explicit nil

### UnsetTerminalId
`func (o *AlipayTradeQueryResponse) UnsetTerminalId()`

UnsetTerminalId ensures that no value is present for TerminalId, not even an explicit nil
### GetTotalAmount

`func (o *AlipayTradeQueryResponse) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *AlipayTradeQueryResponse) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *AlipayTradeQueryResponse) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *AlipayTradeQueryResponse) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *AlipayTradeQueryResponse) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *AlipayTradeQueryResponse) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetTradeNo

`func (o *AlipayTradeQueryResponse) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *AlipayTradeQueryResponse) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *AlipayTradeQueryResponse) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.

### HasTradeNo

`func (o *AlipayTradeQueryResponse) HasTradeNo() bool`

HasTradeNo returns a boolean if a field has been set.

### SetTradeNoNil

`func (o *AlipayTradeQueryResponse) SetTradeNoNil(b bool)`

 SetTradeNoNil sets the value for TradeNo to be an explicit nil

### UnsetTradeNo
`func (o *AlipayTradeQueryResponse) UnsetTradeNo()`

UnsetTradeNo ensures that no value is present for TradeNo, not even an explicit nil
### GetTradeSettleInfo

`func (o *AlipayTradeQueryResponse) GetTradeSettleInfo() TradeSettleInfo`

GetTradeSettleInfo returns the TradeSettleInfo field if non-nil, zero value otherwise.

### GetTradeSettleInfoOk

`func (o *AlipayTradeQueryResponse) GetTradeSettleInfoOk() (*TradeSettleInfo, bool)`

GetTradeSettleInfoOk returns a tuple with the TradeSettleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSettleInfo

`func (o *AlipayTradeQueryResponse) SetTradeSettleInfo(v TradeSettleInfo)`

SetTradeSettleInfo sets TradeSettleInfo field to given value.

### HasTradeSettleInfo

`func (o *AlipayTradeQueryResponse) HasTradeSettleInfo() bool`

HasTradeSettleInfo returns a boolean if a field has been set.

### GetTradeStatus

`func (o *AlipayTradeQueryResponse) GetTradeStatus() string`

GetTradeStatus returns the TradeStatus field if non-nil, zero value otherwise.

### GetTradeStatusOk

`func (o *AlipayTradeQueryResponse) GetTradeStatusOk() (*string, bool)`

GetTradeStatusOk returns a tuple with the TradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeStatus

`func (o *AlipayTradeQueryResponse) SetTradeStatus(v string)`

SetTradeStatus sets TradeStatus field to given value.

### HasTradeStatus

`func (o *AlipayTradeQueryResponse) HasTradeStatus() bool`

HasTradeStatus returns a boolean if a field has been set.

### SetTradeStatusNil

`func (o *AlipayTradeQueryResponse) SetTradeStatusNil(b bool)`

 SetTradeStatusNil sets the value for TradeStatus to be an explicit nil

### UnsetTradeStatus
`func (o *AlipayTradeQueryResponse) UnsetTradeStatus()`

UnsetTradeStatus ensures that no value is present for TradeStatus, not even an explicit nil
### GetTransCurrency

`func (o *AlipayTradeQueryResponse) GetTransCurrency() string`

GetTransCurrency returns the TransCurrency field if non-nil, zero value otherwise.

### GetTransCurrencyOk

`func (o *AlipayTradeQueryResponse) GetTransCurrencyOk() (*string, bool)`

GetTransCurrencyOk returns a tuple with the TransCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransCurrency

`func (o *AlipayTradeQueryResponse) SetTransCurrency(v string)`

SetTransCurrency sets TransCurrency field to given value.

### HasTransCurrency

`func (o *AlipayTradeQueryResponse) HasTransCurrency() bool`

HasTransCurrency returns a boolean if a field has been set.

### SetTransCurrencyNil

`func (o *AlipayTradeQueryResponse) SetTransCurrencyNil(b bool)`

 SetTransCurrencyNil sets the value for TransCurrency to be an explicit nil

### UnsetTransCurrency
`func (o *AlipayTradeQueryResponse) UnsetTransCurrency()`

UnsetTransCurrency ensures that no value is present for TransCurrency, not even an explicit nil
### GetTransPayRate

`func (o *AlipayTradeQueryResponse) GetTransPayRate() string`

GetTransPayRate returns the TransPayRate field if non-nil, zero value otherwise.

### GetTransPayRateOk

`func (o *AlipayTradeQueryResponse) GetTransPayRateOk() (*string, bool)`

GetTransPayRateOk returns a tuple with the TransPayRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransPayRate

`func (o *AlipayTradeQueryResponse) SetTransPayRate(v string)`

SetTransPayRate sets TransPayRate field to given value.

### HasTransPayRate

`func (o *AlipayTradeQueryResponse) HasTransPayRate() bool`

HasTransPayRate returns a boolean if a field has been set.

### SetTransPayRateNil

`func (o *AlipayTradeQueryResponse) SetTransPayRateNil(b bool)`

 SetTransPayRateNil sets the value for TransPayRate to be an explicit nil

### UnsetTransPayRate
`func (o *AlipayTradeQueryResponse) UnsetTransPayRate()`

UnsetTransPayRate ensures that no value is present for TransPayRate, not even an explicit nil
### GetVoucherDetailList

`func (o *AlipayTradeQueryResponse) GetVoucherDetailList() []VoucherDetail`

GetVoucherDetailList returns the VoucherDetailList field if non-nil, zero value otherwise.

### GetVoucherDetailListOk

`func (o *AlipayTradeQueryResponse) GetVoucherDetailListOk() (*[]VoucherDetail, bool)`

GetVoucherDetailListOk returns a tuple with the VoucherDetailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherDetailList

`func (o *AlipayTradeQueryResponse) SetVoucherDetailList(v []VoucherDetail)`

SetVoucherDetailList sets VoucherDetailList field to given value.

### HasVoucherDetailList

`func (o *AlipayTradeQueryResponse) HasVoucherDetailList() bool`

HasVoucherDetailList returns a boolean if a field has been set.

### SetVoucherDetailListNil

`func (o *AlipayTradeQueryResponse) SetVoucherDetailListNil(b bool)`

 SetVoucherDetailListNil sets the value for VoucherDetailList to be an explicit nil

### UnsetVoucherDetailList
`func (o *AlipayTradeQueryResponse) UnsetVoucherDetailList()`

UnsetVoucherDetailList ensures that no value is present for VoucherDetailList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


