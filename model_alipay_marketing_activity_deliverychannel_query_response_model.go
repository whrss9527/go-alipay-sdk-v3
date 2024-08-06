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

// checks if the AlipayMarketingActivityDeliverychannelQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityDeliverychannelQueryResponseModel{}

// AlipayMarketingActivityDeliverychannelQueryResponseModel struct for AlipayMarketingActivityDeliverychannelQueryResponseModel
type AlipayMarketingActivityDeliverychannelQueryResponseModel struct {
	// \"可投放的渠道信息列表。 用于表达当前可以投放的渠道列表信息\"
	DeliveryChannelInfoList []DeliveryChannelInfo `json:"delivery_channel_info_list,omitempty"`
	// 查询的页码。 特别说明： 页码从1开始。
	PageNum *string `json:"page_num,omitempty"`
	// 每页查询个数
	PageSize *string `json:"page_size,omitempty"`
	// 可返回的渠道总数
	TotalSize *string `json:"total_size,omitempty"`
}

// NewAlipayMarketingActivityDeliverychannelQueryResponseModel instantiates a new AlipayMarketingActivityDeliverychannelQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityDeliverychannelQueryResponseModel() *AlipayMarketingActivityDeliverychannelQueryResponseModel {
	this := AlipayMarketingActivityDeliverychannelQueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityDeliverychannelQueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityDeliverychannelQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityDeliverychannelQueryResponseModelWithDefaults() *AlipayMarketingActivityDeliverychannelQueryResponseModel {
	this := AlipayMarketingActivityDeliverychannelQueryResponseModel{}
	return &this
}

// GetDeliveryChannelInfoList returns the DeliveryChannelInfoList field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetDeliveryChannelInfoList() []DeliveryChannelInfo {
	if o == nil || IsNil(o.DeliveryChannelInfoList) {
		var ret []DeliveryChannelInfo
		return ret
	}
	return o.DeliveryChannelInfoList
}

// GetDeliveryChannelInfoListOk returns a tuple with the DeliveryChannelInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetDeliveryChannelInfoListOk() ([]DeliveryChannelInfo, bool) {
	if o == nil || IsNil(o.DeliveryChannelInfoList) {
		return nil, false
	}
	return o.DeliveryChannelInfoList, true
}

// HasDeliveryChannelInfoList returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) HasDeliveryChannelInfoList() bool {
	if o != nil && !IsNil(o.DeliveryChannelInfoList) {
		return true
	}

	return false
}

// SetDeliveryChannelInfoList gets a reference to the given []DeliveryChannelInfo and assigns it to the DeliveryChannelInfoList field.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) SetDeliveryChannelInfoList(v []DeliveryChannelInfo) {
	o.DeliveryChannelInfoList = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetPageNum() string {
	if o == nil || IsNil(o.PageNum) {
		var ret string
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetPageNumOk() (*string, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given string and assigns it to the PageNum field.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) SetPageNum(v string) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) SetPageSize(v string) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *AlipayMarketingActivityDeliverychannelQueryResponseModel) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o AlipayMarketingActivityDeliverychannelQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityDeliverychannelQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryChannelInfoList) {
		toSerialize["delivery_channel_info_list"] = o.DeliveryChannelInfoList
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

type NullableAlipayMarketingActivityDeliverychannelQueryResponseModel struct {
	value *AlipayMarketingActivityDeliverychannelQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) Get() *AlipayMarketingActivityDeliverychannelQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) Set(val *AlipayMarketingActivityDeliverychannelQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityDeliverychannelQueryResponseModel(val *AlipayMarketingActivityDeliverychannelQueryResponseModel) *NullableAlipayMarketingActivityDeliverychannelQueryResponseModel {
	return &NullableAlipayMarketingActivityDeliverychannelQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityDeliverychannelQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
