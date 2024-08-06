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


// AlipayOpenServicemarketCommodityShopAPIService AlipayOpenServicemarketCommodityShopAPI service
type AlipayOpenServicemarketCommodityShopAPIService service

type ApiAlipayOpenServicemarketCommodityShopOfflineRequest struct {
	ctx context.Context
	ApiService *AlipayOpenServicemarketCommodityShopAPIService
	alipayOpenServicemarketCommodityShopOfflineModel *AlipayOpenServicemarketCommodityShopOfflineModel
}

func (r ApiAlipayOpenServicemarketCommodityShopOfflineRequest) AlipayOpenServicemarketCommodityShopOfflineModel(alipayOpenServicemarketCommodityShopOfflineModel AlipayOpenServicemarketCommodityShopOfflineModel) ApiAlipayOpenServicemarketCommodityShopOfflineRequest {
	r.alipayOpenServicemarketCommodityShopOfflineModel = &alipayOpenServicemarketCommodityShopOfflineModel
	return r
}

func (r ApiAlipayOpenServicemarketCommodityShopOfflineRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenServicemarketCommodityShopOfflineExecute(r)
}

/*
AlipayOpenServicemarketCommodityShopOffline 门店插件下架操作

本接口需要商户授权服务商门店操作权限后，服务商可对商户门店上的插件进行下架操作

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenServicemarketCommodityShopOfflineRequest
*/
func (r *AlipayOpenServicemarketCommodityShopAPIService) AlipayOpenServicemarketCommodityShopOffline(ctx context.Context) ApiAlipayOpenServicemarketCommodityShopOfflineRequest {
	return ApiAlipayOpenServicemarketCommodityShopOfflineRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayOpenServicemarketCommodityShopAPIService) AlipayOpenServicemarketCommodityShopOfflineExecute(r ApiAlipayOpenServicemarketCommodityShopOfflineRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenServicemarketCommodityShopAPIService.AlipayOpenServicemarketCommodityShopOffline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/servicemarket/commodity/shop/offline"

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
	localVarPostBody = r.alipayOpenServicemarketCommodityShopOfflineModel

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
			var v AlipayOpenServicemarketCommodityShopOfflineDefaultResponse
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


type ApiAlipayOpenServicemarketCommodityShopOnlineRequest struct {
	ctx context.Context
	ApiService *AlipayOpenServicemarketCommodityShopAPIService
	alipayOpenServicemarketCommodityShopOnlineModel *AlipayOpenServicemarketCommodityShopOnlineModel
}

func (r ApiAlipayOpenServicemarketCommodityShopOnlineRequest) AlipayOpenServicemarketCommodityShopOnlineModel(alipayOpenServicemarketCommodityShopOnlineModel AlipayOpenServicemarketCommodityShopOnlineModel) ApiAlipayOpenServicemarketCommodityShopOnlineRequest {
	r.alipayOpenServicemarketCommodityShopOnlineModel = &alipayOpenServicemarketCommodityShopOnlineModel
	return r
}

func (r ApiAlipayOpenServicemarketCommodityShopOnlineRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenServicemarketCommodityShopOnlineExecute(r)
}

/*
AlipayOpenServicemarketCommodityShopOnline 门店插件上架操作

本接口用于上架商户门店订购的服务。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenServicemarketCommodityShopOnlineRequest
*/
func (r *AlipayOpenServicemarketCommodityShopAPIService) AlipayOpenServicemarketCommodityShopOnline(ctx context.Context) ApiAlipayOpenServicemarketCommodityShopOnlineRequest {
	return ApiAlipayOpenServicemarketCommodityShopOnlineRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayOpenServicemarketCommodityShopAPIService) AlipayOpenServicemarketCommodityShopOnlineExecute(r ApiAlipayOpenServicemarketCommodityShopOnlineRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenServicemarketCommodityShopAPIService.AlipayOpenServicemarketCommodityShopOnline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/servicemarket/commodity/shop/online"

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
	localVarPostBody = r.alipayOpenServicemarketCommodityShopOnlineModel

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
			var v AlipayOpenServicemarketCommodityShopOnlineDefaultResponse
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


func (a *AlipayOpenServicemarketCommodityShopAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenServicemarketCommodityShopAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


