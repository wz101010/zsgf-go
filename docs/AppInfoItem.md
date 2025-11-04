# AppInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppKey** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Logo** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectID** | Pointer to **int64** |  | [optional] 

## Methods

### NewAppInfoItem

`func NewAppInfoItem() *AppInfoItem`

NewAppInfoItem instantiates a new AppInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInfoItemWithDefaults

`func NewAppInfoItemWithDefaults() *AppInfoItem`

NewAppInfoItemWithDefaults instantiates a new AppInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppKey

`func (o *AppInfoItem) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *AppInfoItem) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *AppInfoItem) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *AppInfoItem) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### SetAppKeyNil

`func (o *AppInfoItem) SetAppKeyNil(b bool)`

 SetAppKeyNil sets the value for AppKey to be an explicit nil

### UnsetAppKey
`func (o *AppInfoItem) UnsetAppKey()`

UnsetAppKey ensures that no value is present for AppKey, not even an explicit nil
### GetId

`func (o *AppInfoItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppInfoItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppInfoItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppInfoItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppInfoItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppInfoItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppInfoItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppInfoItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AppInfoItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AppInfoItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogo

`func (o *AppInfoItem) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AppInfoItem) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AppInfoItem) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AppInfoItem) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *AppInfoItem) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *AppInfoItem) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetTags

`func (o *AppInfoItem) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AppInfoItem) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AppInfoItem) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AppInfoItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AppInfoItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AppInfoItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetWebsite

`func (o *AppInfoItem) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AppInfoItem) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AppInfoItem) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AppInfoItem) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *AppInfoItem) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *AppInfoItem) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetDescription

`func (o *AppInfoItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppInfoItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppInfoItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppInfoItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AppInfoItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppInfoItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectID

`func (o *AppInfoItem) GetProjectID() int64`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *AppInfoItem) GetProjectIDOk() (*int64, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *AppInfoItem) SetProjectID(v int64)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *AppInfoItem) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


