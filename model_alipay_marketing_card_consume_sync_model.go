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

// checks if the AlipayMarketingCardConsumeSyncModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardConsumeSyncModel{}

// AlipayMarketingCardConsumeSyncModel struct for AlipayMarketingCardConsumeSyncModel
type AlipayMarketingCardConsumeSyncModel struct {
	// 用户实际付的现金金额  1.针对预付卡面额的核销金额在use_benefit_list展现，作为权益金额  2.权益金额不叠加在该金额上
	ActPayAmount *string       `json:"act_pay_amount,omitempty"`
	CardInfo     *MerchantCard `json:"card_info,omitempty"`
	// 获取权益列表
	GainBenefitList []BenefitInfoDetail `json:"gain_benefit_list,omitempty"`
	// 备注信息，现有直接填写门店信息
	Memo *string `json:"memo,omitempty"`
	// 门店编号
	ShopCode *string `json:"shop_code,omitempty"`
	// 产生该笔交易时，用户出具的凭证类型。枚举支持： *ALIPAY：支付宝电子卡； *ENTITY：实体卡； *OTHER：其他。
	SwipeCertType *string `json:"swipe_cert_type,omitempty"`
	// 支付宝业务卡号，即通过<a  href=\"https://opendocs.alipay.com/apis/009zw3\">alipay.marketing.card.open</a>(会员卡开卡)接口开卡后获取的  card_info.biz_card_no 值。
	TargetCardNo *string `json:"target_card_no,omitempty"`
	// 卡号ID类型。支持： *BIZ_CARD：支付宝业务卡号（商户会员卡场景使用）。
	TargetCardNoType *string `json:"target_card_no_type,omitempty"`
	// 交易金额：本次交易的实际总金额（可认为标价金额）
	TradeAmount *string `json:"trade_amount,omitempty"`
	// 交易名称。为空时支付宝将根据交易类型提供默认名称。
	TradeName *string `json:"trade_name,omitempty"`
	// 支付宝交易号
	TradeNo *string `json:"trade_no,omitempty"`
	// 线下交易时间，为用户付款的交易时间。 说明：当交易时间晚于上次消费记录同步时间时，将变更会员卡信息。
	TradeTime *string `json:"trade_time,omitempty"`
	// 交易类型。枚举支持： *开卡：OPEN； *消费：TRADE； *充值：DEPOSIT； *退卡：RETURN。
	TradeType *string `json:"trade_type,omitempty"`
	// 实际消耗的权益
	UseBenefitList []BenefitInfoDetail `json:"use_benefit_list,omitempty"`
}

// NewAlipayMarketingCardConsumeSyncModel instantiates a new AlipayMarketingCardConsumeSyncModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardConsumeSyncModel() *AlipayMarketingCardConsumeSyncModel {
	this := AlipayMarketingCardConsumeSyncModel{}
	return &this
}

// NewAlipayMarketingCardConsumeSyncModelWithDefaults instantiates a new AlipayMarketingCardConsumeSyncModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardConsumeSyncModelWithDefaults() *AlipayMarketingCardConsumeSyncModel {
	this := AlipayMarketingCardConsumeSyncModel{}
	return &this
}

// GetActPayAmount returns the ActPayAmount field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetActPayAmount() string {
	if o == nil || IsNil(o.ActPayAmount) {
		var ret string
		return ret
	}
	return *o.ActPayAmount
}

// GetActPayAmountOk returns a tuple with the ActPayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetActPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ActPayAmount) {
		return nil, false
	}
	return o.ActPayAmount, true
}

// HasActPayAmount returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasActPayAmount() bool {
	if o != nil && !IsNil(o.ActPayAmount) {
		return true
	}

	return false
}

// SetActPayAmount gets a reference to the given string and assigns it to the ActPayAmount field.
func (o *AlipayMarketingCardConsumeSyncModel) SetActPayAmount(v string) {
	o.ActPayAmount = &v
}

// GetCardInfo returns the CardInfo field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetCardInfo() MerchantCard {
	if o == nil || IsNil(o.CardInfo) {
		var ret MerchantCard
		return ret
	}
	return *o.CardInfo
}

// GetCardInfoOk returns a tuple with the CardInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetCardInfoOk() (*MerchantCard, bool) {
	if o == nil || IsNil(o.CardInfo) {
		return nil, false
	}
	return o.CardInfo, true
}

// HasCardInfo returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasCardInfo() bool {
	if o != nil && !IsNil(o.CardInfo) {
		return true
	}

	return false
}

// SetCardInfo gets a reference to the given MerchantCard and assigns it to the CardInfo field.
func (o *AlipayMarketingCardConsumeSyncModel) SetCardInfo(v MerchantCard) {
	o.CardInfo = &v
}

