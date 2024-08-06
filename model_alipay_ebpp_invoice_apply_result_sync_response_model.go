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

// checks if the AlipayEbppInvoiceApplyResultSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceApplyResultSyncResponseModel{}

// AlipayEbppInvoiceApplyResultSyncResponseModel struct for AlipayEbppInvoiceApplyResultSyncResponseModel
type AlipayEbppInvoiceApplyResultSyncResponseModel struct {
	// 标注是否需要调用方重试。  当结果返回失败时（code 不等于 10000 且 msg 不等于 SUCCESS  ），如果该字段返回true表明该失败的情况通过重试补偿可解决，为false表明失败情况通过重试无法解决，可以停止重试，根据返回的错误码尝试解决。
	RetryFlag *bool `json:"retry_flag,omitempty"`
}

// NewAlipayEbppInvoiceApplyResultSyncResponseModel instantiates a new AlipayEbppInvoiceApplyResultSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceApplyResultSyncResponseModel() *AlipayEbppInvoiceApplyResultSyncResponseModel {
	this := AlipayEbppInvoiceApplyResultSyncResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceApplyResultSyncResponseModelWithDefaults instantiates a new AlipayEbppInvoiceApplyResultSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceApplyResultSyncResponseModelWithDefaults() *AlipayEbppInvoiceApplyResultSyncResponseModel {
	this := AlipayEbppInvoiceApplyResultSyncResponseModel{}
	return &this
}

// GetRetryFlag returns the RetryFlag field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyResultSyncResponseModel) GetRetryFlag() bool {
	if o == nil || IsNil(o.RetryFlag) {
		var ret bool
		return ret
	}
	return *o.RetryFlag
}

// GetRetryFlagOk returns a tuple with the RetryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyResultSyncResponseModel) GetRetryFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.RetryFlag) {
		return nil, false
	}
	return o.RetryFlag, true
}

// HasRetryFlag returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyResultSyncResponseModel) HasRetryFlag() bool {
	if o != nil && !IsNil(o.RetryFlag) {
		return true
	}

	return false
}

// SetRetryFlag gets a reference to the given bool and assigns it to the RetryFlag field.
func (o *AlipayEbppInvoiceApplyResultSyncResponseModel) SetRetryFlag(v bool) {
	o.RetryFlag = &v
}

func (o AlipayEbppInvoiceApplyResultSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceApplyResultSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetryFlag) {
		toSerialize["retry_flag"] = o.RetryFlag
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceApplyResultSyncResponseModel struct {
	value *AlipayEbppInvoiceApplyResultSyncResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceApplyResultSyncResponseModel) Get() *AlipayEbppInvoiceApplyResultSyncResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncResponseModel) Set(val *AlipayEbppInvoiceApplyResultSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceApplyResultSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceApplyResultSyncResponseModel(val *AlipayEbppInvoiceApplyResultSyncResponseModel) *NullableAlipayEbppInvoiceApplyResultSyncResponseModel {
	return &NullableAlipayEbppInvoiceApplyResultSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceApplyResultSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceApplyResultSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
