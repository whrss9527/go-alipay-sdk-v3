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

// AlipayOpenAppApiFieldAPIService AlipayOpenAppApiFieldAPI service
type AlipayOpenAppApiFieldAPIService service

type ApiAlipayOpenAppApiFieldApplyRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenAppApiFieldAPIService
	data       *AlipayOpenAppApiFieldApplyModel
	picture1   *os.File
	picture2   *os.File
	picture3   *os.File
	picture4   *os.File
	picture5   *os.File
	video      *os.File
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Data(data AlipayOpenAppApiFieldApplyModel) ApiAlipayOpenAppApiFieldApplyRequest {
	r.data = &data
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Picture1(picture1 *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.picture1 = picture1
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Picture2(picture2 *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.picture2 = picture2
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Picture3(picture3 *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.picture3 = picture3
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Picture4(picture4 *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.picture4 = picture4
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Picture5(picture5 *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.picture5 = picture5
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Video(video *os.File) ApiAlipayOpenAppApiFieldApplyRequest {
	r.video = video
	return r
}

func (r ApiAlipayOpenAppApiFieldApplyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenAppApiFieldApplyExecute(r)
}

/*
AlipayOpenAppApiFieldApply 申请获取接口用户敏感信息字段

申请获取接口出参中用户敏感信息字段，应用使用视频或图片（至少选择一项上传），单个图片或视频不超过10M，最多可上传5张图片示例。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenAppApiFieldApplyRequest
*/
func (r *AlipayOpenAppApiFieldAPIService) AlipayOpenAppApiFieldApply(ctx context.Context) ApiAlipayOpenAppApiFieldApplyRequest {
	return ApiAlipayOpenAppApiFieldApplyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayOpenAppApiFieldAPIService) AlipayOpenAppApiFieldApplyExecute(r ApiAlipayOpenAppApiFieldApplyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenAppApiFieldAPIService.AlipayOpenAppApiFieldApply")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/app/api/field/apply"

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
	if r.data != nil {
		paramJson, err := parameterToJson(*r.data)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("data", paramJson)
	}
	var picture1LocalVarFormFileName string
	var picture1LocalVarFileName string
	var picture1LocalVarFileBytes []byte

	picture1LocalVarFormFileName = "picture_1"
	picture1LocalVarFile := r.picture1

	if picture1LocalVarFile != nil {
		fbs, _ := io.ReadAll(picture1LocalVarFile)

		picture1LocalVarFileBytes = fbs
		picture1LocalVarFileName = picture1LocalVarFile.Name()
		picture1LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: picture1LocalVarFileBytes, fileName: picture1LocalVarFileName, formFileName: picture1LocalVarFormFileName})
	}
	var picture2LocalVarFormFileName string
	var picture2LocalVarFileName string
	var picture2LocalVarFileBytes []byte

	picture2LocalVarFormFileName = "picture_2"
	picture2LocalVarFile := r.picture2

	if picture2LocalVarFile != nil {
		fbs, _ := io.ReadAll(picture2LocalVarFile)

		picture2LocalVarFileBytes = fbs
		picture2LocalVarFileName = picture2LocalVarFile.Name()
		picture2LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: picture2LocalVarFileBytes, fileName: picture2LocalVarFileName, formFileName: picture2LocalVarFormFileName})
	}
	var picture3LocalVarFormFileName string
	var picture3LocalVarFileName string
	var picture3LocalVarFileBytes []byte

	picture3LocalVarFormFileName = "picture_3"
	picture3LocalVarFile := r.picture3

	if picture3LocalVarFile != nil {
		fbs, _ := io.ReadAll(picture3LocalVarFile)

		picture3LocalVarFileBytes = fbs
		picture3LocalVarFileName = picture3LocalVarFile.Name()
		picture3LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: picture3LocalVarFileBytes, fileName: picture3LocalVarFileName, formFileName: picture3LocalVarFormFileName})
	}
	var picture4LocalVarFormFileName string
	var picture4LocalVarFileName string
	var picture4LocalVarFileBytes []byte

	picture4LocalVarFormFileName = "picture_4"
	picture4LocalVarFile := r.picture4

	if picture4LocalVarFile != nil {
		fbs, _ := io.ReadAll(picture4LocalVarFile)

		picture4LocalVarFileBytes = fbs
		picture4LocalVarFileName = picture4LocalVarFile.Name()
		picture4LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: picture4LocalVarFileBytes, fileName: picture4LocalVarFileName, formFileName: picture4LocalVarFormFileName})
	}
	var picture5LocalVarFormFileName string
	var picture5LocalVarFileName string
	var picture5LocalVarFileBytes []byte

	picture5LocalVarFormFileName = "picture_5"
	picture5LocalVarFile := r.picture5

	if picture5LocalVarFile != nil {
		fbs, _ := io.ReadAll(picture5LocalVarFile)

		picture5LocalVarFileBytes = fbs
		picture5LocalVarFileName = picture5LocalVarFile.Name()
		picture5LocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: picture5LocalVarFileBytes, fileName: picture5LocalVarFileName, formFileName: picture5LocalVarFormFileName})
	}
	var videoLocalVarFormFileName string
	var videoLocalVarFileName string
	var videoLocalVarFileBytes []byte

	videoLocalVarFormFileName = "video"
	videoLocalVarFile := r.video

	if videoLocalVarFile != nil {
		fbs, _ := io.ReadAll(videoLocalVarFile)

		videoLocalVarFileBytes = fbs
		videoLocalVarFileName = videoLocalVarFile.Name()
		videoLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: videoLocalVarFileBytes, fileName: videoLocalVarFileName, formFileName: videoLocalVarFormFileName})
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
		var v AlipayOpenAppApiFieldApplyDefaultResponse
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

type ApiAlipayOpenAppApiFieldQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenAppApiFieldAPIService
}

func (r ApiAlipayOpenAppApiFieldQueryRequest) Execute() (*AlipayOpenAppApiFieldQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenAppApiFieldQueryExecute(r)
}

/*
AlipayOpenAppApiFieldQuery 用户信息申请记录查询

查询用户信息申请记录，开发者发起请求时会基于当前发起请求的app_id进行查询，若为isv发起的三方代调用则以授权的商户app_id进行查询。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenAppApiFieldQueryRequest
*/
func (r *AlipayOpenAppApiFieldAPIService) AlipayOpenAppApiFieldQuery(ctx context.Context) ApiAlipayOpenAppApiFieldQueryRequest {
	return ApiAlipayOpenAppApiFieldQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenAppApiFieldQueryResponseModel
func (a *AlipayOpenAppApiFieldAPIService) AlipayOpenAppApiFieldQueryExecute(r ApiAlipayOpenAppApiFieldQueryRequest) (*AlipayOpenAppApiFieldQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenAppApiFieldQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenAppApiFieldAPIService.AlipayOpenAppApiFieldQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/app/api/field/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		var v AlipayOpenAppApiFieldQueryDefaultResponse
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

func (a *AlipayOpenAppApiFieldAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenAppApiFieldAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
