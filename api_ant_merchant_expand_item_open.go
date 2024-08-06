/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)


// AntMerchantExpandItemOpenAPIService AntMerchantExpandItemOpenAPI service
type AntMerchantExpandItemOpenAPIService service

type ApiAntMerchantExpandItemOpenBatchqueryRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandItemOpenAPIService
	antMerchantExpandItemOpenBatchqueryModel *AntMerchantExpandItemOpenBatchqueryModel
}

func (r ApiAntMerchantExpandItemOpenBatchqueryRequest) AntMerchantExpandItemOpenBatchqueryModel(antMerchantExpandItemOpenBatchqueryModel AntMerchantExpandItemOpenBatchqueryModel) ApiAntMerchantExpandItemOpenBatchqueryRequest {
	r.antMerchantExpandItemOpenBatchqueryModel = &antMerchantExpandItemOpenBatchqueryModel
	return r
}

func (r ApiAntMerchantExpandItemOpenBatchqueryRequest) Execute() (*AntMerchantExpandItemOpenBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandItemOpenBatchqueryExecute(r)
}

/*
AntMerchantExpandItemOpenBatchquery 批量查询商品接口

用于ISV或商户批量查询商品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandItemOpenBatchqueryRequest
*/
func (r *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenBatchquery(ctx context.Context) ApiAntMerchantExpandItemOpenBatchqueryRequest {
	return ApiAntMerchantExpandItemOpenBatchqueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandItemOpenBatchqueryResponseModel
func (a *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenBatchqueryExecute(r ApiAntMerchantExpandItemOpenBatchqueryRequest) (*AntMerchantExpandItemOpenBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandItemOpenBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandItemOpenAPIService.AntMerchantExpandItemOpenBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/item/open/batchquery"

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
	localVarPostBody = r.antMerchantExpandItemOpenBatchqueryModel

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v AntMerchantExpandItemOpenBatchqueryDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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


type ApiAntMerchantExpandItemOpenCreateRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandItemOpenAPIService
	antMerchantExpandItemOpenCreateModel *AntMerchantExpandItemOpenCreateModel
}

func (r ApiAntMerchantExpandItemOpenCreateRequest) AntMerchantExpandItemOpenCreateModel(antMerchantExpandItemOpenCreateModel AntMerchantExpandItemOpenCreateModel) ApiAntMerchantExpandItemOpenCreateRequest {
	r.antMerchantExpandItemOpenCreateModel = &antMerchantExpandItemOpenCreateModel
	return r
}

func (r ApiAntMerchantExpandItemOpenCreateRequest) Execute() (*AntMerchantExpandItemOpenCreateResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandItemOpenCreateExecute(r)
}

/*
AntMerchantExpandItemOpenCreate 创建商品接口

用于ISV或商户创建商品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandItemOpenCreateRequest
*/
func (r *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenCreate(ctx context.Context) ApiAntMerchantExpandItemOpenCreateRequest {
	return ApiAntMerchantExpandItemOpenCreateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandItemOpenCreateResponseModel
func (a *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenCreateExecute(r ApiAntMerchantExpandItemOpenCreateRequest) (*AntMerchantExpandItemOpenCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandItemOpenCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandItemOpenAPIService.AntMerchantExpandItemOpenCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/item/open/create"

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
	localVarPostBody = r.antMerchantExpandItemOpenCreateModel

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v AntMerchantExpandItemOpenCreateDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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


type ApiAntMerchantExpandItemOpenDeleteRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandItemOpenAPIService
	itemId *string
}

// 商品ID
func (r ApiAntMerchantExpandItemOpenDeleteRequest) ItemId(itemId string) ApiAntMerchantExpandItemOpenDeleteRequest {
	r.itemId = &itemId
	return r
}

func (r ApiAntMerchantExpandItemOpenDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AntMerchantExpandItemOpenDeleteExecute(r)
}

/*
AntMerchantExpandItemOpenDelete 删除商品接口

用于ISV或商户删除商品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandItemOpenDeleteRequest
*/
func (r *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenDelete(ctx context.Context) ApiAntMerchantExpandItemOpenDeleteRequest {
	return ApiAntMerchantExpandItemOpenDeleteRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenDeleteExecute(r ApiAntMerchantExpandItemOpenDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandItemOpenAPIService.AntMerchantExpandItemOpenDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/item/open/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.itemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "item_id", r.itemId, "form", "")
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

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v AntMerchantExpandItemOpenDeleteDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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


type ApiAntMerchantExpandItemOpenModifyRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandItemOpenAPIService
	antMerchantExpandItemOpenModifyModel *AntMerchantExpandItemOpenModifyModel
}

func (r ApiAntMerchantExpandItemOpenModifyRequest) AntMerchantExpandItemOpenModifyModel(antMerchantExpandItemOpenModifyModel AntMerchantExpandItemOpenModifyModel) ApiAntMerchantExpandItemOpenModifyRequest {
	r.antMerchantExpandItemOpenModifyModel = &antMerchantExpandItemOpenModifyModel
	return r
}

func (r ApiAntMerchantExpandItemOpenModifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AntMerchantExpandItemOpenModifyExecute(r)
}

/*
AntMerchantExpandItemOpenModify 修改商品接口

用于ISV或商户修改商品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandItemOpenModifyRequest
*/
func (r *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenModify(ctx context.Context) ApiAntMerchantExpandItemOpenModifyRequest {
	return ApiAntMerchantExpandItemOpenModifyRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenModifyExecute(r ApiAntMerchantExpandItemOpenModifyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandItemOpenAPIService.AntMerchantExpandItemOpenModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/item/open/modify"

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
	localVarPostBody = r.antMerchantExpandItemOpenModifyModel

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v AntMerchantExpandItemOpenModifyDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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


type ApiAntMerchantExpandItemOpenQueryRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandItemOpenAPIService
	targetId *string
	scene *string
	targetType *string
	status *string
}

// 商品归属主体ID 例：商品归属主体类型target_type为店铺，则商品归属主体ID为店铺ID（支付宝侧店铺ID）；归属主体类型target_type为小程序，则归属主体ID为小程序ID
func (r ApiAntMerchantExpandItemOpenQueryRequest) TargetId(targetId string) ApiAntMerchantExpandItemOpenQueryRequest {
	r.targetId = &targetId
	return r
}

// 场景码（具体值请参见产品文档）。
func (r ApiAntMerchantExpandItemOpenQueryRequest) Scene(scene string) ApiAntMerchantExpandItemOpenQueryRequest {
	r.scene = &scene
	return r
}

// 商品归属主体类型。枚举如下： 5：店铺。 8：小程序。
func (r ApiAntMerchantExpandItemOpenQueryRequest) TargetType(targetType string) ApiAntMerchantExpandItemOpenQueryRequest {
	r.targetType = &targetType
	return r
}

// 商品状态
func (r ApiAntMerchantExpandItemOpenQueryRequest) Status(status string) ApiAntMerchantExpandItemOpenQueryRequest {
	r.status = &status
	return r
}

func (r ApiAntMerchantExpandItemOpenQueryRequest) Execute() (*AntMerchantExpandItemOpenQueryResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandItemOpenQueryExecute(r)
}

/*
AntMerchantExpandItemOpenQuery 查询商品接口

用于ISV或商户查询商品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandItemOpenQueryRequest
*/
func (r *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenQuery(ctx context.Context) ApiAntMerchantExpandItemOpenQueryRequest {
	return ApiAntMerchantExpandItemOpenQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandItemOpenQueryResponseModel
func (a *AntMerchantExpandItemOpenAPIService) AntMerchantExpandItemOpenQueryExecute(r ApiAntMerchantExpandItemOpenQueryRequest) (*AntMerchantExpandItemOpenQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandItemOpenQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandItemOpenAPIService.AntMerchantExpandItemOpenQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/item/open/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.targetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "target_id", r.targetId, "form", "")
	}
	if r.scene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scene", r.scene, "form", "")
	}
	if r.targetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "target_type", r.targetType, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
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

	// Add signing logic
	err = a.signRequest(req)
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

	// Add verification logic
	err = a.verifyResponse(localVarHTTPResponse, localVarBody)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v AntMerchantExpandItemOpenQueryDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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


func (a *AntMerchantExpandItemOpenAPIService) signRequest(req *http.Request) error {
	appID := a.client.cfg.AppID
	appCertSN := a.client.cfg.AppCertSN
	privateKey := a.client.cfg.PrivateKey

	nonce := generateUUID()
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	authString := fmt.Sprintf("app_id=%s", appID)
	if appCertSN != "" {
		authString += fmt.Sprintf(",app_cert_sn=%s", appCertSN)
	}
	authString += fmt.Sprintf(",nonce=%s,timestamp=%s", nonce, timestamp)

	httpMethod := req.Method
	httpRequestUri := req.URL.Path
	if req.URL.RawQuery != "" {
		httpRequestUri += "?" + req.URL.RawQuery
	}

	var httpRequestBody string
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return err
		}
		httpRequestBody = string(bodyBytes)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	content := authString + "\n" +
		httpMethod + "\n" +
		httpRequestUri + "\n" +
		httpRequestBody + "\n"

	if appAuthToken := req.Header.Get("alipay-app-auth-token"); appAuthToken != "" {
		content += appAuthToken + "\n"
	}

	signature, err := signWithRSA(content, privateKey)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ALIPAY-SHA256withRSA %s,sign=%s", authString, signature))
	return nil
}

func (a *AntMerchantExpandItemOpenAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


