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


// AlipayFundJointaccountSignAPIService AlipayFundJointaccountSignAPI service
type AlipayFundJointaccountSignAPIService service

type ApiAlipayFundJointaccountSignQueryRequest struct {
	ctx context.Context
	ApiService *AlipayFundJointaccountSignAPIService
	productCode *string
	bizScene *string
	outBizNo *string
	accountId *string
}

// 产品码，默认值 ENTERPRISE_PAY
func (r ApiAlipayFundJointaccountSignQueryRequest) ProductCode(productCode string) ApiAlipayFundJointaccountSignQueryRequest {
	r.productCode = &productCode
	return r
}

// 场景码，联系支付宝分配
func (r ApiAlipayFundJointaccountSignQueryRequest) BizScene(bizScene string) ApiAlipayFundJointaccountSignQueryRequest {
	r.bizScene = &bizScene
	return r
}

// 外部业务号，查询对应账户状态，优先级高于account_id
func (r ApiAlipayFundJointaccountSignQueryRequest) OutBizNo(outBizNo string) ApiAlipayFundJointaccountSignQueryRequest {
	r.outBizNo = &outBizNo
	return r
}

// 企业签约账户ID,用于外部商户已获取企业签约ID，查询账户状态
func (r ApiAlipayFundJointaccountSignQueryRequest) AccountId(accountId string) ApiAlipayFundJointaccountSignQueryRequest {
	r.accountId = &accountId
	return r
}

func (r ApiAlipayFundJointaccountSignQueryRequest) Execute() (*AlipayFundJointaccountSignQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayFundJointaccountSignQueryExecute(r)
}

/*
AlipayFundJointaccountSignQuery 企业签约结果查询

通过外部订单号（out_biz_no）主动查询企业签约结果，可用于未收到签约结果通知的补偿操作

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayFundJointaccountSignQueryRequest
*/
func (r *AlipayFundJointaccountSignAPIService) AlipayFundJointaccountSignQuery(ctx context.Context) ApiAlipayFundJointaccountSignQueryRequest {
	return ApiAlipayFundJointaccountSignQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayFundJointaccountSignQueryResponseModel
func (a *AlipayFundJointaccountSignAPIService) AlipayFundJointaccountSignQueryExecute(r ApiAlipayFundJointaccountSignQueryRequest) (*AlipayFundJointaccountSignQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayFundJointaccountSignQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayFundJointaccountSignAPIService.AlipayFundJointaccountSignQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/fund/jointaccount/sign/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.productCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_code", r.productCode, "form", "")
	}
	if r.bizScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_scene", r.bizScene, "form", "")
	}
	if r.outBizNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_biz_no", r.outBizNo, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
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
			var v AlipayFundJointaccountSignQueryDefaultResponse
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


func (a *AlipayFundJointaccountSignAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayFundJointaccountSignAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


