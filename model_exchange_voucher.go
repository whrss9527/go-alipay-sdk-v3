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

// checks if the ExchangeVoucher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeVoucher{}

// ExchangeVoucher struct for ExchangeVoucher
type ExchangeVoucher struct {
	// 券的价值
	Amount *string `json:"amount,omitempty"`
	// 兑换券业务类型。  枚举值 团购券：GROUP_BUY_EXCHANGE_VOUCHER 代金券：FIX_EXCHANGE_VOUCHER 注意：兑换券通过<a href=\"https://opendocs.alipay.com/pre-open/repo-00tbta\">大促活动权益报名</a>能力推广至支付宝会场时，本参数必填。
	BizType *string `json:"biz_type,omitempty"`
	// 客服电话
	CustomerServiceMobile *string `json:"customer_service_mobile,omitempty"`
	// 客服链接
	CustomerServiceUrl *string `json:"customer_service_url,omitempty"`
	// 门槛金额。说明：该字段可不填，认为无门槛;
	FloorAmount *string `json:"floor_amount,omitempty"`
	// 是否支持优惠券过期后，自动退款给用户。 不填默认否，枚举值： true：是 false：否
	OverdueRefundable *bool `json:"overdue_refundable,omitempty"`
	// 收款账号。目前的结算规则是，每核销一笔优惠券，支付宝会打款到该收款账户。
	PayeePid *string `json:"payee_pid,omitempty"`
	// 购买的优惠券是否允许退款。不填默认否，枚举值： true：是 false：否
	Refundable *bool `json:"refundable,omitempty"`
	// 用户购买优惠券需要支付的金额
	SaleAmount *string `json:"sale_amount,omitempty"`
	// 解决类型。
	SettleType *string `json:"settle_type,omitempty"`
	// 领(购)券详情页链接，从支付宝公域跳转到服务商(商户)自定义领(购)券详情页。
	VoucherDetailUrl *string `json:"voucher_detail_url,omitempty"`
	// 对消费者展示的券(商品)名称。
	VoucherName *string `json:"voucher_name,omitempty"`
}

// NewExchangeVoucher instantiates a new ExchangeVoucher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeVoucher() *ExchangeVoucher {
	this := ExchangeVoucher{}
	return &this
}

// NewExchangeVoucherWithDefaults instantiates a new ExchangeVoucher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeVoucherWithDefaults() *ExchangeVoucher {
	this := ExchangeVoucher{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ExchangeVoucher) SetAmount(v string) {
	o.Amount = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *ExchangeVoucher) SetBizType(v string) {
	o.BizType = &v
}

// GetCustomerServiceMobile returns the CustomerServiceMobile field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetCustomerServiceMobile() string {
	if o == nil || IsNil(o.CustomerServiceMobile) {
		var ret string
		return ret
	}
	return *o.CustomerServiceMobile
}

// GetCustomerServiceMobileOk returns a tuple with the CustomerServiceMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetCustomerServiceMobileOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerServiceMobile) {
		return nil, false
	}
	return o.CustomerServiceMobile, true
}

// HasCustomerServiceMobile returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasCustomerServiceMobile() bool {
	if o != nil && !IsNil(o.CustomerServiceMobile) {
		return true
	}

	return false
}

// SetCustomerServiceMobile gets a reference to the given string and assigns it to the CustomerServiceMobile field.
func (o *ExchangeVoucher) SetCustomerServiceMobile(v string) {
	o.CustomerServiceMobile = &v
}

// GetCustomerServiceUrl returns the CustomerServiceUrl field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetCustomerServiceUrl() string {
	if o == nil || IsNil(o.CustomerServiceUrl) {
		var ret string
		return ret
	}
	return *o.CustomerServiceUrl
}

// GetCustomerServiceUrlOk returns a tuple with the CustomerServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetCustomerServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerServiceUrl) {
		return nil, false
	}
	return o.CustomerServiceUrl, true
}

// HasCustomerServiceUrl returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasCustomerServiceUrl() bool {
	if o != nil && !IsNil(o.CustomerServiceUrl) {
		return true
	}

	return false
}

