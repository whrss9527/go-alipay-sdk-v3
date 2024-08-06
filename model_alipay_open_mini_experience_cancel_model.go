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

// checks if the AlipayOpenMiniExperienceCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniExperienceCancelModel{}

// AlipayOpenMiniExperienceCancelModel struct for AlipayOpenMiniExperienceCancelModel
type AlipayOpenMiniExperienceCancelModel struct {
	// 小程序版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 小程序客户端类型，默认为支付宝端。常见支持如下客户端： com.alipay.alipaywallet：支付宝端； com.alibaba.android.rimet：DINGDING端； com.amap.app：高德端； com.alibaba.ailabs.genie.webapps：天猫精灵端； com.alipay.iot.xpaas：支付宝IoT端。 如需更多端投放，请联系业务BD。
	BundleId *string `json:"bundle_id,omitempty"`
}

// NewAlipayOpenMiniExperienceCancelModel instantiates a new AlipayOpenMiniExperienceCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniExperienceCancelModel() *AlipayOpenMiniExperienceCancelModel {
	this := AlipayOpenMiniExperienceCancelModel{}
	return &this
}

// NewAlipayOpenMiniExperienceCancelModelWithDefaults instantiates a new AlipayOpenMiniExperienceCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniExperienceCancelModelWithDefaults() *AlipayOpenMiniExperienceCancelModel {
	this := AlipayOpenMiniExperienceCancelModel{}
	return &this
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniExperienceCancelModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniExperienceCancelModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniExperienceCancelModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniExperienceCancelModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniExperienceCancelModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniExperienceCancelModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniExperienceCancelModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniExperienceCancelModel) SetBundleId(v string) {
	o.BundleId = &v
}

func (o AlipayOpenMiniExperienceCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniExperienceCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniExperienceCancelModel struct {
	value *AlipayOpenMiniExperienceCancelModel
	isSet bool
}

func (v NullableAlipayOpenMiniExperienceCancelModel) Get() *AlipayOpenMiniExperienceCancelModel {
	return v.value
}

func (v *NullableAlipayOpenMiniExperienceCancelModel) Set(val *AlipayOpenMiniExperienceCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniExperienceCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniExperienceCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniExperienceCancelModel(val *AlipayOpenMiniExperienceCancelModel) *NullableAlipayOpenMiniExperienceCancelModel {
	return &NullableAlipayOpenMiniExperienceCancelModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniExperienceCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniExperienceCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


