# VoucherDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Memo** | Pointer to **NullableString** |  | [optional] 
**MerchantContribute** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OtherContribute** | Pointer to **NullableString** |  | [optional] 
**OtherContributeDetail** | Pointer to [**[]ContributeDetail**](ContributeDetail.md) |  | [optional] 
**PurchaseAntContribute** | Pointer to **NullableString** |  | [optional] 
**PurchaseBuyerContribute** | Pointer to **NullableString** |  | [optional] 
**PurchaseMerchantContribute** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoucherDetail

`func NewVoucherDetail() *VoucherDetail`

NewVoucherDetail instantiates a new VoucherDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherDetailWithDefaults

`func NewVoucherDetailWithDefaults() *VoucherDetail`

NewVoucherDetailWithDefaults instantiates a new VoucherDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *VoucherDetail) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VoucherDetail) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VoucherDetail) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VoucherDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VoucherDetail) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VoucherDetail) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetId

`func (o *VoucherDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoucherDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoucherDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoucherDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VoucherDetail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VoucherDetail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMemo

`func (o *VoucherDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *VoucherDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *VoucherDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *VoucherDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *VoucherDetail) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *VoucherDetail) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetMerchantContribute

`func (o *VoucherDetail) GetMerchantContribute() string`

GetMerchantContribute returns the MerchantContribute field if non-nil, zero value otherwise.

### GetMerchantContributeOk

`func (o *VoucherDetail) GetMerchantContributeOk() (*string, bool)`

GetMerchantContributeOk returns a tuple with the MerchantContribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantContribute

`func (o *VoucherDetail) SetMerchantContribute(v string)`

SetMerchantContribute sets MerchantContribute field to given value.

### HasMerchantContribute

`func (o *VoucherDetail) HasMerchantContribute() bool`

HasMerchantContribute returns a boolean if a field has been set.

### SetMerchantContributeNil

`func (o *VoucherDetail) SetMerchantContributeNil(b bool)`

 SetMerchantContributeNil sets the value for MerchantContribute to be an explicit nil

### UnsetMerchantContribute
`func (o *VoucherDetail) UnsetMerchantContribute()`

UnsetMerchantContribute ensures that no value is present for MerchantContribute, not even an explicit nil
### GetName

