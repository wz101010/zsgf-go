# \UserAPI

All URIs are relative to *https://api.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**User**](UserAPI.md#User) | **Get** /User/{appKey}/{id} | 获取用户详情
[**UserClear**](UserAPI.md#UserClear) | **Delete** /User/{appKey}/Clear | 清空用户数据
[**UserCommonInterests**](UserAPI.md#UserCommonInterests) | **Get** /User/{appKey}/Friends/CommonInterests | 有共同爱好的用户推荐
[**UserDelete**](UserAPI.md#UserDelete) | **Delete** /User/{appKey}/{id} | 删除用户
[**UserEmailSignIn**](UserAPI.md#UserEmailSignIn) | **Post** /User/{appKey}/EmailSignIn | 邮箱登录
[**UserEmailSignUp**](UserAPI.md#UserEmailSignUp) | **Post** /User/{appKey}/EmailSignUp | 邮箱注册
[**UserExport**](UserAPI.md#UserExport) | **Get** /User/{appKey}/Export | 导出用户数据
[**UserFollowUser**](UserAPI.md#UserFollowUser) | **Post** /User/{appKey}/Follower/{userId} | 关注用户
[**UserFollowerPut**](UserAPI.md#UserFollowerPut) | **Put** /User/{appKey}/Follower/{id} | 更新粉丝
[**UserFollowers**](UserAPI.md#UserFollowers) | **Get** /User/{appKey}/Followers | 获取我的粉丝列表
[**UserFollowing**](UserAPI.md#UserFollowing) | **Get** /User/{appKey}/Following | 获取我的关注列表
[**UserFriendsNearBy**](UserAPI.md#UserFriendsNearBy) | **Get** /User/{appKey}/Friends/NearBy | 附近的用户推荐
[**UserImport**](UserAPI.md#UserImport) | **Post** /User/{appKey}/Import | 导入用户数据
[**UserLocation**](UserAPI.md#UserLocation) | **Get** /User/{appKey}/Location/{id} | 获取位置详情
[**UserLocationDelete**](UserAPI.md#UserLocationDelete) | **Delete** /User/{appKey}/Location/{id} | 删除位置
[**UserLocationPost**](UserAPI.md#UserLocationPost) | **Post** /User/{appKey}/Location | 添加位置
[**UserLocationPut**](UserAPI.md#UserLocationPut) | **Put** /User/{appKey}/Location/{id} | 更新位置
[**UserLocations**](UserAPI.md#UserLocations) | **Get** /User/{appKey}/Locations | 获取位置列表
[**UserMutualFollowers**](UserAPI.md#UserMutualFollowers) | **Get** /User/{appKey}/Friends/MutualFollowers | 有共同粉丝的用户推荐
[**UserMutualFollowings**](UserAPI.md#UserMutualFollowings) | **Get** /User/{appKey}/Friends/MutualFollowings | 有共同关注的用户推荐
[**UserOAuthAccountBind**](UserAPI.md#UserOAuthAccountBind) | **Post** /User/{appKey}/OAuthAccountBind | 外部账号 绑定
[**UserOAuthAccountSignIn**](UserAPI.md#UserOAuthAccountSignIn) | **Post** /User/{appKey}/OAuthAccountSignIn | 外部账号 登录
[**UserOAuthAccounts**](UserAPI.md#UserOAuthAccounts) | **Get** /User/{appKey}/OAuthAccounts | 外部账号 绑定列表
[**UserOAuthAccountsPutBind**](UserAPI.md#UserOAuthAccountsPutBind) | **Put** /User/{appKey}/OAuthAccounts/{id} | 外部账号 更新绑定
[**UserOAuthAccountsUnBind**](UserAPI.md#UserOAuthAccountsUnBind) | **Delete** /User/{appKey}/OAuthAccounts/{id} | 外部账号 删除绑定
[**UserPhoneSignIn**](UserAPI.md#UserPhoneSignIn) | **Post** /User/{appKey}/PhoneSignIn | 手机登录
[**UserPhoneSignUp**](UserAPI.md#UserPhoneSignUp) | **Post** /User/{appKey}/PhoneSignUp | 手机注册
[**UserProfile**](UserAPI.md#UserProfile) | **Get** /User/{appKey}/Profile | 获取个人资料
[**UserPut**](UserAPI.md#UserPut) | **Put** /User/{appKey}/{id} | 更新用户信息
[**UserQRCodePreSignIn**](UserAPI.md#UserQRCodePreSignIn) | **Post** /User/{appKey}/QRCodePreSignIn | 微信小程序 - 预登陆
[**UserQRCodeScan**](UserAPI.md#UserQRCodeScan) | **Post** /User/{appKey}/QRCodeScan | 微信小程序 - 扫码
[**UserQRCodeSignIn**](UserAPI.md#UserQRCodeSignIn) | **Post** /User/{appKey}/QRCodeSignIn | 微信小程序 - 网页登录
[**UserQRCodeSignUp**](UserAPI.md#UserQRCodeSignUp) | **Post** /User/{appKey}/QRCodeSignUp | 微信小程序 - 注册
[**UserResetPwd**](UserAPI.md#UserResetPwd) | **Post** /User/{appKey}/ResetPwd | 重置密码
[**UserSendEmailCode**](UserAPI.md#UserSendEmailCode) | **Post** /User/{appKey}/SendEmailCode | 发送邮箱验证码
[**UserSendSMSCode**](UserAPI.md#UserSendSMSCode) | **Post** /User/{appKey}/SendSMSCode | 发送手机验证码
[**UserSignIn**](UserAPI.md#UserSignIn) | **Post** /User/{appKey}/SignIn | 账号密码 登录
[**UserSignUp**](UserAPI.md#UserSignUp) | **Post** /User/{appKey}/SignUp | 账号密码 注册
[**UserTwoFactorAuth**](UserAPI.md#UserTwoFactorAuth) | **Get** /User/{appKey}/TwoFactorAuth | 双因素验证
[**UserUnfollowUser**](UserAPI.md#UserUnfollowUser) | **Delete** /User/{appKey}/Follower/{userId} | 取消关注
[**UserUnionIDSignIn**](UserAPI.md#UserUnionIDSignIn) | **Post** /User/{appKey}/UnionIDSignIn | UnionID登录
[**UserUnionIDSignUp**](UserAPI.md#UserUnionIDSignUp) | **Post** /User/{appKey}/UnionIDSignUp | UnionID注册
[**UserUpdateProfile**](UserAPI.md#UserUpdateProfile) | **Put** /User/{appKey}/Profile | 更新个人资料
[**Users**](UserAPI.md#Users) | **Get** /User/{appKey} | 获取用户列表



## User

> UserApiResponse User(ctx, id, appKey).Execute()

获取用户详情



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
	resp, r, err := apiClient.UserAPI.User(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.User``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `User`: UserApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.User`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserApiResponse**](UserApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserClear

> BooleanApiResponse UserClear(ctx, appKey).Execute()

清空用户数据



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
	resp, r, err := apiClient.UserAPI.UserClear(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserClear`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserClear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserClearRequest struct via the builder pattern


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


## UserCommonInterests

> UserCommonInterestsResultApiResponse UserCommonInterests(ctx, appKey).Tag(tag).Skip(skip).Take(take).Execute()

有共同爱好的用户推荐



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
	resp, r, err := apiClient.UserAPI.UserCommonInterests(context.Background(), appKey).Tag(tag).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserCommonInterests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCommonInterests`: UserCommonInterestsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserCommonInterests`: %v\n", resp)
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


## UserDelete

> BooleanApiResponse UserDelete(ctx, id, appKey).Execute()

删除用户



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
	resp, r, err := apiClient.UserAPI.UserDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteRequest struct via the builder pattern


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


## UserEmailSignIn

> TokenModelApiResponse UserEmailSignIn(ctx, appKey).EmailSignInRequest(emailSignInRequest).Execute()

邮箱登录



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
	emailSignInRequest := *openapiclient.NewEmailSignInRequest("Email_example", "VerifyCode_example") // EmailSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserEmailSignIn(context.Background(), appKey).EmailSignInRequest(emailSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserEmailSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserEmailSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserEmailSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEmailSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailSignInRequest** | [**EmailSignInRequest**](EmailSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserEmailSignUp

> TokenModelApiResponse UserEmailSignUp(ctx, appKey).EmailSignUpRequest(emailSignUpRequest).Execute()

邮箱注册



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
	emailSignUpRequest := *openapiclient.NewEmailSignUpRequest("Email_example", "Pwd_example") // EmailSignUpRequest | 注册请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserEmailSignUp(context.Background(), appKey).EmailSignUpRequest(emailSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserEmailSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserEmailSignUp`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserEmailSignUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserEmailSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailSignUpRequest** | [**EmailSignUpRequest**](EmailSignUpRequest.md) | 注册请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserExport

> *os.File UserExport(ctx, appKey).Execute()

导出用户数据



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
	resp, r, err := apiClient.UserAPI.UserExport(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserFollowUser

> BooleanApiResponse UserFollowUser(ctx, userId, appKey).Execute()

关注用户



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
	resp, r, err := apiClient.UserAPI.UserFollowUser(context.Background(), userId, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserFollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowUser`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserFollowUser`: %v\n", resp)
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

更新粉丝



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
	resp, r, err := apiClient.UserAPI.UserFollowerPut(context.Background(), id, appKey).FollowerPutModel(followerPutModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserFollowerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowerPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserFollowerPut`: %v\n", resp)
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

> UserFollowersResultApiResponse UserFollowers(ctx, appKey).Tag(tag).Status(status).Skip(skip).Take(take).Execute()

获取我的粉丝列表



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
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserFollowers(context.Background(), appKey).Tag(tag).Status(status).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowers`: UserFollowersResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserFollowers`: %v\n", resp)
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

> UserFollowingResultApiResponse UserFollowing(ctx, appKey).Tag(tag).Status(status).Skip(skip).Take(take).Execute()

获取我的关注列表



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
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserFollowing(context.Background(), appKey).Tag(tag).Status(status).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFollowing`: UserFollowingResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserFollowing`: %v\n", resp)
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

 **tag** | **string** | 标签 | 
 **status** | **string** | 状态 | 
 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

### Return type

[**UserFollowingResultApiResponse**](UserFollowingResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserFriendsNearBy

> UserFriendsNearByResultApiResponse UserFriendsNearBy(ctx, appKey).X(x).Y(y).Distance(distance).Gender(gender).AgeS(ageS).AgeE(ageE).Tag(tag).Type_(type_).Skip(skip).Take(take).Execute()

附近的用户推荐



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
	x := float64(1.2) // float64 | 纬度
	y := float64(1.2) // float64 | 经度
	distance := int64(789) // int64 | 附近距离，单位：米
	appKey := "appKey_example" // string | 
	gender := "gender_example" // string | 性别 (optional)
	ageS := int32(56) // int32 | 年龄段起始 (optional)
	ageE := int32(56) // int32 | 年龄段结束 (optional)
	tag := "tag_example" // string | 兴趣标签 (optional)
	type_ := "type__example" // string | 分类 (optional)
	skip := int32(56) // int32 | 跳过的记录数 (optional) (default to 0)
	take := int32(56) // int32 | 获取的记录数 (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserFriendsNearBy(context.Background(), appKey).X(x).Y(y).Distance(distance).Gender(gender).AgeS(ageS).AgeE(ageE).Tag(tag).Type_(type_).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserFriendsNearBy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFriendsNearBy`: UserFriendsNearByResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserFriendsNearBy`: %v\n", resp)
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
 **x** | **float64** | 纬度 | 
 **y** | **float64** | 经度 | 
 **distance** | **int64** | 附近距离，单位：米 | 

 **gender** | **string** | 性别 | 
 **ageS** | **int32** | 年龄段起始 | 
 **ageE** | **int32** | 年龄段结束 | 
 **tag** | **string** | 兴趣标签 | 
 **type_** | **string** | 分类 | 
 **skip** | **int32** | 跳过的记录数 | [default to 0]
 **take** | **int32** | 获取的记录数 | [default to 10]

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


## UserImport

> BooleanApiResponse UserImport(ctx, appKey).Check(check).File(file).Execute()

导入用户数据



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
	check := true // bool | 是否进行检查 (optional)
	file := os.NewFile(1234, "some_file") // *os.File | 导入的文件 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserImport(context.Background(), appKey).Check(check).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserImport`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **check** | **bool** | 是否进行检查 | 
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
	resp, r, err := apiClient.UserAPI.UserLocation(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocation`: GeoLocationModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserLocation`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.UserLocationDelete(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserLocationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationDelete`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserLocationDelete`: %v\n", resp)
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
	geoLocationModel := *openapiclient.NewGeoLocationModel() // GeoLocationModel | 位置信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserLocationPost(context.Background(), appKey).GeoLocationModel(geoLocationModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserLocationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationPost`: UserLocationPostResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserLocationPost`: %v\n", resp)
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
	geoLocationModel := *openapiclient.NewGeoLocationModel() // GeoLocationModel | 位置信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserLocationPut(context.Background(), id, appKey).GeoLocationModel(geoLocationModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserLocationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocationPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserLocationPut`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.UserLocations(context.Background(), appKey).Tag(tag).Type_(type_).X(x).Y(y).Sphere(sphere).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserLocations`: UserLocationsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserLocations`: %v\n", resp)
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


## UserMutualFollowers

> UserMutualFollowersResultApiResponse UserMutualFollowers(ctx, appKey).Skip(skip).Take(take).Execute()

有共同粉丝的用户推荐



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
	resp, r, err := apiClient.UserAPI.UserMutualFollowers(context.Background(), appKey).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMutualFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMutualFollowers`: UserMutualFollowersResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMutualFollowers`: %v\n", resp)
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

有共同关注的用户推荐



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
	resp, r, err := apiClient.UserAPI.UserMutualFollowings(context.Background(), appKey).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserMutualFollowings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserMutualFollowings`: UserMutualFollowingsResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserMutualFollowings`: %v\n", resp)
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


## UserOAuthAccountBind

> BooleanApiResponse UserOAuthAccountBind(ctx, appKey).OAuthAccountBindRequest(oAuthAccountBindRequest).Execute()

外部账号 绑定



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
	oAuthAccountBindRequest := *openapiclient.NewOAuthAccountBindRequest("UnionID_example", "Platform_example", "PlatformName_example") // OAuthAccountBindRequest | 绑定请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserOAuthAccountBind(context.Background(), appKey).OAuthAccountBindRequest(oAuthAccountBindRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserOAuthAccountBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserOAuthAccountBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthAccountBindRequest** | [**OAuthAccountBindRequest**](OAuthAccountBindRequest.md) | 绑定请求参数 | 

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


## UserOAuthAccountSignIn

> TokenModelApiResponse UserOAuthAccountSignIn(ctx, appKey).OAuthAccountSignInRequest(oAuthAccountSignInRequest).Execute()

外部账号 登录



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
	oAuthAccountSignInRequest := *openapiclient.NewOAuthAccountSignInRequest("UnionID_example", "Platform_example") // OAuthAccountSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserOAuthAccountSignIn(context.Background(), appKey).OAuthAccountSignInRequest(oAuthAccountSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserOAuthAccountSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserOAuthAccountSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuthAccountSignInRequest** | [**OAuthAccountSignInRequest**](OAuthAccountSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOAuthAccounts

> UserLoginsListApiResponse UserOAuthAccounts(ctx, appKey).Execute()

外部账号 绑定列表



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
	resp, r, err := apiClient.UserAPI.UserOAuthAccounts(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserOAuthAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccounts`: UserLoginsListApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserOAuthAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserLoginsListApiResponse**](UserLoginsListApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserOAuthAccountsPutBind

> BooleanApiResponse UserOAuthAccountsPutBind(ctx, id, appKey).OAuthAccountPutBindRequest(oAuthAccountPutBindRequest).Execute()

外部账号 更新绑定



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
	id := int64(789) // int64 | 绑定ID
	appKey := "appKey_example" // string | 
	oAuthAccountPutBindRequest := *openapiclient.NewOAuthAccountPutBindRequest() // OAuthAccountPutBindRequest | 更新请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserOAuthAccountsPutBind(context.Background(), id, appKey).OAuthAccountPutBindRequest(oAuthAccountPutBindRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserOAuthAccountsPutBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountsPutBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserOAuthAccountsPutBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 绑定ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsPutBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oAuthAccountPutBindRequest** | [**OAuthAccountPutBindRequest**](OAuthAccountPutBindRequest.md) | 更新请求参数 | 

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


## UserOAuthAccountsUnBind

> BooleanApiResponse UserOAuthAccountsUnBind(ctx, id, appKey).Execute()

外部账号 删除绑定



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
	id := int64(789) // int64 | 绑定ID
	appKey := "appKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserOAuthAccountsUnBind(context.Background(), id, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserOAuthAccountsUnBind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserOAuthAccountsUnBind`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserOAuthAccountsUnBind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 绑定ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserOAuthAccountsUnBindRequest struct via the builder pattern


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


## UserPhoneSignIn

> TokenModelApiResponse UserPhoneSignIn(ctx, appKey).PhoneSignInRequest(phoneSignInRequest).Execute()

手机登录



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
	phoneSignInRequest := *openapiclient.NewPhoneSignInRequest("Phone_example", "VerifyCode_example") // PhoneSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPhoneSignIn(context.Background(), appKey).PhoneSignInRequest(phoneSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPhoneSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPhoneSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPhoneSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserPhoneSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneSignInRequest** | [**PhoneSignInRequest**](PhoneSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPhoneSignUp

> TokenModelApiResponse UserPhoneSignUp(ctx, appKey).PhoneSignUpRequest(phoneSignUpRequest).Execute()

手机注册



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
	phoneSignUpRequest := *openapiclient.NewPhoneSignUpRequest("Phone_example", "PhoneCode_example", "Pwd_example") // PhoneSignUpRequest | 注册请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPhoneSignUp(context.Background(), appKey).PhoneSignUpRequest(phoneSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPhoneSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPhoneSignUp`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPhoneSignUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserPhoneSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneSignUpRequest** | [**PhoneSignUpRequest**](PhoneSignUpRequest.md) | 注册请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserProfile

> UserProfileResultApiResponse UserProfile(ctx, appKey).Execute()

获取个人资料



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
	resp, r, err := apiClient.UserAPI.UserProfile(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserProfile`: UserProfileResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfileResultApiResponse**](UserProfileResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserPut

> BooleanApiResponse UserPut(ctx, id, appKey).User(user).Execute()

更新用户信息



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
	user := *openapiclient.NewUser() // User | 用户信息 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserPut(context.Background(), id, appKey).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPut`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | 用户ID | 
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | [**User**](User.md) | 用户信息 | 

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


## UserQRCodePreSignIn

> Int64ApiResponse UserQRCodePreSignIn(ctx, appKey).QRCodePreSignInRequest(qRCodePreSignInRequest).Execute()

微信小程序 - 预登陆



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
	qRCodePreSignInRequest := *openapiclient.NewQRCodePreSignInRequest() // QRCodePreSignInRequest | 预登陆请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserQRCodePreSignIn(context.Background(), appKey).QRCodePreSignInRequest(qRCodePreSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserQRCodePreSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserQRCodePreSignIn`: Int64ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserQRCodePreSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserQRCodePreSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodePreSignInRequest** | [**QRCodePreSignInRequest**](QRCodePreSignInRequest.md) | 预登陆请求参数 | 

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


## UserQRCodeScan

> UserQRCodeScanResultApiResponse UserQRCodeScan(ctx, appKey).QRCodeScanRequest(qRCodeScanRequest).Execute()

微信小程序 - 扫码



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
	qRCodeScanRequest := *openapiclient.NewQRCodeScanRequest(int64(123)) // QRCodeScanRequest | 扫码请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserQRCodeScan(context.Background(), appKey).QRCodeScanRequest(qRCodeScanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserQRCodeScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserQRCodeScan`: UserQRCodeScanResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserQRCodeScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserQRCodeScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeScanRequest** | [**QRCodeScanRequest**](QRCodeScanRequest.md) | 扫码请求参数 | 

### Return type

[**UserQRCodeScanResultApiResponse**](UserQRCodeScanResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserQRCodeSignIn

> TokenModelApiResponse UserQRCodeSignIn(ctx, appKey).QRCodeSignInRequest(qRCodeSignInRequest).Execute()

微信小程序 - 网页登录



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
	qRCodeSignInRequest := *openapiclient.NewQRCodeSignInRequest(int64(123)) // QRCodeSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserQRCodeSignIn(context.Background(), appKey).QRCodeSignInRequest(qRCodeSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserQRCodeSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserQRCodeSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserQRCodeSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserQRCodeSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeSignInRequest** | [**QRCodeSignInRequest**](QRCodeSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserQRCodeSignUp

> TokenModelApiResponse UserQRCodeSignUp(ctx, appKey).QRCodeSignUpRequest(qRCodeSignUpRequest).Execute()

微信小程序 - 注册



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
	qRCodeSignUpRequest := *openapiclient.NewQRCodeSignUpRequest("UnionID_example") // QRCodeSignUpRequest | 注册请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserQRCodeSignUp(context.Background(), appKey).QRCodeSignUpRequest(qRCodeSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserQRCodeSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserQRCodeSignUp`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserQRCodeSignUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserQRCodeSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **qRCodeSignUpRequest** | [**QRCodeSignUpRequest**](QRCodeSignUpRequest.md) | 注册请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserResetPwd

> BooleanApiResponse UserResetPwd(ctx, appKey).AppUserResetPwdRequest(appUserResetPwdRequest).Execute()

重置密码



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
	appUserResetPwdRequest := *openapiclient.NewAppUserResetPwdRequest("Code_example", "Pwd_example") // AppUserResetPwdRequest | 重置密码的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserResetPwd(context.Background(), appKey).AppUserResetPwdRequest(appUserResetPwdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserResetPwd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserResetPwd`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserResetPwd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserResetPwdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserResetPwdRequest** | [**AppUserResetPwdRequest**](AppUserResetPwdRequest.md) | 重置密码的请求参数 | 

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


## UserSendEmailCode

> BooleanApiResponse UserSendEmailCode(ctx, appKey).SendEmailCodeRequest(sendEmailCodeRequest).Execute()

发送邮箱验证码



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
	sendEmailCodeRequest := *openapiclient.NewSendEmailCodeRequest("Email_example", "Type_example") // SendEmailCodeRequest | 发送邮箱验证码的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSendEmailCode(context.Background(), appKey).SendEmailCodeRequest(sendEmailCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSendEmailCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSendEmailCode`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSendEmailCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSendEmailCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendEmailCodeRequest** | [**SendEmailCodeRequest**](SendEmailCodeRequest.md) | 发送邮箱验证码的请求参数 | 

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


## UserSendSMSCode

> BooleanApiResponse UserSendSMSCode(ctx, appKey).SendSMSCodeRequest(sendSMSCodeRequest).Execute()

发送手机验证码



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
	sendSMSCodeRequest := *openapiclient.NewSendSMSCodeRequest("Phone_example", "Type_example") // SendSMSCodeRequest | 发送手机验证码的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSendSMSCode(context.Background(), appKey).SendSMSCodeRequest(sendSMSCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSendSMSCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSendSMSCode`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSendSMSCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSendSMSCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendSMSCodeRequest** | [**SendSMSCodeRequest**](SendSMSCodeRequest.md) | 发送手机验证码的请求参数 | 

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


## UserSignIn

> TokenModelApiResponse UserSignIn(ctx, appKey).SignInRequest(signInRequest).Execute()

账号密码 登录



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
	signInRequest := *openapiclient.NewSignInRequest("UserName_example", "Pwd_example") // SignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSignIn(context.Background(), appKey).SignInRequest(signInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signInRequest** | [**SignInRequest**](SignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSignUp

> TokenModelApiResponse UserSignUp(ctx, appKey).SignUpRequest(signUpRequest).Execute()

账号密码 注册



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
	signUpRequest := *openapiclient.NewSignUpRequest("UserName_example", "Pwd_example") // SignUpRequest | 注册请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSignUp(context.Background(), appKey).SignUpRequest(signUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSignUp`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSignUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signUpRequest** | [**SignUpRequest**](SignUpRequest.md) | 注册请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTwoFactorAuth

> SetupCodeApiResponse UserTwoFactorAuth(ctx, appKey).Execute()

双因素验证



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
	resp, r, err := apiClient.UserAPI.UserTwoFactorAuth(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserTwoFactorAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTwoFactorAuth`: SetupCodeApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserTwoFactorAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTwoFactorAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SetupCodeApiResponse**](SetupCodeApiResponse.md)

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
	resp, r, err := apiClient.UserAPI.UserUnfollowUser(context.Background(), userId, appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserUnfollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUnfollowUser`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserUnfollowUser`: %v\n", resp)
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


## UserUnionIDSignIn

> TokenModelApiResponse UserUnionIDSignIn(ctx, appKey).UnionIDSignInRequest(unionIDSignInRequest).Execute()

UnionID登录



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
	unionIDSignInRequest := *openapiclient.NewUnionIDSignInRequest("UnionID_example", "Platform_example") // UnionIDSignInRequest | 登录请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserUnionIDSignIn(context.Background(), appKey).UnionIDSignInRequest(unionIDSignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserUnionIDSignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUnionIDSignIn`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserUnionIDSignIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUnionIDSignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unionIDSignInRequest** | [**UnionIDSignInRequest**](UnionIDSignInRequest.md) | 登录请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUnionIDSignUp

> TokenModelApiResponse UserUnionIDSignUp(ctx, appKey).UnionIDSignUpRequest(unionIDSignUpRequest).Execute()

UnionID注册



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
	unionIDSignUpRequest := *openapiclient.NewUnionIDSignUpRequest("UnionID_example", "Platform_example", "Pwd_example") // UnionIDSignUpRequest | 注册请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserUnionIDSignUp(context.Background(), appKey).UnionIDSignUpRequest(unionIDSignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserUnionIDSignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUnionIDSignUp`: TokenModelApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserUnionIDSignUp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUnionIDSignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unionIDSignUpRequest** | [**UnionIDSignUpRequest**](UnionIDSignUpRequest.md) | 注册请求参数 | 

### Return type

[**TokenModelApiResponse**](TokenModelApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdateProfile

> BooleanApiResponse UserUpdateProfile(ctx, appKey).UpdateProfileRequest(updateProfileRequest).Execute()

更新个人资料



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
	updateProfileRequest := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | 更新个人资料的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserUpdateProfile(context.Background(), appKey).UpdateProfileRequest(updateProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUpdateProfile`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserUpdateProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) | 更新个人资料的请求参数 | 

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


## Users

> UserListResultApiResponse Users(ctx, appKey).UserName(userName).Email(email).Phone(phone).Platform(platform).UnionId(unionId).Role(role).Skip(skip).Take(take).Execute()

获取用户列表



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
	userName := "userName_example" // string | 用户名 (optional)
	email := "email_example" // string | 邮箱 (optional)
	phone := "phone_example" // string | 电话 (optional)
	platform := "platform_example" // string | 平台 (optional)
	unionId := "unionId_example" // string | 联合ID (optional)
	role := "role_example" // string | 角色 (optional)
	skip := int32(56) // int32 | 跳过的记录数 (optional)
	take := int32(56) // int32 | 获取的记录数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.Users(context.Background(), appKey).UserName(userName).Email(email).Phone(phone).Platform(platform).UnionId(unionId).Role(role).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Users`: UserListResultApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.Users`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userName** | **string** | 用户名 | 
 **email** | **string** | 邮箱 | 
 **phone** | **string** | 电话 | 
 **platform** | **string** | 平台 | 
 **unionId** | **string** | 联合ID | 
 **role** | **string** | 角色 | 
 **skip** | **int32** | 跳过的记录数 | 
 **take** | **int32** | 获取的记录数 | 

### Return type

[**UserListResultApiResponse**](UserListResultApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

