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

// checks if the AlipayOpenMiniInnerMembersAddModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerMembersAddModel{}

// AlipayOpenMiniInnerMembersAddModel struct for AlipayOpenMiniInnerMembersAddModel
type AlipayOpenMiniInnerMembersAddModel struct {
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 域账号（不为空时会添加联系人）
	DomainAccount *string `json:"domain_account,omitempty"`
	// 支付宝登陆账号
	LoginId *string `json:"login_id,omitempty"`
	// 业务appId
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 成员类型
	Type *string `json:"type,omitempty"`
}

// NewAlipayOpenMiniInnerMembersAddModel instantiates a new AlipayOpenMiniInnerMembersAddModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerMembersAddModel() *AlipayOpenMiniInnerMembersAddModel {
	this := AlipayOpenMiniInnerMembersAddModel{}
	return &this
}

// NewAlipayOpenMiniInnerMembersAddModelWithDefaults instantiates a new AlipayOpenMiniInnerMembersAddModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerMembersAddModelWithDefaults() *AlipayOpenMiniInnerMembersAddModel {
	this := AlipayOpenMiniInnerMembersAddModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersAddModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerMembersAddModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetDomainAccount returns the DomainAccount field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersAddModel) GetDomainAccount() string {
	if o == nil || IsNil(o.DomainAccount) {
		var ret string
		return ret
	}
	return *o.DomainAccount
}

// GetDomainAccountOk returns a tuple with the DomainAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) GetDomainAccountOk() (*string, bool) {
	if o == nil || IsNil(o.DomainAccount) {
		return nil, false
	}
	return o.DomainAccount, true
}

// HasDomainAccount returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) HasDomainAccount() bool {
	if o != nil && !IsNil(o.DomainAccount) {
		return true
	}

	return false
}

// SetDomainAccount gets a reference to the given string and assigns it to the DomainAccount field.
func (o *AlipayOpenMiniInnerMembersAddModel) SetDomainAccount(v string) {
	o.DomainAccount = &v
}

// GetLoginId returns the LoginId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersAddModel) GetLoginId() string {
	if o == nil || IsNil(o.LoginId) {
		var ret string
		return ret
	}
	return *o.LoginId
}

// GetLoginIdOk returns a tuple with the LoginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) GetLoginIdOk() (*string, bool) {
	if o == nil || IsNil(o.LoginId) {
		return nil, false
	}
	return o.LoginId, true
}

// HasLoginId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) HasLoginId() bool {
	if o != nil && !IsNil(o.LoginId) {
		return true
	}

	return false
}

// SetLoginId gets a reference to the given string and assigns it to the LoginId field.
func (o *AlipayOpenMiniInnerMembersAddModel) SetLoginId(v string) {
	o.LoginId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersAddModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerMembersAddModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersAddModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersAddModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayOpenMiniInnerMembersAddModel) SetType(v string) {
	o.Type = &v
}

func (o AlipayOpenMiniInnerMembersAddModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerMembersAddModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.DomainAccount) {
		toSerialize["domain_account"] = o.DomainAccount
	}
	if !IsNil(o.LoginId) {
		toSerialize["login_id"] = o.LoginId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerMembersAddModel struct {
	value *AlipayOpenMiniInnerMembersAddModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerMembersAddModel) Get() *AlipayOpenMiniInnerMembersAddModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerMembersAddModel) Set(val *AlipayOpenMiniInnerMembersAddModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerMembersAddModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerMembersAddModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerMembersAddModel(val *AlipayOpenMiniInnerMembersAddModel) *NullableAlipayOpenMiniInnerMembersAddModel {
	return &NullableAlipayOpenMiniInnerMembersAddModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerMembersAddModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerMembersAddModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
