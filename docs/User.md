# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 用户唯一标识 | [optional] 
**Platform** | Pointer to **NullableString** | 用户所在平台 | [optional] 
**UnionID** | Pointer to **NullableString** | 用户所在平台的唯一标识 | [optional] 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 其他数据 | [optional] 
**UserName** | Pointer to **NullableString** | 用户名 | [optional] 
**Pwd** | Pointer to **NullableString** | 用户密码 | [optional] 
**Email** | Pointer to **NullableString** | 邮箱地址 | [optional] 
**EmailIsValid** | Pointer to **bool** | 邮箱已验证 | [optional] 
**Phone** | Pointer to **NullableString** | 手机号码 | [optional] 
**PhoneIsValid** | Pointer to **bool** | 手机号码已验证 | [optional] 
**RelationChain** | Pointer to **NullableString** | 关系链 | [optional] 
**InterestTags** | Pointer to **NullableString** | 兴趣标签 | [optional] 
**Biography** | Pointer to **NullableString** | 个人简介 | [optional] 
**Gender** | Pointer to **NullableString** | 性别 | [optional] 
**Birthday** | Pointer to **NullableTime** | 生日 | [optional] 
**Occupation** | Pointer to **NullableString** | 职业 | [optional] 
**Education** | Pointer to **NullableString** | 学历 | [optional] 
**Contact** | Pointer to **NullableString** | 联系方式 | [optional] 
**Languages** | Pointer to **NullableString** | 语言 | [optional] 
**SocialLinks** | Pointer to **NullableString** | 社交网络链接 | [optional] 
**RelationshipStatus** | Pointer to **NullableString** | 婚姻状态 | [optional] 
**Company** | Pointer to **NullableString** | 公司 | [optional] 
**Industry** | Pointer to **NullableString** | 行业 | [optional] 
**CompanyPosition** | Pointer to **NullableString** | 行业职位 | [optional] 
**PrivateSettings** | Pointer to **NullableString** | 私密设置 | [optional] 
**IsLock** | Pointer to **bool** | 账户是否锁定 | [optional] 
**LockUntil** | Pointer to **time.Time** | 账户锁定截止时间 | [optional] 
**EnableUserNameSignIn** | Pointer to **bool** | 能使用用户名登录 | [optional] 
**EnableEmailSignIn** | Pointer to **bool** | 能使用邮箱登录 | [optional] 
**EnablePhoneSignIn** | Pointer to **bool** | 能使用电话号码登录 | [optional] 
**EnableUnionIDSignIn** | Pointer to **bool** | 能使用联合身份标识登录 | [optional] 
**EnableOAuth** | Pointer to **bool** | 能使用OAuth认证方式登录 | [optional] 
**Enable2FAAuth** | Pointer to **bool** | 启用双因素认证登录 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新时间 | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *User) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *User) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *User) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *User) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *User) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *User) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetUnionID

`func (o *User) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *User) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *User) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.

### HasUnionID

`func (o *User) HasUnionID() bool`

HasUnionID returns a boolean if a field has been set.

### SetUnionIDNil

`func (o *User) SetUnionIDNil(b bool)`

 SetUnionIDNil sets the value for UnionID to be an explicit nil

### UnsetUnionID
`func (o *User) UnsetUnionID()`

UnsetUnionID ensures that no value is present for UnionID, not even an explicit nil
### GetNickName

`func (o *User) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *User) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *User) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *User) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *User) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *User) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *User) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *User) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *User) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *User) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *User) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *User) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *User) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *User) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *User) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetUserName

`func (o *User) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *User) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *User) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *User) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *User) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *User) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPwd

`func (o *User) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *User) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *User) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *User) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### SetPwdNil

`func (o *User) SetPwdNil(b bool)`

 SetPwdNil sets the value for Pwd to be an explicit nil

### UnsetPwd
`func (o *User) UnsetPwd()`

UnsetPwd ensures that no value is present for Pwd, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailIsValid

`func (o *User) GetEmailIsValid() bool`

GetEmailIsValid returns the EmailIsValid field if non-nil, zero value otherwise.

### GetEmailIsValidOk

`func (o *User) GetEmailIsValidOk() (*bool, bool)`

GetEmailIsValidOk returns a tuple with the EmailIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIsValid

`func (o *User) SetEmailIsValid(v bool)`

SetEmailIsValid sets EmailIsValid field to given value.

