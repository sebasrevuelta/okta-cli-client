/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AuthorizationServerScopesAPI interface {

	/*
		CreateOAuth2Scope Create a Custom Token Scope

		Creates a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@return ApiCreateOAuth2ScopeRequest
	*/
	CreateOAuth2Scope(ctx context.Context, authServerId string) ApiCreateOAuth2ScopeRequest

	// CreateOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	CreateOAuth2ScopeExecute(r ApiCreateOAuth2ScopeRequest) (*APIResponse, error)

	/*
		DeleteOAuth2Scope Delete a Custom Token Scope

		Deletes a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiDeleteOAuth2ScopeRequest
	*/
	DeleteOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiDeleteOAuth2ScopeRequest

	// DeleteOAuth2ScopeExecute executes the request
	DeleteOAuth2ScopeExecute(r ApiDeleteOAuth2ScopeRequest) (*APIResponse, error)

	/*
		GetOAuth2Scope Retrieve a Custom Token Scope

		Retrieves a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiGetOAuth2ScopeRequest
	*/
	GetOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiGetOAuth2ScopeRequest

	// GetOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	GetOAuth2ScopeExecute(r ApiGetOAuth2ScopeRequest) (*APIResponse, error)

	/*
		ListOAuth2Scopes List all Custom Token Scopes

		Lists all custom token scopes

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@return ApiListOAuth2ScopesRequest
	*/
	ListOAuth2Scopes(ctx context.Context, authServerId string) ApiListOAuth2ScopesRequest

	// ListOAuth2ScopesExecute executes the request
	//  @return []OAuth2Scope
	ListOAuth2ScopesExecute(r ApiListOAuth2ScopesRequest) (*APIResponse, error)

	/*
		ReplaceOAuth2Scope Replace a Custom Token Scope

		Replaces a custom token scope

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authServerId `id` of the Authorization Server
		@param scopeId `id` of Scope
		@return ApiReplaceOAuth2ScopeRequest
	*/
	ReplaceOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiReplaceOAuth2ScopeRequest

	// ReplaceOAuth2ScopeExecute executes the request
	//  @return OAuth2Scope
	ReplaceOAuth2ScopeExecute(r ApiReplaceOAuth2ScopeRequest) (*APIResponse, error)
}

// AuthorizationServerScopesAPIService AuthorizationServerScopesAPI service
type AuthorizationServerScopesAPIService service

type ApiCreateOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	oAuth2Scope  *OAuth2Scope
	data         interface{}
	retryCount   int32
}

func (r ApiCreateOAuth2ScopeRequest) OAuth2Scope(oAuth2Scope OAuth2Scope) ApiCreateOAuth2ScopeRequest {
	r.oAuth2Scope = &oAuth2Scope
	return r
}

func (r ApiCreateOAuth2ScopeRequest) Data(data interface{}) ApiCreateOAuth2ScopeRequest {
	r.data = data
	return r
}

func (r ApiCreateOAuth2ScopeRequest) Execute() (*APIResponse, error) {
	return r.ApiService.CreateOAuth2ScopeExecute(r)
}

/*
CreateOAuth2Scope Create a Custom Token Scope

Creates a custom token scope

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @return ApiCreateOAuth2ScopeRequest
*/

func (a *AuthorizationServerScopesAPIService) CreateOAuth2Scope(ctx context.Context, authServerId string) ApiCreateOAuth2ScopeRequest {
	return ApiCreateOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		retryCount:   0,
	}
}

// Execute executes the request
//  @return OAuth2Scope

func (a *AuthorizationServerScopesAPIService) CreateOAuth2ScopeExecute(r ApiCreateOAuth2ScopeRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.CreateOAuth2Scope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// localVarPostBody = r.oAuth2Scope
	localVarPostBody = r.data
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiDeleteOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	data         interface{}
	retryCount   int32
}

func (r ApiDeleteOAuth2ScopeRequest) Data(data interface{}) ApiDeleteOAuth2ScopeRequest {
	r.data = data
	return r
}

func (r ApiDeleteOAuth2ScopeRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeleteOAuth2ScopeExecute(r)
}

/*
DeleteOAuth2Scope Delete a Custom Token Scope

Deletes a custom token scope

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param scopeId `id` of Scope
 @return ApiDeleteOAuth2ScopeRequest
*/

