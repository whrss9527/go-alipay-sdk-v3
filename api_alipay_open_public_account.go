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


// AlipayOpenPublicAccountAPIService AlipayOpenPublicAccountAPI service
type AlipayOpenPublicAccountAPIService service

type ApiAlipayOpenPublicAccountCreateRequest struct {
	ctx context.Context
	ApiService *AlipayOpenPublicAccountAPIService
	alipayOpenPublicAccountCreateModel *AlipayOpenPublicAccountCreateModel
}

func (r ApiAlipayOpenPublicAccountCreateRequest) AlipayOpenPublicAccountCreateModel(alipayOpenPublicAccountCreateModel AlipayOpenPublicAccountCreateModel) ApiAlipayOpenPublicAccountCreateRequest {
	r.alipayOpenPublicAccountCreateModel = &alipayOpenPublicAccountCreateModel
	return r
}

func (r ApiAlipayOpenPublicAccountCreateRequest) Execute() (*AlipayOpenPublicAccountCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicAccountCreateExecute(r)
}

/*
AlipayOpenPublicAccountCreate 添加绑定商户会员号

当用户成为商户的关注用户后，可以在商户的服务窗中点击“添加绑定商户会员号”功能，支付宝系统收到操作请求后将该动作通知给商户（使用用户发送信息到商户接口，其中eventType（事件类型）为click，actionParam（按钮标识）为authentication），商户根据此通知调用商户回复消息接口（其中须包含Url链接地址），支付宝收到商户的回复消息中的链接地址后，自动跳转至商户平台的上商户会员绑定界面中，让用户完成账户绑定。 当用户有效完成账户绑定后，商户调用本接口，把绑定结果数据通知给支付宝。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenPublicAccountCreateRequest
*/
func (r *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountCreate(ctx context.Context) ApiAlipayOpenPublicAccountCreateRequest {
	return ApiAlipayOpenPublicAccountCreateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayOpenPublicAccountCreateResponseModel
func (a *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountCreateExecute(r ApiAlipayOpenPublicAccountCreateRequest) (*AlipayOpenPublicAccountCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayOpenPublicAccountCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicAccountAPIService.AlipayOpenPublicAccountCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/account/create"

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
	localVarPostBody = r.alipayOpenPublicAccountCreateModel

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
			var v AlipayOpenPublicAccountCreateDefaultResponse
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


type ApiAlipayOpenPublicAccountDeleteRequest struct {
	ctx context.Context
	ApiService *AlipayOpenPublicAccountAPIService
	agreementId *string
	bindAccountNo *string
	fromUserId *string
	openId *string
}

// 协议号，商户会员在支付宝服务窗账号中的唯一标识，与bind_account_no不能同时为空
func (r ApiAlipayOpenPublicAccountDeleteRequest) AgreementId(agreementId string) ApiAlipayOpenPublicAccountDeleteRequest {
	r.agreementId = &agreementId
	return r
}

// 绑定帐号，建议在开发者的系统中保持唯一性，与agreement_id不能同时为空
func (r ApiAlipayOpenPublicAccountDeleteRequest) BindAccountNo(bindAccountNo string) ApiAlipayOpenPublicAccountDeleteRequest {
	r.bindAccountNo = &bindAccountNo
	return r
}

// 绑定用户的支付宝userid，2088开头16位长度的字符串，与agreementId不能同时为空
func (r ApiAlipayOpenPublicAccountDeleteRequest) FromUserId(fromUserId string) ApiAlipayOpenPublicAccountDeleteRequest {
	r.fromUserId = &fromUserId
	return r
}

// 支付宝用户的唯一标识
func (r ApiAlipayOpenPublicAccountDeleteRequest) OpenId(openId string) ApiAlipayOpenPublicAccountDeleteRequest {
	r.openId = &openId
	return r
}

func (r ApiAlipayOpenPublicAccountDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicAccountDeleteExecute(r)
}

/*
AlipayOpenPublicAccountDelete 解除绑定商户会员号

在支付宝服务窗中目前一共有三种解除绑定商户会员号的场景，具体如下：
用户取消关注服务窗
用户在服务窗内手动执行解除绑定操作
开发者调用解除绑定商户会员号接口解除绑定

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenPublicAccountDeleteRequest
*/
func (r *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountDelete(ctx context.Context) ApiAlipayOpenPublicAccountDeleteRequest {
	return ApiAlipayOpenPublicAccountDeleteRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountDeleteExecute(r ApiAlipayOpenPublicAccountDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicAccountAPIService.AlipayOpenPublicAccountDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/account/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.agreementId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_id", r.agreementId, "form", "")
	}
	if r.bindAccountNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bind_account_no", r.bindAccountNo, "form", "")
	}
	if r.fromUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_user_id", r.fromUserId, "form", "")
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
			var v AlipayOpenPublicAccountDeleteDefaultResponse
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


type ApiAlipayOpenPublicAccountQueryRequest struct {
	ctx context.Context
	ApiService *AlipayOpenPublicAccountAPIService
	userId *string
	openId *string
}

// 用户的支付宝用户号，2088开头。
func (r ApiAlipayOpenPublicAccountQueryRequest) UserId(userId string) ApiAlipayOpenPublicAccountQueryRequest {
	r.userId = &userId
	return r
}

// 支付宝用户的唯一标识
func (r ApiAlipayOpenPublicAccountQueryRequest) OpenId(openId string) ApiAlipayOpenPublicAccountQueryRequest {
	r.openId = &openId
	return r
}

func (r ApiAlipayOpenPublicAccountQueryRequest) Execute() (*AlipayOpenPublicAccountQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicAccountQueryExecute(r)
}

/*
AlipayOpenPublicAccountQuery 查询绑定商户会员号

当用户成为商户的关注用户之后，则商户可以通过本接口查询关注者的绑定账户，以便补全异常情况下的单边账户数据。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenPublicAccountQueryRequest
*/
func (r *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountQuery(ctx context.Context) ApiAlipayOpenPublicAccountQueryRequest {
	return ApiAlipayOpenPublicAccountQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayOpenPublicAccountQueryResponseModel
func (a *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountQueryExecute(r ApiAlipayOpenPublicAccountQueryRequest) (*AlipayOpenPublicAccountQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayOpenPublicAccountQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicAccountAPIService.AlipayOpenPublicAccountQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/account/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
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
			var v AlipayOpenPublicAccountQueryDefaultResponse
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


type ApiAlipayOpenPublicAccountResetRequest struct {
	ctx context.Context
	ApiService *AlipayOpenPublicAccountAPIService
	alipayOpenPublicAccountResetModel *AlipayOpenPublicAccountResetModel
}

func (r ApiAlipayOpenPublicAccountResetRequest) AlipayOpenPublicAccountResetModel(alipayOpenPublicAccountResetModel AlipayOpenPublicAccountResetModel) ApiAlipayOpenPublicAccountResetRequest {
	r.alipayOpenPublicAccountResetModel = &alipayOpenPublicAccountResetModel
	return r
}

func (r ApiAlipayOpenPublicAccountResetRequest) Execute() (*AlipayOpenPublicAccountResetResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicAccountResetExecute(r)
}

/*
AlipayOpenPublicAccountReset 重新设置绑定商家会员号

如果商户想要重置已经添加的外部账户，可以通过该接口完成。重置后，原有的外部户将删除，新的外部户添加进去。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayOpenPublicAccountResetRequest
*/
func (r *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountReset(ctx context.Context) ApiAlipayOpenPublicAccountResetRequest {
	return ApiAlipayOpenPublicAccountResetRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayOpenPublicAccountResetResponseModel
func (a *AlipayOpenPublicAccountAPIService) AlipayOpenPublicAccountResetExecute(r ApiAlipayOpenPublicAccountResetRequest) (*AlipayOpenPublicAccountResetResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayOpenPublicAccountResetResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicAccountAPIService.AlipayOpenPublicAccountReset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/account/reset"

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
	localVarPostBody = r.alipayOpenPublicAccountResetModel

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
			var v AlipayOpenPublicAccountResetDefaultResponse
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


func (a *AlipayOpenPublicAccountAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenPublicAccountAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

