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

// checks if the ZMGOCycleFlexConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGOCycleFlexConfig{}

// ZMGOCycleFlexConfig struct for ZMGOCycleFlexConfig
type ZMGOCycleFlexConfig struct {
	// 周期灵活扣的费用名称
	CycleFlexWithholdFeeName *string `json:"cycle_flex_withhold_fee_name,omitempty"`
	// 周期灵活扣单期扣款的最大额度
	CycleFlexWithholdMaxPrice *string `json:"cycle_flex_withhold_max_price,omitempty"`
	// 周期灵活扣的总期数
	CycleFlexWithholdTotalPeriodCount *int32 `json:"cycle_flex_withhold_total_period_count,omitempty"`
}

// NewZMGOCycleFlexConfig instantiates a new ZMGOCycleFlexConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGOCycleFlexConfig() *ZMGOCycleFlexConfig {
	this := ZMGOCycleFlexConfig{}
	return &this
}

// NewZMGOCycleFlexConfigWithDefaults instantiates a new ZMGOCycleFlexConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGOCycleFlexConfigWithDefaults() *ZMGOCycleFlexConfig {
	this := ZMGOCycleFlexConfig{}
	return &this
}

// GetCycleFlexWithholdFeeName returns the CycleFlexWithholdFeeName field value if set, zero value otherwise.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdFeeName() string {
	if o == nil || IsNil(o.CycleFlexWithholdFeeName) {
		var ret string
		return ret
	}
	return *o.CycleFlexWithholdFeeName
}

// GetCycleFlexWithholdFeeNameOk returns a tuple with the CycleFlexWithholdFeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdFeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.CycleFlexWithholdFeeName) {
		return nil, false
	}
	return o.CycleFlexWithholdFeeName, true
}

// HasCycleFlexWithholdFeeName returns a boolean if a field has been set.
func (o *ZMGOCycleFlexConfig) HasCycleFlexWithholdFeeName() bool {
	if o != nil && !IsNil(o.CycleFlexWithholdFeeName) {
		return true
	}

	return false
}

// SetCycleFlexWithholdFeeName gets a reference to the given string and assigns it to the CycleFlexWithholdFeeName field.
func (o *ZMGOCycleFlexConfig) SetCycleFlexWithholdFeeName(v string) {
	o.CycleFlexWithholdFeeName = &v
}

// GetCycleFlexWithholdMaxPrice returns the CycleFlexWithholdMaxPrice field value if set, zero value otherwise.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdMaxPrice() string {
	if o == nil || IsNil(o.CycleFlexWithholdMaxPrice) {
		var ret string
		return ret
	}
	return *o.CycleFlexWithholdMaxPrice
}

// GetCycleFlexWithholdMaxPriceOk returns a tuple with the CycleFlexWithholdMaxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdMaxPriceOk() (*string, bool) {
	if o == nil || IsNil(o.CycleFlexWithholdMaxPrice) {
		return nil, false
	}
	return o.CycleFlexWithholdMaxPrice, true
}

// HasCycleFlexWithholdMaxPrice returns a boolean if a field has been set.
func (o *ZMGOCycleFlexConfig) HasCycleFlexWithholdMaxPrice() bool {
	if o != nil && !IsNil(o.CycleFlexWithholdMaxPrice) {
		return true
	}

	return false
}

// SetCycleFlexWithholdMaxPrice gets a reference to the given string and assigns it to the CycleFlexWithholdMaxPrice field.
func (o *ZMGOCycleFlexConfig) SetCycleFlexWithholdMaxPrice(v string) {
	o.CycleFlexWithholdMaxPrice = &v
}

// GetCycleFlexWithholdTotalPeriodCount returns the CycleFlexWithholdTotalPeriodCount field value if set, zero value otherwise.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdTotalPeriodCount() int32 {
	if o == nil || IsNil(o.CycleFlexWithholdTotalPeriodCount) {
		var ret int32
		return ret
	}
	return *o.CycleFlexWithholdTotalPeriodCount
}

// GetCycleFlexWithholdTotalPeriodCountOk returns a tuple with the CycleFlexWithholdTotalPeriodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleFlexConfig) GetCycleFlexWithholdTotalPeriodCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CycleFlexWithholdTotalPeriodCount) {
		return nil, false
	}
	return o.CycleFlexWithholdTotalPeriodCount, true
}

// HasCycleFlexWithholdTotalPeriodCount returns a boolean if a field has been set.
func (o *ZMGOCycleFlexConfig) HasCycleFlexWithholdTotalPeriodCount() bool {
	if o != nil && !IsNil(o.CycleFlexWithholdTotalPeriodCount) {
		return true
	}

	return false
}

// SetCycleFlexWithholdTotalPeriodCount gets a reference to the given int32 and assigns it to the CycleFlexWithholdTotalPeriodCount field.
func (o *ZMGOCycleFlexConfig) SetCycleFlexWithholdTotalPeriodCount(v int32) {
	o.CycleFlexWithholdTotalPeriodCount = &v
}

func (o ZMGOCycleFlexConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGOCycleFlexConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CycleFlexWithholdFeeName) {
		toSerialize["cycle_flex_withhold_fee_name"] = o.CycleFlexWithholdFeeName
	}
	if !IsNil(o.CycleFlexWithholdMaxPrice) {
		toSerialize["cycle_flex_withhold_max_price"] = o.CycleFlexWithholdMaxPrice
	}
	if !IsNil(o.CycleFlexWithholdTotalPeriodCount) {
		toSerialize["cycle_flex_withhold_total_period_count"] = o.CycleFlexWithholdTotalPeriodCount
	}
	return toSerialize, nil
}

type NullableZMGOCycleFlexConfig struct {
	value *ZMGOCycleFlexConfig
	isSet bool
}

func (v NullableZMGOCycleFlexConfig) Get() *ZMGOCycleFlexConfig {
	return v.value
}

func (v *NullableZMGOCycleFlexConfig) Set(val *ZMGOCycleFlexConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGOCycleFlexConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGOCycleFlexConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGOCycleFlexConfig(val *ZMGOCycleFlexConfig) *NullableZMGOCycleFlexConfig {
	return &NullableZMGOCycleFlexConfig{value: val, isSet: true}
}

func (v NullableZMGOCycleFlexConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGOCycleFlexConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


