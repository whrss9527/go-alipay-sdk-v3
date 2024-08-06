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

// checks if the MerchantCardMsgInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantCardMsgInfo{}

// MerchantCardMsgInfo struct for MerchantCardMsgInfo
type MerchantCardMsgInfo struct {
	// 本次消息通知的变动积分，积分必须为数字型（可为浮点型，带2位小数点），格式非法则不发送消息。例如： 100，则代表增加100积分。 -100，则代表减少100积分。
	ChangedPoint *string `json:"changed_point,omitempty"`
}

// NewMerchantCardMsgInfo instantiates a new MerchantCardMsgInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantCardMsgInfo() *MerchantCardMsgInfo {
	this := MerchantCardMsgInfo{}
	return &this
}

// NewMerchantCardMsgInfoWithDefaults instantiates a new MerchantCardMsgInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantCardMsgInfoWithDefaults() *MerchantCardMsgInfo {
	this := MerchantCardMsgInfo{}
	return &this
}

// GetChangedPoint returns the ChangedPoint field value if set, zero value otherwise.
func (o *MerchantCardMsgInfo) GetChangedPoint() string {
	if o == nil || IsNil(o.ChangedPoint) {
		var ret string
		return ret
	}
	return *o.ChangedPoint
}

// GetChangedPointOk returns a tuple with the ChangedPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantCardMsgInfo) GetChangedPointOk() (*string, bool) {
	if o == nil || IsNil(o.ChangedPoint) {
		return nil, false
	}
	return o.ChangedPoint, true
}

// HasChangedPoint returns a boolean if a field has been set.
func (o *MerchantCardMsgInfo) HasChangedPoint() bool {
	if o != nil && !IsNil(o.ChangedPoint) {
		return true
	}

	return false
}

// SetChangedPoint gets a reference to the given string and assigns it to the ChangedPoint field.
func (o *MerchantCardMsgInfo) SetChangedPoint(v string) {
	o.ChangedPoint = &v
}

func (o MerchantCardMsgInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantCardMsgInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangedPoint) {
		toSerialize["changed_point"] = o.ChangedPoint
	}
	return toSerialize, nil
}

type NullableMerchantCardMsgInfo struct {
	value *MerchantCardMsgInfo
	isSet bool
}

func (v NullableMerchantCardMsgInfo) Get() *MerchantCardMsgInfo {
	return v.value
}

func (v *NullableMerchantCardMsgInfo) Set(val *MerchantCardMsgInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantCardMsgInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantCardMsgInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantCardMsgInfo(val *MerchantCardMsgInfo) *NullableMerchantCardMsgInfo {
	return &NullableMerchantCardMsgInfo{value: val, isSet: true}
}

func (v NullableMerchantCardMsgInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantCardMsgInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


