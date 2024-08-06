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


// AlipayCommerceEcDepartmentInfoAPIService AlipayCommerceEcDepartmentInfoAPI service
type AlipayCommerceEcDepartmentInfoAPIService service

type ApiAlipayCommerceEcDepartmentInfoModifyRequest struct {
	ctx context.Context
	ApiService *AlipayCommerceEcDepartmentInfoAPIService
	alipayCommerceEcDepartmentInfoModifyModel *AlipayCommerceEcDepartmentInfoModifyModel
}

func (r ApiAlipayCommerceEcDepartmentInfoModifyRequest) AlipayCommerceEcDepartmentInfoModifyModel(alipayCommerceEcDepartmentInfoModifyModel AlipayCommerceEcDepartmentInfoModifyModel) ApiAlipayCommerceEcDepartmentInfoModifyRequest {
	r.alipayCommerceEcDepartmentInfoModifyModel = &alipayCommerceEcDepartmentInfoModifyModel
	return r
}

func (r ApiAlipayCommerceEcDepartmentInfoModifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcDepartmentInfoModifyExecute(r)
}

/*
AlipayCommerceEcDepartmentInfoModify 企业部门信息修改

修改企业部门信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayCommerceEcDepartmentInfoModifyRequest
*/
func (r *AlipayCommerceEcDepartmentInfoAPIService) AlipayCommerceEcDepartmentInfoModify(ctx context.Context) ApiAlipayCommerceEcDepartmentInfoModifyRequest {
	return ApiAlipayCommerceEcDepartmentInfoModifyRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayCommerceEcDepartmentInfoAPIService) AlipayCommerceEcDepartmentInfoModifyExecute(r ApiAlipayCommerceEcDepartmentInfoModifyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcDepartmentInfoAPIService.AlipayCommerceEcDepartmentInfoModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/department"

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
	localVarPostBody = r.alipayCommerceEcDepartmentInfoModifyModel

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
			var v AlipayCommerceEcDepartmentInfoModifyDefaultResponse
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


type ApiAlipayCommerceEcDepartmentInfoQueryRequest struct {
	ctx context.Context
	ApiService *AlipayCommerceEcDepartmentInfoAPIService
	enterpriseId *string
	departmentId *string
}

// 企业id
func (r ApiAlipayCommerceEcDepartmentInfoQueryRequest) EnterpriseId(enterpriseId string) ApiAlipayCommerceEcDepartmentInfoQueryRequest {
	r.enterpriseId = &enterpriseId
	return r
}

// 部门id
func (r ApiAlipayCommerceEcDepartmentInfoQueryRequest) DepartmentId(departmentId string) ApiAlipayCommerceEcDepartmentInfoQueryRequest {
	r.departmentId = &departmentId
	return r
}

func (r ApiAlipayCommerceEcDepartmentInfoQueryRequest) Execute() (*AlipayCommerceEcDepartmentInfoQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcDepartmentInfoQueryExecute(r)
}

/*
AlipayCommerceEcDepartmentInfoQuery 查询部门详情

根据部门id查询企业下某个部门的详细信息

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayCommerceEcDepartmentInfoQueryRequest
*/
func (r *AlipayCommerceEcDepartmentInfoAPIService) AlipayCommerceEcDepartmentInfoQuery(ctx context.Context) ApiAlipayCommerceEcDepartmentInfoQueryRequest {
	return ApiAlipayCommerceEcDepartmentInfoQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayCommerceEcDepartmentInfoQueryResponseModel
func (a *AlipayCommerceEcDepartmentInfoAPIService) AlipayCommerceEcDepartmentInfoQueryExecute(r ApiAlipayCommerceEcDepartmentInfoQueryRequest) (*AlipayCommerceEcDepartmentInfoQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayCommerceEcDepartmentInfoQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcDepartmentInfoAPIService.AlipayCommerceEcDepartmentInfoQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/department/info/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.enterpriseId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enterprise_id", r.enterpriseId, "form", "")
	}
	if r.departmentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "department_id", r.departmentId, "form", "")
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
			var v AlipayCommerceEcDepartmentInfoQueryDefaultResponse
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


func (a *AlipayCommerceEcDepartmentInfoAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceEcDepartmentInfoAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

