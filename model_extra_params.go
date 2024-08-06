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

// checks if the ExtraParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtraParams{}

// ExtraParams struct for ExtraParams
type ExtraParams struct {
	// 轻会员场景下协议结算周期，由协议生效日期到失效日期拼接(yyyyMMddHHmmss-yyyyMMddHHmmss)，精确到秒。
	Period *string `json:"period,omitempty"`
	// 轻会员费用结算原始信息。详细字段说明：total_real_pay_amount轻会员周期内累计支付宝支付金额， total_pay_count轻会员周期内累计支付宝支付次数，total_discount_amount轻会员周期内累计享受的轻会员优惠。（上面金额字段单位为元，精确小数点后两位）
	PeriodSummaryInfo *string `json:"period_summary_info,omitempty"`
	// 用户主动意愿退出：USER_CANCEL_QUIT; 商户结算退出：SETTLE_APPLY_QUIT; 默认值为SETTLE_APPLY_QUIT；这个字段会影响用户在芝麻信用合约的状态
	QuitType *string `json:"quit_type,omitempty"`
	// 代扣期数，周期扣场景PERIOD_SETTLE下需要传递
	WithholdIndex *string `json:"withhold_index,omitempty"`
}

// NewExtraParams instantiates a new ExtraParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraParams() *ExtraParams {
	this := ExtraParams{}
	return &this
}

// NewExtraParamsWithDefaults instantiates a new ExtraParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraParamsWithDefaults() *ExtraParams {
	this := ExtraParams{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *ExtraParams) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraParams) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *ExtraParams) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *ExtraParams) SetPeriod(v string) {
	o.Period = &v
}

// GetPeriodSummaryInfo returns the PeriodSummaryInfo field value if set, zero value otherwise.
func (o *ExtraParams) GetPeriodSummaryInfo() string {
	if o == nil || IsNil(o.PeriodSummaryInfo) {
		var ret string
		return ret
	}
	return *o.PeriodSummaryInfo
}

// GetPeriodSummaryInfoOk returns a tuple with the PeriodSummaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraParams) GetPeriodSummaryInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodSummaryInfo) {
		return nil, false
	}
	return o.PeriodSummaryInfo, true
}

// HasPeriodSummaryInfo returns a boolean if a field has been set.
func (o *ExtraParams) HasPeriodSummaryInfo() bool {
	if o != nil && !IsNil(o.PeriodSummaryInfo) {
		return true
	}

	return false
}

// SetPeriodSummaryInfo gets a reference to the given string and assigns it to the PeriodSummaryInfo field.
func (o *ExtraParams) SetPeriodSummaryInfo(v string) {
	o.PeriodSummaryInfo = &v
}

// GetQuitType returns the QuitType field value if set, zero value otherwise.
func (o *ExtraParams) GetQuitType() string {
	if o == nil || IsNil(o.QuitType) {
		var ret string
		return ret
	}
	return *o.QuitType
}

// GetQuitTypeOk returns a tuple with the QuitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraParams) GetQuitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuitType) {
		return nil, false
	}
	return o.QuitType, true
}

// HasQuitType returns a boolean if a field has been set.
func (o *ExtraParams) HasQuitType() bool {
	if o != nil && !IsNil(o.QuitType) {
		return true
	}

	return false
}

// SetQuitType gets a reference to the given string and assigns it to the QuitType field.
func (o *ExtraParams) SetQuitType(v string) {
	o.QuitType = &v
}

// GetWithholdIndex returns the WithholdIndex field value if set, zero value otherwise.
func (o *ExtraParams) GetWithholdIndex() string {
	if o == nil || IsNil(o.WithholdIndex) {
		var ret string
		return ret
	}
	return *o.WithholdIndex
}

// GetWithholdIndexOk returns a tuple with the WithholdIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraParams) GetWithholdIndexOk() (*string, bool) {
	if o == nil || IsNil(o.WithholdIndex) {
		return nil, false
	}
	return o.WithholdIndex, true
}

// HasWithholdIndex returns a boolean if a field has been set.
func (o *ExtraParams) HasWithholdIndex() bool {
	if o != nil && !IsNil(o.WithholdIndex) {
		return true
	}

	return false
}

// SetWithholdIndex gets a reference to the given string and assigns it to the WithholdIndex field.
func (o *ExtraParams) SetWithholdIndex(v string) {
	o.WithholdIndex = &v
}

func (o ExtraParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtraParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.PeriodSummaryInfo) {
		toSerialize["period_summary_info"] = o.PeriodSummaryInfo
	}
	if !IsNil(o.QuitType) {
		toSerialize["quit_type"] = o.QuitType
	}
	if !IsNil(o.WithholdIndex) {
		toSerialize["withhold_index"] = o.WithholdIndex
	}
	return toSerialize, nil
}

type NullableExtraParams struct {
	value *ExtraParams
	isSet bool
}

func (v NullableExtraParams) Get() *ExtraParams {
	return v.value
}

func (v *NullableExtraParams) Set(val *ExtraParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraParams(val *ExtraParams) *NullableExtraParams {
	return &NullableExtraParams{value: val, isSet: true}
}

func (v NullableExtraParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


