# \StorageAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageAggregate**](StorageAPI.md#StorageAggregate) | **Get** /Storage/{appKey}/{table}/Aggregate | 聚合查询
[**StorageClear**](StorageAPI.md#StorageClear) | **Delete** /Storage/{appKey}/{table}/Clear | 清空表数据
[**StorageDelete**](StorageAPI.md#StorageDelete) | **Delete** /Storage/{appKey}/{table}/{id} | 删除数据
[**StorageDeleteIndex**](StorageAPI.md#StorageDeleteIndex) | **Delete** /Storage/{appKey}/{table}/indexes | 删除索引
[**StorageDetail**](StorageAPI.md#StorageDetail) | **Get** /Storage/{appKey}/{table}/{id} | 数据详情
[**StorageExecuteFunction**](StorageAPI.md#StorageExecuteFunction) | **Get** /Storage/{appKey}/ExecuteFunction | 执行函数
[**StorageExport**](StorageAPI.md#StorageExport) | **Get** /Storage/{appKey}/{table}/Export | 导出数据
[**StorageImport**](StorageAPI.md#StorageImport) | **Post** /Storage/{appKey}/{table}/Import | 导入数据
[**StorageIndexes**](StorageAPI.md#StorageIndexes) | **Get** /Storage/{appKey}/{table}/Indexes | 获取索引
[**StorageList**](StorageAPI.md#StorageList) | **Get** /Storage/{appKey}/{table} | 查询数据
[**StoragePost**](StorageAPI.md#StoragePost) | **Post** /Storage/{appKey}/{table} | 添加数据
[**StoragePostIndex**](StorageAPI.md#StoragePostIndex) | **Post** /Storage/{appKey}/{table}/indexes | 添加索引
[**StoragePut**](StorageAPI.md#StoragePut) | **Put** /Storage/{appKey}/{table}/{id} | 更新数据
[**StorageRemove**](StorageAPI.md#StorageRemove) | **Delete** /Storage/{appKey}/{table}/Remove | 删除表
[**StorageStats**](StorageAPI.md#StorageStats) | **Get** /Storage/{appKey}/{table}/Stats | 数据表统计
[**StorageTables**](StorageAPI.md#StorageTables) | **Get** /Storage/{appKey}/Tables | 获取数据表



## StorageAggregate

> ObjectListApiResponse StorageAggregate(ctx, table, appKey).Pipeline(pipeline).Execute()

聚合查询



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	pipeline := "pipeline_example" // string | 构建聚合查询 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageAggregate(context.Background(), table, appKey).Pipeline(pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageAggregate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageAggregate`: ObjectListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageAggregate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageAggregateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipeline** | **string** | 构建聚合查询 | 

### Return type

[**ObjectListApiResponse**](ObjectListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageClear

> Int64ApiResponse StorageClear(ctx, table, appKey).Filter(filter).Execute()

清空表数据



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	filter := "filter_example" // string | 筛选条件，json格式 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageClear(context.Background(), table, appKey).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageClear`: Int64ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageClear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | 筛选条件，json格式 | 

### Return type

[**Int64ApiResponse**](Int64ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageDelete

> BooleanApiResponse StorageDelete(ctx, table, id, appKey).Execute()

删除数据



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
	table := "table_example" // string | 表名称
	id := "id_example" // string | 数据ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageDelete(context.Background(), table, id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**id** | **string** | 数据ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageDeleteRequest struct via the builder pattern


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


## StorageDeleteIndex

> BooleanApiResponse StorageDeleteIndex(ctx, table, appKey).IndexName(indexName).Execute()

删除索引



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	indexName := "indexName_example" // string | 索引名称 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageDeleteIndex(context.Background(), table, appKey).IndexName(indexName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageDeleteIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageDeleteIndex`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageDeleteIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageDeleteIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **indexName** | **string** | 索引名称 | 

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


## StorageDetail

> ObjectApiResponse StorageDetail(ctx, table, id, appKey).Project(project).Execute()

数据详情



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
	table := "table_example" // string | 表名称
	id := "id_example" // string | 数据ID
	appKey := "appKey_example" // string | 
	project := "project_example" // string | json格式 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageDetail(context.Background(), table, id, appKey).Project(project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageDetail`: ObjectApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**id** | **string** | 数据ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **project** | **string** | json格式 | 

### Return type

[**ObjectApiResponse**](ObjectApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageExecuteFunction

> ObjectApiResponse StorageExecuteFunction(ctx, appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).ExecuteFunctionRequest(executeFunctionRequest).Execute()

执行函数



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
	nonce := "nonce_example" // string | 随机字符串
	timestamp := int64(789) // int64 | 时间戳
	signature := "signature_example" // string | 签名
	appKey := "appKey_example" // string | 
	executeFunctionRequest := *openapiclient.NewExecuteFunctionRequest() // ExecuteFunctionRequest | 函数请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageExecuteFunction(context.Background(), appKey).Nonce(nonce).Timestamp(timestamp).Signature(signature).ExecuteFunctionRequest(executeFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageExecuteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageExecuteFunction`: ObjectApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageExecuteFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageExecuteFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonce** | **string** | 随机字符串 | 
 **timestamp** | **int64** | 时间戳 | 
 **signature** | **string** | 签名 | 

 **executeFunctionRequest** | [**ExecuteFunctionRequest**](ExecuteFunctionRequest.md) | 函数请求参数 | 

### Return type

[**ObjectApiResponse**](ObjectApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageExport

> *os.File StorageExport(ctx, table, appKey).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()

导出数据



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	filter := "filter_example" // string | 筛选条件，json格式 (optional)
	startTime := time.Now() // time.Time | 开始时间 (optional)
	endTime := time.Now() // time.Time | 结束时间 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageExport(context.Background(), table, appKey).Filter(filter).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | 筛选条件，json格式 | 
 **startTime** | **time.Time** | 开始时间 | 
 **endTime** | **time.Time** | 结束时间 | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageImport

> BooleanApiResponse StorageImport(ctx, table, appKey).File(file).Execute()

导入数据



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 导入的文件 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageImport(context.Background(), table, appKey).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageImport`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | 导入的文件 | 

### Return type

[**BooleanApiResponse**](BooleanApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageIndexes

> ObjectListApiResponse StorageIndexes(ctx, table, appKey).Execute()

获取索引



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageIndexes(context.Background(), table, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageIndexes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageIndexes`: ObjectListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageIndexes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageIndexesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectListApiResponse**](ObjectListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageList

> StorageListResultApiResponse StorageList(ctx, table, appKey).Filter(filter).Project(project).Sort(sort).StartTime(startTime).EndTime(endTime).Explain(explain).Take(take).Skip(skip).Execute()

查询数据



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 
	filter := "filter_example" // string | 过滤，json格式 (optional)
	project := "project_example" // string | 字段映射，json格式 (optional)
	sort := "sort_example" // string | 排序，json格式 (optional)
	startTime := time.Now() // time.Time | 开始时间 (optional)
	endTime := time.Now() // time.Time | 结束时间 (optional)
	explain := true // bool | 查看执行计划 (optional) (default to false)
	take := int32(56) // int32 | 默认为10 (optional) (default to 10)
	skip := int32(56) // int32 | 默认为0 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageList(context.Background(), table, appKey).Filter(filter).Project(project).Sort(sort).StartTime(startTime).EndTime(endTime).Explain(explain).Take(take).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageList`: StorageListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | 过滤，json格式 | 
 **project** | **string** | 字段映射，json格式 | 
 **sort** | **string** | 排序，json格式 | 
 **startTime** | **time.Time** | 开始时间 | 
 **endTime** | **time.Time** | 结束时间 | 
 **explain** | **bool** | 查看执行计划 | [default to false]
 **take** | **int32** | 默认为10 | [default to 10]
 **skip** | **int32** | 默认为0 | [default to 0]

### Return type

[**StorageListResultApiResponse**](StorageListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoragePost

> StringApiResponse StoragePost(ctx, table, appKey).RequestBody(requestBody).Execute()

添加数据



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
	table := "table_example" // string | 表名称（小写字母加数字,1-50位）
	appKey := "appKey_example" // string | 
	requestBody := []interface{}{interface{}(123)} // []interface{} | json对象 / json数组

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StoragePost(context.Background(), table, appKey).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StoragePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragePost`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StoragePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称（小写字母加数字,1-50位） | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]interface{}** | json对象 / json数组 | 

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


## StoragePostIndex

> StringApiResponse StoragePostIndex(ctx, table, appKey).PostIndexRequest(postIndexRequest).Execute()

添加索引



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
	table := "table_example" // string | 表名称（小写字母加数字,1-50位）
	appKey := "appKey_example" // string | 
	postIndexRequest := *openapiclient.NewPostIndexRequest(map[string][]interface{}{"key": []interface{}{nil}}) // PostIndexRequest | json对象 / json数组 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StoragePostIndex(context.Background(), table, appKey).PostIndexRequest(postIndexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StoragePostIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragePostIndex`: StringApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StoragePostIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称（小写字母加数字,1-50位） | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragePostIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postIndexRequest** | [**PostIndexRequest**](PostIndexRequest.md) | json对象 / json数组 | 

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


## StoragePut

> BooleanApiResponse StoragePut(ctx, table, id, appKey).Body(body).Replace(replace).Execute()

更新数据



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
	table := "table_example" // string | 表名称
	id := "id_example" // string | 数据ID
	appKey := "appKey_example" // string | 
	body := interface{}(987) // interface{} | json格式
	replace := true // bool | 是否为全量更新，默认为false (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StoragePut(context.Background(), table, id, appKey).Body(body).Replace(replace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StoragePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoragePut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StoragePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**id** | **string** | 数据ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoragePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **interface{}** | json格式 | 
 **replace** | **bool** | 是否为全量更新，默认为false | [default to false]

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


## StorageRemove

> BooleanApiResponse StorageRemove(ctx, table, appKey).Execute()

删除表



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageRemove(context.Background(), table, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageRemove`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageRemoveRequest struct via the builder pattern


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


## StorageStats

> ObjectApiResponse StorageStats(ctx, table, appKey).Execute()

数据表统计



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
	table := "table_example" // string | 表名称
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageStats(context.Background(), table, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageStats`: ObjectApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | 表名称 | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObjectApiResponse**](ObjectApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageTables

> StringListApiResponse StorageTables(ctx, appKey).Execute()

获取数据表



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
	resp, r, err := apiClient.StorageAPI.StorageTables(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageTables`: StringListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StringListApiResponse**](StringListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

