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

// checks if the IndirectAuthOrderFailedReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndirectAuthOrderFailedReason{}

// IndirectAuthOrderFailedReason struct for IndirectAuthOrderFailedReason
type IndirectAuthOrderFailedReason struct {
	// 审核失败字段
	FailParam *string `json:"fail_param,omitempty"`
	// 描述申请单审核失败原因
	FailReason *string `json:"fail_reason,omitempty"`
}

// NewIndirectAuthOrderFailedReason instantiates a new IndirectAuthOrderFailedReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndirectAuthOrderFailedReason() *IndirectAuthOrderFailedReason {
	this := IndirectAuthOrderFailedReason{}
	return &this
}

// NewIndirectAuthOrderFailedReasonWithDefaults instantiates a new IndirectAuthOrderFailedReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndirectAuthOrderFailedReasonWithDefaults() *IndirectAuthOrderFailedReason {
	this := IndirectAuthOrderFailedReason{}
	return &this
}

// GetFailParam returns the FailParam field value if set, zero value otherwise.
func (o *IndirectAuthOrderFailedReason) GetFailParam() string {
	if o == nil || IsNil(o.FailParam) {
		var ret string
		return ret
	}
	return *o.FailParam
}

// GetFailParamOk returns a tuple with the FailParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectAuthOrderFailedReason) GetFailParamOk() (*string, bool) {
	if o == nil || IsNil(o.FailParam) {
		return nil, false
	}
	return o.FailParam, true
}

// HasFailParam returns a boolean if a field has been set.
func (o *IndirectAuthOrderFailedReason) HasFailParam() bool {
	if o != nil && !IsNil(o.FailParam) {
		return true
	}

	return false
}

// SetFailParam gets a reference to the given string and assigns it to the FailParam field.
func (o *IndirectAuthOrderFailedReason) SetFailParam(v string) {
	o.FailParam = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *IndirectAuthOrderFailedReason) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndirectAuthOrderFailedReason) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *IndirectAuthOrderFailedReason) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *IndirectAuthOrderFailedReason) SetFailReason(v string) {
	o.FailReason = &v
}

func (o IndirectAuthOrderFailedReason) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndirectAuthOrderFailedReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailParam) {
		toSerialize["fail_param"] = o.FailParam
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	return toSerialize, nil
}

type NullableIndirectAuthOrderFailedReason struct {
	value *IndirectAuthOrderFailedReason
	isSet bool
}

func (v NullableIndirectAuthOrderFailedReason) Get() *IndirectAuthOrderFailedReason {
	return v.value
}

func (v *NullableIndirectAuthOrderFailedReason) Set(val *IndirectAuthOrderFailedReason) {
	v.value = val
	v.isSet = true
}

func (v NullableIndirectAuthOrderFailedReason) IsSet() bool {
	return v.isSet
}

func (v *NullableIndirectAuthOrderFailedReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndirectAuthOrderFailedReason(val *IndirectAuthOrderFailedReason) *NullableIndirectAuthOrderFailedReason {
	return &NullableIndirectAuthOrderFailedReason{value: val, isSet: true}
}

func (v NullableIndirectAuthOrderFailedReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndirectAuthOrderFailedReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

