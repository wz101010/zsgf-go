# \OrderAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Order**](OrderAPI.md#Order) | **Get** /Order/{appKey}/{id} | 获取订单详情
[**OrderCreate**](OrderAPI.md#OrderCreate) | **Post** /Order/{appKey}/Create | 创建订单
[**Orders**](OrderAPI.md#Orders) | **Get** /Order/{appKey} | 获取订单列表



## Order

> OrderApiResponse Order(ctx, id, appKey).Execute()

获取订单详情



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
	id := int64(789) // int64 | 订单ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.Order(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.Order``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Order`: OrderApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.Order`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 订单ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrderApiResponse**](OrderApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCreate

> CreateOrderResultApiResponse OrderCreate(ctx, appKey).CreateOrderRequest(createOrderRequest).Execute()

创建订单



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
	createOrderRequest := *openapiclient.NewCreateOrderRequest(float64(123), "ProductName_example", "ProductType_example", "ProductID_example") // CreateOrderRequest | 订单创建请求 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.OrderCreate(context.Background(), appKey).CreateOrderRequest(createOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.OrderCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderCreate`: CreateOrderResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.OrderCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrderRequest** | [**CreateOrderRequest**](CreateOrderRequest.md) | 订单创建请求 | 

### Return type

[**CreateOrderResultApiResponse**](CreateOrderResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Orders

> OrderListResultApiResponse Orders(ctx, appKey).Status(status).OrderNo(orderNo).TradeNo(tradeNo).UserId(userId).PctType(pctType).PctId(pctId).PctName(pctName).Skip(skip).Take(take).Execute()

获取订单列表



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
	status := "status_example" // string | 订单状态 (optional)
	orderNo := "orderNo_example" // string | 系统订单号 (optional)
	tradeNo := "tradeNo_example" // string | 支付平台单号 (optional)
	userId := int64(789) // int64 | 用户ID (optional)
	pctType := "pctType_example" // string | 商品类型 (optional)
	pctId := "pctId_example" // string | 商品ID (optional)
	pctName := "pctName_example" // string | 商品名称 (optional)
	skip := int32(56) // int32 | 跳过的条数 (optional)
	take := int32(56) // int32 | 拉取的条数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.Orders(context.Background(), appKey).Status(status).OrderNo(orderNo).TradeNo(tradeNo).UserId(userId).PctType(pctType).PctId(pctId).PctName(pctName).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.Orders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Orders`: OrderListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.Orders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | 订单状态 | 
 **orderNo** | **string** | 系统订单号 | 
 **tradeNo** | **string** | 支付平台单号 | 
 **userId** | **int64** | 用户ID | 
 **pctType** | **string** | 商品类型 | 
 **pctId** | **string** | 商品ID | 
 **pctName** | **string** | 商品名称 | 
 **skip** | **int32** | 跳过的条数 | 
 **take** | **int32** | 拉取的条数 | 

### Return type

[**OrderListResultApiResponse**](OrderListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

