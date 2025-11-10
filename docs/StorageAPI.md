# \StorageAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageAggregate**](StorageAPI.md#StorageAggregate) | **Get** /Storage/{appKey}/{table}/Aggregate | 聚合查询
[**StorageDelete**](StorageAPI.md#StorageDelete) | **Delete** /Storage/{appKey}/{table}/{id} | 删除数据
[**StorageDetail**](StorageAPI.md#StorageDetail) | **Get** /Storage/{appKey}/{table}/{id} | 数据详情
[**StorageList**](StorageAPI.md#StorageList) | **Get** /Storage/{appKey}/{table} | 查询数据
[**StoragePost**](StorageAPI.md#StoragePost) | **Post** /Storage/{appKey}/{table} | 添加数据
[**StoragePut**](StorageAPI.md#StoragePut) | **Put** /Storage/{appKey}/{table}/{id} | 更新数据



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


## StoragePut

> BooleanApiResponse StoragePut(ctx, table, id, appKey).RequestBody(requestBody).Replace(replace).Execute()

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
	requestBody := []interface{}{interface{}(123)} // []interface{} | json格式
	replace := true // bool | 是否为全量更新，默认为false (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StoragePut(context.Background(), table, id, appKey).RequestBody(requestBody).Replace(replace).Execute()
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



 **requestBody** | **[]interface{}** | json格式 | 
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

