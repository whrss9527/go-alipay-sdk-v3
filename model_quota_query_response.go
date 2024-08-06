/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the QuotaQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaQueryResponse{}

// QuotaQueryResponse struct for QuotaQueryResponse
type QuotaQueryResponse struct {
	// 协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 单个协议查询结果错误码
	ErrorCode *string `json:"error_code,omitempty"`
	// 单个协议响应错误描述
	FailReason   *string             `json:"fail_reason,omitempty"`
	QuotaDetails *AccountQuotaDetail `json:"quota_details,omitempty"`
	// 单个协议查询结果是否成功
	Success *string `json:"success,omitempty"`
}

// NewQuotaQueryResponse instantiates a new QuotaQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaQueryResponse() *QuotaQueryResponse {
	this := QuotaQueryResponse{}
	return &this
}

// NewQuotaQueryResponseWithDefaults instantiates a new QuotaQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaQueryResponseWithDefaults() *QuotaQueryResponse {
	this := QuotaQueryResponse{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *QuotaQueryResponse) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaQueryResponse) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *QuotaQueryResponse) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *QuotaQueryResponse) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *QuotaQueryResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaQueryResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *QuotaQueryResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *QuotaQueryResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *QuotaQueryResponse) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaQueryResponse) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *QuotaQueryResponse) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *QuotaQueryResponse) SetFailReason(v string) {
	o.FailReason = &v
}

// GetQuotaDetails returns the QuotaDetails field value if set, zero value otherwise.
func (o *QuotaQueryResponse) GetQuotaDetails() AccountQuotaDetail {
	if o == nil || IsNil(o.QuotaDetails) {
		var ret AccountQuotaDetail
		return ret
	}
	return *o.QuotaDetails
}

// GetQuotaDetailsOk returns a tuple with the QuotaDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaQueryResponse) GetQuotaDetailsOk() (*AccountQuotaDetail, bool) {
	if o == nil || IsNil(o.QuotaDetails) {
		return nil, false
	}
	return o.QuotaDetails, true
}

// HasQuotaDetails returns a boolean if a field has been set.
func (o *QuotaQueryResponse) HasQuotaDetails() bool {
	if o != nil && !IsNil(o.QuotaDetails) {
		return true
	}

	return false
}

// SetQuotaDetails gets a reference to the given AccountQuotaDetail and assigns it to the QuotaDetails field.
func (o *QuotaQueryResponse) SetQuotaDetails(v AccountQuotaDetail) {
	o.QuotaDetails = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *QuotaQueryResponse) GetSuccess() string {
	if o == nil || IsNil(o.Success) {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaQueryResponse) GetSuccessOk() (*string, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *QuotaQueryResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given string and assigns it to the Success field.
func (o *QuotaQueryResponse) SetSuccess(v string) {
	o.Success = &v
}

func (o QuotaQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.QuotaDetails) {
		toSerialize["quota_details"] = o.QuotaDetails
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableQuotaQueryResponse struct {
	value *QuotaQueryResponse
	isSet bool
}

func (v NullableQuotaQueryResponse) Get() *QuotaQueryResponse {
	return v.value
}

func (v *NullableQuotaQueryResponse) Set(val *QuotaQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaQueryResponse(val *QuotaQueryResponse) *NullableQuotaQueryResponse {
	return &NullableQuotaQueryResponse{value: val, isSet: true}
}

func (v NullableQuotaQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
