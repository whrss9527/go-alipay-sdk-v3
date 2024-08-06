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

// checks if the AlipayUserCertifyOpenInitializeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserCertifyOpenInitializeModel{}

// AlipayUserCertifyOpenInitializeModel struct for AlipayUserCertifyOpenInitializeModel
type AlipayUserCertifyOpenInitializeModel struct {
	// 认证场景码。入参支持的认证场景码和商户签约的认证场景相关，取值如下: FACE：多因子人脸认证 CERT_PHOTO：多因子证照认证 CERT_PHOTO_FACE ：多因子证照和人脸认证 SMART_FACE：多因子快捷认证
	BizCode *string `json:"biz_code,omitempty"`
	// 自定义人脸比对图片的base64编码格式的string字符串
	FaceContrastPicture *string `json:"face_contrast_picture,omitempty"`
	IdentityParam *OpenCertifyIdentityParam `json:"identity_param,omitempty"`
	MerchantConfig *OpenCertifyMerchantConfig `json:"merchant_config,omitempty"`
	// 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	OuterOrderNo *string `json:"outer_order_no,omitempty"`
}

// NewAlipayUserCertifyOpenInitializeModel instantiates a new AlipayUserCertifyOpenInitializeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserCertifyOpenInitializeModel() *AlipayUserCertifyOpenInitializeModel {
	this := AlipayUserCertifyOpenInitializeModel{}
	return &this
}

// NewAlipayUserCertifyOpenInitializeModelWithDefaults instantiates a new AlipayUserCertifyOpenInitializeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserCertifyOpenInitializeModelWithDefaults() *AlipayUserCertifyOpenInitializeModel {
	this := AlipayUserCertifyOpenInitializeModel{}
	return &this
}

// GetBizCode returns the BizCode field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeModel) GetBizCode() string {
	if o == nil || IsNil(o.BizCode) {
		var ret string
		return ret
	}
	return *o.BizCode
}

// GetBizCodeOk returns a tuple with the BizCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeModel) GetBizCodeOk() (*string, bool) {
	if o == nil || IsNil(o.BizCode) {
		return nil, false
	}
	return o.BizCode, true
}

// HasBizCode returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeModel) HasBizCode() bool {
	if o != nil && !IsNil(o.BizCode) {
		return true
	}

	return false
}

// SetBizCode gets a reference to the given string and assigns it to the BizCode field.
func (o *AlipayUserCertifyOpenInitializeModel) SetBizCode(v string) {
	o.BizCode = &v
}

// GetFaceContrastPicture returns the FaceContrastPicture field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeModel) GetFaceContrastPicture() string {
	if o == nil || IsNil(o.FaceContrastPicture) {
		var ret string
		return ret
	}
	return *o.FaceContrastPicture
}

// GetFaceContrastPictureOk returns a tuple with the FaceContrastPicture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeModel) GetFaceContrastPictureOk() (*string, bool) {
	if o == nil || IsNil(o.FaceContrastPicture) {
		return nil, false
	}
	return o.FaceContrastPicture, true
}

// HasFaceContrastPicture returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeModel) HasFaceContrastPicture() bool {
	if o != nil && !IsNil(o.FaceContrastPicture) {
		return true
	}

	return false
}

// SetFaceContrastPicture gets a reference to the given string and assigns it to the FaceContrastPicture field.
func (o *AlipayUserCertifyOpenInitializeModel) SetFaceContrastPicture(v string) {
	o.FaceContrastPicture = &v
}

// GetIdentityParam returns the IdentityParam field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeModel) GetIdentityParam() OpenCertifyIdentityParam {
	if o == nil || IsNil(o.IdentityParam) {
		var ret OpenCertifyIdentityParam
		return ret
	}
	return *o.IdentityParam
}

// GetIdentityParamOk returns a tuple with the IdentityParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeModel) GetIdentityParamOk() (*OpenCertifyIdentityParam, bool) {
	if o == nil || IsNil(o.IdentityParam) {
		return nil, false
	}
	return o.IdentityParam, true
}

