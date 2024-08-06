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

// checks if the AlipayOpenPublicAdvertCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicAdvertCreateModel{}

// AlipayOpenPublicAdvertCreateModel struct for AlipayOpenPublicAdvertCreateModel
type AlipayOpenPublicAdvertCreateModel struct {
	//  广告位轮播内容列表。数量限制：大于1个，小于5个。广告位轮播内容顺序：与接口传入的顺序排列一致。
	AdvertItems []AdvertItem `json:"advert_items,omitempty"`
}

// NewAlipayOpenPublicAdvertCreateModel instantiates a new AlipayOpenPublicAdvertCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicAdvertCreateModel() *AlipayOpenPublicAdvertCreateModel {
	this := AlipayOpenPublicAdvertCreateModel{}
	return &this
}

// NewAlipayOpenPublicAdvertCreateModelWithDefaults instantiates a new AlipayOpenPublicAdvertCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicAdvertCreateModelWithDefaults() *AlipayOpenPublicAdvertCreateModel {
	this := AlipayOpenPublicAdvertCreateModel{}
	return &this
}

// GetAdvertItems returns the AdvertItems field value if set, zero value otherwise.
func (o *AlipayOpenPublicAdvertCreateModel) GetAdvertItems() []AdvertItem {
	if o == nil || IsNil(o.AdvertItems) {
		var ret []AdvertItem
		return ret
	}
	return o.AdvertItems
}

// GetAdvertItemsOk returns a tuple with the AdvertItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicAdvertCreateModel) GetAdvertItemsOk() ([]AdvertItem, bool) {
	if o == nil || IsNil(o.AdvertItems) {
		return nil, false
	}
	return o.AdvertItems, true
}

// HasAdvertItems returns a boolean if a field has been set.
func (o *AlipayOpenPublicAdvertCreateModel) HasAdvertItems() bool {
	if o != nil && !IsNil(o.AdvertItems) {
		return true
	}

	return false
}

// SetAdvertItems gets a reference to the given []AdvertItem and assigns it to the AdvertItems field.
func (o *AlipayOpenPublicAdvertCreateModel) SetAdvertItems(v []AdvertItem) {
	o.AdvertItems = v
}

func (o AlipayOpenPublicAdvertCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicAdvertCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvertItems) {
		toSerialize["advert_items"] = o.AdvertItems
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicAdvertCreateModel struct {
	value *AlipayOpenPublicAdvertCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicAdvertCreateModel) Get() *AlipayOpenPublicAdvertCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicAdvertCreateModel) Set(val *AlipayOpenPublicAdvertCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAdvertCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAdvertCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAdvertCreateModel(val *AlipayOpenPublicAdvertCreateModel) *NullableAlipayOpenPublicAdvertCreateModel {
	return &NullableAlipayOpenPublicAdvertCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAdvertCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAdvertCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


