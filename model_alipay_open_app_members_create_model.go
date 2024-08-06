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

// checks if the AlipayOpenAppMembersCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAppMembersCreateModel{}

// AlipayOpenAppMembersCreateModel struct for AlipayOpenAppMembersCreateModel
type AlipayOpenAppMembersCreateModel struct {
	// 支付宝登录账号。
	LogonId *string `json:"logon_id,omitempty"`
	// 为成员添加的角色类型
	Role *string `json:"role,omitempty"`
}

// NewAlipayOpenAppMembersCreateModel instantiates a new AlipayOpenAppMembersCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAppMembersCreateModel() *AlipayOpenAppMembersCreateModel {
	this := AlipayOpenAppMembersCreateModel{}
	return &this
}

// NewAlipayOpenAppMembersCreateModelWithDefaults instantiates a new AlipayOpenAppMembersCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAppMembersCreateModelWithDefaults() *AlipayOpenAppMembersCreateModel {
	this := AlipayOpenAppMembersCreateModel{}
	return &this
}

// GetLogonId returns the LogonId field value if set, zero value otherwise.
func (o *AlipayOpenAppMembersCreateModel) GetLogonId() string {
	if o == nil || IsNil(o.LogonId) {
		var ret string
		return ret
	}
	return *o.LogonId
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppMembersCreateModel) GetLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogonId) {
		return nil, false
	}
	return o.LogonId, true
}

// HasLogonId returns a boolean if a field has been set.
func (o *AlipayOpenAppMembersCreateModel) HasLogonId() bool {
	if o != nil && !IsNil(o.LogonId) {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given string and assigns it to the LogonId field.
func (o *AlipayOpenAppMembersCreateModel) SetLogonId(v string) {
	o.LogonId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AlipayOpenAppMembersCreateModel) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAppMembersCreateModel) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AlipayOpenAppMembersCreateModel) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AlipayOpenAppMembersCreateModel) SetRole(v string) {
	o.Role = &v
}

func (o AlipayOpenAppMembersCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAppMembersCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogonId) {
		toSerialize["logon_id"] = o.LogonId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableAlipayOpenAppMembersCreateModel struct {
	value *AlipayOpenAppMembersCreateModel
	isSet bool
}

func (v NullableAlipayOpenAppMembersCreateModel) Get() *AlipayOpenAppMembersCreateModel {
	return v.value
}

func (v *NullableAlipayOpenAppMembersCreateModel) Set(val *AlipayOpenAppMembersCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppMembersCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppMembersCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppMembersCreateModel(val *AlipayOpenAppMembersCreateModel) *NullableAlipayOpenAppMembersCreateModel {
	return &NullableAlipayOpenAppMembersCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAppMembersCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppMembersCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
