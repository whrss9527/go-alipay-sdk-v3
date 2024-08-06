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


// AntMerchantExpandShopAPIService AntMerchantExpandShopAPI service
type AntMerchantExpandShopAPIService service

type ApiAntMerchantExpandShopCloseRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandShopAPIService
	antMerchantExpandShopCloseModel *AntMerchantExpandShopCloseModel
}

func (r ApiAntMerchantExpandShopCloseRequest) AntMerchantExpandShopCloseModel(antMerchantExpandShopCloseModel AntMerchantExpandShopCloseModel) ApiAntMerchantExpandShopCloseRequest {
	r.antMerchantExpandShopCloseModel = &antMerchantExpandShopCloseModel
	return r
}

func (r ApiAntMerchantExpandShopCloseRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AntMerchantExpandShopCloseExecute(r)
}

/*
AntMerchantExpandShopClose 蚂蚁店铺关闭

通过shop_id，关闭蚂蚁店铺。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandShopCloseRequest
*/
func (r *AntMerchantExpandShopAPIService) AntMerchantExpandShopClose(ctx context.Context) ApiAntMerchantExpandShopCloseRequest {
	return ApiAntMerchantExpandShopCloseRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AntMerchantExpandShopAPIService) AntMerchantExpandShopCloseExecute(r ApiAntMerchantExpandShopCloseRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandShopAPIService.AntMerchantExpandShopClose")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/shop/close"

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
	localVarPostBody = r.antMerchantExpandShopCloseModel

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
			var v AntMerchantExpandShopCloseDefaultResponse
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


type ApiAntMerchantExpandShopConsultRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandShopAPIService
	antMerchantExpandShopConsultModel *AntMerchantExpandShopConsultModel
}

func (r ApiAntMerchantExpandShopConsultRequest) AntMerchantExpandShopConsultModel(antMerchantExpandShopConsultModel AntMerchantExpandShopConsultModel) ApiAntMerchantExpandShopConsultRequest {
	r.antMerchantExpandShopConsultModel = &antMerchantExpandShopConsultModel
	return r
}

func (r ApiAntMerchantExpandShopConsultRequest) Execute() (*AntMerchantExpandShopConsultResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandShopConsultExecute(r)
}

