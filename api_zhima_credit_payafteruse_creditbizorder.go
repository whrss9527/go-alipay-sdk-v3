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


// ZhimaCreditPayafteruseCreditbizorderAPIService ZhimaCreditPayafteruseCreditbizorderAPI service
type ZhimaCreditPayafteruseCreditbizorderAPIService service

type ApiZhimaCreditPayafteruseCreditbizorderFinishRequest struct {
	ctx context.Context
	ApiService *ZhimaCreditPayafteruseCreditbizorderAPIService
	zhimaCreditPayafteruseCreditbizorderFinishModel *ZhimaCreditPayafteruseCreditbizorderFinishModel
}

func (r ApiZhimaCreditPayafteruseCreditbizorderFinishRequest) ZhimaCreditPayafteruseCreditbizorderFinishModel(zhimaCreditPayafteruseCreditbizorderFinishModel ZhimaCreditPayafteruseCreditbizorderFinishModel) ApiZhimaCreditPayafteruseCreditbizorderFinishRequest {
	r.zhimaCreditPayafteruseCreditbizorderFinishModel = &zhimaCreditPayafteruseCreditbizorderFinishModel
	return r
}

func (r ApiZhimaCreditPayafteruseCreditbizorderFinishRequest) Execute() (*ZhimaCreditPayafteruseCreditbizorderFinishResponseModel, *http.Response, error) {
	return r.ApiService.ZhimaCreditPayafteruseCreditbizorderFinishExecute(r)
}

/*
ZhimaCreditPayafteruseCreditbizorderFinish 结束信用服务订单

结束信用服务订单

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiZhimaCreditPayafteruseCreditbizorderFinishRequest
*/
func (r *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderFinish(ctx context.Context) ApiZhimaCreditPayafteruseCreditbizorderFinishRequest {
	return ApiZhimaCreditPayafteruseCreditbizorderFinishRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ZhimaCreditPayafteruseCreditbizorderFinishResponseModel
func (a *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderFinishExecute(r ApiZhimaCreditPayafteruseCreditbizorderFinishRequest) (*ZhimaCreditPayafteruseCreditbizorderFinishResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ZhimaCreditPayafteruseCreditbizorderFinishResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ZhimaCreditPayafteruseCreditbizorderAPIService.ZhimaCreditPayafteruseCreditbizorderFinish")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/zhima/credit/payafteruse/creditbizorder/finish"

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
	localVarPostBody = r.zhimaCreditPayafteruseCreditbizorderFinishModel

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
			var v ZhimaCreditPayafteruseCreditbizorderFinishDefaultResponse
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


type ApiZhimaCreditPayafteruseCreditbizorderOrderRequest struct {
	ctx context.Context
	ApiService *ZhimaCreditPayafteruseCreditbizorderAPIService
	zhimaCreditPayafteruseCreditbizorderOrderModel *ZhimaCreditPayafteruseCreditbizorderOrderModel
}

func (r ApiZhimaCreditPayafteruseCreditbizorderOrderRequest) ZhimaCreditPayafteruseCreditbizorderOrderModel(zhimaCreditPayafteruseCreditbizorderOrderModel ZhimaCreditPayafteruseCreditbizorderOrderModel) ApiZhimaCreditPayafteruseCreditbizorderOrderRequest {
	r.zhimaCreditPayafteruseCreditbizorderOrderModel = &zhimaCreditPayafteruseCreditbizorderOrderModel
	return r
}

func (r ApiZhimaCreditPayafteruseCreditbizorderOrderRequest) Execute() (*ZhimaCreditPayafteruseCreditbizorderOrderResponseModel, *http.Response, error) {
	return r.ApiService.ZhimaCreditPayafteruseCreditbizorderOrderExecute(r)
}

/*
ZhimaCreditPayafteruseCreditbizorderOrder 芝麻信用服务下单（免用户确认场景）

芝麻信用产品免密下单，不需要唤起支付宝APP，通过服务端调用完成下单。
涉及芝麻信用服务产品、芝麻风险评估产品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiZhimaCreditPayafteruseCreditbizorderOrderRequest
*/
func (r *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderOrder(ctx context.Context) ApiZhimaCreditPayafteruseCreditbizorderOrderRequest {
	return ApiZhimaCreditPayafteruseCreditbizorderOrderRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ZhimaCreditPayafteruseCreditbizorderOrderResponseModel
func (a *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderOrderExecute(r ApiZhimaCreditPayafteruseCreditbizorderOrderRequest) (*ZhimaCreditPayafteruseCreditbizorderOrderResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ZhimaCreditPayafteruseCreditbizorderOrderResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ZhimaCreditPayafteruseCreditbizorderAPIService.ZhimaCreditPayafteruseCreditbizorderOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/zhima/credit/payafteruse/creditbizorder/order"

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
	localVarPostBody = r.zhimaCreditPayafteruseCreditbizorderOrderModel

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
			var v ZhimaCreditPayafteruseCreditbizorderOrderDefaultResponse
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


type ApiZhimaCreditPayafteruseCreditbizorderQueryRequest struct {
	ctx context.Context
	ApiService *ZhimaCreditPayafteruseCreditbizorderAPIService
	creditBizOrderId *string
	outOrderNo *string
}

// 信用服务订单号，out_order_no与credit_biz_order_id至少传一个
func (r ApiZhimaCreditPayafteruseCreditbizorderQueryRequest) CreditBizOrderId(creditBizOrderId string) ApiZhimaCreditPayafteruseCreditbizorderQueryRequest {
	r.creditBizOrderId = &creditBizOrderId
	return r
}

// 商户外部单号，out_order_no与credit_biz_order_id至少传一个
func (r ApiZhimaCreditPayafteruseCreditbizorderQueryRequest) OutOrderNo(outOrderNo string) ApiZhimaCreditPayafteruseCreditbizorderQueryRequest {
	r.outOrderNo = &outOrderNo
	return r
}

func (r ApiZhimaCreditPayafteruseCreditbizorderQueryRequest) Execute() (*ZhimaCreditPayafteruseCreditbizorderQueryResponseModel, *http.Response, error) {
	return r.ApiService.ZhimaCreditPayafteruseCreditbizorderQueryExecute(r)
}

/*
ZhimaCreditPayafteruseCreditbizorderQuery 信用服务订单查询

信用服务订单查询

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiZhimaCreditPayafteruseCreditbizorderQueryRequest
*/
func (r *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderQuery(ctx context.Context) ApiZhimaCreditPayafteruseCreditbizorderQueryRequest {
	return ApiZhimaCreditPayafteruseCreditbizorderQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ZhimaCreditPayafteruseCreditbizorderQueryResponseModel
func (a *ZhimaCreditPayafteruseCreditbizorderAPIService) ZhimaCreditPayafteruseCreditbizorderQueryExecute(r ApiZhimaCreditPayafteruseCreditbizorderQueryRequest) (*ZhimaCreditPayafteruseCreditbizorderQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ZhimaCreditPayafteruseCreditbizorderQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ZhimaCreditPayafteruseCreditbizorderAPIService.ZhimaCreditPayafteruseCreditbizorderQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/zhima/credit/payafteruse/creditbizorder/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.creditBizOrderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credit_biz_order_id", r.creditBizOrderId, "form", "")
	}
	if r.outOrderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_order_no", r.outOrderNo, "form", "")
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
			var v ZhimaCreditPayafteruseCreditbizorderQueryDefaultResponse
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


func (a *ZhimaCreditPayafteruseCreditbizorderAPIService) signRequest(req *http.Request) error {
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

func (a *ZhimaCreditPayafteruseCreditbizorderAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


