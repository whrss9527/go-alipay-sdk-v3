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

// checks if the JourneyMerchantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JourneyMerchantInfo{}

// JourneyMerchantInfo struct for JourneyMerchantInfo
type JourneyMerchantInfo struct {
	// 扩展信息
	ExtInfo []OrderExtInfo `json:"ext_info,omitempty"`
	// 商户logo链接
	Logo *string `json:"logo,omitempty"`
	// 商家名称
	Name *string `json:"name,omitempty"`
	// 商家简称
	ShortName *string `json:"short_name,omitempty"`
}

// NewJourneyMerchantInfo instantiates a new JourneyMerchantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJourneyMerchantInfo() *JourneyMerchantInfo {
	this := JourneyMerchantInfo{}
	return &this
}

// NewJourneyMerchantInfoWithDefaults instantiates a new JourneyMerchantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJourneyMerchantInfoWithDefaults() *JourneyMerchantInfo {
	this := JourneyMerchantInfo{}
	return &this
}

// GetExtInfo returns the ExtInfo field value if set, zero value otherwise.
func (o *JourneyMerchantInfo) GetExtInfo() []OrderExtInfo {
	if o == nil || IsNil(o.ExtInfo) {
		var ret []OrderExtInfo
		return ret
	}
	return o.ExtInfo
}

// GetExtInfoOk returns a tuple with the ExtInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMerchantInfo) GetExtInfoOk() ([]OrderExtInfo, bool) {
	if o == nil || IsNil(o.ExtInfo) {
		return nil, false
	}
	return o.ExtInfo, true
}

// HasExtInfo returns a boolean if a field has been set.
func (o *JourneyMerchantInfo) HasExtInfo() bool {
	if o != nil && !IsNil(o.ExtInfo) {
		return true
	}

	return false
}

// SetExtInfo gets a reference to the given []OrderExtInfo and assigns it to the ExtInfo field.
func (o *JourneyMerchantInfo) SetExtInfo(v []OrderExtInfo) {
	o.ExtInfo = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *JourneyMerchantInfo) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMerchantInfo) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *JourneyMerchantInfo) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *JourneyMerchantInfo) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JourneyMerchantInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMerchantInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JourneyMerchantInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JourneyMerchantInfo) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *JourneyMerchantInfo) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JourneyMerchantInfo) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *JourneyMerchantInfo) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *JourneyMerchantInfo) SetShortName(v string) {
	o.ShortName = &v
}

func (o JourneyMerchantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JourneyMerchantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtInfo) {
		toSerialize["ext_info"] = o.ExtInfo
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	return toSerialize, nil
}

type NullableJourneyMerchantInfo struct {
	value *JourneyMerchantInfo
	isSet bool
}

func (v NullableJourneyMerchantInfo) Get() *JourneyMerchantInfo {
	return v.value
}

func (v *NullableJourneyMerchantInfo) Set(val *JourneyMerchantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableJourneyMerchantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableJourneyMerchantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJourneyMerchantInfo(val *JourneyMerchantInfo) *NullableJourneyMerchantInfo {
	return &NullableJourneyMerchantInfo{value: val, isSet: true}
}

func (v NullableJourneyMerchantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJourneyMerchantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