func (a *AuthorizationServerScopesAPIService) DeleteOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiDeleteOAuth2ScopeRequest {
	return ApiDeleteOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request

func (a *AuthorizationServerScopesAPIService) DeleteOAuth2ScopeExecute(r ApiDeleteOAuth2ScopeRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.DeleteOAuth2Scope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiGetOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	data         interface{}
	retryCount   int32
}

func (r ApiGetOAuth2ScopeRequest) Data(data interface{}) ApiGetOAuth2ScopeRequest {
	r.data = data
	return r
}

func (r ApiGetOAuth2ScopeRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetOAuth2ScopeExecute(r)
}

/*
GetOAuth2Scope Retrieve a Custom Token Scope

Retrieves a custom token scope

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param scopeId `id` of Scope
 @return ApiGetOAuth2ScopeRequest
*/

func (a *AuthorizationServerScopesAPIService) GetOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiGetOAuth2ScopeRequest {
	return ApiGetOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request
//  @return OAuth2Scope

func (a *AuthorizationServerScopesAPIService) GetOAuth2ScopeExecute(r ApiGetOAuth2ScopeRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.GetOAuth2Scope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiListOAuth2ScopesRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	q            *string
	filter       *string
	cursor       *string
	limit        *int32
	data         interface{}
	retryCount   int32
}

func (r ApiListOAuth2ScopesRequest) Q(q string) ApiListOAuth2ScopesRequest {
	r.q = &q
	return r
}

func (r ApiListOAuth2ScopesRequest) Filter(filter string) ApiListOAuth2ScopesRequest {
	r.filter = &filter
	return r
}

func (r ApiListOAuth2ScopesRequest) Cursor(cursor string) ApiListOAuth2ScopesRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListOAuth2ScopesRequest) Limit(limit int32) ApiListOAuth2ScopesRequest {
	r.limit = &limit
	return r
}

func (r ApiListOAuth2ScopesRequest) Data(data interface{}) ApiListOAuth2ScopesRequest {
	r.data = data
	return r
}

func (r ApiListOAuth2ScopesRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ListOAuth2ScopesExecute(r)
}

/*
ListOAuth2Scopes List all Custom Token Scopes

Lists all custom token scopes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @return ApiListOAuth2ScopesRequest
*/

func (a *AuthorizationServerScopesAPIService) ListOAuth2Scopes(ctx context.Context, authServerId string) ApiListOAuth2ScopesRequest {
	return ApiListOAuth2ScopesRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		retryCount:   0,
	}
}

// Execute executes the request
//  @return []OAuth2Scope

func (a *AuthorizationServerScopesAPIService) ListOAuth2ScopesExecute(r ApiListOAuth2ScopesRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.ListOAuth2Scopes")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiReplaceOAuth2ScopeRequest struct {
	ctx          context.Context
	ApiService   AuthorizationServerScopesAPI
	authServerId string
	scopeId      string
	oAuth2Scope  *OAuth2Scope
	data         interface{}
	retryCount   int32
}

func (r ApiReplaceOAuth2ScopeRequest) OAuth2Scope(oAuth2Scope OAuth2Scope) ApiReplaceOAuth2ScopeRequest {
	r.oAuth2Scope = &oAuth2Scope
	return r
}

func (r ApiReplaceOAuth2ScopeRequest) Data(data interface{}) ApiReplaceOAuth2ScopeRequest {
	r.data = data
	return r
}

func (r ApiReplaceOAuth2ScopeRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ReplaceOAuth2ScopeExecute(r)
}

/*
ReplaceOAuth2Scope Replace a Custom Token Scope

Replaces a custom token scope

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authServerId `id` of the Authorization Server
 @param scopeId `id` of Scope
 @return ApiReplaceOAuth2ScopeRequest
*/

func (a *AuthorizationServerScopesAPIService) ReplaceOAuth2Scope(ctx context.Context, authServerId string, scopeId string) ApiReplaceOAuth2ScopeRequest {
	return ApiReplaceOAuth2ScopeRequest{
		ApiService:   a,
		ctx:          ctx,
		authServerId: authServerId,
		scopeId:      scopeId,
		retryCount:   0,
	}
}

// Execute executes the request
//  @return OAuth2Scope

func (a *AuthorizationServerScopesAPIService) ReplaceOAuth2ScopeExecute(r ApiReplaceOAuth2ScopeRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationServerScopesAPIService.ReplaceOAuth2Scope")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/authorizationServers/{authServerId}/scopes/{scopeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"authServerId"+"}", url.PathEscape(parameterToString(r.authServerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scopeId"+"}", url.PathEscape(parameterToString(r.scopeId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// localVarPostBody = r.oAuth2Scope
	localVarPostBody = r.data
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}