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

// checks if the InvoiceTitleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceTitleModel{}

// InvoiceTitleModel struct for InvoiceTitleModel
type InvoiceTitleModel struct {
	// 是否为用户设置默认抬头  字段值包括两种情况：  false（非默认）  true（默认抬头）
	IsDefault *bool `json:"is_default,omitempty"`
	// 支付宝用户登录名
	LogonId *string `json:"logon_id,omitempty"`
	// 银行账号
	OpenBankAccount *string `json:"open_bank_account,omitempty"`
	// 开户银行
	OpenBankName *string `json:"open_bank_name,omitempty"`
	// 支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// 纳税人识别号
	TaxRegisterNo *string `json:"tax_register_no,omitempty"`
	// 用户私人手机号
	TelePhoneNo *string `json:"tele_phone_no,omitempty"`
	// 抬头名称
	TitleName *string `json:"title_name,omitempty"`
	// 抬头类型 字段值有两种情况抬: PERSONAL（个人）  CORPORATION（企业）
	TitleType *string `json:"title_type,omitempty"`
	// 地址
	UserAddress *string `json:"user_address,omitempty"`
	// 邮箱
	UserEmail *string `json:"user_email,omitempty"`
	// 支付宝用户id
	UserId *string `json:"user_id,omitempty"`
	// 电话号码
	UserMobile *string `json:"user_mobile,omitempty"`
}

// NewInvoiceTitleModel instantiates a new InvoiceTitleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceTitleModel() *InvoiceTitleModel {
	this := InvoiceTitleModel{}
	return &this
}

// NewInvoiceTitleModelWithDefaults instantiates a new InvoiceTitleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceTitleModelWithDefaults() *InvoiceTitleModel {
	this := InvoiceTitleModel{}
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *InvoiceTitleModel) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLogonId returns the LogonId field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetLogonId() string {
	if o == nil || IsNil(o.LogonId) {
		var ret string
		return ret
	}
	return *o.LogonId
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogonId) {
		return nil, false
	}
	return o.LogonId, true
}

// HasLogonId returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasLogonId() bool {
	if o != nil && !IsNil(o.LogonId) {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given string and assigns it to the LogonId field.
func (o *InvoiceTitleModel) SetLogonId(v string) {
	o.LogonId = &v
}

// GetOpenBankAccount returns the OpenBankAccount field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetOpenBankAccount() string {
	if o == nil || IsNil(o.OpenBankAccount) {
		var ret string
		return ret
	}
	return *o.OpenBankAccount
}

// GetOpenBankAccountOk returns a tuple with the OpenBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetOpenBankAccountOk() (*string, bool) {
	if o == nil || IsNil(o.OpenBankAccount) {
		return nil, false
	}
	return o.OpenBankAccount, true
}

// HasOpenBankAccount returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasOpenBankAccount() bool {
	if o != nil && !IsNil(o.OpenBankAccount) {
		return true
	}

	return false
}

// SetOpenBankAccount gets a reference to the given string and assigns it to the OpenBankAccount field.
func (o *InvoiceTitleModel) SetOpenBankAccount(v string) {
	o.OpenBankAccount = &v
}

// GetOpenBankName returns the OpenBankName field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetOpenBankName() string {
	if o == nil || IsNil(o.OpenBankName) {
		var ret string
		return ret
	}
	return *o.OpenBankName
}

// GetOpenBankNameOk returns a tuple with the OpenBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetOpenBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.OpenBankName) {
		return nil, false
	}
	return o.OpenBankName, true
}

// HasOpenBankName returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasOpenBankName() bool {
	if o != nil && !IsNil(o.OpenBankName) {
		return true
	}

	return false
}

// SetOpenBankName gets a reference to the given string and assigns it to the OpenBankName field.
func (o *InvoiceTitleModel) SetOpenBankName(v string) {
	o.OpenBankName = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *InvoiceTitleModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetTaxRegisterNo returns the TaxRegisterNo field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetTaxRegisterNo() string {
	if o == nil || IsNil(o.TaxRegisterNo) {
		var ret string
		return ret
	}
	return *o.TaxRegisterNo
}

// GetTaxRegisterNoOk returns a tuple with the TaxRegisterNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetTaxRegisterNoOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegisterNo) {
		return nil, false
	}
	return o.TaxRegisterNo, true
}

// HasTaxRegisterNo returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasTaxRegisterNo() bool {
	if o != nil && !IsNil(o.TaxRegisterNo) {
		return true
	}

	return false
}

// SetTaxRegisterNo gets a reference to the given string and assigns it to the TaxRegisterNo field.
func (o *InvoiceTitleModel) SetTaxRegisterNo(v string) {
	o.TaxRegisterNo = &v
}

// GetTelePhoneNo returns the TelePhoneNo field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetTelePhoneNo() string {
	if o == nil || IsNil(o.TelePhoneNo) {
		var ret string
		return ret
	}
	return *o.TelePhoneNo
}

// GetTelePhoneNoOk returns a tuple with the TelePhoneNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetTelePhoneNoOk() (*string, bool) {
	if o == nil || IsNil(o.TelePhoneNo) {
		return nil, false
	}
	return o.TelePhoneNo, true
}

