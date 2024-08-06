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

// AlipayOpenMiniIsvFastregisterAPIService AlipayOpenMiniIsvFastregisterAPI service
type AlipayOpenMiniIsvFastregisterAPIService service

type ApiAlipayOpenMiniIsvFastregisterCreateRequest struct {
	ctx                                      context.Context
	ApiService                               *AlipayOpenMiniIsvFastregisterAPIService
	alipayOpenMiniIsvFastregisterCreateModel *AlipayOpenMiniIsvFastregisterCreateModel
}

func (r ApiAlipayOpenMiniIsvFastregisterCreateRequest) AlipayOpenMiniIsvFastregisterCreateModel(alipayOpenMiniIsvFastregisterCreateModel AlipayOpenMiniIsvFastregisterCreateModel) ApiAlipayOpenMiniIsvFastregisterCreateRequest {
	r.alipayOpenMiniIsvFastregisterCreateModel = &alipayOpenMiniIsvFastregisterCreateModel
	return r
}

func (r ApiAlipayOpenMiniIsvFastregisterCreateRequest) Execute() (*AlipayOpenMiniIsvFastregisterCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenMiniIsvFastregisterCreateExecute(r)
}

/*
AlipayOpenMiniIsvFastregisterCreate ISV代创建试用小程序

ISV代商户创建试用小程序

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenMiniIsvFastregisterCreateRequest
*/
func (r *AlipayOpenMiniIsvFastregisterAPIService) AlipayOpenMiniIsvFastregisterCreate(ctx context.Context) ApiAlipayOpenMiniIsvFastregisterCreateRequest {
	return ApiAlipayOpenMiniIsvFastregisterCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenMiniIsvFastregisterCreateResponseModel
func (a *AlipayOpenMiniIsvFastregisterAPIService) AlipayOpenMiniIsvFastregisterCreateExecute(r ApiAlipayOpenMiniIsvFastregisterCreateRequest) (*AlipayOpenMiniIsvFastregisterCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenMiniIsvFastregisterCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenMiniIsvFastregisterAPIService.AlipayOpenMiniIsvFastregisterCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/mini/isv/fastregister/create"

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
	localVarPostBody = r.alipayOpenMiniIsvFastregisterCreateModel

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
		var v AlipayOpenMiniIsvFastregisterCreateDefaultResponse
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

type ApiAlipayOpenMiniIsvFastregisterQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenMiniIsvFastregisterAPIService
	orderNo    *string
}

// 代创建试用小程序单号
func (r ApiAlipayOpenMiniIsvFastregisterQueryRequest) OrderNo(orderNo string) ApiAlipayOpenMiniIsvFastregisterQueryRequest {
	r.orderNo = &orderNo
	return r
}

func (r ApiAlipayOpenMiniIsvFastregisterQueryRequest) Execute() (*AlipayOpenMiniIsvFastregisterQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenMiniIsvFastregisterQueryExecute(r)
}

/*
AlipayOpenMiniIsvFastregisterQuery 查询试用小程序

查询试用小程序

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenMiniIsvFastregisterQueryRequest
*/
func (r *AlipayOpenMiniIsvFastregisterAPIService) AlipayOpenMiniIsvFastregisterQuery(ctx context.Context) ApiAlipayOpenMiniIsvFastregisterQueryRequest {
	return ApiAlipayOpenMiniIsvFastregisterQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenMiniIsvFastregisterQueryResponseModel
func (a *AlipayOpenMiniIsvFastregisterAPIService) AlipayOpenMiniIsvFastregisterQueryExecute(r ApiAlipayOpenMiniIsvFastregisterQueryRequest) (*AlipayOpenMiniIsvFastregisterQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenMiniIsvFastregisterQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenMiniIsvFastregisterAPIService.AlipayOpenMiniIsvFastregisterQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/mini/isv/fastregister/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.orderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_no", r.orderNo, "form", "")
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
		var v AlipayOpenMiniIsvFastregisterQueryDefaultResponse
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

func (a *AlipayOpenMiniIsvFastregisterAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenMiniIsvFastregisterAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
