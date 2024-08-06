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

// checks if the AlipayOpenPublicSettingCategoryQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicSettingCategoryQueryResponseModel{}

// AlipayOpenPublicSettingCategoryQueryResponseModel struct for AlipayOpenPublicSettingCategoryQueryResponseModel
type AlipayOpenPublicSettingCategoryQueryResponseModel struct {
	// 已设置的一级行业分类名称。
	PrimaryCategory *string `json:"primary_category,omitempty"`
	// 已设置的二级行业分类名称。
	SecondaryCategory *string `json:"secondary_category,omitempty"`
}

// NewAlipayOpenPublicSettingCategoryQueryResponseModel instantiates a new AlipayOpenPublicSettingCategoryQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicSettingCategoryQueryResponseModel() *AlipayOpenPublicSettingCategoryQueryResponseModel {
	this := AlipayOpenPublicSettingCategoryQueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicSettingCategoryQueryResponseModelWithDefaults instantiates a new AlipayOpenPublicSettingCategoryQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicSettingCategoryQueryResponseModelWithDefaults() *AlipayOpenPublicSettingCategoryQueryResponseModel {
	this := AlipayOpenPublicSettingCategoryQueryResponseModel{}
	return &this
}

// GetPrimaryCategory returns the PrimaryCategory field value if set, zero value otherwise.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) GetPrimaryCategory() string {
	if o == nil || IsNil(o.PrimaryCategory) {
		var ret string
		return ret
	}
	return *o.PrimaryCategory
}

// GetPrimaryCategoryOk returns a tuple with the PrimaryCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) GetPrimaryCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryCategory) {
		return nil, false
	}
	return o.PrimaryCategory, true
}

// HasPrimaryCategory returns a boolean if a field has been set.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) HasPrimaryCategory() bool {
	if o != nil && !IsNil(o.PrimaryCategory) {
		return true
	}

	return false
}

// SetPrimaryCategory gets a reference to the given string and assigns it to the PrimaryCategory field.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) SetPrimaryCategory(v string) {
	o.PrimaryCategory = &v
}

// GetSecondaryCategory returns the SecondaryCategory field value if set, zero value otherwise.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) GetSecondaryCategory() string {
	if o == nil || IsNil(o.SecondaryCategory) {
		var ret string
		return ret
	}
	return *o.SecondaryCategory
}

// GetSecondaryCategoryOk returns a tuple with the SecondaryCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) GetSecondaryCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryCategory) {
		return nil, false
	}
	return o.SecondaryCategory, true
}

// HasSecondaryCategory returns a boolean if a field has been set.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) HasSecondaryCategory() bool {
	if o != nil && !IsNil(o.SecondaryCategory) {
		return true
	}

	return false
}

// SetSecondaryCategory gets a reference to the given string and assigns it to the SecondaryCategory field.
func (o *AlipayOpenPublicSettingCategoryQueryResponseModel) SetSecondaryCategory(v string) {
	o.SecondaryCategory = &v
}

func (o AlipayOpenPublicSettingCategoryQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicSettingCategoryQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryCategory) {
		toSerialize["primary_category"] = o.PrimaryCategory
	}
	if !IsNil(o.SecondaryCategory) {
		toSerialize["secondary_category"] = o.SecondaryCategory
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicSettingCategoryQueryResponseModel struct {
	value *AlipayOpenPublicSettingCategoryQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicSettingCategoryQueryResponseModel) Get() *AlipayOpenPublicSettingCategoryQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryResponseModel) Set(val *AlipayOpenPublicSettingCategoryQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicSettingCategoryQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicSettingCategoryQueryResponseModel(val *AlipayOpenPublicSettingCategoryQueryResponseModel) *NullableAlipayOpenPublicSettingCategoryQueryResponseModel {
	return &NullableAlipayOpenPublicSettingCategoryQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicSettingCategoryQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicSettingCategoryQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
