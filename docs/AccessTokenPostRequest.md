# AccessTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Tags** | Pointer to **NullableString** |  | [optional] 
**UserID** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**ExpireInDays** | Pointer to **int32** |  | [optional] 

## Methods

### NewAccessTokenPostRequest

`func NewAccessTokenPostRequest(title string, ) *AccessTokenPostRequest`

NewAccessTokenPostRequest instantiates a new AccessTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenPostRequestWithDefaults

`func NewAccessTokenPostRequestWithDefaults() *AccessTokenPostRequest`

NewAccessTokenPostRequestWithDefaults instantiates a new AccessTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AccessTokenPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccessTokenPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccessTokenPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTags

`func (o *AccessTokenPostRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccessTokenPostRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccessTokenPostRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccessTokenPostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccessTokenPostRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccessTokenPostRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUserID

`func (o *AccessTokenPostRequest) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AccessTokenPostRequest) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AccessTokenPostRequest) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AccessTokenPostRequest) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetDescription

`func (o *AccessTokenPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessTokenPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessTokenPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessTokenPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccessTokenPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccessTokenPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPermissions

`func (o *AccessTokenPostRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccessTokenPostRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccessTokenPostRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccessTokenPostRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AccessTokenPostRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AccessTokenPostRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetExpireInDays

`func (o *AccessTokenPostRequest) GetExpireInDays() int32`

GetExpireInDays returns the ExpireInDays field if non-nil, zero value otherwise.

### GetExpireInDaysOk

`func (o *AccessTokenPostRequest) GetExpireInDaysOk() (*int32, bool)`

GetExpireInDaysOk returns a tuple with the ExpireInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInDays

`func (o *AccessTokenPostRequest) SetExpireInDays(v int32)`

SetExpireInDays sets ExpireInDays field to given value.

### HasExpireInDays

`func (o *AccessTokenPostRequest) HasExpireInDays() bool`

HasExpireInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


