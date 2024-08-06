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


// AlipayEcoEduKtParentAPIService AlipayEcoEduKtParentAPI service
type AlipayEcoEduKtParentAPIService service

type ApiAlipayEcoEduKtParentQueryRequest struct {
	ctx context.Context
	ApiService *AlipayEcoEduKtParentAPIService
	schoolPid *string
	schoolNo *string
	partnerId *string
	childName *string
	userMobile *string
	studentCode *string
	studentIdentify *string
}

// 学校支付宝pid
func (r ApiAlipayEcoEduKtParentQueryRequest) SchoolPid(schoolPid string) ApiAlipayEcoEduKtParentQueryRequest {
	r.schoolPid = &schoolPid
	return r
}

// 学校编码，录入学校接口返回的school_no参数
func (r ApiAlipayEcoEduKtParentQueryRequest) SchoolNo(schoolNo string) ApiAlipayEcoEduKtParentQueryRequest {
	r.schoolNo = &schoolNo
	return r
}

// Isv的支付宝pid
func (r ApiAlipayEcoEduKtParentQueryRequest) PartnerId(partnerId string) ApiAlipayEcoEduKtParentQueryRequest {
	r.partnerId = &partnerId
	return r
}

// 孩子或学生姓名
func (r ApiAlipayEcoEduKtParentQueryRequest) ChildName(childName string) ApiAlipayEcoEduKtParentQueryRequest {
	r.childName = &childName
	return r
}

// 用户手机号，发账单时填写users数组中的一个手机号。结果返回用户是否通过此手机号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
func (r ApiAlipayEcoEduKtParentQueryRequest) UserMobile(userMobile string) ApiAlipayEcoEduKtParentQueryRequest {
	r.userMobile = &userMobile
	return r
}

// 学生的学号，一个学校的学号必须是唯一。结果返回用户是否通过此学号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
func (r ApiAlipayEcoEduKtParentQueryRequest) StudentCode(studentCode string) ApiAlipayEcoEduKtParentQueryRequest {
	r.studentCode = &studentCode
	return r
}

// 学生的身份证号，使用身份证规则验证。结果返回用户是否通过此身份证号添加此学生的缴费账户。user_mobile、student_code 、student_identify 必须并且只能填一项。
func (r ApiAlipayEcoEduKtParentQueryRequest) StudentIdentify(studentIdentify string) ApiAlipayEcoEduKtParentQueryRequest {
	r.studentIdentify = &studentIdentify
	return r
}

func (r ApiAlipayEcoEduKtParentQueryRequest) Execute() (*AlipayEcoEduKtParentQueryResponseModel, *http.Response, error) {
	return r.ApiService.AlipayEcoEduKtParentQueryExecute(r)
}

/*
AlipayEcoEduKtParentQuery 查询学生家长状态接口

ISV发送账单后，用户在支付宝缴费成功。ISV可以通过此接口查询，用户缴费的时候，是否创建了缴费账户，如果已经添加，学校使用ISV的系统再次发账单就可以通过支付宝的通知触达用户；如果未添加，学校即可联系家长，督促家长在支付宝添加学生的缴费账户。
目前的出现的场景是家长第一次缴费成功后，学校使用ISV的系统再次发账单，认为家长应该收到通知，实际上只有建立缴费账户的家长才能收到，未建立缴费账户的是收不到通知，这部分家长有可能就缴费不及时或不通过支付宝缴费。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAlipayEcoEduKtParentQueryRequest
*/
func (r *AlipayEcoEduKtParentAPIService) AlipayEcoEduKtParentQuery(ctx context.Context) ApiAlipayEcoEduKtParentQueryRequest {
	return ApiAlipayEcoEduKtParentQueryRequest{
		ApiService: r,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AlipayEcoEduKtParentQueryResponseModel
func (a *AlipayEcoEduKtParentAPIService) AlipayEcoEduKtParentQueryExecute(r ApiAlipayEcoEduKtParentQueryRequest) (*AlipayEcoEduKtParentQueryResponseModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AlipayEcoEduKtParentQueryResponseModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AlipayEcoEduKtParentAPIService.AlipayEcoEduKtParentQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/alipay/eco/edu/kt/parent/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	if r.schoolPid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "school_pid", r.schoolPid, "form", "")
	}
	if r.schoolNo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "school_no", r.schoolNo, "form", "")
	}
	if r.partnerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "partner_id", r.partnerId, "form", "")
	}
	if r.childName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "child_name", r.childName, "form", "")
	}
	if r.userMobile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_mobile", r.userMobile, "form", "")
	}
	if r.studentCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "student_code", r.studentCode, "form", "")
	}
	if r.studentIdentify != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "student_identify", r.studentIdentify, "form", "")
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
			var v AlipayEcoEduKtParentQueryDefaultResponse
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


func (a *AlipayEcoEduKtParentAPIService) signRequest(req *http.Request) error {
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

func (a *AlipayEcoEduKtParentAPIService) verifyResponse(resp *http.Response, body []byte) error {
	timestamp := resp.Header.Get("alipay-timestamp")
	nonce := resp.Header.Get("alipay-nonce")
	sign := resp.Header.Get("alipay-signature")

	content := timestamp + "\n" +
		nonce + "\n" +
		string(body) + "\n"

	publicKey := a.client.cfg.PublicKey

	return verifyWithRSA(content, sign, publicKey)
}

