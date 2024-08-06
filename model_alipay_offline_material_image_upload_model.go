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

// checks if the AlipayOfflineMaterialImageUploadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOfflineMaterialImageUploadModel{}

// AlipayOfflineMaterialImageUploadModel struct for AlipayOfflineMaterialImageUploadModel
type AlipayOfflineMaterialImageUploadModel struct {
	// 图片/视频名称
	ImageName *string `json:"image_name,omitempty"`
	// 用于显示指定图片/视频所属的partnerId（支付宝内部使用，外部商户无需填写此字段）
	ImagePid *string `json:"image_pid,omitempty"`
	// 图片/视频格式
	ImageType *string `json:"image_type,omitempty"`
}

// NewAlipayOfflineMaterialImageUploadModel instantiates a new AlipayOfflineMaterialImageUploadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOfflineMaterialImageUploadModel() *AlipayOfflineMaterialImageUploadModel {
	this := AlipayOfflineMaterialImageUploadModel{}
	return &this
}

// NewAlipayOfflineMaterialImageUploadModelWithDefaults instantiates a new AlipayOfflineMaterialImageUploadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOfflineMaterialImageUploadModelWithDefaults() *AlipayOfflineMaterialImageUploadModel {
	this := AlipayOfflineMaterialImageUploadModel{}
	return &this
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *AlipayOfflineMaterialImageUploadModel) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMaterialImageUploadModel) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *AlipayOfflineMaterialImageUploadModel) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *AlipayOfflineMaterialImageUploadModel) SetImageName(v string) {
	o.ImageName = &v
}

// GetImagePid returns the ImagePid field value if set, zero value otherwise.
func (o *AlipayOfflineMaterialImageUploadModel) GetImagePid() string {
	if o == nil || IsNil(o.ImagePid) {
		var ret string
		return ret
	}
	return *o.ImagePid
}

// GetImagePidOk returns a tuple with the ImagePid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMaterialImageUploadModel) GetImagePidOk() (*string, bool) {
	if o == nil || IsNil(o.ImagePid) {
		return nil, false
	}
	return o.ImagePid, true
}

// HasImagePid returns a boolean if a field has been set.
func (o *AlipayOfflineMaterialImageUploadModel) HasImagePid() bool {
	if o != nil && !IsNil(o.ImagePid) {
		return true
	}

	return false
}

// SetImagePid gets a reference to the given string and assigns it to the ImagePid field.
func (o *AlipayOfflineMaterialImageUploadModel) SetImagePid(v string) {
	o.ImagePid = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *AlipayOfflineMaterialImageUploadModel) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMaterialImageUploadModel) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *AlipayOfflineMaterialImageUploadModel) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *AlipayOfflineMaterialImageUploadModel) SetImageType(v string) {
	o.ImageType = &v
}

func (o AlipayOfflineMaterialImageUploadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOfflineMaterialImageUploadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !IsNil(o.ImagePid) {
		toSerialize["image_pid"] = o.ImagePid
	}
	if !IsNil(o.ImageType) {
		toSerialize["image_type"] = o.ImageType
	}
	return toSerialize, nil
}

type NullableAlipayOfflineMaterialImageUploadModel struct {
	value *AlipayOfflineMaterialImageUploadModel
	isSet bool
}

func (v NullableAlipayOfflineMaterialImageUploadModel) Get() *AlipayOfflineMaterialImageUploadModel {
	return v.value
}

func (v *NullableAlipayOfflineMaterialImageUploadModel) Set(val *AlipayOfflineMaterialImageUploadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMaterialImageUploadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMaterialImageUploadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMaterialImageUploadModel(val *AlipayOfflineMaterialImageUploadModel) *NullableAlipayOfflineMaterialImageUploadModel {
	return &NullableAlipayOfflineMaterialImageUploadModel{value: val, isSet: true}
}

func (v NullableAlipayOfflineMaterialImageUploadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMaterialImageUploadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
