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

// checks if the SearchBoxAccountModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxAccountModule{}

// SearchBoxAccountModule struct for SearchBoxAccountModule
type SearchBoxAccountModule struct {
	// 应用信息列表
	AppInfos []SearchBoxAppInfo `json:"app_infos,omitempty"`
	// 模块配置ID
	ModuleId *string `json:"module_id,omitempty"`
	// 搜索直达模块类型
	ModuleType *string `json:"module_type,omitempty"`
}

// NewSearchBoxAccountModule instantiates a new SearchBoxAccountModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxAccountModule() *SearchBoxAccountModule {
	this := SearchBoxAccountModule{}
	return &this
}

// NewSearchBoxAccountModuleWithDefaults instantiates a new SearchBoxAccountModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxAccountModuleWithDefaults() *SearchBoxAccountModule {
	this := SearchBoxAccountModule{}
	return &this
}

// GetAppInfos returns the AppInfos field value if set, zero value otherwise.
func (o *SearchBoxAccountModule) GetAppInfos() []SearchBoxAppInfo {
	if o == nil || IsNil(o.AppInfos) {
		var ret []SearchBoxAppInfo
		return ret
	}
	return o.AppInfos
}

// GetAppInfosOk returns a tuple with the AppInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAccountModule) GetAppInfosOk() ([]SearchBoxAppInfo, bool) {
	if o == nil || IsNil(o.AppInfos) {
		return nil, false
	}
	return o.AppInfos, true
}

// HasAppInfos returns a boolean if a field has been set.
func (o *SearchBoxAccountModule) HasAppInfos() bool {
	if o != nil && !IsNil(o.AppInfos) {
		return true
	}

	return false
}

// SetAppInfos gets a reference to the given []SearchBoxAppInfo and assigns it to the AppInfos field.
func (o *SearchBoxAccountModule) SetAppInfos(v []SearchBoxAppInfo) {
	o.AppInfos = v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *SearchBoxAccountModule) GetModuleId() string {
	if o == nil || IsNil(o.ModuleId) {
		var ret string
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAccountModule) GetModuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *SearchBoxAccountModule) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given string and assigns it to the ModuleId field.
func (o *SearchBoxAccountModule) SetModuleId(v string) {
	o.ModuleId = &v
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise.
func (o *SearchBoxAccountModule) GetModuleType() string {
	if o == nil || IsNil(o.ModuleType) {
		var ret string
		return ret
	}
	return *o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxAccountModule) GetModuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleType) {
		return nil, false
	}
	return o.ModuleType, true
}

// HasModuleType returns a boolean if a field has been set.
func (o *SearchBoxAccountModule) HasModuleType() bool {
	if o != nil && !IsNil(o.ModuleType) {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given string and assigns it to the ModuleType field.
func (o *SearchBoxAccountModule) SetModuleType(v string) {
	o.ModuleType = &v
}

func (o SearchBoxAccountModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxAccountModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppInfos) {
		toSerialize["app_infos"] = o.AppInfos
	}
	if !IsNil(o.ModuleId) {
		toSerialize["module_id"] = o.ModuleId
	}
	if !IsNil(o.ModuleType) {
		toSerialize["module_type"] = o.ModuleType
	}
	return toSerialize, nil
}

type NullableSearchBoxAccountModule struct {
	value *SearchBoxAccountModule
	isSet bool
}

func (v NullableSearchBoxAccountModule) Get() *SearchBoxAccountModule {
	return v.value
}

func (v *NullableSearchBoxAccountModule) Set(val *SearchBoxAccountModule) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxAccountModule) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxAccountModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxAccountModule(val *SearchBoxAccountModule) *NullableSearchBoxAccountModule {
	return &NullableSearchBoxAccountModule{value: val, isSet: true}
}

func (v NullableSearchBoxAccountModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxAccountModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


