# UserProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppKey** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**UnionID** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**PhoneIsValid** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**EmailIsValid** | Pointer to **bool** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Currencies** | Pointer to [**[]UserCurrency**](UserCurrency.md) |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserProfileResult

`func NewUserProfileResult() *UserProfileResult`

NewUserProfileResult instantiates a new UserProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileResultWithDefaults

`func NewUserProfileResultWithDefaults() *UserProfileResult`

NewUserProfileResultWithDefaults instantiates a new UserProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppKey

`func (o *UserProfileResult) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *UserProfileResult) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *UserProfileResult) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *UserProfileResult) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### SetAppKeyNil

`func (o *UserProfileResult) SetAppKeyNil(b bool)`

 SetAppKeyNil sets the value for AppKey to be an explicit nil

### UnsetAppKey
`func (o *UserProfileResult) UnsetAppKey()`

UnsetAppKey ensures that no value is present for AppKey, not even an explicit nil
### GetPlatform

`func (o *UserProfileResult) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UserProfileResult) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UserProfileResult) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UserProfileResult) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *UserProfileResult) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *UserProfileResult) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetUnionID

`func (o *UserProfileResult) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *UserProfileResult) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *UserProfileResult) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.

### HasUnionID

`func (o *UserProfileResult) HasUnionID() bool`

HasUnionID returns a boolean if a field has been set.

### SetUnionIDNil

`func (o *UserProfileResult) SetUnionIDNil(b bool)`

 SetUnionIDNil sets the value for UnionID to be an explicit nil

### UnsetUnionID
`func (o *UserProfileResult) UnsetUnionID()`

UnsetUnionID ensures that no value is present for UnionID, not even an explicit nil
### GetPhone

`func (o *UserProfileResult) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserProfileResult) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserProfileResult) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserProfileResult) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UserProfileResult) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UserProfileResult) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCreateDate

`func (o *UserProfileResult) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *UserProfileResult) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *UserProfileResult) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *UserProfileResult) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUserName

`func (o *UserProfileResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserProfileResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserProfileResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserProfileResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserProfileResult) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserProfileResult) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPhoneIsValid

`func (o *UserProfileResult) GetPhoneIsValid() bool`

GetPhoneIsValid returns the PhoneIsValid field if non-nil, zero value otherwise.

### GetPhoneIsValidOk

`func (o *UserProfileResult) GetPhoneIsValidOk() (*bool, bool)`

GetPhoneIsValidOk returns a tuple with the PhoneIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneIsValid

`func (o *UserProfileResult) SetPhoneIsValid(v bool)`

SetPhoneIsValid sets PhoneIsValid field to given value.

### HasPhoneIsValid

`func (o *UserProfileResult) HasPhoneIsValid() bool`

HasPhoneIsValid returns a boolean if a field has been set.

### GetData

`func (o *UserProfileResult) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserProfileResult) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserProfileResult) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UserProfileResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserProfileResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserProfileResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *UserProfileResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfileResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfileResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProfileResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserProfileResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserProfileResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailIsValid

`func (o *UserProfileResult) GetEmailIsValid() bool`

GetEmailIsValid returns the EmailIsValid field if non-nil, zero value otherwise.

### GetEmailIsValidOk

`func (o *UserProfileResult) GetEmailIsValidOk() (*bool, bool)`

GetEmailIsValidOk returns a tuple with the EmailIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIsValid

`func (o *UserProfileResult) SetEmailIsValid(v bool)`

SetEmailIsValid sets EmailIsValid field to given value.

### HasEmailIsValid

`func (o *UserProfileResult) HasEmailIsValid() bool`

HasEmailIsValid returns a boolean if a field has been set.

### GetLastUpdate

`func (o *UserProfileResult) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserProfileResult) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *UserProfileResult) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *UserProfileResult) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetNickName

`func (o *UserProfileResult) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *UserProfileResult) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *UserProfileResult) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *UserProfileResult) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *UserProfileResult) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *UserProfileResult) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetId

`func (o *UserProfileResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfileResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfileResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserProfileResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAvatar

`func (o *UserProfileResult) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserProfileResult) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserProfileResult) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserProfileResult) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UserProfileResult) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UserProfileResult) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCurrencies

`func (o *UserProfileResult) GetCurrencies() []UserCurrency`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *UserProfileResult) GetCurrenciesOk() (*[]UserCurrency, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *UserProfileResult) SetCurrencies(v []UserCurrency)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *UserProfileResult) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### SetCurrenciesNil

`func (o *UserProfileResult) SetCurrenciesNil(b bool)`

 SetCurrenciesNil sets the value for Currencies to be an explicit nil

### UnsetCurrencies
`func (o *UserProfileResult) UnsetCurrencies()`

UnsetCurrencies ensures that no value is present for Currencies, not even an explicit nil
### GetRole

`func (o *UserProfileResult) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserProfileResult) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserProfileResult) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserProfileResult) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *UserProfileResult) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *UserProfileResult) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


