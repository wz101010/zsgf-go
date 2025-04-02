# \UserSettingAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSetting**](UserSettingAPI.md#UserSetting) | **Get** /UserSetting/{appKey}/{id} | 获取用户配置项详情
[**UserSettingDelete**](UserSettingAPI.md#UserSettingDelete) | **Delete** /UserSetting/{appKey}/{id} | 删除用户配置项
[**UserSettingPost**](UserSettingAPI.md#UserSettingPost) | **Post** /UserSetting/{appKey} | 添加用户配置项
[**UserSettingPut**](UserSettingAPI.md#UserSettingPut) | **Put** /UserSetting/{appKey}/{id} | 更新用户配置项
[**UserSettings**](UserSettingAPI.md#UserSettings) | **Get** /UserSetting/{appKey} | 获取用户配置列表



## UserSetting

> UserSettingApiResponse UserSetting(ctx, id, appKey).Execute()

获取用户配置项详情



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
	id := int64(789) // int64 | 配置项ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingAPI.UserSetting(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingAPI.UserSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSetting`: UserSettingApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingAPI.UserSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserSettingApiResponse**](UserSettingApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingDelete

> BooleanApiResponse UserSettingDelete(ctx, id, appKey).Execute()

删除用户配置项



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
	id := int64(789) // int64 | 配置项ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingAPI.UserSettingDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingAPI.UserSettingDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingAPI.UserSettingDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## UserSettingPost

> UserSettingPostResultApiResponse UserSettingPost(ctx, appKey).UserSetting(userSetting).Execute()

添加用户配置项



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
	userSetting := *openapiclient.NewUserSetting(int64(123), "Code_example", "Value_example") // UserSetting | 配置项内容 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingAPI.UserSettingPost(context.Background(), appKey).UserSetting(userSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingAPI.UserSettingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingPost`: UserSettingPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingAPI.UserSettingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSetting** | [**UserSetting**](UserSetting.md) | 配置项内容 | 

### Return type

[**UserSettingPostResultApiResponse**](UserSettingPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingPut

> BooleanApiResponse UserSettingPut(ctx, id, appKey).UserSetting(userSetting).Execute()

更新用户配置项



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
	id := int64(789) // int64 | 配置项ID
	appKey := "appKey_example" // string | 
	userSetting := *openapiclient.NewUserSetting(int64(123), "Code_example", "Value_example") // UserSetting | 配置项内容 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingAPI.UserSettingPut(context.Background(), id, appKey).UserSetting(userSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingAPI.UserSettingPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingAPI.UserSettingPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userSetting** | [**UserSetting**](UserSetting.md) | 配置项内容 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettings

> UserSettingListApiResponse UserSettings(ctx, appKey).UserId(userId).Code(code).Tag(tag).Execute()

获取用户配置列表



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
	userId := int64(789) // int64 | 用户ID (optional)
	code := "code_example" // string | 配置项代码 (optional)
	tag := "tag_example" // string | 配置项标签 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingAPI.UserSettings(context.Background(), appKey).UserId(userId).Code(code).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingAPI.UserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettings`: UserSettingListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserSettingAPI.UserSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int64** | 用户ID | 
 **code** | **string** | 配置项代码 | 
 **tag** | **string** | 配置项标签 | 

### Return type

[**UserSettingListApiResponse**](UserSettingListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

