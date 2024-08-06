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

// checks if the AlipayMarketingActivityUserBatchqueryvoucherResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityUserBatchqueryvoucherResponseModel{}

// AlipayMarketingActivityUserBatchqueryvoucherResponseModel struct for AlipayMarketingActivityUserBatchqueryvoucherResponseModel
type AlipayMarketingActivityUserBatchqueryvoucherResponseModel struct {
	// 分页查询页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询单页数据条数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总数量
	TotalSize *int32 `json:"total_size,omitempty"`
	// 用户券实例
	UserVoucherInfos []UserVoucherInfo `json:"user_voucher_infos,omitempty"`
}

// NewAlipayMarketingActivityUserBatchqueryvoucherResponseModel instantiates a new AlipayMarketingActivityUserBatchqueryvoucherResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityUserBatchqueryvoucherResponseModel() *AlipayMarketingActivityUserBatchqueryvoucherResponseModel {
	this := AlipayMarketingActivityUserBatchqueryvoucherResponseModel{}
	return &this
}

// NewAlipayMarketingActivityUserBatchqueryvoucherResponseModelWithDefaults instantiates a new AlipayMarketingActivityUserBatchqueryvoucherResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityUserBatchqueryvoucherResponseModelWithDefaults() *AlipayMarketingActivityUserBatchqueryvoucherResponseModel {
	this := AlipayMarketingActivityUserBatchqueryvoucherResponseModel{}
	return &this
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

// GetUserVoucherInfos returns the UserVoucherInfos field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetUserVoucherInfos() []UserVoucherInfo {
	if o == nil || IsNil(o.UserVoucherInfos) {
		var ret []UserVoucherInfo
		return ret
	}
	return o.UserVoucherInfos
}

// GetUserVoucherInfosOk returns a tuple with the UserVoucherInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) GetUserVoucherInfosOk() ([]UserVoucherInfo, bool) {
	if o == nil || IsNil(o.UserVoucherInfos) {
		return nil, false
	}
	return o.UserVoucherInfos, true
}

// HasUserVoucherInfos returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) HasUserVoucherInfos() bool {
	if o != nil && !IsNil(o.UserVoucherInfos) {
		return true
	}

	return false
}

// SetUserVoucherInfos gets a reference to the given []UserVoucherInfo and assigns it to the UserVoucherInfos field.
func (o *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) SetUserVoucherInfos(v []UserVoucherInfo) {
	o.UserVoucherInfos = v
}

func (o AlipayMarketingActivityUserBatchqueryvoucherResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityUserBatchqueryvoucherResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	if !IsNil(o.UserVoucherInfos) {
		toSerialize["user_voucher_infos"] = o.UserVoucherInfos
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel struct {
	value *AlipayMarketingActivityUserBatchqueryvoucherResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) Get() *AlipayMarketingActivityUserBatchqueryvoucherResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) Set(val *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel(val *AlipayMarketingActivityUserBatchqueryvoucherResponseModel) *NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel {
	return &NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityUserBatchqueryvoucherResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