// HasTelePhoneNo returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasTelePhoneNo() bool {
	if o != nil && !IsNil(o.TelePhoneNo) {
		return true
	}

	return false
}

// SetTelePhoneNo gets a reference to the given string and assigns it to the TelePhoneNo field.
func (o *InvoiceTitleModel) SetTelePhoneNo(v string) {
	o.TelePhoneNo = &v
}

// GetTitleName returns the TitleName field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetTitleName() string {
	if o == nil || IsNil(o.TitleName) {
		var ret string
		return ret
	}
	return *o.TitleName
}

// GetTitleNameOk returns a tuple with the TitleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetTitleNameOk() (*string, bool) {
	if o == nil || IsNil(o.TitleName) {
		return nil, false
	}
	return o.TitleName, true
}

// HasTitleName returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasTitleName() bool {
	if o != nil && !IsNil(o.TitleName) {
		return true
	}

	return false
}

// SetTitleName gets a reference to the given string and assigns it to the TitleName field.
func (o *InvoiceTitleModel) SetTitleName(v string) {
	o.TitleName = &v
}

// GetTitleType returns the TitleType field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetTitleType() string {
	if o == nil || IsNil(o.TitleType) {
		var ret string
		return ret
	}
	return *o.TitleType
}

// GetTitleTypeOk returns a tuple with the TitleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetTitleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TitleType) {
		return nil, false
	}
	return o.TitleType, true
}

// HasTitleType returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasTitleType() bool {
	if o != nil && !IsNil(o.TitleType) {
		return true
	}

	return false
}

// SetTitleType gets a reference to the given string and assigns it to the TitleType field.
func (o *InvoiceTitleModel) SetTitleType(v string) {
	o.TitleType = &v
}

// GetUserAddress returns the UserAddress field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetUserAddress() string {
	if o == nil || IsNil(o.UserAddress) {
		var ret string
		return ret
	}
	return *o.UserAddress
}

// GetUserAddressOk returns a tuple with the UserAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetUserAddressOk() (*string, bool) {
	if o == nil || IsNil(o.UserAddress) {
		return nil, false
	}
	return o.UserAddress, true
}

// HasUserAddress returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasUserAddress() bool {
	if o != nil && !IsNil(o.UserAddress) {
		return true
	}

	return false
}

// SetUserAddress gets a reference to the given string and assigns it to the UserAddress field.
func (o *InvoiceTitleModel) SetUserAddress(v string) {
	o.UserAddress = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetUserEmail() string {
	if o == nil || IsNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasUserEmail() bool {
	if o != nil && !IsNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *InvoiceTitleModel) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *InvoiceTitleModel) SetUserId(v string) {
	o.UserId = &v
}

// GetUserMobile returns the UserMobile field value if set, zero value otherwise.
func (o *InvoiceTitleModel) GetUserMobile() string {
	if o == nil || IsNil(o.UserMobile) {
		var ret string
		return ret
	}
	return *o.UserMobile
}

// GetUserMobileOk returns a tuple with the UserMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceTitleModel) GetUserMobileOk() (*string, bool) {
	if o == nil || IsNil(o.UserMobile) {
		return nil, false
	}
	return o.UserMobile, true
}

// HasUserMobile returns a boolean if a field has been set.
func (o *InvoiceTitleModel) HasUserMobile() bool {
	if o != nil && !IsNil(o.UserMobile) {
		return true
	}

	return false
}

// SetUserMobile gets a reference to the given string and assigns it to the UserMobile field.
func (o *InvoiceTitleModel) SetUserMobile(v string) {
	o.UserMobile = &v
}

func (o InvoiceTitleModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceTitleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.LogonId) {
		toSerialize["logon_id"] = o.LogonId
	}
	if !IsNil(o.OpenBankAccount) {
		toSerialize["open_bank_account"] = o.OpenBankAccount
	}
	if !IsNil(o.OpenBankName) {
		toSerialize["open_bank_name"] = o.OpenBankName
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.TaxRegisterNo) {
		toSerialize["tax_register_no"] = o.TaxRegisterNo
	}
	if !IsNil(o.TelePhoneNo) {
		toSerialize["tele_phone_no"] = o.TelePhoneNo
	}
	if !IsNil(o.TitleName) {
		toSerialize["title_name"] = o.TitleName
	}
	if !IsNil(o.TitleType) {
		toSerialize["title_type"] = o.TitleType
	}
	if !IsNil(o.UserAddress) {
		toSerialize["user_address"] = o.UserAddress
	}
	if !IsNil(o.UserEmail) {
		toSerialize["user_email"] = o.UserEmail
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserMobile) {
		toSerialize["user_mobile"] = o.UserMobile
	}
	return toSerialize, nil
}

type NullableInvoiceTitleModel struct {
	value *InvoiceTitleModel
	isSet bool
}

func (v NullableInvoiceTitleModel) Get() *InvoiceTitleModel {
	return v.value
}

func (v *NullableInvoiceTitleModel) Set(val *InvoiceTitleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceTitleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceTitleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceTitleModel(val *InvoiceTitleModel) *NullableInvoiceTitleModel {
	return &NullableInvoiceTitleModel{value: val, isSet: true}
}

func (v NullableInvoiceTitleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceTitleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

