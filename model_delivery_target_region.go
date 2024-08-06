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

// checks if the DeliveryTargetRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryTargetRegion{}

// DeliveryTargetRegion struct for DeliveryTargetRegion
type DeliveryTargetRegion struct {
	// 区域编码，参考<a href=\"https://opendocs.alipay.com/pre-open/03144c\" target=\"_blank\">取值说明</a>
	RegionCode *string `json:"region_code,omitempty"`
	// 区域名，参考<a href=\"https://opendocs.alipay.com/pre-open/03144c\" target=\"_blank\">取值说明</a>
	RegionName *string `json:"region_name,omitempty"`
	// 区域类型，1代表全国，2代表省级区域，3代表市级区域，参考<a href=\"https://opendocs.alipay.com/pre-open/03144c\" target=\"_blank\">取值说明</a>
	RegionType *string `json:"region_type,omitempty"`
}

// NewDeliveryTargetRegion instantiates a new DeliveryTargetRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryTargetRegion() *DeliveryTargetRegion {
	this := DeliveryTargetRegion{}
	return &this
}

// NewDeliveryTargetRegionWithDefaults instantiates a new DeliveryTargetRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryTargetRegionWithDefaults() *DeliveryTargetRegion {
	this := DeliveryTargetRegion{}
	return &this
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *DeliveryTargetRegion) GetRegionCode() string {
	if o == nil || IsNil(o.RegionCode) {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTargetRegion) GetRegionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionCode) {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *DeliveryTargetRegion) HasRegionCode() bool {
	if o != nil && !IsNil(o.RegionCode) {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *DeliveryTargetRegion) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *DeliveryTargetRegion) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTargetRegion) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DeliveryTargetRegion) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DeliveryTargetRegion) SetRegionName(v string) {
	o.RegionName = &v
}

// GetRegionType returns the RegionType field value if set, zero value otherwise.
func (o *DeliveryTargetRegion) GetRegionType() string {
	if o == nil || IsNil(o.RegionType) {
		var ret string
		return ret
	}
	return *o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryTargetRegion) GetRegionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RegionType) {
		return nil, false
	}
	return o.RegionType, true
}

// HasRegionType returns a boolean if a field has been set.
func (o *DeliveryTargetRegion) HasRegionType() bool {
	if o != nil && !IsNil(o.RegionType) {
		return true
	}

	return false
}

// SetRegionType gets a reference to the given string and assigns it to the RegionType field.
func (o *DeliveryTargetRegion) SetRegionType(v string) {
	o.RegionType = &v
}

func (o DeliveryTargetRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryTargetRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegionCode) {
		toSerialize["region_code"] = o.RegionCode
	}
	if !IsNil(o.RegionName) {
		toSerialize["region_name"] = o.RegionName
	}
	if !IsNil(o.RegionType) {
		toSerialize["region_type"] = o.RegionType
	}
	return toSerialize, nil
}

type NullableDeliveryTargetRegion struct {
	value *DeliveryTargetRegion
	isSet bool
}

func (v NullableDeliveryTargetRegion) Get() *DeliveryTargetRegion {
	return v.value
}

func (v *NullableDeliveryTargetRegion) Set(val *DeliveryTargetRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryTargetRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryTargetRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryTargetRegion(val *DeliveryTargetRegion) *NullableDeliveryTargetRegion {
	return &NullableDeliveryTargetRegion{value: val, isSet: true}
}

func (v NullableDeliveryTargetRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryTargetRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

