# GetUserProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Currencies** | Pointer to [**[]UserCurrency**](UserCurrency.md) |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 

## Methods

### NewGetUserProfileResult

`func NewGetUserProfileResult() *GetUserProfileResult`

NewGetUserProfileResult instantiates a new GetUserProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserProfileResultWithDefaults

`func NewGetUserProfileResultWithDefaults() *GetUserProfileResult`

NewGetUserProfileResultWithDefaults instantiates a new GetUserProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GetUserProfileResult) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetUserProfileResult) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetUserProfileResult) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *GetUserProfileResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCurrencies

`func (o *GetUserProfileResult) GetCurrencies() []UserCurrency`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *GetUserProfileResult) GetCurrenciesOk() (*[]UserCurrency, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *GetUserProfileResult) SetCurrencies(v []UserCurrency)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *GetUserProfileResult) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### SetCurrenciesNil

`func (o *GetUserProfileResult) SetCurrenciesNil(b bool)`

 SetCurrenciesNil sets the value for Currencies to be an explicit nil

### UnsetCurrencies
`func (o *GetUserProfileResult) UnsetCurrencies()`

UnsetCurrencies ensures that no value is present for Currencies, not even an explicit nil
### GetRole

`func (o *GetUserProfileResult) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetUserProfileResult) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetUserProfileResult) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetUserProfileResult) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *GetUserProfileResult) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GetUserProfileResult) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetLocation

`func (o *GetUserProfileResult) GetLocation() GeoLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetUserProfileResult) GetLocationOk() (*GeoLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetUserProfileResult) SetLocation(v GeoLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetUserProfileResult) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


