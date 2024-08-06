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

// AlipayCommerceIotSdarttoolMessageAPIService AlipayCommerceIotSdarttoolMessageAPI service
type AlipayCommerceIotSdarttoolMessageAPIService service

type ApiAlipayCommerceIotSdarttoolMessageQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayCommerceIotSdarttoolMessageAPIService
	messageNo  *string
}

// 消息ID
func (r ApiAlipayCommerceIotSdarttoolMessageQueryRequest) MessageNo(messageNo string) ApiAlipayCommerceIotSdarttoolMessageQueryRequest {
	r.messageNo = &messageNo
	return r
}

func (r ApiAlipayCommerceIotSdarttoolMessageQueryRequest) Execute() (*AlipayCommerceIotSdarttoolMessageQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceIotSdarttoolMessageQueryExecute(r)
}

/*
AlipayCommerceIotSdarttoolMessageQuery 消息查询

isv通过消息id查询消息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceIotSdarttoolMessageQueryRequest
*/
func (r *AlipayCommerceIotSdarttoolMessageAPIService) AlipayCommerceIotSdarttoolMessageQuery(ctx context.Context) ApiAlipayCommerceIotSdarttoolMessageQueryRequest {
	return ApiAlipayCommerceIotSdarttoolMessageQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceIotSdarttoolMessageQueryResponseModel
func (a *AlipayCommerceIotSdarttoolMessageAPIService) AlipayCommerceIotSdarttoolMessageQueryExecute(r ApiAlipayCommerceIotSdarttoolMessageQueryRequest) (*AlipayCommerceIotSdarttoolMessageQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceIotSdarttoolMessageQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceIotSdarttoolMessageAPIService.AlipayCommerceIotSdarttoolMessageQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/iot/sdarttool/message/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.messageNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "message_no", r.messageNo, "form", "")
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
		var v AlipayCommerceIotSdarttoolMessageQueryDefaultResponse
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

type ApiAlipayCommerceIotSdarttoolMessageSendRequest struct {
	ctx                                        context.Context
	ApiService                                 *AlipayCommerceIotSdarttoolMessageAPIService
	alipayCommerceIotSdarttoolMessageSendModel *AlipayCommerceIotSdarttoolMessageSendModel
}

func (r ApiAlipayCommerceIotSdarttoolMessageSendRequest) AlipayCommerceIotSdarttoolMessageSendModel(alipayCommerceIotSdarttoolMessageSendModel AlipayCommerceIotSdarttoolMessageSendModel) ApiAlipayCommerceIotSdarttoolMessageSendRequest {
	r.alipayCommerceIotSdarttoolMessageSendModel = &alipayCommerceIotSdarttoolMessageSendModel
	return r
}

func (r ApiAlipayCommerceIotSdarttoolMessageSendRequest) Execute() (*AlipayCommerceIotSdarttoolMessageSendResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceIotSdarttoolMessageSendExecute(r)
}

/*
AlipayCommerceIotSdarttoolMessageSend IOT消息服务

ISV通过此接口向IOT设备推送消息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceIotSdarttoolMessageSendRequest
*/
func (r *AlipayCommerceIotSdarttoolMessageAPIService) AlipayCommerceIotSdarttoolMessageSend(ctx context.Context) ApiAlipayCommerceIotSdarttoolMessageSendRequest {
	return ApiAlipayCommerceIotSdarttoolMessageSendRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceIotSdarttoolMessageSendResponseModel
func (a *AlipayCommerceIotSdarttoolMessageAPIService) AlipayCommerceIotSdarttoolMessageSendExecute(r ApiAlipayCommerceIotSdarttoolMessageSendRequest) (*AlipayCommerceIotSdarttoolMessageSendResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceIotSdarttoolMessageSendResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceIotSdarttoolMessageAPIService.AlipayCommerceIotSdarttoolMessageSend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/iot/sdarttool/message/send"

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
	localVarPostBody = r.alipayCommerceIotSdarttoolMessageSendModel

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
		var v AlipayCommerceIotSdarttoolMessageSendDefaultResponse
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

func (a *AlipayCommerceIotSdarttoolMessageAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceIotSdarttoolMessageAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
