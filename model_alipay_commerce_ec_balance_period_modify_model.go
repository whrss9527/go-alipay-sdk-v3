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

// checks if the AlipayCommerceEcBalancePeriodModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcBalancePeriodModifyModel{}

// AlipayCommerceEcBalancePeriodModifyModel struct for AlipayCommerceEcBalancePeriodModifyModel
type AlipayCommerceEcBalancePeriodModifyModel struct {
	// 共同账户id - 适用于在企业码小程序创建的共同账户，和agreement_no搭配使用
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 月账单账期
	BillDay *int32 `json:"bill_day,omitempty"`
	// 企业ID - 适用于在企业码PC端创建的企业账号
	EnterpriseId *string `json:"enterprise_id,omitempty"`
}

// NewAlipayCommerceEcBalancePeriodModifyModel instantiates a new AlipayCommerceEcBalancePeriodModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcBalancePeriodModifyModel() *AlipayCommerceEcBalancePeriodModifyModel {
	this := AlipayCommerceEcBalancePeriodModifyModel{}
	return &this
}

// NewAlipayCommerceEcBalancePeriodModifyModelWithDefaults instantiates a new AlipayCommerceEcBalancePeriodModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcBalancePeriodModifyModelWithDefaults() *AlipayCommerceEcBalancePeriodModifyModel {
	this := AlipayCommerceEcBalancePeriodModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayCommerceEcBalancePeriodModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayCommerceEcBalancePeriodModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetBillDay returns the BillDay field value if set, zero value otherwise.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetBillDay() int32 {
	if o == nil || IsNil(o.BillDay) {
		var ret int32
		return ret
	}
	return *o.BillDay
}

// GetBillDayOk returns a tuple with the BillDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetBillDayOk() (*int32, bool) {
	if o == nil || IsNil(o.BillDay) {
		return nil, false
	}
	return o.BillDay, true
}

// HasBillDay returns a boolean if a field has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) HasBillDay() bool {
	if o != nil && !IsNil(o.BillDay) {
		return true
	}

	return false
}

// SetBillDay gets a reference to the given int32 and assigns it to the BillDay field.
func (o *AlipayCommerceEcBalancePeriodModifyModel) SetBillDay(v int32) {
	o.BillDay = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayCommerceEcBalancePeriodModifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayCommerceEcBalancePeriodModifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

func (o AlipayCommerceEcBalancePeriodModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcBalancePeriodModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.BillDay) {
		toSerialize["bill_day"] = o.BillDay
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcBalancePeriodModifyModel struct {
	value *AlipayCommerceEcBalancePeriodModifyModel
	isSet bool
}

func (v NullableAlipayCommerceEcBalancePeriodModifyModel) Get() *AlipayCommerceEcBalancePeriodModifyModel {
	return v.value
}

func (v *NullableAlipayCommerceEcBalancePeriodModifyModel) Set(val *AlipayCommerceEcBalancePeriodModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcBalancePeriodModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcBalancePeriodModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcBalancePeriodModifyModel(val *AlipayCommerceEcBalancePeriodModifyModel) *NullableAlipayCommerceEcBalancePeriodModifyModel {
	return &NullableAlipayCommerceEcBalancePeriodModifyModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcBalancePeriodModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcBalancePeriodModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


