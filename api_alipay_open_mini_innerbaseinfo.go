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


// AlipayOpenMiniInnerbaseinfoAPIService AlipayOpenMiniInnerbaseinfoAPI service
type AlipayOpenMiniInnerbaseinfoAPIService service

type ApiAlipayOpenMiniInnerbaseinfoQueryRequest struct {
	ctx context.Context
	ApiService *AlipayOpenMiniInnerbaseinfoAPIService
	miniAppId *string
	instCode *string
	miniAppName *string
	appSubType *string
}

// 小程序ID，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。
func (r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) MiniAppId(miniAppId string) ApiAlipayOpenMiniInnerbaseinfoQueryRequest {
	r.miniAppId = &miniAppId
	return r
}

// 租户code，alipay or taobao
func (r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) InstCode(instCode string) ApiAlipayOpenMiniInnerbaseinfoQueryRequest {
	r.instCode = &instCode
	return r
}

// 小程序name，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。
func (r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) MiniAppName(miniAppName string) ApiAlipayOpenMiniInnerbaseinfoQueryRequest {
	r.miniAppName = &miniAppName
	return r
}

// 小程序类型，TINYAPP_TEMPLATE，TINYAPP_NORMAL，TINYAPP_PLUGIN，使用mini_app_name查询的时候，该字段要求必传。
func (r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) AppSubType(appSubType string) ApiAlipayOpenMiniInnerbaseinfoQueryRequest {
	r.appSubType = &appSubType
	return r
}

func (r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) Execute() (*AlipayOpenMiniInnerbaseinfoQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenMiniInnerbaseinfoQueryExecute(r)
}

/*
AlipayOpenMiniInnerbaseinfoQuery 内部小程序-应用信息查询

查询小程序应用信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenMiniInnerbaseinfoQueryRequest
*/
func (r *AlipayOpenMiniInnerbaseinfoAPIService) AlipayOpenMiniInnerbaseinfoQuery(ctx context.Context) ApiAlipayOpenMiniInnerbaseinfoQueryRequest {
	return ApiAlipayOpenMiniInnerbaseinfoQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayOpenMiniInnerbaseinfoQueryResponseModel
func (a *AlipayOpenMiniInnerbaseinfoAPIService) AlipayOpenMiniInnerbaseinfoQueryExecute(r ApiAlipayOpenMiniInnerbaseinfoQueryRequest) (*AlipayOpenMiniInnerbaseinfoQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayOpenMiniInnerbaseinfoQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenMiniInnerbaseinfoAPIService.AlipayOpenMiniInnerbaseinfoQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/mini/innerbaseinfo/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.miniAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mini_app_id", r.miniAppId, "form", "")
	}
	if r.instCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inst_code", r.instCode, "form", "")
	}
	if r.miniAppName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mini_app_name", r.miniAppName, "form", "")
	}
	if r.appSubType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_sub_type", r.appSubType, "form", "")
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
			var v AlipayOpenMiniInnerbaseinfoQueryDefaultResponse
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


func (a *AlipayOpenMiniInnerbaseinfoAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenMiniInnerbaseinfoAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


