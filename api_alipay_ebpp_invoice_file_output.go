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

// AlipayEbppInvoiceFileOutputAPIService AlipayEbppInvoiceFileOutputAPI service
type AlipayEbppInvoiceFileOutputAPIService service

type ApiAlipayEbppInvoiceFileOutputQueryRequest struct {
	ctx                     context.Context
	ApiService              *AlipayEbppInvoiceFileOutputAPIService
	userId                  *string
	openId                  *string
	invoiceCode             *string
	invoiceNo               *string
	scene                   *string
	skipExpenseProgressSync *bool
}

// 发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) UserId(userId string) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.userId = &userId
	return r
}

// 发票归属用户 id，用户在支付宝的唯一标识，以 2088 开头的 16 位纯数字组成。
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) OpenId(openId string) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.openId = &openId
	return r
}

// 发票代码 字段长度（10-12位），全电票时为空
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) InvoiceCode(invoiceCode string) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.invoiceCode = &invoiceCode
	return r
}

// 发票号码 字段长度（8-10位），全电票时为20位
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) InvoiceNo(invoiceNo string) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.invoiceNo = &invoiceNo
	return r
}

// 发票pdf文件下载应用场景。固定为 INVOICE_EXPENSE  表示发票报销。
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) Scene(scene string) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.scene = &scene
	return r
}

// 是否跳过发票报销状态同步；当为true时，跳过报销状态同步校验。默认为false，需要先做报销状态同步
func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) SkipExpenseProgressSync(skipExpenseProgressSync bool) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	r.skipExpenseProgressSync = &skipExpenseProgressSync
	return r
}

func (r ApiAlipayEbppInvoiceFileOutputQueryRequest) Execute() (*AlipayEbppInvoiceFileOutputQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEbppInvoiceFileOutputQueryExecute(r)
}

/*
AlipayEbppInvoiceFileOutputQuery 获取电子发票原文件

根据发票代码，发票号码获取用户的普通增值税电子发票pdf文件
查询权限要求： 发票属于当前isv报销单据进行状态（支持的状态有 EXPENSE_APPLY－用户已提交申请 EXPENSE_APPROVAL_PASS －报销审核通过 EXPENSE_FINISHED－报销完结）

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayEbppInvoiceFileOutputQueryRequest
*/
func (r *AlipayEbppInvoiceFileOutputAPIService) AlipayEbppInvoiceFileOutputQuery(ctx context.Context) ApiAlipayEbppInvoiceFileOutputQueryRequest {
	return ApiAlipayEbppInvoiceFileOutputQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayEbppInvoiceFileOutputQueryResponseModel
func (a *AlipayEbppInvoiceFileOutputAPIService) AlipayEbppInvoiceFileOutputQueryExecute(r ApiAlipayEbppInvoiceFileOutputQueryRequest) (*AlipayEbppInvoiceFileOutputQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayEbppInvoiceFileOutputQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEbppInvoiceFileOutputAPIService.AlipayEbppInvoiceFileOutputQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/ebpp/invoice/file/output/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	}
	if r.openId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "open_id", r.openId, "form", "")
	}
	if r.invoiceCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoice_code", r.invoiceCode, "form", "")
	}
	if r.invoiceNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoice_no", r.invoiceNo, "form", "")
	}
	if r.scene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scene", r.scene, "form", "")
	}
	if r.skipExpenseProgressSync != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip_expense_progress_sync", r.skipExpenseProgressSync, "form", "")
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
		var v AlipayEbppInvoiceFileOutputQueryDefaultResponse
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

func (a *AlipayEbppInvoiceFileOutputAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEbppInvoiceFileOutputAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
