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

// checks if the AlipayMarketingMaterialImageUploadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingMaterialImageUploadModel{}

// AlipayMarketingMaterialImageUploadModel struct for AlipayMarketingMaterialImageUploadModel
type AlipayMarketingMaterialImageUploadModel struct {
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// 文件业务标识。  枚举值 alipay.marketing.activity.delivery.create接口中 delivery_base_info.delivery_material.delivery_single_material.delivery_image 当delivery_booth_code=PUBLIC_UNION，上传图片接口需指定file_key=PUBLIC_UNION_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb； 当delivery_booth_code=PAYMENT_RESULT，上传图片接口需指定file_key=DELIVERY_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb。 上传图片更多要求参考文档： https://render.alipay.com/p/c/18tpirlg12e8?operateFrom=BALIPAY  alipay.marketing.activity.ordervoucher.create接口中 voucher_display_info.brand_logo字段,file_key=PROMO_BRAND_LOGO，上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB voucher_display_info.voucher_image字段,file_key=PROMO_VOUCHER_IMAGE,上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB
	FileKey *string `json:"file_key,omitempty"`
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
}

// NewAlipayMarketingMaterialImageUploadModel instantiates a new AlipayMarketingMaterialImageUploadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingMaterialImageUploadModel() *AlipayMarketingMaterialImageUploadModel {
	this := AlipayMarketingMaterialImageUploadModel{}
	return &this
}

// NewAlipayMarketingMaterialImageUploadModelWithDefaults instantiates a new AlipayMarketingMaterialImageUploadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingMaterialImageUploadModelWithDefaults() *AlipayMarketingMaterialImageUploadModel {
	this := AlipayMarketingMaterialImageUploadModel{}
	return &this
}

// GetBelongMerchantInfo returns the BelongMerchantInfo field value if set, zero value otherwise.
func (o *AlipayMarketingMaterialImageUploadModel) GetBelongMerchantInfo() BelongMerchantInfo {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		var ret BelongMerchantInfo
		return ret
	}
	return *o.BelongMerchantInfo
}

// GetBelongMerchantInfoOk returns a tuple with the BelongMerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingMaterialImageUploadModel) GetBelongMerchantInfoOk() (*BelongMerchantInfo, bool) {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		return nil, false
	}
	return o.BelongMerchantInfo, true
}

// HasBelongMerchantInfo returns a boolean if a field has been set.
func (o *AlipayMarketingMaterialImageUploadModel) HasBelongMerchantInfo() bool {
	if o != nil && !IsNil(o.BelongMerchantInfo) {
		return true
	}

	return false
}

// SetBelongMerchantInfo gets a reference to the given BelongMerchantInfo and assigns it to the BelongMerchantInfo field.
func (o *AlipayMarketingMaterialImageUploadModel) SetBelongMerchantInfo(v BelongMerchantInfo) {
	o.BelongMerchantInfo = &v
}

// GetFileKey returns the FileKey field value if set, zero value otherwise.
func (o *AlipayMarketingMaterialImageUploadModel) GetFileKey() string {
	if o == nil || IsNil(o.FileKey) {
		var ret string
		return ret
	}
	return *o.FileKey
}

// GetFileKeyOk returns a tuple with the FileKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingMaterialImageUploadModel) GetFileKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FileKey) {
		return nil, false
	}
	return o.FileKey, true
}

// HasFileKey returns a boolean if a field has been set.
func (o *AlipayMarketingMaterialImageUploadModel) HasFileKey() bool {
	if o != nil && !IsNil(o.FileKey) {
		return true
	}

	return false
}

// SetFileKey gets a reference to the given string and assigns it to the FileKey field.
func (o *AlipayMarketingMaterialImageUploadModel) SetFileKey(v string) {
	o.FileKey = &v
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingMaterialImageUploadModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingMaterialImageUploadModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingMaterialImageUploadModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingMaterialImageUploadModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

func (o AlipayMarketingMaterialImageUploadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingMaterialImageUploadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BelongMerchantInfo) {
		toSerialize["belong_merchant_info"] = o.BelongMerchantInfo
	}
	if !IsNil(o.FileKey) {
		toSerialize["file_key"] = o.FileKey
	}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	return toSerialize, nil
}

type NullableAlipayMarketingMaterialImageUploadModel struct {
	value *AlipayMarketingMaterialImageUploadModel
	isSet bool
}

func (v NullableAlipayMarketingMaterialImageUploadModel) Get() *AlipayMarketingMaterialImageUploadModel {
	return v.value
}

func (v *NullableAlipayMarketingMaterialImageUploadModel) Set(val *AlipayMarketingMaterialImageUploadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingMaterialImageUploadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingMaterialImageUploadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingMaterialImageUploadModel(val *AlipayMarketingMaterialImageUploadModel) *NullableAlipayMarketingMaterialImageUploadModel {
	return &NullableAlipayMarketingMaterialImageUploadModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingMaterialImageUploadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingMaterialImageUploadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
