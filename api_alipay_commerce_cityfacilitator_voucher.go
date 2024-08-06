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


// AlipayCommerceCityfacilitatorVoucherAPIService AlipayCommerceCityfacilitatorVoucherAPI service
type AlipayCommerceCityfacilitatorVoucherAPIService service

type ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest struct {
	ctx context.Context
	ApiService *AlipayCommerceCityfacilitatorVoucherAPIService
	alipayCommerceCityfacilitatorVoucherBatchqueryModel *AlipayCommerceCityfacilitatorVoucherBatchqueryModel
}

func (r ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest) AlipayCommerceCityfacilitatorVoucherBatchqueryModel(alipayCommerceCityfacilitatorVoucherBatchqueryModel AlipayCommerceCityfacilitatorVoucherBatchqueryModel) ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest {
	r.alipayCommerceCityfacilitatorVoucherBatchqueryModel = &alipayCommerceCityfacilitatorVoucherBatchqueryModel
	return r
}

func (r ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest) Execute() (*AlipayCommerceCityfacilitatorVoucherBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceCityfacilitatorVoucherBatchqueryExecute(r)
}

/*
AlipayCommerceCityfacilitatorVoucherBatchquery 地铁购票订单批量查询

商户APP调该接口批量查询订单状态

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest
*/
func (r *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherBatchquery(ctx context.Context) ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest {
	return ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayCommerceCityfacilitatorVoucherBatchqueryResponseModel
func (a *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherBatchqueryExecute(r ApiAlipayCommerceCityfacilitatorVoucherBatchqueryRequest) (*AlipayCommerceCityfacilitatorVoucherBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayCommerceCityfacilitatorVoucherBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceCityfacilitatorVoucherAPIService.AlipayCommerceCityfacilitatorVoucherBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/cityfacilitator/voucher/batchquery"

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
	localVarPostBody = r.alipayCommerceCityfacilitatorVoucherBatchqueryModel

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
			var v AlipayCommerceCityfacilitatorVoucherBatchqueryDefaultResponse
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


type ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest struct {
	ctx context.Context
	ApiService *AlipayCommerceCityfacilitatorVoucherAPIService
	alipayCommerceCityfacilitatorVoucherGenerateModel *AlipayCommerceCityfacilitatorVoucherGenerateModel
}

func (r ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest) AlipayCommerceCityfacilitatorVoucherGenerateModel(alipayCommerceCityfacilitatorVoucherGenerateModel AlipayCommerceCityfacilitatorVoucherGenerateModel) ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest {
	r.alipayCommerceCityfacilitatorVoucherGenerateModel = &alipayCommerceCityfacilitatorVoucherGenerateModel
	return r
}

func (r ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest) Execute() (*AlipayCommerceCityfacilitatorVoucherGenerateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceCityfacilitatorVoucherGenerateExecute(r)
}

/*
AlipayCommerceCityfacilitatorVoucherGenerate 地铁购票核销码发码

商户APP调快捷支付付款后，获取地铁购票的核销码

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest
*/
func (r *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherGenerate(ctx context.Context) ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest {
	return ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayCommerceCityfacilitatorVoucherGenerateResponseModel
func (a *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherGenerateExecute(r ApiAlipayCommerceCityfacilitatorVoucherGenerateRequest) (*AlipayCommerceCityfacilitatorVoucherGenerateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayCommerceCityfacilitatorVoucherGenerateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceCityfacilitatorVoucherAPIService.AlipayCommerceCityfacilitatorVoucherGenerate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/cityfacilitator/voucher/generate"

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
	localVarPostBody = r.alipayCommerceCityfacilitatorVoucherGenerateModel

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
			var v AlipayCommerceCityfacilitatorVoucherGenerateDefaultResponse
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


type ApiAlipayCommerceCityfacilitatorVoucherRefundRequest struct {
	ctx context.Context
	ApiService *AlipayCommerceCityfacilitatorVoucherAPIService
	alipayCommerceCityfacilitatorVoucherRefundModel *AlipayCommerceCityfacilitatorVoucherRefundModel
}

func (r ApiAlipayCommerceCityfacilitatorVoucherRefundRequest) AlipayCommerceCityfacilitatorVoucherRefundModel(alipayCommerceCityfacilitatorVoucherRefundModel AlipayCommerceCityfacilitatorVoucherRefundModel) ApiAlipayCommerceCityfacilitatorVoucherRefundRequest {
	r.alipayCommerceCityfacilitatorVoucherRefundModel = &alipayCommerceCityfacilitatorVoucherRefundModel
	return r
}

func (r ApiAlipayCommerceCityfacilitatorVoucherRefundRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayCommerceCityfacilitatorVoucherRefundExecute(r)
}

/*
AlipayCommerceCityfacilitatorVoucherRefund 地铁购票发码退款

地铁购票，商户发码后，调该接口从中间户退款

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayCommerceCityfacilitatorVoucherRefundRequest
*/
func (r *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherRefund(ctx context.Context) ApiAlipayCommerceCityfacilitatorVoucherRefundRequest {
	return ApiAlipayCommerceCityfacilitatorVoucherRefundRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayCommerceCityfacilitatorVoucherAPIService) AlipayCommerceCityfacilitatorVoucherRefundExecute(r ApiAlipayCommerceCityfacilitatorVoucherRefundRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceCityfacilitatorVoucherAPIService.AlipayCommerceCityfacilitatorVoucherRefund")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/cityfacilitator/voucher/refund"

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
	localVarPostBody = r.alipayCommerceCityfacilitatorVoucherRefundModel

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
			var v AlipayCommerceCityfacilitatorVoucherRefundDefaultResponse
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


func (a *AlipayCommerceCityfacilitatorVoucherAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceCityfacilitatorVoucherAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


