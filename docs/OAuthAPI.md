# \OAuthAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OAuthAuthorize**](OAuthAPI.md#OAuthAuthorize) | **Post** /OAuth/{appKey}/Authorize | 获取access_token
[**OAuthConsents**](OAuthAPI.md#OAuthConsents) | **Get** /OAuth/{appKey}/Consents | 授权记录
[**OAuthDeleteConsent**](OAuthAPI.md#OAuthDeleteConsent) | **Delete** /OAuth/{appKey}/Consents/{id} | 删除授权记录
[**OAuthGrantCode**](OAuthAPI.md#OAuthGrantCode) | **Post** /OAuth/{appKey}/GrantCode | 申请授权码
[**OAuthProfile**](OAuthAPI.md#OAuthProfile) | **Get** /OAuth/{appKey}/Profile | 获取个人资料



## OAuthAuthorize

> AuthorizeResultApiResponse OAuthAuthorize(ctx, appKey).Scheme(scheme).Code(code).Execute()

获取access_token

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
	scheme := "scheme_example" // string | 身份源 (optional)
	code := "code_example" // string | 授权码 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.OAuthAuthorize(context.Background(), appKey).Scheme(scheme).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.OAuthAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OAuthAuthorize`: AuthorizeResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.OAuthAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOAuthAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheme** | **string** | 身份源 | 
 **code** | **string** | 授权码 | 

### Return type

[**AuthorizeResultApiResponse**](AuthorizeResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OAuthConsents

> AppUserConsentModelListApiResponse OAuthConsents(ctx, appKey).Execute()

授权记录

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.OAuthConsents(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.OAuthConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OAuthConsents`: AppUserConsentModelListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.OAuthConsents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOAuthConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUserConsentModelListApiResponse**](AppUserConsentModelListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OAuthDeleteConsent

> BooleanApiResponse OAuthDeleteConsent(ctx, id, appKey).Execute()

删除授权记录

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
	id := int64(789) // int64 | 
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.OAuthDeleteConsent(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.OAuthDeleteConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OAuthDeleteConsent`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.OAuthDeleteConsent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOAuthDeleteConsentRequest struct via the builder pattern


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


## OAuthGrantCode

> GrantResultApiResponse OAuthGrantCode(ctx, appKey).Scheme(scheme).GrantRequest(grantRequest).Execute()

申请授权码

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
	scheme := "scheme_example" // string | 身份源，固定填：app (optional)
	grantRequest := *openapiclient.NewGrantRequest("GrantType_example", "Scopes_example") // GrantRequest | 授权详情 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.OAuthGrantCode(context.Background(), appKey).Scheme(scheme).GrantRequest(grantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.OAuthGrantCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OAuthGrantCode`: GrantResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.OAuthGrantCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOAuthGrantCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheme** | **string** | 身份源，固定填：app | 
 **grantRequest** | [**GrantRequest**](GrantRequest.md) | 授权详情 | 

### Return type

[**GrantResultApiResponse**](GrantResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OAuthProfile

> ProfileResultApiResponse OAuthProfile(ctx, appKey).Execute()

获取个人资料

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.OAuthProfile(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.OAuthProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OAuthProfile`: ProfileResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.OAuthProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOAuthProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProfileResultApiResponse**](ProfileResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

