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

// AlipayOpenPublicMatchuserLabelAPIService AlipayOpenPublicMatchuserLabelAPI service
type AlipayOpenPublicMatchuserLabelAPIService service

type ApiAlipayOpenPublicMatchuserLabelCreateRequest struct {
	ctx                                       context.Context
	ApiService                                *AlipayOpenPublicMatchuserLabelAPIService
	alipayOpenPublicMatchuserLabelCreateModel *AlipayOpenPublicMatchuserLabelCreateModel
}

func (r ApiAlipayOpenPublicMatchuserLabelCreateRequest) AlipayOpenPublicMatchuserLabelCreateModel(alipayOpenPublicMatchuserLabelCreateModel AlipayOpenPublicMatchuserLabelCreateModel) ApiAlipayOpenPublicMatchuserLabelCreateRequest {
	r.alipayOpenPublicMatchuserLabelCreateModel = &alipayOpenPublicMatchuserLabelCreateModel
	return r
}

func (r ApiAlipayOpenPublicMatchuserLabelCreateRequest) Execute() (*AlipayOpenPublicMatchuserLabelCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicMatchuserLabelCreateExecute(r)
}

/*
AlipayOpenPublicMatchuserLabelCreate 用户打标接口

开发者可以使用此接口向匹配到的用户添加指定标签值，每个用户每个标签只能有一个标签值，当向同一个用户添加多个标签值时，最新的标签值会覆盖之前的标签值

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenPublicMatchuserLabelCreateRequest
*/
func (r *AlipayOpenPublicMatchuserLabelAPIService) AlipayOpenPublicMatchuserLabelCreate(ctx context.Context) ApiAlipayOpenPublicMatchuserLabelCreateRequest {
	return ApiAlipayOpenPublicMatchuserLabelCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenPublicMatchuserLabelCreateResponseModel
func (a *AlipayOpenPublicMatchuserLabelAPIService) AlipayOpenPublicMatchuserLabelCreateExecute(r ApiAlipayOpenPublicMatchuserLabelCreateRequest) (*AlipayOpenPublicMatchuserLabelCreateResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenPublicMatchuserLabelCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicMatchuserLabelAPIService.AlipayOpenPublicMatchuserLabelCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/matchuser/label/create"

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
	localVarPostBody = r.alipayOpenPublicMatchuserLabelCreateModel

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
		var v AlipayOpenPublicMatchuserLabelCreateDefaultResponse
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

type ApiAlipayOpenPublicMatchuserLabelDeleteRequest struct {
	ctx                                       context.Context
	ApiService                                *AlipayOpenPublicMatchuserLabelAPIService
	alipayOpenPublicMatchuserLabelDeleteModel *AlipayOpenPublicMatchuserLabelDeleteModel
}

func (r ApiAlipayOpenPublicMatchuserLabelDeleteRequest) AlipayOpenPublicMatchuserLabelDeleteModel(alipayOpenPublicMatchuserLabelDeleteModel AlipayOpenPublicMatchuserLabelDeleteModel) ApiAlipayOpenPublicMatchuserLabelDeleteRequest {
	r.alipayOpenPublicMatchuserLabelDeleteModel = &alipayOpenPublicMatchuserLabelDeleteModel
	return r
}

func (r ApiAlipayOpenPublicMatchuserLabelDeleteRequest) Execute() (*AlipayOpenPublicMatchuserLabelDeleteResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicMatchuserLabelDeleteExecute(r)
}

/*
AlipayOpenPublicMatchuserLabelDelete 用户取消标签接口

开发者可以使用此接口删除匹配到的支付宝用户的指定标签值

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenPublicMatchuserLabelDeleteRequest
*/
func (r *AlipayOpenPublicMatchuserLabelAPIService) AlipayOpenPublicMatchuserLabelDelete(ctx context.Context) ApiAlipayOpenPublicMatchuserLabelDeleteRequest {
	return ApiAlipayOpenPublicMatchuserLabelDeleteRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenPublicMatchuserLabelDeleteResponseModel
func (a *AlipayOpenPublicMatchuserLabelAPIService) AlipayOpenPublicMatchuserLabelDeleteExecute(r ApiAlipayOpenPublicMatchuserLabelDeleteRequest) (*AlipayOpenPublicMatchuserLabelDeleteResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenPublicMatchuserLabelDeleteResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicMatchuserLabelAPIService.AlipayOpenPublicMatchuserLabelDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/matchuser/label/delete"

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
	localVarPostBody = r.alipayOpenPublicMatchuserLabelDeleteModel

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
		var v AlipayOpenPublicMatchuserLabelDeleteDefaultResponse
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

func (a *AlipayOpenPublicMatchuserLabelAPIService) signRequest(req *http.Request) error {
	appID := a.client.cfg.AppID
	appCertSN := a.client.cfg.AppCertSN
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

	signature, err := signWithRSA(content, a.client.cfg.privateKey)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("ALIPAY-SHA256withRSA %s,sign=%s", authString, signature))
	return nil
}

func (a *AlipayOpenPublicMatchuserLabelAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
