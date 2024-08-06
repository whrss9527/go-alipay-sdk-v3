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

// checks if the AlipayOpenPublicLabelCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicLabelCreateModel{}

// AlipayOpenPublicLabelCreateModel struct for AlipayOpenPublicLabelCreateModel
type AlipayOpenPublicLabelCreateModel struct {
	// 标签名
	Name *string `json:"name,omitempty"`
}

// NewAlipayOpenPublicLabelCreateModel instantiates a new AlipayOpenPublicLabelCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicLabelCreateModel() *AlipayOpenPublicLabelCreateModel {
	this := AlipayOpenPublicLabelCreateModel{}
	return &this
}

// NewAlipayOpenPublicLabelCreateModelWithDefaults instantiates a new AlipayOpenPublicLabelCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicLabelCreateModelWithDefaults() *AlipayOpenPublicLabelCreateModel {
	this := AlipayOpenPublicLabelCreateModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayOpenPublicLabelCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicLabelCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayOpenPublicLabelCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayOpenPublicLabelCreateModel) SetName(v string) {
	o.Name = &v
}

func (o AlipayOpenPublicLabelCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicLabelCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicLabelCreateModel struct {
	value *AlipayOpenPublicLabelCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicLabelCreateModel) Get() *AlipayOpenPublicLabelCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicLabelCreateModel) Set(val *AlipayOpenPublicLabelCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLabelCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLabelCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLabelCreateModel(val *AlipayOpenPublicLabelCreateModel) *NullableAlipayOpenPublicLabelCreateModel {
	return &NullableAlipayOpenPublicLabelCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLabelCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLabelCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
