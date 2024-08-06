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

// checks if the AlipayMarketingActivityShopBatchqueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityShopBatchqueryResponseModel{}

// AlipayMarketingActivityShopBatchqueryResponseModel struct for AlipayMarketingActivityShopBatchqueryResponseModel
type AlipayMarketingActivityShopBatchqueryResponseModel struct {
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 分页查询页码。
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询单页数据条数。
	PageSize *int32 `json:"page_size,omitempty"`
	// 可用门店列表
	ShopInfos []ActivityShopInfo `json:"shop_infos,omitempty"`
	// 可用门店总数量
	TotalSize *int32 `json:"total_size,omitempty"`
}

// NewAlipayMarketingActivityShopBatchqueryResponseModel instantiates a new AlipayMarketingActivityShopBatchqueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityShopBatchqueryResponseModel() *AlipayMarketingActivityShopBatchqueryResponseModel {
	this := AlipayMarketingActivityShopBatchqueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityShopBatchqueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityShopBatchqueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityShopBatchqueryResponseModelWithDefaults() *AlipayMarketingActivityShopBatchqueryResponseModel {
	this := AlipayMarketingActivityShopBatchqueryResponseModel{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetShopInfos returns the ShopInfos field value if set, zero value otherwise.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetShopInfos() []ActivityShopInfo {
	if o == nil || IsNil(o.ShopInfos) {
		var ret []ActivityShopInfo
		return ret
	}
	return o.ShopInfos
}

// GetShopInfosOk returns a tuple with the ShopInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetShopInfosOk() ([]ActivityShopInfo, bool) {
	if o == nil || IsNil(o.ShopInfos) {
		return nil, false
	}
	return o.ShopInfos, true
}

// HasShopInfos returns a boolean if a field has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) HasShopInfos() bool {
	if o != nil && !IsNil(o.ShopInfos) {
		return true
	}

	return false
}

// SetShopInfos gets a reference to the given []ActivityShopInfo and assigns it to the ShopInfos field.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) SetShopInfos(v []ActivityShopInfo) {
	o.ShopInfos = v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetTotalSize() int32 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) GetTotalSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityShopBatchqueryResponseModel) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o AlipayMarketingActivityShopBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityShopBatchqueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.ShopInfos) {
		toSerialize["shop_infos"] = o.ShopInfos
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityShopBatchqueryResponseModel struct {
	value *AlipayMarketingActivityShopBatchqueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityShopBatchqueryResponseModel) Get() *AlipayMarketingActivityShopBatchqueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityShopBatchqueryResponseModel) Set(val *AlipayMarketingActivityShopBatchqueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityShopBatchqueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityShopBatchqueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityShopBatchqueryResponseModel(val *AlipayMarketingActivityShopBatchqueryResponseModel) *NullableAlipayMarketingActivityShopBatchqueryResponseModel {
	return &NullableAlipayMarketingActivityShopBatchqueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityShopBatchqueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityShopBatchqueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

