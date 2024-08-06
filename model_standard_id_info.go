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

// checks if the StandardIdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardIdInfo{}

// StandardIdInfo struct for StandardIdInfo
type StandardIdInfo struct {
	// 外部使用规则id
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// 使用规则id
	StandardId *string `json:"standard_id,omitempty"`
}

// NewStandardIdInfo instantiates a new StandardIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardIdInfo() *StandardIdInfo {
	this := StandardIdInfo{}
	return &this
}

// NewStandardIdInfoWithDefaults instantiates a new StandardIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardIdInfoWithDefaults() *StandardIdInfo {
	this := StandardIdInfo{}
	return &this
}

// GetOuterSourceId returns the OuterSourceId field value if set, zero value otherwise.
func (o *StandardIdInfo) GetOuterSourceId() string {
	if o == nil || IsNil(o.OuterSourceId) {
		var ret string
		return ret
	}
	return *o.OuterSourceId
}

// GetOuterSourceIdOk returns a tuple with the OuterSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardIdInfo) GetOuterSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OuterSourceId) {
		return nil, false
	}
	return o.OuterSourceId, true
}

// HasOuterSourceId returns a boolean if a field has been set.
func (o *StandardIdInfo) HasOuterSourceId() bool {
	if o != nil && !IsNil(o.OuterSourceId) {
		return true
	}

	return false
}

// SetOuterSourceId gets a reference to the given string and assigns it to the OuterSourceId field.
func (o *StandardIdInfo) SetOuterSourceId(v string) {
	o.OuterSourceId = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *StandardIdInfo) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardIdInfo) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *StandardIdInfo) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *StandardIdInfo) SetStandardId(v string) {
	o.StandardId = &v
}

func (o StandardIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardIdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OuterSourceId) {
		toSerialize["outer_source_id"] = o.OuterSourceId
	}
	if !IsNil(o.StandardId) {
		toSerialize["standard_id"] = o.StandardId
	}
	return toSerialize, nil
}

type NullableStandardIdInfo struct {
	value *StandardIdInfo
	isSet bool
}

func (v NullableStandardIdInfo) Get() *StandardIdInfo {
	return v.value
}

func (v *NullableStandardIdInfo) Set(val *StandardIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardIdInfo(val *StandardIdInfo) *NullableStandardIdInfo {
	return &NullableStandardIdInfo{value: val, isSet: true}
}

func (v NullableStandardIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


