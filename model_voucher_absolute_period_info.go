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

// checks if the VoucherAbsolutePeriodInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAbsolutePeriodInfo{}

// VoucherAbsolutePeriodInfo struct for VoucherAbsolutePeriodInfo
type VoucherAbsolutePeriodInfo struct {
	TimeRestrictInfo *TimeRestrictInfo `json:"time_restrict_info,omitempty"`
	// 券可使用的开始时间。 格式为：yyyy-MM-dd HH:mm:ss。
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// 券可使用的结束时间。 格式为yyyy-MM-dd HH:mm:ss。
	ValidEndTime *string `json:"valid_end_time,omitempty"`
}

// NewVoucherAbsolutePeriodInfo instantiates a new VoucherAbsolutePeriodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAbsolutePeriodInfo() *VoucherAbsolutePeriodInfo {
	this := VoucherAbsolutePeriodInfo{}
	return &this
}

// NewVoucherAbsolutePeriodInfoWithDefaults instantiates a new VoucherAbsolutePeriodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAbsolutePeriodInfoWithDefaults() *VoucherAbsolutePeriodInfo {
	this := VoucherAbsolutePeriodInfo{}
	return &this
}

// GetTimeRestrictInfo returns the TimeRestrictInfo field value if set, zero value otherwise.
func (o *VoucherAbsolutePeriodInfo) GetTimeRestrictInfo() TimeRestrictInfo {
	if o == nil || IsNil(o.TimeRestrictInfo) {
		var ret TimeRestrictInfo
		return ret
	}
	return *o.TimeRestrictInfo
}

// GetTimeRestrictInfoOk returns a tuple with the TimeRestrictInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAbsolutePeriodInfo) GetTimeRestrictInfoOk() (*TimeRestrictInfo, bool) {
	if o == nil || IsNil(o.TimeRestrictInfo) {
		return nil, false
	}
	return o.TimeRestrictInfo, true
}

// HasTimeRestrictInfo returns a boolean if a field has been set.
func (o *VoucherAbsolutePeriodInfo) HasTimeRestrictInfo() bool {
	if o != nil && !IsNil(o.TimeRestrictInfo) {
		return true
	}

	return false
}

// SetTimeRestrictInfo gets a reference to the given TimeRestrictInfo and assigns it to the TimeRestrictInfo field.
func (o *VoucherAbsolutePeriodInfo) SetTimeRestrictInfo(v TimeRestrictInfo) {
	o.TimeRestrictInfo = &v
}

// GetValidBeginTime returns the ValidBeginTime field value if set, zero value otherwise.
func (o *VoucherAbsolutePeriodInfo) GetValidBeginTime() string {
	if o == nil || IsNil(o.ValidBeginTime) {
		var ret string
		return ret
	}
	return *o.ValidBeginTime
}

// GetValidBeginTimeOk returns a tuple with the ValidBeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAbsolutePeriodInfo) GetValidBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidBeginTime) {
		return nil, false
	}
	return o.ValidBeginTime, true
}

// HasValidBeginTime returns a boolean if a field has been set.
func (o *VoucherAbsolutePeriodInfo) HasValidBeginTime() bool {
	if o != nil && !IsNil(o.ValidBeginTime) {
		return true
	}

	return false
}

// SetValidBeginTime gets a reference to the given string and assigns it to the ValidBeginTime field.
func (o *VoucherAbsolutePeriodInfo) SetValidBeginTime(v string) {
	o.ValidBeginTime = &v
}

// GetValidEndTime returns the ValidEndTime field value if set, zero value otherwise.
func (o *VoucherAbsolutePeriodInfo) GetValidEndTime() string {
	if o == nil || IsNil(o.ValidEndTime) {
		var ret string
		return ret
	}
	return *o.ValidEndTime
}

// GetValidEndTimeOk returns a tuple with the ValidEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAbsolutePeriodInfo) GetValidEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidEndTime) {
		return nil, false
	}
	return o.ValidEndTime, true
}

// HasValidEndTime returns a boolean if a field has been set.
func (o *VoucherAbsolutePeriodInfo) HasValidEndTime() bool {
	if o != nil && !IsNil(o.ValidEndTime) {
		return true
	}

	return false
}

// SetValidEndTime gets a reference to the given string and assigns it to the ValidEndTime field.
func (o *VoucherAbsolutePeriodInfo) SetValidEndTime(v string) {
	o.ValidEndTime = &v
}

func (o VoucherAbsolutePeriodInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAbsolutePeriodInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TimeRestrictInfo) {
		toSerialize["time_restrict_info"] = o.TimeRestrictInfo
	}
	if !IsNil(o.ValidBeginTime) {
		toSerialize["valid_begin_time"] = o.ValidBeginTime
	}
	if !IsNil(o.ValidEndTime) {
		toSerialize["valid_end_time"] = o.ValidEndTime
	}
	return toSerialize, nil
}

type NullableVoucherAbsolutePeriodInfo struct {
	value *VoucherAbsolutePeriodInfo
	isSet bool
}

func (v NullableVoucherAbsolutePeriodInfo) Get() *VoucherAbsolutePeriodInfo {
	return v.value
}

func (v *NullableVoucherAbsolutePeriodInfo) Set(val *VoucherAbsolutePeriodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAbsolutePeriodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAbsolutePeriodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAbsolutePeriodInfo(val *VoucherAbsolutePeriodInfo) *NullableVoucherAbsolutePeriodInfo {
	return &NullableVoucherAbsolutePeriodInfo{value: val, isSet: true}
}

func (v NullableVoucherAbsolutePeriodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAbsolutePeriodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
