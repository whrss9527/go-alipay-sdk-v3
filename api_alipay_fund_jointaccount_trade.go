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

// AlipayFundJointaccountTradeAPIService AlipayFundJointaccountTradeAPI service
type AlipayFundJointaccountTradeAPIService service

type ApiAlipayFundJointaccountTradeQueryRequest struct {
	ctx             context.Context
	ApiService      *AlipayFundJointaccountTradeAPIService
	productCode     *string
	bizScene        *string
	memberId        *string
	memberOpenId    *string
	accountId       *string
	agreementNo     *string
	tradeNo         *string
	platformOrderId *string
}

// 销售产品码
func (r ApiAlipayFundJointaccountTradeQueryRequest) ProductCode(productCode string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.productCode = &productCode
	return r
}

// 业务场景
func (r ApiAlipayFundJointaccountTradeQueryRequest) BizScene(bizScene string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.bizScene = &bizScene
	return r
}

// 成员ID，消费发起人
func (r ApiAlipayFundJointaccountTradeQueryRequest) MemberId(memberId string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.memberId = &memberId
	return r
}

// 消费发起人的openId
func (r ApiAlipayFundJointaccountTradeQueryRequest) MemberOpenId(memberOpenId string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.memberOpenId = &memberOpenId
	return r
}

// 企业账户ID
func (r ApiAlipayFundJointaccountTradeQueryRequest) AccountId(accountId string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.accountId = &accountId
	return r
}

// 三方授权协议号
func (r ApiAlipayFundJointaccountTradeQueryRequest) AgreementNo(agreementNo string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 支付宝交易订单号
func (r ApiAlipayFundJointaccountTradeQueryRequest) TradeNo(tradeNo string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.tradeNo = &tradeNo
	return r
}

// 外部平台订单号。使用该参数查询时，返回结果见trade_info_list
func (r ApiAlipayFundJointaccountTradeQueryRequest) PlatformOrderId(platformOrderId string) ApiAlipayFundJointaccountTradeQueryRequest {
	r.platformOrderId = &platformOrderId
	return r
}

func (r ApiAlipayFundJointaccountTradeQueryRequest) Execute() (*AlipayFundJointaccountTradeQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundJointaccountTradeQueryExecute(r)
}

/*
AlipayFundJointaccountTradeQuery 共同账户交易查询

共同账户交易查询

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayFundJointaccountTradeQueryRequest
*/
func (r *AlipayFundJointaccountTradeAPIService) AlipayFundJointaccountTradeQuery(ctx context.Context) ApiAlipayFundJointaccountTradeQueryRequest {
	return ApiAlipayFundJointaccountTradeQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayFundJointaccountTradeQueryResponseModel
func (a *AlipayFundJointaccountTradeAPIService) AlipayFundJointaccountTradeQueryExecute(r ApiAlipayFundJointaccountTradeQueryRequest) (*AlipayFundJointaccountTradeQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayFundJointaccountTradeQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundJointaccountTradeAPIService.AlipayFundJointaccountTradeQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/jointaccount/trade/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.memberId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "member_id", r.memberId, "form", "")
	}
	if r.memberOpenId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "member_open_id", r.memberOpenId, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.agreementNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "agreement_no", r.agreementNo, "form", "")
	}
	if r.tradeNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "trade_no", r.tradeNo, "form", "")
	}
	if r.platformOrderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform_order_id", r.platformOrderId, "form", "")
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
		var v AlipayFundJointaccountTradeQueryDefaultResponse
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

func (a *AlipayFundJointaccountTradeAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundJointaccountTradeAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
