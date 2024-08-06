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


// AlipayUserAgreementAPIService AlipayUserAgreementAPI service
type AlipayUserAgreementAPIService service

type ApiAlipayUserAgreementMigrateRequest struct {
	ctx context.Context
	ApiService *AlipayUserAgreementAPIService
	alipayUserAgreementMigrateModel *AlipayUserAgreementMigrateModel
}

func (r ApiAlipayUserAgreementMigrateRequest) AlipayUserAgreementMigrateModel(alipayUserAgreementMigrateModel AlipayUserAgreementMigrateModel) ApiAlipayUserAgreementMigrateRequest {
	r.alipayUserAgreementMigrateModel = &alipayUserAgreementMigrateModel
	return r
}

func (r ApiAlipayUserAgreementMigrateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayUserAgreementMigrateExecute(r)
}

/*
AlipayUserAgreementMigrate 代扣协议迁移

由商户调用，将商户与用户签署的代扣协议内容进行迁移，包括协议主体迁移等

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayUserAgreementMigrateRequest
*/
func (r *AlipayUserAgreementAPIService) AlipayUserAgreementMigrate(ctx context.Context) ApiAlipayUserAgreementMigrateRequest {
	return ApiAlipayUserAgreementMigrateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayUserAgreementAPIService) AlipayUserAgreementMigrateExecute(r ApiAlipayUserAgreementMigrateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayUserAgreementAPIService.AlipayUserAgreementMigrate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/user/agreement/migrate"

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
	localVarPostBody = r.alipayUserAgreementMigrateModel

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
			var v AlipayUserAgreementMigrateDefaultResponse
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


type ApiAlipayUserAgreementQueryRequest struct {
	ctx context.Context
	ApiService *AlipayUserAgreementAPIService
	personalProductCode *string
	alipayUserId *string
	alipayOpenId *string
	alipayLogonId *string
	signScene *string
	externalAgreementNo *string
	thirdPartyType *string
	agreementNo *string
}

// 协议产品码，商户和支付宝签约时确定，商户可咨询技术支持。
func (r ApiAlipayUserAgreementQueryRequest) PersonalProductCode(personalProductCode string) ApiAlipayUserAgreementQueryRequest {
	r.personalProductCode = &personalProductCode
	return r
}

// 用户的支付宝账号对应 的支付宝唯一用户号，以 2088 开头的 16 位纯数字 组成。 本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
func (r ApiAlipayUserAgreementQueryRequest) AlipayUserId(alipayUserId string) ApiAlipayUserAgreementQueryRequest {
	r.alipayUserId = &alipayUserId
	return r
}

// 用户的支付宝账号对应 的支付宝唯一用户号， 本参数与alipay_logon_id若都填写，则以本参数为准，优先级高于 alipay_logon_id。
func (r ApiAlipayUserAgreementQueryRequest) AlipayOpenId(alipayOpenId string) ApiAlipayUserAgreementQueryRequest {
	r.alipayOpenId = &alipayOpenId
	return r
}

// 用户的支付宝登录账号，支持邮箱或手机号码格式。本参数与alipay_open_id 或 alipay_user_id 同时填写，优先按照 alipay_open_id 或 alipay_user_id 处理。
func (r ApiAlipayUserAgreementQueryRequest) AlipayLogonId(alipayLogonId string) ApiAlipayUserAgreementQueryRequest {
	r.alipayLogonId = &alipayLogonId
	return r
}

// 签约场景码，该值需要与系统/页面签约接口调用时传入的值保持一 致。如：周期扣款场景与调用 alipay.user.agreement.page.sign(支付宝个人协议页面签约接口) 签约时的 sign_scene 相同。  注意：当传入商户签约号 external_agreement_no 时，该值不能为空或默认值 DEFAULT|DEFAULT。 
func (r ApiAlipayUserAgreementQueryRequest) SignScene(signScene string) ApiAlipayUserAgreementQueryRequest {
	r.signScene = &signScene
	return r
}

// 代扣协议中标示用户的唯一签约号(确保在商户系统中 唯一)。 格式规则:支持大写小写字母和数字，最长 32 位。
func (r ApiAlipayUserAgreementQueryRequest) ExternalAgreementNo(externalAgreementNo string) ApiAlipayUserAgreementQueryRequest {
	r.externalAgreementNo = &externalAgreementNo
	return r
}

// 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 默认为PARTNER。
func (r ApiAlipayUserAgreementQueryRequest) ThirdPartyType(thirdPartyType string) ApiAlipayUserAgreementQueryRequest {
	r.thirdPartyType = &thirdPartyType
	return r
}

// 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ） ，如果传了该参数，其他参数会被忽略
func (r ApiAlipayUserAgreementQueryRequest) AgreementNo(agreementNo string) ApiAlipayUserAgreementQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

func (r ApiAlipayUserAgreementQueryRequest) Execute() (*AlipayUserAgreementQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayUserAgreementQueryExecute(r)
}

/*
AlipayUserAgreementQuery 支付宝个人代扣协议查询接口

支付宝个人代扣协议查询接口

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayUserAgreementQueryRequest
*/
func (r *AlipayUserAgreementAPIService) AlipayUserAgreementQuery(ctx context.Context) ApiAlipayUserAgreementQueryRequest {
	return ApiAlipayUserAgreementQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayUserAgreementQueryResponseModel
func (a *AlipayUserAgreementAPIService) AlipayUserAgreementQueryExecute(r ApiAlipayUserAgreementQueryRequest) (*AlipayUserAgreementQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayUserAgreementQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayUserAgreementAPIService.AlipayUserAgreementQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/user/agreement/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.personalProductCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "personal_product_code", r.personalProductCode, "form", "")
	}
	if r.alipayUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_user_id", r.alipayUserId, "form", "")
	}
	if r.alipayOpenId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_open_id", r.alipayOpenId, "form", "")
	}
	if r.alipayLogonId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_logon_id", r.alipayLogonId, "form", "")
	}
	if r.signScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sign_scene", r.signScene, "form", "")
	}
	if r.externalAgreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_agreement_no", r.externalAgreementNo, "form", "")
	}
	if r.thirdPartyType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "third_party_type", r.thirdPartyType, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
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
			var v AlipayUserAgreementQueryDefaultResponse
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


