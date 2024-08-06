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

// checks if the AlipayEbppPdeductBillPayStatusResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductBillPayStatusResponseModel{}

// AlipayEbppPdeductBillPayStatusResponseModel struct for AlipayEbppPdeductBillPayStatusResponseModel
type AlipayEbppPdeductBillPayStatusResponseModel struct {
	// 支付宝协议流水
	AgreementId *string `json:"agreement_id,omitempty"`
	// 用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// 支付宝流billNo
	OrderNo *string `json:"order_no,omitempty"`
	// 订单的结果码
	OrderResultCode *string `json:"order_result_code,omitempty"`
	// 订单的结果描述
	OrderResultMsg *string `json:"order_result_msg,omitempty"`
	// 外部订单流水
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 支付宝订单支付状态。  0：未知状态。  1：支付成功。  2：支付失败。
	Status *string `json:"status,omitempty"`
}

// NewAlipayEbppPdeductBillPayStatusResponseModel instantiates a new AlipayEbppPdeductBillPayStatusResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductBillPayStatusResponseModel() *AlipayEbppPdeductBillPayStatusResponseModel {
	this := AlipayEbppPdeductBillPayStatusResponseModel{}
	return &this
}

// NewAlipayEbppPdeductBillPayStatusResponseModelWithDefaults instantiates a new AlipayEbppPdeductBillPayStatusResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductBillPayStatusResponseModelWithDefaults() *AlipayEbppPdeductBillPayStatusResponseModel {
	this := AlipayEbppPdeductBillPayStatusResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetOrderResultCode returns the OrderResultCode field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderResultCode() string {
	if o == nil || IsNil(o.OrderResultCode) {
		var ret string
		return ret
	}
	return *o.OrderResultCode
}

// GetOrderResultCodeOk returns a tuple with the OrderResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderResultCode) {
		return nil, false
	}
	return o.OrderResultCode, true
}

// HasOrderResultCode returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasOrderResultCode() bool {
	if o != nil && !IsNil(o.OrderResultCode) {
		return true
	}

	return false
}

// SetOrderResultCode gets a reference to the given string and assigns it to the OrderResultCode field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetOrderResultCode(v string) {
	o.OrderResultCode = &v
}

// GetOrderResultMsg returns the OrderResultMsg field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderResultMsg() string {
	if o == nil || IsNil(o.OrderResultMsg) {
		var ret string
		return ret
	}
	return *o.OrderResultMsg
}

// GetOrderResultMsgOk returns a tuple with the OrderResultMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOrderResultMsgOk() (*string, bool) {
	if o == nil || IsNil(o.OrderResultMsg) {
		return nil, false
	}
	return o.OrderResultMsg, true
}

// HasOrderResultMsg returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasOrderResultMsg() bool {
	if o != nil && !IsNil(o.OrderResultMsg) {
		return true
	}

	return false
}

// SetOrderResultMsg gets a reference to the given string and assigns it to the OrderResultMsg field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetOrderResultMsg(v string) {
	o.OrderResultMsg = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayEbppPdeductBillPayStatusResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayEbppPdeductBillPayStatusResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductBillPayStatusResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.OrderResultCode) {
		toSerialize["order_result_code"] = o.OrderResultCode
	}
	if !IsNil(o.OrderResultMsg) {
		toSerialize["order_result_msg"] = o.OrderResultMsg
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductBillPayStatusResponseModel struct {
	value *AlipayEbppPdeductBillPayStatusResponseModel
	isSet bool
}

func (v NullableAlipayEbppPdeductBillPayStatusResponseModel) Get() *AlipayEbppPdeductBillPayStatusResponseModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductBillPayStatusResponseModel) Set(val *AlipayEbppPdeductBillPayStatusResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductBillPayStatusResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductBillPayStatusResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductBillPayStatusResponseModel(val *AlipayEbppPdeductBillPayStatusResponseModel) *NullableAlipayEbppPdeductBillPayStatusResponseModel {
	return &NullableAlipayEbppPdeductBillPayStatusResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductBillPayStatusResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductBillPayStatusResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