/*
AntMerchantExpandShopConsult 蚂蚁店铺创建咨询

蚂蚁店铺创建咨询

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandShopConsultRequest
*/
func (r *AntMerchantExpandShopAPIService) AntMerchantExpandShopConsult(ctx context.Context) ApiAntMerchantExpandShopConsultRequest {
	return ApiAntMerchantExpandShopConsultRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandShopConsultResponseModel
func (a *AntMerchantExpandShopAPIService) AntMerchantExpandShopConsultExecute(r ApiAntMerchantExpandShopConsultRequest) (*AntMerchantExpandShopConsultResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandShopConsultResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandShopAPIService.AntMerchantExpandShopConsult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/expand/shop/consult"

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
	localVarPostBody = r.antMerchantExpandShopConsultModel

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
			var v AntMerchantExpandShopConsultDefaultResponse
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


type ApiAntMerchantExpandShopCreateRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandShopAPIService
	antMerchantExpandShopCreateModel *AntMerchantExpandShopCreateModel
}

func (r ApiAntMerchantExpandShopCreateRequest) AntMerchantExpandShopCreateModel(antMerchantExpandShopCreateModel AntMerchantExpandShopCreateModel) ApiAntMerchantExpandShopCreateRequest {
	r.antMerchantExpandShopCreateModel = &antMerchantExpandShopCreateModel
	return r
}

func (r ApiAntMerchantExpandShopCreateRequest) Execute() (*AntMerchantExpandShopCreateResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandShopCreateExecute(r)
}

/*
AntMerchantExpandShopCreate 蚂蚁店铺创建

创建蚂蚁店铺

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandShopCreateRequest
*/
func (r *AntMerchantExpandShopAPIService) AntMerchantExpandShopCreate(ctx context.Context) ApiAntMerchantExpandShopCreateRequest {
	return ApiAntMerchantExpandShopCreateRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandShopCreateResponseModel
func (a *AntMerchantExpandShopAPIService) AntMerchantExpandShopCreateExecute(r ApiAntMerchantExpandShopCreateRequest) (*AntMerchantExpandShopCreateResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandShopCreateResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandShopAPIService.AntMerchantExpandShopCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/shop"

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
	localVarPostBody = r.antMerchantExpandShopCreateModel

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
			var v AntMerchantExpandShopCreateDefaultResponse
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


type ApiAntMerchantExpandShopModifyRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandShopAPIService
	antMerchantExpandShopModifyModel *AntMerchantExpandShopModifyModel
}

func (r ApiAntMerchantExpandShopModifyRequest) AntMerchantExpandShopModifyModel(antMerchantExpandShopModifyModel AntMerchantExpandShopModifyModel) ApiAntMerchantExpandShopModifyRequest {
	r.antMerchantExpandShopModifyModel = &antMerchantExpandShopModifyModel
	return r
}

func (r ApiAntMerchantExpandShopModifyRequest) Execute() (*AntMerchantExpandShopModifyResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandShopModifyExecute(r)
}

/*
AntMerchantExpandShopModify 修改蚂蚁店铺

修改蚂蚁店铺，按信息项修改。若无特殊说明，如果某项存在但是没填写，则不会覆盖掉原来的值

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandShopModifyRequest
*/
func (r *AntMerchantExpandShopAPIService) AntMerchantExpandShopModify(ctx context.Context) ApiAntMerchantExpandShopModifyRequest {
	return ApiAntMerchantExpandShopModifyRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandShopModifyResponseModel
func (a *AntMerchantExpandShopAPIService) AntMerchantExpandShopModifyExecute(r ApiAntMerchantExpandShopModifyRequest) (*AntMerchantExpandShopModifyResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandShopModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandShopAPIService.AntMerchantExpandShopModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/shop"

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
	localVarPostBody = r.antMerchantExpandShopModifyModel

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
			var v AntMerchantExpandShopModifyDefaultResponse
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


type ApiAntMerchantExpandShopQueryRequest struct {
	ctx context.Context
	ApiService *AntMerchantExpandShopAPIService
	shopId *string
	storeId *string
	ipRoleId *string
	addressVersion *string
	needRecommend *string
}

// 蚂蚁店铺id
func (r ApiAntMerchantExpandShopQueryRequest) ShopId(shopId string) ApiAntMerchantExpandShopQueryRequest {
	r.shopId = &shopId
	return r
}

// 门店编号，表示该门店在该商户角色id(直连pid，间连smid)下，由商户自己定义的外部门店编号
func (r ApiAntMerchantExpandShopQueryRequest) StoreId(storeId string) ApiAntMerchantExpandShopQueryRequest {
	r.storeId = &storeId
	return r
}

// 商户角色id，表示将要开的店属于哪个商户角色。对于直连开店场景，填写商户pid；对于间连开店场景（线上、线下、直付通），填写商户smid
func (r ApiAntMerchantExpandShopQueryRequest) IpRoleId(ipRoleId string) ApiAntMerchantExpandShopQueryRequest {
	r.ipRoleId = &ipRoleId
	return r
}

// 行政区划版本，当前可传空值(取默认版本)、2022Q2、UPTODATE(取最新版本)，其中空值默认为：2020Q1版本（ address_version&#x3D;&#39;&#39;或null），想要查看版本是2022年2季度版本则传入:(address_version&#x3D;&#39;2022Q2&#39;)，想要获取最新版本则传入:(address_version &#x3D;&#39;UPTODATE&#39;)
func (r ApiAntMerchantExpandShopQueryRequest) AddressVersion(addressVersion string) ApiAntMerchantExpandShopQueryRequest {
	r.addressVersion = &addressVersion
	return r
}

// 门店不置信时，是否需要返回shop_recommend_info
func (r ApiAntMerchantExpandShopQueryRequest) NeedRecommend(needRecommend string) ApiAntMerchantExpandShopQueryRequest {
	r.needRecommend = &needRecommend
	return r
}

func (r ApiAntMerchantExpandShopQueryRequest) Execute() (*AntMerchantExpandShopQueryResponseModel, *http.Response, error) {
	return r.ApiService.AntMerchantExpandShopQueryExecute(r)
}

/*
AntMerchantExpandShopQuery 店铺查询接口

用于服务商或商户查询其自己的店铺信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAntMerchantExpandShopQueryRequest
*/
func (r *AntMerchantExpandShopAPIService) AntMerchantExpandShopQuery(ctx context.Context) ApiAntMerchantExpandShopQueryRequest {
	return ApiAntMerchantExpandShopQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AntMerchantExpandShopQueryResponseModel
func (a *AntMerchantExpandShopAPIService) AntMerchantExpandShopQueryExecute(r ApiAntMerchantExpandShopQueryRequest) (*AntMerchantExpandShopQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AntMerchantExpandShopQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AntMerchantExpandShopAPIService.AntMerchantExpandShopQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/ant/merchant/shop"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.shopId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shop_id", r.shopId, "form", "")
	}
	if r.storeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "store_id", r.storeId, "form", "")
	}
	if r.ipRoleId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ip_role_id", r.ipRoleId, "form", "")
	}
	if r.addressVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address_version", r.addressVersion, "form", "")
	}
	if r.needRecommend != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "need_recommend", r.needRecommend, "form", "")
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
			var v AntMerchantExpandShopQueryDefaultResponse
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


func (a *AntMerchantExpandShopAPIService) signRequest(req *http.Request) error {
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

func (a *AntMerchantExpandShopAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