// HasIdentityParam returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeModel) HasIdentityParam() bool {
	if o != nil && !IsNil(o.IdentityParam) {
		return true
	}

	return false
}

// SetIdentityParam gets a reference to the given OpenCertifyIdentityParam and assigns it to the IdentityParam field.
func (o *AlipayUserCertifyOpenInitializeModel) SetIdentityParam(v OpenCertifyIdentityParam) {
	o.IdentityParam = &v
}

// GetMerchantConfig returns the MerchantConfig field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeModel) GetMerchantConfig() OpenCertifyMerchantConfig {
	if o == nil || IsNil(o.MerchantConfig) {
		var ret OpenCertifyMerchantConfig
		return ret
	}
	return *o.MerchantConfig
}

// GetMerchantConfigOk returns a tuple with the MerchantConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeModel) GetMerchantConfigOk() (*OpenCertifyMerchantConfig, bool) {
	if o == nil || IsNil(o.MerchantConfig) {
		return nil, false
	}
	return o.MerchantConfig, true
}

// HasMerchantConfig returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeModel) HasMerchantConfig() bool {
	if o != nil && !IsNil(o.MerchantConfig) {
		return true
	}

	return false
}

// SetMerchantConfig gets a reference to the given OpenCertifyMerchantConfig and assigns it to the MerchantConfig field.
func (o *AlipayUserCertifyOpenInitializeModel) SetMerchantConfig(v OpenCertifyMerchantConfig) {
	o.MerchantConfig = &v
}

// GetOuterOrderNo returns the OuterOrderNo field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeModel) GetOuterOrderNo() string {
	if o == nil || IsNil(o.OuterOrderNo) {
		var ret string
		return ret
	}
	return *o.OuterOrderNo
}

// GetOuterOrderNoOk returns a tuple with the OuterOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeModel) GetOuterOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OuterOrderNo) {
		return nil, false
	}
	return o.OuterOrderNo, true
}

// HasOuterOrderNo returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeModel) HasOuterOrderNo() bool {
	if o != nil && !IsNil(o.OuterOrderNo) {
		return true
	}

	return false
}

// SetOuterOrderNo gets a reference to the given string and assigns it to the OuterOrderNo field.
func (o *AlipayUserCertifyOpenInitializeModel) SetOuterOrderNo(v string) {
	o.OuterOrderNo = &v
}

func (o AlipayUserCertifyOpenInitializeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserCertifyOpenInitializeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizCode) {
		toSerialize["biz_code"] = o.BizCode
	}
	if !IsNil(o.FaceContrastPicture) {
		toSerialize["face_contrast_picture"] = o.FaceContrastPicture
	}
	if !IsNil(o.IdentityParam) {
		toSerialize["identity_param"] = o.IdentityParam
	}
	if !IsNil(o.MerchantConfig) {
		toSerialize["merchant_config"] = o.MerchantConfig
	}
	if !IsNil(o.OuterOrderNo) {
		toSerialize["outer_order_no"] = o.OuterOrderNo
	}
	return toSerialize, nil
}

type NullableAlipayUserCertifyOpenInitializeModel struct {
	value *AlipayUserCertifyOpenInitializeModel
	isSet bool
}

func (v NullableAlipayUserCertifyOpenInitializeModel) Get() *AlipayUserCertifyOpenInitializeModel {
	return v.value
}

func (v *NullableAlipayUserCertifyOpenInitializeModel) Set(val *AlipayUserCertifyOpenInitializeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserCertifyOpenInitializeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserCertifyOpenInitializeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserCertifyOpenInitializeModel(val *AlipayUserCertifyOpenInitializeModel) *NullableAlipayUserCertifyOpenInitializeModel {
	return &NullableAlipayUserCertifyOpenInitializeModel{value: val, isSet: true}
}

func (v NullableAlipayUserCertifyOpenInitializeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserCertifyOpenInitializeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


