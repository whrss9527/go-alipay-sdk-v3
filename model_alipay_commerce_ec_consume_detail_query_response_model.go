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

// checks if the AlipayCommerceEcConsumeDetailQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcConsumeDetailQueryResponseModel{}

// AlipayCommerceEcConsumeDetailQueryResponseModel struct for AlipayCommerceEcConsumeDetailQueryResponseModel
type AlipayCommerceEcConsumeDetailQueryResponseModel struct {
	ConsumeInfo      *EcConsumeInfo `json:"consume_info,omitempty"`
	RelatedOrderInfo *EcOrderInfo   `json:"related_order_info,omitempty"`
	// 关联账单列表
	RelatedRefundList []EcConsumeInfo `json:"related_refund_list,omitempty"`
	// 关联凭证详情列表
	RelatedVoucherList []EcVoucherInfo `json:"related_voucher_list,omitempty"`
}

// NewAlipayCommerceEcConsumeDetailQueryResponseModel instantiates a new AlipayCommerceEcConsumeDetailQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcConsumeDetailQueryResponseModel() *AlipayCommerceEcConsumeDetailQueryResponseModel {
	this := AlipayCommerceEcConsumeDetailQueryResponseModel{}
	return &this
}

// NewAlipayCommerceEcConsumeDetailQueryResponseModelWithDefaults instantiates a new AlipayCommerceEcConsumeDetailQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcConsumeDetailQueryResponseModelWithDefaults() *AlipayCommerceEcConsumeDetailQueryResponseModel {
	this := AlipayCommerceEcConsumeDetailQueryResponseModel{}
	return &this
}

// GetConsumeInfo returns the ConsumeInfo field value if set, zero value otherwise.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetConsumeInfo() EcConsumeInfo {
	if o == nil || IsNil(o.ConsumeInfo) {
		var ret EcConsumeInfo
		return ret
	}
	return *o.ConsumeInfo
}

// GetConsumeInfoOk returns a tuple with the ConsumeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetConsumeInfoOk() (*EcConsumeInfo, bool) {
	if o == nil || IsNil(o.ConsumeInfo) {
		return nil, false
	}
	return o.ConsumeInfo, true
}

// HasConsumeInfo returns a boolean if a field has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) HasConsumeInfo() bool {
	if o != nil && !IsNil(o.ConsumeInfo) {
		return true
	}

	return false
}

// SetConsumeInfo gets a reference to the given EcConsumeInfo and assigns it to the ConsumeInfo field.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) SetConsumeInfo(v EcConsumeInfo) {
	o.ConsumeInfo = &v
}

// GetRelatedOrderInfo returns the RelatedOrderInfo field value if set, zero value otherwise.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedOrderInfo() EcOrderInfo {
	if o == nil || IsNil(o.RelatedOrderInfo) {
		var ret EcOrderInfo
		return ret
	}
	return *o.RelatedOrderInfo
}

// GetRelatedOrderInfoOk returns a tuple with the RelatedOrderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedOrderInfoOk() (*EcOrderInfo, bool) {
	if o == nil || IsNil(o.RelatedOrderInfo) {
		return nil, false
	}
	return o.RelatedOrderInfo, true
}

// HasRelatedOrderInfo returns a boolean if a field has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) HasRelatedOrderInfo() bool {
	if o != nil && !IsNil(o.RelatedOrderInfo) {
		return true
	}

	return false
}

// SetRelatedOrderInfo gets a reference to the given EcOrderInfo and assigns it to the RelatedOrderInfo field.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) SetRelatedOrderInfo(v EcOrderInfo) {
	o.RelatedOrderInfo = &v
}

// GetRelatedRefundList returns the RelatedRefundList field value if set, zero value otherwise.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedRefundList() []EcConsumeInfo {
	if o == nil || IsNil(o.RelatedRefundList) {
		var ret []EcConsumeInfo
		return ret
	}
	return o.RelatedRefundList
}

// GetRelatedRefundListOk returns a tuple with the RelatedRefundList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedRefundListOk() ([]EcConsumeInfo, bool) {
	if o == nil || IsNil(o.RelatedRefundList) {
		return nil, false
	}
	return o.RelatedRefundList, true
}

// HasRelatedRefundList returns a boolean if a field has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) HasRelatedRefundList() bool {
	if o != nil && !IsNil(o.RelatedRefundList) {
		return true
	}

	return false
}

// SetRelatedRefundList gets a reference to the given []EcConsumeInfo and assigns it to the RelatedRefundList field.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) SetRelatedRefundList(v []EcConsumeInfo) {
	o.RelatedRefundList = v
}

// GetRelatedVoucherList returns the RelatedVoucherList field value if set, zero value otherwise.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedVoucherList() []EcVoucherInfo {
	if o == nil || IsNil(o.RelatedVoucherList) {
		var ret []EcVoucherInfo
		return ret
	}
	return o.RelatedVoucherList
}

// GetRelatedVoucherListOk returns a tuple with the RelatedVoucherList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) GetRelatedVoucherListOk() ([]EcVoucherInfo, bool) {
	if o == nil || IsNil(o.RelatedVoucherList) {
		return nil, false
	}
	return o.RelatedVoucherList, true
}

// HasRelatedVoucherList returns a boolean if a field has been set.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) HasRelatedVoucherList() bool {
	if o != nil && !IsNil(o.RelatedVoucherList) {
		return true
	}

	return false
}

// SetRelatedVoucherList gets a reference to the given []EcVoucherInfo and assigns it to the RelatedVoucherList field.
func (o *AlipayCommerceEcConsumeDetailQueryResponseModel) SetRelatedVoucherList(v []EcVoucherInfo) {
	o.RelatedVoucherList = v
}

func (o AlipayCommerceEcConsumeDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcConsumeDetailQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumeInfo) {
		toSerialize["consume_info"] = o.ConsumeInfo
	}
	if !IsNil(o.RelatedOrderInfo) {
		toSerialize["related_order_info"] = o.RelatedOrderInfo
	}
	if !IsNil(o.RelatedRefundList) {
		toSerialize["related_refund_list"] = o.RelatedRefundList
	}
	if !IsNil(o.RelatedVoucherList) {
		toSerialize["related_voucher_list"] = o.RelatedVoucherList
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcConsumeDetailQueryResponseModel struct {
	value *AlipayCommerceEcConsumeDetailQueryResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcConsumeDetailQueryResponseModel) Get() *AlipayCommerceEcConsumeDetailQueryResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcConsumeDetailQueryResponseModel) Set(val *AlipayCommerceEcConsumeDetailQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcConsumeDetailQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcConsumeDetailQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcConsumeDetailQueryResponseModel(val *AlipayCommerceEcConsumeDetailQueryResponseModel) *NullableAlipayCommerceEcConsumeDetailQueryResponseModel {
	return &NullableAlipayCommerceEcConsumeDetailQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcConsumeDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcConsumeDetailQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
