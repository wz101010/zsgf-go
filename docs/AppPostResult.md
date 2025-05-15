# AppPostResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ServerPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppPostResult

`func NewAppPostResult() *AppPostResult`

NewAppPostResult instantiates a new AppPostResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPostResultWithDefaults

`func NewAppPostResultWithDefaults() *AppPostResult`

NewAppPostResultWithDefaults instantiates a new AppPostResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppPostResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPostResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPostResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AppPostResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerPath

`func (o *AppPostResult) GetServerPath() string`

GetServerPath returns the ServerPath field if non-nil, zero value otherwise.

### GetServerPathOk

`func (o *AppPostResult) GetServerPathOk() (*string, bool)`

GetServerPathOk returns a tuple with the ServerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPath

`func (o *AppPostResult) SetServerPath(v string)`

SetServerPath sets ServerPath field to given value.

### HasServerPath

`func (o *AppPostResult) HasServerPath() bool`

HasServerPath returns a boolean if a field has been set.

### SetServerPathNil

`func (o *AppPostResult) SetServerPathNil(b bool)`

 SetServerPathNil sets the value for ServerPath to be an explicit nil

### UnsetServerPath
`func (o *AppPostResult) UnsetServerPath()`

UnsetServerPath ensures that no value is present for ServerPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


