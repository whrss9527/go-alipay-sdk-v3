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

// checks if the AuthInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthInfoDTO{}

// AuthInfoDTO struct for AuthInfoDTO
type AuthInfoDTO struct {
	// 授权协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 授权的支付宝账号，脱敏。
	AlipayLogonId *string `json:"alipay_logon_id,omitempty"`
}

// NewAuthInfoDTO instantiates a new AuthInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthInfoDTO() *AuthInfoDTO {
	this := AuthInfoDTO{}
	return &this
}

// NewAuthInfoDTOWithDefaults instantiates a new AuthInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthInfoDTOWithDefaults() *AuthInfoDTO {
	this := AuthInfoDTO{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AuthInfoDTO) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthInfoDTO) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AuthInfoDTO) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AuthInfoDTO) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetAlipayLogonId returns the AlipayLogonId field value if set, zero value otherwise.
func (o *AuthInfoDTO) GetAlipayLogonId() string {
	if o == nil || IsNil(o.AlipayLogonId) {
		var ret string
		return ret
	}
	return *o.AlipayLogonId
}

// GetAlipayLogonIdOk returns a tuple with the AlipayLogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthInfoDTO) GetAlipayLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayLogonId) {
		return nil, false
	}
	return o.AlipayLogonId, true
}

// HasAlipayLogonId returns a boolean if a field has been set.
func (o *AuthInfoDTO) HasAlipayLogonId() bool {
	if o != nil && !IsNil(o.AlipayLogonId) {
		return true
	}

	return false
}

// SetAlipayLogonId gets a reference to the given string and assigns it to the AlipayLogonId field.
func (o *AuthInfoDTO) SetAlipayLogonId(v string) {
	o.AlipayLogonId = &v
}

func (o AuthInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.AlipayLogonId) {
		toSerialize["alipay_logon_id"] = o.AlipayLogonId
	}
	return toSerialize, nil
}

type NullableAuthInfoDTO struct {
	value *AuthInfoDTO
	isSet bool
}

func (v NullableAuthInfoDTO) Get() *AuthInfoDTO {
	return v.value
}

func (v *NullableAuthInfoDTO) Set(val *AuthInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthInfoDTO(val *AuthInfoDTO) *NullableAuthInfoDTO {
	return &NullableAuthInfoDTO{value: val, isSet: true}
}

func (v NullableAuthInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


