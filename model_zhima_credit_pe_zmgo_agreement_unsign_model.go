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

// checks if the ZhimaCreditPeZmgoAgreementUnsignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPeZmgoAgreementUnsignModel{}

// ZhimaCreditPeZmgoAgreementUnsignModel struct for ZhimaCreditPeZmgoAgreementUnsignModel
type ZhimaCreditPeZmgoAgreementUnsignModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号。
	AgreementId *string `json:"agreement_id,omitempty"`
	// 支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// 支付宝的用户id
	OpenId *string `json:"open_id,omitempty"`
	// 商户ID
	PartnerId *string `json:"partner_id,omitempty"`
	// quit_type为USER_CANCEL_QUIT或者SETTLE_APPLY_QUIT
	QuitType *string `json:"quit_type,omitempty"`
}

// NewZhimaCreditPeZmgoAgreementUnsignModel instantiates a new ZhimaCreditPeZmgoAgreementUnsignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPeZmgoAgreementUnsignModel() *ZhimaCreditPeZmgoAgreementUnsignModel {
	this := ZhimaCreditPeZmgoAgreementUnsignModel{}
	return &this
}

// NewZhimaCreditPeZmgoAgreementUnsignModelWithDefaults instantiates a new ZhimaCreditPeZmgoAgreementUnsignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPeZmgoAgreementUnsignModelWithDefaults() *ZhimaCreditPeZmgoAgreementUnsignModel {
	this := ZhimaCreditPeZmgoAgreementUnsignModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetQuitType returns the QuitType field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetQuitType() string {
	if o == nil || IsNil(o.QuitType) {
		var ret string
		return ret
	}
	return *o.QuitType
}

// GetQuitTypeOk returns a tuple with the QuitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) GetQuitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuitType) {
		return nil, false
	}
	return o.QuitType, true
}

// HasQuitType returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) HasQuitType() bool {
	if o != nil && !IsNil(o.QuitType) {
		return true
	}

	return false
}

// SetQuitType gets a reference to the given string and assigns it to the QuitType field.
func (o *ZhimaCreditPeZmgoAgreementUnsignModel) SetQuitType(v string) {
	o.QuitType = &v
}

func (o ZhimaCreditPeZmgoAgreementUnsignModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPeZmgoAgreementUnsignModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.QuitType) {
		toSerialize["quit_type"] = o.QuitType
	}
	return toSerialize, nil
}

type NullableZhimaCreditPeZmgoAgreementUnsignModel struct {
	value *ZhimaCreditPeZmgoAgreementUnsignModel
	isSet bool
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignModel) Get() *ZhimaCreditPeZmgoAgreementUnsignModel {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignModel) Set(val *ZhimaCreditPeZmgoAgreementUnsignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoAgreementUnsignModel(val *ZhimaCreditPeZmgoAgreementUnsignModel) *NullableZhimaCreditPeZmgoAgreementUnsignModel {
	return &NullableZhimaCreditPeZmgoAgreementUnsignModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoAgreementUnsignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoAgreementUnsignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
