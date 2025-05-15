# SystemDirectoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**FileSize** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewSystemDirectoryItem

`func NewSystemDirectoryItem() *SystemDirectoryItem`

NewSystemDirectoryItem instantiates a new SystemDirectoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemDirectoryItemWithDefaults

`func NewSystemDirectoryItemWithDefaults() *SystemDirectoryItem`

NewSystemDirectoryItemWithDefaults instantiates a new SystemDirectoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemDirectoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemDirectoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemDirectoryItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemDirectoryItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SystemDirectoryItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SystemDirectoryItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreateDate

`func (o *SystemDirectoryItem) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *SystemDirectoryItem) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *SystemDirectoryItem) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *SystemDirectoryItem) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *SystemDirectoryItem) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *SystemDirectoryItem) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *SystemDirectoryItem) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *SystemDirectoryItem) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetFileSize

`func (o *SystemDirectoryItem) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *SystemDirectoryItem) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *SystemDirectoryItem) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *SystemDirectoryItem) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *SystemDirectoryItem) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *SystemDirectoryItem) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