// SetCustomerServiceUrl gets a reference to the given string and assigns it to the CustomerServiceUrl field.
func (o *ExchangeVoucher) SetCustomerServiceUrl(v string) {
	o.CustomerServiceUrl = &v
}

// GetFloorAmount returns the FloorAmount field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetFloorAmount() string {
	if o == nil || IsNil(o.FloorAmount) {
		var ret string
		return ret
	}
	return *o.FloorAmount
}

// GetFloorAmountOk returns a tuple with the FloorAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetFloorAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FloorAmount) {
		return nil, false
	}
	return o.FloorAmount, true
}

// HasFloorAmount returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasFloorAmount() bool {
	if o != nil && !IsNil(o.FloorAmount) {
		return true
	}

	return false
}

// SetFloorAmount gets a reference to the given string and assigns it to the FloorAmount field.
func (o *ExchangeVoucher) SetFloorAmount(v string) {
	o.FloorAmount = &v
}

// GetOverdueRefundable returns the OverdueRefundable field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetOverdueRefundable() bool {
	if o == nil || IsNil(o.OverdueRefundable) {
		var ret bool
		return ret
	}
	return *o.OverdueRefundable
}

// GetOverdueRefundableOk returns a tuple with the OverdueRefundable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetOverdueRefundableOk() (*bool, bool) {
	if o == nil || IsNil(o.OverdueRefundable) {
		return nil, false
	}
	return o.OverdueRefundable, true
}

// HasOverdueRefundable returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasOverdueRefundable() bool {
	if o != nil && !IsNil(o.OverdueRefundable) {
		return true
	}

	return false
}

// SetOverdueRefundable gets a reference to the given bool and assigns it to the OverdueRefundable field.
func (o *ExchangeVoucher) SetOverdueRefundable(v bool) {
	o.OverdueRefundable = &v
}

// GetPayeePid returns the PayeePid field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetPayeePid() string {
	if o == nil || IsNil(o.PayeePid) {
		var ret string
		return ret
	}
	return *o.PayeePid
}

// GetPayeePidOk returns a tuple with the PayeePid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetPayeePidOk() (*string, bool) {
	if o == nil || IsNil(o.PayeePid) {
		return nil, false
	}
	return o.PayeePid, true
}

// HasPayeePid returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasPayeePid() bool {
	if o != nil && !IsNil(o.PayeePid) {
		return true
	}

	return false
}

// SetPayeePid gets a reference to the given string and assigns it to the PayeePid field.
func (o *ExchangeVoucher) SetPayeePid(v string) {
	o.PayeePid = &v
}

// GetRefundable returns the Refundable field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetRefundable() bool {
	if o == nil || IsNil(o.Refundable) {
		var ret bool
		return ret
	}
	return *o.Refundable
}

// GetRefundableOk returns a tuple with the Refundable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetRefundableOk() (*bool, bool) {
	if o == nil || IsNil(o.Refundable) {
		return nil, false
	}
	return o.Refundable, true
}

// HasRefundable returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasRefundable() bool {
	if o != nil && !IsNil(o.Refundable) {
		return true
	}

	return false
}

// SetRefundable gets a reference to the given bool and assigns it to the Refundable field.
func (o *ExchangeVoucher) SetRefundable(v bool) {
	o.Refundable = &v
}

// GetSaleAmount returns the SaleAmount field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetSaleAmount() string {
	if o == nil || IsNil(o.SaleAmount) {
		var ret string
		return ret
	}
	return *o.SaleAmount
}

// GetSaleAmountOk returns a tuple with the SaleAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetSaleAmountOk() (*string, bool) {
	if o == nil || IsNil(o.SaleAmount) {
		return nil, false
	}
	return o.SaleAmount, true
}

// HasSaleAmount returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasSaleAmount() bool {
	if o != nil && !IsNil(o.SaleAmount) {
		return true
	}

	return false
}

// SetSaleAmount gets a reference to the given string and assigns it to the SaleAmount field.
func (o *ExchangeVoucher) SetSaleAmount(v string) {
	o.SaleAmount = &v
}

