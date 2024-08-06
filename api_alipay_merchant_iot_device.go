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


// AlipayMerchantIotDeviceAPIService AlipayMerchantIotDeviceAPI service
type AlipayMerchantIotDeviceAPIService service

type ApiAlipayMerchantIotDeviceBindRequest struct {
	ctx context.Context
	ApiService *AlipayMerchantIotDeviceAPIService
	alipayMerchantIotDeviceBindModel *AlipayMerchantIotDeviceBindModel
}

func (r ApiAlipayMerchantIotDeviceBindRequest) AlipayMerchantIotDeviceBindModel(alipayMerchantIotDeviceBindModel AlipayMerchantIotDeviceBindModel) ApiAlipayMerchantIotDeviceBindRequest {
	r.alipayMerchantIotDeviceBindModel = &alipayMerchantIotDeviceBindModel
	return r
}

func (r ApiAlipayMerchantIotDeviceBindRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayMerchantIotDeviceBindExecute(r)
}

/*
AlipayMerchantIotDeviceBind IoT设备绑定门店

同步IoT设备、商户和门店的绑定关系

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMerchantIotDeviceBindRequest
*/
func (r *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceBind(ctx context.Context) ApiAlipayMerchantIotDeviceBindRequest {
	return ApiAlipayMerchantIotDeviceBindRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceBindExecute(r ApiAlipayMerchantIotDeviceBindRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMerchantIotDeviceAPIService.AlipayMerchantIotDeviceBind")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/merchant/iot/device/bind"

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
	localVarPostBody = r.alipayMerchantIotDeviceBindModel

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
			var v AlipayMerchantIotDeviceBindDefaultResponse
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


type ApiAlipayMerchantIotDeviceQueryRequest struct {
	ctx context.Context
	ApiService *AlipayMerchantIotDeviceAPIService
	deviceIdType *string
	bizTid *string
	supplierId *string
	deviceSn *string
}

// 可选方式 [ID,SN]。ID-使用biztid作为设备唯一识别标识；SN-使用supplier_id、device_sn联合作为设备唯一识别标识。由于不同机型的supplier_id不同，推荐使用 ID 。
func (r ApiAlipayMerchantIotDeviceQueryRequest) DeviceIdType(deviceIdType string) ApiAlipayMerchantIotDeviceQueryRequest {
	r.deviceIdType = &deviceIdType
	return r
}

// 设备 ID ，device_id_type 为 ID 时填写。
func (r ApiAlipayMerchantIotDeviceQueryRequest) BizTid(bizTid string) ApiAlipayMerchantIotDeviceQueryRequest {
	r.bizTid = &bizTid
	return r
}

// 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
func (r ApiAlipayMerchantIotDeviceQueryRequest) SupplierId(supplierId string) ApiAlipayMerchantIotDeviceQueryRequest {
	r.supplierId = &supplierId
	return r
}

// 设备序列号 ，device_id_type 为 SN 时填写。需配合supplier_id使用。
func (r ApiAlipayMerchantIotDeviceQueryRequest) DeviceSn(deviceSn string) ApiAlipayMerchantIotDeviceQueryRequest {
	r.deviceSn = &deviceSn
	return r
}

func (r ApiAlipayMerchantIotDeviceQueryRequest) Execute() (*AlipayMerchantIotDeviceQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayMerchantIotDeviceQueryExecute(r)
}

/*
AlipayMerchantIotDeviceQuery IoT设备绑定关系查询

通过设备唯一标识查询设备绑定关系

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMerchantIotDeviceQueryRequest
*/
func (r *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceQuery(ctx context.Context) ApiAlipayMerchantIotDeviceQueryRequest {
	return ApiAlipayMerchantIotDeviceQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayMerchantIotDeviceQueryResponseModel
func (a *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceQueryExecute(r ApiAlipayMerchantIotDeviceQueryRequest) (*AlipayMerchantIotDeviceQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayMerchantIotDeviceQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMerchantIotDeviceAPIService.AlipayMerchantIotDeviceQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/merchant/iot/device/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.deviceIdType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_id_type", r.deviceIdType, "form", "")
	}
	if r.bizTid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "biz_tid", r.bizTid, "form", "")
	}
	if r.supplierId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supplier_id", r.supplierId, "form", "")
	}
	if r.deviceSn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "device_sn", r.deviceSn, "form", "")
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
			var v AlipayMerchantIotDeviceQueryDefaultResponse
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


type ApiAlipayMerchantIotDeviceVerifyRequest struct {
	ctx context.Context
	ApiService *AlipayMerchantIotDeviceAPIService
	alipayMerchantIotDeviceVerifyModel *AlipayMerchantIotDeviceVerifyModel
}

func (r ApiAlipayMerchantIotDeviceVerifyRequest) AlipayMerchantIotDeviceVerifyModel(alipayMerchantIotDeviceVerifyModel AlipayMerchantIotDeviceVerifyModel) ApiAlipayMerchantIotDeviceVerifyRequest {
	r.alipayMerchantIotDeviceVerifyModel = &alipayMerchantIotDeviceVerifyModel
	return r
}

func (r ApiAlipayMerchantIotDeviceVerifyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AlipayMerchantIotDeviceVerifyExecute(r)
}

/*
AlipayMerchantIotDeviceVerify IoT设备绑定校验

校验IoT设备和商户收单账号的对应关系

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayMerchantIotDeviceVerifyRequest
*/
func (r *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceVerify(ctx context.Context) ApiAlipayMerchantIotDeviceVerifyRequest {
	return ApiAlipayMerchantIotDeviceVerifyRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AlipayMerchantIotDeviceAPIService) AlipayMerchantIotDeviceVerifyExecute(r ApiAlipayMerchantIotDeviceVerifyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayMerchantIotDeviceAPIService.AlipayMerchantIotDeviceVerify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/merchant/iot/device/verify"

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
	localVarPostBody = r.alipayMerchantIotDeviceVerifyModel

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
			var v AlipayMerchantIotDeviceVerifyDefaultResponse
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


func (a *AlipayMerchantIotDeviceAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayMerchantIotDeviceAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

