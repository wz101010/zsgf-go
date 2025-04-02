# OAuthAccountBindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnionID** | **string** |  | 
**Platform** | **string** |  | 
**PlatformName** | **string** |  | 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuthAccountBindRequest

`func NewOAuthAccountBindRequest(unionID string, platform string, platformName string, ) *OAuthAccountBindRequest`

NewOAuthAccountBindRequest instantiates a new OAuthAccountBindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthAccountBindRequestWithDefaults

`func NewOAuthAccountBindRequestWithDefaults() *OAuthAccountBindRequest`

NewOAuthAccountBindRequestWithDefaults instantiates a new OAuthAccountBindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnionID

`func (o *OAuthAccountBindRequest) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *OAuthAccountBindRequest) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *OAuthAccountBindRequest) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.


### GetPlatform

`func (o *OAuthAccountBindRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OAuthAccountBindRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OAuthAccountBindRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformName

`func (o *OAuthAccountBindRequest) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *OAuthAccountBindRequest) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *OAuthAccountBindRequest) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetAvatar

`func (o *OAuthAccountBindRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *OAuthAccountBindRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *OAuthAccountBindRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *OAuthAccountBindRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *OAuthAccountBindRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *OAuthAccountBindRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *OAuthAccountBindRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OAuthAccountBindRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OAuthAccountBindRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *OAuthAccountBindRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *OAuthAccountBindRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OAuthAccountBindRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