### HasEmailIsValid

`func (o *User) HasEmailIsValid() bool`

HasEmailIsValid returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *User) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *User) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneIsValid

`func (o *User) GetPhoneIsValid() bool`

GetPhoneIsValid returns the PhoneIsValid field if non-nil, zero value otherwise.

### GetPhoneIsValidOk

`func (o *User) GetPhoneIsValidOk() (*bool, bool)`

GetPhoneIsValidOk returns a tuple with the PhoneIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneIsValid

`func (o *User) SetPhoneIsValid(v bool)`

SetPhoneIsValid sets PhoneIsValid field to given value.

### HasPhoneIsValid

`func (o *User) HasPhoneIsValid() bool`

HasPhoneIsValid returns a boolean if a field has been set.

### GetRelationChain

`func (o *User) GetRelationChain() string`

GetRelationChain returns the RelationChain field if non-nil, zero value otherwise.

### GetRelationChainOk

`func (o *User) GetRelationChainOk() (*string, bool)`

GetRelationChainOk returns a tuple with the RelationChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationChain

`func (o *User) SetRelationChain(v string)`

SetRelationChain sets RelationChain field to given value.

### HasRelationChain

`func (o *User) HasRelationChain() bool`

HasRelationChain returns a boolean if a field has been set.

### SetRelationChainNil

`func (o *User) SetRelationChainNil(b bool)`

 SetRelationChainNil sets the value for RelationChain to be an explicit nil

### UnsetRelationChain
`func (o *User) UnsetRelationChain()`

UnsetRelationChain ensures that no value is present for RelationChain, not even an explicit nil
### GetInterestTags

`func (o *User) GetInterestTags() string`

GetInterestTags returns the InterestTags field if non-nil, zero value otherwise.

### GetInterestTagsOk

`func (o *User) GetInterestTagsOk() (*string, bool)`

GetInterestTagsOk returns a tuple with the InterestTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestTags

`func (o *User) SetInterestTags(v string)`

SetInterestTags sets InterestTags field to given value.

### HasInterestTags

`func (o *User) HasInterestTags() bool`

HasInterestTags returns a boolean if a field has been set.

### SetInterestTagsNil

`func (o *User) SetInterestTagsNil(b bool)`

 SetInterestTagsNil sets the value for InterestTags to be an explicit nil

### UnsetInterestTags
`func (o *User) UnsetInterestTags()`

UnsetInterestTags ensures that no value is present for InterestTags, not even an explicit nil
### GetBiography

`func (o *User) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *User) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *User) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *User) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### SetBiographyNil

`func (o *User) SetBiographyNil(b bool)`

 SetBiographyNil sets the value for Biography to be an explicit nil

### UnsetBiography
`func (o *User) UnsetBiography()`

UnsetBiography ensures that no value is present for Biography, not even an explicit nil
### GetGender

`func (o *User) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *User) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *User) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *User) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *User) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *User) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetBirthday

`func (o *User) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *User) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *User) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *User) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *User) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *User) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetOccupation

`func (o *User) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *User) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *User) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *User) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### SetOccupationNil

`func (o *User) SetOccupationNil(b bool)`

 SetOccupationNil sets the value for Occupation to be an explicit nil

### UnsetOccupation
`func (o *User) UnsetOccupation()`

UnsetOccupation ensures that no value is present for Occupation, not even an explicit nil
### GetEducation

`func (o *User) GetEducation() string`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *User) GetEducationOk() (*string, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *User) SetEducation(v string)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *User) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### SetEducationNil

`func (o *User) SetEducationNil(b bool)`

 SetEducationNil sets the value for Education to be an explicit nil

### UnsetEducation
`func (o *User) UnsetEducation()`

UnsetEducation ensures that no value is present for Education, not even an explicit nil
### GetContact

`func (o *User) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *User) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *User) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *User) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *User) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *User) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetLanguages

`func (o *User) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *User) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *User) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *User) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *User) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *User) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetSocialLinks

`func (o *User) GetSocialLinks() string`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *User) GetSocialLinksOk() (*string, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *User) SetSocialLinks(v string)`

SetSocialLinks sets SocialLinks field to given value.

### HasSocialLinks

`func (o *User) HasSocialLinks() bool`

HasSocialLinks returns a boolean if a field has been set.

### SetSocialLinksNil

`func (o *User) SetSocialLinksNil(b bool)`

 SetSocialLinksNil sets the value for SocialLinks to be an explicit nil

### UnsetSocialLinks
`func (o *User) UnsetSocialLinks()`

UnsetSocialLinks ensures that no value is present for SocialLinks, not even an explicit nil
### GetRelationshipStatus

`func (o *User) GetRelationshipStatus() string`

GetRelationshipStatus returns the RelationshipStatus field if non-nil, zero value otherwise.

### GetRelationshipStatusOk

`func (o *User) GetRelationshipStatusOk() (*string, bool)`

GetRelationshipStatusOk returns a tuple with the RelationshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipStatus

`func (o *User) SetRelationshipStatus(v string)`

SetRelationshipStatus sets RelationshipStatus field to given value.

### HasRelationshipStatus

`func (o *User) HasRelationshipStatus() bool`

HasRelationshipStatus returns a boolean if a field has been set.

### SetRelationshipStatusNil

`func (o *User) SetRelationshipStatusNil(b bool)`

 SetRelationshipStatusNil sets the value for RelationshipStatus to be an explicit nil

### UnsetRelationshipStatus
`func (o *User) UnsetRelationshipStatus()`

UnsetRelationshipStatus ensures that no value is present for RelationshipStatus, not even an explicit nil
### GetCompany

`func (o *User) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *User) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *User) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *User) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *User) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *User) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndustry

`func (o *User) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *User) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *User) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *User) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### SetIndustryNil

`func (o *User) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *User) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetCompanyPosition

`func (o *User) GetCompanyPosition() string`

GetCompanyPosition returns the CompanyPosition field if non-nil, zero value otherwise.

### GetCompanyPositionOk

`func (o *User) GetCompanyPositionOk() (*string, bool)`

GetCompanyPositionOk returns a tuple with the CompanyPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPosition

`func (o *User) SetCompanyPosition(v string)`

SetCompanyPosition sets CompanyPosition field to given value.

### HasCompanyPosition

`func (o *User) HasCompanyPosition() bool`

HasCompanyPosition returns a boolean if a field has been set.

### SetCompanyPositionNil

`func (o *User) SetCompanyPositionNil(b bool)`

 SetCompanyPositionNil sets the value for CompanyPosition to be an explicit nil

### UnsetCompanyPosition
`func (o *User) UnsetCompanyPosition()`

UnsetCompanyPosition ensures that no value is present for CompanyPosition, not even an explicit nil
### GetPrivateSettings

`func (o *User) GetPrivateSettings() string`

GetPrivateSettings returns the PrivateSettings field if non-nil, zero value otherwise.

### GetPrivateSettingsOk

`func (o *User) GetPrivateSettingsOk() (*string, bool)`

GetPrivateSettingsOk returns a tuple with the PrivateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSettings

`func (o *User) SetPrivateSettings(v string)`

SetPrivateSettings sets PrivateSettings field to given value.

### HasPrivateSettings

`func (o *User) HasPrivateSettings() bool`

HasPrivateSettings returns a boolean if a field has been set.

### SetPrivateSettingsNil

`func (o *User) SetPrivateSettingsNil(b bool)`

 SetPrivateSettingsNil sets the value for PrivateSettings to be an explicit nil

### UnsetPrivateSettings
`func (o *User) UnsetPrivateSettings()`

UnsetPrivateSettings ensures that no value is present for PrivateSettings, not even an explicit nil
### GetIsLock

`func (o *User) GetIsLock() bool`

GetIsLock returns the IsLock field if non-nil, zero value otherwise.

### GetIsLockOk

`func (o *User) GetIsLockOk() (*bool, bool)`

GetIsLockOk returns a tuple with the IsLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLock

`func (o *User) SetIsLock(v bool)`

SetIsLock sets IsLock field to given value.

### HasIsLock

`func (o *User) HasIsLock() bool`

HasIsLock returns a boolean if a field has been set.

### GetLockUntil

`func (o *User) GetLockUntil() time.Time`

GetLockUntil returns the LockUntil field if non-nil, zero value otherwise.

### GetLockUntilOk

`func (o *User) GetLockUntilOk() (*time.Time, bool)`

GetLockUntilOk returns a tuple with the LockUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockUntil

`func (o *User) SetLockUntil(v time.Time)`

SetLockUntil sets LockUntil field to given value.

### HasLockUntil

`func (o *User) HasLockUntil() bool`

HasLockUntil returns a boolean if a field has been set.

