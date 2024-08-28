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

type ApplicationConnectionsAPI interface {

	/*
		ActivateDefaultProvisioningConnectionForApplication Activate the default Provisioning Connection

		Activates the default Provisioning Connection for an app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appId Application ID
		@return ApiActivateDefaultProvisioningConnectionForApplicationRequest
	*/
	ActivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiActivateDefaultProvisioningConnectionForApplicationRequest

	// ActivateDefaultProvisioningConnectionForApplicationExecute executes the request
	ActivateDefaultProvisioningConnectionForApplicationExecute(r ApiActivateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error)

	/*
		DeactivateDefaultProvisioningConnectionForApplication Deactivate the default Provisioning Connection

		Deactivates the default Provisioning Connection for an app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appId Application ID
		@return ApiDeactivateDefaultProvisioningConnectionForApplicationRequest
	*/
	DeactivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiDeactivateDefaultProvisioningConnectionForApplicationRequest

	// DeactivateDefaultProvisioningConnectionForApplicationExecute executes the request
	DeactivateDefaultProvisioningConnectionForApplicationExecute(r ApiDeactivateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error)

	/*
		GetDefaultProvisioningConnectionForApplication Retrieve the default Provisioning Connection

		Retrieves the default Provisioning Connection for an app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appId Application ID
		@return ApiGetDefaultProvisioningConnectionForApplicationRequest
	*/
	GetDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiGetDefaultProvisioningConnectionForApplicationRequest

	// GetDefaultProvisioningConnectionForApplicationExecute executes the request
	//  @return ProvisioningConnectionResponse
	GetDefaultProvisioningConnectionForApplicationExecute(r ApiGetDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error)

	/*
		UpdateDefaultProvisioningConnectionForApplication Update the default Provisioning Connection

		Updates the default Provisioning Connection for an app

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param appId Application ID
		@return ApiUpdateDefaultProvisioningConnectionForApplicationRequest
	*/
	UpdateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiUpdateDefaultProvisioningConnectionForApplicationRequest

	// UpdateDefaultProvisioningConnectionForApplicationExecute executes the request
	//  @return ProvisioningConnectionResponse
	UpdateDefaultProvisioningConnectionForApplicationExecute(r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error)

	/*
			VerifyProvisioningConnectionForApplication Verify the Provisioning Connection

			Verifies the OAuth 2.0-based connection as part of the OAuth 2.0 consent flow. The validation of the consent flow is the last step of the provisioning setup for an OAuth 2.0-based connection.
		Currently, this operation only supports `office365`,`google`, `zoomus`, and `slack` apps.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param appName
			@param appId Application ID
			@return ApiVerifyProvisioningConnectionForApplicationRequest
	*/
	VerifyProvisioningConnectionForApplication(ctx context.Context, appName string, appId string) ApiVerifyProvisioningConnectionForApplicationRequest

	// VerifyProvisioningConnectionForApplicationExecute executes the request
	VerifyProvisioningConnectionForApplicationExecute(r ApiVerifyProvisioningConnectionForApplicationRequest) (*APIResponse, error)
}

// ApplicationConnectionsAPIService ApplicationConnectionsAPI service
type ApplicationConnectionsAPIService service

type ApiActivateDefaultProvisioningConnectionForApplicationRequest struct {
	ctx        context.Context
	ApiService ApplicationConnectionsAPI
	appId      string
	data       interface{}
	retryCount int32
}

func (r ApiActivateDefaultProvisioningConnectionForApplicationRequest) Data(data interface{}) ApiActivateDefaultProvisioningConnectionForApplicationRequest {
	r.data = data
	return r
}

func (r ApiActivateDefaultProvisioningConnectionForApplicationRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ActivateDefaultProvisioningConnectionForApplicationExecute(r)
}

/*
ActivateDefaultProvisioningConnectionForApplication Activate the default Provisioning Connection

Activates the default Provisioning Connection for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId Application ID
 @return ApiActivateDefaultProvisioningConnectionForApplicationRequest
*/

