# UserQRCodeScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppID** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **NullableString** |  | [optional] 
**Scopes** | Pointer to **NullableString** |  | [optional] 
**Remark** | Pointer to **NullableString** |  | [optional] 
**Scheme** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserQRCodeScanResult

`func NewUserQRCodeScanResult() *UserQRCodeScanResult`

NewUserQRCodeScanResult instantiates a new UserQRCodeScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserQRCodeScanResultWithDefaults

`func NewUserQRCodeScanResultWithDefaults() *UserQRCodeScanResult`

NewUserQRCodeScanResultWithDefaults instantiates a new UserQRCodeScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppID

`func (o *UserQRCodeScanResult) GetAppID() int64`

GetAppID returns the AppID field if non-nil, zero value otherwise.

### GetAppIDOk

`func (o *UserQRCodeScanResult) GetAppIDOk() (*int64, bool)`

GetAppIDOk returns a tuple with the AppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppID

`func (o *UserQRCodeScanResult) SetAppID(v int64)`

SetAppID sets AppID field to given value.

### HasAppID

`func (o *UserQRCodeScanResult) HasAppID() bool`

HasAppID returns a boolean if a field has been set.

### GetName

`func (o *UserQRCodeScanResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserQRCodeScanResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserQRCodeScanResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserQRCodeScanResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserQRCodeScanResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserQRCodeScanResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogo

`func (o *UserQRCodeScanResult) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UserQRCodeScanResult) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UserQRCodeScanResult) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UserQRCodeScanResult) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *UserQRCodeScanResult) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *UserQRCodeScanResult) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetWebsite

`func (o *UserQRCodeScanResult) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *UserQRCodeScanResult) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *UserQRCodeScanResult) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *UserQRCodeScanResult) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *UserQRCodeScanResult) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *UserQRCodeScanResult) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetDescription

`func (o *UserQRCodeScanResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserQRCodeScanResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserQRCodeScanResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserQRCodeScanResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserQRCodeScanResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserQRCodeScanResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *UserQRCodeScanResult) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserQRCodeScanResult) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserQRCodeScanResult) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserQRCodeScanResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UserQRCodeScanResult) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UserQRCodeScanResult) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetScopes

`func (o *UserQRCodeScanResult) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UserQRCodeScanResult) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UserQRCodeScanResult) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UserQRCodeScanResult) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopesNil

`func (o *UserQRCodeScanResult) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *UserQRCodeScanResult) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetRemark

`func (o *UserQRCodeScanResult) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *UserQRCodeScanResult) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *UserQRCodeScanResult) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *UserQRCodeScanResult) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### SetRemarkNil

`func (o *UserQRCodeScanResult) SetRemarkNil(b bool)`

 SetRemarkNil sets the value for Remark to be an explicit nil

### UnsetRemark
`func (o *UserQRCodeScanResult) UnsetRemark()`

UnsetRemark ensures that no value is present for Remark, not even an explicit nil
### GetScheme

`func (o *UserQRCodeScanResult) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *UserQRCodeScanResult) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *UserQRCodeScanResult) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *UserQRCodeScanResult) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *UserQRCodeScanResult) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *UserQRCodeScanResult) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


