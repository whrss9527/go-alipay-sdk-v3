/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AlipayUserAgreementPermissionCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserAgreementPermissionCreateModel{}

// AlipayUserAgreementPermissionCreateModel struct for AlipayUserAgreementPermissionCreateModel
type AlipayUserAgreementPermissionCreateModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号（用户签约成功后的协议号 ）
	AgreementNo *string `json:"agreement_no,omitempty"`
	BusinessParams *BusinessParamsMap `json:"business_params,omitempty"`
	// 目前共两种类型SERVIOCE_NOTICE和DEFAULT_PERMISSION，如果是服务变更提醒，需传入SERVIOCE_NOTICE
	NoticeTemplate *string `json:"notice_template,omitempty"`
	// 商户请求号。 由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 订单总金额。 单位为元，精确到小数点后两位，取值范围：[0.01,100000000] 。
	TotalAmount *string `json:"total_amount,omitempty"`
}

// NewAlipayUserAgreementPermissionCreateModel instantiates a new AlipayUserAgreementPermissionCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserAgreementPermissionCreateModel() *AlipayUserAgreementPermissionCreateModel {
	this := AlipayUserAgreementPermissionCreateModel{}
	return &this
}

// NewAlipayUserAgreementPermissionCreateModelWithDefaults instantiates a new AlipayUserAgreementPermissionCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserAgreementPermissionCreateModelWithDefaults() *AlipayUserAgreementPermissionCreateModel {
	this := AlipayUserAgreementPermissionCreateModel{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayUserAgreementPermissionCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBusinessParams returns the BusinessParams field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateModel) GetBusinessParams() BusinessParamsMap {
	if o == nil || IsNil(o.BusinessParams) {
		var ret BusinessParamsMap
		return ret
	}
	return *o.BusinessParams
}

// GetBusinessParamsOk returns a tuple with the BusinessParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateModel) GetBusinessParamsOk() (*BusinessParamsMap, bool) {
	if o == nil || IsNil(o.BusinessParams) {
		return nil, false
	}
	return o.BusinessParams, true
}

// HasBusinessParams returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateModel) HasBusinessParams() bool {
	if o != nil && !IsNil(o.BusinessParams) {
		return true
	}

	return false
}

// SetBusinessParams gets a reference to the given BusinessParamsMap and assigns it to the BusinessParams field.
func (o *AlipayUserAgreementPermissionCreateModel) SetBusinessParams(v BusinessParamsMap) {
	o.BusinessParams = &v
}

// GetNoticeTemplate returns the NoticeTemplate field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateModel) GetNoticeTemplate() string {
	if o == nil || IsNil(o.NoticeTemplate) {
		var ret string
		return ret
	}
	return *o.NoticeTemplate
}

// GetNoticeTemplateOk returns a tuple with the NoticeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateModel) GetNoticeTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.NoticeTemplate) {
		return nil, false
	}
	return o.NoticeTemplate, true
}

// HasNoticeTemplate returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateModel) HasNoticeTemplate() bool {
	if o != nil && !IsNil(o.NoticeTemplate) {
		return true
	}

	return false
}

// SetNoticeTemplate gets a reference to the given string and assigns it to the NoticeTemplate field.
func (o *AlipayUserAgreementPermissionCreateModel) SetNoticeTemplate(v string) {
	o.NoticeTemplate = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayUserAgreementPermissionCreateModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateModel) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateModel) GetTotalAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateModel) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *AlipayUserAgreementPermissionCreateModel) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

func (o AlipayUserAgreementPermissionCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserAgreementPermissionCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BusinessParams) {
		toSerialize["business_params"] = o.BusinessParams
	}
	if !IsNil(o.NoticeTemplate) {
		toSerialize["notice_template"] = o.NoticeTemplate
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["total_amount"] = o.TotalAmount
	}
	return toSerialize, nil
}

type NullableAlipayUserAgreementPermissionCreateModel struct {
	value *AlipayUserAgreementPermissionCreateModel
	isSet bool
}

func (v NullableAlipayUserAgreementPermissionCreateModel) Get() *AlipayUserAgreementPermissionCreateModel {
	return v.value
}

func (v *NullableAlipayUserAgreementPermissionCreateModel) Set(val *AlipayUserAgreementPermissionCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementPermissionCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementPermissionCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementPermissionCreateModel(val *AlipayUserAgreementPermissionCreateModel) *NullableAlipayUserAgreementPermissionCreateModel {
	return &NullableAlipayUserAgreementPermissionCreateModel{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementPermissionCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementPermissionCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


