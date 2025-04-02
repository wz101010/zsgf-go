# \AlipayAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlipayCreateOrder**](AlipayAPI.md#AlipayCreateOrder) | **Post** /Alipay/{appKey}/CreateOrder | 创建订单 - 当面付
[**AlipayCreateOrderPagePay**](AlipayAPI.md#AlipayCreateOrderPagePay) | **Post** /Alipay/{appKey}/CreateOrderPagePay | 创建订单 - PC支付
[**AlipayCreateOrderWapPay**](AlipayAPI.md#AlipayCreateOrderWapPay) | **Post** /Alipay/{appKey}/CreateOrderWapPay | 创建订单 - WAP支付
[**AlipayOrderDetail**](AlipayAPI.md#AlipayOrderDetail) | **Get** /Alipay/{appKey}/OrderDetail | 订单详情
[**AlipayOrderRefund**](AlipayAPI.md#AlipayOrderRefund) | **Post** /Alipay/{appKey}/OrderRefund | 订单退款
[**AlipayReturnPageNotify**](AlipayAPI.md#AlipayReturnPageNotify) | **Post** /Alipay/{appKey}/ReturnPageNotify | 支付成功同步通知



## AlipayCreateOrder

> StringApiResponse AlipayCreateOrder(ctx, appKey).AlipayCreateOrderRequest(alipayCreateOrderRequest).Execute()

创建订单 - 当面付



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
	alipayCreateOrderRequest := *openapiclient.NewAlipayCreateOrderRequest("OrderNo_example", float64(123), "Subject_example") // AlipayCreateOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayCreateOrder(context.Background(), appKey).AlipayCreateOrderRequest(alipayCreateOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayCreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayCreateOrder`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayCreateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alipayCreateOrderRequest** | [**AlipayCreateOrderRequest**](AlipayCreateOrderRequest.md) |  | 

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


## AlipayCreateOrderPagePay

> StringApiResponse AlipayCreateOrderPagePay(ctx, appKey).AlipayCreateOrderPagePayRequest(alipayCreateOrderPagePayRequest).Execute()

创建订单 - PC支付



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
	alipayCreateOrderPagePayRequest := *openapiclient.NewAlipayCreateOrderPagePayRequest("OrderNo_example", float64(123), "Subject_example") // AlipayCreateOrderPagePayRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayCreateOrderPagePay(context.Background(), appKey).AlipayCreateOrderPagePayRequest(alipayCreateOrderPagePayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayCreateOrderPagePay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayCreateOrderPagePay`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayCreateOrderPagePay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayCreateOrderPagePayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alipayCreateOrderPagePayRequest** | [**AlipayCreateOrderPagePayRequest**](AlipayCreateOrderPagePayRequest.md) |  | 

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


## AlipayCreateOrderWapPay

> StringApiResponse AlipayCreateOrderWapPay(ctx, appKey).AlipayCreateOrderWapPayRequest(alipayCreateOrderWapPayRequest).Execute()

创建订单 - WAP支付



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
	alipayCreateOrderWapPayRequest := *openapiclient.NewAlipayCreateOrderWapPayRequest("OrderNo_example", float64(123), "Subject_example") // AlipayCreateOrderWapPayRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayCreateOrderWapPay(context.Background(), appKey).AlipayCreateOrderWapPayRequest(alipayCreateOrderWapPayRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayCreateOrderWapPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayCreateOrderWapPay`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayCreateOrderWapPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayCreateOrderWapPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alipayCreateOrderWapPayRequest** | [**AlipayCreateOrderWapPayRequest**](AlipayCreateOrderWapPayRequest.md) |  | 

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


## AlipayOrderDetail

> AlipayTradeQueryResponseApiResponse AlipayOrderDetail(ctx, appKey).OrderNo(orderNo).Execute()

订单详情



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
	orderNo := "orderNo_example" // string | 订单号 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayOrderDetail(context.Background(), appKey).OrderNo(orderNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayOrderDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayOrderDetail`: AlipayTradeQueryResponseApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayOrderDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayOrderDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderNo** | **string** | 订单号 | 

### Return type

[**AlipayTradeQueryResponseApiResponse**](AlipayTradeQueryResponseApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlipayOrderRefund

> AlipayTradeRefundResponseApiResponse AlipayOrderRefund(ctx, appKey).Amount(amount).OrderNo(orderNo).Execute()

订单退款



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
	amount := "amount_example" // string | 退款金额 (optional)
	orderNo := "orderNo_example" // string | 订单号 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayOrderRefund(context.Background(), appKey).Amount(amount).OrderNo(orderNo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayOrderRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayOrderRefund`: AlipayTradeRefundResponseApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayOrderRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayOrderRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **string** | 退款金额 | 
 **orderNo** | **string** | 订单号 | 

### Return type

[**AlipayTradeRefundResponseApiResponse**](AlipayTradeRefundResponseApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlipayReturnPageNotify

> BooleanApiResponse AlipayReturnPageNotify(ctx, appKey).ReturnPageNotifyRequest(returnPageNotifyRequest).Execute()

支付成功同步通知



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
	returnPageNotifyRequest := *openapiclient.NewReturnPageNotifyRequest() // ReturnPageNotifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlipayAPI.AlipayReturnPageNotify(context.Background(), appKey).ReturnPageNotifyRequest(returnPageNotifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlipayAPI.AlipayReturnPageNotify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlipayReturnPageNotify`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AlipayAPI.AlipayReturnPageNotify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlipayReturnPageNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnPageNotifyRequest** | [**ReturnPageNotifyRequest**](ReturnPageNotifyRequest.md) |  | 

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

