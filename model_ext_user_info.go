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

// checks if the ExtUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtUserInfo{}

// ExtUserInfo struct for ExtUserInfo
type ExtUserInfo struct {
	// 买家证件号。 注：need_check_info=T或fix_buyer=T时该参数才有效，支付宝会比较买家在支付宝留存的证件号码与该参数传入的值是否匹配。
	CertNo *string `json:"cert_no,omitempty"`
	// 指定买家证件类型。 枚举值： IDENTITY_CARD：身份证； PASSPORT：护照； OFFICER_CARD：军官证； SOLDIER_CARD：士兵证； HOKOU：户口本； PERMANENT_RESIDENCE_FOREIGNER：外国人永久居留身份证。 如有其它类型需要支持，请与蚂蚁金服工作人员联系。 注： need_check_info=T或fix_buyer=T时该参数才有效，支付宝会比较买家在支付宝留存的证件类型与该参数传入的值是否匹配。
	CertType *string `json:"cert_type,omitempty"`
	// 是否强制校验买家身份。 需要强制校验传：T; 不需要强制校验传：F或者不传； 当传T时，接口上必须指定cert_type、cert_no和name信息且支付宝会校验传入的信息跟支付买家的信息都匹配，否则报错。 默认为不校验。
	FixBuyer *string `json:"fix_buyer,omitempty"`
	// 买家加密身份信息。当指定了此参数且指定need_check_info=T时，支付宝会对买家身份进行校验，校验逻辑为买家姓名、买家证件号拼接后的字符串，以sha256算法utf-8编码计算hash，若与传入的值不匹配则会拦截本次支付。注意：如果同时指定了用户明文身份信息（name，cert_type，cert_no中任意一个），则忽略identity_hash以明文参数校验。
	IdentityHash *string `json:"identity_hash,omitempty"`
	// 允许的最小买家年龄。 买家年龄必须大于等于所传数值  注： 1. need_check_info=T时该参数才有效 2. min_age为整数，必须大于等于0
	MinAge *string `json:"min_age,omitempty"`
	// 指定买家手机号。 注：该参数暂不校验
	Mobile *string `json:"mobile,omitempty"`
	// 指定买家姓名。 注： need_check_info=T或fix_buyer=T时该参数才有效
	Name *string `json:"name,omitempty"`
	// 是否强制校验买家信息； 需要强制校验传：T; 不需要强制校验传：F或者不传； 当传T时，支付宝会校验支付买家的信息与接口上传递的cert_type、cert_no、name或age是否匹配，只有接口传递了信息才会进行对应项的校验；只要有任何一项信息校验不匹配交易都会失败。如果传递了need_check_info，但是没有传任何校验项，则不进行任何校验。 默认为不校验。
	NeedCheckInfo *string `json:"need_check_info,omitempty"`
}

// NewExtUserInfo instantiates a new ExtUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtUserInfo() *ExtUserInfo {
	this := ExtUserInfo{}
	return &this
}

// NewExtUserInfoWithDefaults instantiates a new ExtUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtUserInfoWithDefaults() *ExtUserInfo {
	this := ExtUserInfo{}
	return &this
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *ExtUserInfo) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *ExtUserInfo) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *ExtUserInfo) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *ExtUserInfo) GetCertType() string {
	if o == nil || IsNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *ExtUserInfo) HasCertType() bool {
	if o != nil && !IsNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *ExtUserInfo) SetCertType(v string) {
	o.CertType = &v
}

// GetFixBuyer returns the FixBuyer field value if set, zero value otherwise.
func (o *ExtUserInfo) GetFixBuyer() string {
	if o == nil || IsNil(o.FixBuyer) {
		var ret string
		return ret
	}
	return *o.FixBuyer
}

// GetFixBuyerOk returns a tuple with the FixBuyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetFixBuyerOk() (*string, bool) {
	if o == nil || IsNil(o.FixBuyer) {
		return nil, false
	}
	return o.FixBuyer, true
}

// HasFixBuyer returns a boolean if a field has been set.
func (o *ExtUserInfo) HasFixBuyer() bool {
	if o != nil && !IsNil(o.FixBuyer) {
		return true
	}

	return false
}

