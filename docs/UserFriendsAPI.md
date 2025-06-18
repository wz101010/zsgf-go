# \UserFriendsAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserCommonInterests**](UserFriendsAPI.md#UserCommonInterests) | **Get** /UserFriends/{appKey}/CommonInterests | 推荐相似兴趣用户
[**UserFollowUser**](UserFriendsAPI.md#UserFollowUser) | **Post** /UserFriends/{appKey}/Follower/{userId} | 添加关注
[**UserFollowerPut**](UserFriendsAPI.md#UserFollowerPut) | **Put** /UserFriends/{appKey}/Follower/{id} | 刷新粉丝数据
[**UserFollowers**](UserFriendsAPI.md#UserFollowers) | **Get** /UserFriends/{appKey}/Followers | 获取粉丝列表
[**UserFollowing**](UserFriendsAPI.md#UserFollowing) | **Get** /UserFriends/{appKey}/Following | 获取关注列表 / 判断是否关注
[**UserFriendsNearBy**](UserFriendsAPI.md#UserFriendsNearBy) | **Get** /UserFriends/{appKey}/NearBy | 推荐附近用户
[**UserMutualFollowers**](UserFriendsAPI.md#UserMutualFollowers) | **Get** /UserFriends/{appKey}/MutualFollowers | 推荐共同粉丝用户
[**UserMutualFollowings**](UserFriendsAPI.md#UserMutualFollowings) | **Get** /UserFriends/{appKey}/MutualFollowings | 推荐共同关注用户
[**UserProfileById**](UserFriendsAPI.md#UserProfileById) | **Get** /UserFriends/{appKey}/Profile/{userId} | 获取用户资料
[**UserUnfollowUser**](UserFriendsAPI.md#UserUnfollowUser) | **Delete** /UserFriends/{appKey}/Follower/{userId} | 取消关注



## UserCommonInterests

> UserCommonInterestsResultApiResponse UserCommonInterests(ctx, appKey).Tag(tag).Skip(skip).Take(take).Execute()

推荐相似兴趣用户



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
	tag := "tag_example" // string | 兴趣标签 (optional)
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserCommonInterests(context.Background(), appKey).Tag(tag).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserCommonInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCommonInterests`: UserCommonInterestsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserCommonInterests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCommonInterestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** | 兴趣标签 | 
 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserCommonInterestsResultApiResponse**](UserCommonInterestsResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserFollowUser

> BooleanApiResponse UserFollowUser(ctx, userId, appKey).Execute()

添加关注



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
	userId := int64(789) // int64 | 要关注的用户ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserFollowUser(context.Background(), userId, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserFollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowUser`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserFollowUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | 要关注的用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowUserRequest struct via the builder pattern


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


## UserFollowerPut

> BooleanApiResponse UserFollowerPut(ctx, id, appKey).FollowerPutModel(followerPutModel).Execute()

刷新粉丝数据



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
	id := int64(789) // int64 | 粉丝ID
	appKey := "appKey_example" // string | 
	followerPutModel := *openapiclient.NewFollowerPutModel() // FollowerPutModel | 更新粉丝的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserFollowerPut(context.Background(), id, appKey).FollowerPutModel(followerPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserFollowerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowerPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserFollowerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 粉丝ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **followerPutModel** | [**FollowerPutModel**](FollowerPutModel.md) | 更新粉丝的请求参数 | 

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


## UserFollowers

> UserFollowersResultApiResponse UserFollowers(ctx, appKey).Tag(tag).Status(status).TargetUserId(targetUserId).Skip(skip).Take(take).Execute()

获取粉丝列表



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
	status := "status_example" // string | 状态 (optional)
	targetUserId := int64(789) // int64 | 指定用户的粉丝 (optional) (default to 0)
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserFollowers(context.Background(), appKey).Tag(tag).Status(status).TargetUserId(targetUserId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowers`: UserFollowersResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** | 标签 | 
 **status** | **string** | 状态 | 
 **targetUserId** | **int64** | 指定用户的粉丝 | [default to 0]
 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserFollowersResultApiResponse**](UserFollowersResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserFollowing

> BooleanApiResponse UserFollowing(ctx, appKey).Tag(tag).Status(status).TargetUserId(targetUserId).Skip(skip).Take(take).CheckUserId(checkUserId).OnlyIDs(onlyIDs).Execute()

获取关注列表 / 判断是否关注



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
	tag := "tag_example" // string | 用于过滤关注列表的标签（可选）。 (optional)
	status := "status_example" // string | 用于过滤关注列表的状态（可选）。 (optional)
	targetUserId := int64(789) // int64 | 指定用户的关注记录，如果不提供则默认为当前用户的关注。 (optional) (default to 0)
	skip := int32(56) // int32 | 跳过的记录数，用于分页（默认0）。 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数，用于分页（默认10）。 (optional) (default to 10)
	checkUserId := int64(789) // int64 | 要判断是否关注的目标用户ID。如果提供此参数，方法将返回一个布尔值，表示当前用户是否关注该目标用户。 (optional)
	onlyIDs := true // bool | 是否只返回关注用户的ID集合，默认为false（即返回完整的关注用户信息）。 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserFollowing(context.Background(), appKey).Tag(tag).Status(status).TargetUserId(targetUserId).Skip(skip).Take(take).CheckUserId(checkUserId).OnlyIDs(onlyIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowing`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserFollowing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **string** | 用于过滤关注列表的标签（可选）。 | 
 **status** | **string** | 用于过滤关注列表的状态（可选）。 | 
 **targetUserId** | **int64** | 指定用户的关注记录，如果不提供则默认为当前用户的关注。 | [default to 0]
 **skip** | **int32** | 跳过的记录数，用于分页（默认0）。 | [default to 0]
 **take** | **int32** | 获取的记录数，用于分页（默认10）。 | [default to 10]
 **checkUserId** | **int64** | 要判断是否关注的目标用户ID。如果提供此参数，方法将返回一个布尔值，表示当前用户是否关注该目标用户。 | 
 **onlyIDs** | **bool** | 是否只返回关注用户的ID集合，默认为false（即返回完整的关注用户信息）。 | [default to false]

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


## UserFriendsNearBy

> UserFriendsNearByResultApiResponse UserFriendsNearBy(ctx, appKey).Longitude(longitude).Latitude(latitude).Country(country).State(state).City(city).District(district).Gender(gender).AgeS(ageS).AgeE(ageE).Tag(tag).Distance(distance).Skip(skip).Take(take).Execute()

推荐附近用户



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
	longitude := float64(1.2) // float64 | 当前用户经度坐标(WGS84坐标系)
	latitude := float64(1.2) // float64 | 当前用户纬度坐标(WGS84坐标系)
	appKey := "appKey_example" // string | 
	country := "country_example" // string | 国家过滤条件（精确匹配） (optional)
	state := "state_example" // string | 省份过滤条件（精确匹配） (optional)
	city := "city_example" // string | 城市过滤条件（精确匹配） (optional)
	district := "district_example" // string | 区县过滤条件（精确匹配） (optional)
	gender := "gender_example" // string | 性别过滤条件（可选值示例：Male/Female/Other） (optional)
	ageS := int32(56) // int32 | 年龄起始范围（包含，0表示不限制） (optional)
	ageE := int32(56) // int32 | 年龄结束范围（包含，0表示不限制） (optional)
	tag := "tag_example" // string | 兴趣标签过滤（支持模糊匹配，如：\"运动\"） (optional)
	distance := int64(789) // int64 | 搜索半径（单位：米，0表示不限制距离） (optional) (default to 0)
	skip := int32(56) // int32 | 跳过的记录数（分页起始位置，默认0） (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数（分页大小，默认10，最大100） (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserFriendsNearBy(context.Background(), appKey).Longitude(longitude).Latitude(latitude).Country(country).State(state).City(city).District(district).Gender(gender).AgeS(ageS).AgeE(ageE).Tag(tag).Distance(distance).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserFriendsNearBy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFriendsNearBy`: UserFriendsNearByResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserFriendsNearBy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserFriendsNearByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **longitude** | **float64** | 当前用户经度坐标(WGS84坐标系) | 
 **latitude** | **float64** | 当前用户纬度坐标(WGS84坐标系) | 

 **country** | **string** | 国家过滤条件（精确匹配） | 
 **state** | **string** | 省份过滤条件（精确匹配） | 
 **city** | **string** | 城市过滤条件（精确匹配） | 
 **district** | **string** | 区县过滤条件（精确匹配） | 
 **gender** | **string** | 性别过滤条件（可选值示例：Male/Female/Other） | 
 **ageS** | **int32** | 年龄起始范围（包含，0表示不限制） | 
 **ageE** | **int32** | 年龄结束范围（包含，0表示不限制） | 
 **tag** | **string** | 兴趣标签过滤（支持模糊匹配，如：\&quot;运动\&quot;） | 
 **distance** | **int64** | 搜索半径（单位：米，0表示不限制距离） | [default to 0]
 **skip** | **int32** | 跳过的记录数（分页起始位置，默认0） | [default to 0]
 **take** | **int32** | 获取的记录数（分页大小，默认10，最大100） | [default to 10]

### Return type

[**UserFriendsNearByResultApiResponse**](UserFriendsNearByResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMutualFollowers

> UserMutualFollowersResultApiResponse UserMutualFollowers(ctx, appKey).Skip(skip).Take(take).Execute()

推荐共同粉丝用户



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
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserMutualFollowers(context.Background(), appKey).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserMutualFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMutualFollowers`: UserMutualFollowersResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserMutualFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserMutualFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserMutualFollowersResultApiResponse**](UserMutualFollowersResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserMutualFollowings

> UserMutualFollowingsResultApiResponse UserMutualFollowings(ctx, appKey).Skip(skip).Take(take).Execute()

推荐共同关注用户



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
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserMutualFollowings(context.Background(), appKey).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserMutualFollowings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMutualFollowings`: UserMutualFollowingsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserMutualFollowings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserMutualFollowingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserMutualFollowingsResultApiResponse**](UserMutualFollowingsResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserProfileById

> GetUserProfileResultApiResponse UserProfileById(ctx, userId, appKey).Execute()

获取用户资料



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
	userId := int64(789) // int64 | 用户ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserProfileById(context.Background(), userId, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserProfileById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserProfileById`: GetUserProfileResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | 用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserProfileResultApiResponse**](GetUserProfileResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUnfollowUser

> BooleanApiResponse UserUnfollowUser(ctx, userId, appKey).Execute()

取消关注



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
	userId := int64(789) // int64 | 要取消关注的用户ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFriendsAPI.UserUnfollowUser(context.Background(), userId, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFriendsAPI.UserUnfollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUnfollowUser`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserFriendsAPI.UserUnfollowUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | 要取消关注的用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUnfollowUserRequest struct via the builder pattern


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