// GetSettleType returns the SettleType field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetSettleType() string {
	if o == nil || IsNil(o.SettleType) {
		var ret string
		return ret
	}
	return *o.SettleType
}

// GetSettleTypeOk returns a tuple with the SettleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetSettleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SettleType) {
		return nil, false
	}
	return o.SettleType, true
}

// HasSettleType returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasSettleType() bool {
	if o != nil && !IsNil(o.SettleType) {
		return true
	}

	return false
}

// SetSettleType gets a reference to the given string and assigns it to the SettleType field.
func (o *ExchangeVoucher) SetSettleType(v string) {
	o.SettleType = &v
}

// GetVoucherDetailUrl returns the VoucherDetailUrl field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetVoucherDetailUrl() string {
	if o == nil || IsNil(o.VoucherDetailUrl) {
		var ret string
		return ret
	}
	return *o.VoucherDetailUrl
}

// GetVoucherDetailUrlOk returns a tuple with the VoucherDetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetVoucherDetailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherDetailUrl) {
		return nil, false
	}
	return o.VoucherDetailUrl, true
}

// HasVoucherDetailUrl returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasVoucherDetailUrl() bool {
	if o != nil && !IsNil(o.VoucherDetailUrl) {
		return true
	}

	return false
}

// SetVoucherDetailUrl gets a reference to the given string and assigns it to the VoucherDetailUrl field.
func (o *ExchangeVoucher) SetVoucherDetailUrl(v string) {
	o.VoucherDetailUrl = &v
}

// GetVoucherName returns the VoucherName field value if set, zero value otherwise.
func (o *ExchangeVoucher) GetVoucherName() string {
	if o == nil || IsNil(o.VoucherName) {
		var ret string
		return ret
	}
	return *o.VoucherName
}

// GetVoucherNameOk returns a tuple with the VoucherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucher) GetVoucherNameOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherName) {
		return nil, false
	}
	return o.VoucherName, true
}

// HasVoucherName returns a boolean if a field has been set.
func (o *ExchangeVoucher) HasVoucherName() bool {
	if o != nil && !IsNil(o.VoucherName) {
		return true
	}

	return false
}

// SetVoucherName gets a reference to the given string and assigns it to the VoucherName field.
func (o *ExchangeVoucher) SetVoucherName(v string) {
	o.VoucherName = &v
}

func (o ExchangeVoucher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeVoucher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.CustomerServiceMobile) {
		toSerialize["customer_service_mobile"] = o.CustomerServiceMobile
	}
	if !IsNil(o.CustomerServiceUrl) {
		toSerialize["customer_service_url"] = o.CustomerServiceUrl
	}
	if !IsNil(o.FloorAmount) {
		toSerialize["floor_amount"] = o.FloorAmount
	}
	if !IsNil(o.OverdueRefundable) {
		toSerialize["overdue_refundable"] = o.OverdueRefundable
	}
	if !IsNil(o.PayeePid) {
		toSerialize["payee_pid"] = o.PayeePid
	}
	if !IsNil(o.Refundable) {
		toSerialize["refundable"] = o.Refundable
	}
	if !IsNil(o.SaleAmount) {
		toSerialize["sale_amount"] = o.SaleAmount
	}
	if !IsNil(o.SettleType) {
		toSerialize["settle_type"] = o.SettleType
	}
	if !IsNil(o.VoucherDetailUrl) {
		toSerialize["voucher_detail_url"] = o.VoucherDetailUrl
	}
	if !IsNil(o.VoucherName) {
		toSerialize["voucher_name"] = o.VoucherName
	}
	return toSerialize, nil
}

type NullableExchangeVoucher struct {
	value *ExchangeVoucher
	isSet bool
}

func (v NullableExchangeVoucher) Get() *ExchangeVoucher {
	return v.value
}

func (v *NullableExchangeVoucher) Set(val *ExchangeVoucher) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeVoucher) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeVoucher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeVoucher(val *ExchangeVoucher) *NullableExchangeVoucher {
	return &NullableExchangeVoucher{value: val, isSet: true}
}

func (v NullableExchangeVoucher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeVoucher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


