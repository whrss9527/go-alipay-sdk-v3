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

// checks if the AlipayOpenPublicMenuBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMenuBatchqueryResponseModel{}

// AlipayOpenPublicMenuBatchqueryResponseModel struct for AlipayOpenPublicMenuBatchqueryResponseModel
type AlipayOpenPublicMenuBatchqueryResponseModel struct {
	// 菜单数量，包括默认菜单和个性化菜单
	Count *string `json:"count,omitempty"`
	// 菜单列表
	Menus []QueryMenu `json:"menus,omitempty"`
}

// NewAlipayOpenPublicMenuBatchqueryResponseModel instantiates a new AlipayOpenPublicMenuBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMenuBatchqueryResponseModel() *AlipayOpenPublicMenuBatchqueryResponseModel {
	this := AlipayOpenPublicMenuBatchqueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicMenuBatchqueryResponseModelWithDefaults instantiates a new AlipayOpenPublicMenuBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMenuBatchqueryResponseModelWithDefaults() *AlipayOpenPublicMenuBatchqueryResponseModel {
	this := AlipayOpenPublicMenuBatchqueryResponseModel{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) SetCount(v string) {
	o.Count = &v
}

// GetMenus returns the Menus field value if set, zero value otherwise.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) GetMenus() []QueryMenu {
	if o == nil || IsNil(o.Menus) {
		var ret []QueryMenu
		return ret
	}
	return o.Menus
}

// GetMenusOk returns a tuple with the Menus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) GetMenusOk() ([]QueryMenu, bool) {
	if o == nil || IsNil(o.Menus) {
		return nil, false
	}
	return o.Menus, true
}

// HasMenus returns a boolean if a field has been set.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) HasMenus() bool {
	if o != nil && !IsNil(o.Menus) {
		return true
	}

	return false
}

// SetMenus gets a reference to the given []QueryMenu and assigns it to the Menus field.
func (o *AlipayOpenPublicMenuBatchqueryResponseModel) SetMenus(v []QueryMenu) {
	o.Menus = v
}

func (o AlipayOpenPublicMenuBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMenuBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Menus) {
		toSerialize["menus"] = o.Menus
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMenuBatchqueryResponseModel struct {
	value *AlipayOpenPublicMenuBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicMenuBatchqueryResponseModel) Get() *AlipayOpenPublicMenuBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuBatchqueryResponseModel) Set(val *AlipayOpenPublicMenuBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuBatchqueryResponseModel(val *AlipayOpenPublicMenuBatchqueryResponseModel) *NullableAlipayOpenPublicMenuBatchqueryResponseModel {
	return &NullableAlipayOpenPublicMenuBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
