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

// checks if the ZMGOCycleWithholdConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGOCycleWithholdConfig{}

// ZMGOCycleWithholdConfig struct for ZMGOCycleWithholdConfig
type ZMGOCycleWithholdConfig struct {
	// 扣款计划
	DeductPlan []string `json:"deduct_plan,omitempty"`
	// 豁免期天数
	ExemptionPeriod *string `json:"exemption_period,omitempty"`
	// 周期
	Period *string `json:"period,omitempty"`
	// 扣款周期类型，支持DAY-天模式，MONTH-月模式
	PeriodType *string `json:"period_type,omitempty"`
	// 是否支持周期扣高级模式，高级模式下扣款计划中内一起的扣款金额可以进行定制
	SupportCycleWithholdHighMode *bool `json:"support_cycle_withhold_high_mode,omitempty"`
	// 是否支持豁免期，如果支持豁免期，用户开通后再豁免期内退出无需扣款
	SupportExemptionPeriod *bool `json:"support_exemption_period,omitempty"`
	// 设置扣款模式，目前支持每月固定日期或固定时间间隔
	WithholdMode *string `json:"withhold_mode,omitempty"`
}

// NewZMGOCycleWithholdConfig instantiates a new ZMGOCycleWithholdConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGOCycleWithholdConfig() *ZMGOCycleWithholdConfig {
	this := ZMGOCycleWithholdConfig{}
	return &this
}

// NewZMGOCycleWithholdConfigWithDefaults instantiates a new ZMGOCycleWithholdConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGOCycleWithholdConfigWithDefaults() *ZMGOCycleWithholdConfig {
	this := ZMGOCycleWithholdConfig{}
	return &this
}

// GetDeductPlan returns the DeductPlan field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetDeductPlan() []string {
	if o == nil || IsNil(o.DeductPlan) {
		var ret []string
		return ret
	}
	return o.DeductPlan
}

// GetDeductPlanOk returns a tuple with the DeductPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetDeductPlanOk() ([]string, bool) {
	if o == nil || IsNil(o.DeductPlan) {
		return nil, false
	}
	return o.DeductPlan, true
}

// HasDeductPlan returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasDeductPlan() bool {
	if o != nil && !IsNil(o.DeductPlan) {
		return true
	}

	return false
}

// SetDeductPlan gets a reference to the given []string and assigns it to the DeductPlan field.
func (o *ZMGOCycleWithholdConfig) SetDeductPlan(v []string) {
	o.DeductPlan = v
}

// GetExemptionPeriod returns the ExemptionPeriod field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetExemptionPeriod() string {
	if o == nil || IsNil(o.ExemptionPeriod) {
		var ret string
		return ret
	}
	return *o.ExemptionPeriod
}

// GetExemptionPeriodOk returns a tuple with the ExemptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetExemptionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ExemptionPeriod) {
		return nil, false
	}
	return o.ExemptionPeriod, true
}

// HasExemptionPeriod returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasExemptionPeriod() bool {
	if o != nil && !IsNil(o.ExemptionPeriod) {
		return true
	}

	return false
}

// SetExemptionPeriod gets a reference to the given string and assigns it to the ExemptionPeriod field.
func (o *ZMGOCycleWithholdConfig) SetExemptionPeriod(v string) {
	o.ExemptionPeriod = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *ZMGOCycleWithholdConfig) SetPeriod(v string) {
	o.Period = &v
}

// GetPeriodType returns the PeriodType field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetPeriodType() string {
	if o == nil || IsNil(o.PeriodType) {
		var ret string
		return ret
	}
	return *o.PeriodType
}

// GetPeriodTypeOk returns a tuple with the PeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetPeriodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodType) {
		return nil, false
	}
	return o.PeriodType, true
}

// HasPeriodType returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasPeriodType() bool {
	if o != nil && !IsNil(o.PeriodType) {
		return true
	}

	return false
}

// SetPeriodType gets a reference to the given string and assigns it to the PeriodType field.
func (o *ZMGOCycleWithholdConfig) SetPeriodType(v string) {
	o.PeriodType = &v
}

// GetSupportCycleWithholdHighMode returns the SupportCycleWithholdHighMode field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetSupportCycleWithholdHighMode() bool {
	if o == nil || IsNil(o.SupportCycleWithholdHighMode) {
		var ret bool
		return ret
	}
	return *o.SupportCycleWithholdHighMode
}

