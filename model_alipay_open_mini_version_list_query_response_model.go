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

// checks if the AlipayOpenMiniVersionListQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniVersionListQueryResponseModel{}

// AlipayOpenMiniVersionListQueryResponseModel struct for AlipayOpenMiniVersionListQueryResponseModel
type AlipayOpenMiniVersionListQueryResponseModel struct {
	// 版本列表，根据版本号倒叙排列，即版本号大的在前面；如果不存在任何版本，返回空列表
	AppVersionInfos []AppVersionInfo `json:"app_version_infos,omitempty"`
	// 小程序支付宝端每个状态的最新版本号列表(历史返回值，已不推荐使用，逐渐废弃)
	AppVersions []string `json:"app_versions,omitempty"`
}

// NewAlipayOpenMiniVersionListQueryResponseModel instantiates a new AlipayOpenMiniVersionListQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniVersionListQueryResponseModel() *AlipayOpenMiniVersionListQueryResponseModel {
	this := AlipayOpenMiniVersionListQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniVersionListQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniVersionListQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniVersionListQueryResponseModelWithDefaults() *AlipayOpenMiniVersionListQueryResponseModel {
	this := AlipayOpenMiniVersionListQueryResponseModel{}
	return &this
}

// GetAppVersionInfos returns the AppVersionInfos field value if set, zero value otherwise.
func (o *AlipayOpenMiniVersionListQueryResponseModel) GetAppVersionInfos() []AppVersionInfo {
	if o == nil || IsNil(o.AppVersionInfos) {
		var ret []AppVersionInfo
		return ret
	}
	return o.AppVersionInfos
}

// GetAppVersionInfosOk returns a tuple with the AppVersionInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniVersionListQueryResponseModel) GetAppVersionInfosOk() ([]AppVersionInfo, bool) {
	if o == nil || IsNil(o.AppVersionInfos) {
		return nil, false
	}
	return o.AppVersionInfos, true
}

// HasAppVersionInfos returns a boolean if a field has been set.
func (o *AlipayOpenMiniVersionListQueryResponseModel) HasAppVersionInfos() bool {
	if o != nil && !IsNil(o.AppVersionInfos) {
		return true
	}

	return false
}

// SetAppVersionInfos gets a reference to the given []AppVersionInfo and assigns it to the AppVersionInfos field.
func (o *AlipayOpenMiniVersionListQueryResponseModel) SetAppVersionInfos(v []AppVersionInfo) {
	o.AppVersionInfos = v
}

// GetAppVersions returns the AppVersions field value if set, zero value otherwise.
func (o *AlipayOpenMiniVersionListQueryResponseModel) GetAppVersions() []string {
	if o == nil || IsNil(o.AppVersions) {
		var ret []string
		return ret
	}
	return o.AppVersions
}

// GetAppVersionsOk returns a tuple with the AppVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniVersionListQueryResponseModel) GetAppVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppVersions) {
		return nil, false
	}
	return o.AppVersions, true
}

// HasAppVersions returns a boolean if a field has been set.
func (o *AlipayOpenMiniVersionListQueryResponseModel) HasAppVersions() bool {
	if o != nil && !IsNil(o.AppVersions) {
		return true
	}

	return false
}

// SetAppVersions gets a reference to the given []string and assigns it to the AppVersions field.
func (o *AlipayOpenMiniVersionListQueryResponseModel) SetAppVersions(v []string) {
	o.AppVersions = v
}

func (o AlipayOpenMiniVersionListQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniVersionListQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppVersionInfos) {
		toSerialize["app_version_infos"] = o.AppVersionInfos
	}
	if !IsNil(o.AppVersions) {
		toSerialize["app_versions"] = o.AppVersions
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniVersionListQueryResponseModel struct {
	value *AlipayOpenMiniVersionListQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniVersionListQueryResponseModel) Get() *AlipayOpenMiniVersionListQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniVersionListQueryResponseModel) Set(val *AlipayOpenMiniVersionListQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniVersionListQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniVersionListQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniVersionListQueryResponseModel(val *AlipayOpenMiniVersionListQueryResponseModel) *NullableAlipayOpenMiniVersionListQueryResponseModel {
	return &NullableAlipayOpenMiniVersionListQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniVersionListQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniVersionListQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


