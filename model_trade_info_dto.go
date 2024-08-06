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

// checks if the TradeInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TradeInfoDTO{}

// TradeInfoDTO struct for TradeInfoDTO
type TradeInfoDTO struct {
	// 买家ID
	BuyerId *string `json:"buyer_id,omitempty"`
	// 交易创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 外部平台订单号
	PlatformOrderId *string `json:"platform_order_id,omitempty"`
	// 订单总金额
	TotalAmount *string `json:"total_amount,omitempty"`
	// 订单总金额
	TradeAmount *string `json:"trade_amount,omitempty"`
	// 资金单明细
	TradeFundBillList []TradeFundBillDetail `json:"trade_fund_bill_list,omitempty"`
	// 交易单号
	TradeNo *string `json:"trade_no,omitempty"`
	// 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	TradeStatus *string `json:"trade_status,omitempty"`
}

// NewTradeInfoDTO instantiates a new TradeInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeInfoDTO() *TradeInfoDTO {
	this := TradeInfoDTO{}
	return &this
}

// NewTradeInfoDTOWithDefaults instantiates a new TradeInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeInfoDTOWithDefaults() *TradeInfoDTO {
	this := TradeInfoDTO{}
	return &this
}

// GetBuyerId returns the BuyerId field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetBuyerId() string {
	if o == nil || IsNil(o.BuyerId) {
		var ret string
		return ret
	}
	return *o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetBuyerIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerId) {
		return nil, false
	}
	return o.BuyerId, true
}

// HasBuyerId returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasBuyerId() bool {
	if o != nil && !IsNil(o.BuyerId) {
		return true
	}

	return false
}

// SetBuyerId gets a reference to the given string and assigns it to the BuyerId field.
func (o *TradeInfoDTO) SetBuyerId(v string) {
	o.BuyerId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *TradeInfoDTO) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetPlatformOrderId returns the PlatformOrderId field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetPlatformOrderId() string {
	if o == nil || IsNil(o.PlatformOrderId) {
		var ret string
		return ret
	}
	return *o.PlatformOrderId
}

// GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetPlatformOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformOrderId) {
		return nil, false
	}
	return o.PlatformOrderId, true
}

// HasPlatformOrderId returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasPlatformOrderId() bool {
	if o != nil && !IsNil(o.PlatformOrderId) {
		return true
	}

	return false
}

// SetPlatformOrderId gets a reference to the given string and assigns it to the PlatformOrderId field.
func (o *TradeInfoDTO) SetPlatformOrderId(v string) {
	o.PlatformOrderId = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetTotalAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *TradeInfoDTO) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetTradeAmount returns the TradeAmount field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetTradeAmount() string {
	if o == nil || IsNil(o.TradeAmount) {
		var ret string
		return ret
	}
	return *o.TradeAmount
}

// GetTradeAmountOk returns a tuple with the TradeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetTradeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TradeAmount) {
		return nil, false
	}
	return o.TradeAmount, true
}

// HasTradeAmount returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasTradeAmount() bool {
	if o != nil && !IsNil(o.TradeAmount) {
		return true
	}

	return false
}

// SetTradeAmount gets a reference to the given string and assigns it to the TradeAmount field.
func (o *TradeInfoDTO) SetTradeAmount(v string) {
	o.TradeAmount = &v
}

// GetTradeFundBillList returns the TradeFundBillList field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetTradeFundBillList() []TradeFundBillDetail {
	if o == nil || IsNil(o.TradeFundBillList) {
		var ret []TradeFundBillDetail
		return ret
	}
	return o.TradeFundBillList
}

// GetTradeFundBillListOk returns a tuple with the TradeFundBillList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetTradeFundBillListOk() ([]TradeFundBillDetail, bool) {
	if o == nil || IsNil(o.TradeFundBillList) {
		return nil, false
	}
	return o.TradeFundBillList, true
}

// HasTradeFundBillList returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasTradeFundBillList() bool {
	if o != nil && !IsNil(o.TradeFundBillList) {
		return true
	}

	return false
}

// SetTradeFundBillList gets a reference to the given []TradeFundBillDetail and assigns it to the TradeFundBillList field.
func (o *TradeInfoDTO) SetTradeFundBillList(v []TradeFundBillDetail) {
	o.TradeFundBillList = v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *TradeInfoDTO) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetTradeStatus returns the TradeStatus field value if set, zero value otherwise.
func (o *TradeInfoDTO) GetTradeStatus() string {
	if o == nil || IsNil(o.TradeStatus) {
		var ret string
		return ret
	}
	return *o.TradeStatus
}

// GetTradeStatusOk returns a tuple with the TradeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeInfoDTO) GetTradeStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TradeStatus) {
		return nil, false
	}
	return o.TradeStatus, true
}

// HasTradeStatus returns a boolean if a field has been set.
func (o *TradeInfoDTO) HasTradeStatus() bool {
	if o != nil && !IsNil(o.TradeStatus) {
		return true
	}

	return false
}

// SetTradeStatus gets a reference to the given string and assigns it to the TradeStatus field.
func (o *TradeInfoDTO) SetTradeStatus(v string) {
	o.TradeStatus = &v
}

func (o TradeInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyerId) {
		toSerialize["buyer_id"] = o.BuyerId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.PlatformOrderId) {
		toSerialize["platform_order_id"] = o.PlatformOrderId
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if !IsNil(o.TradeAmount) {
		toSerialize["trade_amount"] = o.TradeAmount
	}
	if !IsNil(o.TradeFundBillList) {
		toSerialize["trade_fund_bill_list"] = o.TradeFundBillList
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !IsNil(o.TradeStatus) {
		toSerialize["trade_status"] = o.TradeStatus
	}
	return toSerialize, nil
}

type NullableTradeInfoDTO struct {
	value *TradeInfoDTO
	isSet bool
}

func (v NullableTradeInfoDTO) Get() *TradeInfoDTO {
	return v.value
}

func (v *NullableTradeInfoDTO) Set(val *TradeInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeInfoDTO(val *TradeInfoDTO) *NullableTradeInfoDTO {
	return &NullableTradeInfoDTO{value: val, isSet: true}
}

func (v NullableTradeInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


