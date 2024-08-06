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

// checks if the AlipayFundAuthOperationCancelResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAuthOperationCancelResponseModel{}

// AlipayFundAuthOperationCancelResponseModel struct for AlipayFundAuthOperationCancelResponseModel
type AlipayFundAuthOperationCancelResponseModel struct {
	// 本次撤销触发的资金动作  close：关闭冻结明细，无资金解冻  unfreeze：产生了资金解冻
	Action *string `json:"action,omitempty"`
	// 支付宝资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// 支付宝的冻结操作流水号。
	OperationId *string `json:"operation_id,omitempty"`
	// 商户的授权资金订单号。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 商户的冻结操作流水号 。
	OutRequestNo *string `json:"out_request_no,omitempty"`
}

// NewAlipayFundAuthOperationCancelResponseModel instantiates a new AlipayFundAuthOperationCancelResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAuthOperationCancelResponseModel() *AlipayFundAuthOperationCancelResponseModel {
	this := AlipayFundAuthOperationCancelResponseModel{}
	return &this
}

// NewAlipayFundAuthOperationCancelResponseModelWithDefaults instantiates a new AlipayFundAuthOperationCancelResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAuthOperationCancelResponseModelWithDefaults() *AlipayFundAuthOperationCancelResponseModel {
	this := AlipayFundAuthOperationCancelResponseModel{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelResponseModel) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AlipayFundAuthOperationCancelResponseModel) SetAction(v string) {
	o.Action = &v
}

// GetAuthNo returns the AuthNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelResponseModel) GetAuthNo() string {
	if o == nil || IsNil(o.AuthNo) {
		var ret string
		return ret
	}
	return *o.AuthNo
}

// GetAuthNoOk returns a tuple with the AuthNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) GetAuthNoOk() (*string, bool) {
	if o == nil || IsNil(o.AuthNo) {
		return nil, false
	}
	return o.AuthNo, true
}

// HasAuthNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) HasAuthNo() bool {
	if o != nil && !IsNil(o.AuthNo) {
		return true
	}

	return false
}

// SetAuthNo gets a reference to the given string and assigns it to the AuthNo field.
func (o *AlipayFundAuthOperationCancelResponseModel) SetAuthNo(v string) {
	o.AuthNo = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOperationId() string {
	if o == nil || IsNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOperationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) HasOperationId() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *AlipayFundAuthOperationCancelResponseModel) SetOperationId(v string) {
	o.OperationId = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayFundAuthOperationCancelResponseModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelResponseModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayFundAuthOperationCancelResponseModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

func (o AlipayFundAuthOperationCancelResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAuthOperationCancelResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AuthNo) {
		toSerialize["auth_no"] = o.AuthNo
	}
	if !IsNil(o.OperationId) {
		toSerialize["operation_id"] = o.OperationId
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	return toSerialize, nil
}

type NullableAlipayFundAuthOperationCancelResponseModel struct {
	value *AlipayFundAuthOperationCancelResponseModel
	isSet bool
}

func (v NullableAlipayFundAuthOperationCancelResponseModel) Get() *AlipayFundAuthOperationCancelResponseModel {
	return v.value
}

func (v *NullableAlipayFundAuthOperationCancelResponseModel) Set(val *AlipayFundAuthOperationCancelResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOperationCancelResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOperationCancelResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOperationCancelResponseModel(val *AlipayFundAuthOperationCancelResponseModel) *NullableAlipayFundAuthOperationCancelResponseModel {
	return &NullableAlipayFundAuthOperationCancelResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOperationCancelResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOperationCancelResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


