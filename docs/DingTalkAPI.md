# \DingTalkAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DingTalkUserInfo**](DingTalkAPI.md#DingTalkUserInfo) | **Get** /DingTalk/{appKey}/UserInfo | 获取用户资料



## DingTalkUserInfo

> StringApiResponse DingTalkUserInfo(ctx, appKey).Code(code).Execute()

获取用户资料



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
	code := "code_example" // string | 钉钉用户授权码 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DingTalkAPI.DingTalkUserInfo(context.Background(), appKey).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DingTalkAPI.DingTalkUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DingTalkUserInfo`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `DingTalkAPI.DingTalkUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDingTalkUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** | 钉钉用户授权码 | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

