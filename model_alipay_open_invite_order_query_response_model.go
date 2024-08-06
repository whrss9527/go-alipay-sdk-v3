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

// checks if the AlipayOpenInviteOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenInviteOrderQueryResponseModel{}

// AlipayOpenInviteOrderQueryResponseModel struct for AlipayOpenInviteOrderQueryResponseModel
type AlipayOpenInviteOrderQueryResponseModel struct {
	// 商家支付宝账号对应的ID，2088开头
	MerchantPid *string `json:"merchant_pid,omitempty"`
	// 产品签约状态对象
	SignStatusList []ProductInviteStatusInfo `json:"sign_status_list,omitempty"`
}

// NewAlipayOpenInviteOrderQueryResponseModel instantiates a new AlipayOpenInviteOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenInviteOrderQueryResponseModel() *AlipayOpenInviteOrderQueryResponseModel {
	this := AlipayOpenInviteOrderQueryResponseModel{}
	return &this
}

// NewAlipayOpenInviteOrderQueryResponseModelWithDefaults instantiates a new AlipayOpenInviteOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenInviteOrderQueryResponseModelWithDefaults() *AlipayOpenInviteOrderQueryResponseModel {
	this := AlipayOpenInviteOrderQueryResponseModel{}
	return &this
}

// GetMerchantPid returns the MerchantPid field value if set, zero value otherwise.
func (o *AlipayOpenInviteOrderQueryResponseModel) GetMerchantPid() string {
	if o == nil || IsNil(o.MerchantPid) {
		var ret string
		return ret
	}
	return *o.MerchantPid
}

// GetMerchantPidOk returns a tuple with the MerchantPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInviteOrderQueryResponseModel) GetMerchantPidOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantPid) {
		return nil, false
	}
	return o.MerchantPid, true
}

// HasMerchantPid returns a boolean if a field has been set.
func (o *AlipayOpenInviteOrderQueryResponseModel) HasMerchantPid() bool {
	if o != nil && !IsNil(o.MerchantPid) {
		return true
	}

	return false
}

// SetMerchantPid gets a reference to the given string and assigns it to the MerchantPid field.
func (o *AlipayOpenInviteOrderQueryResponseModel) SetMerchantPid(v string) {
	o.MerchantPid = &v
}

// GetSignStatusList returns the SignStatusList field value if set, zero value otherwise.
func (o *AlipayOpenInviteOrderQueryResponseModel) GetSignStatusList() []ProductInviteStatusInfo {
	if o == nil || IsNil(o.SignStatusList) {
		var ret []ProductInviteStatusInfo
		return ret
	}
	return o.SignStatusList
}

// GetSignStatusListOk returns a tuple with the SignStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenInviteOrderQueryResponseModel) GetSignStatusListOk() ([]ProductInviteStatusInfo, bool) {
	if o == nil || IsNil(o.SignStatusList) {
		return nil, false
	}
	return o.SignStatusList, true
}

// HasSignStatusList returns a boolean if a field has been set.
func (o *AlipayOpenInviteOrderQueryResponseModel) HasSignStatusList() bool {
	if o != nil && !IsNil(o.SignStatusList) {
		return true
	}

	return false
}

// SetSignStatusList gets a reference to the given []ProductInviteStatusInfo and assigns it to the SignStatusList field.
func (o *AlipayOpenInviteOrderQueryResponseModel) SetSignStatusList(v []ProductInviteStatusInfo) {
	o.SignStatusList = v
}

func (o AlipayOpenInviteOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenInviteOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantPid) {
		toSerialize["merchant_pid"] = o.MerchantPid
	}
	if !IsNil(o.SignStatusList) {
		toSerialize["sign_status_list"] = o.SignStatusList
	}
	return toSerialize, nil
}

type NullableAlipayOpenInviteOrderQueryResponseModel struct {
	value *AlipayOpenInviteOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenInviteOrderQueryResponseModel) Get() *AlipayOpenInviteOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenInviteOrderQueryResponseModel) Set(val *AlipayOpenInviteOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenInviteOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenInviteOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenInviteOrderQueryResponseModel(val *AlipayOpenInviteOrderQueryResponseModel) *NullableAlipayOpenInviteOrderQueryResponseModel {
	return &NullableAlipayOpenInviteOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenInviteOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenInviteOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
