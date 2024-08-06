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

// checks if the AlipayOpenMiniInnerversionUpgradeModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerversionUpgradeModel{}

// AlipayOpenMiniInnerversionUpgradeModel struct for AlipayOpenMiniInnerversionUpgradeModel
type AlipayOpenMiniInnerversionUpgradeModel struct {
	// 来源类型，新接入方需要向支付宝申请专用来源，否则不予接入。
	AppOrigin *string `json:"app_origin,omitempty"`
	// 端ID，多端场景下区分不同端
	BundleId *string `json:"bundle_id,omitempty"`
	// 小程序ID，特殊场景专用，普通业务方无需关注该参数。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// 用于升级的模板id
	TemplateId *string `json:"template_id,omitempty"`
	// 用于升级的模板版本号，版本号必须满足 x.y.z, 且均为数字
	TemplateVersion *string `json:"template_version,omitempty"`
}

// NewAlipayOpenMiniInnerversionUpgradeModel instantiates a new AlipayOpenMiniInnerversionUpgradeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerversionUpgradeModel() *AlipayOpenMiniInnerversionUpgradeModel {
	this := AlipayOpenMiniInnerversionUpgradeModel{}
	return &this
}

// NewAlipayOpenMiniInnerversionUpgradeModelWithDefaults instantiates a new AlipayOpenMiniInnerversionUpgradeModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerversionUpgradeModelWithDefaults() *AlipayOpenMiniInnerversionUpgradeModel {
	this := AlipayOpenMiniInnerversionUpgradeModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerversionUpgradeModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerversionUpgradeModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerversionUpgradeModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *AlipayOpenMiniInnerversionUpgradeModel) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateVersion returns the TemplateVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetTemplateVersion() string {
	if o == nil || IsNil(o.TemplateVersion) {
		var ret string
		return ret
	}
	return *o.TemplateVersion
}

// GetTemplateVersionOk returns a tuple with the TemplateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) GetTemplateVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateVersion) {
		return nil, false
	}
	return o.TemplateVersion, true
}

// HasTemplateVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerversionUpgradeModel) HasTemplateVersion() bool {
	if o != nil && !IsNil(o.TemplateVersion) {
		return true
	}

	return false
}

// SetTemplateVersion gets a reference to the given string and assigns it to the TemplateVersion field.
func (o *AlipayOpenMiniInnerversionUpgradeModel) SetTemplateVersion(v string) {
	o.TemplateVersion = &v
}

func (o AlipayOpenMiniInnerversionUpgradeModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerversionUpgradeModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.TemplateVersion) {
		toSerialize["template_version"] = o.TemplateVersion
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerversionUpgradeModel struct {
	value *AlipayOpenMiniInnerversionUpgradeModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionUpgradeModel) Get() *AlipayOpenMiniInnerversionUpgradeModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionUpgradeModel) Set(val *AlipayOpenMiniInnerversionUpgradeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionUpgradeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionUpgradeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionUpgradeModel(val *AlipayOpenMiniInnerversionUpgradeModel) *NullableAlipayOpenMiniInnerversionUpgradeModel {
	return &NullableAlipayOpenMiniInnerversionUpgradeModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionUpgradeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionUpgradeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
