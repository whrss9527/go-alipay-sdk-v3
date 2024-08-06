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


// AlipayDataBillBuyAPIService AlipayDataBillBuyAPI service
type AlipayDataBillBuyAPIService service

type ApiAlipayDataBillBuyQueryRequest struct {
	ctx context.Context
	ApiService *AlipayDataBillBuyAPIService
	startTime *string
	endTime *string
	alipayOrderNo *string
	merchantOrderNo *string
	storeNo *string
	pageNo *string
	pageSize *string
}

// 交易流水创建时间的起始范围
func (r ApiAlipayDataBillBuyQueryRequest) StartTime(startTime string) ApiAlipayDataBillBuyQueryRequest {
	r.startTime = &startTime
	return r
}

// 交易流水创建时间的结束范围。与起始时间间隔不超过31天。查询结果为起始时间至结束时间的左闭右开区间
func (r ApiAlipayDataBillBuyQueryRequest) EndTime(endTime string) ApiAlipayDataBillBuyQueryRequest {
	r.endTime = &endTime
	return r
}

// 支付宝交易流水号。如果查询参数中指定流水号，则只查询流水号相关的记录
func (r ApiAlipayDataBillBuyQueryRequest) AlipayOrderNo(alipayOrderNo string) ApiAlipayDataBillBuyQueryRequest {
	r.alipayOrderNo = &alipayOrderNo
	return r
}

// 商户交易号。如果查询参数中指定交易号，则只查询相关的记录
func (r ApiAlipayDataBillBuyQueryRequest) MerchantOrderNo(merchantOrderNo string) ApiAlipayDataBillBuyQueryRequest {
	r.merchantOrderNo = &merchantOrderNo
	return r
}

// 门店编号，模糊搜索
func (r ApiAlipayDataBillBuyQueryRequest) StoreNo(storeNo string) ApiAlipayDataBillBuyQueryRequest {
	r.storeNo = &storeNo
	return r
}

// 分页号，从1开始
func (r ApiAlipayDataBillBuyQueryRequest) PageNo(pageNo string) ApiAlipayDataBillBuyQueryRequest {
	r.pageNo = &pageNo
	return r
}

// 分页大小1000-2000，默认2000
func (r ApiAlipayDataBillBuyQueryRequest) PageSize(pageSize string) ApiAlipayDataBillBuyQueryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiAlipayDataBillBuyQueryRequest) Execute() (*AlipayDataBillBuyQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayDataBillBuyQueryExecute(r)
}

/*
AlipayDataBillBuyQuery 支付宝商家账户买入交易查询

为支付宝商家提供支付宝账户买入交易信息，供对账使用。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayDataBillBuyQueryRequest
*/
func (r *AlipayDataBillBuyAPIService) AlipayDataBillBuyQuery(ctx context.Context) ApiAlipayDataBillBuyQueryRequest {
	return ApiAlipayDataBillBuyQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayDataBillBuyQueryResponseModel
func (a *AlipayDataBillBuyAPIService) AlipayDataBillBuyQueryExecute(r ApiAlipayDataBillBuyQueryRequest) (*AlipayDataBillBuyQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayDataBillBuyQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayDataBillBuyAPIService.AlipayDataBillBuyQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/data/bill/buy/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime, "form", "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime, "form", "")
	}
	if r.alipayOrderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "alipay_order_no", r.alipayOrderNo, "form", "")
	}
	if r.merchantOrderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "merchant_order_no", r.merchantOrderNo, "form", "")
	}
	if r.storeNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_no", r.storeNo, "form", "")
	}
	if r.pageNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_no", r.pageNo, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
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
			var v AlipayDataBillBuyQueryDefaultResponse
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


func (a *AlipayDataBillBuyAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayDataBillBuyAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


