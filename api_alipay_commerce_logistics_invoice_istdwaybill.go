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

// AlipayCommerceLogisticsInvoiceIstdwaybillAPIService AlipayCommerceLogisticsInvoiceIstdwaybillAPI service
type AlipayCommerceLogisticsInvoiceIstdwaybillAPIService service

type ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest struct {
	ctx                                                  context.Context
	ApiService                                           *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService
	alipayCommerceLogisticsInvoiceIstdwaybillCreateModel *AlipayCommerceLogisticsInvoiceIstdwaybillCreateModel
}

func (r ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest) AlipayCommerceLogisticsInvoiceIstdwaybillCreateModel(alipayCommerceLogisticsInvoiceIstdwaybillCreateModel AlipayCommerceLogisticsInvoiceIstdwaybillCreateModel) ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest {
	r.alipayCommerceLogisticsInvoiceIstdwaybillCreateModel = &alipayCommerceLogisticsInvoiceIstdwaybillCreateModel
	return r
}

func (r ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest) Execute() (*AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsInvoiceIstdwaybillCreateExecute(r)
}

/*
AlipayCommerceLogisticsInvoiceIstdwaybillCreate 开即时配送订单发票

开即时配送订单发票

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest
*/
func (r *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) AlipayCommerceLogisticsInvoiceIstdwaybillCreate(ctx context.Context) ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest {
	return ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel
func (a *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) AlipayCommerceLogisticsInvoiceIstdwaybillCreateExecute(r ApiAlipayCommerceLogisticsInvoiceIstdwaybillCreateRequest) (*AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsInvoiceIstdwaybillAPIService.AlipayCommerceLogisticsInvoiceIstdwaybillCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/invoice/istdwaybill/create"

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
	localVarPostBody = r.alipayCommerceLogisticsInvoiceIstdwaybillCreateModel

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
		var v AlipayCommerceLogisticsInvoiceIstdwaybillCreateDefaultResponse
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

type ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest struct {
	ctx                 context.Context
	ApiService          *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService
	logisticsCode       *string
	outInvoiceRequestNo *string
}

// 即时配送公司编码
func (r ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest) LogisticsCode(logisticsCode string) ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest {
	r.logisticsCode = &logisticsCode
	return r
}

// 开票请求流水号
func (r ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest) OutInvoiceRequestNo(outInvoiceRequestNo string) ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest {
	r.outInvoiceRequestNo = &outInvoiceRequestNo
	return r
}

func (r ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest) Execute() (*AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsInvoiceIstdwaybillQueryExecute(r)
}

/*
AlipayCommerceLogisticsInvoiceIstdwaybillQuery 查询即时配送订单的开票结果

查询即时配送订单的开票结果

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest
*/
func (r *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) AlipayCommerceLogisticsInvoiceIstdwaybillQuery(ctx context.Context) ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest {
	return ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel
func (a *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) AlipayCommerceLogisticsInvoiceIstdwaybillQueryExecute(r ApiAlipayCommerceLogisticsInvoiceIstdwaybillQueryRequest) (*AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsInvoiceIstdwaybillAPIService.AlipayCommerceLogisticsInvoiceIstdwaybillQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/invoice/istdwaybill/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.logisticsCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "logistics_code", r.logisticsCode, "form", "")
	}
	if r.outInvoiceRequestNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_invoice_request_no", r.outInvoiceRequestNo, "form", "")
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
		var v AlipayCommerceLogisticsInvoiceIstdwaybillQueryDefaultResponse
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

func (a *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) signRequest(req *http.Request) error {
	appID := a.client.cfg.AppID
	appCertSN := a.client.cfg.AppCertSN
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

	signature, err := signWithRSA(content, a.client.cfg.privateKey)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ALIPAY-SHA256withRSA %s,sign=%s", authString, signature))
	return nil
}

func (a *AlipayCommerceLogisticsInvoiceIstdwaybillAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