type ApiAlipayUserAgreementTransferRequest struct {
	ctx context.Context
	ApiService *AlipayUserAgreementAPIService
	alipayUserAgreementTransferModel *AlipayUserAgreementTransferModel
}

func (r ApiAlipayUserAgreementTransferRequest) AlipayUserAgreementTransferModel(alipayUserAgreementTransferModel AlipayUserAgreementTransferModel) ApiAlipayUserAgreementTransferRequest {
	r.alipayUserAgreementTransferModel = &alipayUserAgreementTransferModel
	return r
}

func (r ApiAlipayUserAgreementTransferRequest) Execute() (*AlipayUserAgreementTransferResponseModel, *http.Response, error) {
	return r.ApiService.AlipayUserAgreementTransferExecute(r)
}

/*
AlipayUserAgreementTransfer 协议由普通通用代扣协议产品转移到周期扣协议产品

由商户调用，将商户之前通用代扣产品转移到周期扣的协议产品

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayUserAgreementTransferRequest
*/
func (r *AlipayUserAgreementAPIService) AlipayUserAgreementTransfer(ctx context.Context) ApiAlipayUserAgreementTransferRequest {
	return ApiAlipayUserAgreementTransferRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayUserAgreementTransferResponseModel
func (a *AlipayUserAgreementAPIService) AlipayUserAgreementTransferExecute(r ApiAlipayUserAgreementTransferRequest) (*AlipayUserAgreementTransferResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayUserAgreementTransferResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayUserAgreementAPIService.AlipayUserAgreementTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/user/agreement/transfer"

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
	localVarPostBody = r.alipayUserAgreementTransferModel

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
			var v AlipayUserAgreementTransferDefaultResponse
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


type ApiAlipayUserAgreementUnsignRequest struct {
	ctx context.Context
	ApiService *AlipayUserAgreementAPIService
	alipayUserAgreementUnsignModel *AlipayUserAgreementUnsignModel
}

func (r ApiAlipayUserAgreementUnsignRequest) AlipayUserAgreementUnsignModel(alipayUserAgreementUnsignModel AlipayUserAgreementUnsignModel) ApiAlipayUserAgreementUnsignRequest {
	r.alipayUserAgreementUnsignModel = &alipayUserAgreementUnsignModel
	return r
}

func (r ApiAlipayUserAgreementUnsignRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayUserAgreementUnsignExecute(r)
}

/*
AlipayUserAgreementUnsign 支付宝个人代扣协议解约接口

支付宝个人代扣协议解约接口

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayUserAgreementUnsignRequest
*/
func (r *AlipayUserAgreementAPIService) AlipayUserAgreementUnsign(ctx context.Context) ApiAlipayUserAgreementUnsignRequest {
	return ApiAlipayUserAgreementUnsignRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayUserAgreementAPIService) AlipayUserAgreementUnsignExecute(r ApiAlipayUserAgreementUnsignRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayUserAgreementAPIService.AlipayUserAgreementUnsign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/user/agreement/unsign"

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
	localVarPostBody = r.alipayUserAgreementUnsignModel

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
			var v AlipayUserAgreementUnsignDefaultResponse
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


func (a *AlipayUserAgreementAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayUserAgreementAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

