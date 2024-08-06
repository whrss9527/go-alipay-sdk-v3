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

// checks if the DeliverySearchBoxRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverySearchBoxRule{}

// DeliverySearchBoxRule struct for DeliverySearchBoxRule
type DeliverySearchBoxRule struct {
	// 指定搜索直达区域
	BoxId *string `json:"box_id,omitempty"`
}

// NewDeliverySearchBoxRule instantiates a new DeliverySearchBoxRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverySearchBoxRule() *DeliverySearchBoxRule {
	this := DeliverySearchBoxRule{}
	return &this
}

// NewDeliverySearchBoxRuleWithDefaults instantiates a new DeliverySearchBoxRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverySearchBoxRuleWithDefaults() *DeliverySearchBoxRule {
	this := DeliverySearchBoxRule{}
	return &this
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *DeliverySearchBoxRule) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliverySearchBoxRule) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *DeliverySearchBoxRule) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *DeliverySearchBoxRule) SetBoxId(v string) {
	o.BoxId = &v
}

func (o DeliverySearchBoxRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverySearchBoxRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	return toSerialize, nil
}

type NullableDeliverySearchBoxRule struct {
	value *DeliverySearchBoxRule
	isSet bool
}

func (v NullableDeliverySearchBoxRule) Get() *DeliverySearchBoxRule {
	return v.value
}

func (v *NullableDeliverySearchBoxRule) Set(val *DeliverySearchBoxRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverySearchBoxRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverySearchBoxRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverySearchBoxRule(val *DeliverySearchBoxRule) *NullableDeliverySearchBoxRule {
	return &NullableDeliverySearchBoxRule{value: val, isSet: true}
}

func (v NullableDeliverySearchBoxRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverySearchBoxRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


