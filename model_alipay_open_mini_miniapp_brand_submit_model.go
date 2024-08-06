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

// checks if the AlipayOpenMiniMiniappBrandSubmitModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniMiniappBrandSubmitModel{}

// AlipayOpenMiniMiniappBrandSubmitModel struct for AlipayOpenMiniMiniappBrandSubmitModel
type AlipayOpenMiniMiniappBrandSubmitModel struct {
	// 申请的资质类型 SELF_BRAND_APPLY 自有品牌申请(当需要一个新的品牌时选择此项) AUTHORIZE_BRAND_APPLY 授权品牌申请(当需要授权一个已有的品牌给商户时选择此项) EXIST_BRAND_BIND 品牌绑定，适用于品牌已申请成功的情况(当需要将已有品牌绑定到无品牌小程序时选择此项)
	ApplyType     *string        `json:"apply_type,omitempty"`
	AuthorizeInfo *AuthorizeInfo `json:"authorize_info,omitempty"`
	// 商户已有品牌Id(只有申请的资质类型为EXIST_BRAND_BIND 时需要填写)
	BrandId *string `json:"brand_id,omitempty"`
	// 品牌名称
	BrandName             *string                `json:"brand_name,omitempty"`
	BrandRegistrationInfo *BrandRegistrationInfo `json:"brand_registration_info,omitempty"`
	// 申请人身份证明材料(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片，可以上传多个身份证明材料)
	IdMaterials []string `json:"id_materials,omitempty"`
}

// NewAlipayOpenMiniMiniappBrandSubmitModel instantiates a new AlipayOpenMiniMiniappBrandSubmitModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniMiniappBrandSubmitModel() *AlipayOpenMiniMiniappBrandSubmitModel {
	this := AlipayOpenMiniMiniappBrandSubmitModel{}
	return &this
}

// NewAlipayOpenMiniMiniappBrandSubmitModelWithDefaults instantiates a new AlipayOpenMiniMiniappBrandSubmitModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniMiniappBrandSubmitModelWithDefaults() *AlipayOpenMiniMiniappBrandSubmitModel {
	this := AlipayOpenMiniMiniappBrandSubmitModel{}
	return &this
}

// GetApplyType returns the ApplyType field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetApplyType() string {
	if o == nil || IsNil(o.ApplyType) {
		var ret string
		return ret
	}
	return *o.ApplyType
}

// GetApplyTypeOk returns a tuple with the ApplyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetApplyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyType) {
		return nil, false
	}
	return o.ApplyType, true
}

// HasApplyType returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasApplyType() bool {
	if o != nil && !IsNil(o.ApplyType) {
		return true
	}

	return false
}

// SetApplyType gets a reference to the given string and assigns it to the ApplyType field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetApplyType(v string) {
	o.ApplyType = &v
}

// GetAuthorizeInfo returns the AuthorizeInfo field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetAuthorizeInfo() AuthorizeInfo {
	if o == nil || IsNil(o.AuthorizeInfo) {
		var ret AuthorizeInfo
		return ret
	}
	return *o.AuthorizeInfo
}

// GetAuthorizeInfoOk returns a tuple with the AuthorizeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetAuthorizeInfoOk() (*AuthorizeInfo, bool) {
	if o == nil || IsNil(o.AuthorizeInfo) {
		return nil, false
	}
	return o.AuthorizeInfo, true
}

// HasAuthorizeInfo returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasAuthorizeInfo() bool {
	if o != nil && !IsNil(o.AuthorizeInfo) {
		return true
	}

	return false
}

// SetAuthorizeInfo gets a reference to the given AuthorizeInfo and assigns it to the AuthorizeInfo field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetAuthorizeInfo(v AuthorizeInfo) {
	o.AuthorizeInfo = &v
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetBrandId(v string) {
	o.BrandId = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetBrandName(v string) {
	o.BrandName = &v
}

// GetBrandRegistrationInfo returns the BrandRegistrationInfo field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandRegistrationInfo() BrandRegistrationInfo {
	if o == nil || IsNil(o.BrandRegistrationInfo) {
		var ret BrandRegistrationInfo
		return ret
	}
	return *o.BrandRegistrationInfo
}

// GetBrandRegistrationInfoOk returns a tuple with the BrandRegistrationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetBrandRegistrationInfoOk() (*BrandRegistrationInfo, bool) {
	if o == nil || IsNil(o.BrandRegistrationInfo) {
		return nil, false
	}
	return o.BrandRegistrationInfo, true
}

// HasBrandRegistrationInfo returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasBrandRegistrationInfo() bool {
	if o != nil && !IsNil(o.BrandRegistrationInfo) {
		return true
	}

	return false
}

// SetBrandRegistrationInfo gets a reference to the given BrandRegistrationInfo and assigns it to the BrandRegistrationInfo field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetBrandRegistrationInfo(v BrandRegistrationInfo) {
	o.BrandRegistrationInfo = &v
}

// GetIdMaterials returns the IdMaterials field value if set, zero value otherwise.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetIdMaterials() []string {
	if o == nil || IsNil(o.IdMaterials) {
		var ret []string
		return ret
	}
	return o.IdMaterials
}

// GetIdMaterialsOk returns a tuple with the IdMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) GetIdMaterialsOk() ([]string, bool) {
	if o == nil || IsNil(o.IdMaterials) {
		return nil, false
	}
	return o.IdMaterials, true
}

// HasIdMaterials returns a boolean if a field has been set.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) HasIdMaterials() bool {
	if o != nil && !IsNil(o.IdMaterials) {
		return true
	}

	return false
}

// SetIdMaterials gets a reference to the given []string and assigns it to the IdMaterials field.
func (o *AlipayOpenMiniMiniappBrandSubmitModel) SetIdMaterials(v []string) {
	o.IdMaterials = v
}

func (o AlipayOpenMiniMiniappBrandSubmitModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniMiniappBrandSubmitModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyType) {
		toSerialize["apply_type"] = o.ApplyType
	}
	if !IsNil(o.AuthorizeInfo) {
		toSerialize["authorize_info"] = o.AuthorizeInfo
	}
	if !IsNil(o.BrandId) {
		toSerialize["brand_id"] = o.BrandId
	}
	if !IsNil(o.BrandName) {
		toSerialize["brand_name"] = o.BrandName
	}
	if !IsNil(o.BrandRegistrationInfo) {
		toSerialize["brand_registration_info"] = o.BrandRegistrationInfo
	}
	if !IsNil(o.IdMaterials) {
		toSerialize["id_materials"] = o.IdMaterials
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniMiniappBrandSubmitModel struct {
	value *AlipayOpenMiniMiniappBrandSubmitModel
	isSet bool
}

func (v NullableAlipayOpenMiniMiniappBrandSubmitModel) Get() *AlipayOpenMiniMiniappBrandSubmitModel {
	return v.value
}

func (v *NullableAlipayOpenMiniMiniappBrandSubmitModel) Set(val *AlipayOpenMiniMiniappBrandSubmitModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniMiniappBrandSubmitModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniMiniappBrandSubmitModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniMiniappBrandSubmitModel(val *AlipayOpenMiniMiniappBrandSubmitModel) *NullableAlipayOpenMiniMiniappBrandSubmitModel {
	return &NullableAlipayOpenMiniMiniappBrandSubmitModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniMiniappBrandSubmitModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniMiniappBrandSubmitModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
