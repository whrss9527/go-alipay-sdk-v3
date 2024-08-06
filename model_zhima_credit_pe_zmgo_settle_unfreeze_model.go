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

// checks if the ZhimaCreditPeZmgoSettleUnfreezeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCreditPeZmgoSettleUnfreezeModel{}

// ZhimaCreditPeZmgoSettleUnfreezeModel struct for ZhimaCreditPeZmgoSettleUnfreezeModel
type ZhimaCreditPeZmgoSettleUnfreezeModel struct {
	// 支付宝系统中用以唯一标识用户签约记录的编号，即花呗先享协议号
	AgreementId *string `json:"agreement_id,omitempty"`
	// 买家在支付宝的用户id
	AlipayOpenId *string `json:"alipay_open_id,omitempty"`
	// 买家在支付宝的用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// 解冻成功时间
	BizTime *string `json:"biz_time,omitempty"`
	// 解冻的描述
	OrderTitle *string `json:"order_title,omitempty"`
	// 商户本次操作的请求流水号，用于标识请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。由商户传入，最终返回给商户。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户ID
	PartnerId *string `json:"partner_id,omitempty"`
	// 解冻金额
	UnfreezeAmount *string `json:"unfreeze_amount,omitempty"`
	UnfreezeExtendParams *UnfreezeExtendParams `json:"unfreeze_extend_params,omitempty"`
}

// NewZhimaCreditPeZmgoSettleUnfreezeModel instantiates a new ZhimaCreditPeZmgoSettleUnfreezeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCreditPeZmgoSettleUnfreezeModel() *ZhimaCreditPeZmgoSettleUnfreezeModel {
	this := ZhimaCreditPeZmgoSettleUnfreezeModel{}
	return &this
}

// NewZhimaCreditPeZmgoSettleUnfreezeModelWithDefaults instantiates a new ZhimaCreditPeZmgoSettleUnfreezeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCreditPeZmgoSettleUnfreezeModelWithDefaults() *ZhimaCreditPeZmgoSettleUnfreezeModel {
	this := ZhimaCreditPeZmgoSettleUnfreezeModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAlipayOpenId returns the AlipayOpenId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAlipayOpenId() string {
	if o == nil || IsNil(o.AlipayOpenId) {
		var ret string
		return ret
	}
	return *o.AlipayOpenId
}

// GetAlipayOpenIdOk returns a tuple with the AlipayOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAlipayOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayOpenId) {
		return nil, false
	}
	return o.AlipayOpenId, true
}

// HasAlipayOpenId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasAlipayOpenId() bool {
	if o != nil && !IsNil(o.AlipayOpenId) {
		return true
	}

	return false
}

// SetAlipayOpenId gets a reference to the given string and assigns it to the AlipayOpenId field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetAlipayOpenId(v string) {
	o.AlipayOpenId = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetBizTime returns the BizTime field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetBizTime() string {
	if o == nil || IsNil(o.BizTime) {
		var ret string
		return ret
	}
	return *o.BizTime
}

// GetBizTimeOk returns a tuple with the BizTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetBizTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BizTime) {
		return nil, false
	}
	return o.BizTime, true
}

// HasBizTime returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasBizTime() bool {
	if o != nil && !IsNil(o.BizTime) {
		return true
	}

	return false
}

// SetBizTime gets a reference to the given string and assigns it to the BizTime field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetBizTime(v string) {
	o.BizTime = &v
}

// GetOrderTitle returns the OrderTitle field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetOrderTitle() string {
	if o == nil || IsNil(o.OrderTitle) {
		var ret string
		return ret
	}
	return *o.OrderTitle
}

// GetOrderTitleOk returns a tuple with the OrderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetOrderTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OrderTitle) {
		return nil, false
	}
	return o.OrderTitle, true
}

// HasOrderTitle returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasOrderTitle() bool {
	if o != nil && !IsNil(o.OrderTitle) {
		return true
	}

	return false
}

// SetOrderTitle gets a reference to the given string and assigns it to the OrderTitle field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetOrderTitle(v string) {
	o.OrderTitle = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetUnfreezeAmount returns the UnfreezeAmount field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetUnfreezeAmount() string {
	if o == nil || IsNil(o.UnfreezeAmount) {
		var ret string
		return ret
	}
	return *o.UnfreezeAmount
}

// GetUnfreezeAmountOk returns a tuple with the UnfreezeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetUnfreezeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.UnfreezeAmount) {
		return nil, false
	}
	return o.UnfreezeAmount, true
}

// HasUnfreezeAmount returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasUnfreezeAmount() bool {
	if o != nil && !IsNil(o.UnfreezeAmount) {
		return true
	}

	return false
}

// SetUnfreezeAmount gets a reference to the given string and assigns it to the UnfreezeAmount field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetUnfreezeAmount(v string) {
	o.UnfreezeAmount = &v
}

// GetUnfreezeExtendParams returns the UnfreezeExtendParams field value if set, zero value otherwise.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetUnfreezeExtendParams() UnfreezeExtendParams {
	if o == nil || IsNil(o.UnfreezeExtendParams) {
		var ret UnfreezeExtendParams
		return ret
	}
	return *o.UnfreezeExtendParams
}

// GetUnfreezeExtendParamsOk returns a tuple with the UnfreezeExtendParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) GetUnfreezeExtendParamsOk() (*UnfreezeExtendParams, bool) {
	if o == nil || IsNil(o.UnfreezeExtendParams) {
		return nil, false
	}
	return o.UnfreezeExtendParams, true
}

// HasUnfreezeExtendParams returns a boolean if a field has been set.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) HasUnfreezeExtendParams() bool {
	if o != nil && !IsNil(o.UnfreezeExtendParams) {
		return true
	}

	return false
}

// SetUnfreezeExtendParams gets a reference to the given UnfreezeExtendParams and assigns it to the UnfreezeExtendParams field.
func (o *ZhimaCreditPeZmgoSettleUnfreezeModel) SetUnfreezeExtendParams(v UnfreezeExtendParams) {
	o.UnfreezeExtendParams = &v
}

func (o ZhimaCreditPeZmgoSettleUnfreezeModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCreditPeZmgoSettleUnfreezeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AlipayOpenId) {
		toSerialize["alipay_open_id"] = o.AlipayOpenId
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.BizTime) {
		toSerialize["biz_time"] = o.BizTime
	}
	if !IsNil(o.OrderTitle) {
		toSerialize["order_title"] = o.OrderTitle
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partner_id"] = o.PartnerId
	}
	if !IsNil(o.UnfreezeAmount) {
		toSerialize["unfreeze_amount"] = o.UnfreezeAmount
	}
	if !IsNil(o.UnfreezeExtendParams) {
		toSerialize["unfreeze_extend_params"] = o.UnfreezeExtendParams
	}
	return toSerialize, nil
}

type NullableZhimaCreditPeZmgoSettleUnfreezeModel struct {
	value *ZhimaCreditPeZmgoSettleUnfreezeModel
	isSet bool
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeModel) Get() *ZhimaCreditPeZmgoSettleUnfreezeModel {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeModel) Set(val *ZhimaCreditPeZmgoSettleUnfreezeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoSettleUnfreezeModel(val *ZhimaCreditPeZmgoSettleUnfreezeModel) *NullableZhimaCreditPeZmgoSettleUnfreezeModel {
	return &NullableZhimaCreditPeZmgoSettleUnfreezeModel{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoSettleUnfreezeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoSettleUnfreezeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


