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

// AlipayFundAccountAPIService AlipayFundAccountAPI service
type AlipayFundAccountAPIService service

type ApiAlipayFundAccountQueryRequest struct {
	ctx                context.Context
	ApiService         *AlipayFundAccountAPIService
	merchantUserId     *string
	alipayUserId       *string
	alipayOpenId       *string
	accountProductCode *string
	accountType        *string
	accountSceneCode   *string
	extInfo            *string
}

// 商户会员的唯一标识。如果传入的user_id为虚拟账户userId，此字段必传并比对一致性。
func (r ApiAlipayFundAccountQueryRequest) MerchantUserId(merchantUserId string) ApiAlipayFundAccountQueryRequest {
	r.merchantUserId = &merchantUserId
	return r
}

// 支付宝会员 id。
func (r ApiAlipayFundAccountQueryRequest) AlipayUserId(alipayUserId string) ApiAlipayFundAccountQueryRequest {
	r.alipayUserId = &alipayUserId
	return r
}

// 支付宝openId
func (r ApiAlipayFundAccountQueryRequest) AlipayOpenId(alipayOpenId string) ApiAlipayFundAccountQueryRequest {
	r.alipayOpenId = &alipayOpenId
	return r
}

// 开户产品码。如果查询托管子户余额，必传且必须传入与开户时传入的值一致。
func (r ApiAlipayFundAccountQueryRequest) AccountProductCode(accountProductCode string) ApiAlipayFundAccountQueryRequest {
	r.accountProductCode = &accountProductCode
	return r
}

// 查询的账号类型，查询余额账户值为ACCTRANS_ACCOUNT。必填。
func (r ApiAlipayFundAccountQueryRequest) AccountType(accountType string) ApiAlipayFundAccountQueryRequest {
	r.accountType = &accountType
	return r
}

// 开户场景码，与开户产品码不可同时传递。
func (r ApiAlipayFundAccountQueryRequest) AccountSceneCode(accountSceneCode string) ApiAlipayFundAccountQueryRequest {
	r.accountSceneCode = &accountSceneCode
	return r
}

// JSON格式，传递业务扩展参数。
func (r ApiAlipayFundAccountQueryRequest) ExtInfo(extInfo string) ApiAlipayFundAccountQueryRequest {
	r.extInfo = &extInfo
	return r
}

func (r ApiAlipayFundAccountQueryRequest) Execute() (*AlipayFundAccountQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundAccountQueryExecute(r)
}

/*
AlipayFundAccountQuery 支付宝资金账户资产查询接口

可查询请求方的支付宝账户余额信息。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundAccountQueryRequest
*/
func (r *AlipayFundAccountAPIService) AlipayFundAccountQuery(ctx context.Context) ApiAlipayFundAccountQueryRequest {
	return ApiAlipayFundAccountQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundAccountQueryResponseModel
func (a *AlipayFundAccountAPIService) AlipayFundAccountQueryExecute(r ApiAlipayFundAccountQueryRequest) (*AlipayFundAccountQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundAccountQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundAccountAPIService.AlipayFundAccountQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/account/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.merchantUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merchant_user_id", r.merchantUserId, "form", "")
	}
	if r.alipayUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_user_id", r.alipayUserId, "form", "")
	}
	if r.alipayOpenId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_open_id", r.alipayOpenId, "form", "")
	}
	if r.accountProductCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_product_code", r.accountProductCode, "form", "")
	}
	if r.accountType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType, "form", "")
	}
	if r.accountSceneCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_scene_code", r.accountSceneCode, "form", "")
	}
	if r.extInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ext_info", r.extInfo, "form", "")
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
		var v AlipayFundAccountQueryDefaultResponse
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

func (a *AlipayFundAccountAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundAccountAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