// SetFixBuyer gets a reference to the given string and assigns it to the FixBuyer field.
func (o *ExtUserInfo) SetFixBuyer(v string) {
	o.FixBuyer = &v
}

// GetIdentityHash returns the IdentityHash field value if set, zero value otherwise.
func (o *ExtUserInfo) GetIdentityHash() string {
	if o == nil || IsNil(o.IdentityHash) {
		var ret string
		return ret
	}
	return *o.IdentityHash
}

// GetIdentityHashOk returns a tuple with the IdentityHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetIdentityHashOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityHash) {
		return nil, false
	}
	return o.IdentityHash, true
}

// HasIdentityHash returns a boolean if a field has been set.
func (o *ExtUserInfo) HasIdentityHash() bool {
	if o != nil && !IsNil(o.IdentityHash) {
		return true
	}

	return false
}

// SetIdentityHash gets a reference to the given string and assigns it to the IdentityHash field.
func (o *ExtUserInfo) SetIdentityHash(v string) {
	o.IdentityHash = &v
}

// GetMinAge returns the MinAge field value if set, zero value otherwise.
func (o *ExtUserInfo) GetMinAge() string {
	if o == nil || IsNil(o.MinAge) {
		var ret string
		return ret
	}
	return *o.MinAge
}

// GetMinAgeOk returns a tuple with the MinAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetMinAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MinAge) {
		return nil, false
	}
	return o.MinAge, true
}

// HasMinAge returns a boolean if a field has been set.
func (o *ExtUserInfo) HasMinAge() bool {
	if o != nil && !IsNil(o.MinAge) {
		return true
	}

	return false
}

// SetMinAge gets a reference to the given string and assigns it to the MinAge field.
func (o *ExtUserInfo) SetMinAge(v string) {
	o.MinAge = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ExtUserInfo) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ExtUserInfo) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *ExtUserInfo) SetMobile(v string) {
	o.Mobile = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExtUserInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExtUserInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExtUserInfo) SetName(v string) {
	o.Name = &v
}

// GetNeedCheckInfo returns the NeedCheckInfo field value if set, zero value otherwise.
func (o *ExtUserInfo) GetNeedCheckInfo() string {
	if o == nil || IsNil(o.NeedCheckInfo) {
		var ret string
		return ret
	}
	return *o.NeedCheckInfo
}

// GetNeedCheckInfoOk returns a tuple with the NeedCheckInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtUserInfo) GetNeedCheckInfoOk() (*string, bool) {
	if o == nil || IsNil(o.NeedCheckInfo) {
		return nil, false
	}
	return o.NeedCheckInfo, true
}

// HasNeedCheckInfo returns a boolean if a field has been set.
func (o *ExtUserInfo) HasNeedCheckInfo() bool {
	if o != nil && !IsNil(o.NeedCheckInfo) {
		return true
	}

	return false
}

// SetNeedCheckInfo gets a reference to the given string and assigns it to the NeedCheckInfo field.
func (o *ExtUserInfo) SetNeedCheckInfo(v string) {
	o.NeedCheckInfo = &v
}

func (o ExtUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.CertType) {
		toSerialize["cert_type"] = o.CertType
	}
	if !IsNil(o.FixBuyer) {
		toSerialize["fix_buyer"] = o.FixBuyer
	}
	if !IsNil(o.IdentityHash) {
		toSerialize["identity_hash"] = o.IdentityHash
	}
	if !IsNil(o.MinAge) {
		toSerialize["min_age"] = o.MinAge
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NeedCheckInfo) {
		toSerialize["need_check_info"] = o.NeedCheckInfo
	}
	return toSerialize, nil
}

type NullableExtUserInfo struct {
	value *ExtUserInfo
	isSet bool
}

func (v NullableExtUserInfo) Get() *ExtUserInfo {
	return v.value
}

func (v *NullableExtUserInfo) Set(val *ExtUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExtUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExtUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtUserInfo(val *ExtUserInfo) *NullableExtUserInfo {
	return &NullableExtUserInfo{value: val, isSet: true}
}

func (v NullableExtUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
