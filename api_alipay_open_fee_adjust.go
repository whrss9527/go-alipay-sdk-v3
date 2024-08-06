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

// AlipayOpenFeeAdjustAPIService AlipayOpenFeeAdjustAPI service
type AlipayOpenFeeAdjustAPIService service

type ApiAlipayOpenFeeAdjustApplyRequest struct {
	ctx          context.Context
	ApiService   *AlipayOpenFeeAdjustAPIService
	attachment   *os.File
	certPic      *os.File
	data         *AlipayOpenFeeAdjustApplyModel
	shopScenePic *os.File
	shopSignPic  *os.File
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) Attachment(attachment *os.File) ApiAlipayOpenFeeAdjustApplyRequest {
	r.attachment = attachment
	return r
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) CertPic(certPic *os.File) ApiAlipayOpenFeeAdjustApplyRequest {
	r.certPic = certPic
	return r
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) Data(data AlipayOpenFeeAdjustApplyModel) ApiAlipayOpenFeeAdjustApplyRequest {
	r.data = &data
	return r
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) ShopScenePic(shopScenePic *os.File) ApiAlipayOpenFeeAdjustApplyRequest {
	r.shopScenePic = shopScenePic
	return r
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) ShopSignPic(shopSignPic *os.File) ApiAlipayOpenFeeAdjustApplyRequest {
	r.shopSignPic = shopSignPic
	return r
}

func (r ApiAlipayOpenFeeAdjustApplyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayOpenFeeAdjustApplyExecute(r)
}

/*
AlipayOpenFeeAdjustApply 特殊费率申请

服务商替代签约或交易传参的商户申请费率

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenFeeAdjustApplyRequest
*/
func (r *AlipayOpenFeeAdjustAPIService) AlipayOpenFeeAdjustApply(ctx context.Context) ApiAlipayOpenFeeAdjustApplyRequest {
	return ApiAlipayOpenFeeAdjustApplyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *AlipayOpenFeeAdjustAPIService) AlipayOpenFeeAdjustApplyExecute(r ApiAlipayOpenFeeAdjustApplyRequest) (map[string]interface{}, *http.Response, error) {
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

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenFeeAdjustAPIService.AlipayOpenFeeAdjustApply")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/fee/adjust/apply"

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
	var attachmentLocalVarFormFileName string
	var attachmentLocalVarFileName string
	var attachmentLocalVarFileBytes []byte

	attachmentLocalVarFormFileName = "attachment"
	attachmentLocalVarFile := r.attachment

	if attachmentLocalVarFile != nil {
		fbs, _ := io.ReadAll(attachmentLocalVarFile)

		attachmentLocalVarFileBytes = fbs
		attachmentLocalVarFileName = attachmentLocalVarFile.Name()
		attachmentLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: attachmentLocalVarFileBytes, fileName: attachmentLocalVarFileName, formFileName: attachmentLocalVarFormFileName})
	}
	var certPicLocalVarFormFileName string
	var certPicLocalVarFileName string
	var certPicLocalVarFileBytes []byte

	certPicLocalVarFormFileName = "cert_pic"
	certPicLocalVarFile := r.certPic

	if certPicLocalVarFile != nil {
		fbs, _ := io.ReadAll(certPicLocalVarFile)

		certPicLocalVarFileBytes = fbs
		certPicLocalVarFileName = certPicLocalVarFile.Name()
		certPicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: certPicLocalVarFileBytes, fileName: certPicLocalVarFileName, formFileName: certPicLocalVarFormFileName})
	}
	if r.data != nil {
		paramJson, err := parameterToJson(*r.data)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("data", paramJson)
	}
	var shopScenePicLocalVarFormFileName string
	var shopScenePicLocalVarFileName string
	var shopScenePicLocalVarFileBytes []byte

	shopScenePicLocalVarFormFileName = "shop_scene_pic"
	shopScenePicLocalVarFile := r.shopScenePic

	if shopScenePicLocalVarFile != nil {
		fbs, _ := io.ReadAll(shopScenePicLocalVarFile)

		shopScenePicLocalVarFileBytes = fbs
		shopScenePicLocalVarFileName = shopScenePicLocalVarFile.Name()
		shopScenePicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: shopScenePicLocalVarFileBytes, fileName: shopScenePicLocalVarFileName, formFileName: shopScenePicLocalVarFormFileName})
	}
	var shopSignPicLocalVarFormFileName string
	var shopSignPicLocalVarFileName string
	var shopSignPicLocalVarFileBytes []byte

	shopSignPicLocalVarFormFileName = "shop_sign_pic"
	shopSignPicLocalVarFile := r.shopSignPic

	if shopSignPicLocalVarFile != nil {
		fbs, _ := io.ReadAll(shopSignPicLocalVarFile)

		shopSignPicLocalVarFileBytes = fbs
		shopSignPicLocalVarFileName = shopSignPicLocalVarFile.Name()
		shopSignPicLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: shopSignPicLocalVarFileBytes, fileName: shopSignPicLocalVarFileName, formFileName: shopSignPicLocalVarFormFileName})
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
		var v AlipayOpenFeeAdjustApplyDefaultResponse
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

func (a *AlipayOpenFeeAdjustAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenFeeAdjustAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
