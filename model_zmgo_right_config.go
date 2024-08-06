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

// checks if the ZMGORightConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZMGORightConfig{}

// ZMGORightConfig struct for ZMGORightConfig
type ZMGORightConfig struct {
	// 芝麻GO管理页已享优惠进度的重定向链接
	CumulativePreferentialRedirectSchema *string `json:"cumulative_preferential_redirect_schema,omitempty"`
	// 权益描述
	CustomBenefitDesc *string `json:"custom_benefit_desc,omitempty"`
	// 权益描述子标题
	CustomSubBenefitDesc *string `json:"custom_sub_benefit_desc,omitempty"`
}

// NewZMGORightConfig instantiates a new ZMGORightConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZMGORightConfig() *ZMGORightConfig {
	this := ZMGORightConfig{}
	return &this
}

// NewZMGORightConfigWithDefaults instantiates a new ZMGORightConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZMGORightConfigWithDefaults() *ZMGORightConfig {
	this := ZMGORightConfig{}
	return &this
}

// GetCumulativePreferentialRedirectSchema returns the CumulativePreferentialRedirectSchema field value if set, zero value otherwise.
func (o *ZMGORightConfig) GetCumulativePreferentialRedirectSchema() string {
	if o == nil || IsNil(o.CumulativePreferentialRedirectSchema) {
		var ret string
		return ret
	}
	return *o.CumulativePreferentialRedirectSchema
}

// GetCumulativePreferentialRedirectSchemaOk returns a tuple with the CumulativePreferentialRedirectSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGORightConfig) GetCumulativePreferentialRedirectSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.CumulativePreferentialRedirectSchema) {
		return nil, false
	}
	return o.CumulativePreferentialRedirectSchema, true
}

// HasCumulativePreferentialRedirectSchema returns a boolean if a field has been set.
func (o *ZMGORightConfig) HasCumulativePreferentialRedirectSchema() bool {
	if o != nil && !IsNil(o.CumulativePreferentialRedirectSchema) {
		return true
	}

	return false
}

// SetCumulativePreferentialRedirectSchema gets a reference to the given string and assigns it to the CumulativePreferentialRedirectSchema field.
func (o *ZMGORightConfig) SetCumulativePreferentialRedirectSchema(v string) {
	o.CumulativePreferentialRedirectSchema = &v
}

// GetCustomBenefitDesc returns the CustomBenefitDesc field value if set, zero value otherwise.
func (o *ZMGORightConfig) GetCustomBenefitDesc() string {
	if o == nil || IsNil(o.CustomBenefitDesc) {
		var ret string
		return ret
	}
	return *o.CustomBenefitDesc
}

// GetCustomBenefitDescOk returns a tuple with the CustomBenefitDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGORightConfig) GetCustomBenefitDescOk() (*string, bool) {
	if o == nil || IsNil(o.CustomBenefitDesc) {
		return nil, false
	}
	return o.CustomBenefitDesc, true
}

// HasCustomBenefitDesc returns a boolean if a field has been set.
func (o *ZMGORightConfig) HasCustomBenefitDesc() bool {
	if o != nil && !IsNil(o.CustomBenefitDesc) {
		return true
	}

	return false
}

// SetCustomBenefitDesc gets a reference to the given string and assigns it to the CustomBenefitDesc field.
func (o *ZMGORightConfig) SetCustomBenefitDesc(v string) {
	o.CustomBenefitDesc = &v
}

// GetCustomSubBenefitDesc returns the CustomSubBenefitDesc field value if set, zero value otherwise.
func (o *ZMGORightConfig) GetCustomSubBenefitDesc() string {
	if o == nil || IsNil(o.CustomSubBenefitDesc) {
		var ret string
		return ret
	}
	return *o.CustomSubBenefitDesc
}

// GetCustomSubBenefitDescOk returns a tuple with the CustomSubBenefitDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZMGORightConfig) GetCustomSubBenefitDescOk() (*string, bool) {
	if o == nil || IsNil(o.CustomSubBenefitDesc) {
		return nil, false
	}
	return o.CustomSubBenefitDesc, true
}

// HasCustomSubBenefitDesc returns a boolean if a field has been set.
func (o *ZMGORightConfig) HasCustomSubBenefitDesc() bool {
	if o != nil && !IsNil(o.CustomSubBenefitDesc) {
		return true
	}

	return false
}

// SetCustomSubBenefitDesc gets a reference to the given string and assigns it to the CustomSubBenefitDesc field.
func (o *ZMGORightConfig) SetCustomSubBenefitDesc(v string) {
	o.CustomSubBenefitDesc = &v
}

func (o ZMGORightConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZMGORightConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CumulativePreferentialRedirectSchema) {
		toSerialize["cumulative_preferential_redirect_schema"] = o.CumulativePreferentialRedirectSchema
	}
	if !IsNil(o.CustomBenefitDesc) {
		toSerialize["custom_benefit_desc"] = o.CustomBenefitDesc
	}
	if !IsNil(o.CustomSubBenefitDesc) {
		toSerialize["custom_sub_benefit_desc"] = o.CustomSubBenefitDesc
	}
	return toSerialize, nil
}

type NullableZMGORightConfig struct {
	value *ZMGORightConfig
	isSet bool
}

func (v NullableZMGORightConfig) Get() *ZMGORightConfig {
	return v.value
}

func (v *NullableZMGORightConfig) Set(val *ZMGORightConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableZMGORightConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableZMGORightConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZMGORightConfig(val *ZMGORightConfig) *NullableZMGORightConfig {
	return &NullableZMGORightConfig{value: val, isSet: true}
}

func (v NullableZMGORightConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZMGORightConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

