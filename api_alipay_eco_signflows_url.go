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

// AlipayEcoSignflowsUrlAPIService AlipayEcoSignflowsUrlAPI service
type AlipayEcoSignflowsUrlAPIService service

type ApiAlipayEcoSignflowsUrlQueryRequest struct {
	ctx                 context.Context
	ApiService          *AlipayEcoSignflowsUrlAPIService
	targetAppId         *string
	flowId              *string
	thirdPartyUserId    *string
	orgThirdPartyUserId *string
}

// 目标isv应用ID
func (r ApiAlipayEcoSignflowsUrlQueryRequest) TargetAppId(targetAppId string) ApiAlipayEcoSignflowsUrlQueryRequest {
	r.targetAppId = &targetAppId
	return r
}

// 流程id，通过 &lt;a href &#x3D;\&quot;https://opendocs.alipay.com/apis/api_50/alipay.eco.contract.signflows.create\&quot;&gt;创建电子合同签署流程&lt;/a&gt;(alipay.eco.contract.signflows.create)接口获取。
func (r ApiAlipayEcoSignflowsUrlQueryRequest) FlowId(flowId string) ApiAlipayEcoSignflowsUrlQueryRequest {
	r.flowId = &flowId
	return r
}

// 创建流程时指定个人唯一标识
func (r ApiAlipayEcoSignflowsUrlQueryRequest) ThirdPartyUserId(thirdPartyUserId string) ApiAlipayEcoSignflowsUrlQueryRequest {
	r.thirdPartyUserId = &thirdPartyUserId
	return r
}

// 创建流程时指定企业唯一标识
func (r ApiAlipayEcoSignflowsUrlQueryRequest) OrgThirdPartyUserId(orgThirdPartyUserId string) ApiAlipayEcoSignflowsUrlQueryRequest {
	r.orgThirdPartyUserId = &orgThirdPartyUserId
	return r
}

func (r ApiAlipayEcoSignflowsUrlQueryRequest) Execute() (*AlipayEcoSignflowsUrlQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoSignflowsUrlQueryExecute(r)
}

/*
AlipayEcoSignflowsUrlQuery 获取签署地址

创建流程后，获取指定签署人的签署链接地址，可在应用内集成H5签署页，或者通过短信发送签署链接。
传入个人唯一标识，则获取的签署任务链接仅包含个人人的签署任务；如同时传入企业唯一标识，则获取的签署任务链接包含企业与个人的签署任务。
预览链接：支持签署人先查看合同原文，后进行登录并完成签署。适用于应用内集成场景。
签署链接：签署人需要登录后查看合同原文并签署。适用用短信发送场景。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEcoSignflowsUrlQueryRequest
*/
func (r *AlipayEcoSignflowsUrlAPIService) AlipayEcoSignflowsUrlQuery(ctx context.Context) ApiAlipayEcoSignflowsUrlQueryRequest {
	return ApiAlipayEcoSignflowsUrlQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEcoSignflowsUrlQueryResponseModel
func (a *AlipayEcoSignflowsUrlAPIService) AlipayEcoSignflowsUrlQueryExecute(r ApiAlipayEcoSignflowsUrlQueryRequest) (*AlipayEcoSignflowsUrlQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEcoSignflowsUrlQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoSignflowsUrlAPIService.AlipayEcoSignflowsUrlQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/signflows/url/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.targetAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "target_app_id", r.targetAppId, "form", "")
	}
	if r.flowId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_id", r.flowId, "form", "")
	}
	if r.thirdPartyUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "third_party_user_id", r.thirdPartyUserId, "form", "")
	}
	if r.orgThirdPartyUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "org_third_party_user_id", r.orgThirdPartyUserId, "form", "")
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
		var v AlipayEcoSignflowsUrlQueryDefaultResponse
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

func (a *AlipayEcoSignflowsUrlAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoSignflowsUrlAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
