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


// AlipayDataDataserviceAdPromotepageAPIService AlipayDataDataserviceAdPromotepageAPI service
type AlipayDataDataserviceAdPromotepageAPIService service

type ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest struct {
	ctx context.Context
	ApiService *AlipayDataDataserviceAdPromotepageAPIService
	bizToken *string
	principalTag *string
	type_ *string
	pageNo *int32
	pageSize *int32
}

// 代理商访问灯火平台的token
func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) BizToken(bizToken string) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	r.bizToken = &bizToken
	return r
}

// 商家标志
func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) PrincipalTag(principalTag string) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	r.principalTag = &principalTag
	return r
}

// 推广页类型：COLLECT_INFO -  免费留资；TRADE - 付费留资； OPERATION_PAID - 运营商付费留资；待扩展
func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) Type_(type_ string) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	r.type_ = &type_
	return r
}

// 分页参数之页数，从1开始
func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) PageNo(pageNo int32) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	r.pageNo = &pageNo
	return r
}

// 分页参数之每页大小，最大1000
func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) PageSize(pageSize int32) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) Execute() (*AlipayDataDataserviceAdPromotepageBatchqueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayDataDataserviceAdPromotepageBatchqueryExecute(r)
}

/*
AlipayDataDataserviceAdPromotepageBatchquery 自建推广页列表批量查询

用于获取指定商家的自建推广页列表

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest
*/
func (r *AlipayDataDataserviceAdPromotepageAPIService) AlipayDataDataserviceAdPromotepageBatchquery(ctx context.Context) ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest {
	return ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayDataDataserviceAdPromotepageBatchqueryResponseModel
func (a *AlipayDataDataserviceAdPromotepageAPIService) AlipayDataDataserviceAdPromotepageBatchqueryExecute(r ApiAlipayDataDataserviceAdPromotepageBatchqueryRequest) (*AlipayDataDataserviceAdPromotepageBatchqueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayDataDataserviceAdPromotepageBatchqueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayDataDataserviceAdPromotepageAPIService.AlipayDataDataserviceAdPromotepageBatchquery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/data/dataservice/ad/promotepage/batchquery"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.bizToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_token", r.bizToken, "form", "")
	}
	if r.principalTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "principal_tag", r.principalTag, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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
			var v AlipayDataDataserviceAdPromotepageBatchqueryDefaultResponse
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


type ApiAlipayDataDataserviceAdPromotepageDownloadRequest struct {
	ctx context.Context
	ApiService *AlipayDataDataserviceAdPromotepageAPIService
	startDate *string
	endDate *string
	pageNo *int32
	pageSize *int32
	bizToken *string
	principalTag *string
	promotePageId *int32
}

// 留资开始日期，格式：yyyy-mm-dd，不能早于30天前
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) StartDate(startDate string) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.startDate = &startDate
	return r
}

// 留资结束日期，格式：yyyy-mm-dd；不能晚于当天
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) EndDate(endDate string) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.endDate = &endDate
	return r
}

// 分页参数之页数，从1开始
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) PageNo(pageNo int32) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.pageNo = &pageNo
	return r
}

// 分页参数之每页大小，最大1000
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) PageSize(pageSize int32) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.pageSize = &pageSize
	return r
}

// 代理商访问灯火平台的token
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) BizToken(bizToken string) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.bizToken = &bizToken
	return r
}

// 商家标志
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) PrincipalTag(principalTag string) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.principalTag = &principalTag
	return r
}

// 推广页id
func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) PromotePageId(promotePageId int32) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	r.promotePageId = &promotePageId
	return r
}

func (r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) Execute() (*AlipayDataDataserviceAdPromotepageDownloadResponseModel, *http.Response, error) {
	return r.ApiService.AlipayDataDataserviceAdPromotepageDownloadExecute(r)
}

/*
AlipayDataDataserviceAdPromotepageDownload 自建推广页留资数据查询

用于获取指定商家指定推广页的留资数据，注意调用频次小于20qps

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayDataDataserviceAdPromotepageDownloadRequest
*/
func (r *AlipayDataDataserviceAdPromotepageAPIService) AlipayDataDataserviceAdPromotepageDownload(ctx context.Context) ApiAlipayDataDataserviceAdPromotepageDownloadRequest {
	return ApiAlipayDataDataserviceAdPromotepageDownloadRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayDataDataserviceAdPromotepageDownloadResponseModel
func (a *AlipayDataDataserviceAdPromotepageAPIService) AlipayDataDataserviceAdPromotepageDownloadExecute(r ApiAlipayDataDataserviceAdPromotepageDownloadRequest) (*AlipayDataDataserviceAdPromotepageDownloadResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayDataDataserviceAdPromotepageDownloadResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayDataDataserviceAdPromotepageAPIService.AlipayDataDataserviceAdPromotepageDownload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/data/dataservice/ad/promotepage/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "form", "")
	}
	if r.pageNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_no", r.pageNo, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.bizToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_token", r.bizToken, "form", "")
	}
	if r.principalTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "principal_tag", r.principalTag, "form", "")
	}
	if r.promotePageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promote_page_id", r.promotePageId, "form", "")
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
			var v AlipayDataDataserviceAdPromotepageDownloadDefaultResponse
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


func (a *AlipayDataDataserviceAdPromotepageAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayDataDataserviceAdPromotepageAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}


