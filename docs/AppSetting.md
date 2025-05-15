# AppSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 唯一标识 | [optional] 
**ProviderCode** | Pointer to **NullableString** | 提供商代码 | [optional] 
**GroupCode** | Pointer to **NullableString** | 分组代码 | [optional] 
**Code** | **string** | 设置项代码 | 
**Value** | **string** | 设置值 | 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] 
**FrontendUsable** | Pointer to **bool** | 是否在前端可用 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新时间 | [optional] 

## Methods

### NewAppSetting

`func NewAppSetting(code string, value string, ) *AppSetting`

NewAppSetting instantiates a new AppSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSettingWithDefaults

`func NewAppSettingWithDefaults() *AppSetting`

NewAppSettingWithDefaults instantiates a new AppSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppSetting) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderCode

`func (o *AppSetting) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *AppSetting) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *AppSetting) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *AppSetting) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *AppSetting) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *AppSetting) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetGroupCode

`func (o *AppSetting) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *AppSetting) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *AppSetting) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *AppSetting) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### SetGroupCodeNil

`func (o *AppSetting) SetGroupCodeNil(b bool)`

 SetGroupCodeNil sets the value for GroupCode to be an explicit nil

### UnsetGroupCode
`func (o *AppSetting) UnsetGroupCode()`

UnsetGroupCode ensures that no value is present for GroupCode, not even an explicit nil
### GetCode

`func (o *AppSetting) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppSetting) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppSetting) SetCode(v string)`

SetCode sets Code field to given value.


### GetValue

`func (o *AppSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppSetting) SetValue(v string)`

SetValue sets Value field to given value.


### GetTags

`func (o *AppSetting) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AppSetting) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AppSetting) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AppSetting) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AppSetting) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AppSetting) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *AppSetting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppSetting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppSetting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppSetting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppSetting) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppSetting) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFrontendUsable

`func (o *AppSetting) GetFrontendUsable() bool`

GetFrontendUsable returns the FrontendUsable field if non-nil, zero value otherwise.

### GetFrontendUsableOk

`func (o *AppSetting) GetFrontendUsableOk() (*bool, bool)`

GetFrontendUsableOk returns a tuple with the FrontendUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendUsable

`func (o *AppSetting) SetFrontendUsable(v bool)`

SetFrontendUsable sets FrontendUsable field to given value.

### HasFrontendUsable

`func (o *AppSetting) HasFrontendUsable() bool`

HasFrontendUsable returns a boolean if a field has been set.

### GetCreateDate

`func (o *AppSetting) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *AppSetting) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *AppSetting) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *AppSetting) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *AppSetting) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *AppSetting) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *AppSetting) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *AppSetting) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


