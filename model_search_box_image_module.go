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

// checks if the SearchBoxImageModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxImageModule{}

// SearchBoxImageModule struct for SearchBoxImageModule
type SearchBoxImageModule struct {
	// 申请单号
	ApplyNo *string `json:"apply_no,omitempty"`
	// 审核失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 模块配置ID
	ModuleId *string `json:"module_id,omitempty"`
	// 搜索直达模块类型
	ModuleType *string `json:"module_type,omitempty"`
	// 状态，INITIAL-初始/AUDIT-审核中/CANCEL-已取消/ONLINE-已上架/REJECT-驳回/OFFLINE-已下架/EXPIRE-已失效
	Status *string `json:"status,omitempty"`
}

// NewSearchBoxImageModule instantiates a new SearchBoxImageModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxImageModule() *SearchBoxImageModule {
	this := SearchBoxImageModule{}
	return &this
}

// NewSearchBoxImageModuleWithDefaults instantiates a new SearchBoxImageModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxImageModuleWithDefaults() *SearchBoxImageModule {
	this := SearchBoxImageModule{}
	return &this
}

// GetApplyNo returns the ApplyNo field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetApplyNo() string {
	if o == nil || IsNil(o.ApplyNo) {
		var ret string
		return ret
	}
	return *o.ApplyNo
}

// GetApplyNoOk returns a tuple with the ApplyNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetApplyNoOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyNo) {
		return nil, false
	}
	return o.ApplyNo, true
}

// HasApplyNo returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasApplyNo() bool {
	if o != nil && !IsNil(o.ApplyNo) {
		return true
	}

	return false
}

// SetApplyNo gets a reference to the given string and assigns it to the ApplyNo field.
func (o *SearchBoxImageModule) SetApplyNo(v string) {
	o.ApplyNo = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *SearchBoxImageModule) SetFailReason(v string) {
	o.FailReason = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *SearchBoxImageModule) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetModuleId() string {
	if o == nil || IsNil(o.ModuleId) {
		var ret string
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetModuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given string and assigns it to the ModuleId field.
func (o *SearchBoxImageModule) SetModuleId(v string) {
	o.ModuleId = &v
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetModuleType() string {
	if o == nil || IsNil(o.ModuleType) {
		var ret string
		return ret
	}
	return *o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetModuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleType) {
		return nil, false
	}
	return o.ModuleType, true
}

// HasModuleType returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasModuleType() bool {
	if o != nil && !IsNil(o.ModuleType) {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given string and assigns it to the ModuleType field.
func (o *SearchBoxImageModule) SetModuleType(v string) {
	o.ModuleType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchBoxImageModule) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxImageModule) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchBoxImageModule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SearchBoxImageModule) SetStatus(v string) {
	o.Status = &v
}

func (o SearchBoxImageModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxImageModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyNo) {
		toSerialize["apply_no"] = o.ApplyNo
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.ModuleId) {
		toSerialize["module_id"] = o.ModuleId
	}
	if !IsNil(o.ModuleType) {
		toSerialize["module_type"] = o.ModuleType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableSearchBoxImageModule struct {
	value *SearchBoxImageModule
	isSet bool
}

func (v NullableSearchBoxImageModule) Get() *SearchBoxImageModule {
	return v.value
}

func (v *NullableSearchBoxImageModule) Set(val *SearchBoxImageModule) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxImageModule) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxImageModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxImageModule(val *SearchBoxImageModule) *NullableSearchBoxImageModule {
	return &NullableSearchBoxImageModule{value: val, isSet: true}
}

func (v NullableSearchBoxImageModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxImageModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

