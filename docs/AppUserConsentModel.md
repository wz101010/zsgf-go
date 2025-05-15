# AppUserConsentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**GrantType** | Pointer to **NullableString** |  | [optional] 
**RedirectUri** | Pointer to **NullableString** |  | [optional] 
**Remark** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppUserConsentModel

`func NewAppUserConsentModel() *AppUserConsentModel`

NewAppUserConsentModel instantiates a new AppUserConsentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserConsentModelWithDefaults

`func NewAppUserConsentModelWithDefaults() *AppUserConsentModel`

NewAppUserConsentModelWithDefaults instantiates a new AppUserConsentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppUserConsentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUserConsentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUserConsentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppUserConsentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateDate

`func (o *AppUserConsentModel) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *AppUserConsentModel) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *AppUserConsentModel) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *AppUserConsentModel) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *AppUserConsentModel) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *AppUserConsentModel) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *AppUserConsentModel) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *AppUserConsentModel) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetGrantType

`func (o *AppUserConsentModel) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AppUserConsentModel) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AppUserConsentModel) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *AppUserConsentModel) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### SetGrantTypeNil

`func (o *AppUserConsentModel) SetGrantTypeNil(b bool)`

 SetGrantTypeNil sets the value for GrantType to be an explicit nil

### UnsetGrantType
`func (o *AppUserConsentModel) UnsetGrantType()`

UnsetGrantType ensures that no value is present for GrantType, not even an explicit nil
### GetRedirectUri

`func (o *AppUserConsentModel) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AppUserConsentModel) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AppUserConsentModel) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *AppUserConsentModel) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *AppUserConsentModel) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *AppUserConsentModel) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
### GetRemark

`func (o *AppUserConsentModel) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *AppUserConsentModel) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *AppUserConsentModel) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *AppUserConsentModel) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *AppUserConsentModel) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *AppUserConsentModel) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetScopes

`func (o *AppUserConsentModel) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AppUserConsentModel) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AppUserConsentModel) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AppUserConsentModel) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *AppUserConsentModel) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *AppUserConsentModel) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


