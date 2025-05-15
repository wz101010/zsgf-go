# App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 唯一标识符 | [optional] 
**UserID** | Pointer to **int64** | 用户 ID | [optional] 
**ProjectID** | Pointer to **int64** | 项目 ID | [optional] 
**Website** | Pointer to **NullableString** | 网站默认域名 | [optional] 
**Name** | Pointer to **NullableString** | 名称 | [optional] 
**Logo** | Pointer to **NullableString** | Logo | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] 
**Tags** | Pointer to **NullableString** | 标签 | [optional] 
**AppKey** | Pointer to **NullableString** | 应用公钥 | [optional] 
**AppSecret** | Pointer to **NullableString** | 应用私密密钥 | [optional] 
**ClientSecret** | Pointer to **NullableString** | 客户端密钥 | [optional] 
**SshPublickey** | Pointer to **NullableString** | SSH公钥 | [optional] 
**Share** | Pointer to **bool** | 是否共享 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新时间 | [optional] 
**Show** | Pointer to **bool** | 是否显示 | [optional] 
**ShowIndex** | Pointer to **int64** | 显示索引 | [optional] 

## Methods

### NewApp

`func NewApp() *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *App) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *App) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *App) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *App) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *App) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetProjectID

`func (o *App) GetProjectID() int64`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *App) GetProjectIDOk() (*int64, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *App) SetProjectID(v int64)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *App) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### GetWebsite

`func (o *App) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *App) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *App) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *App) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *App) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *App) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *App) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *App) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *App) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogo

`func (o *App) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *App) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *App) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *App) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *App) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *App) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *App) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *App) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *App) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *App) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *App) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *App) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *App) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *App) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *App) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAppKey

`func (o *App) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *App) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *App) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *App) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### SetAppKeyNil

`func (o *App) SetAppKeyNil(b bool)`

 SetAppKeyNil sets the value for AppKey to be an explicit nil

### UnsetAppKey
`func (o *App) UnsetAppKey()`

UnsetAppKey ensures that no value is present for AppKey, not even an explicit nil
### GetAppSecret

`func (o *App) GetAppSecret() string`

GetAppSecret returns the AppSecret field if non-nil, zero value otherwise.

### GetAppSecretOk

`func (o *App) GetAppSecretOk() (*string, bool)`

GetAppSecretOk returns a tuple with the AppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSecret

`func (o *App) SetAppSecret(v string)`

SetAppSecret sets AppSecret field to given value.

### HasAppSecret

`func (o *App) HasAppSecret() bool`

HasAppSecret returns a boolean if a field has been set.

### SetAppSecretNil

`func (o *App) SetAppSecretNil(b bool)`

 SetAppSecretNil sets the value for AppSecret to be an explicit nil

### UnsetAppSecret
`func (o *App) UnsetAppSecret()`

UnsetAppSecret ensures that no value is present for AppSecret, not even an explicit nil
### GetClientSecret

`func (o *App) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *App) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *App) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *App) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *App) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *App) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetSshPublickey

`func (o *App) GetSshPublickey() string`

GetSshPublickey returns the SshPublickey field if non-nil, zero value otherwise.

### GetSshPublickeyOk

`func (o *App) GetSshPublickeyOk() (*string, bool)`

GetSshPublickeyOk returns a tuple with the SshPublickey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublickey

`func (o *App) SetSshPublickey(v string)`

SetSshPublickey sets SshPublickey field to given value.

### HasSshPublickey

`func (o *App) HasSshPublickey() bool`

HasSshPublickey returns a boolean if a field has been set.

### SetSshPublickeyNil

`func (o *App) SetSshPublickeyNil(b bool)`

 SetSshPublickeyNil sets the value for SshPublickey to be an explicit nil

### UnsetSshPublickey
`func (o *App) UnsetSshPublickey()`

UnsetSshPublickey ensures that no value is present for SshPublickey, not even an explicit nil
### GetShare

`func (o *App) GetShare() bool`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *App) GetShareOk() (*bool, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *App) SetShare(v bool)`

SetShare sets Share field to given value.

### HasShare

`func (o *App) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetCreateDate

`func (o *App) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *App) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *App) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *App) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *App) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *App) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *App) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *App) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetShow

`func (o *App) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *App) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *App) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *App) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetShowIndex

`func (o *App) GetShowIndex() int64`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *App) GetShowIndexOk() (*int64, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *App) SetShowIndex(v int64)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *App) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


