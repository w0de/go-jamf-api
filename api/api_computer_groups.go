/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type ComputerGroupsAPI interface {

	/*
	V1ComputerGroupsGet Returns the list of all computer groups 

	Use it to get the list of all computer groups.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ComputerGroupsAPIV1ComputerGroupsGetRequest
	*/
	V1ComputerGroupsGet(ctx context.Context) ComputerGroupsAPIV1ComputerGroupsGetRequest

	// V1ComputerGroupsGetExecute executes the request
	//  @return []ComputerGroup
	V1ComputerGroupsGetExecute(r ComputerGroupsAPIV1ComputerGroupsGetRequest) ([]ComputerGroup, *http.Response, error)
}

// ComputerGroupsAPIService ComputerGroupsAPI service
type ComputerGroupsAPIService service

type ComputerGroupsAPIV1ComputerGroupsGetRequest struct {
	ctx context.Context
	ApiService ComputerGroupsAPI
}

func (r ComputerGroupsAPIV1ComputerGroupsGetRequest) Execute() ([]ComputerGroup, *http.Response, error) {
	return r.ApiService.V1ComputerGroupsGetExecute(r)
}

/*
V1ComputerGroupsGet Returns the list of all computer groups 

Use it to get the list of all computer groups.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ComputerGroupsAPIV1ComputerGroupsGetRequest
*/
func (a *ComputerGroupsAPIService) V1ComputerGroupsGet(ctx context.Context) ComputerGroupsAPIV1ComputerGroupsGetRequest {
	return ComputerGroupsAPIV1ComputerGroupsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ComputerGroup
func (a *ComputerGroupsAPIService) V1ComputerGroupsGetExecute(r ComputerGroupsAPIV1ComputerGroupsGetRequest) ([]ComputerGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ComputerGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComputerGroupsAPIService.V1ComputerGroupsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/computer-groups"

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