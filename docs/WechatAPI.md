# \WechatAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmQRCodeLogin**](WechatAPI.md#ConfirmQRCodeLogin) | **Post** /Wechat/{appKey}/QR-Auth/Confirm-Login | 确认二维码登录请求
[**ConfirmQRCodeRegistration**](WechatAPI.md#ConfirmQRCodeRegistration) | **Post** /Wechat/{appKey}/QR-Auth/Confirm-Register | 确认二维码注册请求
[**InitiateQRAuthSession**](WechatAPI.md#InitiateQRAuthSession) | **Post** /Wechat/{appKey}/QR-Auth/Initiate | 初始化二维码认证会话
[**ScanQRCodeForAuth**](WechatAPI.md#ScanQRCodeForAuth) | **Post** /Wechat/{appKey}/QR-Auth/Scan | 验证二维码扫描结果
[**WechatDecrypt**](WechatAPI.md#WechatDecrypt) | **Get** /Wechat/{appKey}/Decrypt | 解密小程序用户数据
[**WechatGenerateScheme**](WechatAPI.md#WechatGenerateScheme) | **Post** /Wechat/{appKey}/GenerateScheme | 生成小程序Scheme码
[**WechatJSCode2Session**](WechatAPI.md#WechatJSCode2Session) | **Get** /Wechat/{appKey}/JSCode2Session | 校验小程序登录状态
[**WechatJSConfig**](WechatAPI.md#WechatJSConfig) | **Get** /Wechat/{appKey}/JSConfig | 配置公众号JS SDK
[**WechatMsgSecCheck**](WechatAPI.md#WechatMsgSecCheck) | **Post** /Wechat/{appKey}/MsgSecCheck | 小程序内容安全检测
[**WechatSubscribeMSG**](WechatAPI.md#WechatSubscribeMSG) | **Post** /Wechat/{appKey}/SubscribeMSG | 发送公众号一次性订阅消息
[**WechatSubscribeSend**](WechatAPI.md#WechatSubscribeSend) | **Post** /Wechat/{appKey}/SubscribeSend | 发送小程序订阅消息
[**WechatUrlLinkGenerate**](WechatAPI.md#WechatUrlLinkGenerate) | **Post** /Wechat/{appKey}/UrlLinkGenerate | 生成小程序URL跳转链接
[**WechatUserInfo**](WechatAPI.md#WechatUserInfo) | **Get** /Wechat/{appKey}/UserInfo | 获取公众号H5 UnionID
[**WechatWXACodeGet**](WechatAPI.md#WechatWXACodeGet) | **Post** /Wechat/{appKey}/WXACodeGet | 获取小程序码（普通）
[**WechatWXACodeGetUnlimited**](WechatAPI.md#WechatWXACodeGetUnlimited) | **Post** /Wechat/{appKey}/WXACodeGetUnlimited | 获取小程序码（无限制）



## ConfirmQRCodeLogin

> TokenModelApiResponse ConfirmQRCodeLogin(ctx, appKey).QRCodeSignInRequest(qRCodeSignInRequest).Execute()

确认二维码登录请求



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
	qRCodeSignInRequest := *openapiclient.NewQRCodeSignInRequest(int64(123)) // QRCodeSignInRequest | 登录确认请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.ConfirmQRCodeLogin(context.Background(), appKey).QRCodeSignInRequest(qRCodeSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.ConfirmQRCodeLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmQRCodeLogin`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.ConfirmQRCodeLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmQRCodeLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeSignInRequest** | [**QRCodeSignInRequest**](QRCodeSignInRequest.md) | 登录确认请求参数 | 

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


## ConfirmQRCodeRegistration

> TokenModelApiResponse ConfirmQRCodeRegistration(ctx, appKey).QRCodeSignUpRequest(qRCodeSignUpRequest).Execute()

确认二维码注册请求



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
	qRCodeSignUpRequest := *openapiclient.NewQRCodeSignUpRequest("UnionID_example") // QRCodeSignUpRequest | 注册确认请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.ConfirmQRCodeRegistration(context.Background(), appKey).QRCodeSignUpRequest(qRCodeSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.ConfirmQRCodeRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmQRCodeRegistration`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.ConfirmQRCodeRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmQRCodeRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeSignUpRequest** | [**QRCodeSignUpRequest**](QRCodeSignUpRequest.md) | 注册确认请求参数 | 

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


## InitiateQRAuthSession

> Int64ApiResponse InitiateQRAuthSession(ctx, appKey).QRCodePreSignInRequest(qRCodePreSignInRequest).Execute()

初始化二维码认证会话



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
	qRCodePreSignInRequest := *openapiclient.NewQRCodePreSignInRequest() // QRCodePreSignInRequest | 认证会话初始化请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.InitiateQRAuthSession(context.Background(), appKey).QRCodePreSignInRequest(qRCodePreSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.InitiateQRAuthSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateQRAuthSession`: Int64ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.InitiateQRAuthSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateQRAuthSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodePreSignInRequest** | [**QRCodePreSignInRequest**](QRCodePreSignInRequest.md) | 认证会话初始化请求参数 | 

### Return type

[**Int64ApiResponse**](Int64ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanQRCodeForAuth

> UserQRCodeScanResultApiResponse ScanQRCodeForAuth(ctx, appKey).QRCodeScanRequest(qRCodeScanRequest).Execute()

验证二维码扫描结果



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
	qRCodeScanRequest := *openapiclient.NewQRCodeScanRequest(int64(123)) // QRCodeScanRequest | 二维码扫描请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.ScanQRCodeForAuth(context.Background(), appKey).QRCodeScanRequest(qRCodeScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.ScanQRCodeForAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanQRCodeForAuth`: UserQRCodeScanResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.ScanQRCodeForAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanQRCodeForAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeScanRequest** | [**QRCodeScanRequest**](QRCodeScanRequest.md) | 二维码扫描请求参数 | 

### Return type

[**UserQRCodeScanResultApiResponse**](UserQRCodeScanResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatDecrypt

> StringApiResponse WechatDecrypt(ctx, appKey).EncryptedData(encryptedData).Iv(iv).SessionKey(sessionKey).Execute()

解密小程序用户数据



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
	encryptedData := "encryptedData_example" // string | 加密的数据 (optional)
	iv := "iv_example" // string | 加密算法的初始向量 (optional)
	sessionKey := "sessionKey_example" // string | 会话密钥 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatDecrypt(context.Background(), appKey).EncryptedData(encryptedData).Iv(iv).SessionKey(sessionKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatDecrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatDecrypt`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatDecrypt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatDecryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encryptedData** | **string** | 加密的数据 | 
 **iv** | **string** | 加密算法的初始向量 | 
 **sessionKey** | **string** | 会话密钥 | 

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


## WechatGenerateScheme

> StringApiResponse WechatGenerateScheme(ctx, appKey).RequestBody(requestBody).Execute()

生成小程序Scheme码



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | scheme码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatGenerateScheme(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatGenerateScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatGenerateScheme`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatGenerateScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatGenerateSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | scheme码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatJSCode2Session

> StringApiResponse WechatJSCode2Session(ctx, appKey).JsCode(jsCode).Execute()

校验小程序登录状态



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
	jsCode := "jsCode_example" // string | 登录凭证，开发参考：https://dwz.cn/icNajFh7 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatJSCode2Session(context.Background(), appKey).JsCode(jsCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatJSCode2Session``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatJSCode2Session`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatJSCode2Session`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatJSCode2SessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsCode** | **string** | 登录凭证，开发参考：https://dwz.cn/icNajFh7 | 

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


## WechatJSConfig

> WechatJSConfigResultApiResponse WechatJSConfig(ctx, appKey).Url(url).Execute()

配置公众号JS SDK



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
	url := "url_example" // string | 当前网页的URL (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatJSConfig(context.Background(), appKey).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatJSConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatJSConfig`: WechatJSConfigResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatJSConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatJSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **string** | 当前网页的URL | 

### Return type

[**WechatJSConfigResultApiResponse**](WechatJSConfigResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatMsgSecCheck

> interface{} WechatMsgSecCheck(ctx, appKey).RequestBody(requestBody).Execute()

小程序内容安全检测



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 消息内容数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatMsgSecCheck(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatMsgSecCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatMsgSecCheck`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatMsgSecCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatMsgSecCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 消息内容数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/msgSecCheck.html | 

### Return type

**interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatSubscribeMSG

> StringApiResponse WechatSubscribeMSG(ctx, appKey).RequestBody(requestBody).Execute()

发送公众号一次性订阅消息



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 订阅消息数据，开发参考：https://dwz.cn/IXptek5n (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatSubscribeMSG(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatSubscribeMSG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatSubscribeMSG`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatSubscribeMSG`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatSubscribeMSGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 订阅消息数据，开发参考：https://dwz.cn/IXptek5n | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatSubscribeSend

> StringApiResponse WechatSubscribeSend(ctx, appKey).RequestBody(requestBody).Execute()

发送小程序订阅消息



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 订阅消息数据，开发参考：https://dwz.cn/bohXaCnp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatSubscribeSend(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatSubscribeSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatSubscribeSend`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatSubscribeSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatSubscribeSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 订阅消息数据，开发参考：https://dwz.cn/bohXaCnp | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatUrlLinkGenerate

> StringApiResponse WechatUrlLinkGenerate(ctx, appKey).RequestBody(requestBody).Execute()

生成小程序URL跳转链接



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 跳转地址数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatUrlLinkGenerate(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatUrlLinkGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatUrlLinkGenerate`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatUrlLinkGenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatUrlLinkGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 跳转地址数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html | 

### Return type

[**StringApiResponse**](StringApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatUserInfo

> StringApiResponse WechatUserInfo(ctx, appKey).Openid(openid).Execute()

获取公众号H5 UnionID



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
	openid := "openid_example" // string | 用户的OpenID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatUserInfo(context.Background(), appKey).Openid(openid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatUserInfo`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatUserInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openid** | **string** | 用户的OpenID | 

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


## WechatWXACodeGet

> *os.File WechatWXACodeGet(ctx, appKey).RequestBody(requestBody).Execute()

获取小程序码（普通）



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatWXACodeGet(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatWXACodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatWXACodeGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatWXACodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatWXACodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatWXACodeGetUnlimited

> *os.File WechatWXACodeGetUnlimited(ctx, appKey).RequestBody(requestBody).Execute()

获取小程序码（无限制）



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
	requestBody := []interface{}{interface{}(123)} // []interface{} | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatWXACodeGetUnlimited(context.Background(), appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WechatAPI.WechatWXACodeGetUnlimited``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WechatWXACodeGetUnlimited`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `WechatAPI.WechatWXACodeGetUnlimited`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWechatWXACodeGetUnlimitedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]interface{}** | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

