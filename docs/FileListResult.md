# FileListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directories** | Pointer to [**[]DirectoryItem**](DirectoryItem.md) |  | [optional] 
**Files** | Pointer to [**[]FileItem**](FileItem.md) |  | [optional] 
**TotalDirectories** | Pointer to **int32** |  | [optional] 
**TotalFiles** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewFileListResult

`func NewFileListResult() *FileListResult`

NewFileListResult instantiates a new FileListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileListResultWithDefaults

`func NewFileListResultWithDefaults() *FileListResult`

NewFileListResultWithDefaults instantiates a new FileListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectories

`func (o *FileListResult) GetDirectories() []DirectoryItem`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *FileListResult) GetDirectoriesOk() (*[]DirectoryItem, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *FileListResult) SetDirectories(v []DirectoryItem)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *FileListResult) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### SetDirectoriesNil

`func (o *FileListResult) SetDirectoriesNil(b bool)`

 SetDirectoriesNil sets the value for Directories to be an explicit nil

### UnsetDirectories
`func (o *FileListResult) UnsetDirectories()`

UnsetDirectories ensures that no value is present for Directories, not even an explicit nil
### GetFiles

`func (o *FileListResult) GetFiles() []FileItem`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FileListResult) GetFilesOk() (*[]FileItem, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FileListResult) SetFiles(v []FileItem)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FileListResult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FileListResult) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FileListResult) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetTotalDirectories

`func (o *FileListResult) GetTotalDirectories() int32`

GetTotalDirectories returns the TotalDirectories field if non-nil, zero value otherwise.

### GetTotalDirectoriesOk

`func (o *FileListResult) GetTotalDirectoriesOk() (*int32, bool)`

GetTotalDirectoriesOk returns a tuple with the TotalDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDirectories

`func (o *FileListResult) SetTotalDirectories(v int32)`

SetTotalDirectories sets TotalDirectories field to given value.

### HasTotalDirectories

`func (o *FileListResult) HasTotalDirectories() bool`

HasTotalDirectories returns a boolean if a field has been set.

### GetTotalFiles

`func (o *FileListResult) GetTotalFiles() int32`

GetTotalFiles returns the TotalFiles field if non-nil, zero value otherwise.

### GetTotalFilesOk

`func (o *FileListResult) GetTotalFilesOk() (*int32, bool)`

GetTotalFilesOk returns a tuple with the TotalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFiles

`func (o *FileListResult) SetTotalFiles(v int32)`

SetTotalFiles sets TotalFiles field to given value.

### HasTotalFiles

`func (o *FileListResult) HasTotalFiles() bool`

HasTotalFiles returns a boolean if a field has been set.

### GetSize

`func (o *FileListResult) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileListResult) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileListResult) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileListResult) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


