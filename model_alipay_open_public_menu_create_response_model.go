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

// checks if the AlipayOpenPublicMenuCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMenuCreateResponseModel{}

// AlipayOpenPublicMenuCreateResponseModel struct for AlipayOpenPublicMenuCreateResponseModel
type AlipayOpenPublicMenuCreateResponseModel struct {
	// 默认菜单菜单key，文本菜单为“default”，icon菜单为“iconDefault”
	MenuKey *string `json:"menu_key,omitempty"`
}

// NewAlipayOpenPublicMenuCreateResponseModel instantiates a new AlipayOpenPublicMenuCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMenuCreateResponseModel() *AlipayOpenPublicMenuCreateResponseModel {
	this := AlipayOpenPublicMenuCreateResponseModel{}
	return &this
}

// NewAlipayOpenPublicMenuCreateResponseModelWithDefaults instantiates a new AlipayOpenPublicMenuCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMenuCreateResponseModelWithDefaults() *AlipayOpenPublicMenuCreateResponseModel {
	this := AlipayOpenPublicMenuCreateResponseModel{}
	return &this
}

// GetMenuKey returns the MenuKey field value if set, zero value otherwise.
func (o *AlipayOpenPublicMenuCreateResponseModel) GetMenuKey() string {
	if o == nil || IsNil(o.MenuKey) {
		var ret string
		return ret
	}
	return *o.MenuKey
}

// GetMenuKeyOk returns a tuple with the MenuKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMenuCreateResponseModel) GetMenuKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MenuKey) {
		return nil, false
	}
	return o.MenuKey, true
}

// HasMenuKey returns a boolean if a field has been set.
func (o *AlipayOpenPublicMenuCreateResponseModel) HasMenuKey() bool {
	if o != nil && !IsNil(o.MenuKey) {
		return true
	}

	return false
}

// SetMenuKey gets a reference to the given string and assigns it to the MenuKey field.
func (o *AlipayOpenPublicMenuCreateResponseModel) SetMenuKey(v string) {
	o.MenuKey = &v
}

func (o AlipayOpenPublicMenuCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMenuCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MenuKey) {
		toSerialize["menu_key"] = o.MenuKey
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMenuCreateResponseModel struct {
	value *AlipayOpenPublicMenuCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicMenuCreateResponseModel) Get() *AlipayOpenPublicMenuCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuCreateResponseModel) Set(val *AlipayOpenPublicMenuCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuCreateResponseModel(val *AlipayOpenPublicMenuCreateResponseModel) *NullableAlipayOpenPublicMenuCreateResponseModel {
	return &NullableAlipayOpenPublicMenuCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
