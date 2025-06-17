# \UserLocationAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserLocation**](UserLocationAPI.md#UserLocation) | **Get** /UserLocation/{appKey}/{id} | 获取位置详情
[**UserLocationDelete**](UserLocationAPI.md#UserLocationDelete) | **Delete** /UserLocation/{appKey}/{id} | 删除位置
[**UserLocationPost**](UserLocationAPI.md#UserLocationPost) | **Post** /UserLocation/{appKey} | 添加位置
[**UserLocationPut**](UserLocationAPI.md#UserLocationPut) | **Put** /UserLocation/{appKey}/{id} | 更新位置
[**UserLocations**](UserLocationAPI.md#UserLocations) | **Get** /UserLocation/{appKey} | 获取位置列表



## UserLocation

> GeoLocationModelApiResponse UserLocation(ctx, id, appKey).Execute()

获取位置详情



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
	id := int64(789) // int64 | 位置ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLocationAPI.UserLocation(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLocationAPI.UserLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocation`: GeoLocationModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserLocationAPI.UserLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 位置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GeoLocationModelApiResponse**](GeoLocationModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserLocationDelete

> BooleanApiResponse UserLocationDelete(ctx, id, appKey).Execute()

删除位置



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
	id := int64(789) // int64 | 位置ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLocationAPI.UserLocationDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLocationAPI.UserLocationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserLocationAPI.UserLocationDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 位置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserLocationDeleteRequest struct via the builder pattern


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


## UserLocationPost

> UserLocationPostResultApiResponse UserLocationPost(ctx, appKey).GeoLocationModel(geoLocationModel).Execute()

添加位置



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
	geoLocationModel := *openapiclient.NewGeoLocationModel(float64(123), float64(123), "LocationType_example") // GeoLocationModel | 位置信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLocationAPI.UserLocationPost(context.Background(), appKey).GeoLocationModel(geoLocationModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLocationAPI.UserLocationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationPost`: UserLocationPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserLocationAPI.UserLocationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserLocationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **geoLocationModel** | [**GeoLocationModel**](GeoLocationModel.md) | 位置信息 | 

### Return type

[**UserLocationPostResultApiResponse**](UserLocationPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserLocationPut

> BooleanApiResponse UserLocationPut(ctx, id, appKey).GeoLocationModel(geoLocationModel).Execute()

更新位置



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
	id := int64(789) // int64 | 位置ID
	appKey := "appKey_example" // string | 
	geoLocationModel := *openapiclient.NewGeoLocationModel(float64(123), float64(123), "LocationType_example") // GeoLocationModel | 位置信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLocationAPI.UserLocationPut(context.Background(), id, appKey).GeoLocationModel(geoLocationModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLocationAPI.UserLocationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserLocationAPI.UserLocationPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 位置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserLocationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **geoLocationModel** | [**GeoLocationModel**](GeoLocationModel.md) | 位置信息 | 

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


## UserLocations

> UserLocationsResultApiResponse UserLocations(ctx, appKey).Tag(tag).Type_(type_).X(x).Y(y).Sphere(sphere).Skip(skip).Take(take).Execute()

获取位置列表



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
	tag := "tag_example" // string | 标签 (optional)
	type_ := "type__example" // string | 分类 (optional)
	x := float64(1.2) // float64 | 纬度 (optional)
	y := float64(1.2) // float64 | 经度 (optional)
	sphere := int64(789) // int64 | 附近距离，单位：米 (optional)
	skip := int32(56) // int32 | 跳过的记录数 (optional)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLocationAPI.UserLocations(context.Background(), appKey).Tag(tag).Type_(type_).X(x).Y(y).Sphere(sphere).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLocationAPI.UserLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocations`: UserLocationsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserLocationAPI.UserLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** | 标签 | 
 **type_** | **string** | 分类 | 
 **x** | **float64** | 纬度 | 
 **y** | **float64** | 经度 | 
 **sphere** | **int64** | 附近距离，单位：米 | 
 **skip** | **int32** | 跳过的记录数 | 
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserLocationsResultApiResponse**](UserLocationsResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

