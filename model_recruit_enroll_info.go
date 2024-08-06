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

// checks if the RecruitEnrollInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecruitEnrollInfo{}

// RecruitEnrollInfo struct for RecruitEnrollInfo
type RecruitEnrollInfo struct {
	// 报名提交的活动城市码，详见<a href=\"https://opendocs.alipay.com/pre-open/02r07u\" target=\"_blank\">活动城市码</a>
	Cities []string `json:"cities,omitempty"`
	EnrollMerchant *RecruitEnrollMerchant `json:"enroll_merchant,omitempty"`
	// 报名提交的素材
	Materials []RecruitMaterial `json:"materials,omitempty"`
	// 报名提交的小程序信息，是否必选取决于方案要求
	MiniApps []RecruitMiniApp `json:"mini_apps,omitempty"`
	// 报名提交的券信息，是否必选取决于方案要求
	Vouchers []RecruitVoucher `json:"vouchers,omitempty"`
}

// NewRecruitEnrollInfo instantiates a new RecruitEnrollInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecruitEnrollInfo() *RecruitEnrollInfo {
	this := RecruitEnrollInfo{}
	return &this
}

// NewRecruitEnrollInfoWithDefaults instantiates a new RecruitEnrollInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecruitEnrollInfoWithDefaults() *RecruitEnrollInfo {
	this := RecruitEnrollInfo{}
	return &this
}

// GetCities returns the Cities field value if set, zero value otherwise.
func (o *RecruitEnrollInfo) GetCities() []string {
	if o == nil || IsNil(o.Cities) {
		var ret []string
		return ret
	}
	return o.Cities
}

// GetCitiesOk returns a tuple with the Cities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollInfo) GetCitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cities) {
		return nil, false
	}
	return o.Cities, true
}

// HasCities returns a boolean if a field has been set.
func (o *RecruitEnrollInfo) HasCities() bool {
	if o != nil && !IsNil(o.Cities) {
		return true
	}

	return false
}

// SetCities gets a reference to the given []string and assigns it to the Cities field.
func (o *RecruitEnrollInfo) SetCities(v []string) {
	o.Cities = v
}

// GetEnrollMerchant returns the EnrollMerchant field value if set, zero value otherwise.
func (o *RecruitEnrollInfo) GetEnrollMerchant() RecruitEnrollMerchant {
	if o == nil || IsNil(o.EnrollMerchant) {
		var ret RecruitEnrollMerchant
		return ret
	}
	return *o.EnrollMerchant
}

// GetEnrollMerchantOk returns a tuple with the EnrollMerchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollInfo) GetEnrollMerchantOk() (*RecruitEnrollMerchant, bool) {
	if o == nil || IsNil(o.EnrollMerchant) {
		return nil, false
	}
	return o.EnrollMerchant, true
}

// HasEnrollMerchant returns a boolean if a field has been set.
func (o *RecruitEnrollInfo) HasEnrollMerchant() bool {
	if o != nil && !IsNil(o.EnrollMerchant) {
		return true
	}

	return false
}

// SetEnrollMerchant gets a reference to the given RecruitEnrollMerchant and assigns it to the EnrollMerchant field.
func (o *RecruitEnrollInfo) SetEnrollMerchant(v RecruitEnrollMerchant) {
	o.EnrollMerchant = &v
}

// GetMaterials returns the Materials field value if set, zero value otherwise.
func (o *RecruitEnrollInfo) GetMaterials() []RecruitMaterial {
	if o == nil || IsNil(o.Materials) {
		var ret []RecruitMaterial
		return ret
	}
	return o.Materials
}

// GetMaterialsOk returns a tuple with the Materials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollInfo) GetMaterialsOk() ([]RecruitMaterial, bool) {
	if o == nil || IsNil(o.Materials) {
		return nil, false
	}
	return o.Materials, true
}

// HasMaterials returns a boolean if a field has been set.
func (o *RecruitEnrollInfo) HasMaterials() bool {
	if o != nil && !IsNil(o.Materials) {
		return true
	}

	return false
}

// SetMaterials gets a reference to the given []RecruitMaterial and assigns it to the Materials field.
func (o *RecruitEnrollInfo) SetMaterials(v []RecruitMaterial) {
	o.Materials = v
}

// GetMiniApps returns the MiniApps field value if set, zero value otherwise.
func (o *RecruitEnrollInfo) GetMiniApps() []RecruitMiniApp {
	if o == nil || IsNil(o.MiniApps) {
		var ret []RecruitMiniApp
		return ret
	}
	return o.MiniApps
}

// GetMiniAppsOk returns a tuple with the MiniApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollInfo) GetMiniAppsOk() ([]RecruitMiniApp, bool) {
	if o == nil || IsNil(o.MiniApps) {
		return nil, false
	}
	return o.MiniApps, true
}

// HasMiniApps returns a boolean if a field has been set.
func (o *RecruitEnrollInfo) HasMiniApps() bool {
	if o != nil && !IsNil(o.MiniApps) {
		return true
	}

	return false
}

// SetMiniApps gets a reference to the given []RecruitMiniApp and assigns it to the MiniApps field.
func (o *RecruitEnrollInfo) SetMiniApps(v []RecruitMiniApp) {
	o.MiniApps = v
}

// GetVouchers returns the Vouchers field value if set, zero value otherwise.
func (o *RecruitEnrollInfo) GetVouchers() []RecruitVoucher {
	if o == nil || IsNil(o.Vouchers) {
		var ret []RecruitVoucher
		return ret
	}
	return o.Vouchers
}

// GetVouchersOk returns a tuple with the Vouchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollInfo) GetVouchersOk() ([]RecruitVoucher, bool) {
	if o == nil || IsNil(o.Vouchers) {
		return nil, false
	}
	return o.Vouchers, true
}

// HasVouchers returns a boolean if a field has been set.
func (o *RecruitEnrollInfo) HasVouchers() bool {
	if o != nil && !IsNil(o.Vouchers) {
		return true
	}

	return false
}

// SetVouchers gets a reference to the given []RecruitVoucher and assigns it to the Vouchers field.
func (o *RecruitEnrollInfo) SetVouchers(v []RecruitVoucher) {
	o.Vouchers = v
}

func (o RecruitEnrollInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecruitEnrollInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cities) {
		toSerialize["cities"] = o.Cities
	}
	if !IsNil(o.EnrollMerchant) {
		toSerialize["enroll_merchant"] = o.EnrollMerchant
	}
	if !IsNil(o.Materials) {
		toSerialize["materials"] = o.Materials
	}
	if !IsNil(o.MiniApps) {
		toSerialize["mini_apps"] = o.MiniApps
	}
	if !IsNil(o.Vouchers) {
		toSerialize["vouchers"] = o.Vouchers
	}
	return toSerialize, nil
}

type NullableRecruitEnrollInfo struct {
	value *RecruitEnrollInfo
	isSet bool
}

func (v NullableRecruitEnrollInfo) Get() *RecruitEnrollInfo {
	return v.value
}

func (v *NullableRecruitEnrollInfo) Set(val *RecruitEnrollInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRecruitEnrollInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRecruitEnrollInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecruitEnrollInfo(val *RecruitEnrollInfo) *NullableRecruitEnrollInfo {
	return &NullableRecruitEnrollInfo{value: val, isSet: true}
}

func (v NullableRecruitEnrollInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecruitEnrollInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


