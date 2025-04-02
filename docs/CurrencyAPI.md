# \CurrencyAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Currencies**](CurrencyAPI.md#Currencies) | **Get** /Currency/{appKey} | 获取货币列表
[**Currency**](CurrencyAPI.md#Currency) | **Get** /Currency/{appKey}/{id} | 获取货币详情
[**CurrencyDelete**](CurrencyAPI.md#CurrencyDelete) | **Delete** /Currency/{appKey}/{id} | 删除货币
[**CurrencyExchangeRateDelete**](CurrencyAPI.md#CurrencyExchangeRateDelete) | **Delete** /Currency/{appKey}/ExchangeRates/{id} | 删除汇率
[**CurrencyExchangeRatePut**](CurrencyAPI.md#CurrencyExchangeRatePut) | **Put** /Currency/{appKey}/ExchangeRates/{code} | 更新汇率
[**CurrencyExchangeRates**](CurrencyAPI.md#CurrencyExchangeRates) | **Get** /Currency/{appKey}/ExchangeRates/{code} | 获取汇率列表
[**CurrencyPost**](CurrencyAPI.md#CurrencyPost) | **Post** /Currency/{appKey} | 创建新货币
[**CurrencyPut**](CurrencyAPI.md#CurrencyPut) | **Put** /Currency/{appKey}/{id} | 更新货币信息
[**CurrencyTransactions**](CurrencyAPI.md#CurrencyTransactions) | **Get** /Currency/{appKey}/Transactions | 获取货币交易记录



## Currencies

> CurrencyListApiResponse Currencies(ctx, appKey).Execute()

获取货币列表



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
	resp, r, err := apiClient.CurrencyAPI.Currencies(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.Currencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Currencies`: CurrencyListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.Currencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencyListApiResponse**](CurrencyListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Currency

> CurrencyApiResponse Currency(ctx, id, appKey).Execute()

获取货币详情



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
	id := int64(789) // int64 | 货币ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.Currency(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.Currency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Currency`: CurrencyApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.Currency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 货币ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrencyApiResponse**](CurrencyApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrencyDelete

> BooleanApiResponse CurrencyDelete(ctx, id, appKey).Execute()

删除货币



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
	id := int64(789) // int64 | 货币ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 货币ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyDeleteRequest struct via the builder pattern


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


## CurrencyExchangeRateDelete

> BooleanApiResponse CurrencyExchangeRateDelete(ctx, id, appKey).Execute()

删除汇率



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
	id := int64(789) // int64 | 汇率ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyExchangeRateDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyExchangeRateDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyExchangeRateDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyExchangeRateDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 汇率ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyExchangeRateDeleteRequest struct via the builder pattern


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


## CurrencyExchangeRatePut

> Int64ApiResponse CurrencyExchangeRatePut(ctx, code, appKey).ExchangeRatePutRequest(exchangeRatePutRequest).Execute()

更新汇率



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
	code := "code_example" // string | 货币代码
	appKey := "appKey_example" // string | 
	exchangeRatePutRequest := *openapiclient.NewExchangeRatePutRequest() // ExchangeRatePutRequest | 汇率信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyExchangeRatePut(context.Background(), code, appKey).ExchangeRatePutRequest(exchangeRatePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyExchangeRatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyExchangeRatePut`: Int64ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyExchangeRatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | 货币代码 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyExchangeRatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exchangeRatePutRequest** | [**ExchangeRatePutRequest**](ExchangeRatePutRequest.md) | 汇率信息 | 

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


## CurrencyExchangeRates

> CurrencyExchangeRateApiResponse CurrencyExchangeRates(ctx, code, appKey).Execute()

获取汇率列表



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
	code := "code_example" // string | 货币代码
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyExchangeRates(context.Background(), code, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyExchangeRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyExchangeRates`: CurrencyExchangeRateApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyExchangeRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | 货币代码 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyExchangeRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CurrencyExchangeRateApiResponse**](CurrencyExchangeRateApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrencyPost

> Int64ApiResponse CurrencyPost(ctx, appKey).Currency(currency).Execute()

创建新货币



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
	currency := *openapiclient.NewCurrency() // Currency | 货币信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyPost(context.Background(), appKey).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyPost`: Int64ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currency** | [**Currency**](Currency.md) | 货币信息 | 

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


## CurrencyPut

> BooleanApiResponse CurrencyPut(ctx, id, appKey).Currency(currency).Execute()

更新货币信息



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
	id := int64(789) // int64 | 货币ID
	appKey := "appKey_example" // string | 
	currency := *openapiclient.NewCurrency() // Currency | 货币信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyPut(context.Background(), id, appKey).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 货币ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **currency** | [**Currency**](Currency.md) | 货币信息 | 

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


## CurrencyTransactions

> CurrencyTransactionListApiResponse CurrencyTransactions(ctx, appKey).UserId(userId).TransType(transType).CurCode(curCode).StartTime(startTime).EndTime(endTime).Skip(skip).Take(take).Execute()

获取货币交易记录



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
	userId := int64(789) // int64 | 用户ID (optional)
	transType := "transType_example" // string | 交易类型 (optional)
	curCode := "curCode_example" // string | 货币代码 (optional)
	startTime := time.Now() // time.Time | 开始时间 (optional)
	endTime := time.Now() // time.Time | 结束时间 (optional)
	skip := int32(56) // int32 | 跳过的条数 (optional)
	take := int32(56) // int32 | 拉取的条数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurrencyAPI.CurrencyTransactions(context.Background(), appKey).UserId(userId).TransType(transType).CurCode(curCode).StartTime(startTime).EndTime(endTime).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrencyAPI.CurrencyTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrencyTransactions`: CurrencyTransactionListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CurrencyAPI.CurrencyTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrencyTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int64** | 用户ID | 
 **transType** | **string** | 交易类型 | 
 **curCode** | **string** | 货币代码 | 
 **startTime** | **time.Time** | 开始时间 | 
 **endTime** | **time.Time** | 结束时间 | 
 **skip** | **int32** | 跳过的条数 | 
 **take** | **int32** | 拉取的条数 | 

### Return type

[**CurrencyTransactionListApiResponse**](CurrencyTransactionListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

