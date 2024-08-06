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

// checks if the AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel{}

// AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel struct for AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel
type AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel struct {
	// 共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 传入操作人员标识
	OperatorId *string     `json:"operator_id,omitempty"`
	ShopInfo   *EcShopInfo `json:"shop_info,omitempty"`
	// 交易流水号
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel instantiates a new AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel {
	this := AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseMerchantrelationCreateModelWithDefaults() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel {
	this := AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetOperatorId() string {
	if o == nil || IsNil(o.OperatorId) {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetOperatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorId) {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) HasOperatorId() bool {
	if o != nil && !IsNil(o.OperatorId) {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetShopInfo returns the ShopInfo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetShopInfo() EcShopInfo {
	if o == nil || IsNil(o.ShopInfo) {
		var ret EcShopInfo
		return ret
	}
	return *o.ShopInfo
}

// GetShopInfoOk returns a tuple with the ShopInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetShopInfoOk() (*EcShopInfo, bool) {
	if o == nil || IsNil(o.ShopInfo) {
		return nil, false
	}
	return o.ShopInfo, true
}

// HasShopInfo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) HasShopInfo() bool {
	if o != nil && !IsNil(o.ShopInfo) {
		return true
	}

	return false
}

// SetShopInfo gets a reference to the given EcShopInfo and assigns it to the ShopInfo field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) SetShopInfo(v EcShopInfo) {
	o.ShopInfo = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.OperatorId) {
		toSerialize["operator_id"] = o.OperatorId
	}
	if !IsNil(o.ShopInfo) {
		toSerialize["shop_info"] = o.ShopInfo
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel struct {
	value *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) Get() *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) Set(val *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel(val *AlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel {
	return &NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseMerchantrelationCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
