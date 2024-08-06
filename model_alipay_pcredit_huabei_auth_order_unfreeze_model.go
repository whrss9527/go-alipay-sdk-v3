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

// checks if the AlipayPcreditHuabeiAuthOrderUnfreezeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayPcreditHuabeiAuthOrderUnfreezeModel{}

// AlipayPcreditHuabeiAuthOrderUnfreezeModel struct for AlipayPcreditHuabeiAuthOrderUnfreezeModel
type AlipayPcreditHuabeiAuthOrderUnfreezeModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// true表示需要解约。false或者不填写表示仅解冻。
	NeedTerminated *string `json:"need_terminated,omitempty"`
	// 买家在支付宝的用户id
	OpenId *string `json:"open_id,omitempty"`
	// 商户业务订单的简单描述，如商品名称等，长度不超过100个字母或50个汉字
	OrderTitle *string `json:"order_title,omitempty"`
	// 商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户的支付宝用户id。如果该值为空，则默认为商户签约账号对应的支付宝用户ID。
	SellerId *string `json:"seller_id,omitempty"`
	// 需要解冻的金额，单位为：元（人民币），精确到小数点后两位
	UnfreezeAmount *string `json:"unfreeze_amount,omitempty"`
}

// NewAlipayPcreditHuabeiAuthOrderUnfreezeModel instantiates a new AlipayPcreditHuabeiAuthOrderUnfreezeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayPcreditHuabeiAuthOrderUnfreezeModel() *AlipayPcreditHuabeiAuthOrderUnfreezeModel {
	this := AlipayPcreditHuabeiAuthOrderUnfreezeModel{}
	return &this
}

// NewAlipayPcreditHuabeiAuthOrderUnfreezeModelWithDefaults instantiates a new AlipayPcreditHuabeiAuthOrderUnfreezeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayPcreditHuabeiAuthOrderUnfreezeModelWithDefaults() *AlipayPcreditHuabeiAuthOrderUnfreezeModel {
	this := AlipayPcreditHuabeiAuthOrderUnfreezeModel{}
	return &this
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetNeedTerminated returns the NeedTerminated field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetNeedTerminated() string {
	if o == nil || IsNil(o.NeedTerminated) {
		var ret string
		return ret
	}
	return *o.NeedTerminated
}

// GetNeedTerminatedOk returns a tuple with the NeedTerminated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetNeedTerminatedOk() (*string, bool) {
	if o == nil || IsNil(o.NeedTerminated) {
		return nil, false
	}
	return o.NeedTerminated, true
}

// HasNeedTerminated returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasNeedTerminated() bool {
	if o != nil && !IsNil(o.NeedTerminated) {
		return true
	}

	return false
}

// SetNeedTerminated gets a reference to the given string and assigns it to the NeedTerminated field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetNeedTerminated(v string) {
	o.NeedTerminated = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOrderTitle returns the OrderTitle field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOrderTitle() string {
	if o == nil || IsNil(o.OrderTitle) {
		var ret string
		return ret
	}
	return *o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOrderTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OrderTitle) {
		return nil, false
	}
	return o.OrderTitle, true
}

// HasOrderTitle returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasOrderTitle() bool {
	if o != nil && !IsNil(o.OrderTitle) {
		return true
	}

	return false
}

// SetOrderTitle gets a reference to the given string and assigns it to the OrderTitle field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetOrderTitle(v string) {
	o.OrderTitle = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetSellerId returns the SellerId field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetSellerId() string {
	if o == nil || IsNil(o.SellerId) {
		var ret string
		return ret
	}
	return *o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetSellerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SellerId) {
		return nil, false
	}
	return o.SellerId, true
}

// HasSellerId returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasSellerId() bool {
	if o != nil && !IsNil(o.SellerId) {
		return true
	}

	return false
}

// SetSellerId gets a reference to the given string and assigns it to the SellerId field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetSellerId(v string) {
	o.SellerId = &v
}

// GetUnfreezeAmount returns the UnfreezeAmount field value if set, zero value otherwise.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetUnfreezeAmount() string {
	if o == nil || IsNil(o.UnfreezeAmount) {
		var ret string
		return ret
	}
	return *o.UnfreezeAmount
}

// GetUnfreezeAmountOk returns a tuple with the UnfreezeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) GetUnfreezeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.UnfreezeAmount) {
		return nil, false
	}
	return o.UnfreezeAmount, true
}

// HasUnfreezeAmount returns a boolean if a field has been set.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) HasUnfreezeAmount() bool {
	if o != nil && !IsNil(o.UnfreezeAmount) {
		return true
	}

	return false
}

// SetUnfreezeAmount gets a reference to the given string and assigns it to the UnfreezeAmount field.
func (o *AlipayPcreditHuabeiAuthOrderUnfreezeModel) SetUnfreezeAmount(v string) {
	o.UnfreezeAmount = &v
}

func (o AlipayPcreditHuabeiAuthOrderUnfreezeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayPcreditHuabeiAuthOrderUnfreezeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.NeedTerminated) {
		toSerialize["need_terminated"] = o.NeedTerminated
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OrderTitle) {
		toSerialize["order_title"] = o.OrderTitle
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.SellerId) {
		toSerialize["seller_id"] = o.SellerId
	}
	if !IsNil(o.UnfreezeAmount) {
		toSerialize["unfreeze_amount"] = o.UnfreezeAmount
	}
	return toSerialize, nil
}

type NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel struct {
	value *AlipayPcreditHuabeiAuthOrderUnfreezeModel
	isSet bool
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) Get() *AlipayPcreditHuabeiAuthOrderUnfreezeModel {
	return v.value
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) Set(val *AlipayPcreditHuabeiAuthOrderUnfreezeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayPcreditHuabeiAuthOrderUnfreezeModel(val *AlipayPcreditHuabeiAuthOrderUnfreezeModel) *NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel {
	return &NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel{value: val, isSet: true}
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


