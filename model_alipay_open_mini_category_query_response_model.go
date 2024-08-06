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

// checks if the AlipayOpenMiniCategoryQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniCategoryQueryResponseModel{}

// AlipayOpenMiniCategoryQueryResponseModel struct for AlipayOpenMiniCategoryQueryResponseModel
type AlipayOpenMiniCategoryQueryResponseModel struct {
	// 小程序类目列表
	CategoryList []MiniAppCategory `json:"category_list,omitempty"`
	// 小程序新类目列表。小程序新类目列表，为三级类目。创建、修改、提审传入mini_category_ids，资质相关请参考<a href=\"https://opendocs.alipay.com/mini/operation/material \">https://opendocs.alipay.com/mini/operation/material</a>
	MiniCategoryList []MiniAppCategory `json:"mini_category_list,omitempty"`
}

// NewAlipayOpenMiniCategoryQueryResponseModel instantiates a new AlipayOpenMiniCategoryQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniCategoryQueryResponseModel() *AlipayOpenMiniCategoryQueryResponseModel {
	this := AlipayOpenMiniCategoryQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniCategoryQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniCategoryQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniCategoryQueryResponseModelWithDefaults() *AlipayOpenMiniCategoryQueryResponseModel {
	this := AlipayOpenMiniCategoryQueryResponseModel{}
	return &this
}

// GetCategoryList returns the CategoryList field value if set, zero value otherwise.
func (o *AlipayOpenMiniCategoryQueryResponseModel) GetCategoryList() []MiniAppCategory {
	if o == nil || IsNil(o.CategoryList) {
		var ret []MiniAppCategory
		return ret
	}
	return o.CategoryList
}

// GetCategoryListOk returns a tuple with the CategoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniCategoryQueryResponseModel) GetCategoryListOk() ([]MiniAppCategory, bool) {
	if o == nil || IsNil(o.CategoryList) {
		return nil, false
	}
	return o.CategoryList, true
}

// HasCategoryList returns a boolean if a field has been set.
func (o *AlipayOpenMiniCategoryQueryResponseModel) HasCategoryList() bool {
	if o != nil && !IsNil(o.CategoryList) {
		return true
	}

	return false
}

// SetCategoryList gets a reference to the given []MiniAppCategory and assigns it to the CategoryList field.
func (o *AlipayOpenMiniCategoryQueryResponseModel) SetCategoryList(v []MiniAppCategory) {
	o.CategoryList = v
}

// GetMiniCategoryList returns the MiniCategoryList field value if set, zero value otherwise.
func (o *AlipayOpenMiniCategoryQueryResponseModel) GetMiniCategoryList() []MiniAppCategory {
	if o == nil || IsNil(o.MiniCategoryList) {
		var ret []MiniAppCategory
		return ret
	}
	return o.MiniCategoryList
}

// GetMiniCategoryListOk returns a tuple with the MiniCategoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniCategoryQueryResponseModel) GetMiniCategoryListOk() ([]MiniAppCategory, bool) {
	if o == nil || IsNil(o.MiniCategoryList) {
		return nil, false
	}
	return o.MiniCategoryList, true
}

// HasMiniCategoryList returns a boolean if a field has been set.
func (o *AlipayOpenMiniCategoryQueryResponseModel) HasMiniCategoryList() bool {
	if o != nil && !IsNil(o.MiniCategoryList) {
		return true
	}

	return false
}

// SetMiniCategoryList gets a reference to the given []MiniAppCategory and assigns it to the MiniCategoryList field.
func (o *AlipayOpenMiniCategoryQueryResponseModel) SetMiniCategoryList(v []MiniAppCategory) {
	o.MiniCategoryList = v
}

func (o AlipayOpenMiniCategoryQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniCategoryQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryList) {
		toSerialize["category_list"] = o.CategoryList
	}
	if !IsNil(o.MiniCategoryList) {
		toSerialize["mini_category_list"] = o.MiniCategoryList
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniCategoryQueryResponseModel struct {
	value *AlipayOpenMiniCategoryQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniCategoryQueryResponseModel) Get() *AlipayOpenMiniCategoryQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniCategoryQueryResponseModel) Set(val *AlipayOpenMiniCategoryQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniCategoryQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniCategoryQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniCategoryQueryResponseModel(val *AlipayOpenMiniCategoryQueryResponseModel) *NullableAlipayOpenMiniCategoryQueryResponseModel {
	return &NullableAlipayOpenMiniCategoryQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniCategoryQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniCategoryQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


