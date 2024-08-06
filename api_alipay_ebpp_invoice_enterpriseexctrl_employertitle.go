/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

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

// AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPI service
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService service

type ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest struct {
	ctx                                                           context.Context
	ApiService                                                    *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService
	alipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel *AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest) AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel(alipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest {
	r.alipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel = &alipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel
	return r
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest) Execute() (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryExecute(r)
}

/*
AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchquery 批量查询企业抬头

分页查询企业抬头列表

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest
*/
func (r *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchquery(ctx context.Context) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest {
	return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryResponseModel
func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryExecute(r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryRequest) (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery"

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
	localVarPostBody = r.alipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryModel

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
		var v AlipayEbppInvoiceEnterpriseexctrlEmployertitleBatchqueryDefaultResponse
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

type ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest struct {
	ctx                                                       context.Context
	ApiService                                                *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService
	alipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest) AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel(alipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest {
	r.alipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel = &alipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel
	return r
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest) Execute() (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateExecute(r)
}

/*
AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreate 新增企业抬头

新增一条企业开票抬头

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest
*/
func (r *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreate(ctx context.Context) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest {
	return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel
func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateExecute(r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateRequest) (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create"

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
	localVarPostBody = r.alipayEbppInvoiceEnterpriseexctrlEmployertitleCreateModel

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
		var v AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse
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

type ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest struct {
	ctx                                                       context.Context
	ApiService                                                *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService
	alipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest) AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel(alipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest {
	r.alipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel = &alipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel
	return r
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest) Execute() (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyExecute(r)
}

/*
AlipayEbppInvoiceEnterpriseexctrlEmployertitleModify 修改企业抬头

修改企业开票抬头

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest
*/
func (r *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleModify(ctx context.Context) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest {
	return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel
func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyExecute(r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyRequest) (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle"

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
	localVarPostBody = r.alipayEbppInvoiceEnterpriseexctrlEmployertitleModifyModel

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
		var v AlipayEbppInvoiceEnterpriseexctrlEmployertitleModifyDefaultResponse
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

type ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService
	accountId    *string
	enterpriseId *string
	agreementNo  *string
	titleId      *string
}

// 企业共同账户id
func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) AccountId(accountId string) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest {
	r.accountId = &accountId
	return r
}

// 企业id
func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) EnterpriseId(enterpriseId string) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest {
	r.enterpriseId = &enterpriseId
	return r
}

// 授权签约协议号
func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) AgreementNo(agreementNo string) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 抬头ID
func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) TitleId(titleId string) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest {
	r.titleId = &titleId
	return r
}

func (r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) Execute() (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryExecute(r)
}

/*
AlipayEbppInvoiceEnterpriseexctrlEmployertitleQuery 查询企业抬头

根据抬头ID查询企业抬头详情

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest
*/
func (r *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleQuery(ctx context.Context) ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest {
	return ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryResponseModel
func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryExecute(r ApiAlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryRequest) (*AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService.AlipayEbppInvoiceEnterpriseexctrlEmployertitleQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.enterpriseId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enterprise_id", r.enterpriseId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.titleId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title_id", r.titleId, "form", "")
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
		var v AlipayEbppInvoiceEnterpriseexctrlEmployertitleQueryDefaultResponse
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

func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceEnterpriseexctrlEmployertitleAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
