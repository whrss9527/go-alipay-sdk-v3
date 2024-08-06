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

// checks if the LogisticInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogisticInfo{}

// LogisticInfo struct for LogisticInfo
type LogisticInfo struct {
	// 物流公司名称
	Channel *string `json:"channel,omitempty"`
	// 物流详情
	Detail *string `json:"detail,omitempty"`
	// 物流id
	LogisticId *string `json:"logistic_id,omitempty"`
	// 发货地
	ShipArea *string `json:"ship_area,omitempty"`
	// 发货时效
	ShipPeriod *string `json:"ship_period,omitempty"`
	// 物流状态
	Status *string `json:"status,omitempty"`
	// 物流停更时间
	StopUpdateTime *string `json:"stop_update_time,omitempty"`
}

// NewLogisticInfo instantiates a new LogisticInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogisticInfo() *LogisticInfo {
	this := LogisticInfo{}
	return &this
}

// NewLogisticInfoWithDefaults instantiates a new LogisticInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogisticInfoWithDefaults() *LogisticInfo {
	this := LogisticInfo{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *LogisticInfo) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *LogisticInfo) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *LogisticInfo) SetChannel(v string) {
	o.Channel = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *LogisticInfo) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *LogisticInfo) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *LogisticInfo) SetDetail(v string) {
	o.Detail = &v
}

// GetLogisticId returns the LogisticId field value if set, zero value otherwise.
func (o *LogisticInfo) GetLogisticId() string {
	if o == nil || IsNil(o.LogisticId) {
		var ret string
		return ret
	}
	return *o.LogisticId
}

// GetLogisticIdOk returns a tuple with the LogisticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetLogisticIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogisticId) {
		return nil, false
	}
	return o.LogisticId, true
}

// HasLogisticId returns a boolean if a field has been set.
func (o *LogisticInfo) HasLogisticId() bool {
	if o != nil && !IsNil(o.LogisticId) {
		return true
	}

	return false
}

// SetLogisticId gets a reference to the given string and assigns it to the LogisticId field.
func (o *LogisticInfo) SetLogisticId(v string) {
	o.LogisticId = &v
}

// GetShipArea returns the ShipArea field value if set, zero value otherwise.
func (o *LogisticInfo) GetShipArea() string {
	if o == nil || IsNil(o.ShipArea) {
		var ret string
		return ret
	}
	return *o.ShipArea
}

// GetShipAreaOk returns a tuple with the ShipArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetShipAreaOk() (*string, bool) {
	if o == nil || IsNil(o.ShipArea) {
		return nil, false
	}
	return o.ShipArea, true
}

// HasShipArea returns a boolean if a field has been set.
func (o *LogisticInfo) HasShipArea() bool {
	if o != nil && !IsNil(o.ShipArea) {
		return true
	}

	return false
}

// SetShipArea gets a reference to the given string and assigns it to the ShipArea field.
func (o *LogisticInfo) SetShipArea(v string) {
	o.ShipArea = &v
}

// GetShipPeriod returns the ShipPeriod field value if set, zero value otherwise.
func (o *LogisticInfo) GetShipPeriod() string {
	if o == nil || IsNil(o.ShipPeriod) {
		var ret string
		return ret
	}
	return *o.ShipPeriod
}

// GetShipPeriodOk returns a tuple with the ShipPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetShipPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ShipPeriod) {
		return nil, false
	}
	return o.ShipPeriod, true
}

// HasShipPeriod returns a boolean if a field has been set.
func (o *LogisticInfo) HasShipPeriod() bool {
	if o != nil && !IsNil(o.ShipPeriod) {
		return true
	}

	return false
}

// SetShipPeriod gets a reference to the given string and assigns it to the ShipPeriod field.
func (o *LogisticInfo) SetShipPeriod(v string) {
	o.ShipPeriod = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogisticInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogisticInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LogisticInfo) SetStatus(v string) {
	o.Status = &v
}

// GetStopUpdateTime returns the StopUpdateTime field value if set, zero value otherwise.
func (o *LogisticInfo) GetStopUpdateTime() string {
	if o == nil || IsNil(o.StopUpdateTime) {
		var ret string
		return ret
	}
	return *o.StopUpdateTime
}

// GetStopUpdateTimeOk returns a tuple with the StopUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogisticInfo) GetStopUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StopUpdateTime) {
		return nil, false
	}
	return o.StopUpdateTime, true
}

// HasStopUpdateTime returns a boolean if a field has been set.
func (o *LogisticInfo) HasStopUpdateTime() bool {
	if o != nil && !IsNil(o.StopUpdateTime) {
		return true
	}

	return false
}

// SetStopUpdateTime gets a reference to the given string and assigns it to the StopUpdateTime field.
func (o *LogisticInfo) SetStopUpdateTime(v string) {
	o.StopUpdateTime = &v
}

func (o LogisticInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogisticInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.LogisticId) {
		toSerialize["logistic_id"] = o.LogisticId
	}
	if !IsNil(o.ShipArea) {
		toSerialize["ship_area"] = o.ShipArea
	}
	if !IsNil(o.ShipPeriod) {
		toSerialize["ship_period"] = o.ShipPeriod
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StopUpdateTime) {
		toSerialize["stop_update_time"] = o.StopUpdateTime
	}
	return toSerialize, nil
}

type NullableLogisticInfo struct {
	value *LogisticInfo
	isSet bool
}

func (v NullableLogisticInfo) Get() *LogisticInfo {
	return v.value
}

func (v *NullableLogisticInfo) Set(val *LogisticInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogisticInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogisticInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogisticInfo(val *LogisticInfo) *NullableLogisticInfo {
	return &NullableLogisticInfo{value: val, isSet: true}
}

func (v NullableLogisticInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogisticInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
