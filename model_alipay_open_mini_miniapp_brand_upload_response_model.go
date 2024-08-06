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

// checks if the AlipayOpenMiniMiniappBrandUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniMiniappBrandUploadResponseModel{}

// AlipayOpenMiniMiniappBrandUploadResponseModel struct for AlipayOpenMiniMiniappBrandUploadResponseModel
type AlipayOpenMiniMiniappBrandUploadResponseModel struct {
	// 图片上传成功后的图片存储键(可用于填写调用alipay.open.mini.miniapp.brand.submit时需要的图片类型的字段)
	FileKey *string `json:"file_key,omitempty"`
}

// NewAlipayOpenMiniMiniappBrandUploadResponseModel instantiates a new AlipayOpenMiniMiniappBrandUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniMiniappBrandUploadResponseModel() *AlipayOpenMiniMiniappBrandUploadResponseModel {
	this := AlipayOpenMiniMiniappBrandUploadResponseModel{}
	return &this
}

// NewAlipayOpenMiniMiniappBrandUploadResponseModelWithDefaults instantiates a new AlipayOpenMiniMiniappBrandUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniMiniappBrandUploadResponseModelWithDefaults() *AlipayOpenMiniMiniappBrandUploadResponseModel {
	this := AlipayOpenMiniMiniappBrandUploadResponseModel{}
	return &this
}

// GetFileKey returns the FileKey field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandUploadResponseModel) GetFileKey() string {
	if o == nil || IsNil(o.FileKey) {
		var ret string
		return ret
	}
	return *o.FileKey
}

// GetFileKeyOk returns a tuple with the FileKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandUploadResponseModel) GetFileKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FileKey) {
		return nil, false
	}
	return o.FileKey, true
}

// HasFileKey returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandUploadResponseModel) HasFileKey() bool {
	if o != nil && !IsNil(o.FileKey) {
		return true
	}

	return false
}

// SetFileKey gets a reference to the given string and assigns it to the FileKey field.
func (o *AlipayOpenMiniMiniappBrandUploadResponseModel) SetFileKey(v string) {
	o.FileKey = &v
}

func (o AlipayOpenMiniMiniappBrandUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniMiniappBrandUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileKey) {
		toSerialize["file_key"] = o.FileKey
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniMiniappBrandUploadResponseModel struct {
	value *AlipayOpenMiniMiniappBrandUploadResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniMiniappBrandUploadResponseModel) Get() *AlipayOpenMiniMiniappBrandUploadResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniMiniappBrandUploadResponseModel) Set(val *AlipayOpenMiniMiniappBrandUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniMiniappBrandUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniMiniappBrandUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniMiniappBrandUploadResponseModel(val *AlipayOpenMiniMiniappBrandUploadResponseModel) *NullableAlipayOpenMiniMiniappBrandUploadResponseModel {
	return &NullableAlipayOpenMiniMiniappBrandUploadResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniMiniappBrandUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniMiniappBrandUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


