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

// AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService AlipayEbppInvoiceExpensecontrolEmployeequotaAPI service
type AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService service

type ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest struct {
	ctx                                                     context.Context
	ApiService                                              *AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService
	alipayEbppInvoiceExpensecontrolEmployeequotaCreateModel *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel
}

func (r ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest) AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel(alipayEbppInvoiceExpensecontrolEmployeequotaCreateModel AlipayEbppInvoiceExpensecontrolEmployeequotaCreateModel) ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest {
	r.alipayEbppInvoiceExpensecontrolEmployeequotaCreateModel = &alipayEbppInvoiceExpensecontrolEmployeequotaCreateModel
	return r
}

func (r ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest) Execute() (*AlipayEbppInvoiceExpensecontrolEmployeequotaCreateResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceExpensecontrolEmployeequotaCreateExecute(r)
}

/*
AlipayEbppInvoiceExpensecontrolEmployeequotaCreate 创建余额（特定场景）

创建余额，该接口开放给特定场景。如果需要使用，请联系技术支持。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest
*/
func (r *AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService) AlipayEbppInvoiceExpensecontrolEmployeequotaCreate(ctx context.Context) ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest {
	return ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceExpensecontrolEmployeequotaCreateResponseModel
func (a *AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService) AlipayEbppInvoiceExpensecontrolEmployeequotaCreateExecute(r ApiAlipayEbppInvoiceExpensecontrolEmployeequotaCreateRequest) (*AlipayEbppInvoiceExpensecontrolEmployeequotaCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceExpensecontrolEmployeequotaCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService.AlipayEbppInvoiceExpensecontrolEmployeequotaCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/expensecontrol/employeequota/create"

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
	localVarPostBody = r.alipayEbppInvoiceExpensecontrolEmployeequotaCreateModel

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
		var v AlipayEbppInvoiceExpensecontrolEmployeequotaCreateDefaultResponse
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

func (a *AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceExpensecontrolEmployeequotaAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
