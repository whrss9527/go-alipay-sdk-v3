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

// checks if the QuotaModifyDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaModifyDetail{}

// QuotaModifyDetail struct for QuotaModifyDetail
type QuotaModifyDetail struct {
	// 授权协议号，和入参的协议号对应
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 明细处理错误码，当success为false时有值
	ErrorCode *string `json:"error_code,omitempty"`
	// 错误描述，当success为false时有值
	FailReason *string `json:"fail_reason,omitempty"`
	// 当前协议下的明细处理结果
	Success *bool `json:"success,omitempty"`
}

// NewQuotaModifyDetail instantiates a new QuotaModifyDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaModifyDetail() *QuotaModifyDetail {
	this := QuotaModifyDetail{}
	return &this
}

// NewQuotaModifyDetailWithDefaults instantiates a new QuotaModifyDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaModifyDetailWithDefaults() *QuotaModifyDetail {
	this := QuotaModifyDetail{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *QuotaModifyDetail) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaModifyDetail) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *QuotaModifyDetail) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *QuotaModifyDetail) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *QuotaModifyDetail) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaModifyDetail) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *QuotaModifyDetail) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *QuotaModifyDetail) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *QuotaModifyDetail) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaModifyDetail) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *QuotaModifyDetail) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *QuotaModifyDetail) SetFailReason(v string) {
	o.FailReason = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *QuotaModifyDetail) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaModifyDetail) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *QuotaModifyDetail) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *QuotaModifyDetail) SetSuccess(v bool) {
	o.Success = &v
}

func (o QuotaModifyDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaModifyDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableQuotaModifyDetail struct {
	value *QuotaModifyDetail
	isSet bool
}

func (v NullableQuotaModifyDetail) Get() *QuotaModifyDetail {
	return v.value
}

func (v *NullableQuotaModifyDetail) Set(val *QuotaModifyDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaModifyDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaModifyDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaModifyDetail(val *QuotaModifyDetail) *NullableQuotaModifyDetail {
	return &NullableQuotaModifyDetail{value: val, isSet: true}
}

func (v NullableQuotaModifyDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaModifyDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
