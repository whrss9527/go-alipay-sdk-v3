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

// checks if the AlipayIserviceCcmRolePageQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmRolePageQueryResponseModel{}

// AlipayIserviceCcmRolePageQueryResponseModel struct for AlipayIserviceCcmRolePageQueryResponseModel
type AlipayIserviceCcmRolePageQueryResponseModel struct {
	// 查询结果的页码
	PageNum *int32 `json:"page_num,omitempty"`
	// 分页查询时设置的每页记录数
	PageSize *int32 `json:"page_size,omitempty"`
	// 角色列表
	Roles []Role `json:"roles,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewAlipayIserviceCcmRolePageQueryResponseModel instantiates a new AlipayIserviceCcmRolePageQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmRolePageQueryResponseModel() *AlipayIserviceCcmRolePageQueryResponseModel {
	this := AlipayIserviceCcmRolePageQueryResponseModel{}
	return &this
}

// NewAlipayIserviceCcmRolePageQueryResponseModelWithDefaults instantiates a new AlipayIserviceCcmRolePageQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmRolePageQueryResponseModelWithDefaults() *AlipayIserviceCcmRolePageQueryResponseModel {
	this := AlipayIserviceCcmRolePageQueryResponseModel{}
	return &this
}

// GetPageNum returns the PageNum field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetPageNum() int32 {
	if o == nil || IsNil(o.PageNum) {
		var ret int32
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetPageNumOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}
	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int32 and assigns it to the PageNum field.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) SetPageNum(v int32) {
	o.PageNum = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) SetRoles(v []Role) {
	o.Roles = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *AlipayIserviceCcmRolePageQueryResponseModel) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o AlipayIserviceCcmRolePageQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmRolePageQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNum) {
		toSerialize["page_num"] = o.PageNum
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmRolePageQueryResponseModel struct {
	value *AlipayIserviceCcmRolePageQueryResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmRolePageQueryResponseModel) Get() *AlipayIserviceCcmRolePageQueryResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmRolePageQueryResponseModel) Set(val *AlipayIserviceCcmRolePageQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmRolePageQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmRolePageQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmRolePageQueryResponseModel(val *AlipayIserviceCcmRolePageQueryResponseModel) *NullableAlipayIserviceCcmRolePageQueryResponseModel {
	return &NullableAlipayIserviceCcmRolePageQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmRolePageQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmRolePageQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
