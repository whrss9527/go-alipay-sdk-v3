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

// checks if the WaybillInvoiceQueryIstd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WaybillInvoiceQueryIstd{}

// WaybillInvoiceQueryIstd struct for WaybillInvoiceQueryIstd
type WaybillInvoiceQueryIstd struct {
	// 失败原因，需要传单个即时配送运单的驳回原因
	Reason *string `json:"reason,omitempty"`
	// 商家门店编号
	ShopNo *string `json:"shop_no,omitempty"`
	// 即时配送运单金额，waybill_invoce_status为1的情况下不能为空
	WaybillAmount *string `json:"waybill_amount,omitempty"`
	// 明细的运单开票状态，1：开票成功 2：不可开票 3：可开票；整体开票状态为0的情况下，无开票明细；整体开票状态为1，明细开票状态全部是1；整体开票状态为2，明细开票状态全部为2
	WaybillInvoiceStatus *int32 `json:"waybill_invoice_status,omitempty"`
	// 即时配送运单编号
	WaybillNo *string `json:"waybill_no,omitempty"`
}

// NewWaybillInvoiceQueryIstd instantiates a new WaybillInvoiceQueryIstd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaybillInvoiceQueryIstd() *WaybillInvoiceQueryIstd {
	this := WaybillInvoiceQueryIstd{}
	return &this
}

// NewWaybillInvoiceQueryIstdWithDefaults instantiates a new WaybillInvoiceQueryIstd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaybillInvoiceQueryIstdWithDefaults() *WaybillInvoiceQueryIstd {
	this := WaybillInvoiceQueryIstd{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *WaybillInvoiceQueryIstd) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaybillInvoiceQueryIstd) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *WaybillInvoiceQueryIstd) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *WaybillInvoiceQueryIstd) SetReason(v string) {
	o.Reason = &v
}

// GetShopNo returns the ShopNo field value if set, zero value otherwise.
func (o *WaybillInvoiceQueryIstd) GetShopNo() string {
	if o == nil || IsNil(o.ShopNo) {
		var ret string
		return ret
	}
	return *o.ShopNo
}

// GetShopNoOk returns a tuple with the ShopNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaybillInvoiceQueryIstd) GetShopNoOk() (*string, bool) {
	if o == nil || IsNil(o.ShopNo) {
		return nil, false
	}
	return o.ShopNo, true
}

// HasShopNo returns a boolean if a field has been set.
func (o *WaybillInvoiceQueryIstd) HasShopNo() bool {
	if o != nil && !IsNil(o.ShopNo) {
		return true
	}

	return false
}

// SetShopNo gets a reference to the given string and assigns it to the ShopNo field.
func (o *WaybillInvoiceQueryIstd) SetShopNo(v string) {
	o.ShopNo = &v
}

// GetWaybillAmount returns the WaybillAmount field value if set, zero value otherwise.
func (o *WaybillInvoiceQueryIstd) GetWaybillAmount() string {
	if o == nil || IsNil(o.WaybillAmount) {
		var ret string
		return ret
	}
	return *o.WaybillAmount
}

// GetWaybillAmountOk returns a tuple with the WaybillAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaybillInvoiceQueryIstd) GetWaybillAmountOk() (*string, bool) {
	if o == nil || IsNil(o.WaybillAmount) {
		return nil, false
	}
	return o.WaybillAmount, true
}

// HasWaybillAmount returns a boolean if a field has been set.
func (o *WaybillInvoiceQueryIstd) HasWaybillAmount() bool {
	if o != nil && !IsNil(o.WaybillAmount) {
		return true
	}

	return false
}

// SetWaybillAmount gets a reference to the given string and assigns it to the WaybillAmount field.
func (o *WaybillInvoiceQueryIstd) SetWaybillAmount(v string) {
	o.WaybillAmount = &v
}

// GetWaybillInvoiceStatus returns the WaybillInvoiceStatus field value if set, zero value otherwise.
func (o *WaybillInvoiceQueryIstd) GetWaybillInvoiceStatus() int32 {
	if o == nil || IsNil(o.WaybillInvoiceStatus) {
		var ret int32
		return ret
	}
	return *o.WaybillInvoiceStatus
}

// GetWaybillInvoiceStatusOk returns a tuple with the WaybillInvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaybillInvoiceQueryIstd) GetWaybillInvoiceStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.WaybillInvoiceStatus) {
		return nil, false
	}
	return o.WaybillInvoiceStatus, true
}

// HasWaybillInvoiceStatus returns a boolean if a field has been set.
func (o *WaybillInvoiceQueryIstd) HasWaybillInvoiceStatus() bool {
	if o != nil && !IsNil(o.WaybillInvoiceStatus) {
		return true
	}

	return false
}

// SetWaybillInvoiceStatus gets a reference to the given int32 and assigns it to the WaybillInvoiceStatus field.
func (o *WaybillInvoiceQueryIstd) SetWaybillInvoiceStatus(v int32) {
	o.WaybillInvoiceStatus = &v
}

// GetWaybillNo returns the WaybillNo field value if set, zero value otherwise.
func (o *WaybillInvoiceQueryIstd) GetWaybillNo() string {
	if o == nil || IsNil(o.WaybillNo) {
		var ret string
		return ret
	}
	return *o.WaybillNo
}

// GetWaybillNoOk returns a tuple with the WaybillNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaybillInvoiceQueryIstd) GetWaybillNoOk() (*string, bool) {
	if o == nil || IsNil(o.WaybillNo) {
		return nil, false
	}
	return o.WaybillNo, true
}

// HasWaybillNo returns a boolean if a field has been set.
func (o *WaybillInvoiceQueryIstd) HasWaybillNo() bool {
	if o != nil && !IsNil(o.WaybillNo) {
		return true
	}

	return false
}

// SetWaybillNo gets a reference to the given string and assigns it to the WaybillNo field.
func (o *WaybillInvoiceQueryIstd) SetWaybillNo(v string) {
	o.WaybillNo = &v
}

func (o WaybillInvoiceQueryIstd) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WaybillInvoiceQueryIstd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ShopNo) {
		toSerialize["shop_no"] = o.ShopNo
	}
	if !IsNil(o.WaybillAmount) {
		toSerialize["waybill_amount"] = o.WaybillAmount
	}
	if !IsNil(o.WaybillInvoiceStatus) {
		toSerialize["waybill_invoice_status"] = o.WaybillInvoiceStatus
	}
	if !IsNil(o.WaybillNo) {
		toSerialize["waybill_no"] = o.WaybillNo
	}
	return toSerialize, nil
}

type NullableWaybillInvoiceQueryIstd struct {
	value *WaybillInvoiceQueryIstd
	isSet bool
}

func (v NullableWaybillInvoiceQueryIstd) Get() *WaybillInvoiceQueryIstd {
	return v.value
}

func (v *NullableWaybillInvoiceQueryIstd) Set(val *WaybillInvoiceQueryIstd) {
	v.value = val
	v.isSet = true
}

func (v NullableWaybillInvoiceQueryIstd) IsSet() bool {
	return v.isSet
}

func (v *NullableWaybillInvoiceQueryIstd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaybillInvoiceQueryIstd(val *WaybillInvoiceQueryIstd) *NullableWaybillInvoiceQueryIstd {
	return &NullableWaybillInvoiceQueryIstd{value: val, isSet: true}
}

func (v NullableWaybillInvoiceQueryIstd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaybillInvoiceQueryIstd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
