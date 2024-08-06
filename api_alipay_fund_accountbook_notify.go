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

// AlipayFundAccountbookNotifyAPIService AlipayFundAccountbookNotifyAPI service
type AlipayFundAccountbookNotifyAPIService service

type ApiAlipayFundAccountbookNotifyQueryRequest struct {
	ctx           context.Context
	ApiService    *AlipayFundAccountbookNotifyAPIService
	productCode   *string
	bizScene      *string
	accountBookId *string
	agreementNo   *string
}

// 产品码。固定为SATF_FUND_BOOK
func (r ApiAlipayFundAccountbookNotifyQueryRequest) ProductCode(productCode string) ApiAlipayFundAccountbookNotifyQueryRequest {
	r.productCode = &productCode
	return r
}

// 场景码。固定为DEFAULT
func (r ApiAlipayFundAccountbookNotifyQueryRequest) BizScene(bizScene string) ApiAlipayFundAccountbookNotifyQueryRequest {
	r.bizScene = &bizScene
	return r
}

// 记账本ID
func (r ApiAlipayFundAccountbookNotifyQueryRequest) AccountBookId(accountBookId string) ApiAlipayFundAccountbookNotifyQueryRequest {
	r.accountBookId = &accountBookId
	return r
}

// 协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传；
func (r ApiAlipayFundAccountbookNotifyQueryRequest) AgreementNo(agreementNo string) ApiAlipayFundAccountbookNotifyQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

func (r ApiAlipayFundAccountbookNotifyQueryRequest) Execute() (*AlipayFundAccountbookNotifyQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundAccountbookNotifyQueryExecute(r)
}

/*
AlipayFundAccountbookNotifyQuery 记账本通知订阅关系查询接口

ISV通过此接口查询指定记账本的入金和出金通知的订阅关系

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundAccountbookNotifyQueryRequest
*/
func (r *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifyQuery(ctx context.Context) ApiAlipayFundAccountbookNotifyQueryRequest {
	return ApiAlipayFundAccountbookNotifyQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundAccountbookNotifyQueryResponseModel
func (a *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifyQueryExecute(r ApiAlipayFundAccountbookNotifyQueryRequest) (*AlipayFundAccountbookNotifyQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundAccountbookNotifyQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundAccountbookNotifyAPIService.AlipayFundAccountbookNotifyQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/accountbook/notify/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.accountBookId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_book_id", r.accountBookId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
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
		var v AlipayFundAccountbookNotifyQueryDefaultResponse
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

type ApiAlipayFundAccountbookNotifySubscribeRequest struct {
	ctx                                       context.Context
	ApiService                                *AlipayFundAccountbookNotifyAPIService
	alipayFundAccountbookNotifySubscribeModel *AlipayFundAccountbookNotifySubscribeModel
}

func (r ApiAlipayFundAccountbookNotifySubscribeRequest) AlipayFundAccountbookNotifySubscribeModel(alipayFundAccountbookNotifySubscribeModel AlipayFundAccountbookNotifySubscribeModel) ApiAlipayFundAccountbookNotifySubscribeRequest {
	r.alipayFundAccountbookNotifySubscribeModel = &alipayFundAccountbookNotifySubscribeModel
	return r
}

func (r ApiAlipayFundAccountbookNotifySubscribeRequest) Execute() (*AlipayFundAccountbookNotifySubscribeResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundAccountbookNotifySubscribeExecute(r)
}

/*
AlipayFundAccountbookNotifySubscribe 记账本通知订阅接口

ISV通过此接口订阅指定记账本的入金和出金通知。目前只有非OpenAPI方式的充值通知（比如大额来账、贷款入金）

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundAccountbookNotifySubscribeRequest
*/
func (r *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifySubscribe(ctx context.Context) ApiAlipayFundAccountbookNotifySubscribeRequest {
	return ApiAlipayFundAccountbookNotifySubscribeRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundAccountbookNotifySubscribeResponseModel
func (a *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifySubscribeExecute(r ApiAlipayFundAccountbookNotifySubscribeRequest) (*AlipayFundAccountbookNotifySubscribeResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundAccountbookNotifySubscribeResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundAccountbookNotifyAPIService.AlipayFundAccountbookNotifySubscribe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/accountbook/notify/subscribe"

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
	localVarPostBody = r.alipayFundAccountbookNotifySubscribeModel

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
		var v AlipayFundAccountbookNotifySubscribeDefaultResponse
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

type ApiAlipayFundAccountbookNotifyUnsubscribeRequest struct {
	ctx                                         context.Context
	ApiService                                  *AlipayFundAccountbookNotifyAPIService
	alipayFundAccountbookNotifyUnsubscribeModel *AlipayFundAccountbookNotifyUnsubscribeModel
}

func (r ApiAlipayFundAccountbookNotifyUnsubscribeRequest) AlipayFundAccountbookNotifyUnsubscribeModel(alipayFundAccountbookNotifyUnsubscribeModel AlipayFundAccountbookNotifyUnsubscribeModel) ApiAlipayFundAccountbookNotifyUnsubscribeRequest {
	r.alipayFundAccountbookNotifyUnsubscribeModel = &alipayFundAccountbookNotifyUnsubscribeModel
	return r
}

func (r ApiAlipayFundAccountbookNotifyUnsubscribeRequest) Execute() (*AlipayFundAccountbookNotifyUnsubscribeResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundAccountbookNotifyUnsubscribeExecute(r)
}

/*
AlipayFundAccountbookNotifyUnsubscribe 记账本通知取消订阅接口

ISV通过此接口取消订阅指定记账本的入金和出金通知

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundAccountbookNotifyUnsubscribeRequest
*/
func (r *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifyUnsubscribe(ctx context.Context) ApiAlipayFundAccountbookNotifyUnsubscribeRequest {
	return ApiAlipayFundAccountbookNotifyUnsubscribeRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundAccountbookNotifyUnsubscribeResponseModel
func (a *AlipayFundAccountbookNotifyAPIService) AlipayFundAccountbookNotifyUnsubscribeExecute(r ApiAlipayFundAccountbookNotifyUnsubscribeRequest) (*AlipayFundAccountbookNotifyUnsubscribeResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundAccountbookNotifyUnsubscribeResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundAccountbookNotifyAPIService.AlipayFundAccountbookNotifyUnsubscribe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/accountbook/notify/unsubscribe"

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
	localVarPostBody = r.alipayFundAccountbookNotifyUnsubscribeModel

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
		var v AlipayFundAccountbookNotifyUnsubscribeDefaultResponse
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

func (a *AlipayFundAccountbookNotifyAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundAccountbookNotifyAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
