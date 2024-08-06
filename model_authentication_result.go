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

// checks if the AuthenticationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationResult{}

// AuthenticationResult struct for AuthenticationResult
type AuthenticationResult struct {
	// 已加密的鉴权数据
	AuthenticationData *string `json:"authentication_data,omitempty"`
	// 鉴权类型，如：支付密码、数字签名
	AuthenticationMechanism *string `json:"authentication_mechanism,omitempty"`
}

// NewAuthenticationResult instantiates a new AuthenticationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationResult() *AuthenticationResult {
	this := AuthenticationResult{}
	return &this
}

// NewAuthenticationResultWithDefaults instantiates a new AuthenticationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationResultWithDefaults() *AuthenticationResult {
	this := AuthenticationResult{}
	return &this
}

// GetAuthenticationData returns the AuthenticationData field value if set, zero value otherwise.
func (o *AuthenticationResult) GetAuthenticationData() string {
	if o == nil || IsNil(o.AuthenticationData) {
		var ret string
		return ret
	}
	return *o.AuthenticationData
}

// GetAuthenticationDataOk returns a tuple with the AuthenticationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationResult) GetAuthenticationDataOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationData) {
		return nil, false
	}
	return o.AuthenticationData, true
}

// HasAuthenticationData returns a boolean if a field has been set.
func (o *AuthenticationResult) HasAuthenticationData() bool {
	if o != nil && !IsNil(o.AuthenticationData) {
		return true
	}

	return false
}

// SetAuthenticationData gets a reference to the given string and assigns it to the AuthenticationData field.
func (o *AuthenticationResult) SetAuthenticationData(v string) {
	o.AuthenticationData = &v
}

// GetAuthenticationMechanism returns the AuthenticationMechanism field value if set, zero value otherwise.
func (o *AuthenticationResult) GetAuthenticationMechanism() string {
	if o == nil || IsNil(o.AuthenticationMechanism) {
		var ret string
		return ret
	}
	return *o.AuthenticationMechanism
}

// GetAuthenticationMechanismOk returns a tuple with the AuthenticationMechanism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationResult) GetAuthenticationMechanismOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMechanism) {
		return nil, false
	}
	return o.AuthenticationMechanism, true
}

// HasAuthenticationMechanism returns a boolean if a field has been set.
func (o *AuthenticationResult) HasAuthenticationMechanism() bool {
	if o != nil && !IsNil(o.AuthenticationMechanism) {
		return true
	}

	return false
}

// SetAuthenticationMechanism gets a reference to the given string and assigns it to the AuthenticationMechanism field.
func (o *AuthenticationResult) SetAuthenticationMechanism(v string) {
	o.AuthenticationMechanism = &v
}

func (o AuthenticationResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationData) {
		toSerialize["authentication_data"] = o.AuthenticationData
	}
	if !IsNil(o.AuthenticationMechanism) {
		toSerialize["authentication_mechanism"] = o.AuthenticationMechanism
	}
	return toSerialize, nil
}

type NullableAuthenticationResult struct {
	value *AuthenticationResult
	isSet bool
}

func (v NullableAuthenticationResult) Get() *AuthenticationResult {
	return v.value
}

func (v *NullableAuthenticationResult) Set(val *AuthenticationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationResult(val *AuthenticationResult) *NullableAuthenticationResult {
	return &NullableAuthenticationResult{value: val, isSet: true}
}

func (v NullableAuthenticationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
