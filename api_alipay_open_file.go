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

// AlipayOpenFileAPIService AlipayOpenFileAPI service
type AlipayOpenFileAPIService service

type ApiAlipayOpenFileUploadRequest struct {
	ctx         context.Context
	ApiService  *AlipayOpenFileAPIService
	data        *AlipayOpenFileUploadModel
	fileContent *os.File
}

func (r ApiAlipayOpenFileUploadRequest) Data(data AlipayOpenFileUploadModel) ApiAlipayOpenFileUploadRequest {
	r.data = &data
	return r
}

func (r ApiAlipayOpenFileUploadRequest) FileContent(fileContent *os.File) ApiAlipayOpenFileUploadRequest {
	r.fileContent = fileContent
	return r
}

func (r ApiAlipayOpenFileUploadRequest) Execute() (*AlipayOpenFileUploadResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenFileUploadExecute(r)
}

/*
AlipayOpenFileUpload 支付宝文件上传接口

支付宝通用文件上传接口，支付宝开发者可以调用此接口上传文件，文件可以提供给其他业务使用，例如：上传素材（视频/图片），素材可以供给到生活号+/小程序等其他开放应用后正式发布。
该接口支持使用支付宝SDK和HTTP POST请求，两种方式示例如下：
1. 使用支付宝SDK时，可以直接使用FileItem上传，代码片段参考下文
2. 使用直接使用http协议时，需要忽略file_content参数，文件放入表单中，其余参数放入URL中，其中将参数（不包括文件和file_content）按<a href="https://opendocs.alipay.com/common/02kf5q#%E8%87%AA%E8%A1%8C%E5%AE%9E%E7%8E%B0%E7%AD%BE%E5%90%8D">自行实现签名</a>文档实现，示例：
curl --request POST 'http://openapi.alipay.com/gateway.do?charset=GBK&biz_content=%7B%22biz_code%22%3A%22content_creation%22%2C%22extra_parameters%22%3A%7B%22extern_upload%22%3A%22youku%22%2C%22alipay_user_id%22%3A%22xxxx%22%7D%7D&method=alipay.open.file.upload&format=json&sign=xxx&app_id=2014060600164699&version=1.0&sign_type=RSA2&timestamp=xxx'  --form 'file=@"xx.mp4"'

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenFileUploadRequest
*/
func (r *AlipayOpenFileAPIService) AlipayOpenFileUpload(ctx context.Context) ApiAlipayOpenFileUploadRequest {
	return ApiAlipayOpenFileUploadRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenFileUploadResponseModel
func (a *AlipayOpenFileAPIService) AlipayOpenFileUploadExecute(r ApiAlipayOpenFileUploadRequest) (*AlipayOpenFileUploadResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenFileUploadResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenFileAPIService.AlipayOpenFileUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/file/upload"

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
	var fileContentLocalVarFormFileName string
	var fileContentLocalVarFileName string
	var fileContentLocalVarFileBytes []byte

	fileContentLocalVarFormFileName = "file_content"
	fileContentLocalVarFile := r.fileContent

	if fileContentLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileContentLocalVarFile)

		fileContentLocalVarFileBytes = fbs
		fileContentLocalVarFileName = fileContentLocalVarFile.Name()
		fileContentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileContentLocalVarFileBytes, fileName: fileContentLocalVarFileName, formFileName: fileContentLocalVarFormFileName})
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
		var v AlipayOpenFileUploadDefaultResponse
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

func (a *AlipayOpenFileAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenFileAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
