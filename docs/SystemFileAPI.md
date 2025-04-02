# \SystemFileAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemFileCreateFolder**](SystemFileAPI.md#SystemFileCreateFolder) | **Post** /SystemFile/CreateFolder | 创建文件夹
[**SystemFileDelete**](SystemFileAPI.md#SystemFileDelete) | **Delete** /SystemFile | 删除文件
[**SystemFileRename**](SystemFileAPI.md#SystemFileRename) | **Post** /SystemFile/Rename | 重命名文件
[**SystemFileUpload**](SystemFileAPI.md#SystemFileUpload) | **Post** /SystemFile | 上传文件
[**SystemFiles**](SystemFileAPI.md#SystemFiles) | **Get** /SystemFile | 获取文件列表



## SystemFileCreateFolder

> BooleanApiResponse SystemFileCreateFolder(ctx).Path(path).Execute()

创建文件夹



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	path := "path_example" // string | 文件夹路径 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFileAPI.SystemFileCreateFolder(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFileAPI.SystemFileCreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemFileCreateFolder`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFileAPI.SystemFileCreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemFileCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 文件夹路径 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemFileDelete

> BooleanApiResponse SystemFileDelete(ctx).Path(path).Execute()

删除文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	path := "path_example" // string | 文件路径 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFileAPI.SystemFileDelete(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFileAPI.SystemFileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemFileDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFileAPI.SystemFileDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemFileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 文件路径 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemFileRename

> BooleanApiResponse SystemFileRename(ctx).SourceName(sourceName).DestName(destName).Execute()

重命名文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	sourceName := "sourceName_example" // string | 源文件名 (optional)
	destName := "destName_example" // string | 目标文件名 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFileAPI.SystemFileRename(context.Background()).SourceName(sourceName).DestName(destName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFileAPI.SystemFileRename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemFileRename`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFileAPI.SystemFileRename`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemFileRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceName** | **string** | 源文件名 | 
 **destName** | **string** | 目标文件名 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemFileUpload

> StringApiResponse SystemFileUpload(ctx).Path(path).File(file).Execute()

上传文件



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	path := "path_example" // string | 上传的路径 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 上传的文件 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFileAPI.SystemFileUpload(context.Background()).Path(path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFileAPI.SystemFileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemFileUpload`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFileAPI.SystemFileUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 上传的路径 | 
 **file** | ***os.File** | 上传的文件 | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemFiles

> SystemFileListResultApiResponse SystemFiles(ctx).Path(path).Execute()

获取文件列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	path := "path_example" // string | 文件路径 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemFileAPI.SystemFiles(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemFileAPI.SystemFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemFiles`: SystemFileListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SystemFileAPI.SystemFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | 文件路径 | 

### Return type

[**SystemFileListResultApiResponse**](SystemFileListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

