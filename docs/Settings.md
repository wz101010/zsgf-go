# Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 配置项的唯一标识符。 | [optional] 
**BizCode** | Pointer to **NullableString** | 配置项所属的业务代码，用于分类管理。 | [optional] 
**BizIdentity** | Pointer to **NullableString** | 配置项所属的业务标识，用于唯一标识业务。 | [optional] 
**ProviderCode** | Pointer to **NullableString** | 配置项的提供者代码，用于标识配置来源。 | [optional] 
**GroupCode** | Pointer to **NullableString** | 配置项的分组代码，用于组织和管理相关配置。 | [optional] 
**Code** | Pointer to **NullableString** | 配置项的唯一代码，用于标识具体的配置项。 | [optional] 
**Value** | Pointer to **NullableString** | 配置项的具体值，存储配置内容。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记配置项的标签。 | [optional] 
**Description** | Pointer to **NullableString** | 配置项的详细描述，说明其用途和作用。 | [optional] 
**FrontendUsable** | Pointer to **bool** | 指示该配置项是否可供前端使用。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 配置项的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 配置项的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewSettings

`func NewSettings() *Settings`

NewSettings instantiates a new Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsWithDefaults

`func NewSettingsWithDefaults() *Settings`

NewSettingsWithDefaults instantiates a new Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Settings) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Settings) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Settings) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Settings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBizCode

`func (o *Settings) GetBizCode() string`

GetBizCode returns the BizCode field if non-nil, zero value otherwise.

### GetBizCodeOk

`func (o *Settings) GetBizCodeOk() (*string, bool)`

GetBizCodeOk returns a tuple with the BizCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizCode

`func (o *Settings) SetBizCode(v string)`

SetBizCode sets BizCode field to given value.

### HasBizCode

`func (o *Settings) HasBizCode() bool`

HasBizCode returns a boolean if a field has been set.

### SetBizCodeNil

`func (o *Settings) SetBizCodeNil(b bool)`

 SetBizCodeNil sets the value for BizCode to be an explicit nil

### UnsetBizCode
`func (o *Settings) UnsetBizCode()`

UnsetBizCode ensures that no value is present for BizCode, not even an explicit nil
### GetBizIdentity

`func (o *Settings) GetBizIdentity() string`

GetBizIdentity returns the BizIdentity field if non-nil, zero value otherwise.

### GetBizIdentityOk

`func (o *Settings) GetBizIdentityOk() (*string, bool)`

GetBizIdentityOk returns a tuple with the BizIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizIdentity

`func (o *Settings) SetBizIdentity(v string)`

SetBizIdentity sets BizIdentity field to given value.

### HasBizIdentity

`func (o *Settings) HasBizIdentity() bool`

HasBizIdentity returns a boolean if a field has been set.

### SetBizIdentityNil

`func (o *Settings) SetBizIdentityNil(b bool)`

 SetBizIdentityNil sets the value for BizIdentity to be an explicit nil

### UnsetBizIdentity
`func (o *Settings) UnsetBizIdentity()`

UnsetBizIdentity ensures that no value is present for BizIdentity, not even an explicit nil
### GetProviderCode

`func (o *Settings) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *Settings) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *Settings) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *Settings) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *Settings) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *Settings) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetGroupCode

`func (o *Settings) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *Settings) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *Settings) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *Settings) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### SetGroupCodeNil

`func (o *Settings) SetGroupCodeNil(b bool)`

 SetGroupCodeNil sets the value for GroupCode to be an explicit nil

### UnsetGroupCode
`func (o *Settings) UnsetGroupCode()`

UnsetGroupCode ensures that no value is present for GroupCode, not even an explicit nil
### GetCode

`func (o *Settings) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Settings) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Settings) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Settings) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Settings) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Settings) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetValue

`func (o *Settings) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Settings) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Settings) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Settings) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Settings) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Settings) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTags

`func (o *Settings) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Settings) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Settings) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Settings) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Settings) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Settings) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *Settings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Settings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Settings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Settings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Settings) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Settings) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFrontendUsable

`func (o *Settings) GetFrontendUsable() bool`

GetFrontendUsable returns the FrontendUsable field if non-nil, zero value otherwise.

### GetFrontendUsableOk

`func (o *Settings) GetFrontendUsableOk() (*bool, bool)`

GetFrontendUsableOk returns a tuple with the FrontendUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendUsable

`func (o *Settings) SetFrontendUsable(v bool)`

SetFrontendUsable sets FrontendUsable field to given value.

### HasFrontendUsable

`func (o *Settings) HasFrontendUsable() bool`

HasFrontendUsable returns a boolean if a field has been set.

### GetCreateDate

`func (o *Settings) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Settings) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Settings) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Settings) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *Settings) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Settings) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Settings) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Settings) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


