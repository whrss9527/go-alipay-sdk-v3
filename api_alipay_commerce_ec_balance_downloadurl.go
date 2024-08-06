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

// AlipayCommerceEcBalanceDownloadurlAPIService AlipayCommerceEcBalanceDownloadurlAPI service
type AlipayCommerceEcBalanceDownloadurlAPIService service

type ApiAlipayCommerceEcBalanceDownloadurlQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayCommerceEcBalanceDownloadurlAPIService
	enterpriseId *string
	accountId    *string
	agreementNo  *string
	billType     *string
	billDate     *string
}

// 企业ID
func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) EnterpriseId(enterpriseId string) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	r.enterpriseId = &enterpriseId
	return r
}

// 共同账户ID
func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) AccountId(accountId string) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	r.accountId = &accountId
	return r
}

// 授权签约协议号
func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) AgreementNo(agreementNo string) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 账单类型
func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) BillType(billType string) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	r.billType = &billType
	return r
}

// 账单时间
func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) BillDate(billDate string) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	r.billDate = &billDate
	return r
}

func (r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) Execute() (*AlipayCommerceEcBalanceDownloadurlQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcBalanceDownloadurlQueryExecute(r)
}

/*
AlipayCommerceEcBalanceDownloadurlQuery 对账单文件下载接口

对账单文件下载接口

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceEcBalanceDownloadurlQueryRequest
*/
func (r *AlipayCommerceEcBalanceDownloadurlAPIService) AlipayCommerceEcBalanceDownloadurlQuery(ctx context.Context) ApiAlipayCommerceEcBalanceDownloadurlQueryRequest {
	return ApiAlipayCommerceEcBalanceDownloadurlQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceEcBalanceDownloadurlQueryResponseModel
func (a *AlipayCommerceEcBalanceDownloadurlAPIService) AlipayCommerceEcBalanceDownloadurlQueryExecute(r ApiAlipayCommerceEcBalanceDownloadurlQueryRequest) (*AlipayCommerceEcBalanceDownloadurlQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceEcBalanceDownloadurlQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcBalanceDownloadurlAPIService.AlipayCommerceEcBalanceDownloadurlQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/balance/downloadurl/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.enterpriseId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enterprise_id", r.enterpriseId, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.billType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bill_type", r.billType, "form", "")
	}
	if r.billDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "bill_date", r.billDate, "form", "")
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
		var v AlipayCommerceEcBalanceDownloadurlQueryDefaultResponse
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

func (a *AlipayCommerceEcBalanceDownloadurlAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceEcBalanceDownloadurlAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
