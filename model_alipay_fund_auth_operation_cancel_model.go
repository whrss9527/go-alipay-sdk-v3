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

// checks if the AlipayFundAuthOperationCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAuthOperationCancelModel{}

// AlipayFundAuthOperationCancelModel struct for AlipayFundAuthOperationCancelModel
type AlipayFundAuthOperationCancelModel struct {
	// 支付宝授权资金订单号。 与商户的授权资金订单号不能同时为空，二者都传入时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。
	AuthNo *string `json:"auth_no,omitempty"`
	// 通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 支付宝的授权资金操作流水号。 与商户的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。
	OperationId *string `json:"operation_id,omitempty"`
	// 商户的授权资金订单号。 与支付宝的授权资金订单号不能同时为空，二者都传入时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。 该值与资金冻结时 out_order_no一致。
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 商户的授权资金操作流水号。 与支付宝的授权资金操作流水号不能同时为空，二者都传入时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。 该值与资金冻结时out_request_no一致
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户对本次撤销操作的附言描述。 长度不超过100个字母或50个汉字。
	Remark *string `json:"remark,omitempty"`
}

// NewAlipayFundAuthOperationCancelModel instantiates a new AlipayFundAuthOperationCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAuthOperationCancelModel() *AlipayFundAuthOperationCancelModel {
	this := AlipayFundAuthOperationCancelModel{}
	return &this
}

// NewAlipayFundAuthOperationCancelModelWithDefaults instantiates a new AlipayFundAuthOperationCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAuthOperationCancelModelWithDefaults() *AlipayFundAuthOperationCancelModel {
	this := AlipayFundAuthOperationCancelModel{}
	return &this
}

// GetAuthNo returns the AuthNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetAuthNo() string {
	if o == nil || IsNil(o.AuthNo) {
		var ret string
		return ret
	}
	return *o.AuthNo
}

// GetAuthNoOk returns a tuple with the AuthNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetAuthNoOk() (*string, bool) {
	if o == nil || IsNil(o.AuthNo) {
		return nil, false
	}
	return o.AuthNo, true
}

// HasAuthNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasAuthNo() bool {
	if o != nil && !IsNil(o.AuthNo) {
		return true
	}

	return false
}

// SetAuthNo gets a reference to the given string and assigns it to the AuthNo field.
func (o *AlipayFundAuthOperationCancelModel) SetAuthNo(v string) {
	o.AuthNo = &v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *AlipayFundAuthOperationCancelModel) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetOperationId() string {
	if o == nil || IsNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetOperationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasOperationId() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *AlipayFundAuthOperationCancelModel) SetOperationId(v string) {
	o.OperationId = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayFundAuthOperationCancelModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayFundAuthOperationCancelModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AlipayFundAuthOperationCancelModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOperationCancelModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AlipayFundAuthOperationCancelModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AlipayFundAuthOperationCancelModel) SetRemark(v string) {
	o.Remark = &v
}

func (o AlipayFundAuthOperationCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAuthOperationCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthNo) {
		toSerialize["auth_no"] = o.AuthNo
	}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
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
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableAlipayFundAuthOperationCancelModel struct {
	value *AlipayFundAuthOperationCancelModel
	isSet bool
}

func (v NullableAlipayFundAuthOperationCancelModel) Get() *AlipayFundAuthOperationCancelModel {
	return v.value
}

func (v *NullableAlipayFundAuthOperationCancelModel) Set(val *AlipayFundAuthOperationCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOperationCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOperationCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOperationCancelModel(val *AlipayFundAuthOperationCancelModel) *NullableAlipayFundAuthOperationCancelModel {
	return &NullableAlipayFundAuthOperationCancelModel{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOperationCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOperationCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


