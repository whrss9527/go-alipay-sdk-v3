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

// checks if the AlipayTradeFastpayRefundQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeFastpayRefundQueryModel{}

// AlipayTradeFastpayRefundQueryModel struct for AlipayTradeFastpayRefundQueryModel
type AlipayTradeFastpayRefundQueryModel struct {
	// 银行间联模式下有用，其它场景请不要使用；  双联通过该参数指定需要查询的交易所属收单机构的pid;
	OrgPid *string `json:"org_pid,omitempty"`
	// 退款请求号。 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的商户订单号。
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// 商户订单号。 订单支付时传入的商户订单号,和支付宝交易号不能同时为空。 trade_no,out_trade_no如果同时存在优先取trade_no
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 查询选项，商户通过上送该参数来定制同步需要额外返回的信息字段，数组格式。枚举支持： refund_detail_item_list：本次退款使用的资金渠道； gmt_refund_pay：退款执行成功的时间； deposit_back_info：银行卡冲退信息；
	QueryOptions []string `json:"query_options,omitempty"`
	// 支付宝交易号。 和商户订单号不能同时为空
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayTradeFastpayRefundQueryModel instantiates a new AlipayTradeFastpayRefundQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeFastpayRefundQueryModel() *AlipayTradeFastpayRefundQueryModel {
	this := AlipayTradeFastpayRefundQueryModel{}
	return &this
}

// NewAlipayTradeFastpayRefundQueryModelWithDefaults instantiates a new AlipayTradeFastpayRefundQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeFastpayRefundQueryModelWithDefaults() *AlipayTradeFastpayRefundQueryModel {
	this := AlipayTradeFastpayRefundQueryModel{}
	return &this
}

// GetOrgPid returns the OrgPid field value if set, zero value otherwise.
func (o *AlipayTradeFastpayRefundQueryModel) GetOrgPid() string {
	if o == nil || IsNil(o.OrgPid) {
		var ret string
		return ret
	}
	return *o.OrgPid
}

// GetOrgPidOk returns a tuple with the OrgPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeFastpayRefundQueryModel) GetOrgPidOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPid) {
		return nil, false
	}
	return o.OrgPid, true
}

// HasOrgPid returns a boolean if a field has been set.
func (o *AlipayTradeFastpayRefundQueryModel) HasOrgPid() bool {
	if o != nil && !IsNil(o.OrgPid) {
		return true
	}

	return false
}

// SetOrgPid gets a reference to the given string and assigns it to the OrgPid field.
func (o *AlipayTradeFastpayRefundQueryModel) SetOrgPid(v string) {
	o.OrgPid = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayTradeFastpayRefundQueryModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeFastpayRefundQueryModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayTradeFastpayRefundQueryModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayTradeFastpayRefundQueryModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayTradeFastpayRefundQueryModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeFastpayRefundQueryModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeFastpayRefundQueryModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayTradeFastpayRefundQueryModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetQueryOptions returns the QueryOptions field value if set, zero value otherwise.
func (o *AlipayTradeFastpayRefundQueryModel) GetQueryOptions() []string {
	if o == nil || IsNil(o.QueryOptions) {
		var ret []string
		return ret
	}
	return o.QueryOptions
}

// GetQueryOptionsOk returns a tuple with the QueryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeFastpayRefundQueryModel) GetQueryOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.QueryOptions) {
		return nil, false
	}
	return o.QueryOptions, true
}

// HasQueryOptions returns a boolean if a field has been set.
func (o *AlipayTradeFastpayRefundQueryModel) HasQueryOptions() bool {
	if o != nil && !IsNil(o.QueryOptions) {
		return true
	}

	return false
}

// SetQueryOptions gets a reference to the given []string and assigns it to the QueryOptions field.
func (o *AlipayTradeFastpayRefundQueryModel) SetQueryOptions(v []string) {
	o.QueryOptions = v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeFastpayRefundQueryModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeFastpayRefundQueryModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeFastpayRefundQueryModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeFastpayRefundQueryModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayTradeFastpayRefundQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeFastpayRefundQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgPid) {
		toSerialize["org_pid"] = o.OrgPid
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.QueryOptions) {
		toSerialize["query_options"] = o.QueryOptions
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayTradeFastpayRefundQueryModel struct {
	value *AlipayTradeFastpayRefundQueryModel
	isSet bool
}

func (v NullableAlipayTradeFastpayRefundQueryModel) Get() *AlipayTradeFastpayRefundQueryModel {
	return v.value
}

func (v *NullableAlipayTradeFastpayRefundQueryModel) Set(val *AlipayTradeFastpayRefundQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeFastpayRefundQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeFastpayRefundQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeFastpayRefundQueryModel(val *AlipayTradeFastpayRefundQueryModel) *NullableAlipayTradeFastpayRefundQueryModel {
	return &NullableAlipayTradeFastpayRefundQueryModel{value: val, isSet: true}
}

func (v NullableAlipayTradeFastpayRefundQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeFastpayRefundQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

