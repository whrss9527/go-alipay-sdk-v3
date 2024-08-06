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

// checks if the AlipayDataBillAccountbookereceiptApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillAccountbookereceiptApplyResponseModel{}

// AlipayDataBillAccountbookereceiptApplyResponseModel struct for AlipayDataBillAccountbookereceiptApplyResponseModel
type AlipayDataBillAccountbookereceiptApplyResponseModel struct {
	// 文件申请号file_id信息。使用file_id可以查询处理状态
	FileId *string `json:"file_id,omitempty"`
}

// NewAlipayDataBillAccountbookereceiptApplyResponseModel instantiates a new AlipayDataBillAccountbookereceiptApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillAccountbookereceiptApplyResponseModel() *AlipayDataBillAccountbookereceiptApplyResponseModel {
	this := AlipayDataBillAccountbookereceiptApplyResponseModel{}
	return &this
}

// NewAlipayDataBillAccountbookereceiptApplyResponseModelWithDefaults instantiates a new AlipayDataBillAccountbookereceiptApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillAccountbookereceiptApplyResponseModelWithDefaults() *AlipayDataBillAccountbookereceiptApplyResponseModel {
	this := AlipayDataBillAccountbookereceiptApplyResponseModel{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *AlipayDataBillAccountbookereceiptApplyResponseModel) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountbookereceiptApplyResponseModel) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *AlipayDataBillAccountbookereceiptApplyResponseModel) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *AlipayDataBillAccountbookereceiptApplyResponseModel) SetFileId(v string) {
	o.FileId = &v
}

func (o AlipayDataBillAccountbookereceiptApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillAccountbookereceiptApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["file_id"] = o.FileId
	}
	return toSerialize, nil
}

type NullableAlipayDataBillAccountbookereceiptApplyResponseModel struct {
	value *AlipayDataBillAccountbookereceiptApplyResponseModel
	isSet bool
}

func (v NullableAlipayDataBillAccountbookereceiptApplyResponseModel) Get() *AlipayDataBillAccountbookereceiptApplyResponseModel {
	return v.value
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyResponseModel) Set(val *AlipayDataBillAccountbookereceiptApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillAccountbookereceiptApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillAccountbookereceiptApplyResponseModel(val *AlipayDataBillAccountbookereceiptApplyResponseModel) *NullableAlipayDataBillAccountbookereceiptApplyResponseModel {
	return &NullableAlipayDataBillAccountbookereceiptApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillAccountbookereceiptApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


