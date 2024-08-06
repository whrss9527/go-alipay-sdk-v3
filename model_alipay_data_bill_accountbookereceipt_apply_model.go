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

// checks if the AlipayDataBillAccountbookereceiptApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataBillAccountbookereceiptApplyModel{}

// AlipayDataBillAccountbookereceiptApplyModel struct for AlipayDataBillAccountbookereceiptApplyModel
type AlipayDataBillAccountbookereceiptApplyModel struct {
	// 协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 明细凭证，传入流水号（转账接口查询结果）。汇总凭证，传入起止时间，格式yyyy-MM-dd_yyyy-MM-dd。
	Key *string `json:"key,omitempty"`
	// 子账本号，或者子账本名称。模糊查询
	StoreNo *string `json:"store_no,omitempty"`
	// 申请的类型。可传入：FUND_PLATFORM_DETAIL - 资金明细证明，FUND_PLATFORM_SUM - 资金汇总证明
	Type *string `json:"type,omitempty"`
}

// NewAlipayDataBillAccountbookereceiptApplyModel instantiates a new AlipayDataBillAccountbookereceiptApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataBillAccountbookereceiptApplyModel() *AlipayDataBillAccountbookereceiptApplyModel {
	this := AlipayDataBillAccountbookereceiptApplyModel{}
	return &this
}

// NewAlipayDataBillAccountbookereceiptApplyModelWithDefaults instantiates a new AlipayDataBillAccountbookereceiptApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataBillAccountbookereceiptApplyModelWithDefaults() *AlipayDataBillAccountbookereceiptApplyModel {
	this := AlipayDataBillAccountbookereceiptApplyModel{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayDataBillAccountbookereceiptApplyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AlipayDataBillAccountbookereceiptApplyModel) SetKey(v string) {
	o.Key = &v
}

// GetStoreNo returns the StoreNo field value if set, zero value otherwise.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetStoreNo() string {
	if o == nil || IsNil(o.StoreNo) {
		var ret string
		return ret
	}
	return *o.StoreNo
}

// GetStoreNoOk returns a tuple with the StoreNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetStoreNoOk() (*string, bool) {
	if o == nil || IsNil(o.StoreNo) {
		return nil, false
	}
	return o.StoreNo, true
}

// HasStoreNo returns a boolean if a field has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) HasStoreNo() bool {
	if o != nil && !IsNil(o.StoreNo) {
		return true
	}

	return false
}

// SetStoreNo gets a reference to the given string and assigns it to the StoreNo field.
func (o *AlipayDataBillAccountbookereceiptApplyModel) SetStoreNo(v string) {
	o.StoreNo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayDataBillAccountbookereceiptApplyModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayDataBillAccountbookereceiptApplyModel) SetType(v string) {
	o.Type = &v
}

func (o AlipayDataBillAccountbookereceiptApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataBillAccountbookereceiptApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.StoreNo) {
		toSerialize["store_no"] = o.StoreNo
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAlipayDataBillAccountbookereceiptApplyModel struct {
	value *AlipayDataBillAccountbookereceiptApplyModel
	isSet bool
}

func (v NullableAlipayDataBillAccountbookereceiptApplyModel) Get() *AlipayDataBillAccountbookereceiptApplyModel {
	return v.value
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyModel) Set(val *AlipayDataBillAccountbookereceiptApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillAccountbookereceiptApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillAccountbookereceiptApplyModel(val *AlipayDataBillAccountbookereceiptApplyModel) *NullableAlipayDataBillAccountbookereceiptApplyModel {
	return &NullableAlipayDataBillAccountbookereceiptApplyModel{value: val, isSet: true}
}

func (v NullableAlipayDataBillAccountbookereceiptApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillAccountbookereceiptApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

