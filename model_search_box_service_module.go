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

// checks if the SearchBoxServiceModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchBoxServiceModule{}

// SearchBoxServiceModule struct for SearchBoxServiceModule
type SearchBoxServiceModule struct {
	// 模块配置ID
	ModuleId *string `json:"module_id,omitempty"`
	// 搜索直达模块类型
	ModuleType *string `json:"module_type,omitempty"`
	// 服务信息列表
	ServiceInfos []SearchBoxServiceInfo `json:"service_infos,omitempty"`
}

// NewSearchBoxServiceModule instantiates a new SearchBoxServiceModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBoxServiceModule() *SearchBoxServiceModule {
	this := SearchBoxServiceModule{}
	return &this
}

// NewSearchBoxServiceModuleWithDefaults instantiates a new SearchBoxServiceModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBoxServiceModuleWithDefaults() *SearchBoxServiceModule {
	this := SearchBoxServiceModule{}
	return &this
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *SearchBoxServiceModule) GetModuleId() string {
	if o == nil || IsNil(o.ModuleId) {
		var ret string
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxServiceModule) GetModuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleId) {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *SearchBoxServiceModule) HasModuleId() bool {
	if o != nil && !IsNil(o.ModuleId) {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given string and assigns it to the ModuleId field.
func (o *SearchBoxServiceModule) SetModuleId(v string) {
	o.ModuleId = &v
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise.
func (o *SearchBoxServiceModule) GetModuleType() string {
	if o == nil || IsNil(o.ModuleType) {
		var ret string
		return ret
	}
	return *o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxServiceModule) GetModuleTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ModuleType) {
		return nil, false
	}
	return o.ModuleType, true
}

// HasModuleType returns a boolean if a field has been set.
func (o *SearchBoxServiceModule) HasModuleType() bool {
	if o != nil && !IsNil(o.ModuleType) {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given string and assigns it to the ModuleType field.
func (o *SearchBoxServiceModule) SetModuleType(v string) {
	o.ModuleType = &v
}

// GetServiceInfos returns the ServiceInfos field value if set, zero value otherwise.
func (o *SearchBoxServiceModule) GetServiceInfos() []SearchBoxServiceInfo {
	if o == nil || IsNil(o.ServiceInfos) {
		var ret []SearchBoxServiceInfo
		return ret
	}
	return o.ServiceInfos
}

// GetServiceInfosOk returns a tuple with the ServiceInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBoxServiceModule) GetServiceInfosOk() ([]SearchBoxServiceInfo, bool) {
	if o == nil || IsNil(o.ServiceInfos) {
		return nil, false
	}
	return o.ServiceInfos, true
}

// HasServiceInfos returns a boolean if a field has been set.
func (o *SearchBoxServiceModule) HasServiceInfos() bool {
	if o != nil && !IsNil(o.ServiceInfos) {
		return true
	}

	return false
}

// SetServiceInfos gets a reference to the given []SearchBoxServiceInfo and assigns it to the ServiceInfos field.
func (o *SearchBoxServiceModule) SetServiceInfos(v []SearchBoxServiceInfo) {
	o.ServiceInfos = v
}

func (o SearchBoxServiceModule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchBoxServiceModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ModuleId) {
		toSerialize["module_id"] = o.ModuleId
	}
	if !IsNil(o.ModuleType) {
		toSerialize["module_type"] = o.ModuleType
	}
	if !IsNil(o.ServiceInfos) {
		toSerialize["service_infos"] = o.ServiceInfos
	}
	return toSerialize, nil
}

type NullableSearchBoxServiceModule struct {
	value *SearchBoxServiceModule
	isSet bool
}

func (v NullableSearchBoxServiceModule) Get() *SearchBoxServiceModule {
	return v.value
}

func (v *NullableSearchBoxServiceModule) Set(val *SearchBoxServiceModule) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBoxServiceModule) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBoxServiceModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBoxServiceModule(val *SearchBoxServiceModule) *NullableSearchBoxServiceModule {
	return &NullableSearchBoxServiceModule{value: val, isSet: true}
}

func (v NullableSearchBoxServiceModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBoxServiceModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


