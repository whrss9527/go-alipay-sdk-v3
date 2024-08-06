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

// checks if the Option type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Option{}

// Option struct for Option
type Option struct {
	// 文本，通常用于理解对应的取值
	Text *string `json:"text,omitempty"`
	// 取值，通常使用简单的数字或字符串
	Value *string `json:"value,omitempty"`
}

// NewOption instantiates a new Option object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOption() *Option {
	this := Option{}
	return &this
}

// NewOptionWithDefaults instantiates a new Option object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionWithDefaults() *Option {
	this := Option{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Option) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Option) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Option) SetText(v string) {
	o.Text = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Option) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Option) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Option) SetValue(v string) {
	o.Value = &v
}

func (o Option) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Option) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableOption struct {
	value *Option
	isSet bool
}

func (v NullableOption) Get() *Option {
	return v.value
}

func (v *NullableOption) Set(val *Option) {
	v.value = val
	v.isSet = true
}

func (v NullableOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOption(val *Option) *NullableOption {
	return &NullableOption{value: val, isSet: true}
}

func (v NullableOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