### GetEnableUserNameSignIn

`func (o *User) GetEnableUserNameSignIn() bool`

GetEnableUserNameSignIn returns the EnableUserNameSignIn field if non-nil, zero value otherwise.

### GetEnableUserNameSignInOk

`func (o *User) GetEnableUserNameSignInOk() (*bool, bool)`

GetEnableUserNameSignInOk returns a tuple with the EnableUserNameSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserNameSignIn

`func (o *User) SetEnableUserNameSignIn(v bool)`

SetEnableUserNameSignIn sets EnableUserNameSignIn field to given value.

### HasEnableUserNameSignIn

`func (o *User) HasEnableUserNameSignIn() bool`

HasEnableUserNameSignIn returns a boolean if a field has been set.

### GetEnableEmailSignIn

`func (o *User) GetEnableEmailSignIn() bool`

GetEnableEmailSignIn returns the EnableEmailSignIn field if non-nil, zero value otherwise.

### GetEnableEmailSignInOk

`func (o *User) GetEnableEmailSignInOk() (*bool, bool)`

GetEnableEmailSignInOk returns a tuple with the EnableEmailSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailSignIn

`func (o *User) SetEnableEmailSignIn(v bool)`

SetEnableEmailSignIn sets EnableEmailSignIn field to given value.

### HasEnableEmailSignIn

`func (o *User) HasEnableEmailSignIn() bool`

HasEnableEmailSignIn returns a boolean if a field has been set.

### GetEnablePhoneSignIn

`func (o *User) GetEnablePhoneSignIn() bool`

GetEnablePhoneSignIn returns the EnablePhoneSignIn field if non-nil, zero value otherwise.

### GetEnablePhoneSignInOk

`func (o *User) GetEnablePhoneSignInOk() (*bool, bool)`

GetEnablePhoneSignInOk returns a tuple with the EnablePhoneSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhoneSignIn

`func (o *User) SetEnablePhoneSignIn(v bool)`

SetEnablePhoneSignIn sets EnablePhoneSignIn field to given value.

### HasEnablePhoneSignIn

`func (o *User) HasEnablePhoneSignIn() bool`

HasEnablePhoneSignIn returns a boolean if a field has been set.

### GetEnableUnionIDSignIn

`func (o *User) GetEnableUnionIDSignIn() bool`

GetEnableUnionIDSignIn returns the EnableUnionIDSignIn field if non-nil, zero value otherwise.

### GetEnableUnionIDSignInOk

`func (o *User) GetEnableUnionIDSignInOk() (*bool, bool)`

GetEnableUnionIDSignInOk returns a tuple with the EnableUnionIDSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUnionIDSignIn

`func (o *User) SetEnableUnionIDSignIn(v bool)`

SetEnableUnionIDSignIn sets EnableUnionIDSignIn field to given value.

### HasEnableUnionIDSignIn

`func (o *User) HasEnableUnionIDSignIn() bool`

HasEnableUnionIDSignIn returns a boolean if a field has been set.

### GetEnableOAuth

`func (o *User) GetEnableOAuth() bool`

GetEnableOAuth returns the EnableOAuth field if non-nil, zero value otherwise.

### GetEnableOAuthOk

`func (o *User) GetEnableOAuthOk() (*bool, bool)`

GetEnableOAuthOk returns a tuple with the EnableOAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOAuth

`func (o *User) SetEnableOAuth(v bool)`

SetEnableOAuth sets EnableOAuth field to given value.

### HasEnableOAuth

`func (o *User) HasEnableOAuth() bool`

HasEnableOAuth returns a boolean if a field has been set.

### GetEnable2FAAuth

`func (o *User) GetEnable2FAAuth() bool`

GetEnable2FAAuth returns the Enable2FAAuth field if non-nil, zero value otherwise.

### GetEnable2FAAuthOk

`func (o *User) GetEnable2FAAuthOk() (*bool, bool)`

GetEnable2FAAuthOk returns a tuple with the Enable2FAAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable2FAAuth

`func (o *User) SetEnable2FAAuth(v bool)`

SetEnable2FAAuth sets Enable2FAAuth field to given value.

### HasEnable2FAAuth

`func (o *User) HasEnable2FAAuth() bool`

HasEnable2FAAuth returns a boolean if a field has been set.

### GetCreateDate

`func (o *User) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *User) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *User) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *User) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *User) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *User) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *User) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *User) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


