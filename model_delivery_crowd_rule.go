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

// checks if the DeliveryCrowdRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryCrowdRule{}

// DeliveryCrowdRule struct for DeliveryCrowdRule
type DeliveryCrowdRule struct {
	// 指定人群ID
	CrowdId *string `json:"crowd_id,omitempty"`
}

// NewDeliveryCrowdRule instantiates a new DeliveryCrowdRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryCrowdRule() *DeliveryCrowdRule {
	this := DeliveryCrowdRule{}
	return &this
}

// NewDeliveryCrowdRuleWithDefaults instantiates a new DeliveryCrowdRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryCrowdRuleWithDefaults() *DeliveryCrowdRule {
	this := DeliveryCrowdRule{}
	return &this
}

// GetCrowdId returns the CrowdId field value if set, zero value otherwise.
func (o *DeliveryCrowdRule) GetCrowdId() string {
	if o == nil || IsNil(o.CrowdId) {
		var ret string
		return ret
	}
	return *o.CrowdId
}

// GetCrowdIdOk returns a tuple with the CrowdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryCrowdRule) GetCrowdIdOk() (*string, bool) {
	if o == nil || IsNil(o.CrowdId) {
		return nil, false
	}
	return o.CrowdId, true
}

// HasCrowdId returns a boolean if a field has been set.
func (o *DeliveryCrowdRule) HasCrowdId() bool {
	if o != nil && !IsNil(o.CrowdId) {
		return true
	}

	return false
}

// SetCrowdId gets a reference to the given string and assigns it to the CrowdId field.
func (o *DeliveryCrowdRule) SetCrowdId(v string) {
	o.CrowdId = &v
}

func (o DeliveryCrowdRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryCrowdRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CrowdId) {
		toSerialize["crowd_id"] = o.CrowdId
	}
	return toSerialize, nil
}

type NullableDeliveryCrowdRule struct {
	value *DeliveryCrowdRule
	isSet bool
}

func (v NullableDeliveryCrowdRule) Get() *DeliveryCrowdRule {
	return v.value
}

func (v *NullableDeliveryCrowdRule) Set(val *DeliveryCrowdRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryCrowdRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryCrowdRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryCrowdRule(val *DeliveryCrowdRule) *NullableDeliveryCrowdRule {
	return &NullableDeliveryCrowdRule{value: val, isSet: true}
}

func (v NullableDeliveryCrowdRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryCrowdRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

