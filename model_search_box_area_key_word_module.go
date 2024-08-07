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

// checks if the SearchBoxAreaKeyWordModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxAreaKeyWordModule{}

// SearchBoxAreaKeyWordModule struct for SearchBoxAreaKeyWordModule
type SearchBoxAreaKeyWordModule struct {
	LatestAuditAreaKeywordInfo *SearchBoxKeywordInfo `json:"latest_audit_area_keyword_info,omitempty"`
	// 模块配置ID
	ModuleId *string `json:"module_id,omitempty"`
	// 搜索直达模块类型
	ModuleType           *string               `json:"module_type,omitempty"`
	ValidAreaKeywordInfo *SearchBoxKeywordInfo `json:"valid_area_keyword_info,omitempty"`
}

// NewSearchBoxAreaKeyWordModule instantiates a new SearchBoxAreaKeyWordModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxAreaKeyWordModule() *SearchBoxAreaKeyWordModule {
	this := SearchBoxAreaKeyWordModule{}
	return &this
}

// NewSearchBoxAreaKeyWordModuleWithDefaults instantiates a new SearchBoxAreaKeyWordModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxAreaKeyWordModuleWithDefaults() *SearchBoxAreaKeyWordModule {
	this := SearchBoxAreaKeyWordModule{}
	return &this
}

// GetLatestAuditAreaKeywordInfo returns the LatestAuditAreaKeywordInfo field value if set, zero value otherwise.
func (o *SearchBoxAreaKeyWordModule) GetLatestAuditAreaKeywordInfo() SearchBoxKeywordInfo {
	if o == nil || IsNil(o.LatestAuditAreaKeywordInfo) {
		var ret SearchBoxKeywordInfo
		return ret
	}
	return *o.LatestAuditAreaKeywordInfo
}

// GetLatestAuditAreaKeywordInfoOk returns a tuple with the LatestAuditAreaKeywordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAreaKeyWordModule) GetLatestAuditAreaKeywordInfoOk() (*SearchBoxKeywordInfo, bool) {
	if o == nil || IsNil(o.LatestAuditAreaKeywordInfo) {
		return nil, false
	}
	return o.LatestAuditAreaKeywordInfo, true
}

// HasLatestAuditAreaKeywordInfo returns a boolean if a field has been set.
func (o *SearchBoxAreaKeyWordModule) HasLatestAuditAreaKeywordInfo() bool {
	if o != nil && !IsNil(o.LatestAuditAreaKeywordInfo) {
		return true
	}

	return false
}

// SetLatestAuditAreaKeywordInfo gets a reference to the given SearchBoxKeywordInfo and assigns it to the LatestAuditAreaKeywordInfo field.
func (o *SearchBoxAreaKeyWordModule) SetLatestAuditAreaKeywordInfo(v SearchBoxKeywordInfo) {
	o.LatestAuditAreaKeywordInfo = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *SearchBoxAreaKeyWordModule) GetModuleId() string {
	if o == nil || IsNil(o.ModuleId) {
		var ret string
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAreaKeyWordModule) GetModuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *SearchBoxAreaKeyWordModule) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given string and assigns it to the ModuleId field.
func (o *SearchBoxAreaKeyWordModule) SetModuleId(v string) {
	o.ModuleId = &v
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise.
func (o *SearchBoxAreaKeyWordModule) GetModuleType() string {
	if o == nil || IsNil(o.ModuleType) {
		var ret string
		return ret
	}
	return *o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAreaKeyWordModule) GetModuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleType) {
		return nil, false
	}
	return o.ModuleType, true
}

// HasModuleType returns a boolean if a field has been set.
func (o *SearchBoxAreaKeyWordModule) HasModuleType() bool {
	if o != nil && !IsNil(o.ModuleType) {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given string and assigns it to the ModuleType field.
func (o *SearchBoxAreaKeyWordModule) SetModuleType(v string) {
	o.ModuleType = &v
}

// GetValidAreaKeywordInfo returns the ValidAreaKeywordInfo field value if set, zero value otherwise.
func (o *SearchBoxAreaKeyWordModule) GetValidAreaKeywordInfo() SearchBoxKeywordInfo {
	if o == nil || IsNil(o.ValidAreaKeywordInfo) {
		var ret SearchBoxKeywordInfo
		return ret
	}
	return *o.ValidAreaKeywordInfo
}

// GetValidAreaKeywordInfoOk returns a tuple with the ValidAreaKeywordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAreaKeyWordModule) GetValidAreaKeywordInfoOk() (*SearchBoxKeywordInfo, bool) {
	if o == nil || IsNil(o.ValidAreaKeywordInfo) {
		return nil, false
	}
	return o.ValidAreaKeywordInfo, true
}

// HasValidAreaKeywordInfo returns a boolean if a field has been set.
func (o *SearchBoxAreaKeyWordModule) HasValidAreaKeywordInfo() bool {
	if o != nil && !IsNil(o.ValidAreaKeywordInfo) {
		return true
	}

	return false
}

// SetValidAreaKeywordInfo gets a reference to the given SearchBoxKeywordInfo and assigns it to the ValidAreaKeywordInfo field.
func (o *SearchBoxAreaKeyWordModule) SetValidAreaKeywordInfo(v SearchBoxKeywordInfo) {
	o.ValidAreaKeywordInfo = &v
}

func (o SearchBoxAreaKeyWordModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxAreaKeyWordModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LatestAuditAreaKeywordInfo) {
		toSerialize["latest_audit_area_keyword_info"] = o.LatestAuditAreaKeywordInfo
	}
	if !IsNil(o.ModuleId) {
		toSerialize["module_id"] = o.ModuleId
	}
	if !IsNil(o.ModuleType) {
		toSerialize["module_type"] = o.ModuleType
	}
	if !IsNil(o.ValidAreaKeywordInfo) {
		toSerialize["valid_area_keyword_info"] = o.ValidAreaKeywordInfo
	}
	return toSerialize, nil
}

type NullableSearchBoxAreaKeyWordModule struct {
	value *SearchBoxAreaKeyWordModule
	isSet bool
}

func (v NullableSearchBoxAreaKeyWordModule) Get() *SearchBoxAreaKeyWordModule {
	return v.value
}

func (v *NullableSearchBoxAreaKeyWordModule) Set(val *SearchBoxAreaKeyWordModule) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxAreaKeyWordModule) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxAreaKeyWordModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxAreaKeyWordModule(val *SearchBoxAreaKeyWordModule) *NullableSearchBoxAreaKeyWordModule {
	return &NullableSearchBoxAreaKeyWordModule{value: val, isSet: true}
}

func (v NullableSearchBoxAreaKeyWordModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxAreaKeyWordModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
