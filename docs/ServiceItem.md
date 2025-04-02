# ServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 服务配置项的唯一标识符。 | [optional] 
**BizCode** | Pointer to **NullableString** | 服务配置项所属的业务代码，用于分类管理。 | [optional] 
**ProviderCode** | Pointer to **NullableString** | 关联的服务商代码，用于标识提供该配置项的服务商。 | [optional] 
**GroupCode** | Pointer to **NullableString** | 服务配置项所属的功能分组代码，用于组织和管理相关配置项。 | [optional] 
**Name** | Pointer to **NullableString** | 服务配置项的名称，用于描述其功能或用途。 | [optional] 
**Code** | Pointer to **NullableString** | 服务配置项的唯一代码，用于系统内部标识。 | [optional] 
**ValueType** | Pointer to **NullableString** | 服务配置项的值类型，例如 &#39;text&#39;, &#39;number&#39;, &#39;boolean&#39; 等。默认为 &#39;text&#39;。 | [optional] 
**Icon** | Pointer to **NullableString** | 服务配置项的图标URL或路径，用于在界面上显示。 | [optional] 
**ValueDefaults** | Pointer to **NullableString** | 服务配置项的默认值，当未设置具体值时使用。 | [optional] 
**Description** | Pointer to **NullableString** | 服务配置项的详细描述信息，用于说明其用途和配置方法。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记服务配置项的标签。 | [optional] 
**IsSystem** | Pointer to **bool** | 指示该配置项是否为系统级别的配置项，默认为 false。 | [optional] 
**Show** | Pointer to **bool** | 指示服务配置项是否在界面上显示，默认为 true。 | [optional] 
**ShowIndex** | Pointer to **int32** | 服务配置项在界面上的显示顺序。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 服务配置项的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 服务配置项的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewServiceItem

`func NewServiceItem() *ServiceItem`

NewServiceItem instantiates a new ServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceItemWithDefaults

`func NewServiceItemWithDefaults() *ServiceItem`

NewServiceItemWithDefaults instantiates a new ServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBizCode

`func (o *ServiceItem) GetBizCode() string`

GetBizCode returns the BizCode field if non-nil, zero value otherwise.

### GetBizCodeOk

`func (o *ServiceItem) GetBizCodeOk() (*string, bool)`

GetBizCodeOk returns a tuple with the BizCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizCode

`func (o *ServiceItem) SetBizCode(v string)`

SetBizCode sets BizCode field to given value.

### HasBizCode

`func (o *ServiceItem) HasBizCode() bool`

HasBizCode returns a boolean if a field has been set.

### SetBizCodeNil

`func (o *ServiceItem) SetBizCodeNil(b bool)`

 SetBizCodeNil sets the value for BizCode to be an explicit nil

### UnsetBizCode
`func (o *ServiceItem) UnsetBizCode()`

UnsetBizCode ensures that no value is present for BizCode, not even an explicit nil
### GetProviderCode

`func (o *ServiceItem) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *ServiceItem) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *ServiceItem) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *ServiceItem) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *ServiceItem) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *ServiceItem) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetGroupCode

`func (o *ServiceItem) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *ServiceItem) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *ServiceItem) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *ServiceItem) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### SetGroupCodeNil

`func (o *ServiceItem) SetGroupCodeNil(b bool)`

 SetGroupCodeNil sets the value for GroupCode to be an explicit nil

### UnsetGroupCode
`func (o *ServiceItem) UnsetGroupCode()`

UnsetGroupCode ensures that no value is present for GroupCode, not even an explicit nil
### GetName

`func (o *ServiceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServiceItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *ServiceItem) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceItem) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceItem) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceItem) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ServiceItem) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ServiceItem) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetValueType

`func (o *ServiceItem) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ServiceItem) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ServiceItem) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ServiceItem) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### SetValueTypeNil

`func (o *ServiceItem) SetValueTypeNil(b bool)`

 SetValueTypeNil sets the value for ValueType to be an explicit nil

### UnsetValueType
`func (o *ServiceItem) UnsetValueType()`

UnsetValueType ensures that no value is present for ValueType, not even an explicit nil
### GetIcon

`func (o *ServiceItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ServiceItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ServiceItem) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ServiceItem) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ServiceItem) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ServiceItem) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetValueDefaults

`func (o *ServiceItem) GetValueDefaults() string`

GetValueDefaults returns the ValueDefaults field if non-nil, zero value otherwise.

### GetValueDefaultsOk

`func (o *ServiceItem) GetValueDefaultsOk() (*string, bool)`

GetValueDefaultsOk returns a tuple with the ValueDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDefaults

`func (o *ServiceItem) SetValueDefaults(v string)`

SetValueDefaults sets ValueDefaults field to given value.

### HasValueDefaults

`func (o *ServiceItem) HasValueDefaults() bool`

HasValueDefaults returns a boolean if a field has been set.

### SetValueDefaultsNil

`func (o *ServiceItem) SetValueDefaultsNil(b bool)`

 SetValueDefaultsNil sets the value for ValueDefaults to be an explicit nil

### UnsetValueDefaults
`func (o *ServiceItem) UnsetValueDefaults()`

UnsetValueDefaults ensures that no value is present for ValueDefaults, not even an explicit nil
### GetDescription

`func (o *ServiceItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *ServiceItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ServiceItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ServiceItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIsSystem

`func (o *ServiceItem) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ServiceItem) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ServiceItem) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *ServiceItem) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetShow

`func (o *ServiceItem) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *ServiceItem) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *ServiceItem) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *ServiceItem) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetShowIndex

`func (o *ServiceItem) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *ServiceItem) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *ServiceItem) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *ServiceItem) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *ServiceItem) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ServiceItem) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ServiceItem) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ServiceItem) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ServiceItem) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ServiceItem) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ServiceItem) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ServiceItem) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


