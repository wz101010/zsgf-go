# \FileAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileCreateFolder**](FileAPI.md#FileCreateFolder) | **Post** /File/{appKey}/CreateFolder | 创建文件夹
[**FileDelete**](FileAPI.md#FileDelete) | **Delete** /File/{appKey} | 删除文件或文件夹
[**FileRename**](FileAPI.md#FileRename) | **Post** /File/{appKey}/Rename | 重命名文件或文件夹
[**FileUpload**](FileAPI.md#FileUpload) | **Post** /File/{appKey}/Upload | 上传文件
[**Files**](FileAPI.md#Files) | **Get** /File/{appKey} | 获取文件列表



## FileCreateFolder

> BooleanApiResponse FileCreateFolder(ctx, appKey).Path(path).Execute()

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
	appKey := "appKey_example" // string | 
	path := "path_example" // string | 文件夹路径 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.FileCreateFolder(context.Background(), appKey).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FileCreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileCreateFolder`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.FileCreateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileCreateFolderRequest struct via the builder pattern


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


## FileDelete

> BooleanApiResponse FileDelete(ctx, appKey).Path(path).Execute()

删除文件或文件夹



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
	appKey := "appKey_example" // string | 
	path := "path_example" // string | 文件或文件夹路径 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.FileDelete(context.Background(), appKey).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.FileDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | 文件或文件夹路径 | 

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


## FileRename

> BooleanApiResponse FileRename(ctx, appKey).SourceName(sourceName).DestName(destName).Execute()

重命名文件或文件夹



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
	appKey := "appKey_example" // string | 
	sourceName := "sourceName_example" // string | 原文件或文件夹名称 (optional)
	destName := "destName_example" // string | 新文件或文件夹名称 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.FileRename(context.Background(), appKey).SourceName(sourceName).DestName(destName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FileRename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileRename`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.FileRename`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceName** | **string** | 原文件或文件夹名称 | 
 **destName** | **string** | 新文件或文件夹名称 | 

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


## FileUpload

> BooleanApiResponse FileUpload(ctx, appKey).Path(path).File(file).Execute()

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
	appKey := "appKey_example" // string | 
	path := "path_example" // string | 文件夹路径 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 上传的文件 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.FileUpload(context.Background(), appKey).Path(path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FileUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileUpload`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.FileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | 文件夹路径 | 
 **file** | ***os.File** | 上传的文件 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Files

> FileListResultApiResponse Files(ctx, appKey).Path(path).Skip(skip).Take(take).Execute()

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
	appKey := "appKey_example" // string | 
	path := "path_example" // string | 文件路径 (optional)
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileAPI.Files(context.Background(), appKey).Path(path).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.Files``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Files`: FileListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `FileAPI.Files`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | 文件路径 | 
 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 100]

### Return type

[**FileListResultApiResponse**](FileListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

