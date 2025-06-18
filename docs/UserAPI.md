# \UserAPI

All URIs are relative to *https://api-dev.zashigaofa.cn*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserDeactivateHard**](UserAPI.md#UserDeactivateHard) | **Delete** /User/{appKey}/DeactivateHard | 注销账号
[**UserEmailSignIn**](UserAPI.md#UserEmailSignIn) | **Post** /User/{appKey}/EmailSignIn | 邮箱登录
[**UserEmailSignUp**](UserAPI.md#UserEmailSignUp) | **Post** /User/{appKey}/EmailSignUp | 邮箱注册
[**UserPhoneSignIn**](UserAPI.md#UserPhoneSignIn) | **Post** /User/{appKey}/PhoneSignIn | 手机登录
[**UserPhoneSignUp**](UserAPI.md#UserPhoneSignUp) | **Post** /User/{appKey}/PhoneSignUp | 手机注册
[**UserProfile**](UserAPI.md#UserProfile) | **Get** /User/{appKey}/Profile | 获取个人资料
[**UserResetEmail**](UserAPI.md#UserResetEmail) | **Put** /User/{appKey}/ResetEmail | 重置邮箱
[**UserResetPhone**](UserAPI.md#UserResetPhone) | **Put** /User/{appKey}/ResetPhone | 重置手机号
[**UserResetPwd**](UserAPI.md#UserResetPwd) | **Post** /User/{appKey}/ResetPwd | 重置密码
[**UserSendEmailCode**](UserAPI.md#UserSendEmailCode) | **Post** /User/{appKey}/SendEmailCode | 发送邮箱验证码
[**UserSendSMSCode**](UserAPI.md#UserSendSMSCode) | **Post** /User/{appKey}/SendSMSCode | 发送手机验证码
[**UserSignIn**](UserAPI.md#UserSignIn) | **Post** /User/{appKey}/SignIn | 密码登录
[**UserSignUp**](UserAPI.md#UserSignUp) | **Post** /User/{appKey}/SignUp | 账号注册
[**UserTwoFactorAuth**](UserAPI.md#UserTwoFactorAuth) | **Get** /User/{appKey}/TwoFactorAuth | 二次验证
[**UserUnionIDSignIn**](UserAPI.md#UserUnionIDSignIn) | **Post** /User/{appKey}/UnionIDSignIn | UnionID登录
[**UserUnionIDSignUp**](UserAPI.md#UserUnionIDSignUp) | **Post** /User/{appKey}/UnionIDSignUp | UnionID注册
[**UserUpdateProfile**](UserAPI.md#UserUpdateProfile) | **Put** /User/{appKey}/Profile | 更新个人资料



## UserDeactivateHard

> BooleanApiResponse UserDeactivateHard(ctx, appKey).Execute()

注销账号



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
	resp, r, err := apiClient.UserAPI.UserDeactivateHard(context.Background(), appKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserDeactivateHard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserDeactivateHard`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserDeactivateHard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeactivateHardRequest struct via the builder pattern


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


## UserResetEmail

> BooleanApiResponse UserResetEmail(ctx, appKey).AppUserResetEmailRequest(appUserResetEmailRequest).Execute()

重置邮箱



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
	appUserResetEmailRequest := *openapiclient.NewAppUserResetEmailRequest("Code_example") // AppUserResetEmailRequest | 重置邮箱的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserResetEmail(context.Background(), appKey).AppUserResetEmailRequest(appUserResetEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserResetEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserResetEmail`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserResetEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserResetEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserResetEmailRequest** | [**AppUserResetEmailRequest**](AppUserResetEmailRequest.md) | 重置邮箱的请求参数 | 

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


## UserResetPhone

> BooleanApiResponse UserResetPhone(ctx, appKey).AppUserResetPhoneRequest(appUserResetPhoneRequest).Execute()

重置手机号



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
	appUserResetPhoneRequest := *openapiclient.NewAppUserResetPhoneRequest("Code_example") // AppUserResetPhoneRequest | 重置手机号的请求参数 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserResetPhone(context.Background(), appKey).AppUserResetPhoneRequest(appUserResetPhoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserResetPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserResetPhone`: BooleanApiResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserResetPhone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserResetPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserResetPhoneRequest** | [**AppUserResetPhoneRequest**](AppUserResetPhoneRequest.md) | 重置手机号的请求参数 | 

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
	appUserResetPwdRequest := *openapiclient.NewAppUserResetPwdRequest("Pwd_example", "Oldpwd_example") // AppUserResetPwdRequest | 重置密码的请求参数 (optional)

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

密码登录



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

账号注册



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

二次验证



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

