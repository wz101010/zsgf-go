# \AppSettingAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppServiceSettingGroup**](AppSettingAPI.md#AppServiceSettingGroup) | **Get** /AppSetting/{appKey}/Groups/{id} | 获取服务分组详情
[**AppServiceSettingGroupDelete**](AppSettingAPI.md#AppServiceSettingGroupDelete) | **Delete** /AppSetting/{appKey}/Groups/{id} | 删除服务分组
[**AppServiceSettingGroupPost**](AppSettingAPI.md#AppServiceSettingGroupPost) | **Post** /AppSetting/{appKey}/Groups | 添加服务分组
[**AppServiceSettingGroupPut**](AppSettingAPI.md#AppServiceSettingGroupPut) | **Put** /AppSetting/{appKey}/Groups/{id} | 更新服务分组
[**AppServiceSettingGroups**](AppSettingAPI.md#AppServiceSettingGroups) | **Get** /AppSetting/{appKey}/Groups | 获取服务分组列表
[**AppServiceSettingItem**](AppSettingAPI.md#AppServiceSettingItem) | **Get** /AppSetting/{appKey}/Items/{id} | 服务配置项详情
[**AppServiceSettingItemDelete**](AppSettingAPI.md#AppServiceSettingItemDelete) | **Delete** /AppSetting/{appKey}/Items/{id} | 删除服务配置项
[**AppServiceSettingItemPost**](AppSettingAPI.md#AppServiceSettingItemPost) | **Post** /AppSetting/{appKey}/Items | 添加服务配置项
[**AppServiceSettingItemPut**](AppSettingAPI.md#AppServiceSettingItemPut) | **Put** /AppSetting/{appKey}/Items/{id} | 更新服务配置项
[**AppServiceSettingItems**](AppSettingAPI.md#AppServiceSettingItems) | **Get** /AppSetting/{appKey}/Items | 服务配置项列表
[**AppServiceSettingProvider**](AppSettingAPI.md#AppServiceSettingProvider) | **Get** /AppSetting/{appKey}/Providers/{id} | 获取服务商详情
[**AppServiceSettingProviderDelete**](AppSettingAPI.md#AppServiceSettingProviderDelete) | **Delete** /AppSetting/{appKey}/Providers/{id} | 删除服务商
[**AppServiceSettingProviderPost**](AppSettingAPI.md#AppServiceSettingProviderPost) | **Post** /AppSetting/{appKey}/Providers | 添加服务商
[**AppServiceSettingProviderPut**](AppSettingAPI.md#AppServiceSettingProviderPut) | **Put** /AppSetting/{appKey}/Providers/{id} | 更新服务商
[**AppServiceSettingProviders**](AppSettingAPI.md#AppServiceSettingProviders) | **Get** /AppSetting/{appKey}/Providers | 获取服务商列表
[**AppSetting**](AppSettingAPI.md#AppSetting) | **Get** /AppSetting/{appKey}/{id} | 配置详情
[**AppSettingDelete**](AppSettingAPI.md#AppSettingDelete) | **Delete** /AppSetting/{appKey}/{id} | 删除配置
[**AppSettingPost**](AppSettingAPI.md#AppSettingPost) | **Post** /AppSetting/{appKey} | 增加配置
[**AppSettingPut**](AppSettingAPI.md#AppSettingPut) | **Put** /AppSetting/{appKey}/{id} | 更新配置
[**AppSettings**](AppSettingAPI.md#AppSettings) | **Get** /AppSetting/{appKey} | 配置列表



## AppServiceSettingGroup

> ServiceGroupApiResponse AppServiceSettingGroup(ctx, id, appKey).Execute()

获取服务分组详情



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
	id := int64(789) // int64 | 服务分组ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingGroup(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingGroup`: ServiceGroupApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceGroupApiResponse**](ServiceGroupApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingGroupDelete

> BooleanApiResponse AppServiceSettingGroupDelete(ctx, id, appKey).Execute()

删除服务分组



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
	id := int64(789) // int64 | 服务分组ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingGroupDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingGroupDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingGroupDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingGroupDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingGroupDeleteRequest struct via the builder pattern


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


## AppServiceSettingGroupPost

> AppSettingGroupPostResultApiResponse AppServiceSettingGroupPost(ctx, appKey).ServiceGroup(serviceGroup).Execute()

添加服务分组



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
	serviceGroup := *openapiclient.NewServiceGroup() // ServiceGroup | 服务分组信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingGroupPost(context.Background(), appKey).ServiceGroup(serviceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingGroupPost`: AppSettingGroupPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceGroup** | [**ServiceGroup**](ServiceGroup.md) | 服务分组信息 | 

### Return type

[**AppSettingGroupPostResultApiResponse**](AppSettingGroupPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingGroupPut

> BooleanApiResponse AppServiceSettingGroupPut(ctx, id, appKey).ServiceGroup(serviceGroup).Execute()

更新服务分组



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
	id := int64(789) // int64 | 服务分组ID
	appKey := "appKey_example" // string | 
	serviceGroup := *openapiclient.NewServiceGroup() // ServiceGroup | 服务分组信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingGroupPut(context.Background(), id, appKey).ServiceGroup(serviceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingGroupPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingGroupPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingGroupPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingGroupPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceGroup** | [**ServiceGroup**](ServiceGroup.md) | 服务分组信息 | 

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


## AppServiceSettingGroups

> ServiceGroupListApiResponse AppServiceSettingGroups(ctx, appKey).ProviderId(providerId).ShowFlag(showFlag).Execute()

获取服务分组列表



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
	providerId := int64(789) // int64 | 服务商ID (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingGroups(context.Background(), appKey).ProviderId(providerId).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingGroups`: ServiceGroupListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerId** | **int64** | 服务商ID | 
 **showFlag** | **bool** | 是否显示 | [default to false]

### Return type

[**ServiceGroupListApiResponse**](ServiceGroupListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingItem

> ServiceItemApiResponse AppServiceSettingItem(ctx, id, appKey).Execute()

服务配置项详情



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
	id := int64(789) // int64 | 服务配置项ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingItem(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingItem`: ServiceItemApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceItemApiResponse**](ServiceItemApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingItemDelete

> BooleanApiResponse AppServiceSettingItemDelete(ctx, id, appKey).Execute()

删除服务配置项



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
	id := int64(789) // int64 | 服务配置项ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingItemDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingItemDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingItemDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingItemDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingItemDeleteRequest struct via the builder pattern


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


## AppServiceSettingItemPost

> AppSettingItemPostResultApiResponse AppServiceSettingItemPost(ctx, appKey).ServiceItem(serviceItem).Execute()

添加服务配置项



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
	serviceItem := *openapiclient.NewServiceItem() // ServiceItem | 服务配置项信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingItemPost(context.Background(), appKey).ServiceItem(serviceItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingItemPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingItemPost`: AppSettingItemPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingItemPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingItemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceItem** | [**ServiceItem**](ServiceItem.md) | 服务配置项信息 | 

### Return type

[**AppSettingItemPostResultApiResponse**](AppSettingItemPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingItemPut

> BooleanApiResponse AppServiceSettingItemPut(ctx, id, appKey).ServiceItem(serviceItem).Execute()

更新服务配置项



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
	id := int64(789) // int64 | 服务配置项ID
	appKey := "appKey_example" // string | 
	serviceItem := *openapiclient.NewServiceItem() // ServiceItem | 服务配置项信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingItemPut(context.Background(), id, appKey).ServiceItem(serviceItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingItemPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingItemPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingItemPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置项ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingItemPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceItem** | [**ServiceItem**](ServiceItem.md) | 服务配置项信息 | 

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


## AppServiceSettingItems

> ServiceItemListApiResponse AppServiceSettingItems(ctx, appKey).BizCode(bizCode).ProviderCode(providerCode).GroupCode(groupCode).ShowFlag(showFlag).Execute()

服务配置项列表



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
	bizCode := "bizCode_example" // string | 业务代码 (optional)
	providerCode := "providerCode_example" // string | 服务商代码 (optional)
	groupCode := "groupCode_example" // string | 分组代码 (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingItems(context.Background(), appKey).BizCode(bizCode).ProviderCode(providerCode).GroupCode(groupCode).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingItems`: ServiceItemListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bizCode** | **string** | 业务代码 | 
 **providerCode** | **string** | 服务商代码 | 
 **groupCode** | **string** | 分组代码 | 
 **showFlag** | **bool** | 是否显示 | [default to false]

### Return type

[**ServiceItemListApiResponse**](ServiceItemListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingProvider

> ServiceProviderApiResponse AppServiceSettingProvider(ctx, id, appKey).Execute()

获取服务商详情



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
	id := int64(789) // int64 | 服务商ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingProvider(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingProvider`: ServiceProviderApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceProviderApiResponse**](ServiceProviderApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingProviderDelete

> BooleanApiResponse AppServiceSettingProviderDelete(ctx, id, appKey).Execute()

删除服务商



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
	id := int64(789) // int64 | 服务商ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingProviderDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingProviderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingProviderDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingProviderDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingProviderDeleteRequest struct via the builder pattern


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


## AppServiceSettingProviderPost

> AppSettingProviderPostResultApiResponse AppServiceSettingProviderPost(ctx, appKey).ServiceProvider(serviceProvider).Execute()

添加服务商



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
	serviceProvider := *openapiclient.NewServiceProvider() // ServiceProvider | 服务商信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingProviderPost(context.Background(), appKey).ServiceProvider(serviceProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingProviderPost`: AppSettingProviderPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingProviderPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceProvider** | [**ServiceProvider**](ServiceProvider.md) | 服务商信息 | 

### Return type

[**AppSettingProviderPostResultApiResponse**](AppSettingProviderPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppServiceSettingProviderPut

> BooleanApiResponse AppServiceSettingProviderPut(ctx, id, appKey).ServiceProvider(serviceProvider).Execute()

更新服务商



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
	id := int64(789) // int64 | 服务商ID
	appKey := "appKey_example" // string | 
	serviceProvider := *openapiclient.NewServiceProvider() // ServiceProvider | 服务商信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingProviderPut(context.Background(), id, appKey).ServiceProvider(serviceProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingProviderPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingProviderPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingProviderPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingProviderPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceProvider** | [**ServiceProvider**](ServiceProvider.md) | 服务商信息 | 

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


## AppServiceSettingProviders

> ServiceProviderListApiResponse AppServiceSettingProviders(ctx, appKey).BizCode(bizCode).ShowFlag(showFlag).Execute()

获取服务商列表



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
	bizCode := "bizCode_example" // string | 业务代码 (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppServiceSettingProviders(context.Background(), appKey).BizCode(bizCode).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppServiceSettingProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppServiceSettingProviders`: ServiceProviderListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppServiceSettingProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppServiceSettingProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bizCode** | **string** | 业务代码 | 
 **showFlag** | **bool** | 是否显示 | [default to false]

### Return type

[**ServiceProviderListApiResponse**](ServiceProviderListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSetting

> AppSettingApiResponse AppSetting(ctx, id, appKey).Execute()

配置详情



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
	id := int64(789) // int64 | 配置ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppSetting(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppSetting`: AppSettingApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppSettingApiResponse**](AppSettingApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingDelete

> BooleanApiResponse AppSettingDelete(ctx, id, appKey).Execute()

删除配置



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
	id := int64(789) // int64 | 配置ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppSettingDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppSettingDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppSettingDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppSettingDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingDeleteRequest struct via the builder pattern


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


## AppSettingPost

> AppSettingSettingPostResultApiResponse AppSettingPost(ctx, appKey).AppSetting(appSetting).Execute()

增加配置



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
	appSetting := *openapiclient.NewAppSetting("Code_example", "Value_example") // AppSetting | 配置内容 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppSettingPost(context.Background(), appKey).AppSetting(appSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppSettingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppSettingPost`: AppSettingSettingPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppSettingPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appSetting** | [**AppSetting**](AppSetting.md) | 配置内容 | 

### Return type

[**AppSettingSettingPostResultApiResponse**](AppSettingSettingPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSettingPut

> BooleanApiResponse AppSettingPut(ctx, id, appKey).AppSetting(appSetting).Execute()

更新配置



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
	id := int64(789) // int64 | 配置ID
	appKey := "appKey_example" // string | 
	appSetting := *openapiclient.NewAppSetting("Code_example", "Value_example") // AppSetting | 配置内容 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppSettingPut(context.Background(), id, appKey).AppSetting(appSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppSettingPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppSettingPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppSettingPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appSetting** | [**AppSetting**](AppSetting.md) | 配置内容 | 

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


## AppSettings

> AppSettingListApiResponse AppSettings(ctx, appKey).ProviderCode(providerCode).GroupCode(groupCode).Tag(tag).Code(code).Execute()

配置列表



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
	providerCode := "providerCode_example" // string | 服务商代码 (optional)
	groupCode := "groupCode_example" // string | 分组代码 (optional)
	tag := "tag_example" // string | 标签 (optional)
	code := "code_example" // string | 配置项代码 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppSettingAPI.AppSettings(context.Background(), appKey).ProviderCode(providerCode).GroupCode(groupCode).Tag(tag).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppSettingAPI.AppSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppSettings`: AppSettingListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AppSettingAPI.AppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerCode** | **string** | 服务商代码 | 
 **groupCode** | **string** | 分组代码 | 
 **tag** | **string** | 标签 | 
 **code** | **string** | 配置项代码 | 

### Return type

[**AppSettingListApiResponse**](AppSettingListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