`func (o *VoucherDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoucherDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoucherDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoucherDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VoucherDetail) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VoucherDetail) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOtherContribute

`func (o *VoucherDetail) GetOtherContribute() string`

GetOtherContribute returns the OtherContribute field if non-nil, zero value otherwise.

### GetOtherContributeOk

`func (o *VoucherDetail) GetOtherContributeOk() (*string, bool)`

GetOtherContributeOk returns a tuple with the OtherContribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherContribute

`func (o *VoucherDetail) SetOtherContribute(v string)`

SetOtherContribute sets OtherContribute field to given value.

### HasOtherContribute

`func (o *VoucherDetail) HasOtherContribute() bool`

HasOtherContribute returns a boolean if a field has been set.

### SetOtherContributeNil

`func (o *VoucherDetail) SetOtherContributeNil(b bool)`

 SetOtherContributeNil sets the value for OtherContribute to be an explicit nil

### UnsetOtherContribute
`func (o *VoucherDetail) UnsetOtherContribute()`

UnsetOtherContribute ensures that no value is present for OtherContribute, not even an explicit nil
### GetOtherContributeDetail

`func (o *VoucherDetail) GetOtherContributeDetail() []ContributeDetail`

GetOtherContributeDetail returns the OtherContributeDetail field if non-nil, zero value otherwise.

### GetOtherContributeDetailOk

`func (o *VoucherDetail) GetOtherContributeDetailOk() (*[]ContributeDetail, bool)`

GetOtherContributeDetailOk returns a tuple with the OtherContributeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherContributeDetail

`func (o *VoucherDetail) SetOtherContributeDetail(v []ContributeDetail)`

SetOtherContributeDetail sets OtherContributeDetail field to given value.

### HasOtherContributeDetail

`func (o *VoucherDetail) HasOtherContributeDetail() bool`

HasOtherContributeDetail returns a boolean if a field has been set.

### SetOtherContributeDetailNil

`func (o *VoucherDetail) SetOtherContributeDetailNil(b bool)`

 SetOtherContributeDetailNil sets the value for OtherContributeDetail to be an explicit nil

### UnsetOtherContributeDetail
`func (o *VoucherDetail) UnsetOtherContributeDetail()`

UnsetOtherContributeDetail ensures that no value is present for OtherContributeDetail, not even an explicit nil
### GetPurchaseAntContribute

`func (o *VoucherDetail) GetPurchaseAntContribute() string`

GetPurchaseAntContribute returns the PurchaseAntContribute field if non-nil, zero value otherwise.

### GetPurchaseAntContributeOk

`func (o *VoucherDetail) GetPurchaseAntContributeOk() (*string, bool)`

GetPurchaseAntContributeOk returns a tuple with the PurchaseAntContribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAntContribute

`func (o *VoucherDetail) SetPurchaseAntContribute(v string)`

SetPurchaseAntContribute sets PurchaseAntContribute field to given value.

### HasPurchaseAntContribute

`func (o *VoucherDetail) HasPurchaseAntContribute() bool`

HasPurchaseAntContribute returns a boolean if a field has been set.

### SetPurchaseAntContributeNil

`func (o *VoucherDetail) SetPurchaseAntContributeNil(b bool)`

 SetPurchaseAntContributeNil sets the value for PurchaseAntContribute to be an explicit nil

### UnsetPurchaseAntContribute
`func (o *VoucherDetail) UnsetPurchaseAntContribute()`

UnsetPurchaseAntContribute ensures that no value is present for PurchaseAntContribute, not even an explicit nil
### GetPurchaseBuyerContribute

`func (o *VoucherDetail) GetPurchaseBuyerContribute() string`

GetPurchaseBuyerContribute returns the PurchaseBuyerContribute field if non-nil, zero value otherwise.

### GetPurchaseBuyerContributeOk

`func (o *VoucherDetail) GetPurchaseBuyerContributeOk() (*string, bool)`

GetPurchaseBuyerContributeOk returns a tuple with the PurchaseBuyerContribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseBuyerContribute

`func (o *VoucherDetail) SetPurchaseBuyerContribute(v string)`

SetPurchaseBuyerContribute sets PurchaseBuyerContribute field to given value.

### HasPurchaseBuyerContribute

`func (o *VoucherDetail) HasPurchaseBuyerContribute() bool`

HasPurchaseBuyerContribute returns a boolean if a field has been set.

### SetPurchaseBuyerContributeNil

`func (o *VoucherDetail) SetPurchaseBuyerContributeNil(b bool)`

 SetPurchaseBuyerContributeNil sets the value for PurchaseBuyerContribute to be an explicit nil

### UnsetPurchaseBuyerContribute
`func (o *VoucherDetail) UnsetPurchaseBuyerContribute()`

UnsetPurchaseBuyerContribute ensures that no value is present for PurchaseBuyerContribute, not even an explicit nil
### GetPurchaseMerchantContribute

`func (o *VoucherDetail) GetPurchaseMerchantContribute() string`

GetPurchaseMerchantContribute returns the PurchaseMerchantContribute field if non-nil, zero value otherwise.

### GetPurchaseMerchantContributeOk

`func (o *VoucherDetail) GetPurchaseMerchantContributeOk() (*string, bool)`

GetPurchaseMerchantContributeOk returns a tuple with the PurchaseMerchantContribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseMerchantContribute

`func (o *VoucherDetail) SetPurchaseMerchantContribute(v string)`

SetPurchaseMerchantContribute sets PurchaseMerchantContribute field to given value.

### HasPurchaseMerchantContribute

`func (o *VoucherDetail) HasPurchaseMerchantContribute() bool`

HasPurchaseMerchantContribute returns a boolean if a field has been set.

### SetPurchaseMerchantContributeNil

`func (o *VoucherDetail) SetPurchaseMerchantContributeNil(b bool)`

 SetPurchaseMerchantContributeNil sets the value for PurchaseMerchantContribute to be an explicit nil

### UnsetPurchaseMerchantContribute
`func (o *VoucherDetail) UnsetPurchaseMerchantContribute()`

UnsetPurchaseMerchantContribute ensures that no value is present for PurchaseMerchantContribute, not even an explicit nil
### GetTemplateId

`func (o *VoucherDetail) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *VoucherDetail) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *VoucherDetail) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *VoucherDetail) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *VoucherDetail) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *VoucherDetail) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetType

`func (o *VoucherDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VoucherDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VoucherDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VoucherDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VoucherDetail) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VoucherDetail) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


