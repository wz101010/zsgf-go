# UserProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
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

### GetUser

`func (o *UserProfileResult) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserProfileResult) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserProfileResult) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserProfileResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

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


