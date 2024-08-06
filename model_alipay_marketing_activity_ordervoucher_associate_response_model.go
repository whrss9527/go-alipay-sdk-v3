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

// checks if the AlipayMarketingActivityOrdervoucherAssociateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityOrdervoucherAssociateResponseModel{}

// AlipayMarketingActivityOrdervoucherAssociateResponseModel struct for AlipayMarketingActivityOrdervoucherAssociateResponseModel
type AlipayMarketingActivityOrdervoucherAssociateResponseModel struct {
	// 支付宝系统关联订单成功的时间。 格式为：yyyy-MM-dd HH:mm:ss
	AssociateTime *string `json:"associate_time,omitempty"`
}

// NewAlipayMarketingActivityOrdervoucherAssociateResponseModel instantiates a new AlipayMarketingActivityOrdervoucherAssociateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityOrdervoucherAssociateResponseModel() *AlipayMarketingActivityOrdervoucherAssociateResponseModel {
	this := AlipayMarketingActivityOrdervoucherAssociateResponseModel{}
	return &this
}

// NewAlipayMarketingActivityOrdervoucherAssociateResponseModelWithDefaults instantiates a new AlipayMarketingActivityOrdervoucherAssociateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityOrdervoucherAssociateResponseModelWithDefaults() *AlipayMarketingActivityOrdervoucherAssociateResponseModel {
	this := AlipayMarketingActivityOrdervoucherAssociateResponseModel{}
	return &this
}

// GetAssociateTime returns the AssociateTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherAssociateResponseModel) GetAssociateTime() string {
	if o == nil || IsNil(o.AssociateTime) {
		var ret string
		return ret
	}
	return *o.AssociateTime
}

// GetAssociateTimeOk returns a tuple with the AssociateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherAssociateResponseModel) GetAssociateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AssociateTime) {
		return nil, false
	}
	return o.AssociateTime, true
}

// HasAssociateTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherAssociateResponseModel) HasAssociateTime() bool {
	if o != nil && !IsNil(o.AssociateTime) {
		return true
	}

	return false
}

// SetAssociateTime gets a reference to the given string and assigns it to the AssociateTime field.
func (o *AlipayMarketingActivityOrdervoucherAssociateResponseModel) SetAssociateTime(v string) {
	o.AssociateTime = &v
}

func (o AlipayMarketingActivityOrdervoucherAssociateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityOrdervoucherAssociateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociateTime) {
		toSerialize["associate_time"] = o.AssociateTime
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel struct {
	value *AlipayMarketingActivityOrdervoucherAssociateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) Get() *AlipayMarketingActivityOrdervoucherAssociateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) Set(val *AlipayMarketingActivityOrdervoucherAssociateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherAssociateResponseModel(val *AlipayMarketingActivityOrdervoucherAssociateResponseModel) *NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel {
	return &NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherAssociateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
