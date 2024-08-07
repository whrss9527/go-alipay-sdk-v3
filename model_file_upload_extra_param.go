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

// checks if the FileUploadExtraParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUploadExtraParam{}

// FileUploadExtraParam struct for FileUploadExtraParam
type FileUploadExtraParam struct {
	// 是否同步上传到其他平台。素材类型为视频且需要同步优酷，请必传入参数youku，素材接口将代上传到优酷账号；素材类型为图片，无需传入参数
	ExternUpload *string `json:"extern_upload,omitempty"`
	// 文件加密类型
	FileEncryptType *string `json:"file_encrypt_type,omitempty"`
}

// NewFileUploadExtraParam instantiates a new FileUploadExtraParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUploadExtraParam() *FileUploadExtraParam {
	this := FileUploadExtraParam{}
	return &this
}

// NewFileUploadExtraParamWithDefaults instantiates a new FileUploadExtraParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUploadExtraParamWithDefaults() *FileUploadExtraParam {
	this := FileUploadExtraParam{}
	return &this
}

// GetExternUpload returns the ExternUpload field value if set, zero value otherwise.
func (o *FileUploadExtraParam) GetExternUpload() string {
	if o == nil || IsNil(o.ExternUpload) {
		var ret string
		return ret
	}
	return *o.ExternUpload
}

// GetExternUploadOk returns a tuple with the ExternUpload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadExtraParam) GetExternUploadOk() (*string, bool) {
	if o == nil || IsNil(o.ExternUpload) {
		return nil, false
	}
	return o.ExternUpload, true
}

// HasExternUpload returns a boolean if a field has been set.
func (o *FileUploadExtraParam) HasExternUpload() bool {
	if o != nil && !IsNil(o.ExternUpload) {
		return true
	}

	return false
}

// SetExternUpload gets a reference to the given string and assigns it to the ExternUpload field.
func (o *FileUploadExtraParam) SetExternUpload(v string) {
	o.ExternUpload = &v
}

// GetFileEncryptType returns the FileEncryptType field value if set, zero value otherwise.
func (o *FileUploadExtraParam) GetFileEncryptType() string {
	if o == nil || IsNil(o.FileEncryptType) {
		var ret string
		return ret
	}
	return *o.FileEncryptType
}

// GetFileEncryptTypeOk returns a tuple with the FileEncryptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadExtraParam) GetFileEncryptTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileEncryptType) {
		return nil, false
	}
	return o.FileEncryptType, true
}

// HasFileEncryptType returns a boolean if a field has been set.
func (o *FileUploadExtraParam) HasFileEncryptType() bool {
	if o != nil && !IsNil(o.FileEncryptType) {
		return true
	}

	return false
}

// SetFileEncryptType gets a reference to the given string and assigns it to the FileEncryptType field.
func (o *FileUploadExtraParam) SetFileEncryptType(v string) {
	o.FileEncryptType = &v
}

func (o FileUploadExtraParam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUploadExtraParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternUpload) {
		toSerialize["extern_upload"] = o.ExternUpload
	}
	if !IsNil(o.FileEncryptType) {
		toSerialize["file_encrypt_type"] = o.FileEncryptType
	}
	return toSerialize, nil
}

type NullableFileUploadExtraParam struct {
	value *FileUploadExtraParam
	isSet bool
}

func (v NullableFileUploadExtraParam) Get() *FileUploadExtraParam {
	return v.value
}

func (v *NullableFileUploadExtraParam) Set(val *FileUploadExtraParam) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUploadExtraParam) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUploadExtraParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUploadExtraParam(val *FileUploadExtraParam) *NullableFileUploadExtraParam {
	return &NullableFileUploadExtraParam{value: val, isSet: true}
}

func (v NullableFileUploadExtraParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUploadExtraParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
