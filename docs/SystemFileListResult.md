# SystemFileListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | Pointer to [**[]SystemDirectoryItem**](SystemDirectoryItem.md) |  | [optional] 
**Files** | Pointer to [**[]SystemFileItem**](SystemFileItem.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewSystemFileListResult

`func NewSystemFileListResult() *SystemFileListResult`

NewSystemFileListResult instantiates a new SystemFileListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemFileListResultWithDefaults

`func NewSystemFileListResultWithDefaults() *SystemFileListResult`

NewSystemFileListResultWithDefaults instantiates a new SystemFileListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *SystemFileListResult) GetDirectories() []SystemDirectoryItem`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *SystemFileListResult) GetDirectoriesOk() (*[]SystemDirectoryItem, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *SystemFileListResult) SetDirectories(v []SystemDirectoryItem)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *SystemFileListResult) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### SetDirectoriesNil

`func (o *SystemFileListResult) SetDirectoriesNil(b bool)`

 SetDirectoriesNil sets the value for Directories to be an explicit nil

### UnsetDirectories
`func (o *SystemFileListResult) UnsetDirectories()`

UnsetDirectories ensures that no value is present for Directories, not even an explicit nil
### GetFiles

`func (o *SystemFileListResult) GetFiles() []SystemFileItem`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SystemFileListResult) GetFilesOk() (*[]SystemFileItem, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SystemFileListResult) SetFiles(v []SystemFileItem)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *SystemFileListResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *SystemFileListResult) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *SystemFileListResult) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetSize

`func (o *SystemFileListResult) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SystemFileListResult) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SystemFileListResult) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SystemFileListResult) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


