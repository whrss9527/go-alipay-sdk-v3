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

// checks if the Participant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Participant{}

// Participant struct for Participant
type Participant struct {
	BankcardExtInfo *BankcardExtInfo `json:"bankcard_ext_info,omitempty"`
	// 参与方的证件号，支持身份证号、护照号。
	CertNo *string `json:"cert_no,omitempty"`
	// 参与方的证件类型。
	CertType *string `json:"cert_type,omitempty"`
	// 描述参与方信息的扩展属性，使用前请与支付宝工程师确认
	ExtInfo *string `json:"ext_info,omitempty"`
	// 参与方的唯一标识
	Identity *string `json:"identity,omitempty"`
	// 参与方的标识类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、ALIPAY_OPEN_ID：支付宝openid
	IdentityType *string `json:"identity_type,omitempty"`
	// 商户端的用户信息描述，目前可指定如下属性：  merchant_user_id：商户端的用户唯一ID  merchant_user_name：商户端的用户名  merchant_user_nickname：商户端的用户昵称  merchant_user_mobile：商户端的手机号
	MerchantUserInfo *string `json:"merchant_user_info,omitempty"`
	// 参与方真实姓名，如果非空，将校验收款支付宝账号姓名一致性。当identity_type=ALIPAY_LOGON_ID时，本字段必填。
	Name *string `json:"name,omitempty"`
}

// NewParticipant instantiates a new Participant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipant() *Participant {
	this := Participant{}
	return &this
}

// NewParticipantWithDefaults instantiates a new Participant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipantWithDefaults() *Participant {
	this := Participant{}
	return &this
}

// GetBankcardExtInfo returns the BankcardExtInfo field value if set, zero value otherwise.
func (o *Participant) GetBankcardExtInfo() BankcardExtInfo {
	if o == nil || IsNil(o.BankcardExtInfo) {
		var ret BankcardExtInfo
		return ret
	}
	return *o.BankcardExtInfo
}

// GetBankcardExtInfoOk returns a tuple with the BankcardExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetBankcardExtInfoOk() (*BankcardExtInfo, bool) {
	if o == nil || IsNil(o.BankcardExtInfo) {
		return nil, false
	}
	return o.BankcardExtInfo, true
}

// HasBankcardExtInfo returns a boolean if a field has been set.
func (o *Participant) HasBankcardExtInfo() bool {
	if o != nil && !IsNil(o.BankcardExtInfo) {
		return true
	}

	return false
}

// SetBankcardExtInfo gets a reference to the given BankcardExtInfo and assigns it to the BankcardExtInfo field.
func (o *Participant) SetBankcardExtInfo(v BankcardExtInfo) {
	o.BankcardExtInfo = &v
}

// GetCertNo returns the CertNo field value if set, zero value otherwise.
func (o *Participant) GetCertNo() string {
	if o == nil || IsNil(o.CertNo) {
		var ret string
		return ret
	}
	return *o.CertNo
}

// GetCertNoOk returns a tuple with the CertNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetCertNoOk() (*string, bool) {
	if o == nil || IsNil(o.CertNo) {
		return nil, false
	}
	return o.CertNo, true
}

// HasCertNo returns a boolean if a field has been set.
func (o *Participant) HasCertNo() bool {
	if o != nil && !IsNil(o.CertNo) {
		return true
	}

	return false
}

// SetCertNo gets a reference to the given string and assigns it to the CertNo field.
func (o *Participant) SetCertNo(v string) {
	o.CertNo = &v
}

// GetCertType returns the CertType field value if set, zero value otherwise.
func (o *Participant) GetCertType() string {
	if o == nil || IsNil(o.CertType) {
		var ret string
		return ret
	}
	return *o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertType) {
		return nil, false
	}
	return o.CertType, true
}

// HasCertType returns a boolean if a field has been set.
func (o *Participant) HasCertType() bool {
	if o != nil && !IsNil(o.CertType) {
		return true
	}

	return false
}

// SetCertType gets a reference to the given string and assigns it to the CertType field.
func (o *Participant) SetCertType(v string) {
	o.CertType = &v
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *Participant) GetExtInfo() string {
	if o == nil || IsNil(o.ExtInfo) {
		var ret string
		return ret
	}
	return *o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetExtInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *Participant) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given string and assigns it to the ExtInfo field.
func (o *Participant) SetExtInfo(v string) {
	o.ExtInfo = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Participant) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Participant) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *Participant) SetIdentity(v string) {
	o.Identity = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *Participant) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *Participant) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *Participant) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetMerchantUserInfo returns the MerchantUserInfo field value if set, zero value otherwise.
func (o *Participant) GetMerchantUserInfo() string {
	if o == nil || IsNil(o.MerchantUserInfo) {
		var ret string
		return ret
	}
	return *o.MerchantUserInfo
}

// GetMerchantUserInfoOk returns a tuple with the MerchantUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetMerchantUserInfoOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantUserInfo) {
		return nil, false
	}
	return o.MerchantUserInfo, true
}

// HasMerchantUserInfo returns a boolean if a field has been set.
func (o *Participant) HasMerchantUserInfo() bool {
	if o != nil && !IsNil(o.MerchantUserInfo) {
		return true
	}

	return false
}

// SetMerchantUserInfo gets a reference to the given string and assigns it to the MerchantUserInfo field.
func (o *Participant) SetMerchantUserInfo(v string) {
	o.MerchantUserInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Participant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Participant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Participant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Participant) SetName(v string) {
	o.Name = &v
}

func (o Participant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Participant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankcardExtInfo) {
		toSerialize["bankcard_ext_info"] = o.BankcardExtInfo
	}
	if !IsNil(o.CertNo) {
		toSerialize["cert_no"] = o.CertNo
	}
	if !IsNil(o.CertType) {
		toSerialize["cert_type"] = o.CertType
	}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.MerchantUserInfo) {
		toSerialize["merchant_user_info"] = o.MerchantUserInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableParticipant struct {
	value *Participant
	isSet bool
}

func (v NullableParticipant) Get() *Participant {
	return v.value
}

func (v *NullableParticipant) Set(val *Participant) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipant) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipant(val *Participant) *NullableParticipant {
	return &NullableParticipant{value: val, isSet: true}
}

func (v NullableParticipant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParticipant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


