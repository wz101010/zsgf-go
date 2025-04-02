# FollowerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TargetUserID** | Pointer to **int64** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**IsMutual** | Pointer to **bool** |  | [optional] 
**ClosenessScore** | Pointer to **int64** |  | [optional] 
**AttentionScore** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Remark** | Pointer to **NullableString** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFollowerModel

`func NewFollowerModel() *FollowerModel`

NewFollowerModel instantiates a new FollowerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowerModelWithDefaults

`func NewFollowerModelWithDefaults() *FollowerModel`

NewFollowerModelWithDefaults instantiates a new FollowerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FollowerModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FollowerModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FollowerModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FollowerModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetUserID

`func (o *FollowerModel) GetTargetUserID() int64`

GetTargetUserID returns the TargetUserID field if non-nil, zero value otherwise.

### GetTargetUserIDOk

`func (o *FollowerModel) GetTargetUserIDOk() (*int64, bool)`

GetTargetUserIDOk returns a tuple with the TargetUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserID

`func (o *FollowerModel) SetTargetUserID(v int64)`

SetTargetUserID sets TargetUserID field to given value.

### HasTargetUserID

`func (o *FollowerModel) HasTargetUserID() bool`

HasTargetUserID returns a boolean if a field has been set.

### GetAlias

`func (o *FollowerModel) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *FollowerModel) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *FollowerModel) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *FollowerModel) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *FollowerModel) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *FollowerModel) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetNickName

`func (o *FollowerModel) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *FollowerModel) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *FollowerModel) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *FollowerModel) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *FollowerModel) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *FollowerModel) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *FollowerModel) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *FollowerModel) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *FollowerModel) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *FollowerModel) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *FollowerModel) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *FollowerModel) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetIsMutual

`func (o *FollowerModel) GetIsMutual() bool`

GetIsMutual returns the IsMutual field if non-nil, zero value otherwise.

### GetIsMutualOk

`func (o *FollowerModel) GetIsMutualOk() (*bool, bool)`

GetIsMutualOk returns a tuple with the IsMutual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutual

`func (o *FollowerModel) SetIsMutual(v bool)`

SetIsMutual sets IsMutual field to given value.

### HasIsMutual

`func (o *FollowerModel) HasIsMutual() bool`

HasIsMutual returns a boolean if a field has been set.

### GetClosenessScore

`func (o *FollowerModel) GetClosenessScore() int64`

GetClosenessScore returns the ClosenessScore field if non-nil, zero value otherwise.

### GetClosenessScoreOk

`func (o *FollowerModel) GetClosenessScoreOk() (*int64, bool)`

GetClosenessScoreOk returns a tuple with the ClosenessScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosenessScore

`func (o *FollowerModel) SetClosenessScore(v int64)`

SetClosenessScore sets ClosenessScore field to given value.

### HasClosenessScore

`func (o *FollowerModel) HasClosenessScore() bool`

HasClosenessScore returns a boolean if a field has been set.

### GetAttentionScore

`func (o *FollowerModel) GetAttentionScore() int64`

GetAttentionScore returns the AttentionScore field if non-nil, zero value otherwise.

### GetAttentionScoreOk

`func (o *FollowerModel) GetAttentionScoreOk() (*int64, bool)`

GetAttentionScoreOk returns a tuple with the AttentionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttentionScore

`func (o *FollowerModel) SetAttentionScore(v int64)`

SetAttentionScore sets AttentionScore field to given value.

### HasAttentionScore

`func (o *FollowerModel) HasAttentionScore() bool`

HasAttentionScore returns a boolean if a field has been set.

### GetTags

`func (o *FollowerModel) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FollowerModel) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FollowerModel) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FollowerModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FollowerModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FollowerModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetStatus

`func (o *FollowerModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FollowerModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FollowerModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FollowerModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FollowerModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FollowerModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRemark

`func (o *FollowerModel) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *FollowerModel) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *FollowerModel) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *FollowerModel) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *FollowerModel) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *FollowerModel) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetCreateDate

`func (o *FollowerModel) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *FollowerModel) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *FollowerModel) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *FollowerModel) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


