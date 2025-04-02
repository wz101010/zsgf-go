# \AccessTokenAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessTokenDelete**](AccessTokenAPI.md#AccessTokenDelete) | **Delete** /AccessToken/{appKey}/{id} | 删除令牌
[**AccessTokenPost**](AccessTokenAPI.md#AccessTokenPost) | **Post** /AccessToken/{appKey} | 创建令牌
[**AccessTokenPut**](AccessTokenAPI.md#AccessTokenPut) | **Put** /AccessToken/{appKey}/{id} | 更新令牌
[**AccessTokens**](AccessTokenAPI.md#AccessTokens) | **Get** /AccessToken/{appKey} | 令牌列表



## AccessTokenDelete

> BooleanApiResponse AccessTokenDelete(ctx, id, appKey).Execute()

删除令牌



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
	id := int64(789) // int64 | 令牌ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokenAPI.AccessTokenDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.AccessTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokenDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.AccessTokenDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 令牌ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokenDeleteRequest struct via the builder pattern


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


## AccessTokenPost

> TokenModelApiResponse AccessTokenPost(ctx, appKey).AccessTokenPostRequest(accessTokenPostRequest).Execute()

创建令牌



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
	accessTokenPostRequest := *openapiclient.NewAccessTokenPostRequest("Title_example") // AccessTokenPostRequest | 令牌创建请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokenAPI.AccessTokenPost(context.Background(), appKey).AccessTokenPostRequest(accessTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.AccessTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokenPost`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.AccessTokenPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessTokenPostRequest** | [**AccessTokenPostRequest**](AccessTokenPostRequest.md) | 令牌创建请求参数 | 

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


## AccessTokenPut

> BooleanApiResponse AccessTokenPut(ctx, id, appKey).AccessTokenPutRequest(accessTokenPutRequest).Execute()

更新令牌



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
	id := int64(789) // int64 | 令牌ID
	appKey := "appKey_example" // string | 
	accessTokenPutRequest := *openapiclient.NewAccessTokenPutRequest("Title_example") // AccessTokenPutRequest | 令牌更新请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokenAPI.AccessTokenPut(context.Background(), id, appKey).AccessTokenPutRequest(accessTokenPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.AccessTokenPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokenPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.AccessTokenPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 令牌ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokenPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessTokenPutRequest** | [**AccessTokenPutRequest**](AccessTokenPutRequest.md) | 令牌更新请求参数 | 

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


## AccessTokens

> AccessTokenListResultApiResponse AccessTokens(ctx, appKey).UserId(userId).Skip(skip).Take(take).Execute()

令牌列表



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
	skip := int32(56) // int32 | 跳过的记录数 (optional)
	take := int32(56) // int32 | 获取的记录数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokenAPI.AccessTokens(context.Background(), appKey).UserId(userId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.AccessTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokens`: AccessTokenListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.AccessTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int64** | 用户ID | 
 **skip** | **int32** | 跳过的记录数 | 
 **take** | **int32** | 获取的记录数 | 

### Return type

[**AccessTokenListResultApiResponse**](AccessTokenListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

