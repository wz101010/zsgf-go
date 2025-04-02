# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** | 用户头像的链接或路径 | [optional] 
**Data** | Pointer to **NullableString** | 用户的其他数据，可以是序列化后的对象或JSON字符串 | [optional] 
**NickName** | Pointer to **NullableString** | 用户的昵称 | [optional] 
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

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UpdateProfileRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UpdateProfileRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UpdateProfileRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UpdateProfileRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UpdateProfileRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UpdateProfileRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *UpdateProfileRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateProfileRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateProfileRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateProfileRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UpdateProfileRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UpdateProfileRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetNickName

`func (o *UpdateProfileRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *UpdateProfileRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *UpdateProfileRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *UpdateProfileRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *UpdateProfileRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *UpdateProfileRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetInterestTags

`func (o *UpdateProfileRequest) GetInterestTags() string`

GetInterestTags returns the InterestTags field if non-nil, zero value otherwise.

### GetInterestTagsOk

`func (o *UpdateProfileRequest) GetInterestTagsOk() (*string, bool)`

GetInterestTagsOk returns a tuple with the InterestTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestTags

`func (o *UpdateProfileRequest) SetInterestTags(v string)`

SetInterestTags sets InterestTags field to given value.

### HasInterestTags

`func (o *UpdateProfileRequest) HasInterestTags() bool`

HasInterestTags returns a boolean if a field has been set.

### SetInterestTagsNil

`func (o *UpdateProfileRequest) SetInterestTagsNil(b bool)`

 SetInterestTagsNil sets the value for InterestTags to be an explicit nil

### UnsetInterestTags
`func (o *UpdateProfileRequest) UnsetInterestTags()`

UnsetInterestTags ensures that no value is present for InterestTags, not even an explicit nil
### GetBiography

`func (o *UpdateProfileRequest) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *UpdateProfileRequest) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *UpdateProfileRequest) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *UpdateProfileRequest) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### SetBiographyNil

`func (o *UpdateProfileRequest) SetBiographyNil(b bool)`

 SetBiographyNil sets the value for Biography to be an explicit nil

### UnsetBiography
`func (o *UpdateProfileRequest) UnsetBiography()`

UnsetBiography ensures that no value is present for Biography, not even an explicit nil
### GetGender

`func (o *UpdateProfileRequest) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UpdateProfileRequest) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UpdateProfileRequest) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UpdateProfileRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *UpdateProfileRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *UpdateProfileRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetBirthday

`func (o *UpdateProfileRequest) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UpdateProfileRequest) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UpdateProfileRequest) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UpdateProfileRequest) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *UpdateProfileRequest) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *UpdateProfileRequest) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetOccupation

`func (o *UpdateProfileRequest) GetOccupation() string`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *UpdateProfileRequest) GetOccupationOk() (*string, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *UpdateProfileRequest) SetOccupation(v string)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *UpdateProfileRequest) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### SetOccupationNil

`func (o *UpdateProfileRequest) SetOccupationNil(b bool)`

 SetOccupationNil sets the value for Occupation to be an explicit nil

### UnsetOccupation
`func (o *UpdateProfileRequest) UnsetOccupation()`

UnsetOccupation ensures that no value is present for Occupation, not even an explicit nil
### GetEducation

`func (o *UpdateProfileRequest) GetEducation() string`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *UpdateProfileRequest) GetEducationOk() (*string, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *UpdateProfileRequest) SetEducation(v string)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *UpdateProfileRequest) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### SetEducationNil

`func (o *UpdateProfileRequest) SetEducationNil(b bool)`

 SetEducationNil sets the value for Education to be an explicit nil

### UnsetEducation
`func (o *UpdateProfileRequest) UnsetEducation()`

UnsetEducation ensures that no value is present for Education, not even an explicit nil
### GetContact

`func (o *UpdateProfileRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *UpdateProfileRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *UpdateProfileRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *UpdateProfileRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *UpdateProfileRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *UpdateProfileRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetLanguages

`func (o *UpdateProfileRequest) GetLanguages() string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *UpdateProfileRequest) GetLanguagesOk() (*string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *UpdateProfileRequest) SetLanguages(v string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *UpdateProfileRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *UpdateProfileRequest) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *UpdateProfileRequest) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetSocialLinks

