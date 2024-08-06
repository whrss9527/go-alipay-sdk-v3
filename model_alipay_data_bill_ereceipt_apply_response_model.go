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

// checks if the AlipayDataBillEreceiptApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillEreceiptApplyResponseModel{}

// AlipayDataBillEreceiptApplyResponseModel struct for AlipayDataBillEreceiptApplyResponseModel
type AlipayDataBillEreceiptApplyResponseModel struct {
	// 文件申请号file_id信息。 使用file_id可以查询处理状态，有效期：2天
	FileId *string `json:"file_id,omitempty"`
}

// NewAlipayDataBillEreceiptApplyResponseModel instantiates a new AlipayDataBillEreceiptApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillEreceiptApplyResponseModel() *AlipayDataBillEreceiptApplyResponseModel {
	this := AlipayDataBillEreceiptApplyResponseModel{}
	return &this
}

// NewAlipayDataBillEreceiptApplyResponseModelWithDefaults instantiates a new AlipayDataBillEreceiptApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillEreceiptApplyResponseModelWithDefaults() *AlipayDataBillEreceiptApplyResponseModel {
	this := AlipayDataBillEreceiptApplyResponseModel{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *AlipayDataBillEreceiptApplyResponseModel) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillEreceiptApplyResponseModel) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *AlipayDataBillEreceiptApplyResponseModel) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *AlipayDataBillEreceiptApplyResponseModel) SetFileId(v string) {
	o.FileId = &v
}

func (o AlipayDataBillEreceiptApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillEreceiptApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	return toSerialize, nil
}

type NullableAlipayDataBillEreceiptApplyResponseModel struct {
	value *AlipayDataBillEreceiptApplyResponseModel
	isSet bool
}

func (v NullableAlipayDataBillEreceiptApplyResponseModel) Get() *AlipayDataBillEreceiptApplyResponseModel {
	return v.value
}

func (v *NullableAlipayDataBillEreceiptApplyResponseModel) Set(val *AlipayDataBillEreceiptApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillEreceiptApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillEreceiptApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillEreceiptApplyResponseModel(val *AlipayDataBillEreceiptApplyResponseModel) *NullableAlipayDataBillEreceiptApplyResponseModel {
	return &NullableAlipayDataBillEreceiptApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillEreceiptApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillEreceiptApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
