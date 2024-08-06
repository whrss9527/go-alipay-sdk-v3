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
	"os"
	"strconv"
	"time"
)

// AlipayOpenPublicLifeMsgAPIService AlipayOpenPublicLifeMsgAPI service
type AlipayOpenPublicLifeMsgAPIService service

type ApiAlipayOpenPublicLifeMsgRecallRequest struct {
	ctx                                context.Context
	ApiService                         *AlipayOpenPublicLifeMsgAPIService
	alipayOpenPublicLifeMsgRecallModel *AlipayOpenPublicLifeMsgRecallModel
}

func (r ApiAlipayOpenPublicLifeMsgRecallRequest) AlipayOpenPublicLifeMsgRecallModel(alipayOpenPublicLifeMsgRecallModel AlipayOpenPublicLifeMsgRecallModel) ApiAlipayOpenPublicLifeMsgRecallRequest {
	r.alipayOpenPublicLifeMsgRecallModel = &alipayOpenPublicLifeMsgRecallModel
	return r
}

func (r ApiAlipayOpenPublicLifeMsgRecallRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicLifeMsgRecallExecute(r)
}

/*
AlipayOpenPublicLifeMsgRecall 生活号消息撤回接口

生活号消息撤回接口

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenPublicLifeMsgRecallRequest
*/
func (r *AlipayOpenPublicLifeMsgAPIService) AlipayOpenPublicLifeMsgRecall(ctx context.Context) ApiAlipayOpenPublicLifeMsgRecallRequest {
	return ApiAlipayOpenPublicLifeMsgRecallRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayOpenPublicLifeMsgAPIService) AlipayOpenPublicLifeMsgRecallExecute(r ApiAlipayOpenPublicLifeMsgRecallRequest) (map[string]interface{}, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicLifeMsgAPIService.AlipayOpenPublicLifeMsgRecall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/life/msg/recall"

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
	localVarPostBody = r.alipayOpenPublicLifeMsgRecallModel

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
		var v AlipayOpenPublicLifeMsgRecallDefaultResponse
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

type ApiAlipayOpenPublicLifeMsgSendRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenPublicLifeMsgAPIService
	cover      *os.File
	data       *AlipayOpenPublicLifeMsgSendModel
}

func (r ApiAlipayOpenPublicLifeMsgSendRequest) Cover(cover *os.File) ApiAlipayOpenPublicLifeMsgSendRequest {
	r.cover = cover
	return r
}

func (r ApiAlipayOpenPublicLifeMsgSendRequest) Data(data AlipayOpenPublicLifeMsgSendModel) ApiAlipayOpenPublicLifeMsgSendRequest {
	r.data = &data
	return r
}

func (r ApiAlipayOpenPublicLifeMsgSendRequest) Execute() (*AlipayOpenPublicLifeMsgSendResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenPublicLifeMsgSendExecute(r)
}

/*
AlipayOpenPublicLifeMsgSend 生活号广播消息发送接口

媒体资讯生活号群发一篇图文或视频类消息，生活号主页以及客户端首页展示消息详情

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenPublicLifeMsgSendRequest
*/
func (r *AlipayOpenPublicLifeMsgAPIService) AlipayOpenPublicLifeMsgSend(ctx context.Context) ApiAlipayOpenPublicLifeMsgSendRequest {
	return ApiAlipayOpenPublicLifeMsgSendRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenPublicLifeMsgSendResponseModel
func (a *AlipayOpenPublicLifeMsgAPIService) AlipayOpenPublicLifeMsgSendExecute(r ApiAlipayOpenPublicLifeMsgSendRequest) (*AlipayOpenPublicLifeMsgSendResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenPublicLifeMsgSendResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenPublicLifeMsgAPIService.AlipayOpenPublicLifeMsgSend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/public/life/msg/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var coverLocalVarFormFileName string
	var coverLocalVarFileName string
	var coverLocalVarFileBytes []byte

	coverLocalVarFormFileName = "cover"
	coverLocalVarFile := r.cover

	if coverLocalVarFile != nil {
		fbs, _ := io.ReadAll(coverLocalVarFile)

		coverLocalVarFileBytes = fbs
		coverLocalVarFileName = coverLocalVarFile.Name()
		coverLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: coverLocalVarFileBytes, fileName: coverLocalVarFileName, formFileName: coverLocalVarFormFileName})
	}
	if r.data != nil {
		paramJson, err := parameterToJson(*r.data)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("data", paramJson)
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
		var v AlipayOpenPublicLifeMsgSendDefaultResponse
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

func (a *AlipayOpenPublicLifeMsgAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenPublicLifeMsgAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
