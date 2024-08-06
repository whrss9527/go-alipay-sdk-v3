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

// checks if the ShopBusinessTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopBusinessTime{}

// ShopBusinessTime struct for ShopBusinessTime
type ShopBusinessTime struct {
	// 关门时间 格式：HH:mm
	CloseTime *string `json:"close_time,omitempty"`
	// 开门时间 格式：HH:mm
	OpenTime *string `json:"open_time,omitempty"`
	// 本对象表示周几的营业时间。1~6表示周一到周六，7表示周日
	WeekDay *int32 `json:"week_day,omitempty"`
}

// NewShopBusinessTime instantiates a new ShopBusinessTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopBusinessTime() *ShopBusinessTime {
	this := ShopBusinessTime{}
	return &this
}

// NewShopBusinessTimeWithDefaults instantiates a new ShopBusinessTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopBusinessTimeWithDefaults() *ShopBusinessTime {
	this := ShopBusinessTime{}
	return &this
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *ShopBusinessTime) GetCloseTime() string {
	if o == nil || IsNil(o.CloseTime) {
		var ret string
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopBusinessTime) GetCloseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *ShopBusinessTime) HasCloseTime() bool {
	if o != nil && !IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given string and assigns it to the CloseTime field.
func (o *ShopBusinessTime) SetCloseTime(v string) {
	o.CloseTime = &v
}

// GetOpenTime returns the OpenTime field value if set, zero value otherwise.
func (o *ShopBusinessTime) GetOpenTime() string {
	if o == nil || IsNil(o.OpenTime) {
		var ret string
		return ret
	}
	return *o.OpenTime
}

// GetOpenTimeOk returns a tuple with the OpenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopBusinessTime) GetOpenTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OpenTime) {
		return nil, false
	}
	return o.OpenTime, true
}

// HasOpenTime returns a boolean if a field has been set.
func (o *ShopBusinessTime) HasOpenTime() bool {
	if o != nil && !IsNil(o.OpenTime) {
		return true
	}

	return false
}

// SetOpenTime gets a reference to the given string and assigns it to the OpenTime field.
func (o *ShopBusinessTime) SetOpenTime(v string) {
	o.OpenTime = &v
}

// GetWeekDay returns the WeekDay field value if set, zero value otherwise.
func (o *ShopBusinessTime) GetWeekDay() int32 {
	if o == nil || IsNil(o.WeekDay) {
		var ret int32
		return ret
	}
	return *o.WeekDay
}

// GetWeekDayOk returns a tuple with the WeekDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopBusinessTime) GetWeekDayOk() (*int32, bool) {
	if o == nil || IsNil(o.WeekDay) {
		return nil, false
	}
	return o.WeekDay, true
}

// HasWeekDay returns a boolean if a field has been set.
func (o *ShopBusinessTime) HasWeekDay() bool {
	if o != nil && !IsNil(o.WeekDay) {
		return true
	}

	return false
}

// SetWeekDay gets a reference to the given int32 and assigns it to the WeekDay field.
func (o *ShopBusinessTime) SetWeekDay(v int32) {
	o.WeekDay = &v
}

func (o ShopBusinessTime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopBusinessTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloseTime) {
		toSerialize["close_time"] = o.CloseTime
	}
	if !IsNil(o.OpenTime) {
		toSerialize["open_time"] = o.OpenTime
	}
	if !IsNil(o.WeekDay) {
		toSerialize["week_day"] = o.WeekDay
	}
	return toSerialize, nil
}

type NullableShopBusinessTime struct {
	value *ShopBusinessTime
	isSet bool
}

func (v NullableShopBusinessTime) Get() *ShopBusinessTime {
	return v.value
}

func (v *NullableShopBusinessTime) Set(val *ShopBusinessTime) {
	v.value = val
	v.isSet = true
}

func (v NullableShopBusinessTime) IsSet() bool {
	return v.isSet
}

func (v *NullableShopBusinessTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopBusinessTime(val *ShopBusinessTime) *NullableShopBusinessTime {
	return &NullableShopBusinessTime{value: val, isSet: true}
}

func (v NullableShopBusinessTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopBusinessTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
