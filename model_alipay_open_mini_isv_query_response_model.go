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

// checks if the AlipayOpenMiniIsvQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniIsvQueryResponseModel{}

// AlipayOpenMiniIsvQueryResponseModel struct for AlipayOpenMiniIsvQueryResponseModel
type AlipayOpenMiniIsvQueryResponseModel struct {
	// 小程序appId，商家确认后非空，若商家未确认或超时返回空
	MinAppId *string `json:"min_app_id,omitempty"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty"`
	// 外部订单号，开发者帐号+外部订单号维度来控制请求务幂等
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 商户pid
	Pid *string `json:"pid,omitempty"`
	// PROCESS处理中，TIMEOUT超时，AGREED同意， REJECTED拒绝
	Status *string `json:"status,omitempty"`
}

// NewAlipayOpenMiniIsvQueryResponseModel instantiates a new AlipayOpenMiniIsvQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniIsvQueryResponseModel() *AlipayOpenMiniIsvQueryResponseModel {
	this := AlipayOpenMiniIsvQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniIsvQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniIsvQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniIsvQueryResponseModelWithDefaults() *AlipayOpenMiniIsvQueryResponseModel {
	this := AlipayOpenMiniIsvQueryResponseModel{}
	return &this
}

// GetMinAppId returns the MinAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetMinAppId() string {
	if o == nil || IsNil(o.MinAppId) {
		var ret string
		return ret
	}
	return *o.MinAppId
}

// GetMinAppIdOk returns a tuple with the MinAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetMinAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MinAppId) {
		return nil, false
	}
	return o.MinAppId, true
}

// HasMinAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) HasMinAppId() bool {
	if o != nil && !IsNil(o.MinAppId) {
		return true
	}

	return false
}

// SetMinAppId gets a reference to the given string and assigns it to the MinAppId field.
func (o *AlipayOpenMiniIsvQueryResponseModel) SetMinAppId(v string) {
	o.MinAppId = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayOpenMiniIsvQueryResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayOpenMiniIsvQueryResponseModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayOpenMiniIsvQueryResponseModel) SetPid(v string) {
	o.Pid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenMiniIsvQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayOpenMiniIsvQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniIsvQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinAppId) {
		toSerialize["min_app_id"] = o.MinAppId
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniIsvQueryResponseModel struct {
	value *AlipayOpenMiniIsvQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniIsvQueryResponseModel) Get() *AlipayOpenMiniIsvQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniIsvQueryResponseModel) Set(val *AlipayOpenMiniIsvQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniIsvQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniIsvQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniIsvQueryResponseModel(val *AlipayOpenMiniIsvQueryResponseModel) *NullableAlipayOpenMiniIsvQueryResponseModel {
	return &NullableAlipayOpenMiniIsvQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniIsvQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniIsvQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
