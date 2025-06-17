# GrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUri** | Pointer to **NullableString** | 返回地址。默认无限制，也可在【安全-开放认证网址白名单】配置 | [optional] 
**GrantType** | **string** | 授权类型。可选：email、phone、unionid、account | 
**Scopes** | **string** | 自定义授权范围，用英文空格分隔 | 
**UserName** | Pointer to **NullableString** | 用户名。授权类型为：email/phone/account必填 | [optional] 
**Password** | Pointer to **NullableString** | 登录密码。授权类型为：email/phone/account必填 | [optional] 
**UnionId** | Pointer to **NullableString** | unionId。授权类型为：unionid必填 | [optional] 
**Platform** | Pointer to **NullableString** | 平台。授权类型为：unionid必填 | [optional] 
**ExpireInDays** | Pointer to **int32** | 授权有效期。1~99天 | [optional] 

## Methods

### NewGrantRequest

`func NewGrantRequest(grantType string, scopes string, ) *GrantRequest`

NewGrantRequest instantiates a new GrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRequestWithDefaults

`func NewGrantRequestWithDefaults() *GrantRequest`

NewGrantRequestWithDefaults instantiates a new GrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUri

`func (o *GrantRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GrantRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GrantRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GrantRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *GrantRequest) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *GrantRequest) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
### GetGrantType

`func (o *GrantRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetScopes

`func (o *GrantRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GrantRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GrantRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.


### GetUserName

`func (o *GrantRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GrantRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GrantRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GrantRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *GrantRequest) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *GrantRequest) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPassword

`func (o *GrantRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GrantRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GrantRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GrantRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GrantRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GrantRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetUnionId

`func (o *GrantRequest) GetUnionId() string`

GetUnionId returns the UnionId field if non-nil, zero value otherwise.

### GetUnionIdOk

`func (o *GrantRequest) GetUnionIdOk() (*string, bool)`

GetUnionIdOk returns a tuple with the UnionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionId

`func (o *GrantRequest) SetUnionId(v string)`

SetUnionId sets UnionId field to given value.

### HasUnionId

`func (o *GrantRequest) HasUnionId() bool`

HasUnionId returns a boolean if a field has been set.

### SetUnionIdNil

`func (o *GrantRequest) SetUnionIdNil(b bool)`

 SetUnionIdNil sets the value for UnionId to be an explicit nil

### UnsetUnionId
`func (o *GrantRequest) UnsetUnionId()`

UnsetUnionId ensures that no value is present for UnionId, not even an explicit nil
### GetPlatform

`func (o *GrantRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GrantRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GrantRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GrantRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GrantRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GrantRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetExpireInDays

`func (o *GrantRequest) GetExpireInDays() int32`

GetExpireInDays returns the ExpireInDays field if non-nil, zero value otherwise.

### GetExpireInDaysOk

`func (o *GrantRequest) GetExpireInDaysOk() (*int32, bool)`

GetExpireInDaysOk returns a tuple with the ExpireInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInDays

`func (o *GrantRequest) SetExpireInDays(v int32)`

SetExpireInDays sets ExpireInDays field to given value.

### HasExpireInDays

`func (o *GrantRequest) HasExpireInDays() bool`

HasExpireInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


