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

// checks if the AlipayMarketingActivityVoucherpackageConsultModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherpackageConsultModel{}

// AlipayMarketingActivityVoucherpackageConsultModel struct for AlipayMarketingActivityVoucherpackageConsultModel
type AlipayMarketingActivityVoucherpackageConsultModel struct {
	// 券包id列表
	VoucherPackageIdList []string `json:"voucher_package_id_list,omitempty"`
}

// NewAlipayMarketingActivityVoucherpackageConsultModel instantiates a new AlipayMarketingActivityVoucherpackageConsultModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherpackageConsultModel() *AlipayMarketingActivityVoucherpackageConsultModel {
	this := AlipayMarketingActivityVoucherpackageConsultModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherpackageConsultModelWithDefaults instantiates a new AlipayMarketingActivityVoucherpackageConsultModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherpackageConsultModelWithDefaults() *AlipayMarketingActivityVoucherpackageConsultModel {
	this := AlipayMarketingActivityVoucherpackageConsultModel{}
	return &this
}

// GetVoucherPackageIdList returns the VoucherPackageIdList field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherpackageConsultModel) GetVoucherPackageIdList() []string {
	if o == nil || IsNil(o.VoucherPackageIdList) {
		var ret []string
		return ret
	}
	return o.VoucherPackageIdList
}

// GetVoucherPackageIdListOk returns a tuple with the VoucherPackageIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherpackageConsultModel) GetVoucherPackageIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.VoucherPackageIdList) {
		return nil, false
	}
	return o.VoucherPackageIdList, true
}

// HasVoucherPackageIdList returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherpackageConsultModel) HasVoucherPackageIdList() bool {
	if o != nil && !IsNil(o.VoucherPackageIdList) {
		return true
	}

	return false
}

// SetVoucherPackageIdList gets a reference to the given []string and assigns it to the VoucherPackageIdList field.
func (o *AlipayMarketingActivityVoucherpackageConsultModel) SetVoucherPackageIdList(v []string) {
	o.VoucherPackageIdList = v
}

func (o AlipayMarketingActivityVoucherpackageConsultModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherpackageConsultModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoucherPackageIdList) {
		toSerialize["voucher_package_id_list"] = o.VoucherPackageIdList
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherpackageConsultModel struct {
	value *AlipayMarketingActivityVoucherpackageConsultModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherpackageConsultModel) Get() *AlipayMarketingActivityVoucherpackageConsultModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherpackageConsultModel) Set(val *AlipayMarketingActivityVoucherpackageConsultModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherpackageConsultModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherpackageConsultModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherpackageConsultModel(val *AlipayMarketingActivityVoucherpackageConsultModel) *NullableAlipayMarketingActivityVoucherpackageConsultModel {
	return &NullableAlipayMarketingActivityVoucherpackageConsultModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherpackageConsultModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherpackageConsultModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


