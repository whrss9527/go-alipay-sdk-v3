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

// checks if the CloudResumeSkillInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudResumeSkillInfo{}

// CloudResumeSkillInfo struct for CloudResumeSkillInfo
type CloudResumeSkillInfo struct {
	// 技能标签名字
	SkillName *string `json:"skill_name,omitempty"`
}

// NewCloudResumeSkillInfo instantiates a new CloudResumeSkillInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResumeSkillInfo() *CloudResumeSkillInfo {
	this := CloudResumeSkillInfo{}
	return &this
}

// NewCloudResumeSkillInfoWithDefaults instantiates a new CloudResumeSkillInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResumeSkillInfoWithDefaults() *CloudResumeSkillInfo {
	this := CloudResumeSkillInfo{}
	return &this
}

// GetSkillName returns the SkillName field value if set, zero value otherwise.
func (o *CloudResumeSkillInfo) GetSkillName() string {
	if o == nil || IsNil(o.SkillName) {
		var ret string
		return ret
	}
	return *o.SkillName
}

// GetSkillNameOk returns a tuple with the SkillName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResumeSkillInfo) GetSkillNameOk() (*string, bool) {
	if o == nil || IsNil(o.SkillName) {
		return nil, false
	}
	return o.SkillName, true
}

// HasSkillName returns a boolean if a field has been set.
func (o *CloudResumeSkillInfo) HasSkillName() bool {
	if o != nil && !IsNil(o.SkillName) {
		return true
	}

	return false
}

// SetSkillName gets a reference to the given string and assigns it to the SkillName field.
func (o *CloudResumeSkillInfo) SetSkillName(v string) {
	o.SkillName = &v
}

func (o CloudResumeSkillInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudResumeSkillInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkillName) {
		toSerialize["skill_name"] = o.SkillName
	}
	return toSerialize, nil
}

type NullableCloudResumeSkillInfo struct {
	value *CloudResumeSkillInfo
	isSet bool
}

func (v NullableCloudResumeSkillInfo) Get() *CloudResumeSkillInfo {
	return v.value
}

func (v *NullableCloudResumeSkillInfo) Set(val *CloudResumeSkillInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResumeSkillInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResumeSkillInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResumeSkillInfo(val *CloudResumeSkillInfo) *NullableCloudResumeSkillInfo {
	return &NullableCloudResumeSkillInfo{value: val, isSet: true}
}

func (v NullableCloudResumeSkillInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResumeSkillInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
