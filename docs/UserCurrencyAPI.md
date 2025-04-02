# \UserCurrencyAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCurrencies**](UserCurrencyAPI.md#UserCurrencies) | **Get** /UserCurrency/{appKey}/{id} | 获取用户资产
[**UserCurrencyConsume**](UserCurrencyAPI.md#UserCurrencyConsume) | **Post** /UserCurrency/{appKey}/CurrencyConsume | 消费虚拟币
[**UserCurrencyExchange**](UserCurrencyAPI.md#UserCurrencyExchange) | **Post** /UserCurrency/{appKey}/CurrencyExchange | 兑换虚拟币
[**UserCurrencyRecharge**](UserCurrencyAPI.md#UserCurrencyRecharge) | **Post** /UserCurrency/{appKey}/CurrencyRecharge | 充值虚拟币
[**UserCurrencyTransactions**](UserCurrencyAPI.md#UserCurrencyTransactions) | **Get** /UserCurrency/{appKey}/CurrencyTransactions | 虚拟币交易记录



## UserCurrencies

> UserCurrencyListApiResponse UserCurrencies(ctx, id, appKey).Execute()

获取用户资产



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
	id := int64(789) // int64 | 用户ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCurrencyAPI.UserCurrencies(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCurrencyAPI.UserCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrencies`: UserCurrencyListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCurrencyAPI.UserCurrencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserCurrencyListApiResponse**](UserCurrencyListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCurrencyConsume

> BooleanApiResponse UserCurrencyConsume(ctx, appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).CurrencyConsumeRequest(currencyConsumeRequest).Execute()

消费虚拟币



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
	nonce := "nonce_example" // string | 随机数
	timestamp := int64(789) // int64 | 时间戳（允许与服务器时间误差在1分钟内）
	signature := "signature_example" // string | 签名
	appKey := "appKey_example" // string | 
	currencyConsumeRequest := *openapiclient.NewCurrencyConsumeRequest("Currency_example", int32(123)) // CurrencyConsumeRequest | 消费请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCurrencyAPI.UserCurrencyConsume(context.Background(), appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).CurrencyConsumeRequest(currencyConsumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCurrencyAPI.UserCurrencyConsume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrencyConsume`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCurrencyAPI.UserCurrencyConsume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrencyConsumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonce** | **string** | 随机数 | 
 **timestamp** | **int64** | 时间戳（允许与服务器时间误差在1分钟内） | 
 **signature** | **string** | 签名 | 

 **currencyConsumeRequest** | [**CurrencyConsumeRequest**](CurrencyConsumeRequest.md) | 消费请求参数 | 

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


## UserCurrencyExchange

> BooleanApiResponse UserCurrencyExchange(ctx, appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).ExchangeCurrencyRequest(exchangeCurrencyRequest).Execute()

兑换虚拟币



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
	nonce := "nonce_example" // string | 随机数
	timestamp := int64(789) // int64 | 时间戳（允许与服务器时间误差在1分钟内）
	signature := "signature_example" // string | 签名
	appKey := "appKey_example" // string | 
	exchangeCurrencyRequest := *openapiclient.NewExchangeCurrencyRequest("FromCurrency_example", "Currency_example", int32(123)) // ExchangeCurrencyRequest | 兑换请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCurrencyAPI.UserCurrencyExchange(context.Background(), appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).ExchangeCurrencyRequest(exchangeCurrencyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCurrencyAPI.UserCurrencyExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrencyExchange`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCurrencyAPI.UserCurrencyExchange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrencyExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonce** | **string** | 随机数 | 
 **timestamp** | **int64** | 时间戳（允许与服务器时间误差在1分钟内） | 
 **signature** | **string** | 签名 | 

 **exchangeCurrencyRequest** | [**ExchangeCurrencyRequest**](ExchangeCurrencyRequest.md) | 兑换请求参数 | 

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


## UserCurrencyRecharge

> BooleanApiResponse UserCurrencyRecharge(ctx, appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).RechargePointRequest(rechargePointRequest).Execute()

充值虚拟币



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
	nonce := "nonce_example" // string | 随机数
	timestamp := int64(789) // int64 | 时间戳（允许与服务器时间误差在1分钟内）
	signature := "signature_example" // string | 签名
	appKey := "appKey_example" // string | 
	rechargePointRequest := *openapiclient.NewRechargePointRequest("Currency_example", int32(123)) // RechargePointRequest | 充值请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCurrencyAPI.UserCurrencyRecharge(context.Background(), appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).RechargePointRequest(rechargePointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCurrencyAPI.UserCurrencyRecharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrencyRecharge`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCurrencyAPI.UserCurrencyRecharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrencyRechargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonce** | **string** | 随机数 | 
 **timestamp** | **int64** | 时间戳（允许与服务器时间误差在1分钟内） | 
 **signature** | **string** | 签名 | 

 **rechargePointRequest** | [**RechargePointRequest**](RechargePointRequest.md) | 充值请求参数 | 

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


## UserCurrencyTransactions

> UserCurrencyCurrencyTransResultApiResponse UserCurrencyTransactions(ctx, appKey).TransType(transType).CurCode(curCode).StartTime(startTime).EndTime(endTime).Skip(skip).Take(take).Execute()

虚拟币交易记录



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "//"
)

func main() {
	appKey := "appKey_example" // string | 
	transType := "transType_example" // string | 交易类型 (optional)
	curCode := "curCode_example" // string | 货币代码 (optional)
	startTime := time.Now() // time.Time | 开始时间 (optional)
	endTime := time.Now() // time.Time | 结束时间 (optional)
	skip := int32(56) // int32 | 跳过的条数 (optional)
	take := int32(56) // int32 | 拉取的条数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserCurrencyAPI.UserCurrencyTransactions(context.Background(), appKey).TransType(transType).CurCode(curCode).StartTime(startTime).EndTime(endTime).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserCurrencyAPI.UserCurrencyTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCurrencyTransactions`: UserCurrencyCurrencyTransResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserCurrencyAPI.UserCurrencyTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCurrencyTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transType** | **string** | 交易类型 | 
 **curCode** | **string** | 货币代码 | 
 **startTime** | **time.Time** | 开始时间 | 
 **endTime** | **time.Time** | 结束时间 | 
 **skip** | **int32** | 跳过的条数 | 
 **take** | **int32** | 拉取的条数 | 

### Return type

[**UserCurrencyCurrencyTransResultApiResponse**](UserCurrencyCurrencyTransResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