// GetGainBenefitList returns the GainBenefitList field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetGainBenefitList() []BenefitInfoDetail {
	if o == nil || IsNil(o.GainBenefitList) {
		var ret []BenefitInfoDetail
		return ret
	}
	return o.GainBenefitList
}

// GetGainBenefitListOk returns a tuple with the GainBenefitList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetGainBenefitListOk() ([]BenefitInfoDetail, bool) {
	if o == nil || IsNil(o.GainBenefitList) {
		return nil, false
	}
	return o.GainBenefitList, true
}

// HasGainBenefitList returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasGainBenefitList() bool {
	if o != nil && !IsNil(o.GainBenefitList) {
		return true
	}

	return false
}

// SetGainBenefitList gets a reference to the given []BenefitInfoDetail and assigns it to the GainBenefitList field.
func (o *AlipayMarketingCardConsumeSyncModel) SetGainBenefitList(v []BenefitInfoDetail) {
	o.GainBenefitList = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AlipayMarketingCardConsumeSyncModel) SetMemo(v string) {
	o.Memo = &v
}

// GetShopCode returns the ShopCode field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetShopCode() string {
	if o == nil || IsNil(o.ShopCode) {
		var ret string
		return ret
	}
	return *o.ShopCode
}

// GetShopCodeOk returns a tuple with the ShopCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetShopCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShopCode) {
		return nil, false
	}
	return o.ShopCode, true
}

// HasShopCode returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasShopCode() bool {
	if o != nil && !IsNil(o.ShopCode) {
		return true
	}

	return false
}

// SetShopCode gets a reference to the given string and assigns it to the ShopCode field.
func (o *AlipayMarketingCardConsumeSyncModel) SetShopCode(v string) {
	o.ShopCode = &v
}

// GetSwipeCertType returns the SwipeCertType field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetSwipeCertType() string {
	if o == nil || IsNil(o.SwipeCertType) {
		var ret string
		return ret
	}
	return *o.SwipeCertType
}

// GetSwipeCertTypeOk returns a tuple with the SwipeCertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetSwipeCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SwipeCertType) {
		return nil, false
	}
	return o.SwipeCertType, true
}

// HasSwipeCertType returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasSwipeCertType() bool {
	if o != nil && !IsNil(o.SwipeCertType) {
		return true
	}

	return false
}

// SetSwipeCertType gets a reference to the given string and assigns it to the SwipeCertType field.
func (o *AlipayMarketingCardConsumeSyncModel) SetSwipeCertType(v string) {
	o.SwipeCertType = &v
}

// GetTargetCardNo returns the TargetCardNo field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTargetCardNo() string {
	if o == nil || IsNil(o.TargetCardNo) {
		var ret string
		return ret
	}
	return *o.TargetCardNo
}

// GetTargetCardNoOk returns a tuple with the TargetCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTargetCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCardNo) {
		return nil, false
	}
	return o.TargetCardNo, true
}

// HasTargetCardNo returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTargetCardNo() bool {
	if o != nil && !IsNil(o.TargetCardNo) {
		return true
	}

	return false
}

// SetTargetCardNo gets a reference to the given string and assigns it to the TargetCardNo field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTargetCardNo(v string) {
	o.TargetCardNo = &v
}

// GetTargetCardNoType returns the TargetCardNoType field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTargetCardNoType() string {
	if o == nil || IsNil(o.TargetCardNoType) {
		var ret string
		return ret
	}
	return *o.TargetCardNoType
}

// GetTargetCardNoTypeOk returns a tuple with the TargetCardNoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTargetCardNoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetCardNoType) {
		return nil, false
	}
	return o.TargetCardNoType, true
}

// HasTargetCardNoType returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTargetCardNoType() bool {
	if o != nil && !IsNil(o.TargetCardNoType) {
		return true
	}

	return false
}

// SetTargetCardNoType gets a reference to the given string and assigns it to the TargetCardNoType field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTargetCardNoType(v string) {
	o.TargetCardNoType = &v
}

// GetTradeAmount returns the TradeAmount field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeAmount() string {
	if o == nil || IsNil(o.TradeAmount) {
		var ret string
		return ret
	}
	return *o.TradeAmount
}

// GetTradeAmountOk returns a tuple with the TradeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TradeAmount) {
		return nil, false
	}
	return o.TradeAmount, true
}

// HasTradeAmount returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTradeAmount() bool {
	if o != nil && !IsNil(o.TradeAmount) {
		return true
	}

	return false
}

// SetTradeAmount gets a reference to the given string and assigns it to the TradeAmount field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTradeAmount(v string) {
	o.TradeAmount = &v
}

