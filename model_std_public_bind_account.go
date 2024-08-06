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

// checks if the StdPublicBindAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StdPublicBindAccount{}

// StdPublicBindAccount struct for StdPublicBindAccount
type StdPublicBindAccount struct {
	// 协议号是商户会员在支付宝公众账号中的唯一标识。
	AgreementId *string `json:"agreement_id,omitempty"`
	// 公众账号ID
	AppId *string `json:"app_id,omitempty"`
	// 绑定的商户会员号
	BindAccountNo *string `json:"bind_account_no,omitempty"`
	// 公众账号期望支付宝用户在公众账号首页看到的关于该用户的显示信息，最长10个汉字。
	DisplayName *string `json:"display_name,omitempty"`
	// 绑定的商户会员对应的支付宝用户号，以2088 开头的16位数字。
	FromUserId *string `json:"from_user_id,omitempty"`
	// 1
	OpenId *string `json:"open_id,omitempty"`
	// 绑定的商户会员的真实姓名，最长10个汉字。
	RealName *string `json:"real_name,omitempty"`
}

// NewStdPublicBindAccount instantiates a new StdPublicBindAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStdPublicBindAccount() *StdPublicBindAccount {
	this := StdPublicBindAccount{}
	return &this
}

// NewStdPublicBindAccountWithDefaults instantiates a new StdPublicBindAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStdPublicBindAccountWithDefaults() *StdPublicBindAccount {
	this := StdPublicBindAccount{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *StdPublicBindAccount) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *StdPublicBindAccount) SetAppId(v string) {
	o.AppId = &v
}

// GetBindAccountNo returns the BindAccountNo field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetBindAccountNo() string {
	if o == nil || IsNil(o.BindAccountNo) {
		var ret string
		return ret
	}
	return *o.BindAccountNo
}

// GetBindAccountNoOk returns a tuple with the BindAccountNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetBindAccountNoOk() (*string, bool) {
	if o == nil || IsNil(o.BindAccountNo) {
		return nil, false
	}
	return o.BindAccountNo, true
}

// HasBindAccountNo returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasBindAccountNo() bool {
	if o != nil && !IsNil(o.BindAccountNo) {
		return true
	}

	return false
}

// SetBindAccountNo gets a reference to the given string and assigns it to the BindAccountNo field.
func (o *StdPublicBindAccount) SetBindAccountNo(v string) {
	o.BindAccountNo = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StdPublicBindAccount) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFromUserId returns the FromUserId field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetFromUserId() string {
	if o == nil || IsNil(o.FromUserId) {
		var ret string
		return ret
	}
	return *o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetFromUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.FromUserId) {
		return nil, false
	}
	return o.FromUserId, true
}

// HasFromUserId returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasFromUserId() bool {
	if o != nil && !IsNil(o.FromUserId) {
		return true
	}

	return false
}

// SetFromUserId gets a reference to the given string and assigns it to the FromUserId field.
func (o *StdPublicBindAccount) SetFromUserId(v string) {
	o.FromUserId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *StdPublicBindAccount) SetOpenId(v string) {
	o.OpenId = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *StdPublicBindAccount) GetRealName() string {
	if o == nil || IsNil(o.RealName) {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StdPublicBindAccount) GetRealNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealName) {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *StdPublicBindAccount) HasRealName() bool {
	if o != nil && !IsNil(o.RealName) {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *StdPublicBindAccount) SetRealName(v string) {
	o.RealName = &v
}

func (o StdPublicBindAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StdPublicBindAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AppId) {
		toSerialize["app_id"] = o.AppId
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
	return toSerialize, nil
}

type NullableStdPublicBindAccount struct {
	value *StdPublicBindAccount
	isSet bool
}

func (v NullableStdPublicBindAccount) Get() *StdPublicBindAccount {
	return v.value
}

func (v *NullableStdPublicBindAccount) Set(val *StdPublicBindAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableStdPublicBindAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableStdPublicBindAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStdPublicBindAccount(val *StdPublicBindAccount) *NullableStdPublicBindAccount {
	return &NullableStdPublicBindAccount{value: val, isSet: true}
}

func (v NullableStdPublicBindAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStdPublicBindAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


