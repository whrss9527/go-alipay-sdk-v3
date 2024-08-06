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

// checks if the AlipayOpenMiniInnerversionGrayPublishModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionGrayPublishModel{}

// AlipayOpenMiniInnerversionGrayPublishModel struct for AlipayOpenMiniInnerversionGrayPublishModel
type AlipayOpenMiniInnerversionGrayPublishModel struct {
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 小程序版本
	AppVersion *string `json:"app_version,omitempty"`
	// 待邀测应用列表，灰度插件时使用，如无邀测必要请不要传入
	BetaAppIdList []string `json:"beta_app_id_list,omitempty"`
	// 端信息
	BundleId *string `json:"bundle_id,omitempty"`
	// 灰度值
	GrayStrategy *string `json:"gray_strategy,omitempty"`
	// 小程序ID，仅特殊场景使用，普通业务方无需关注该参数
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 小程序所属PID
	Pid *string `json:"pid,omitempty"`
}

// NewAlipayOpenMiniInnerversionGrayPublishModel instantiates a new AlipayOpenMiniInnerversionGrayPublishModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionGrayPublishModel() *AlipayOpenMiniInnerversionGrayPublishModel {
	this := AlipayOpenMiniInnerversionGrayPublishModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionGrayPublishModelWithDefaults instantiates a new AlipayOpenMiniInnerversionGrayPublishModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionGrayPublishModelWithDefaults() *AlipayOpenMiniInnerversionGrayPublishModel {
	this := AlipayOpenMiniInnerversionGrayPublishModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBetaAppIdList returns the BetaAppIdList field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetBetaAppIdList() []string {
	if o == nil || IsNil(o.BetaAppIdList) {
		var ret []string
		return ret
	}
	return o.BetaAppIdList
}

// GetBetaAppIdListOk returns a tuple with the BetaAppIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetBetaAppIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.BetaAppIdList) {
		return nil, false
	}
	return o.BetaAppIdList, true
}

// HasBetaAppIdList returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasBetaAppIdList() bool {
	if o != nil && !IsNil(o.BetaAppIdList) {
		return true
	}

	return false
}

// SetBetaAppIdList gets a reference to the given []string and assigns it to the BetaAppIdList field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetBetaAppIdList(v []string) {
	o.BetaAppIdList = v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetGrayStrategy returns the GrayStrategy field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetGrayStrategy() string {
	if o == nil || IsNil(o.GrayStrategy) {
		var ret string
		return ret
	}
	return *o.GrayStrategy
}

// GetGrayStrategyOk returns a tuple with the GrayStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetGrayStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.GrayStrategy) {
		return nil, false
	}
	return o.GrayStrategy, true
}

// HasGrayStrategy returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasGrayStrategy() bool {
	if o != nil && !IsNil(o.GrayStrategy) {
		return true
	}

	return false
}

// SetGrayStrategy gets a reference to the given string and assigns it to the GrayStrategy field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetGrayStrategy(v string) {
	o.GrayStrategy = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayOpenMiniInnerversionGrayPublishModel) SetPid(v string) {
	o.Pid = &v
}

func (o AlipayOpenMiniInnerversionGrayPublishModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionGrayPublishModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BetaAppIdList) {
		toSerialize["beta_app_id_list"] = o.BetaAppIdList
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.GrayStrategy) {
		toSerialize["gray_strategy"] = o.GrayStrategy
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionGrayPublishModel struct {
	value *AlipayOpenMiniInnerversionGrayPublishModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishModel) Get() *AlipayOpenMiniInnerversionGrayPublishModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishModel) Set(val *AlipayOpenMiniInnerversionGrayPublishModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionGrayPublishModel(val *AlipayOpenMiniInnerversionGrayPublishModel) *NullableAlipayOpenMiniInnerversionGrayPublishModel {
	return &NullableAlipayOpenMiniInnerversionGrayPublishModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


