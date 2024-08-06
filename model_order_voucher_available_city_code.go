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

// checks if the OrderVoucherAvailableCityCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVoucherAvailableCityCode{}

// OrderVoucherAvailableCityCode struct for OrderVoucherAvailableCityCode
type OrderVoucherAvailableCityCode struct {
	// 是否全国。选择全国后，无须填写city_codes字段。系统默认填充全国全部城市信息。
	AllCity *bool `json:"all_city,omitempty"`
	// 城市编码。请按照<a href=\"https://gw.alipayobjects.com/os/basement_prod/253c4dcb-b8a4-4a1e-8be2-79e191a9b6db.xlsx\" target=\"_blank\">表格</a>中内容填写。 （<a href=\"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/\" target=\"_blank\">参考资料</a>）
	CityCodes []string `json:"city_codes,omitempty"`
}

// NewOrderVoucherAvailableCityCode instantiates a new OrderVoucherAvailableCityCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVoucherAvailableCityCode() *OrderVoucherAvailableCityCode {
	this := OrderVoucherAvailableCityCode{}
	return &this
}

// NewOrderVoucherAvailableCityCodeWithDefaults instantiates a new OrderVoucherAvailableCityCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVoucherAvailableCityCodeWithDefaults() *OrderVoucherAvailableCityCode {
	this := OrderVoucherAvailableCityCode{}
	return &this
}

// GetAllCity returns the AllCity field value if set, zero value otherwise.
func (o *OrderVoucherAvailableCityCode) GetAllCity() bool {
	if o == nil || IsNil(o.AllCity) {
		var ret bool
		return ret
	}
	return *o.AllCity
}

// GetAllCityOk returns a tuple with the AllCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableCityCode) GetAllCityOk() (*bool, bool) {
	if o == nil || IsNil(o.AllCity) {
		return nil, false
	}
	return o.AllCity, true
}

// HasAllCity returns a boolean if a field has been set.
func (o *OrderVoucherAvailableCityCode) HasAllCity() bool {
	if o != nil && !IsNil(o.AllCity) {
		return true
	}

	return false
}

// SetAllCity gets a reference to the given bool and assigns it to the AllCity field.
func (o *OrderVoucherAvailableCityCode) SetAllCity(v bool) {
	o.AllCity = &v
}

// GetCityCodes returns the CityCodes field value if set, zero value otherwise.
func (o *OrderVoucherAvailableCityCode) GetCityCodes() []string {
	if o == nil || IsNil(o.CityCodes) {
		var ret []string
		return ret
	}
	return o.CityCodes
}

// GetCityCodesOk returns a tuple with the CityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableCityCode) GetCityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.CityCodes) {
		return nil, false
	}
	return o.CityCodes, true
}

// HasCityCodes returns a boolean if a field has been set.
func (o *OrderVoucherAvailableCityCode) HasCityCodes() bool {
	if o != nil && !IsNil(o.CityCodes) {
		return true
	}

	return false
}

// SetCityCodes gets a reference to the given []string and assigns it to the CityCodes field.
func (o *OrderVoucherAvailableCityCode) SetCityCodes(v []string) {
	o.CityCodes = v
}

func (o OrderVoucherAvailableCityCode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVoucherAvailableCityCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllCity) {
		toSerialize["all_city"] = o.AllCity
	}
	if !IsNil(o.CityCodes) {
		toSerialize["city_codes"] = o.CityCodes
	}
	return toSerialize, nil
}

type NullableOrderVoucherAvailableCityCode struct {
	value *OrderVoucherAvailableCityCode
	isSet bool
}

func (v NullableOrderVoucherAvailableCityCode) Get() *OrderVoucherAvailableCityCode {
	return v.value
}

func (v *NullableOrderVoucherAvailableCityCode) Set(val *OrderVoucherAvailableCityCode) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVoucherAvailableCityCode) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVoucherAvailableCityCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVoucherAvailableCityCode(val *OrderVoucherAvailableCityCode) *NullableOrderVoucherAvailableCityCode {
	return &NullableOrderVoucherAvailableCityCode{value: val, isSet: true}
}

func (v NullableOrderVoucherAvailableCityCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVoucherAvailableCityCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
