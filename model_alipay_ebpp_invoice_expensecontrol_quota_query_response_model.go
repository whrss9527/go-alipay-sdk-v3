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

// checks if the AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel{}

// AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel struct for AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel
type AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel struct {
	// 额度列表
	ExpenseQuotaInfoList []ExpenseQuotaInfo `json:"expense_quota_info_list,omitempty"`
	// 当前页数
	PageNum *int32 `json:"page_num,omitempty"`
	// 当前记录数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总页数
	TotalPageCount *int32 `json:"total_page_count,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel instantiates a new AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel() *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModelWithDefaults() *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel {
	this := AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel{}
	return &this
}

// GetExpenseQuotaInfoList returns the ExpenseQuotaInfoList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetExpenseQuotaInfoList() []ExpenseQuotaInfo {
	if o == nil || IsNil(o.ExpenseQuotaInfoList) {
		var ret []ExpenseQuotaInfo
		return ret
	}
	return o.ExpenseQuotaInfoList
}

// GetExpenseQuotaInfoListOk returns a tuple with the ExpenseQuotaInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetExpenseQuotaInfoListOk() ([]ExpenseQuotaInfo, bool) {
	if o == nil || IsNil(o.ExpenseQuotaInfoList) {
		return nil, false
	}
	return o.ExpenseQuotaInfoList, true
}

// HasExpenseQuotaInfoList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) HasExpenseQuotaInfoList() bool {
	if o != nil && !IsNil(o.ExpenseQuotaInfoList) {
		return true
	}

	return false
}

// SetExpenseQuotaInfoList gets a reference to the given []ExpenseQuotaInfo and assigns it to the ExpenseQuotaInfoList field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) SetExpenseQuotaInfoList(v []ExpenseQuotaInfo) {
	o.ExpenseQuotaInfoList = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalPageCount returns the TotalPageCount field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetTotalPageCount() int32 {
	if o == nil || IsNil(o.TotalPageCount) {
		var ret int32
		return ret
	}
	return *o.TotalPageCount
}

// GetTotalPageCountOk returns a tuple with the TotalPageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) GetTotalPageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPageCount) {
		return nil, false
	}
	return o.TotalPageCount, true
}

// HasTotalPageCount returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) HasTotalPageCount() bool {
	if o != nil && !IsNil(o.TotalPageCount) {
		return true
	}

	return false
}

// SetTotalPageCount gets a reference to the given int32 and assigns it to the TotalPageCount field.
func (o *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) SetTotalPageCount(v int32) {
	o.TotalPageCount = &v
}

func (o AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpenseQuotaInfoList) {
		toSerialize["expense_quota_info_list"] = o.ExpenseQuotaInfoList
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalPageCount) {
		toSerialize["total_page_count"] = o.TotalPageCount
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel struct {
	value *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) Get() *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) Set(val *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel(val *AlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel {
	return &NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolQuotaQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

