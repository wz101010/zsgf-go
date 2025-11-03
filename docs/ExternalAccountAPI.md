# \ExternalAccountAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalAccountSignIn**](ExternalAccountAPI.md#ExternalAccountSignIn) | **Post** /ExternalAccount/{appKey}/SignIn | 外部账号登录
[**UserExternalAccountBind**](ExternalAccountAPI.md#UserExternalAccountBind) | **Post** /ExternalAccount/{appKey} | 绑定外部账号
[**UserOAuthAccounts**](ExternalAccountAPI.md#UserOAuthAccounts) | **Get** /ExternalAccount/{appKey} | 外部账号列表
[**UserOAuthAccountsPutBind**](ExternalAccountAPI.md#UserOAuthAccountsPutBind) | **Put** /ExternalAccount/{appKey}/{id} | 更新绑定账号
[**UserOAuthAccountsUnBind**](ExternalAccountAPI.md#UserOAuthAccountsUnBind) | **Delete** /ExternalAccount/{appKey}/{id} | 删除绑定账号



## ExternalAccountSignIn

> TokenModelApiResponse ExternalAccountSignIn(ctx, appKey).UserId(userId).ExternalAccountSignInRequest(externalAccountSignInRequest).Execute()

外部账号登录



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
	userId := "userId_example" // string |  (optional)
	externalAccountSignInRequest := *openapiclient.NewExternalAccountSignInRequest("UnionID_example", "Platform_example") // ExternalAccountSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAccountAPI.ExternalAccountSignIn(context.Background(), appKey).UserId(userId).ExternalAccountSignInRequest(externalAccountSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountAPI.ExternalAccountSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalAccountSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalAccountAPI.ExternalAccountSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalAccountSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 
 **externalAccountSignInRequest** | [**ExternalAccountSignInRequest**](ExternalAccountSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserExternalAccountBind

> BooleanApiResponse UserExternalAccountBind(ctx, appKey).UserId(userId).ExternalAccountBindRequest(externalAccountBindRequest).Execute()

绑定外部账号



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
	userId := "userId_example" // string |  (optional)
	externalAccountBindRequest := *openapiclient.NewExternalAccountBindRequest("UnionID_example", "Platform_example", "PlatformName_example") // ExternalAccountBindRequest | 绑定请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAccountAPI.UserExternalAccountBind(context.Background(), appKey).UserId(userId).ExternalAccountBindRequest(externalAccountBindRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountAPI.UserExternalAccountBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserExternalAccountBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalAccountAPI.UserExternalAccountBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserExternalAccountBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 
 **externalAccountBindRequest** | [**ExternalAccountBindRequest**](ExternalAccountBindRequest.md) | 绑定请求参数 | 

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


## UserOAuthAccounts

> UserLoginsListApiResponse UserOAuthAccounts(ctx, appKey).UserId(userId).Execute()

外部账号列表



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
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAccountAPI.UserOAuthAccounts(context.Background(), appKey).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountAPI.UserOAuthAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccounts`: UserLoginsListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalAccountAPI.UserOAuthAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** |  | 

### Return type

[**UserLoginsListApiResponse**](UserLoginsListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOAuthAccountsPutBind

> BooleanApiResponse UserOAuthAccountsPutBind(ctx, id, appKey).UserId(userId).ExternalAccountPutRequest(externalAccountPutRequest).Execute()

更新绑定账号



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
	id := int64(789) // int64 | 绑定ID
	appKey := "appKey_example" // string | 
	userId := "userId_example" // string |  (optional)
	externalAccountPutRequest := *openapiclient.NewExternalAccountPutRequest() // ExternalAccountPutRequest | 更新请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAccountAPI.UserOAuthAccountsPutBind(context.Background(), id, appKey).UserId(userId).ExternalAccountPutRequest(externalAccountPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountAPI.UserOAuthAccountsPutBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountsPutBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalAccountAPI.UserOAuthAccountsPutBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 绑定ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsPutBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **string** |  | 
 **externalAccountPutRequest** | [**ExternalAccountPutRequest**](ExternalAccountPutRequest.md) | 更新请求参数 | 

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


## UserOAuthAccountsUnBind

> BooleanApiResponse UserOAuthAccountsUnBind(ctx, id, appKey).UserId(userId).Execute()

删除绑定账号



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
	id := int64(789) // int64 | 绑定ID
	appKey := "appKey_example" // string | 
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExternalAccountAPI.UserOAuthAccountsUnBind(context.Background(), id, appKey).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExternalAccountAPI.UserOAuthAccountsUnBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountsUnBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ExternalAccountAPI.UserOAuthAccountsUnBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 绑定ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsUnBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **string** |  | 

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