`func (o *UpdateProfileRequest) GetSocialLinks() string`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *UpdateProfileRequest) GetSocialLinksOk() (*string, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *UpdateProfileRequest) SetSocialLinks(v string)`

SetSocialLinks sets SocialLinks field to given value.

### HasSocialLinks

`func (o *UpdateProfileRequest) HasSocialLinks() bool`

HasSocialLinks returns a boolean if a field has been set.

### SetSocialLinksNil

`func (o *UpdateProfileRequest) SetSocialLinksNil(b bool)`

 SetSocialLinksNil sets the value for SocialLinks to be an explicit nil

### UnsetSocialLinks
`func (o *UpdateProfileRequest) UnsetSocialLinks()`

UnsetSocialLinks ensures that no value is present for SocialLinks, not even an explicit nil
### GetRelationshipStatus

`func (o *UpdateProfileRequest) GetRelationshipStatus() string`

GetRelationshipStatus returns the RelationshipStatus field if non-nil, zero value otherwise.

### GetRelationshipStatusOk

`func (o *UpdateProfileRequest) GetRelationshipStatusOk() (*string, bool)`

GetRelationshipStatusOk returns a tuple with the RelationshipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipStatus

`func (o *UpdateProfileRequest) SetRelationshipStatus(v string)`

SetRelationshipStatus sets RelationshipStatus field to given value.

### HasRelationshipStatus

`func (o *UpdateProfileRequest) HasRelationshipStatus() bool`

HasRelationshipStatus returns a boolean if a field has been set.

### SetRelationshipStatusNil

`func (o *UpdateProfileRequest) SetRelationshipStatusNil(b bool)`

 SetRelationshipStatusNil sets the value for RelationshipStatus to be an explicit nil

### UnsetRelationshipStatus
`func (o *UpdateProfileRequest) UnsetRelationshipStatus()`

UnsetRelationshipStatus ensures that no value is present for RelationshipStatus, not even an explicit nil
### GetCompany

`func (o *UpdateProfileRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateProfileRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateProfileRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateProfileRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *UpdateProfileRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *UpdateProfileRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndustry

`func (o *UpdateProfileRequest) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *UpdateProfileRequest) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *UpdateProfileRequest) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *UpdateProfileRequest) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### SetIndustryNil

`func (o *UpdateProfileRequest) SetIndustryNil(b bool)`

 SetIndustryNil sets the value for Industry to be an explicit nil

### UnsetIndustry
`func (o *UpdateProfileRequest) UnsetIndustry()`

UnsetIndustry ensures that no value is present for Industry, not even an explicit nil
### GetCompanyPosition

`func (o *UpdateProfileRequest) GetCompanyPosition() string`

GetCompanyPosition returns the CompanyPosition field if non-nil, zero value otherwise.

### GetCompanyPositionOk

`func (o *UpdateProfileRequest) GetCompanyPositionOk() (*string, bool)`

GetCompanyPositionOk returns a tuple with the CompanyPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPosition

`func (o *UpdateProfileRequest) SetCompanyPosition(v string)`

SetCompanyPosition sets CompanyPosition field to given value.

### HasCompanyPosition

`func (o *UpdateProfileRequest) HasCompanyPosition() bool`

HasCompanyPosition returns a boolean if a field has been set.

### SetCompanyPositionNil

`func (o *UpdateProfileRequest) SetCompanyPositionNil(b bool)`

 SetCompanyPositionNil sets the value for CompanyPosition to be an explicit nil

### UnsetCompanyPosition
`func (o *UpdateProfileRequest) UnsetCompanyPosition()`

UnsetCompanyPosition ensures that no value is present for CompanyPosition, not even an explicit nil
### GetPrivateSettings

`func (o *UpdateProfileRequest) GetPrivateSettings() string`

GetPrivateSettings returns the PrivateSettings field if non-nil, zero value otherwise.

### GetPrivateSettingsOk

`func (o *UpdateProfileRequest) GetPrivateSettingsOk() (*string, bool)`

GetPrivateSettingsOk returns a tuple with the PrivateSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSettings

`func (o *UpdateProfileRequest) SetPrivateSettings(v string)`

SetPrivateSettings sets PrivateSettings field to given value.

### HasPrivateSettings

`func (o *UpdateProfileRequest) HasPrivateSettings() bool`

HasPrivateSettings returns a boolean if a field has been set.

### SetPrivateSettingsNil

`func (o *UpdateProfileRequest) SetPrivateSettingsNil(b bool)`

 SetPrivateSettingsNil sets the value for PrivateSettings to be an explicit nil

### UnsetPrivateSettings
`func (o *UpdateProfileRequest) UnsetPrivateSettings()`

UnsetPrivateSettings ensures that no value is present for PrivateSettings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


