# \AuthorizePolicyAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizePolicies**](AuthorizePolicyAPI.md#AuthorizePolicies) | **Get** /AuthorizePolicy/{appKey} | 获取鉴权策略列表
[**AuthorizePolicy**](AuthorizePolicyAPI.md#AuthorizePolicy) | **Get** /AuthorizePolicy/{appKey}/{id} | 获取鉴权策略详情
[**AuthorizePolicyDelete**](AuthorizePolicyAPI.md#AuthorizePolicyDelete) | **Delete** /AuthorizePolicy/{appKey}/{id} | 删除鉴权策略
[**AuthorizePolicyPost**](AuthorizePolicyAPI.md#AuthorizePolicyPost) | **Post** /AuthorizePolicy/{appKey} | 添加鉴权策略
[**AuthorizePolicyPut**](AuthorizePolicyAPI.md#AuthorizePolicyPut) | **Put** /AuthorizePolicy/{appKey}/{id} | 更新鉴权策略



## AuthorizePolicies

> AuthorizePolicyListApiResponse AuthorizePolicies(ctx, appKey).AuthType(authType).PolicyName(policyName).Execute()

获取鉴权策略列表



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
	authType := "authType_example" // string | 鉴权类型（access_token、user、role） (optional)
	policyName := "policyName_example" // string | 策略名称 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizePolicyAPI.AuthorizePolicies(context.Background(), appKey).AuthType(authType).PolicyName(policyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizePolicyAPI.AuthorizePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePolicies`: AuthorizePolicyListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizePolicyAPI.AuthorizePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authType** | **string** | 鉴权类型（access_token、user、role） | 
 **policyName** | **string** | 策略名称 | 

### Return type

[**AuthorizePolicyListApiResponse**](AuthorizePolicyListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePolicy

> AuthorizePolicyApiResponse AuthorizePolicy(ctx, id, appKey).Execute()

获取鉴权策略详情



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
	id := int64(789) // int64 | 鉴权策略ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizePolicyAPI.AuthorizePolicy(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizePolicyAPI.AuthorizePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePolicy`: AuthorizePolicyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizePolicyAPI.AuthorizePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 鉴权策略ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizePolicyApiResponse**](AuthorizePolicyApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePolicyDelete

> BooleanApiResponse AuthorizePolicyDelete(ctx, id, appKey).Execute()

删除鉴权策略



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
	id := int64(789) // int64 | 鉴权策略ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizePolicyAPI.AuthorizePolicyDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizePolicyAPI.AuthorizePolicyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePolicyDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizePolicyAPI.AuthorizePolicyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 鉴权策略ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePolicyDeleteRequest struct via the builder pattern


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


## AuthorizePolicyPost

> CreatePostResultApiResponse AuthorizePolicyPost(ctx, appKey).AuthorizePolicy(authorizePolicy).Execute()

添加鉴权策略



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
	authorizePolicy := *openapiclient.NewAuthorizePolicy() // AuthorizePolicy | 鉴权策略对象 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizePolicyAPI.AuthorizePolicyPost(context.Background(), appKey).AuthorizePolicy(authorizePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizePolicyAPI.AuthorizePolicyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePolicyPost`: CreatePostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizePolicyAPI.AuthorizePolicyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePolicyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizePolicy** | [**AuthorizePolicy**](AuthorizePolicy.md) | 鉴权策略对象 | 

### Return type

[**CreatePostResultApiResponse**](CreatePostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorizePolicyPut

> BooleanApiResponse AuthorizePolicyPut(ctx, id, appKey).AuthorizePolicy(authorizePolicy).Execute()

更新鉴权策略



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
	id := int64(789) // int64 | 鉴权策略ID
	appKey := "appKey_example" // string | 
	authorizePolicy := *openapiclient.NewAuthorizePolicy() // AuthorizePolicy | 鉴权策略对象 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizePolicyAPI.AuthorizePolicyPut(context.Background(), id, appKey).AuthorizePolicy(authorizePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizePolicyAPI.AuthorizePolicyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizePolicyPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizePolicyAPI.AuthorizePolicyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 鉴权策略ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePolicyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizePolicy** | [**AuthorizePolicy**](AuthorizePolicy.md) | 鉴权策略对象 | 

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

