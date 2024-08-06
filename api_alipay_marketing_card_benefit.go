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


// AlipayMarketingCardBenefitAPIService AlipayMarketingCardBenefitAPI service
type AlipayMarketingCardBenefitAPIService service

type ApiAlipayMarketingCardBenefitCreateRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingCardBenefitAPIService
	alipayMarketingCardBenefitCreateModel *AlipayMarketingCardBenefitCreateModel
}

func (r ApiAlipayMarketingCardBenefitCreateRequest) AlipayMarketingCardBenefitCreateModel(alipayMarketingCardBenefitCreateModel AlipayMarketingCardBenefitCreateModel) ApiAlipayMarketingCardBenefitCreateRequest {
	r.alipayMarketingCardBenefitCreateModel = &alipayMarketingCardBenefitCreateModel
	return r
}

func (r ApiAlipayMarketingCardBenefitCreateRequest) Execute() (*AlipayMarketingCardBenefitCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingCardBenefitCreateExecute(r)
}

/*
AlipayMarketingCardBenefitCreate 会员卡模板外部权益创建

会员卡模板外部权益创建，创建的权益内容信息将展示在卡包中卡详情页中部区域。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingCardBenefitCreateRequest
*/
func (r *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitCreate(ctx context.Context) ApiAlipayMarketingCardBenefitCreateRequest {
	return ApiAlipayMarketingCardBenefitCreateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingCardBenefitCreateResponseModel
func (a *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitCreateExecute(r ApiAlipayMarketingCardBenefitCreateRequest) (*AlipayMarketingCardBenefitCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingCardBenefitCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingCardBenefitAPIService.AlipayMarketingCardBenefitCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/card/benefit/create"

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
	localVarPostBody = r.alipayMarketingCardBenefitCreateModel

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
			var v AlipayMarketingCardBenefitCreateDefaultResponse
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


type ApiAlipayMarketingCardBenefitDeleteRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingCardBenefitAPIService
	templateId *string
	benefitId *string
}

// 会员卡模板ID，通过 &lt;a  href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.template.create\&quot;&gt;alipay.marketing.card.template.create&lt;/a&gt;（会员卡模板创建)接口创建会员卡模板获取。
func (r ApiAlipayMarketingCardBenefitDeleteRequest) TemplateId(templateId string) ApiAlipayMarketingCardBenefitDeleteRequest {
	r.templateId = &templateId
	return r
}

// 权益ID，通过 &lt;a  href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create\&quot;&gt;alipay.marketing.card.benefit.create&lt;/a&gt;(会员卡模板外部权益创建)接口创建获取。
func (r ApiAlipayMarketingCardBenefitDeleteRequest) BenefitId(benefitId string) ApiAlipayMarketingCardBenefitDeleteRequest {
	r.benefitId = &benefitId
	return r
}

func (r ApiAlipayMarketingCardBenefitDeleteRequest) Execute() (*AlipayMarketingCardBenefitDeleteResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingCardBenefitDeleteExecute(r)
}

/*
AlipayMarketingCardBenefitDelete 会员卡模板外部权益删除

会员卡模板外部权益删除

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingCardBenefitDeleteRequest
*/
func (r *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitDelete(ctx context.Context) ApiAlipayMarketingCardBenefitDeleteRequest {
	return ApiAlipayMarketingCardBenefitDeleteRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingCardBenefitDeleteResponseModel
func (a *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitDeleteExecute(r ApiAlipayMarketingCardBenefitDeleteRequest) (*AlipayMarketingCardBenefitDeleteResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingCardBenefitDeleteResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingCardBenefitAPIService.AlipayMarketingCardBenefitDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/card/benefit/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.templateId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "template_id", r.templateId, "form", "")
	}
	if r.benefitId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "benefit_id", r.benefitId, "form", "")
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
			var v AlipayMarketingCardBenefitDeleteDefaultResponse
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


type ApiAlipayMarketingCardBenefitModifyRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingCardBenefitAPIService
	alipayMarketingCardBenefitModifyModel *AlipayMarketingCardBenefitModifyModel
}

func (r ApiAlipayMarketingCardBenefitModifyRequest) AlipayMarketingCardBenefitModifyModel(alipayMarketingCardBenefitModifyModel AlipayMarketingCardBenefitModifyModel) ApiAlipayMarketingCardBenefitModifyRequest {
	r.alipayMarketingCardBenefitModifyModel = &alipayMarketingCardBenefitModifyModel
	return r
}

func (r ApiAlipayMarketingCardBenefitModifyRequest) Execute() (*AlipayMarketingCardBenefitModifyResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingCardBenefitModifyExecute(r)
}

/*
AlipayMarketingCardBenefitModify 会员卡模板外部权益修改

会员卡模板外部权益修改

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingCardBenefitModifyRequest
*/
func (r *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitModify(ctx context.Context) ApiAlipayMarketingCardBenefitModifyRequest {
	return ApiAlipayMarketingCardBenefitModifyRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingCardBenefitModifyResponseModel
func (a *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitModifyExecute(r ApiAlipayMarketingCardBenefitModifyRequest) (*AlipayMarketingCardBenefitModifyResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingCardBenefitModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingCardBenefitAPIService.AlipayMarketingCardBenefitModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/card/benefit/modify"

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
	localVarPostBody = r.alipayMarketingCardBenefitModifyModel

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
			var v AlipayMarketingCardBenefitModifyDefaultResponse
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


type ApiAlipayMarketingCardBenefitQueryRequest struct {
	ctx context.Context
	ApiService *AlipayMarketingCardBenefitAPIService
	templateId *string
	benefitId *string
}

// 会员卡模板ID，通过 &lt;a  href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.template.create\&quot;&gt;alipay.marketing.card.template.create&lt;/a&gt;（会员卡模板创建)接口创建会员卡模板获取。
func (r ApiAlipayMarketingCardBenefitQueryRequest) TemplateId(templateId string) ApiAlipayMarketingCardBenefitQueryRequest {
	r.templateId = &templateId
	return r
}

// 权益ID，通过 &lt;a  href&#x3D;\&quot;https://opendocs.alipay.com/apis/api_5/alipay.marketing.card.benefit.create\&quot;&gt;alipay.marketing.card.benefit.create&lt;/a&gt;(会员卡模板外部权益创建)接口创建获取。
func (r ApiAlipayMarketingCardBenefitQueryRequest) BenefitId(benefitId string) ApiAlipayMarketingCardBenefitQueryRequest {
	r.benefitId = &benefitId
	return r
}

func (r ApiAlipayMarketingCardBenefitQueryRequest) Execute() (*AlipayMarketingCardBenefitQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMarketingCardBenefitQueryExecute(r)
}

/*
AlipayMarketingCardBenefitQuery 会员卡模板外部权益查询

会员卡模板外部权益查询

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMarketingCardBenefitQueryRequest
*/
func (r *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitQuery(ctx context.Context) ApiAlipayMarketingCardBenefitQueryRequest {
	return ApiAlipayMarketingCardBenefitQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMarketingCardBenefitQueryResponseModel
func (a *AlipayMarketingCardBenefitAPIService) AlipayMarketingCardBenefitQueryExecute(r ApiAlipayMarketingCardBenefitQueryRequest) (*AlipayMarketingCardBenefitQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMarketingCardBenefitQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMarketingCardBenefitAPIService.AlipayMarketingCardBenefitQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/marketing/card/benefit/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.templateId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "template_id", r.templateId, "form", "")
	}
	if r.benefitId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "benefit_id", r.benefitId, "form", "")
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
			var v AlipayMarketingCardBenefitQueryDefaultResponse
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


func (a *AlipayMarketingCardBenefitAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayMarketingCardBenefitAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