// GetTradeName returns the TradeName field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeName() string {
	if o == nil || IsNil(o.TradeName) {
		var ret string
		return ret
	}
	return *o.TradeName
}

// GetTradeNameOk returns a tuple with the TradeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeNameOk() (*string, bool) {
	if o == nil || IsNil(o.TradeName) {
		return nil, false
	}
	return o.TradeName, true
}

// HasTradeName returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTradeName() bool {
	if o != nil && !IsNil(o.TradeName) {
		return true
	}

	return false
}

// SetTradeName gets a reference to the given string and assigns it to the TradeName field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTradeName(v string) {
	o.TradeName = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeTime returns the TradeTime field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeTime() string {
	if o == nil || IsNil(o.TradeTime) {
		var ret string
		return ret
	}
	return *o.TradeTime
}

// GetTradeTimeOk returns a tuple with the TradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeTime) {
		return nil, false
	}
	return o.TradeTime, true
}

// HasTradeTime returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTradeTime() bool {
	if o != nil && !IsNil(o.TradeTime) {
		return true
	}

	return false
}

// SetTradeTime gets a reference to the given string and assigns it to the TradeTime field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTradeTime(v string) {
	o.TradeTime = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeType() string {
	if o == nil || IsNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetTradeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeType) {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasTradeType() bool {
	if o != nil && !IsNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *AlipayMarketingCardConsumeSyncModel) SetTradeType(v string) {
	o.TradeType = &v
}

// GetUseBenefitList returns the UseBenefitList field value if set, zero value otherwise.
func (o *AlipayMarketingCardConsumeSyncModel) GetUseBenefitList() []BenefitInfoDetail {
	if o == nil || IsNil(o.UseBenefitList) {
		var ret []BenefitInfoDetail
		return ret
	}
	return o.UseBenefitList
}

// GetUseBenefitListOk returns a tuple with the UseBenefitList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardConsumeSyncModel) GetUseBenefitListOk() ([]BenefitInfoDetail, bool) {
	if o == nil || IsNil(o.UseBenefitList) {
		return nil, false
	}
	return o.UseBenefitList, true
}

// HasUseBenefitList returns a boolean if a field has been set.
func (o *AlipayMarketingCardConsumeSyncModel) HasUseBenefitList() bool {
	if o != nil && !IsNil(o.UseBenefitList) {
		return true
	}

	return false
}

// SetUseBenefitList gets a reference to the given []BenefitInfoDetail and assigns it to the UseBenefitList field.
func (o *AlipayMarketingCardConsumeSyncModel) SetUseBenefitList(v []BenefitInfoDetail) {
	o.UseBenefitList = v
}

func (o AlipayMarketingCardConsumeSyncModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardConsumeSyncModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActPayAmount) {
		toSerialize["act_pay_amount"] = o.ActPayAmount
	}
	if !IsNil(o.CardInfo) {
		toSerialize["card_info"] = o.CardInfo
	}
	if !IsNil(o.GainBenefitList) {
		toSerialize["gain_benefit_list"] = o.GainBenefitList
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.ShopCode) {
		toSerialize["shop_code"] = o.ShopCode
	}
	if !IsNil(o.SwipeCertType) {
		toSerialize["swipe_cert_type"] = o.SwipeCertType
	}
	if !IsNil(o.TargetCardNo) {
		toSerialize["target_card_no"] = o.TargetCardNo
	}
	if !IsNil(o.TargetCardNoType) {
		toSerialize["target_card_no_type"] = o.TargetCardNoType
	}
	if !IsNil(o.TradeAmount) {
		toSerialize["trade_amount"] = o.TradeAmount
	}
	if !IsNil(o.TradeName) {
		toSerialize["trade_name"] = o.TradeName
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !IsNil(o.TradeTime) {
		toSerialize["trade_time"] = o.TradeTime
	}
	if !IsNil(o.TradeType) {
		toSerialize["trade_type"] = o.TradeType
	}
	if !IsNil(o.UseBenefitList) {
		toSerialize["use_benefit_list"] = o.UseBenefitList
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardConsumeSyncModel struct {
	value *AlipayMarketingCardConsumeSyncModel
	isSet bool
}

func (v NullableAlipayMarketingCardConsumeSyncModel) Get() *AlipayMarketingCardConsumeSyncModel {
	return v.value
}

func (v *NullableAlipayMarketingCardConsumeSyncModel) Set(val *AlipayMarketingCardConsumeSyncModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardConsumeSyncModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardConsumeSyncModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardConsumeSyncModel(val *AlipayMarketingCardConsumeSyncModel) *NullableAlipayMarketingCardConsumeSyncModel {
	return &NullableAlipayMarketingCardConsumeSyncModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardConsumeSyncModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardConsumeSyncModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