// GetSupportCycleWithholdHighModeOk returns a tuple with the SupportCycleWithholdHighMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetSupportCycleWithholdHighModeOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportCycleWithholdHighMode) {
		return nil, false
	}
	return o.SupportCycleWithholdHighMode, true
}

// HasSupportCycleWithholdHighMode returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasSupportCycleWithholdHighMode() bool {
	if o != nil && !IsNil(o.SupportCycleWithholdHighMode) {
		return true
	}

	return false
}

// SetSupportCycleWithholdHighMode gets a reference to the given bool and assigns it to the SupportCycleWithholdHighMode field.
func (o *ZMGOCycleWithholdConfig) SetSupportCycleWithholdHighMode(v bool) {
	o.SupportCycleWithholdHighMode = &v
}

// GetSupportExemptionPeriod returns the SupportExemptionPeriod field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetSupportExemptionPeriod() bool {
	if o == nil || IsNil(o.SupportExemptionPeriod) {
		var ret bool
		return ret
	}
	return *o.SupportExemptionPeriod
}

// GetSupportExemptionPeriodOk returns a tuple with the SupportExemptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetSupportExemptionPeriodOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportExemptionPeriod) {
		return nil, false
	}
	return o.SupportExemptionPeriod, true
}

// HasSupportExemptionPeriod returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasSupportExemptionPeriod() bool {
	if o != nil && !IsNil(o.SupportExemptionPeriod) {
		return true
	}

	return false
}

// SetSupportExemptionPeriod gets a reference to the given bool and assigns it to the SupportExemptionPeriod field.
func (o *ZMGOCycleWithholdConfig) SetSupportExemptionPeriod(v bool) {
	o.SupportExemptionPeriod = &v
}

// GetWithholdMode returns the WithholdMode field value if set, zero value otherwise.
func (o *ZMGOCycleWithholdConfig) GetWithholdMode() string {
	if o == nil || IsNil(o.WithholdMode) {
		var ret string
		return ret
	}
	return *o.WithholdMode
}

// GetWithholdModeOk returns a tuple with the WithholdMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGOCycleWithholdConfig) GetWithholdModeOk() (*string, bool) {
	if o == nil || IsNil(o.WithholdMode) {
		return nil, false
	}
	return o.WithholdMode, true
}

// HasWithholdMode returns a boolean if a field has been set.
func (o *ZMGOCycleWithholdConfig) HasWithholdMode() bool {
	if o != nil && !IsNil(o.WithholdMode) {
		return true
	}

	return false
}

// SetWithholdMode gets a reference to the given string and assigns it to the WithholdMode field.
func (o *ZMGOCycleWithholdConfig) SetWithholdMode(v string) {
	o.WithholdMode = &v
}

func (o ZMGOCycleWithholdConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGOCycleWithholdConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeductPlan) {
		toSerialize["deduct_plan"] = o.DeductPlan
	}
	if !IsNil(o.ExemptionPeriod) {
		toSerialize["exemption_period"] = o.ExemptionPeriod
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.PeriodType) {
		toSerialize["period_type"] = o.PeriodType
	}
	if !IsNil(o.SupportCycleWithholdHighMode) {
		toSerialize["support_cycle_withhold_high_mode"] = o.SupportCycleWithholdHighMode
	}
	if !IsNil(o.SupportExemptionPeriod) {
		toSerialize["support_exemption_period"] = o.SupportExemptionPeriod
	}
	if !IsNil(o.WithholdMode) {
		toSerialize["withhold_mode"] = o.WithholdMode
	}
	return toSerialize, nil
}

type NullableZMGOCycleWithholdConfig struct {
	value *ZMGOCycleWithholdConfig
	isSet bool
}

func (v NullableZMGOCycleWithholdConfig) Get() *ZMGOCycleWithholdConfig {
	return v.value
}

func (v *NullableZMGOCycleWithholdConfig) Set(val *ZMGOCycleWithholdConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGOCycleWithholdConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGOCycleWithholdConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGOCycleWithholdConfig(val *ZMGOCycleWithholdConfig) *NullableZMGOCycleWithholdConfig {
	return &NullableZMGOCycleWithholdConfig{value: val, isSet: true}
}

func (v NullableZMGOCycleWithholdConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGOCycleWithholdConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