func (a *ApplicationConnectionsAPIService) ActivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiActivateDefaultProvisioningConnectionForApplicationRequest {
	return ApiActivateDefaultProvisioningConnectionForApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		retryCount: 0,
	}
}

// Execute executes the request

func (a *ApplicationConnectionsAPIService) ActivateDefaultProvisioningConnectionForApplicationExecute(r ApiActivateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConnectionsAPIService.ActivateDefaultProvisioningConnectionForApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appId}/connections/default/lifecycle/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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

type ApiDeactivateDefaultProvisioningConnectionForApplicationRequest struct {
	ctx        context.Context
	ApiService ApplicationConnectionsAPI
	appId      string
	data       interface{}
	retryCount int32
}

func (r ApiDeactivateDefaultProvisioningConnectionForApplicationRequest) Data(data interface{}) ApiDeactivateDefaultProvisioningConnectionForApplicationRequest {
	r.data = data
	return r
}

func (r ApiDeactivateDefaultProvisioningConnectionForApplicationRequest) Execute() (*APIResponse, error) {
	return r.ApiService.DeactivateDefaultProvisioningConnectionForApplicationExecute(r)
}

/*
DeactivateDefaultProvisioningConnectionForApplication Deactivate the default Provisioning Connection

Deactivates the default Provisioning Connection for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId Application ID
 @return ApiDeactivateDefaultProvisioningConnectionForApplicationRequest
*/

func (a *ApplicationConnectionsAPIService) DeactivateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiDeactivateDefaultProvisioningConnectionForApplicationRequest {
	return ApiDeactivateDefaultProvisioningConnectionForApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		retryCount: 0,
	}
}

// Execute executes the request

func (a *ApplicationConnectionsAPIService) DeactivateDefaultProvisioningConnectionForApplicationExecute(r ApiDeactivateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConnectionsAPIService.DeactivateDefaultProvisioningConnectionForApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appId}/connections/default/lifecycle/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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

type ApiGetDefaultProvisioningConnectionForApplicationRequest struct {
	ctx        context.Context
	ApiService ApplicationConnectionsAPI
	appId      string
	data       interface{}
	retryCount int32
}

func (r ApiGetDefaultProvisioningConnectionForApplicationRequest) Data(data interface{}) ApiGetDefaultProvisioningConnectionForApplicationRequest {
	r.data = data
	return r
}

func (r ApiGetDefaultProvisioningConnectionForApplicationRequest) Execute() (*APIResponse, error) {
	return r.ApiService.GetDefaultProvisioningConnectionForApplicationExecute(r)
}

/*
GetDefaultProvisioningConnectionForApplication Retrieve the default Provisioning Connection

Retrieves the default Provisioning Connection for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId Application ID
 @return ApiGetDefaultProvisioningConnectionForApplicationRequest
*/

func (a *ApplicationConnectionsAPIService) GetDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiGetDefaultProvisioningConnectionForApplicationRequest {
	return ApiGetDefaultProvisioningConnectionForApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ProvisioningConnectionResponse

func (a *ApplicationConnectionsAPIService) GetDefaultProvisioningConnectionForApplicationExecute(r ApiGetDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConnectionsAPIService.GetDefaultProvisioningConnectionForApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appId}/connections/default"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

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

type ApiUpdateDefaultProvisioningConnectionForApplicationRequest struct {
	ctx                                                      context.Context
	ApiService                                               ApplicationConnectionsAPI
	appId                                                    string
	updateDefaultProvisioningConnectionForApplicationRequest *UpdateDefaultProvisioningConnectionForApplicationRequest
	activate                                                 *bool
	data                                                     interface{}
	retryCount                                               int32
}

func (r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) UpdateDefaultProvisioningConnectionForApplicationRequest(updateDefaultProvisioningConnectionForApplicationRequest UpdateDefaultProvisioningConnectionForApplicationRequest) ApiUpdateDefaultProvisioningConnectionForApplicationRequest {
	r.updateDefaultProvisioningConnectionForApplicationRequest = &updateDefaultProvisioningConnectionForApplicationRequest
	return r
}

// Activates the Provisioning Connection
func (r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) Activate(activate bool) ApiUpdateDefaultProvisioningConnectionForApplicationRequest {
	r.activate = &activate
	return r
}

func (r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) Data(data interface{}) ApiUpdateDefaultProvisioningConnectionForApplicationRequest {
	r.data = data
	return r
}

func (r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) Execute() (*APIResponse, error) {
	return r.ApiService.UpdateDefaultProvisioningConnectionForApplicationExecute(r)
}

/*
UpdateDefaultProvisioningConnectionForApplication Update the default Provisioning Connection

Updates the default Provisioning Connection for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId Application ID
 @return ApiUpdateDefaultProvisioningConnectionForApplicationRequest
*/

func (a *ApplicationConnectionsAPIService) UpdateDefaultProvisioningConnectionForApplication(ctx context.Context, appId string) ApiUpdateDefaultProvisioningConnectionForApplicationRequest {
	return ApiUpdateDefaultProvisioningConnectionForApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		appId:      appId,
		retryCount: 0,
	}
}

