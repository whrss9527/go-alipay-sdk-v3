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

// checks if the AlipayOpenMiniInnerExperienceCancelModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerExperienceCancelModel{}

// AlipayOpenMiniInnerExperienceCancelModel struct for AlipayOpenMiniInnerExperienceCancelModel
type AlipayOpenMiniInnerExperienceCancelModel struct {
	// 业务来源
	AppOrigin *string `json:"app_origin,omitempty"`
	// 小程序版本号
	AppVersion *string `json:"app_version,omitempty"`
	// 端id
	BundleId *string `json:"bundle_id,omitempty"`
	// 业务小程序appid
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// oid
	Oid *string `json:"oid,omitempty"`
	// 操作员id
	OperatorId *string `json:"operator_id,omitempty"`
}

// NewAlipayOpenMiniInnerExperienceCancelModel instantiates a new AlipayOpenMiniInnerExperienceCancelModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerExperienceCancelModel() *AlipayOpenMiniInnerExperienceCancelModel {
	this := AlipayOpenMiniInnerExperienceCancelModel{}
	return &this
}

// NewAlipayOpenMiniInnerExperienceCancelModelWithDefaults instantiates a new AlipayOpenMiniInnerExperienceCancelModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerExperienceCancelModelWithDefaults() *AlipayOpenMiniInnerExperienceCancelModel {
	this := AlipayOpenMiniInnerExperienceCancelModel{}
	return &this
}

// GetAppOrigin returns the AppOrigin field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetAppOrigin() string {
	if o == nil || IsNil(o.AppOrigin) {
		var ret string
		return ret
	}
	return *o.AppOrigin
}

// GetAppOriginOk returns a tuple with the AppOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetAppOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AppOrigin) {
		return nil, false
	}
	return o.AppOrigin, true
}

// HasAppOrigin returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasAppOrigin() bool {
	if o != nil && !IsNil(o.AppOrigin) {
		return true
	}

	return false
}

// SetAppOrigin gets a reference to the given string and assigns it to the AppOrigin field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetAppOrigin(v string) {
	o.AppOrigin = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetBundleId(v string) {
	o.BundleId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

// GetOid returns the Oid field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetOid() string {
	if o == nil || IsNil(o.Oid) {
		var ret string
		return ret
	}
	return *o.Oid
}

// GetOidOk returns a tuple with the Oid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetOidOk() (*string, bool) {
	if o == nil || IsNil(o.Oid) {
		return nil, false
	}
	return o.Oid, true
}

// HasOid returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasOid() bool {
	if o != nil && !IsNil(o.Oid) {
		return true
	}

	return false
}

// SetOid gets a reference to the given string and assigns it to the Oid field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetOid(v string) {
	o.Oid = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetOperatorId() string {
	if o == nil || IsNil(o.OperatorId) {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) GetOperatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperatorId) {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerExperienceCancelModel) HasOperatorId() bool {
	if o != nil && !IsNil(o.OperatorId) {
		return true
	}

	return false
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *AlipayOpenMiniInnerExperienceCancelModel) SetOperatorId(v string) {
	o.OperatorId = &v
}

func (o AlipayOpenMiniInnerExperienceCancelModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerExperienceCancelModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppOrigin) {
		toSerialize["app_origin"] = o.AppOrigin
	}
	if !IsNil(o.AppVersion) {
		toSerialize["app_version"] = o.AppVersion
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	if !IsNil(o.Oid) {
		toSerialize["oid"] = o.Oid
	}
	if !IsNil(o.OperatorId) {
		toSerialize["operator_id"] = o.OperatorId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerExperienceCancelModel struct {
	value *AlipayOpenMiniInnerExperienceCancelModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerExperienceCancelModel) Get() *AlipayOpenMiniInnerExperienceCancelModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerExperienceCancelModel) Set(val *AlipayOpenMiniInnerExperienceCancelModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerExperienceCancelModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerExperienceCancelModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerExperienceCancelModel(val *AlipayOpenMiniInnerExperienceCancelModel) *NullableAlipayOpenMiniInnerExperienceCancelModel {
	return &NullableAlipayOpenMiniInnerExperienceCancelModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerExperienceCancelModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerExperienceCancelModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


