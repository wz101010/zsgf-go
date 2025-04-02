# \ServiceSettingAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceSetting**](ServiceSettingAPI.md#ServiceSetting) | **Get** /ServiceSetting/{id} | 获取配置详情
[**ServiceSettingDelete**](ServiceSettingAPI.md#ServiceSettingDelete) | **Delete** /ServiceSetting/{id} | 删除配置
[**ServiceSettingGroup**](ServiceSettingAPI.md#ServiceSettingGroup) | **Get** /ServiceSetting/Groups/{id} | 获取服务分组详情
[**ServiceSettingGroupDelete**](ServiceSettingAPI.md#ServiceSettingGroupDelete) | **Delete** /ServiceSetting/Groups/{id} | 删除服务分组
[**ServiceSettingGroupPost**](ServiceSettingAPI.md#ServiceSettingGroupPost) | **Post** /ServiceSetting/Groups | 添加服务分组
[**ServiceSettingGroupPut**](ServiceSettingAPI.md#ServiceSettingGroupPut) | **Put** /ServiceSetting/Groups/{id} | 更新服务分组
[**ServiceSettingGroups**](ServiceSettingAPI.md#ServiceSettingGroups) | **Get** /ServiceSetting/Groups | 获取服务分组列表
[**ServiceSettingItem**](ServiceSettingAPI.md#ServiceSettingItem) | **Get** /ServiceSetting/Items/{id} | 服务配置详情
[**ServiceSettingItemDelete**](ServiceSettingAPI.md#ServiceSettingItemDelete) | **Delete** /ServiceSetting/Items/{id} | 删除服务配置
[**ServiceSettingItemPost**](ServiceSettingAPI.md#ServiceSettingItemPost) | **Post** /ServiceSetting/Items | 添加服务配置
[**ServiceSettingItemPut**](ServiceSettingAPI.md#ServiceSettingItemPut) | **Put** /ServiceSetting/Items/{id} | 更新服务配置
[**ServiceSettingItems**](ServiceSettingAPI.md#ServiceSettingItems) | **Get** /ServiceSetting/Items | 服务配置列表
[**ServiceSettingPost**](ServiceSettingAPI.md#ServiceSettingPost) | **Post** /ServiceSetting | 增加配置
[**ServiceSettingProvider**](ServiceSettingAPI.md#ServiceSettingProvider) | **Get** /ServiceSetting/Providers/{id} | 获取服务商详情
[**ServiceSettingProviderDelete**](ServiceSettingAPI.md#ServiceSettingProviderDelete) | **Delete** /ServiceSetting/Providers/{id} | 删除服务商
[**ServiceSettingProviderPost**](ServiceSettingAPI.md#ServiceSettingProviderPost) | **Post** /ServiceSetting/Providers | 添加服务商
[**ServiceSettingProviderPut**](ServiceSettingAPI.md#ServiceSettingProviderPut) | **Put** /ServiceSetting/Providers/{id} | 更新服务商
[**ServiceSettingProviders**](ServiceSettingAPI.md#ServiceSettingProviders) | **Get** /ServiceSetting/Providers | 获取服务商列表
[**ServiceSettingPut**](ServiceSettingAPI.md#ServiceSettingPut) | **Put** /ServiceSetting/{id} | 更新配置
[**ServiceSettings**](ServiceSettingAPI.md#ServiceSettings) | **Get** /ServiceSetting | 获取配置列表



## ServiceSetting

> SettingsApiResponse ServiceSetting(ctx, id).Execute()

获取配置详情



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSetting(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSetting`: SettingsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsApiResponse**](SettingsApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceSettingDelete

> BooleanApiResponse ServiceSettingDelete(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingDeleteRequest struct via the builder pattern


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


## ServiceSettingGroup

> ServiceGroupApiResponse ServiceSettingGroup(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingGroup`: ServiceGroupApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingGroupRequest struct via the builder pattern


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


## ServiceSettingGroupDelete

> BooleanApiResponse ServiceSettingGroupDelete(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingGroupDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingGroupDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingGroupDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingGroupDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingGroupDeleteRequest struct via the builder pattern


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


## ServiceSettingGroupPost

> ServiceSettingGroupPostResultApiResponse ServiceSettingGroupPost(ctx).ServiceGroup(serviceGroup).Execute()

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
	serviceGroup := *openapiclient.NewServiceGroup() // ServiceGroup | 服务分组实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingGroupPost(context.Background()).ServiceGroup(serviceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingGroupPost`: ServiceSettingGroupPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingGroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceGroup** | [**ServiceGroup**](ServiceGroup.md) | 服务分组实体 | 

### Return type

[**ServiceSettingGroupPostResultApiResponse**](ServiceSettingGroupPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceSettingGroupPut

> BooleanApiResponse ServiceSettingGroupPut(ctx, id).ServiceGroup(serviceGroup).Execute()

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
	serviceGroup := *openapiclient.NewServiceGroup() // ServiceGroup | 服务分组实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingGroupPut(context.Background(), id).ServiceGroup(serviceGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingGroupPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingGroupPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingGroupPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务分组ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingGroupPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceGroup** | [**ServiceGroup**](ServiceGroup.md) | 服务分组实体 | 

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


## ServiceSettingGroups

> ServiceGroupListApiResponse ServiceSettingGroups(ctx).ProviderId(providerId).ShowFlag(showFlag).Execute()

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
	providerId := int64(789) // int64 | 服务商ID (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingGroups(context.Background()).ProviderId(providerId).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingGroups`: ServiceGroupListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingGroupsRequest struct via the builder pattern


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


## ServiceSettingItem

> ServiceItemApiResponse ServiceSettingItem(ctx, id).Execute()

服务配置详情



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
	id := int64(789) // int64 | 服务配置ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingItem(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingItem`: ServiceItemApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingItemRequest struct via the builder pattern


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


## ServiceSettingItemDelete

> BooleanApiResponse ServiceSettingItemDelete(ctx, id).Execute()

删除服务配置



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
	id := int64(789) // int64 | 服务配置ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingItemDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingItemDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingItemDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingItemDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingItemDeleteRequest struct via the builder pattern


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


## ServiceSettingItemPost

> ServiceSettingItemPostResultApiResponse ServiceSettingItemPost(ctx).ServiceItem(serviceItem).Execute()

添加服务配置



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
	serviceItem := *openapiclient.NewServiceItem() // ServiceItem | 服务配置实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingItemPost(context.Background()).ServiceItem(serviceItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingItemPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingItemPost`: ServiceSettingItemPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingItemPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingItemPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceItem** | [**ServiceItem**](ServiceItem.md) | 服务配置实体 | 

### Return type

[**ServiceSettingItemPostResultApiResponse**](ServiceSettingItemPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceSettingItemPut

> BooleanApiResponse ServiceSettingItemPut(ctx, id).ServiceItem(serviceItem).Execute()

更新服务配置



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
	id := int64(789) // int64 | 服务配置ID
	serviceItem := *openapiclient.NewServiceItem() // ServiceItem | 服务配置实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingItemPut(context.Background(), id).ServiceItem(serviceItem).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingItemPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingItemPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingItemPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingItemPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceItem** | [**ServiceItem**](ServiceItem.md) | 服务配置实体 | 

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


## ServiceSettingItems

> ServiceItemListApiResponse ServiceSettingItems(ctx).BizCode(bizCode).ProviderCode(providerCode).GroupCode(groupCode).ShowFlag(showFlag).Execute()

服务配置列表



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
	bizCode := "bizCode_example" // string | 业务代码 (optional)
	providerCode := "providerCode_example" // string | 服务商代码 (optional)
	groupCode := "groupCode_example" // string | 分组代码 (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingItems(context.Background()).BizCode(bizCode).ProviderCode(providerCode).GroupCode(groupCode).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingItems`: ServiceItemListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingItemsRequest struct via the builder pattern


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


## ServiceSettingPost

> ServiceSettingSettingPostResultApiResponse ServiceSettingPost(ctx).Settings(settings).Execute()

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
	settings := *openapiclient.NewSettings() // Settings | 配置实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingPost(context.Background()).Settings(settings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingPost`: ServiceSettingSettingPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**Settings**](Settings.md) | 配置实体 | 

### Return type

[**ServiceSettingSettingPostResultApiResponse**](ServiceSettingSettingPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceSettingProvider

> ServiceProviderApiResponse ServiceSettingProvider(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingProvider`: ServiceProviderApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingProviderRequest struct via the builder pattern


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


## ServiceSettingProviderDelete

> BooleanApiResponse ServiceSettingProviderDelete(ctx, id).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingProviderDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingProviderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingProviderDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingProviderDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingProviderDeleteRequest struct via the builder pattern


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


## ServiceSettingProviderPost

> ServiceSettingProviderPostResultApiResponse ServiceSettingProviderPost(ctx).ServiceProvider(serviceProvider).Execute()

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
	serviceProvider := *openapiclient.NewServiceProvider() // ServiceProvider | 服务商实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingProviderPost(context.Background()).ServiceProvider(serviceProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingProviderPost`: ServiceSettingProviderPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingProviderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceProvider** | [**ServiceProvider**](ServiceProvider.md) | 服务商实体 | 

### Return type

[**ServiceSettingProviderPostResultApiResponse**](ServiceSettingProviderPostResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceSettingProviderPut

> BooleanApiResponse ServiceSettingProviderPut(ctx, id).ServiceProvider(serviceProvider).Execute()

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
	serviceProvider := *openapiclient.NewServiceProvider() // ServiceProvider | 服务商实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingProviderPut(context.Background(), id).ServiceProvider(serviceProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingProviderPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingProviderPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingProviderPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 服务商ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingProviderPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceProvider** | [**ServiceProvider**](ServiceProvider.md) | 服务商实体 | 

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


## ServiceSettingProviders

> ServiceProviderListApiResponse ServiceSettingProviders(ctx).BizCode(bizCode).ShowFlag(showFlag).Execute()

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
	bizCode := "bizCode_example" // string | 业务代码 (optional)
	showFlag := true // bool | 是否显示 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingProviders(context.Background()).BizCode(bizCode).ShowFlag(showFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingProviders`: ServiceProviderListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingProvidersRequest struct via the builder pattern


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


## ServiceSettingPut

> BooleanApiResponse ServiceSettingPut(ctx, id).Settings(settings).Execute()

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
	settings := *openapiclient.NewSettings() // Settings | 配置实体 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettingPut(context.Background(), id).Settings(settings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettingPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettingPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettingPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 配置ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settings** | [**Settings**](Settings.md) | 配置实体 | 

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


## ServiceSettings

> SettingsListApiResponse ServiceSettings(ctx).BizCode(bizCode).BizId(bizId).ProviderCode(providerCode).GroupCode(groupCode).Tag(tag).Code(code).Execute()

获取配置列表



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
	bizCode := "bizCode_example" // string | 业务代码
	bizId := "bizId_example" // string | 业务标识
	providerCode := "providerCode_example" // string | 服务商代码 (optional)
	groupCode := "groupCode_example" // string | 分组代码 (optional)
	tag := "tag_example" // string | 标签 (optional)
	code := "code_example" // string | 配置项代码 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceSettingAPI.ServiceSettings(context.Background()).BizCode(bizCode).BizId(bizId).ProviderCode(providerCode).GroupCode(groupCode).Tag(tag).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceSettingAPI.ServiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSettings`: SettingsListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ServiceSettingAPI.ServiceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bizCode** | **string** | 业务代码 | 
 **bizId** | **string** | 业务标识 | 
 **providerCode** | **string** | 服务商代码 | 
 **groupCode** | **string** | 分组代码 | 
 **tag** | **string** | 标签 | 
 **code** | **string** | 配置项代码 | 

### Return type

[**SettingsListApiResponse**](SettingsListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