// Execute executes the request
//  @return ProvisioningConnectionResponse

func (a *ApplicationConnectionsAPIService) UpdateDefaultProvisioningConnectionForApplicationExecute(r ApiUpdateDefaultProvisioningConnectionForApplicationRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConnectionsAPIService.UpdateDefaultProvisioningConnectionForApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appId}/connections/default"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.activate != nil {
		localVarQueryParams.Add("activate", parameterToString(*r.activate, ""))
	}
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
	// localVarPostBody = r.updateDefaultProvisioningConnectionForApplicationRequest
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

type ApiVerifyProvisioningConnectionForApplicationRequest struct {
	ctx        context.Context
	ApiService ApplicationConnectionsAPI
	appName    string
	appId      string
	code       *string
	state      *string
	data       interface{}
	retryCount int32
}

func (r ApiVerifyProvisioningConnectionForApplicationRequest) Code(code string) ApiVerifyProvisioningConnectionForApplicationRequest {
	r.code = &code
	return r
}

func (r ApiVerifyProvisioningConnectionForApplicationRequest) State(state string) ApiVerifyProvisioningConnectionForApplicationRequest {
	r.state = &state
	return r
}

func (r ApiVerifyProvisioningConnectionForApplicationRequest) Data(data interface{}) ApiVerifyProvisioningConnectionForApplicationRequest {
	r.data = data
	return r
}

func (r ApiVerifyProvisioningConnectionForApplicationRequest) Execute() (*APIResponse, error) {
	return r.ApiService.VerifyProvisioningConnectionForApplicationExecute(r)
}

/*
VerifyProvisioningConnectionForApplication Verify the Provisioning Connection

Verifies the OAuth 2.0-based connection as part of the OAuth 2.0 consent flow. The validation of the consent flow is the last step of the provisioning setup for an OAuth 2.0-based connection.
Currently, this operation only supports `office365`,`google`, `zoomus`, and `slack` apps.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appName
 @param appId Application ID
 @return ApiVerifyProvisioningConnectionForApplicationRequest
*/

func (a *ApplicationConnectionsAPIService) VerifyProvisioningConnectionForApplication(ctx context.Context, appName string, appId string) ApiVerifyProvisioningConnectionForApplicationRequest {
	return ApiVerifyProvisioningConnectionForApplicationRequest{
		ApiService: a,
		ctx:        ctx,
		appName:    appName,
		appId:      appId,
		retryCount: 0,
	}
}

// Execute executes the request

func (a *ApplicationConnectionsAPIService) VerifyProvisioningConnectionForApplicationExecute(r ApiVerifyProvisioningConnectionForApplicationRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationConnectionsAPIService.VerifyProvisioningConnectionForApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/apps/{appName}/{appId}/oauth2/callback"
	localVarPath = strings.Replace(localVarPath, "{"+"appName"+"}", url.PathEscape(parameterToString(r.appName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterToString(r.appId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.code != nil {
		localVarQueryParams.Add("code", parameterToString(*r.code, ""))
	}
	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, ""))
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
