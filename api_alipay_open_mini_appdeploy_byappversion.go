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


// AlipayOpenMiniAppdeployByappversionAPIService AlipayOpenMiniAppdeployByappversionAPI service
type AlipayOpenMiniAppdeployByappversionAPIService service

type ApiAlipayOpenMiniAppdeployByappversionQueryRequest struct {
	ctx context.Context
	ApiService *AlipayOpenMiniAppdeployByappversionAPIService
	miniAppId *string
	bundleId *string
	instCode *string
	appVersion *string
}

// 小程序ID
func (r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) MiniAppId(miniAppId string) ApiAlipayOpenMiniAppdeployByappversionQueryRequest {
	r.miniAppId = &miniAppId
	return r
}

// 端标识
func (r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) BundleId(bundleId string) ApiAlipayOpenMiniAppdeployByappversionQueryRequest {
	r.bundleId = &bundleId
	return r
}

// 租户
func (r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) InstCode(instCode string) ApiAlipayOpenMiniAppdeployByappversionQueryRequest {
	r.instCode = &instCode
	return r
}

// 小程序版本
func (r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) AppVersion(appVersion string) ApiAlipayOpenMiniAppdeployByappversionQueryRequest {
	r.appVersion = &appVersion
	return r
}

func (r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) Execute() (*AlipayOpenMiniAppdeployByappversionQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenMiniAppdeployByappversionQueryExecute(r)
}

/*
AlipayOpenMiniAppdeployByappversionQuery 通过版本查询小程序发布

伙伴平台使用，用于向开发者展示小程序发布信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenMiniAppdeployByappversionQueryRequest
*/
func (r *AlipayOpenMiniAppdeployByappversionAPIService) AlipayOpenMiniAppdeployByappversionQuery(ctx context.Context) ApiAlipayOpenMiniAppdeployByappversionQueryRequest {
	return ApiAlipayOpenMiniAppdeployByappversionQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayOpenMiniAppdeployByappversionQueryResponseModel
func (a *AlipayOpenMiniAppdeployByappversionAPIService) AlipayOpenMiniAppdeployByappversionQueryExecute(r ApiAlipayOpenMiniAppdeployByappversionQueryRequest) (*AlipayOpenMiniAppdeployByappversionQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayOpenMiniAppdeployByappversionQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenMiniAppdeployByappversionAPIService.AlipayOpenMiniAppdeployByappversionQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/mini/appdeploy/byappversion/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.miniAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mini_app_id", r.miniAppId, "form", "")
	}
	if r.bundleId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bundle_id", r.bundleId, "form", "")
	}
	if r.instCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inst_code", r.instCode, "form", "")
	}
	if r.appVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_version", r.appVersion, "form", "")
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
			var v AlipayOpenMiniAppdeployByappversionQueryDefaultResponse
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


func (a *AlipayOpenMiniAppdeployByappversionAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenMiniAppdeployByappversionAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


