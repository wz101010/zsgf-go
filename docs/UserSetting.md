# UserSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 用户的唯一标识符。 | [optional] 
**UserID** | **int64** | 关联的用户ID，表示该配置属于哪个用户。 | 
**ProviderCode** | Pointer to **NullableString** | 提供商的唯一代码，用于标识服务提供者。 | [optional] 
**GroupCode** | Pointer to **NullableString** | 组的唯一代码，用于分类设置项。 | [optional] 
**Code** | **string** | 设置项的唯一代码或键名，用于标识具体的配置项。 | 
**Value** | **string** | 设置项的具体值或选项。 | 
**Tags** | Pointer to **NullableString** | 用于对设置项进行分类或标记的标签。 | [optional] 
**Description** | Pointer to **NullableString** | 设置项的详细描述，说明其作用或用途。 | [optional] 
**FrontendUsable** | Pointer to **bool** | 指示该设置项是否在前端界面中可用。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 设置项的创建时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 设置项的最后更新时间。 | [optional] 

## Methods

### NewUserSetting

`func NewUserSetting(userID int64, code string, value string, ) *UserSetting`

NewUserSetting instantiates a new UserSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingWithDefaults

`func NewUserSettingWithDefaults() *UserSetting`

NewUserSettingWithDefaults instantiates a new UserSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSetting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSetting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSetting) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *UserSetting) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UserSetting) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UserSetting) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetProviderCode

`func (o *UserSetting) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *UserSetting) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *UserSetting) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *UserSetting) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### SetProviderCodeNil

`func (o *UserSetting) SetProviderCodeNil(b bool)`

 SetProviderCodeNil sets the value for ProviderCode to be an explicit nil

### UnsetProviderCode
`func (o *UserSetting) UnsetProviderCode()`

UnsetProviderCode ensures that no value is present for ProviderCode, not even an explicit nil
### GetGroupCode

`func (o *UserSetting) GetGroupCode() string`

GetGroupCode returns the GroupCode field if non-nil, zero value otherwise.

### GetGroupCodeOk

`func (o *UserSetting) GetGroupCodeOk() (*string, bool)`

GetGroupCodeOk returns a tuple with the GroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCode

`func (o *UserSetting) SetGroupCode(v string)`

SetGroupCode sets GroupCode field to given value.

### HasGroupCode

`func (o *UserSetting) HasGroupCode() bool`

HasGroupCode returns a boolean if a field has been set.

### SetGroupCodeNil

`func (o *UserSetting) SetGroupCodeNil(b bool)`

 SetGroupCodeNil sets the value for GroupCode to be an explicit nil

### UnsetGroupCode
`func (o *UserSetting) UnsetGroupCode()`

UnsetGroupCode ensures that no value is present for GroupCode, not even an explicit nil
### GetCode

`func (o *UserSetting) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserSetting) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserSetting) SetCode(v string)`

SetCode sets Code field to given value.


### GetValue

`func (o *UserSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserSetting) SetValue(v string)`

SetValue sets Value field to given value.


### GetTags

`func (o *UserSetting) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserSetting) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserSetting) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserSetting) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UserSetting) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UserSetting) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *UserSetting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserSetting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserSetting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserSetting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserSetting) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserSetting) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFrontendUsable

`func (o *UserSetting) GetFrontendUsable() bool`

GetFrontendUsable returns the FrontendUsable field if non-nil, zero value otherwise.

### GetFrontendUsableOk

`func (o *UserSetting) GetFrontendUsableOk() (*bool, bool)`

GetFrontendUsableOk returns a tuple with the FrontendUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendUsable

`func (o *UserSetting) SetFrontendUsable(v bool)`

SetFrontendUsable sets FrontendUsable field to given value.

### HasFrontendUsable

`func (o *UserSetting) HasFrontendUsable() bool`

HasFrontendUsable returns a boolean if a field has been set.

### GetCreateDate

`func (o *UserSetting) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *UserSetting) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *UserSetting) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *UserSetting) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *UserSetting) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserSetting) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *UserSetting) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *UserSetting) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


