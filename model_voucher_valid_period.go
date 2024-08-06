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

// checks if the VoucherValidPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherValidPeriod{}

// VoucherValidPeriod struct for VoucherValidPeriod
type VoucherValidPeriod struct {
	// 券有效期类型
	Type *string `json:"type,omitempty"`
	// 券可使用的开始时间。格式为：yyyy-MM-dd HH:mm:ss
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// 券生效后N天内可以使用。可以配合wait_days_after_receive字段使用。
	ValidDaysAfterReceive *int32 `json:"valid_days_after_receive,omitempty"`
	// 券可使用的结束时间。格式为yyyy-MM-dd HH:mm:ss
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// 用户领券后需要等待N天，券才可以生效。如果该字段填入0或者不填写，则用户领券后立刻生效。
	WaitDaysAfterReceive *int32 `json:"wait_days_after_receive,omitempty"`
}

// NewVoucherValidPeriod instantiates a new VoucherValidPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherValidPeriod() *VoucherValidPeriod {
	this := VoucherValidPeriod{}
	return &this
}

// NewVoucherValidPeriodWithDefaults instantiates a new VoucherValidPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherValidPeriodWithDefaults() *VoucherValidPeriod {
	this := VoucherValidPeriod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VoucherValidPeriod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherValidPeriod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VoucherValidPeriod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VoucherValidPeriod) SetType(v string) {
	o.Type = &v
}

// GetValidBeginTime returns the ValidBeginTime field value if set, zero value otherwise.
func (o *VoucherValidPeriod) GetValidBeginTime() string {
	if o == nil || IsNil(o.ValidBeginTime) {
		var ret string
		return ret
	}
	return *o.ValidBeginTime
}

// GetValidBeginTimeOk returns a tuple with the ValidBeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherValidPeriod) GetValidBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidBeginTime) {
		return nil, false
	}
	return o.ValidBeginTime, true
}

// HasValidBeginTime returns a boolean if a field has been set.
func (o *VoucherValidPeriod) HasValidBeginTime() bool {
	if o != nil && !IsNil(o.ValidBeginTime) {
		return true
	}

	return false
}

// SetValidBeginTime gets a reference to the given string and assigns it to the ValidBeginTime field.
func (o *VoucherValidPeriod) SetValidBeginTime(v string) {
	o.ValidBeginTime = &v
}

// GetValidDaysAfterReceive returns the ValidDaysAfterReceive field value if set, zero value otherwise.
func (o *VoucherValidPeriod) GetValidDaysAfterReceive() int32 {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		var ret int32
		return ret
	}
	return *o.ValidDaysAfterReceive
}

// GetValidDaysAfterReceiveOk returns a tuple with the ValidDaysAfterReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherValidPeriod) GetValidDaysAfterReceiveOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		return nil, false
	}
	return o.ValidDaysAfterReceive, true
}

// HasValidDaysAfterReceive returns a boolean if a field has been set.
func (o *VoucherValidPeriod) HasValidDaysAfterReceive() bool {
	if o != nil && !IsNil(o.ValidDaysAfterReceive) {
		return true
	}

	return false
}

// SetValidDaysAfterReceive gets a reference to the given int32 and assigns it to the ValidDaysAfterReceive field.
func (o *VoucherValidPeriod) SetValidDaysAfterReceive(v int32) {
	o.ValidDaysAfterReceive = &v
}

// GetValidEndTime returns the ValidEndTime field value if set, zero value otherwise.
func (o *VoucherValidPeriod) GetValidEndTime() string {
	if o == nil || IsNil(o.ValidEndTime) {
		var ret string
		return ret
	}
	return *o.ValidEndTime
}

// GetValidEndTimeOk returns a tuple with the ValidEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherValidPeriod) GetValidEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidEndTime) {
		return nil, false
	}
	return o.ValidEndTime, true
}

// HasValidEndTime returns a boolean if a field has been set.
func (o *VoucherValidPeriod) HasValidEndTime() bool {
	if o != nil && !IsNil(o.ValidEndTime) {
		return true
	}

	return false
}

// SetValidEndTime gets a reference to the given string and assigns it to the ValidEndTime field.
func (o *VoucherValidPeriod) SetValidEndTime(v string) {
	o.ValidEndTime = &v
}

// GetWaitDaysAfterReceive returns the WaitDaysAfterReceive field value if set, zero value otherwise.
func (o *VoucherValidPeriod) GetWaitDaysAfterReceive() int32 {
	if o == nil || IsNil(o.WaitDaysAfterReceive) {
		var ret int32
		return ret
	}
	return *o.WaitDaysAfterReceive
}

// GetWaitDaysAfterReceiveOk returns a tuple with the WaitDaysAfterReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherValidPeriod) GetWaitDaysAfterReceiveOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitDaysAfterReceive) {
		return nil, false
	}
	return o.WaitDaysAfterReceive, true
}

// HasWaitDaysAfterReceive returns a boolean if a field has been set.
func (o *VoucherValidPeriod) HasWaitDaysAfterReceive() bool {
	if o != nil && !IsNil(o.WaitDaysAfterReceive) {
		return true
	}

	return false
}

// SetWaitDaysAfterReceive gets a reference to the given int32 and assigns it to the WaitDaysAfterReceive field.
func (o *VoucherValidPeriod) SetWaitDaysAfterReceive(v int32) {
	o.WaitDaysAfterReceive = &v
}

func (o VoucherValidPeriod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherValidPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ValidBeginTime) {
		toSerialize["valid_begin_time"] = o.ValidBeginTime
	}
	if !IsNil(o.ValidDaysAfterReceive) {
		toSerialize["valid_days_after_receive"] = o.ValidDaysAfterReceive
	}
	if !IsNil(o.ValidEndTime) {
		toSerialize["valid_end_time"] = o.ValidEndTime
	}
	if !IsNil(o.WaitDaysAfterReceive) {
		toSerialize["wait_days_after_receive"] = o.WaitDaysAfterReceive
	}
	return toSerialize, nil
}

type NullableVoucherValidPeriod struct {
	value *VoucherValidPeriod
	isSet bool
}

func (v NullableVoucherValidPeriod) Get() *VoucherValidPeriod {
	return v.value
}

func (v *NullableVoucherValidPeriod) Set(val *VoucherValidPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherValidPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherValidPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherValidPeriod(val *VoucherValidPeriod) *NullableVoucherValidPeriod {
	return &NullableVoucherValidPeriod{value: val, isSet: true}
}

func (v NullableVoucherValidPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherValidPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
