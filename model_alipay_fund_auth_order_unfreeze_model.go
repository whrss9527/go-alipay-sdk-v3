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

// checks if the AlipayFundAuthOrderUnfreezeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundAuthOrderUnfreezeModel{}

// AlipayFundAuthOrderUnfreezeModel struct for AlipayFundAuthOrderUnfreezeModel
type AlipayFundAuthOrderUnfreezeModel struct {
	// 本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位。 取值范围：[0.01,100000000.00]
	Amount *string `json:"amount,omitempty"`
	// 支付宝资金授权订单号。
	AuthNo *string `json:"auth_no,omitempty"`
	// 解冻扩展信息。map<string, string>的json格式，目前支持如下key： unfreezeBizInfo：由芝麻消费，当前支持value如下： \"bizComplete\":\"true\"——标识本次解冻用户已履约，true表示信用单履约完结
	ExtraParam *string `json:"extra_param,omitempty"`
	// 通知地址
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 解冻请求流水号。 如果是针对同一笔授权单不同的解冻请求，如第一次解冻1元，第二次解冻2元，则解冻请求流水号必须不重复； 如果是针对同一笔解冻请求的多次发起，则需要保证每次发起，解冻请求流水号和解冻金额都相同
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户对本次解冻操作的附言描述
	Remark *string `json:"remark,omitempty"`
}

// NewAlipayFundAuthOrderUnfreezeModel instantiates a new AlipayFundAuthOrderUnfreezeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundAuthOrderUnfreezeModel() *AlipayFundAuthOrderUnfreezeModel {
	this := AlipayFundAuthOrderUnfreezeModel{}
	return &this
}

// NewAlipayFundAuthOrderUnfreezeModelWithDefaults instantiates a new AlipayFundAuthOrderUnfreezeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundAuthOrderUnfreezeModelWithDefaults() *AlipayFundAuthOrderUnfreezeModel {
	this := AlipayFundAuthOrderUnfreezeModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetAmount(v string) {
	o.Amount = &v
}

// GetAuthNo returns the AuthNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetAuthNo() string {
	if o == nil || IsNil(o.AuthNo) {
		var ret string
		return ret
	}
	return *o.AuthNo
}

// GetAuthNoOk returns a tuple with the AuthNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetAuthNoOk() (*string, bool) {
	if o == nil || IsNil(o.AuthNo) {
		return nil, false
	}
	return o.AuthNo, true
}

// HasAuthNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasAuthNo() bool {
	if o != nil && !IsNil(o.AuthNo) {
		return true
	}

	return false
}

// SetAuthNo gets a reference to the given string and assigns it to the AuthNo field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetAuthNo(v string) {
	o.AuthNo = &v
}

// GetExtraParam returns the ExtraParam field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetExtraParam() string {
	if o == nil || IsNil(o.ExtraParam) {
		var ret string
		return ret
	}
	return *o.ExtraParam
}

// GetExtraParamOk returns a tuple with the ExtraParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetExtraParamOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraParam) {
		return nil, false
	}
	return o.ExtraParam, true
}

// HasExtraParam returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasExtraParam() bool {
	if o != nil && !IsNil(o.ExtraParam) {
		return true
	}

	return false
}

// SetExtraParam gets a reference to the given string and assigns it to the ExtraParam field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetExtraParam(v string) {
	o.ExtraParam = &v
}

// GetNotifyUrl returns the NotifyUrl field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetNotifyUrl() string {
	if o == nil || IsNil(o.NotifyUrl) {
		var ret string
		return ret
	}
	return *o.NotifyUrl
}

// GetNotifyUrlOk returns a tuple with the NotifyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetNotifyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyUrl) {
		return nil, false
	}
	return o.NotifyUrl, true
}

// HasNotifyUrl returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasNotifyUrl() bool {
	if o != nil && !IsNil(o.NotifyUrl) {
		return true
	}

	return false
}

// SetNotifyUrl gets a reference to the given string and assigns it to the NotifyUrl field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetNotifyUrl(v string) {
	o.NotifyUrl = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AlipayFundAuthOrderUnfreezeModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AlipayFundAuthOrderUnfreezeModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AlipayFundAuthOrderUnfreezeModel) SetRemark(v string) {
	o.Remark = &v
}

func (o AlipayFundAuthOrderUnfreezeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundAuthOrderUnfreezeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AuthNo) {
		toSerialize["auth_no"] = o.AuthNo
	}
	if !IsNil(o.ExtraParam) {
		toSerialize["extra_param"] = o.ExtraParam
	}
	if !IsNil(o.NotifyUrl) {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	return toSerialize, nil
}

type NullableAlipayFundAuthOrderUnfreezeModel struct {
	value *AlipayFundAuthOrderUnfreezeModel
	isSet bool
}

func (v NullableAlipayFundAuthOrderUnfreezeModel) Get() *AlipayFundAuthOrderUnfreezeModel {
	return v.value
}

func (v *NullableAlipayFundAuthOrderUnfreezeModel) Set(val *AlipayFundAuthOrderUnfreezeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOrderUnfreezeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOrderUnfreezeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOrderUnfreezeModel(val *AlipayFundAuthOrderUnfreezeModel) *NullableAlipayFundAuthOrderUnfreezeModel {
	return &NullableAlipayFundAuthOrderUnfreezeModel{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOrderUnfreezeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOrderUnfreezeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
