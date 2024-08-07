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

// checks if the AlipayMarketingCampaignCashListQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCampaignCashListQueryResponseModel{}

// AlipayMarketingCampaignCashListQueryResponseModel struct for AlipayMarketingCampaignCashListQueryResponseModel
type AlipayMarketingCampaignCashListQueryResponseModel struct {
	// 活动列表
	CampList []CashCampaignInfo `json:"camp_list,omitempty"`
	// 分页的页码,起始从1开始
	PageIndex *string `json:"page_index,omitempty"`
	// 分页每页大小
	PageSize *string `json:"page_size,omitempty"`
	// 活动总个数
	TotalSize *string `json:"total_size,omitempty"`
}

// NewAlipayMarketingCampaignCashListQueryResponseModel instantiates a new AlipayMarketingCampaignCashListQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCampaignCashListQueryResponseModel() *AlipayMarketingCampaignCashListQueryResponseModel {
	this := AlipayMarketingCampaignCashListQueryResponseModel{}
	return &this
}

// NewAlipayMarketingCampaignCashListQueryResponseModelWithDefaults instantiates a new AlipayMarketingCampaignCashListQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCampaignCashListQueryResponseModelWithDefaults() *AlipayMarketingCampaignCashListQueryResponseModel {
	this := AlipayMarketingCampaignCashListQueryResponseModel{}
	return &this
}

// GetCampList returns the CampList field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetCampList() []CashCampaignInfo {
	if o == nil || IsNil(o.CampList) {
		var ret []CashCampaignInfo
		return ret
	}
	return o.CampList
}

// GetCampListOk returns a tuple with the CampList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetCampListOk() ([]CashCampaignInfo, bool) {
	if o == nil || IsNil(o.CampList) {
		return nil, false
	}
	return o.CampList, true
}

// HasCampList returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) HasCampList() bool {
	if o != nil && !IsNil(o.CampList) {
		return true
	}

	return false
}

// SetCampList gets a reference to the given []CashCampaignInfo and assigns it to the CampList field.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) SetCampList(v []CashCampaignInfo) {
	o.CampList = v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetPageIndex() string {
	if o == nil || IsNil(o.PageIndex) {
		var ret string
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetPageIndexOk() (*string, bool) {
	if o == nil || IsNil(o.PageIndex) {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) HasPageIndex() bool {
	if o != nil && !IsNil(o.PageIndex) {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given string and assigns it to the PageIndex field.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) SetPageIndex(v string) {
	o.PageIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetPageSize() string {
	if o == nil || IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetPageSizeOk() (*string, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) SetPageSize(v string) {
	o.PageSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *AlipayMarketingCampaignCashListQueryResponseModel) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o AlipayMarketingCampaignCashListQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCampaignCashListQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampList) {
		toSerialize["camp_list"] = o.CampList
	}
	if !IsNil(o.PageIndex) {
		toSerialize["page_index"] = o.PageIndex
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total_size"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCampaignCashListQueryResponseModel struct {
	value *AlipayMarketingCampaignCashListQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCampaignCashListQueryResponseModel) Get() *AlipayMarketingCampaignCashListQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCampaignCashListQueryResponseModel) Set(val *AlipayMarketingCampaignCashListQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCampaignCashListQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCampaignCashListQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCampaignCashListQueryResponseModel(val *AlipayMarketingCampaignCashListQueryResponseModel) *NullableAlipayMarketingCampaignCashListQueryResponseModel {
	return &NullableAlipayMarketingCampaignCashListQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCampaignCashListQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCampaignCashListQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
