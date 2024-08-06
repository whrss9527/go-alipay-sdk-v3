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

// checks if the AlipayOpenPublicTemplateMessageIndustryModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicTemplateMessageIndustryModifyModel{}

// AlipayOpenPublicTemplateMessageIndustryModifyModel struct for AlipayOpenPublicTemplateMessageIndustryModifyModel
type AlipayOpenPublicTemplateMessageIndustryModifyModel struct {
	// 服务窗消息模板所属主行业一/二级编码，参见 <a href=\"https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043\">查看行业信息</a>。
	PrimaryIndustryCode *string `json:"primary_industry_code,omitempty"`
	// 服务窗消息模板所属主行业一/二级名称，参见 <a href=\"https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043\">查看行业信息</a>。
	PrimaryIndustryName *string `json:"primary_industry_name,omitempty"`
	// 服务窗消息模板所属副行业一/二级编码，参见 <a href=\"https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043\">查看行业信息</a>。
	SecondaryIndustryCode *string `json:"secondary_industry_code,omitempty"`
	// 服务窗消息模板所属副行业一/二级名称，参见 <a href=\"https://alipay.open.taobao.com/doc2/detail?treeId=197&docType=1&articleId=105043\">查看行业信息</a>。
	SecondaryIndustryName *string `json:"secondary_industry_name,omitempty"`
}

// NewAlipayOpenPublicTemplateMessageIndustryModifyModel instantiates a new AlipayOpenPublicTemplateMessageIndustryModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicTemplateMessageIndustryModifyModel() *AlipayOpenPublicTemplateMessageIndustryModifyModel {
	this := AlipayOpenPublicTemplateMessageIndustryModifyModel{}
	return &this
}

// NewAlipayOpenPublicTemplateMessageIndustryModifyModelWithDefaults instantiates a new AlipayOpenPublicTemplateMessageIndustryModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicTemplateMessageIndustryModifyModelWithDefaults() *AlipayOpenPublicTemplateMessageIndustryModifyModel {
	this := AlipayOpenPublicTemplateMessageIndustryModifyModel{}
	return &this
}

// GetPrimaryIndustryCode returns the PrimaryIndustryCode field value if set, zero value otherwise.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetPrimaryIndustryCode() string {
	if o == nil || IsNil(o.PrimaryIndustryCode) {
		var ret string
		return ret
	}
	return *o.PrimaryIndustryCode
}

// GetPrimaryIndustryCodeOk returns a tuple with the PrimaryIndustryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetPrimaryIndustryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryIndustryCode) {
		return nil, false
	}
	return o.PrimaryIndustryCode, true
}

// HasPrimaryIndustryCode returns a boolean if a field has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) HasPrimaryIndustryCode() bool {
	if o != nil && !IsNil(o.PrimaryIndustryCode) {
		return true
	}

	return false
}

// SetPrimaryIndustryCode gets a reference to the given string and assigns it to the PrimaryIndustryCode field.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) SetPrimaryIndustryCode(v string) {
	o.PrimaryIndustryCode = &v
}

// GetPrimaryIndustryName returns the PrimaryIndustryName field value if set, zero value otherwise.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetPrimaryIndustryName() string {
	if o == nil || IsNil(o.PrimaryIndustryName) {
		var ret string
		return ret
	}
	return *o.PrimaryIndustryName
}

// GetPrimaryIndustryNameOk returns a tuple with the PrimaryIndustryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetPrimaryIndustryNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryIndustryName) {
		return nil, false
	}
	return o.PrimaryIndustryName, true
}

// HasPrimaryIndustryName returns a boolean if a field has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) HasPrimaryIndustryName() bool {
	if o != nil && !IsNil(o.PrimaryIndustryName) {
		return true
	}

	return false
}

// SetPrimaryIndustryName gets a reference to the given string and assigns it to the PrimaryIndustryName field.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) SetPrimaryIndustryName(v string) {
	o.PrimaryIndustryName = &v
}

// GetSecondaryIndustryCode returns the SecondaryIndustryCode field value if set, zero value otherwise.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetSecondaryIndustryCode() string {
	if o == nil || IsNil(o.SecondaryIndustryCode) {
		var ret string
		return ret
	}
	return *o.SecondaryIndustryCode
}

// GetSecondaryIndustryCodeOk returns a tuple with the SecondaryIndustryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetSecondaryIndustryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryIndustryCode) {
		return nil, false
	}
	return o.SecondaryIndustryCode, true
}

// HasSecondaryIndustryCode returns a boolean if a field has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) HasSecondaryIndustryCode() bool {
	if o != nil && !IsNil(o.SecondaryIndustryCode) {
		return true
	}

	return false
}

// SetSecondaryIndustryCode gets a reference to the given string and assigns it to the SecondaryIndustryCode field.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) SetSecondaryIndustryCode(v string) {
	o.SecondaryIndustryCode = &v
}

// GetSecondaryIndustryName returns the SecondaryIndustryName field value if set, zero value otherwise.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetSecondaryIndustryName() string {
	if o == nil || IsNil(o.SecondaryIndustryName) {
		var ret string
		return ret
	}
	return *o.SecondaryIndustryName
}

// GetSecondaryIndustryNameOk returns a tuple with the SecondaryIndustryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) GetSecondaryIndustryNameOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryIndustryName) {
		return nil, false
	}
	return o.SecondaryIndustryName, true
}

// HasSecondaryIndustryName returns a boolean if a field has been set.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) HasSecondaryIndustryName() bool {
	if o != nil && !IsNil(o.SecondaryIndustryName) {
		return true
	}

	return false
}

// SetSecondaryIndustryName gets a reference to the given string and assigns it to the SecondaryIndustryName field.
func (o *AlipayOpenPublicTemplateMessageIndustryModifyModel) SetSecondaryIndustryName(v string) {
	o.SecondaryIndustryName = &v
}

func (o AlipayOpenPublicTemplateMessageIndustryModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicTemplateMessageIndustryModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryIndustryCode) {
		toSerialize["primary_industry_code"] = o.PrimaryIndustryCode
	}
	if !IsNil(o.PrimaryIndustryName) {
		toSerialize["primary_industry_name"] = o.PrimaryIndustryName
	}
	if !IsNil(o.SecondaryIndustryCode) {
		toSerialize["secondary_industry_code"] = o.SecondaryIndustryCode
	}
	if !IsNil(o.SecondaryIndustryName) {
		toSerialize["secondary_industry_name"] = o.SecondaryIndustryName
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicTemplateMessageIndustryModifyModel struct {
	value *AlipayOpenPublicTemplateMessageIndustryModifyModel
	isSet bool
}

func (v NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) Get() *AlipayOpenPublicTemplateMessageIndustryModifyModel {
	return v.value
}

func (v *NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) Set(val *AlipayOpenPublicTemplateMessageIndustryModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicTemplateMessageIndustryModifyModel(val *AlipayOpenPublicTemplateMessageIndustryModifyModel) *NullableAlipayOpenPublicTemplateMessageIndustryModifyModel {
	return &NullableAlipayOpenPublicTemplateMessageIndustryModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicTemplateMessageIndustryModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


