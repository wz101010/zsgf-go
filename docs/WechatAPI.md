# \WechatAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatDecrypt**](WechatAPI.md#WechatDecrypt) | **Get** /Wechat/{appKey}/Decrypt | 小程序-解密数据
[**WechatGenerateScheme**](WechatAPI.md#WechatGenerateScheme) | **Post** /Wechat/{appKey}/GenerateScheme | 小程序-生成scheme码(该接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景)
[**WechatJSCode2Session**](WechatAPI.md#WechatJSCode2Session) | **Get** /Wechat/{appKey}/JSCode2Session | 小程序-登录凭证校验
[**WechatJSConfig**](WechatAPI.md#WechatJSConfig) | **Get** /Wechat/{appKey}/JSConfig | 公众号H5-JS SDK Config
[**WechatSubscribeMSG**](WechatAPI.md#WechatSubscribeMSG) | **Post** /Wechat/{appKey}/SubscribeMSG | 公众号H5-发送一次性订阅消息
[**WechatSubscribeSend**](WechatAPI.md#WechatSubscribeSend) | **Post** /Wechat/{appKey}/SubscribeSend | 小程序-发送订阅消息
[**WechatUrlLinkGenerate**](WechatAPI.md#WechatUrlLinkGenerate) | **Post** /Wechat/{appKey}/UrlLinkGenerate | 小程序-生成网页跳转地址(获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景)
[**WechatUserInfo**](WechatAPI.md#WechatUserInfo) | **Get** /Wechat/{appKey}/UserInfo | 公众号H5-获取用户UnionID
[**WechatWXACodeGet**](WechatAPI.md#WechatWXACodeGet) | **Post** /Wechat/{appKey}/WXACodeGet | 小程序-获取小程序码
[**WechatWXACodeGetUnlimited**](WechatAPI.md#WechatWXACodeGetUnlimited) | **Post** /Wechat/{appKey}/WXACodeGetUnlimited | 小程序-获取小程序码(无限制)



## WechatDecrypt

> StringApiResponse WechatDecrypt(ctx, appKey).EncryptedData(encryptedData).Iv(iv).SessionKey(sessionKey).Execute()

小程序-解密数据



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

> StringApiResponse WechatGenerateScheme(ctx, appKey).Body(body).Execute()

小程序-生成scheme码(该接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景)



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
	body := interface{}(987) // interface{} | scheme码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatGenerateScheme(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | scheme码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-scheme/generateScheme.html | 

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

小程序-登录凭证校验



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

公众号H5-JS SDK Config



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


## WechatSubscribeMSG

> StringApiResponse WechatSubscribeMSG(ctx, appKey).Body(body).Execute()

公众号H5-发送一次性订阅消息



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
	body := interface{}(987) // interface{} | 订阅消息数据，开发参考：https://dwz.cn/IXptek5n (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatSubscribeMSG(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | 订阅消息数据，开发参考：https://dwz.cn/IXptek5n | 

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

> StringApiResponse WechatSubscribeSend(ctx, appKey).Body(body).Execute()

小程序-发送订阅消息



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
	body := interface{}(987) // interface{} | 订阅消息数据，开发参考：https://dwz.cn/bohXaCnp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatSubscribeSend(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | 订阅消息数据，开发参考：https://dwz.cn/bohXaCnp | 

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

> StringApiResponse WechatUrlLinkGenerate(ctx, appKey).Body(body).Execute()

小程序-生成网页跳转地址(获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景)



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
	body := interface{}(987) // interface{} | 跳转地址数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatUrlLinkGenerate(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | 跳转地址数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html | 

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

公众号H5-获取用户UnionID



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

> *os.File WechatWXACodeGet(ctx, appKey).Body(body).Execute()

小程序-获取小程序码



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
	body := interface{}(987) // interface{} | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatWXACodeGet(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: image/jpeg, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WechatWXACodeGetUnlimited

> *os.File WechatWXACodeGetUnlimited(ctx, appKey).Body(body).Execute()

小程序-获取小程序码(无限制)



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
	body := interface{}(987) // interface{} | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WechatAPI.WechatWXACodeGetUnlimited(context.Background(), appKey).Body(body).Execute()
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

 **body** | **interface{}** | 小程序码数据，开发参考：https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: image/jpeg, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

