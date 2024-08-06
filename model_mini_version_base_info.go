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

// checks if the MiniVersionBaseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiniVersionBaseInfo{}

// MiniVersionBaseInfo struct for MiniVersionBaseInfo
type MiniVersionBaseInfo struct {
	// 版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 小程序所属来源
	BuildSource *string `json:"build_source,omitempty"`
	// com.alipay.alipaywallet:支付宝，com.amap.app:高德
	BundleId *string `json:"bundle_id,omitempty"`
	// 小程序开发者ID
	DevId *string `json:"dev_id,omitempty"`
	// 小程序灰度值
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// 小程序应用ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 小程序状态
	Status *string `json:"status,omitempty"`
}

// NewMiniVersionBaseInfo instantiates a new MiniVersionBaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiniVersionBaseInfo() *MiniVersionBaseInfo {
	this := MiniVersionBaseInfo{}
	return &this
}

// NewMiniVersionBaseInfoWithDefaults instantiates a new MiniVersionBaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiniVersionBaseInfoWithDefaults() *MiniVersionBaseInfo {
	this := MiniVersionBaseInfo{}
	return &this
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *MiniVersionBaseInfo) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBuildSource returns the BuildSource field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetBuildSource() string {
	if o == nil || IsNil(o.BuildSource) {
		var ret string
		return ret
	}
	return *o.BuildSource
}

// GetBuildSourceOk returns a tuple with the BuildSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetBuildSourceOk() (*string, bool) {
	if o == nil || IsNil(o.BuildSource) {
		return nil, false
	}
	return o.BuildSource, true
}

// HasBuildSource returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasBuildSource() bool {
	if o != nil && !IsNil(o.BuildSource) {
		return true
	}

	return false
}

// SetBuildSource gets a reference to the given string and assigns it to the BuildSource field.
func (o *MiniVersionBaseInfo) SetBuildSource(v string) {
	o.BuildSource = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *MiniVersionBaseInfo) SetBundleId(v string) {
	o.BundleId = &v
}

// GetDevId returns the DevId field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetDevId() string {
	if o == nil || IsNil(o.DevId) {
		var ret string
		return ret
	}
	return *o.DevId
}

// GetDevIdOk returns a tuple with the DevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetDevIdOk() (*string, bool) {
	if o == nil || IsNil(o.DevId) {
		return nil, false
	}
	return o.DevId, true
}

// HasDevId returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasDevId() bool {
	if o != nil && !IsNil(o.DevId) {
		return true
	}

	return false
}

// SetDevId gets a reference to the given string and assigns it to the DevId field.
func (o *MiniVersionBaseInfo) SetDevId(v string) {
	o.DevId = &v
}

// GetGrayStrategy returns the GrayStrategy field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetGrayStrategy() string {
	if o == nil || IsNil(o.GrayStrategy) {
		var ret string
		return ret
	}
	return *o.GrayStrategy
}

// GetGrayStrategyOk returns a tuple with the GrayStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetGrayStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.GrayStrategy) {
		return nil, false
	}
	return o.GrayStrategy, true
}

// HasGrayStrategy returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasGrayStrategy() bool {
	if o != nil && !IsNil(o.GrayStrategy) {
		return true
	}

	return false
}

// SetGrayStrategy gets a reference to the given string and assigns it to the GrayStrategy field.
func (o *MiniVersionBaseInfo) SetGrayStrategy(v string) {
	o.GrayStrategy = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *MiniVersionBaseInfo) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MiniVersionBaseInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiniVersionBaseInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MiniVersionBaseInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MiniVersionBaseInfo) SetStatus(v string) {
	o.Status = &v
}

func (o MiniVersionBaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiniVersionBaseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BuildSource) {
		toSerialize["build_source"] = o.BuildSource
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.DevId) {
		toSerialize["dev_id"] = o.DevId
	}
	if !IsNil(o.GrayStrategy) {
		toSerialize["gray_strategy"] = o.GrayStrategy
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableMiniVersionBaseInfo struct {
	value *MiniVersionBaseInfo
	isSet bool
}

func (v NullableMiniVersionBaseInfo) Get() *MiniVersionBaseInfo {
	return v.value
}

func (v *NullableMiniVersionBaseInfo) Set(val *MiniVersionBaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMiniVersionBaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMiniVersionBaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiniVersionBaseInfo(val *MiniVersionBaseInfo) *NullableMiniVersionBaseInfo {
	return &NullableMiniVersionBaseInfo{value: val, isSet: true}
}

func (v NullableMiniVersionBaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiniVersionBaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
