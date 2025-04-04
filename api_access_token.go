/*
全部  API 文档

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zsgf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AccessTokenAPIService AccessTokenAPI service
type AccessTokenAPIService service

type ApiAccessTokenDeleteRequest struct {
	ctx context.Context
	ApiService *AccessTokenAPIService
	id int64
	appKey string
}

func (r ApiAccessTokenDeleteRequest) Execute() (*BooleanApiResponse, *http.Response, error) {
	return r.ApiService.AccessTokenDeleteExecute(r)
}

/*
AccessTokenDelete 删除令牌

删除用户令牌

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id 令牌ID
 @param appKey
 @return ApiAccessTokenDeleteRequest
*/
func (a *AccessTokenAPIService) AccessTokenDelete(ctx context.Context, id int64, appKey string) ApiAccessTokenDeleteRequest {
	return ApiAccessTokenDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		appKey: appKey,
	}
}

// Execute executes the request
//  @return BooleanApiResponse
func (a *AccessTokenAPIService) AccessTokenDeleteExecute(r ApiAccessTokenDeleteRequest) (*BooleanApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokenAPIService.AccessTokenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/AccessToken/{appKey}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appKey"+"}", url.PathEscape(parameterValueToString(r.appKey, "appKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAccessTokenPostRequest struct {
	ctx context.Context
	ApiService *AccessTokenAPIService
	appKey string
	accessTokenPostRequest *AccessTokenPostRequest
}

// 令牌创建请求参数
func (r ApiAccessTokenPostRequest) AccessTokenPostRequest(accessTokenPostRequest AccessTokenPostRequest) ApiAccessTokenPostRequest {
	r.accessTokenPostRequest = &accessTokenPostRequest
	return r
}

func (r ApiAccessTokenPostRequest) Execute() (*TokenModelApiResponse, *http.Response, error) {
	return r.ApiService.AccessTokenPostExecute(r)
}

/*
AccessTokenPost 创建令牌

创建新的用户令牌

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appKey
 @return ApiAccessTokenPostRequest
*/
func (a *AccessTokenAPIService) AccessTokenPost(ctx context.Context, appKey string) ApiAccessTokenPostRequest {
	return ApiAccessTokenPostRequest{
		ApiService: a,
		ctx: ctx,
		appKey: appKey,
	}
}

// Execute executes the request
//  @return TokenModelApiResponse
func (a *AccessTokenAPIService) AccessTokenPostExecute(r ApiAccessTokenPostRequest) (*TokenModelApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TokenModelApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokenAPIService.AccessTokenPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/AccessToken/{appKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"appKey"+"}", url.PathEscape(parameterValueToString(r.appKey, "appKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accessTokenPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAccessTokenPutRequest struct {
	ctx context.Context
	ApiService *AccessTokenAPIService
	id int64
	appKey string
	accessTokenPutRequest *AccessTokenPutRequest
}

// 令牌更新请求参数
func (r ApiAccessTokenPutRequest) AccessTokenPutRequest(accessTokenPutRequest AccessTokenPutRequest) ApiAccessTokenPutRequest {
	r.accessTokenPutRequest = &accessTokenPutRequest
	return r
}

func (r ApiAccessTokenPutRequest) Execute() (*BooleanApiResponse, *http.Response, error) {
	return r.ApiService.AccessTokenPutExecute(r)
}

/*
AccessTokenPut 更新令牌

更新现有的用户令牌

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id 令牌ID
 @param appKey
 @return ApiAccessTokenPutRequest
*/
func (a *AccessTokenAPIService) AccessTokenPut(ctx context.Context, id int64, appKey string) ApiAccessTokenPutRequest {
	return ApiAccessTokenPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		appKey: appKey,
	}
}

// Execute executes the request
//  @return BooleanApiResponse
func (a *AccessTokenAPIService) AccessTokenPutExecute(r ApiAccessTokenPutRequest) (*BooleanApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokenAPIService.AccessTokenPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/AccessToken/{appKey}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appKey"+"}", url.PathEscape(parameterValueToString(r.appKey, "appKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json", "application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accessTokenPutRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAccessTokensRequest struct {
	ctx context.Context
	ApiService *AccessTokenAPIService
	appKey string
	userId *int64
	skip *int32
	take *int32
}

// 用户ID
func (r ApiAccessTokensRequest) UserId(userId int64) ApiAccessTokensRequest {
	r.userId = &userId
	return r
}

// 跳过的记录数
func (r ApiAccessTokensRequest) Skip(skip int32) ApiAccessTokensRequest {
	r.skip = &skip
	return r
}

// 获取的记录数
func (r ApiAccessTokensRequest) Take(take int32) ApiAccessTokensRequest {
	r.take = &take
	return r
}

func (r ApiAccessTokensRequest) Execute() (*AccessTokenListResultApiResponse, *http.Response, error) {
	return r.ApiService.AccessTokensExecute(r)
}

/*
AccessTokens 令牌列表

获取用户令牌列表

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appKey
 @return ApiAccessTokensRequest
*/
func (a *AccessTokenAPIService) AccessTokens(ctx context.Context, appKey string) ApiAccessTokensRequest {
	return ApiAccessTokensRequest{
		ApiService: a,
		ctx: ctx,
		appKey: appKey,
	}
}

// Execute executes the request
//  @return AccessTokenListResultApiResponse
func (a *AccessTokenAPIService) AccessTokensExecute(r ApiAccessTokensRequest) (*AccessTokenListResultApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccessTokenListResultApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessTokenAPIService.AccessTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/AccessToken/{appKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"appKey"+"}", url.PathEscape(parameterValueToString(r.appKey, "appKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "form", "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
