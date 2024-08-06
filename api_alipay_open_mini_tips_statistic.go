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

// AlipayOpenMiniTipsStatisticAPIService AlipayOpenMiniTipsStatisticAPI service
type AlipayOpenMiniTipsStatisticAPIService service

type ApiAlipayOpenMiniTipsStatisticQueryRequest struct {
	ctx        context.Context
	ApiService *AlipayOpenMiniTipsStatisticAPIService
	deliveryId *string
	queryType  *string
	startDate  *string
	endDate    *string
}

// 收藏引导投放活动ID，供查询收藏引导活动配置接口调用 ，当以小程序维度查询数据（query_type为app）时delivery_id为空
func (r ApiAlipayOpenMiniTipsStatisticQueryRequest) DeliveryId(deliveryId string) ApiAlipayOpenMiniTipsStatisticQueryRequest {
	r.deliveryId = &deliveryId
	return r
}

// 查询类型，表示以当前维度进行数据聚合。
func (r ApiAlipayOpenMiniTipsStatisticQueryRequest) QueryType(queryType string) ApiAlipayOpenMiniTipsStatisticQueryRequest {
	r.queryType = &queryType
	return r
}

// 查询开始日期，精度为天
func (r ApiAlipayOpenMiniTipsStatisticQueryRequest) StartDate(startDate string) ApiAlipayOpenMiniTipsStatisticQueryRequest {
	r.startDate = &startDate
	return r
}

// 查询截止日期，精度为天
func (r ApiAlipayOpenMiniTipsStatisticQueryRequest) EndDate(endDate string) ApiAlipayOpenMiniTipsStatisticQueryRequest {
	r.endDate = &endDate
	return r
}

func (r ApiAlipayOpenMiniTipsStatisticQueryRequest) Execute() (*AlipayOpenMiniTipsStatisticQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayOpenMiniTipsStatisticQueryExecute(r)
}

/*
AlipayOpenMiniTipsStatisticQuery 小程序收藏引导汇总数据查询

小程序收藏引导tips文案投放汇总数据查询接口。可查询小程序维度或活动维度的tips曝光uv，收藏uv，以及收藏转化率。请先配置投放活动，否则查询结果返回参数为空。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayOpenMiniTipsStatisticQueryRequest
*/
func (r *AlipayOpenMiniTipsStatisticAPIService) AlipayOpenMiniTipsStatisticQuery(ctx context.Context) ApiAlipayOpenMiniTipsStatisticQueryRequest {
	return ApiAlipayOpenMiniTipsStatisticQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayOpenMiniTipsStatisticQueryResponseModel
func (a *AlipayOpenMiniTipsStatisticAPIService) AlipayOpenMiniTipsStatisticQueryExecute(r ApiAlipayOpenMiniTipsStatisticQueryRequest) (*AlipayOpenMiniTipsStatisticQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayOpenMiniTipsStatisticQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayOpenMiniTipsStatisticAPIService.AlipayOpenMiniTipsStatisticQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/open/mini/tips/statistic/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deliveryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_id", r.deliveryId, "form", "")
	}
	if r.queryType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query_type", r.queryType, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
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
		var v AlipayOpenMiniTipsStatisticQueryDefaultResponse
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

func (a *AlipayOpenMiniTipsStatisticAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayOpenMiniTipsStatisticAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}
