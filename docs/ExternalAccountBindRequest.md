# ExternalAccountBindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnionID** | **string** |  | 
**Platform** | **string** |  | 
**PlatformName** | **string** |  | 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExternalAccountBindRequest

`func NewExternalAccountBindRequest(unionID string, platform string, platformName string, ) *ExternalAccountBindRequest`

NewExternalAccountBindRequest instantiates a new ExternalAccountBindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountBindRequestWithDefaults

`func NewExternalAccountBindRequestWithDefaults() *ExternalAccountBindRequest`

NewExternalAccountBindRequestWithDefaults instantiates a new ExternalAccountBindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnionID

`func (o *ExternalAccountBindRequest) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *ExternalAccountBindRequest) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *ExternalAccountBindRequest) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.


### GetPlatform

`func (o *ExternalAccountBindRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ExternalAccountBindRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ExternalAccountBindRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformName

`func (o *ExternalAccountBindRequest) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *ExternalAccountBindRequest) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *ExternalAccountBindRequest) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetAvatar

`func (o *ExternalAccountBindRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ExternalAccountBindRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ExternalAccountBindRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ExternalAccountBindRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *ExternalAccountBindRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ExternalAccountBindRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *ExternalAccountBindRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalAccountBindRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalAccountBindRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ExternalAccountBindRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ExternalAccountBindRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ExternalAccountBindRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


