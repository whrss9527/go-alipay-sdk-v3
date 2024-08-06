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

// checks if the PaymentVoucherValidPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherValidPeriod{}

// PaymentVoucherValidPeriod struct for PaymentVoucherValidPeriod
type PaymentVoucherValidPeriod struct {
	// 券有效期类型
	Type *string `json:"type,omitempty"`
	// 券可使用的开始时间。  格式为：yyyy-MM-dd HH:mm:ss  限制： type为ABSOLUTE时该字段必填。
	ValidBeginTime *string `json:"valid_begin_time,omitempty"`
	// 券生效后N天内可以使用。  可以配合wait_days_after_receive字段使用。  限制： 1.type为RELATIVE时必填。 2.valid_days_after_receive必须大于0。
	ValidDaysAfterReceive *int32 `json:"valid_days_after_receive,omitempty"`
	// 券可使用的结束时间。  格式为yyyy-MM-dd HH:mm:ss  限制： 1、type为ABSOLUTE必填。 2、券可使用的结束时间valid_end_time必须大于券的发放结束时间publish_end_time
	ValidEndTime *string `json:"valid_end_time,omitempty"`
	// 用户领券后需要等待N天，券才可以生效。默认用户领券后立刻生效。  限制： 1、type为RELATIVE时可选。
	WaitDaysAfterReceive *int32 `json:"wait_days_after_receive,omitempty"`
}

// NewPaymentVoucherValidPeriod instantiates a new PaymentVoucherValidPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherValidPeriod() *PaymentVoucherValidPeriod {
	this := PaymentVoucherValidPeriod{}
	return &this
}

// NewPaymentVoucherValidPeriodWithDefaults instantiates a new PaymentVoucherValidPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherValidPeriodWithDefaults() *PaymentVoucherValidPeriod {
	this := PaymentVoucherValidPeriod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriod) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriod) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriod) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentVoucherValidPeriod) SetType(v string) {
	o.Type = &v
}

// GetValidBeginTime returns the ValidBeginTime field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriod) GetValidBeginTime() string {
	if o == nil || IsNil(o.ValidBeginTime) {
		var ret string
		return ret
	}
	return *o.ValidBeginTime
}

// GetValidBeginTimeOk returns a tuple with the ValidBeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriod) GetValidBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidBeginTime) {
		return nil, false
	}
	return o.ValidBeginTime, true
}

// HasValidBeginTime returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriod) HasValidBeginTime() bool {
	if o != nil && !IsNil(o.ValidBeginTime) {
		return true
	}

	return false
}

// SetValidBeginTime gets a reference to the given string and assigns it to the ValidBeginTime field.
func (o *PaymentVoucherValidPeriod) SetValidBeginTime(v string) {
	o.ValidBeginTime = &v
}

// GetValidDaysAfterReceive returns the ValidDaysAfterReceive field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriod) GetValidDaysAfterReceive() int32 {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		var ret int32
		return ret
	}
	return *o.ValidDaysAfterReceive
}

// GetValidDaysAfterReceiveOk returns a tuple with the ValidDaysAfterReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriod) GetValidDaysAfterReceiveOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidDaysAfterReceive) {
		return nil, false
	}
	return o.ValidDaysAfterReceive, true
}

// HasValidDaysAfterReceive returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriod) HasValidDaysAfterReceive() bool {
	if o != nil && !IsNil(o.ValidDaysAfterReceive) {
		return true
	}

	return false
}

// SetValidDaysAfterReceive gets a reference to the given int32 and assigns it to the ValidDaysAfterReceive field.
func (o *PaymentVoucherValidPeriod) SetValidDaysAfterReceive(v int32) {
	o.ValidDaysAfterReceive = &v
}

// GetValidEndTime returns the ValidEndTime field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriod) GetValidEndTime() string {
	if o == nil || IsNil(o.ValidEndTime) {
		var ret string
		return ret
	}
	return *o.ValidEndTime
}

// GetValidEndTimeOk returns a tuple with the ValidEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriod) GetValidEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ValidEndTime) {
		return nil, false
	}
	return o.ValidEndTime, true
}

// HasValidEndTime returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriod) HasValidEndTime() bool {
	if o != nil && !IsNil(o.ValidEndTime) {
		return true
	}

	return false
}

// SetValidEndTime gets a reference to the given string and assigns it to the ValidEndTime field.
func (o *PaymentVoucherValidPeriod) SetValidEndTime(v string) {
	o.ValidEndTime = &v
}

// GetWaitDaysAfterReceive returns the WaitDaysAfterReceive field value if set, zero value otherwise.
func (o *PaymentVoucherValidPeriod) GetWaitDaysAfterReceive() int32 {
	if o == nil || IsNil(o.WaitDaysAfterReceive) {
		var ret int32
		return ret
	}
	return *o.WaitDaysAfterReceive
}

// GetWaitDaysAfterReceiveOk returns a tuple with the WaitDaysAfterReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherValidPeriod) GetWaitDaysAfterReceiveOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitDaysAfterReceive) {
		return nil, false
	}
	return o.WaitDaysAfterReceive, true
}

// HasWaitDaysAfterReceive returns a boolean if a field has been set.
func (o *PaymentVoucherValidPeriod) HasWaitDaysAfterReceive() bool {
	if o != nil && !IsNil(o.WaitDaysAfterReceive) {
		return true
	}

	return false
}

// SetWaitDaysAfterReceive gets a reference to the given int32 and assigns it to the WaitDaysAfterReceive field.
func (o *PaymentVoucherValidPeriod) SetWaitDaysAfterReceive(v int32) {
	o.WaitDaysAfterReceive = &v
}

func (o PaymentVoucherValidPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherValidPeriod) ToMap() (map[string]interface{}, error) {
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

type NullablePaymentVoucherValidPeriod struct {
	value *PaymentVoucherValidPeriod
	isSet bool
}

func (v NullablePaymentVoucherValidPeriod) Get() *PaymentVoucherValidPeriod {
	return v.value
}

func (v *NullablePaymentVoucherValidPeriod) Set(val *PaymentVoucherValidPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherValidPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherValidPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherValidPeriod(val *PaymentVoucherValidPeriod) *NullablePaymentVoucherValidPeriod {
	return &NullablePaymentVoucherValidPeriod{value: val, isSet: true}
}

func (v NullablePaymentVoucherValidPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherValidPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


