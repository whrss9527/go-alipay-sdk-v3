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

// checks if the AlipayMarketingActivityMerchantBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityMerchantBatchqueryResponseModel{}

// AlipayMarketingActivityMerchantBatchqueryResponseModel struct for AlipayMarketingActivityMerchantBatchqueryResponseModel
type AlipayMarketingActivityMerchantBatchqueryResponseModel struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 可用商户列表
	MerchantInfos []ActivityMerchantInfo `json:"merchant_infos,omitempty"`
	// 分页查询页码。
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询单页数据条数。
	PageSize *int32 `json:"page_size,omitempty"`
	// 可用商户总数量
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAlipayMarketingActivityMerchantBatchqueryResponseModel instantiates a new AlipayMarketingActivityMerchantBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityMerchantBatchqueryResponseModel() *AlipayMarketingActivityMerchantBatchqueryResponseModel {
	this := AlipayMarketingActivityMerchantBatchqueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityMerchantBatchqueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityMerchantBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityMerchantBatchqueryResponseModelWithDefaults() *AlipayMarketingActivityMerchantBatchqueryResponseModel {
	this := AlipayMarketingActivityMerchantBatchqueryResponseModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetMerchantInfos returns the MerchantInfos field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetMerchantInfos() []ActivityMerchantInfo {
	if o == nil || IsNil(o.MerchantInfos) {
		var ret []ActivityMerchantInfo
		return ret
	}
	return o.MerchantInfos
}

// GetMerchantInfosOk returns a tuple with the MerchantInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetMerchantInfosOk() ([]ActivityMerchantInfo, bool) {
	if o == nil || IsNil(o.MerchantInfos) {
		return nil, false
	}
	return o.MerchantInfos, true
}

// HasMerchantInfos returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) HasMerchantInfos() bool {
	if o != nil && !IsNil(o.MerchantInfos) {
		return true
	}

	return false
}

// SetMerchantInfos gets a reference to the given []ActivityMerchantInfo and assigns it to the MerchantInfos field.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) SetMerchantInfos(v []ActivityMerchantInfo) {
	o.MerchantInfos = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityMerchantBatchqueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AlipayMarketingActivityMerchantBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityMerchantBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.MerchantInfos) {
		toSerialize["merchant_infos"] = o.MerchantInfos
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityMerchantBatchqueryResponseModel struct {
	value *AlipayMarketingActivityMerchantBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) Get() *AlipayMarketingActivityMerchantBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) Set(val *AlipayMarketingActivityMerchantBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityMerchantBatchqueryResponseModel(val *AlipayMarketingActivityMerchantBatchqueryResponseModel) *NullableAlipayMarketingActivityMerchantBatchqueryResponseModel {
	return &NullableAlipayMarketingActivityMerchantBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
