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


// AlipayEbppInvoiceTaxnoAPIService AlipayEbppInvoiceTaxnoAPI service
type AlipayEbppInvoiceTaxnoAPIService service

type ApiAlipayEbppInvoiceTaxnoBatchqueryRequest struct {
	ctx context.Context
	ApiService *AlipayEbppInvoiceTaxnoAPIService
	authToken *string
	alipayEbppInvoiceTaxnoBatchqueryModel *AlipayEbppInvoiceTaxnoBatchqueryModel
}

// 用户授权令牌
func (r ApiAlipayEbppInvoiceTaxnoBatchqueryRequest) AuthToken(authToken string) ApiAlipayEbppInvoiceTaxnoBatchqueryRequest {
	r.authToken = &authToken
	return r
}

func (r ApiAlipayEbppInvoiceTaxnoBatchqueryRequest) AlipayEbppInvoiceTaxnoBatchqueryModel(alipayEbppInvoiceTaxnoBatchqueryModel AlipayEbppInvoiceTaxnoBatchqueryModel) ApiAlipayEbppInvoiceTaxnoBatchqueryRequest {
	r.alipayEbppInvoiceTaxnoBatchqueryModel = &alipayEbppInvoiceTaxnoBatchqueryModel
	return r
}

func (r ApiAlipayEbppInvoiceTaxnoBatchqueryRequest) Execute() (*AlipayEbppInvoiceTaxnoBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceTaxnoBatchqueryExecute(r)
}

/*
AlipayEbppInvoiceTaxnoBatchquery 获取指定企业的发票要素列表

根据购买方税号获取用户的发票关键要素列表，该接口目前适用于发票报销场景

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayEbppInvoiceTaxnoBatchqueryRequest
*/
func (r *AlipayEbppInvoiceTaxnoAPIService) AlipayEbppInvoiceTaxnoBatchquery(ctx context.Context) ApiAlipayEbppInvoiceTaxnoBatchqueryRequest {
	return ApiAlipayEbppInvoiceTaxnoBatchqueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayEbppInvoiceTaxnoBatchqueryResponseModel
func (a *AlipayEbppInvoiceTaxnoAPIService) AlipayEbppInvoiceTaxnoBatchqueryExecute(r ApiAlipayEbppInvoiceTaxnoBatchqueryRequest) (*AlipayEbppInvoiceTaxnoBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayEbppInvoiceTaxnoBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceTaxnoAPIService.AlipayEbppInvoiceTaxnoBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/taxno/batchquery"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.authToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auth_token", r.authToken, "form", "")
	}

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
	localVarPostBody = r.alipayEbppInvoiceTaxnoBatchqueryModel

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
			var v AlipayEbppInvoiceTaxnoBatchqueryDefaultResponse
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


func (a *AlipayEbppInvoiceTaxnoAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceTaxnoAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

