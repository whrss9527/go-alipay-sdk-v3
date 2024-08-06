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

// AlipayFundJointaccountListAPIService AlipayFundJointaccountListAPI service
type AlipayFundJointaccountListAPIService service

type ApiAlipayFundJointaccountListQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayFundJointaccountListAPIService
	productCode  *string
	bizScene     *string
	operateRole  *string
	agreementNo  *string
	identity     *string
	identityType *string
}

// 产品码
func (r ApiAlipayFundJointaccountListQueryRequest) ProductCode(productCode string) ApiAlipayFundJointaccountListQueryRequest {
	r.productCode = &productCode
	return r
}

// 业务场景
func (r ApiAlipayFundJointaccountListQueryRequest) BizScene(bizScene string) ApiAlipayFundJointaccountListQueryRequest {
	r.bizScene = &bizScene
	return r
}

// 角色：创建方(CREATOR)、参与方(PARTICIPANT)
func (r ApiAlipayFundJointaccountListQueryRequest) OperateRole(operateRole string) ApiAlipayFundJointaccountListQueryRequest {
	r.operateRole = &operateRole
	return r
}

// 授权协议号
func (r ApiAlipayFundJointaccountListQueryRequest) AgreementNo(agreementNo string) ApiAlipayFundJointaccountListQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 员工账号：  identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）；  是ALIPAY_LOGON_ID 填支付宝登录号
func (r ApiAlipayFundJointaccountListQueryRequest) Identity(identity string) ApiAlipayFundJointaccountListQueryRequest {
	r.identity = &identity
	return r
}

// 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式
func (r ApiAlipayFundJointaccountListQueryRequest) IdentityType(identityType string) ApiAlipayFundJointaccountListQueryRequest {
	r.identityType = &identityType
	return r
}

func (r ApiAlipayFundJointaccountListQueryRequest) Execute() (*AlipayFundJointaccountListQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundJointaccountListQueryExecute(r)
}

/*
AlipayFundJointaccountListQuery 企业查询代付账户列表

企业查询代付账户列表

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundJointaccountListQueryRequest
*/
func (r *AlipayFundJointaccountListAPIService) AlipayFundJointaccountListQuery(ctx context.Context) ApiAlipayFundJointaccountListQueryRequest {
	return ApiAlipayFundJointaccountListQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundJointaccountListQueryResponseModel
func (a *AlipayFundJointaccountListAPIService) AlipayFundJointaccountListQueryExecute(r ApiAlipayFundJointaccountListQueryRequest) (*AlipayFundJointaccountListQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundJointaccountListQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundJointaccountListAPIService.AlipayFundJointaccountListQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/jointaccount/list/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.operateRole != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "operate_role", r.operateRole, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.identity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identity", r.identity, "form", "")
	}
	if r.identityType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identity_type", r.identityType, "form", "")
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
		var v AlipayFundJointaccountListQueryDefaultResponse
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

func (a *AlipayFundJointaccountListAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundJointaccountListAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
