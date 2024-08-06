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

// checks if the AlipayOfflineMarketApplyorderBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOfflineMarketApplyorderBatchqueryResponseModel{}

// AlipayOfflineMarketApplyorderBatchqueryResponseModel struct for AlipayOfflineMarketApplyorderBatchqueryResponseModel
type AlipayOfflineMarketApplyorderBatchqueryResponseModel struct {
	// 支付宝操作流水信息列表
	BizOrderInfos []BizOrderQueryResponse `json:"biz_order_infos,omitempty"`
	// 当前页码
	CurrentPageNo *int32 `json:"current_page_no,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总记录数
	TotalItems *int32 `json:"total_items,omitempty"`
	// 总页码数目
	TotalPageNo *int32 `json:"total_page_no,omitempty"`
}

// NewAlipayOfflineMarketApplyorderBatchqueryResponseModel instantiates a new AlipayOfflineMarketApplyorderBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOfflineMarketApplyorderBatchqueryResponseModel() *AlipayOfflineMarketApplyorderBatchqueryResponseModel {
	this := AlipayOfflineMarketApplyorderBatchqueryResponseModel{}
	return &this
}

// NewAlipayOfflineMarketApplyorderBatchqueryResponseModelWithDefaults instantiates a new AlipayOfflineMarketApplyorderBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOfflineMarketApplyorderBatchqueryResponseModelWithDefaults() *AlipayOfflineMarketApplyorderBatchqueryResponseModel {
	this := AlipayOfflineMarketApplyorderBatchqueryResponseModel{}
	return &this
}

// GetBizOrderInfos returns the BizOrderInfos field value if set, zero value otherwise.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetBizOrderInfos() []BizOrderQueryResponse {
	if o == nil || IsNil(o.BizOrderInfos) {
		var ret []BizOrderQueryResponse
		return ret
	}
	return o.BizOrderInfos
}

// GetBizOrderInfosOk returns a tuple with the BizOrderInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetBizOrderInfosOk() ([]BizOrderQueryResponse, bool) {
	if o == nil || IsNil(o.BizOrderInfos) {
		return nil, false
	}
	return o.BizOrderInfos, true
}

// HasBizOrderInfos returns a boolean if a field has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) HasBizOrderInfos() bool {
	if o != nil && !IsNil(o.BizOrderInfos) {
		return true
	}

	return false
}

// SetBizOrderInfos gets a reference to the given []BizOrderQueryResponse and assigns it to the BizOrderInfos field.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) SetBizOrderInfos(v []BizOrderQueryResponse) {
	o.BizOrderInfos = v
}

// GetCurrentPageNo returns the CurrentPageNo field value if set, zero value otherwise.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetCurrentPageNo() int32 {
	if o == nil || IsNil(o.CurrentPageNo) {
		var ret int32
		return ret
	}
	return *o.CurrentPageNo
}

// GetCurrentPageNoOk returns a tuple with the CurrentPageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetCurrentPageNoOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPageNo) {
		return nil, false
	}
	return o.CurrentPageNo, true
}

// HasCurrentPageNo returns a boolean if a field has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) HasCurrentPageNo() bool {
	if o != nil && !IsNil(o.CurrentPageNo) {
		return true
	}

	return false
}

// SetCurrentPageNo gets a reference to the given int32 and assigns it to the CurrentPageNo field.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) SetCurrentPageNo(v int32) {
	o.CurrentPageNo = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetTotalPageNo returns the TotalPageNo field value if set, zero value otherwise.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetTotalPageNo() int32 {
	if o == nil || IsNil(o.TotalPageNo) {
		var ret int32
		return ret
	}
	return *o.TotalPageNo
}

// GetTotalPageNoOk returns a tuple with the TotalPageNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) GetTotalPageNoOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPageNo) {
		return nil, false
	}
	return o.TotalPageNo, true
}

// HasTotalPageNo returns a boolean if a field has been set.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) HasTotalPageNo() bool {
	if o != nil && !IsNil(o.TotalPageNo) {
		return true
	}

	return false
}

// SetTotalPageNo gets a reference to the given int32 and assigns it to the TotalPageNo field.
func (o *AlipayOfflineMarketApplyorderBatchqueryResponseModel) SetTotalPageNo(v int32) {
	o.TotalPageNo = &v
}

func (o AlipayOfflineMarketApplyorderBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOfflineMarketApplyorderBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizOrderInfos) {
		toSerialize["biz_order_infos"] = o.BizOrderInfos
	}
	if !IsNil(o.CurrentPageNo) {
		toSerialize["current_page_no"] = o.CurrentPageNo
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.TotalPageNo) {
		toSerialize["total_page_no"] = o.TotalPageNo
	}
	return toSerialize, nil
}

type NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel struct {
	value *AlipayOfflineMarketApplyorderBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) Get() *AlipayOfflineMarketApplyorderBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) Set(val *AlipayOfflineMarketApplyorderBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOfflineMarketApplyorderBatchqueryResponseModel(val *AlipayOfflineMarketApplyorderBatchqueryResponseModel) *NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel {
	return &NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOfflineMarketApplyorderBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
