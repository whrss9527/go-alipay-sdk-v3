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

// AlipayCommerceEcEnterpriseAddressAPIService AlipayCommerceEcEnterpriseAddressAPI service
type AlipayCommerceEcEnterpriseAddressAPIService service

type ApiAlipayCommerceEcEnterpriseAddressAddRequest struct {
	ctx                                       context.Context
	ApiService                                *AlipayCommerceEcEnterpriseAddressAPIService
	alipayCommerceEcEnterpriseAddressAddModel *AlipayCommerceEcEnterpriseAddressAddModel
}

func (r ApiAlipayCommerceEcEnterpriseAddressAddRequest) AlipayCommerceEcEnterpriseAddressAddModel(alipayCommerceEcEnterpriseAddressAddModel AlipayCommerceEcEnterpriseAddressAddModel) ApiAlipayCommerceEcEnterpriseAddressAddRequest {
	r.alipayCommerceEcEnterpriseAddressAddModel = &alipayCommerceEcEnterpriseAddressAddModel
	return r
}

func (r ApiAlipayCommerceEcEnterpriseAddressAddRequest) Execute() (*AlipayCommerceEcEnterpriseAddressAddResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcEnterpriseAddressAddExecute(r)
}

/*
AlipayCommerceEcEnterpriseAddressAdd 企业地址添加

企业地址添加

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceEcEnterpriseAddressAddRequest
*/
func (r *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressAdd(ctx context.Context) ApiAlipayCommerceEcEnterpriseAddressAddRequest {
	return ApiAlipayCommerceEcEnterpriseAddressAddRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceEcEnterpriseAddressAddResponseModel
func (a *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressAddExecute(r ApiAlipayCommerceEcEnterpriseAddressAddRequest) (*AlipayCommerceEcEnterpriseAddressAddResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceEcEnterpriseAddressAddResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcEnterpriseAddressAPIService.AlipayCommerceEcEnterpriseAddressAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/enterprise/address"

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
	localVarPostBody = r.alipayCommerceEcEnterpriseAddressAddModel

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
		var v AlipayCommerceEcEnterpriseAddressAddDefaultResponse
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

type ApiAlipayCommerceEcEnterpriseAddressModifyRequest struct {
	ctx                                          context.Context
	ApiService                                   *AlipayCommerceEcEnterpriseAddressAPIService
	alipayCommerceEcEnterpriseAddressModifyModel *AlipayCommerceEcEnterpriseAddressModifyModel
}

func (r ApiAlipayCommerceEcEnterpriseAddressModifyRequest) AlipayCommerceEcEnterpriseAddressModifyModel(alipayCommerceEcEnterpriseAddressModifyModel AlipayCommerceEcEnterpriseAddressModifyModel) ApiAlipayCommerceEcEnterpriseAddressModifyRequest {
	r.alipayCommerceEcEnterpriseAddressModifyModel = &alipayCommerceEcEnterpriseAddressModifyModel
	return r
}

func (r ApiAlipayCommerceEcEnterpriseAddressModifyRequest) Execute() (*AlipayCommerceEcEnterpriseAddressModifyResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcEnterpriseAddressModifyExecute(r)
}

/*
AlipayCommerceEcEnterpriseAddressModify 企业地址修改

企业地址修改

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceEcEnterpriseAddressModifyRequest
*/
func (r *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressModify(ctx context.Context) ApiAlipayCommerceEcEnterpriseAddressModifyRequest {
	return ApiAlipayCommerceEcEnterpriseAddressModifyRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceEcEnterpriseAddressModifyResponseModel
func (a *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressModifyExecute(r ApiAlipayCommerceEcEnterpriseAddressModifyRequest) (*AlipayCommerceEcEnterpriseAddressModifyResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceEcEnterpriseAddressModifyResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcEnterpriseAddressAPIService.AlipayCommerceEcEnterpriseAddressModify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/enterprise/address"

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
	localVarPostBody = r.alipayCommerceEcEnterpriseAddressModifyModel

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
		var v AlipayCommerceEcEnterpriseAddressModifyDefaultResponse
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

type ApiAlipayCommerceEcEnterpriseAddressQueryRequest struct {
	ctx          context.Context
	ApiService   *AlipayCommerceEcEnterpriseAddressAPIService
	enterpriseId *string
	accountId    *string
	agreementNo  *string
	addressId    *string
	cityCode     *string
	pageNum      *int32
	pageSize     *int32
}

// 通过企业码2.0签约接口签约，只填写企业id，无需填写共同账户id和授权签约协议号。
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) EnterpriseId(enterpriseId string) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.enterpriseId = &enterpriseId
	return r
}

// 通过企业码1.0接口签约的共同账户，和agreement_no搭配使用。
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) AccountId(accountId string) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.accountId = &accountId
	return r
}

// 可通过签约消息获取。配合共同账户id使用，当填写企业共同账户id时，此字段必填。
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) AgreementNo(agreementNo string) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.agreementNo = &agreementNo
	return r
}

// 地址id
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) AddressId(addressId string) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.addressId = &addressId
	return r
}

// 市
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) CityCode(cityCode string) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.cityCode = &cityCode
	return r
}

// 页码从1开始
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) PageNum(pageNum int32) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.pageNum = &pageNum
	return r
}

// 每页数据
func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) PageSize(pageSize int32) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) Execute() (*AlipayCommerceEcEnterpriseAddressQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayCommerceEcEnterpriseAddressQueryExecute(r)
}

/*
AlipayCommerceEcEnterpriseAddressQuery 企业地址查询

查询企业地址信息

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAlipayCommerceEcEnterpriseAddressQueryRequest
*/
func (r *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressQuery(ctx context.Context) ApiAlipayCommerceEcEnterpriseAddressQueryRequest {
	return ApiAlipayCommerceEcEnterpriseAddressQueryRequest{
		ApiService: r,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AlipayCommerceEcEnterpriseAddressQueryResponseModel
func (a *AlipayCommerceEcEnterpriseAddressAPIService) AlipayCommerceEcEnterpriseAddressQueryExecute(r ApiAlipayCommerceEcEnterpriseAddressQueryRequest) (*AlipayCommerceEcEnterpriseAddressQueryResponseModel, *http.Response, error) {
	err := a.client.prepareConfig()
	if err != nil {
		return nil, nil, &GenericOpenAPIError{error: err.Error()}
	}
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AlipayCommerceEcEnterpriseAddressQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayCommerceEcEnterpriseAddressAPIService.AlipayCommerceEcEnterpriseAddressQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/commerce/ec/enterprise/address"

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
	if r.addressId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address_id", r.addressId, "form", "")
	}
	if r.cityCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "city_code", r.cityCode, "form", "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_num", r.pageNum, "form", "")
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
		var v AlipayCommerceEcEnterpriseAddressQueryDefaultResponse
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

func (a *AlipayCommerceEcEnterpriseAddressAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayCommerceEcEnterpriseAddressAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	return verifyWithRSA(content, sign, a.client.cfg.publicKey)
}
