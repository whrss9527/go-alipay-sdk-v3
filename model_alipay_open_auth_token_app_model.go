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

// checks if the AlipayOpenAuthTokenAppModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAuthTokenAppModel{}

// AlipayOpenAuthTokenAppModel struct for AlipayOpenAuthTokenAppModel
type AlipayOpenAuthTokenAppModel struct {
	// 应用授权码，传入应用授权后得到的  app_auth_code。 说明： grant_type 为 authorization_code 时，本参数必填； grant_type 为 refresh_token 时不填。
	Code *string `json:"code,omitempty"`
	// 授权方式。支持： authorization_code：使用应用授权码换取应用授权令牌app_auth_token。 refresh_token：使用app_refresh_token刷新获取新授权令牌。
	GrantType *string `json:"grant_type,omitempty"`
	// 刷新令牌，上次换取访问令牌时得到。本参数在 grant_type 为 authorization_code 时不填；为 refresh_token 时必填，且该值来源于此接口的返回值 app_refresh_token（即至少需要通过 grant_type=authorization_code 调用此接口一次才能获取）。
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewAlipayOpenAuthTokenAppModel instantiates a new AlipayOpenAuthTokenAppModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAuthTokenAppModel() *AlipayOpenAuthTokenAppModel {
	this := AlipayOpenAuthTokenAppModel{}
	return &this
}

// NewAlipayOpenAuthTokenAppModelWithDefaults instantiates a new AlipayOpenAuthTokenAppModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAuthTokenAppModelWithDefaults() *AlipayOpenAuthTokenAppModel {
	this := AlipayOpenAuthTokenAppModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayOpenAuthTokenAppModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAuthTokenAppModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayOpenAuthTokenAppModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayOpenAuthTokenAppModel) SetCode(v string) {
	o.Code = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *AlipayOpenAuthTokenAppModel) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAuthTokenAppModel) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *AlipayOpenAuthTokenAppModel) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *AlipayOpenAuthTokenAppModel) SetGrantType(v string) {
	o.GrantType = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AlipayOpenAuthTokenAppModel) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAuthTokenAppModel) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AlipayOpenAuthTokenAppModel) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AlipayOpenAuthTokenAppModel) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o AlipayOpenAuthTokenAppModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAuthTokenAppModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.GrantType) {
		toSerialize["grant_type"] = o.GrantType
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableAlipayOpenAuthTokenAppModel struct {
	value *AlipayOpenAuthTokenAppModel
	isSet bool
}

func (v NullableAlipayOpenAuthTokenAppModel) Get() *AlipayOpenAuthTokenAppModel {
	return v.value
}

func (v *NullableAlipayOpenAuthTokenAppModel) Set(val *AlipayOpenAuthTokenAppModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAuthTokenAppModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAuthTokenAppModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAuthTokenAppModel(val *AlipayOpenAuthTokenAppModel) *NullableAlipayOpenAuthTokenAppModel {
	return &NullableAlipayOpenAuthTokenAppModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAuthTokenAppModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAuthTokenAppModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
