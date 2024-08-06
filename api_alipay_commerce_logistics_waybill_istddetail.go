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

// AlipayCommerceLogisticsWaybillIstddetailAPIService AlipayCommerceLogisticsWaybillIstddetailAPI service
type AlipayCommerceLogisticsWaybillIstddetailAPIService service

type ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest struct {
	ctx           context.Context
	ApiService    *AlipayCommerceLogisticsWaybillIstddetailAPIService
	shopNo        *string
	outOrderNo    *string
	orderNo       *string
	logisticsCode *string
	waybillNo     *string
}

// 商家门店编号
func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) ShopNo(shopNo string) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	r.shopNo = &shopNo
	return r
}

// 商家订单号
func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) OutOrderNo(outOrderNo string) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	r.outOrderNo = &outOrderNo
	return r
}

// 支付宝订单流水号
func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) OrderNo(orderNo string) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	r.orderNo = &orderNo
	return r
}

// 即时配送公司编码
func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) LogisticsCode(logisticsCode string) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	r.logisticsCode = &logisticsCode
	return r
}

// 即时配送运单编号
func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) WaybillNo(waybillNo string) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	r.waybillNo = &waybillNo
	return r
}

func (r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) Execute() (*AlipayCommerceLogisticsWaybillIstddetailQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceLogisticsWaybillIstddetailQueryExecute(r)
}

/*
AlipayCommerceLogisticsWaybillIstddetailQuery 查询即时配送运单详情

查询即时配送运单详情

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest
*/
func (r *AlipayCommerceLogisticsWaybillIstddetailAPIService) AlipayCommerceLogisticsWaybillIstddetailQuery(ctx context.Context) ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest {
	return ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceLogisticsWaybillIstddetailQueryResponseModel
func (a *AlipayCommerceLogisticsWaybillIstddetailAPIService) AlipayCommerceLogisticsWaybillIstddetailQueryExecute(r ApiAlipayCommerceLogisticsWaybillIstddetailQueryRequest) (*AlipayCommerceLogisticsWaybillIstddetailQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceLogisticsWaybillIstddetailQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceLogisticsWaybillIstddetailAPIService.AlipayCommerceLogisticsWaybillIstddetailQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/logistics/waybill/istddetail/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.shopNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shop_no", r.shopNo, "form", "")
	}
	if r.outOrderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_order_no", r.outOrderNo, "form", "")
	}
	if r.orderNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_no", r.orderNo, "form", "")
	}
	if r.logisticsCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "logistics_code", r.logisticsCode, "form", "")
	}
	if r.waybillNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "waybill_no", r.waybillNo, "form", "")
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
		var v AlipayCommerceLogisticsWaybillIstddetailQueryDefaultResponse
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

func (a *AlipayCommerceLogisticsWaybillIstddetailAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceLogisticsWaybillIstddetailAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
