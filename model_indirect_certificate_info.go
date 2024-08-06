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

// checks if the IndirectCertificateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndirectCertificateInfo{}

// IndirectCertificateInfo struct for IndirectCertificateInfo
type IndirectCertificateInfo struct {
	// 证照图片（使用图片上传接口）
	CertImage *string `json:"cert_image,omitempty"`
	// 证件编号
	CertNo *string `json:"cert_no,omitempty"`
	// 登记证书类型时必填，枚举：统一社会信用代码证书（UNIT_SOCIAL_CREDIT）、慈善组织公开募捐资格证书（CHARITY_ORG_SOLICIT），社会团体法人登记证书（SOCIAL_ORG_LEGAL），民办非企业单位登记证书（CIVIL_UN_ENT），基金会法人登记证书（FOUNDATION_LEGAL_PERSON），农民专业合作社法人营业执照（FARMERS_COOPERATE），宗教活动场所登记证（RELIGION_PLACES），其他证书/批文/证明（OTHER_REG_CERT）
	CertType *string `json:"cert_type,omitempty"`
	// 证照生效时间
	EffectTime *string `json:"effect_time,omitempty"`
	// 证照过期时间
	ExpireTime *string `json:"expire_time,omitempty"`
	// 证照法人姓名
	LegalPersonName *string `json:"legal_person_name,omitempty"`
	// 证照商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// 证照注册地址
	RegisterAddress *string `json:"register_address,omitempty"`
}

// NewIndirectCertificateInfo instantiates a new IndirectCertificateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndirectCertificateInfo() *IndirectCertificateInfo {
	this := IndirectCertificateInfo{}
	return &this
}

// NewIndirectCertificateInfoWithDefaults instantiates a new IndirectCertificateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndirectCertificateInfoWithDefaults() *IndirectCertificateInfo {
	this := IndirectCertificateInfo{}
	return &this
}

// GetCertImage returns the CertImage field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetCertImage() string {
	if o == nil || IsNil(o.CertImage) {
		var ret string
		return ret
	}
	return *o.CertImage
}

// GetCertImageOk returns a tuple with the CertImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetCertImageOk() (*string, bool) {
	if o == nil || IsNil(o.CertImage) {
		return nil, false
	}
	return o.CertImage, true
}

// HasCertImage returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasCertImage() bool {
	if o != nil && !IsNil(o.CertImage) {
		return true
	}

	return false
}

// SetCertImage gets a reference to the given string and assigns it to the CertImage field.
func (o *IndirectCertificateInfo) SetCertImage(v string) {
	o.CertImage = &v
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *IndirectCertificateInfo) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetCertType() string {
	if o == nil || IsNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasCertType() bool {
	if o != nil && !IsNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *IndirectCertificateInfo) SetCertType(v string) {
	o.CertType = &v
}

// GetEffectTime returns the EffectTime field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetEffectTime() string {
	if o == nil || IsNil(o.EffectTime) {
		var ret string
		return ret
	}
	return *o.EffectTime
}

// GetEffectTimeOk returns a tuple with the EffectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetEffectTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectTime) {
		return nil, false
	}
	return o.EffectTime, true
}

// HasEffectTime returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasEffectTime() bool {
	if o != nil && !IsNil(o.EffectTime) {
		return true
	}

	return false
}

// SetEffectTime gets a reference to the given string and assigns it to the EffectTime field.
func (o *IndirectCertificateInfo) SetEffectTime(v string) {
	o.EffectTime = &v
}

// GetExpireTime returns the ExpireTime field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetExpireTime() string {
	if o == nil || IsNil(o.ExpireTime) {
		var ret string
		return ret
	}
	return *o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetExpireTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireTime) {
		return nil, false
	}
	return o.ExpireTime, true
}

// HasExpireTime returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasExpireTime() bool {
	if o != nil && !IsNil(o.ExpireTime) {
		return true
	}

	return false
}

// SetExpireTime gets a reference to the given string and assigns it to the ExpireTime field.
func (o *IndirectCertificateInfo) SetExpireTime(v string) {
	o.ExpireTime = &v
}

// GetLegalPersonName returns the LegalPersonName field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetLegalPersonName() string {
	if o == nil || IsNil(o.LegalPersonName) {
		var ret string
		return ret
	}
	return *o.LegalPersonName
}

// GetLegalPersonNameOk returns a tuple with the LegalPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetLegalPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalPersonName) {
		return nil, false
	}
	return o.LegalPersonName, true
}

// HasLegalPersonName returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasLegalPersonName() bool {
	if o != nil && !IsNil(o.LegalPersonName) {
		return true
	}

	return false
}

// SetLegalPersonName gets a reference to the given string and assigns it to the LegalPersonName field.
func (o *IndirectCertificateInfo) SetLegalPersonName(v string) {
	o.LegalPersonName = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetMerchantName() string {
	if o == nil || IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetMerchantNameOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasMerchantName() bool {
	if o != nil && !IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *IndirectCertificateInfo) SetMerchantName(v string) {
	o.MerchantName = &v
}

// GetRegisterAddress returns the RegisterAddress field value if set, zero value otherwise.
func (o *IndirectCertificateInfo) GetRegisterAddress() string {
	if o == nil || IsNil(o.RegisterAddress) {
		var ret string
		return ret
	}
	return *o.RegisterAddress
}

// GetRegisterAddressOk returns a tuple with the RegisterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectCertificateInfo) GetRegisterAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RegisterAddress) {
		return nil, false
	}
	return o.RegisterAddress, true
}

// HasRegisterAddress returns a boolean if a field has been set.
func (o *IndirectCertificateInfo) HasRegisterAddress() bool {
	if o != nil && !IsNil(o.RegisterAddress) {
		return true
	}

	return false
}

// SetRegisterAddress gets a reference to the given string and assigns it to the RegisterAddress field.
func (o *IndirectCertificateInfo) SetRegisterAddress(v string) {
	o.RegisterAddress = &v
}

func (o IndirectCertificateInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndirectCertificateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertImage) {
		toSerialize["cert_image"] = o.CertImage
	}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.CertType) {
		toSerialize["cert_type"] = o.CertType
	}
	if !IsNil(o.EffectTime) {
		toSerialize["effect_time"] = o.EffectTime
	}
	if !IsNil(o.ExpireTime) {
		toSerialize["expire_time"] = o.ExpireTime
	}
	if !IsNil(o.LegalPersonName) {
		toSerialize["legal_person_name"] = o.LegalPersonName
	}
	if !IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	if !IsNil(o.RegisterAddress) {
		toSerialize["register_address"] = o.RegisterAddress
	}
	return toSerialize, nil
}

type NullableIndirectCertificateInfo struct {
	value *IndirectCertificateInfo
	isSet bool
}

func (v NullableIndirectCertificateInfo) Get() *IndirectCertificateInfo {
	return v.value
}

func (v *NullableIndirectCertificateInfo) Set(val *IndirectCertificateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIndirectCertificateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIndirectCertificateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndirectCertificateInfo(val *IndirectCertificateInfo) *NullableIndirectCertificateInfo {
	return &NullableIndirectCertificateInfo{value: val, isSet: true}
}

func (v NullableIndirectCertificateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndirectCertificateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
