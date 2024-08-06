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

// checks if the OpenFormFieldDO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenFormFieldDO{}

// OpenFormFieldDO struct for OpenFormFieldDO
type OpenFormFieldDO struct {
	// 表单可选字段配置，common_fields属性定义一个表单字段数组，表单字段有效值列表与required字段有效值列表相同。  可选字段配置中不能含有必须字段配置的有效值。
	Optional *string `json:"optional,omitempty"`
	// 表单必填字段配置，common_fields属性定义一个表单字段数组，字段有效值如下列表所示：  OPEN_FORM_FIELD_MOBILE -- 手机号  OPEN_FORM_FIELD_GENDER -- 性别  OPEN_FORM_FIELD_NAME -- 姓名  OPEN_FORM_FIELD_BIRTHDAY -- 生日  OPEN_FORM_FIELD_IDCARD -- 身份证  OPEN_FORM_FIELD_EMAIL -- 邮箱  OPEN_FORM_FIELD_ADDRESS -- 地址  OPEN_FORM_FIELD_CITY -- 城市  OPEN_FORM_FIELD_IS_STUDENT -- 是否学生认证  OPEN_FORM_FIELD_MEMBER_GRADE -- 会员等级
	Required *string `json:"required,omitempty"`
}

// NewOpenFormFieldDO instantiates a new OpenFormFieldDO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenFormFieldDO() *OpenFormFieldDO {
	this := OpenFormFieldDO{}
	return &this
}

// NewOpenFormFieldDOWithDefaults instantiates a new OpenFormFieldDO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenFormFieldDOWithDefaults() *OpenFormFieldDO {
	this := OpenFormFieldDO{}
	return &this
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *OpenFormFieldDO) GetOptional() string {
	if o == nil || IsNil(o.Optional) {
		var ret string
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenFormFieldDO) GetOptionalOk() (*string, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *OpenFormFieldDO) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given string and assigns it to the Optional field.
func (o *OpenFormFieldDO) SetOptional(v string) {
	o.Optional = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *OpenFormFieldDO) GetRequired() string {
	if o == nil || IsNil(o.Required) {
		var ret string
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenFormFieldDO) GetRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *OpenFormFieldDO) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given string and assigns it to the Required field.
func (o *OpenFormFieldDO) SetRequired(v string) {
	o.Required = &v
}

func (o OpenFormFieldDO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenFormFieldDO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableOpenFormFieldDO struct {
	value *OpenFormFieldDO
	isSet bool
}

func (v NullableOpenFormFieldDO) Get() *OpenFormFieldDO {
	return v.value
}

func (v *NullableOpenFormFieldDO) Set(val *OpenFormFieldDO) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenFormFieldDO) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenFormFieldDO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenFormFieldDO(val *OpenFormFieldDO) *NullableOpenFormFieldDO {
	return &NullableOpenFormFieldDO{value: val, isSet: true}
}

func (v NullableOpenFormFieldDO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenFormFieldDO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


