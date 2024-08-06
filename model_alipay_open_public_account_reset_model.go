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

// checks if the AlipayOpenPublicAccountResetModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicAccountResetModel{}

// AlipayOpenPublicAccountResetModel struct for AlipayOpenPublicAccountResetModel
type AlipayOpenPublicAccountResetModel struct {
	// 需要重置的协议号，商户会员在支付宝生活号账号中的唯一标识。
	AgreementId *string `json:"agreement_id,omitempty"`
	// 绑定帐号，建议在开发者的系统中保持唯一性
	BindAccountNo *string `json:"bind_account_no,omitempty"`
	// 商户期望在生活号首页看到的关于该用户的显示信息，最长10个字符。 
	DisplayName *string `json:"display_name,omitempty"`
	// 要绑定的商户会员对应的支付宝唯一标识，2088开头长度为16位的字符串。 
	FromUserId *string `json:"from_user_id,omitempty"`
	// 支付宝用户的唯一标识
	OpenId *string `json:"open_id,omitempty"`
	// 要绑定的商户会员的真实姓名，最长10个汉字
	RealName *string `json:"real_name,omitempty"`
	// 备注信息，开发者可以通过该字段纪录其他的额外信息
	Remark *string `json:"remark,omitempty"`
}

// NewAlipayOpenPublicAccountResetModel instantiates a new AlipayOpenPublicAccountResetModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicAccountResetModel() *AlipayOpenPublicAccountResetModel {
	this := AlipayOpenPublicAccountResetModel{}
	return &this
}

// NewAlipayOpenPublicAccountResetModelWithDefaults instantiates a new AlipayOpenPublicAccountResetModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicAccountResetModelWithDefaults() *AlipayOpenPublicAccountResetModel {
	this := AlipayOpenPublicAccountResetModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayOpenPublicAccountResetModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetBindAccountNo returns the BindAccountNo field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetBindAccountNo() string {
	if o == nil || IsNil(o.BindAccountNo) {
		var ret string
		return ret
	}
	return *o.BindAccountNo
}

// GetBindAccountNoOk returns a tuple with the BindAccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetBindAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.BindAccountNo) {
		return nil, false
	}
	return o.BindAccountNo, true
}

// HasBindAccountNo returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasBindAccountNo() bool {
	if o != nil && !IsNil(o.BindAccountNo) {
		return true
	}

	return false
}

// SetBindAccountNo gets a reference to the given string and assigns it to the BindAccountNo field.
func (o *AlipayOpenPublicAccountResetModel) SetBindAccountNo(v string) {
	o.BindAccountNo = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AlipayOpenPublicAccountResetModel) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFromUserId returns the FromUserId field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetFromUserId() string {
	if o == nil || IsNil(o.FromUserId) {
		var ret string
		return ret
	}
	return *o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetFromUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.FromUserId) {
		return nil, false
	}
	return o.FromUserId, true
}

// HasFromUserId returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasFromUserId() bool {
	if o != nil && !IsNil(o.FromUserId) {
		return true
	}

	return false
}

// SetFromUserId gets a reference to the given string and assigns it to the FromUserId field.
func (o *AlipayOpenPublicAccountResetModel) SetFromUserId(v string) {
	o.FromUserId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayOpenPublicAccountResetModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetRealName() string {
	if o == nil || IsNil(o.RealName) {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetRealNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealName) {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasRealName() bool {
	if o != nil && !IsNil(o.RealName) {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *AlipayOpenPublicAccountResetModel) SetRealName(v string) {
	o.RealName = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AlipayOpenPublicAccountResetModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAccountResetModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AlipayOpenPublicAccountResetModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AlipayOpenPublicAccountResetModel) SetRemark(v string) {
	o.Remark = &v
}

func (o AlipayOpenPublicAccountResetModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicAccountResetModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.BindAccountNo) {
		toSerialize["bind_account_no"] = o.BindAccountNo
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.FromUserId) {
		toSerialize["from_user_id"] = o.FromUserId
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.RealName) {
		toSerialize["real_name"] = o.RealName
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicAccountResetModel struct {
	value *AlipayOpenPublicAccountResetModel
	isSet bool
}

func (v NullableAlipayOpenPublicAccountResetModel) Get() *AlipayOpenPublicAccountResetModel {
	return v.value
}

func (v *NullableAlipayOpenPublicAccountResetModel) Set(val *AlipayOpenPublicAccountResetModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAccountResetModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAccountResetModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAccountResetModel(val *AlipayOpenPublicAccountResetModel) *NullableAlipayOpenPublicAccountResetModel {
	return &NullableAlipayOpenPublicAccountResetModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAccountResetModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAccountResetModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


