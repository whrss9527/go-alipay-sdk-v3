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

// checks if the DeliveryActivityContentInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryActivityContentInfo{}

// DeliveryActivityContentInfo struct for DeliveryActivityContentInfo
type DeliveryActivityContentInfo struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// NewDeliveryActivityContentInfo instantiates a new DeliveryActivityContentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryActivityContentInfo() *DeliveryActivityContentInfo {
	this := DeliveryActivityContentInfo{}
	return &this
}

// NewDeliveryActivityContentInfoWithDefaults instantiates a new DeliveryActivityContentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryActivityContentInfoWithDefaults() *DeliveryActivityContentInfo {
	this := DeliveryActivityContentInfo{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *DeliveryActivityContentInfo) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryActivityContentInfo) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *DeliveryActivityContentInfo) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *DeliveryActivityContentInfo) SetActivityId(v string) {
	o.ActivityId = &v
}

func (o DeliveryActivityContentInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryActivityContentInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	return toSerialize, nil
}

type NullableDeliveryActivityContentInfo struct {
	value *DeliveryActivityContentInfo
	isSet bool
}

func (v NullableDeliveryActivityContentInfo) Get() *DeliveryActivityContentInfo {
	return v.value
}

func (v *NullableDeliveryActivityContentInfo) Set(val *DeliveryActivityContentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryActivityContentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryActivityContentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryActivityContentInfo(val *DeliveryActivityContentInfo) *NullableDeliveryActivityContentInfo {
	return &NullableDeliveryActivityContentInfo{value: val, isSet: true}
}

func (v NullableDeliveryActivityContentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryActivityContentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
