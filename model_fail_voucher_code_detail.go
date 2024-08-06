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

// checks if the FailVoucherCodeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailVoucherCodeDetail{}

// FailVoucherCodeDetail struct for FailVoucherCodeDetail
type FailVoucherCodeDetail struct {
	// 券码导入失败错误码
	ErrorCode *string `json:"error_code,omitempty"`
	// 券码导入失败错误原因描述
	ErrorMsg *string `json:"error_msg,omitempty"`
	// 导入失败的券码
	VoucherCode *string `json:"voucher_code,omitempty"`
}

// NewFailVoucherCodeDetail instantiates a new FailVoucherCodeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailVoucherCodeDetail() *FailVoucherCodeDetail {
	this := FailVoucherCodeDetail{}
	return &this
}

// NewFailVoucherCodeDetailWithDefaults instantiates a new FailVoucherCodeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailVoucherCodeDetailWithDefaults() *FailVoucherCodeDetail {
	this := FailVoucherCodeDetail{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *FailVoucherCodeDetail) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailVoucherCodeDetail) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *FailVoucherCodeDetail) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *FailVoucherCodeDetail) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *FailVoucherCodeDetail) GetErrorMsg() string {
	if o == nil || IsNil(o.ErrorMsg) {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailVoucherCodeDetail) GetErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMsg) {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *FailVoucherCodeDetail) HasErrorMsg() bool {
	if o != nil && !IsNil(o.ErrorMsg) {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *FailVoucherCodeDetail) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetVoucherCode returns the VoucherCode field value if set, zero value otherwise.
func (o *FailVoucherCodeDetail) GetVoucherCode() string {
	if o == nil || IsNil(o.VoucherCode) {
		var ret string
		return ret
	}
	return *o.VoucherCode
}

// GetVoucherCodeOk returns a tuple with the VoucherCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailVoucherCodeDetail) GetVoucherCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherCode) {
		return nil, false
	}
	return o.VoucherCode, true
}

// HasVoucherCode returns a boolean if a field has been set.
func (o *FailVoucherCodeDetail) HasVoucherCode() bool {
	if o != nil && !IsNil(o.VoucherCode) {
		return true
	}

	return false
}

// SetVoucherCode gets a reference to the given string and assigns it to the VoucherCode field.
func (o *FailVoucherCodeDetail) SetVoucherCode(v string) {
	o.VoucherCode = &v
}

func (o FailVoucherCodeDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailVoucherCodeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMsg) {
		toSerialize["error_msg"] = o.ErrorMsg
	}
	if !IsNil(o.VoucherCode) {
		toSerialize["voucher_code"] = o.VoucherCode
	}
	return toSerialize, nil
}

type NullableFailVoucherCodeDetail struct {
	value *FailVoucherCodeDetail
	isSet bool
}

func (v NullableFailVoucherCodeDetail) Get() *FailVoucherCodeDetail {
	return v.value
}

func (v *NullableFailVoucherCodeDetail) Set(val *FailVoucherCodeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFailVoucherCodeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFailVoucherCodeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailVoucherCodeDetail(val *FailVoucherCodeDetail) *NullableFailVoucherCodeDetail {
	return &NullableFailVoucherCodeDetail{value: val, isSet: true}
}

func (v NullableFailVoucherCodeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailVoucherCodeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


