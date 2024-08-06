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

// checks if the OrderApplyStatusBriefDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderApplyStatusBriefDTO{}

// OrderApplyStatusBriefDTO struct for OrderApplyStatusBriefDTO
type OrderApplyStatusBriefDTO struct {
	// SUCCESS：开票成功 FAIL：开票失败 PROCESS：开票中 NOTEXIST：申请不存在
	ApplyStatus *string `json:"apply_status,omitempty"`
	// 开票申请时传入订单号（支持主单号、子单号），不限是否为支付宝体内交易单号
	OrderNo *string `json:"order_no,omitempty"`
}

// NewOrderApplyStatusBriefDTO instantiates a new OrderApplyStatusBriefDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderApplyStatusBriefDTO() *OrderApplyStatusBriefDTO {
	this := OrderApplyStatusBriefDTO{}
	return &this
}

// NewOrderApplyStatusBriefDTOWithDefaults instantiates a new OrderApplyStatusBriefDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderApplyStatusBriefDTOWithDefaults() *OrderApplyStatusBriefDTO {
	this := OrderApplyStatusBriefDTO{}
	return &this
}

// GetApplyStatus returns the ApplyStatus field value if set, zero value otherwise.
func (o *OrderApplyStatusBriefDTO) GetApplyStatus() string {
	if o == nil || IsNil(o.ApplyStatus) {
		var ret string
		return ret
	}
	return *o.ApplyStatus
}

// GetApplyStatusOk returns a tuple with the ApplyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderApplyStatusBriefDTO) GetApplyStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyStatus) {
		return nil, false
	}
	return o.ApplyStatus, true
}

// HasApplyStatus returns a boolean if a field has been set.
func (o *OrderApplyStatusBriefDTO) HasApplyStatus() bool {
	if o != nil && !IsNil(o.ApplyStatus) {
		return true
	}

	return false
}

// SetApplyStatus gets a reference to the given string and assigns it to the ApplyStatus field.
func (o *OrderApplyStatusBriefDTO) SetApplyStatus(v string) {
	o.ApplyStatus = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *OrderApplyStatusBriefDTO) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderApplyStatusBriefDTO) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *OrderApplyStatusBriefDTO) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *OrderApplyStatusBriefDTO) SetOrderNo(v string) {
	o.OrderNo = &v
}

func (o OrderApplyStatusBriefDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderApplyStatusBriefDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyStatus) {
		toSerialize["apply_status"] = o.ApplyStatus
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	return toSerialize, nil
}

type NullableOrderApplyStatusBriefDTO struct {
	value *OrderApplyStatusBriefDTO
	isSet bool
}

func (v NullableOrderApplyStatusBriefDTO) Get() *OrderApplyStatusBriefDTO {
	return v.value
}

func (v *NullableOrderApplyStatusBriefDTO) Set(val *OrderApplyStatusBriefDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderApplyStatusBriefDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderApplyStatusBriefDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderApplyStatusBriefDTO(val *OrderApplyStatusBriefDTO) *NullableOrderApplyStatusBriefDTO {
	return &NullableOrderApplyStatusBriefDTO{value: val, isSet: true}
}

func (v NullableOrderApplyStatusBriefDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderApplyStatusBriefDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

