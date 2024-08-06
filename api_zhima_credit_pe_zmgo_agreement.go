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


// ZhimaCreditPeZmgoAgreementAPIService ZhimaCreditPeZmgoAgreementAPI service
type ZhimaCreditPeZmgoAgreementAPIService service

type ApiZhimaCreditPeZmgoAgreementQueryRequest struct {
	ctx context.Context
	ApiService *ZhimaCreditPeZmgoAgreementAPIService
	agreementId *string
	alipayUserId *string
	openId *string
}

// 支付宝系统中用以唯一标识用户签约记录的编号，即花芝轻会员。传入该参数，会忽略其它所有参数。
func (r ApiZhimaCreditPeZmgoAgreementQueryRequest) AgreementId(agreementId string) ApiZhimaCreditPeZmgoAgreementQueryRequest {
	r.agreementId = &agreementId
	return r
}

// 买家在支付宝的用户id
func (r ApiZhimaCreditPeZmgoAgreementQueryRequest) AlipayUserId(alipayUserId string) ApiZhimaCreditPeZmgoAgreementQueryRequest {
	r.alipayUserId = &alipayUserId
	return r
}

// 买家在支付宝的用户id
func (r ApiZhimaCreditPeZmgoAgreementQueryRequest) OpenId(openId string) ApiZhimaCreditPeZmgoAgreementQueryRequest {
	r.openId = &openId
	return r
}

func (r ApiZhimaCreditPeZmgoAgreementQueryRequest) Execute() (*ZhimaCreditPeZmgoAgreementQueryResponseModel, *http.Response, error) {
	return r.ApiService.ZhimaCreditPeZmgoAgreementQueryExecute(r)
}

/*
ZhimaCreditPeZmgoAgreementQuery 芝麻Go协议查询接口

用户已经开通芝麻GO后，通过此接口查询协议。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiZhimaCreditPeZmgoAgreementQueryRequest
*/
func (r *ZhimaCreditPeZmgoAgreementAPIService) ZhimaCreditPeZmgoAgreementQuery(ctx context.Context) ApiZhimaCreditPeZmgoAgreementQueryRequest {
	return ApiZhimaCreditPeZmgoAgreementQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ZhimaCreditPeZmgoAgreementQueryResponseModel
func (a *ZhimaCreditPeZmgoAgreementAPIService) ZhimaCreditPeZmgoAgreementQueryExecute(r ApiZhimaCreditPeZmgoAgreementQueryRequest) (*ZhimaCreditPeZmgoAgreementQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ZhimaCreditPeZmgoAgreementQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ZhimaCreditPeZmgoAgreementAPIService.ZhimaCreditPeZmgoAgreementQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/zhima/credit/pe/zmgo/agreement/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.agreementId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_id", r.agreementId, "form", "")
	}
	if r.alipayUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_user_id", r.alipayUserId, "form", "")
	}
	if r.openId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open_id", r.openId, "form", "")
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
			var v ZhimaCreditPeZmgoAgreementQueryDefaultResponse
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


type ApiZhimaCreditPeZmgoAgreementUnsignRequest struct {
	ctx context.Context
	ApiService *ZhimaCreditPeZmgoAgreementAPIService
	zhimaCreditPeZmgoAgreementUnsignModel *ZhimaCreditPeZmgoAgreementUnsignModel
}

func (r ApiZhimaCreditPeZmgoAgreementUnsignRequest) ZhimaCreditPeZmgoAgreementUnsignModel(zhimaCreditPeZmgoAgreementUnsignModel ZhimaCreditPeZmgoAgreementUnsignModel) ApiZhimaCreditPeZmgoAgreementUnsignRequest {
	r.zhimaCreditPeZmgoAgreementUnsignModel = &zhimaCreditPeZmgoAgreementUnsignModel
	return r
}

func (r ApiZhimaCreditPeZmgoAgreementUnsignRequest) Execute() (*ZhimaCreditPeZmgoAgreementUnsignResponseModel, *http.Response, error) {
	return r.ApiService.ZhimaCreditPeZmgoAgreementUnsignExecute(r)
}

/*
ZhimaCreditPeZmgoAgreementUnsign 芝麻GO协议解约

芝麻GO协议解约

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiZhimaCreditPeZmgoAgreementUnsignRequest
*/
func (r *ZhimaCreditPeZmgoAgreementAPIService) ZhimaCreditPeZmgoAgreementUnsign(ctx context.Context) ApiZhimaCreditPeZmgoAgreementUnsignRequest {
	return ApiZhimaCreditPeZmgoAgreementUnsignRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ZhimaCreditPeZmgoAgreementUnsignResponseModel
func (a *ZhimaCreditPeZmgoAgreementAPIService) ZhimaCreditPeZmgoAgreementUnsignExecute(r ApiZhimaCreditPeZmgoAgreementUnsignRequest) (*ZhimaCreditPeZmgoAgreementUnsignResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ZhimaCreditPeZmgoAgreementUnsignResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ZhimaCreditPeZmgoAgreementAPIService.ZhimaCreditPeZmgoAgreementUnsign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/zhima/credit/pe/zmgo/agreement/unsign"

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
	localVarPostBody = r.zhimaCreditPeZmgoAgreementUnsignModel

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
			var v ZhimaCreditPeZmgoAgreementUnsignDefaultResponse
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


func (a *ZhimaCreditPeZmgoAgreementAPIService) signRequest(req *http.Request) error {
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

func (a *ZhimaCreditPeZmgoAgreementAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


