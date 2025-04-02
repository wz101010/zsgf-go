# AccessTokenPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | 访问令牌的标题 | [default to "默认标题"]
**Tags** | Pointer to **NullableString** | 访问令牌的标签 | [optional] [default to ""]
**Description** | Pointer to **NullableString** | 访问令牌的描述 | [optional] [default to ""]
**Permissions** | Pointer to **NullableString** | 访问令牌的权限 | [optional] [default to ""]
**Enable** | Pointer to **bool** | 是否启用访问令牌 | [optional] [default to false]

## Methods

### NewAccessTokenPutRequest

`func NewAccessTokenPutRequest(title string, ) *AccessTokenPutRequest`

NewAccessTokenPutRequest instantiates a new AccessTokenPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenPutRequestWithDefaults

`func NewAccessTokenPutRequestWithDefaults() *AccessTokenPutRequest`

NewAccessTokenPutRequestWithDefaults instantiates a new AccessTokenPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AccessTokenPutRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccessTokenPutRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccessTokenPutRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTags

`func (o *AccessTokenPutRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessTokenPutRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessTokenPutRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessTokenPutRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccessTokenPutRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccessTokenPutRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *AccessTokenPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessTokenPutRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessTokenPutRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPermissions

`func (o *AccessTokenPutRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccessTokenPutRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccessTokenPutRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccessTokenPutRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AccessTokenPutRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AccessTokenPutRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetEnable

`func (o *AccessTokenPutRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AccessTokenPutRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AccessTokenPutRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *AccessTokenPutRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


