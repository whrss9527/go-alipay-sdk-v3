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

// checks if the AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel{}

// AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel struct for AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel
type AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel struct {
	// 即时配送运单列表
	Waybills []PreCreateWaybillIstd `json:"waybills,omitempty"`
}

// NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel() *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel{}
	return &this
}

// NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModelWithDefaults() *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel {
	this := AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel{}
	return &this
}

// GetWaybills returns the Waybills field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) GetWaybills() []PreCreateWaybillIstd {
	if o == nil || IsNil(o.Waybills) {
		var ret []PreCreateWaybillIstd
		return ret
	}
	return o.Waybills
}

// GetWaybillsOk returns a tuple with the Waybills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) GetWaybillsOk() ([]PreCreateWaybillIstd, bool) {
	if o == nil || IsNil(o.Waybills) {
		return nil, false
	}
	return o.Waybills, true
}

// HasWaybills returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) HasWaybills() bool {
	if o != nil && !IsNil(o.Waybills) {
		return true
	}

	return false
}

// SetWaybills gets a reference to the given []PreCreateWaybillIstd and assigns it to the Waybills field.
func (o *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) SetWaybills(v []PreCreateWaybillIstd) {
	o.Waybills = v
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Waybills) {
		toSerialize["waybills"] = o.Waybills
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel struct {
	value *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) Get() *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) Set(val *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel(val *AlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel {
	return &NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsOrderInstantdeliveryPrecreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


