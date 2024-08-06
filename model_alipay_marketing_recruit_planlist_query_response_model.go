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

// checks if the AlipayMarketingRecruitPlanlistQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingRecruitPlanlistQueryResponseModel{}

// AlipayMarketingRecruitPlanlistQueryResponseModel struct for AlipayMarketingRecruitPlanlistQueryResponseModel
type AlipayMarketingRecruitPlanlistQueryResponseModel struct {
	// 方案列表
	Data []RecruitPlanLight `json:"data,omitempty"`
	// 第几页，默认1（从1开始计数）
	PageNum *int32 `json:"page_num,omitempty"`
	// 每页记录条数，默认20
	PageSize *int32 `json:"page_size,omitempty"`
	// 总数
	Total *int32 `json:"total,omitempty"`
}

// NewAlipayMarketingRecruitPlanlistQueryResponseModel instantiates a new AlipayMarketingRecruitPlanlistQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingRecruitPlanlistQueryResponseModel() *AlipayMarketingRecruitPlanlistQueryResponseModel {
	this := AlipayMarketingRecruitPlanlistQueryResponseModel{}
	return &this
}

// NewAlipayMarketingRecruitPlanlistQueryResponseModelWithDefaults instantiates a new AlipayMarketingRecruitPlanlistQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingRecruitPlanlistQueryResponseModelWithDefaults() *AlipayMarketingRecruitPlanlistQueryResponseModel {
	this := AlipayMarketingRecruitPlanlistQueryResponseModel{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetData() []RecruitPlanLight {
	if o == nil || IsNil(o.Data) {
		var ret []RecruitPlanLight
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetDataOk() ([]RecruitPlanLight, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RecruitPlanLight and assigns it to the Data field.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) SetData(v []RecruitPlanLight) {
	o.Data = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *AlipayMarketingRecruitPlanlistQueryResponseModel) SetTotal(v int32) {
	o.Total = &v
}

func (o AlipayMarketingRecruitPlanlistQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingRecruitPlanlistQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableAlipayMarketingRecruitPlanlistQueryResponseModel struct {
	value *AlipayMarketingRecruitPlanlistQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingRecruitPlanlistQueryResponseModel) Get() *AlipayMarketingRecruitPlanlistQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingRecruitPlanlistQueryResponseModel) Set(val *AlipayMarketingRecruitPlanlistQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingRecruitPlanlistQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingRecruitPlanlistQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingRecruitPlanlistQueryResponseModel(val *AlipayMarketingRecruitPlanlistQueryResponseModel) *NullableAlipayMarketingRecruitPlanlistQueryResponseModel {
	return &NullableAlipayMarketingRecruitPlanlistQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingRecruitPlanlistQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingRecruitPlanlistQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

