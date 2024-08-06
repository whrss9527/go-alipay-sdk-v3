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

// checks if the AlipayIserviceCcmInstanceQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmInstanceQueryResponseModel{}

// AlipayIserviceCcmInstanceQueryResponseModel struct for AlipayIserviceCcmInstanceQueryResponseModel
type AlipayIserviceCcmInstanceQueryResponseModel struct {
	// 租户实例列表
	Instances []Instance `json:"instances,omitempty"`
	// 查询结果的页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询时设置的每页记录数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAlipayIserviceCcmInstanceQueryResponseModel instantiates a new AlipayIserviceCcmInstanceQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmInstanceQueryResponseModel() *AlipayIserviceCcmInstanceQueryResponseModel {
	this := AlipayIserviceCcmInstanceQueryResponseModel{}
	return &this
}

// NewAlipayIserviceCcmInstanceQueryResponseModelWithDefaults instantiates a new AlipayIserviceCcmInstanceQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmInstanceQueryResponseModelWithDefaults() *AlipayIserviceCcmInstanceQueryResponseModel {
	this := AlipayIserviceCcmInstanceQueryResponseModel{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetInstances() []Instance {
	if o == nil || IsNil(o.Instances) {
		var ret []Instance
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetInstancesOk() ([]Instance, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []Instance and assigns it to the Instances field.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) SetInstances(v []Instance) {
	o.Instances = v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AlipayIserviceCcmInstanceQueryResponseModel) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AlipayIserviceCcmInstanceQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmInstanceQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmInstanceQueryResponseModel struct {
	value *AlipayIserviceCcmInstanceQueryResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmInstanceQueryResponseModel) Get() *AlipayIserviceCcmInstanceQueryResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmInstanceQueryResponseModel) Set(val *AlipayIserviceCcmInstanceQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmInstanceQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmInstanceQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmInstanceQueryResponseModel(val *AlipayIserviceCcmInstanceQueryResponseModel) *NullableAlipayIserviceCcmInstanceQueryResponseModel {
	return &NullableAlipayIserviceCcmInstanceQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmInstanceQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmInstanceQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
