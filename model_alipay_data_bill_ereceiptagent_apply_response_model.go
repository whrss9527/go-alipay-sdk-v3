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

// checks if the AlipayDataBillEreceiptagentApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillEreceiptagentApplyResponseModel{}

// AlipayDataBillEreceiptagentApplyResponseModel struct for AlipayDataBillEreceiptagentApplyResponseModel
type AlipayDataBillEreceiptagentApplyResponseModel struct {
	// 文件申请号file_id信息。使用file_id可以查询处理状态
	FileId *string `json:"file_id,omitempty"`
}

// NewAlipayDataBillEreceiptagentApplyResponseModel instantiates a new AlipayDataBillEreceiptagentApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillEreceiptagentApplyResponseModel() *AlipayDataBillEreceiptagentApplyResponseModel {
	this := AlipayDataBillEreceiptagentApplyResponseModel{}
	return &this
}

// NewAlipayDataBillEreceiptagentApplyResponseModelWithDefaults instantiates a new AlipayDataBillEreceiptagentApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillEreceiptagentApplyResponseModelWithDefaults() *AlipayDataBillEreceiptagentApplyResponseModel {
	this := AlipayDataBillEreceiptagentApplyResponseModel{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *AlipayDataBillEreceiptagentApplyResponseModel) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillEreceiptagentApplyResponseModel) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *AlipayDataBillEreceiptagentApplyResponseModel) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *AlipayDataBillEreceiptagentApplyResponseModel) SetFileId(v string) {
	o.FileId = &v
}

func (o AlipayDataBillEreceiptagentApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillEreceiptagentApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	return toSerialize, nil
}

type NullableAlipayDataBillEreceiptagentApplyResponseModel struct {
	value *AlipayDataBillEreceiptagentApplyResponseModel
	isSet bool
}

func (v NullableAlipayDataBillEreceiptagentApplyResponseModel) Get() *AlipayDataBillEreceiptagentApplyResponseModel {
	return v.value
}

func (v *NullableAlipayDataBillEreceiptagentApplyResponseModel) Set(val *AlipayDataBillEreceiptagentApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillEreceiptagentApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillEreceiptagentApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillEreceiptagentApplyResponseModel(val *AlipayDataBillEreceiptagentApplyResponseModel) *NullableAlipayDataBillEreceiptagentApplyResponseModel {
	return &NullableAlipayDataBillEreceiptagentApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillEreceiptagentApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillEreceiptagentApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

