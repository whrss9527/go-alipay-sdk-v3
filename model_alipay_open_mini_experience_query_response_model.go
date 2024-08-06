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

// checks if the AlipayOpenMiniExperienceQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniExperienceQueryResponseModel{}

// AlipayOpenMiniExperienceQueryResponseModel struct for AlipayOpenMiniExperienceQueryResponseModel
type AlipayOpenMiniExperienceQueryResponseModel struct {
	// 小程序体验版二维码地址
	ExpQrCodeUrl *string `json:"exp_qr_code_url,omitempty"`
	// 体验版schema
	ExpSchemaUrl *string `json:"exp_schema_url,omitempty"`
	// 体验版打包状态。
	Status *string `json:"status,omitempty"`
}

// NewAlipayOpenMiniExperienceQueryResponseModel instantiates a new AlipayOpenMiniExperienceQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniExperienceQueryResponseModel() *AlipayOpenMiniExperienceQueryResponseModel {
	this := AlipayOpenMiniExperienceQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniExperienceQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniExperienceQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniExperienceQueryResponseModelWithDefaults() *AlipayOpenMiniExperienceQueryResponseModel {
	this := AlipayOpenMiniExperienceQueryResponseModel{}
	return &this
}

// GetExpQrCodeUrl returns the ExpQrCodeUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetExpQrCodeUrl() string {
	if o == nil || IsNil(o.ExpQrCodeUrl) {
		var ret string
		return ret
	}
	return *o.ExpQrCodeUrl
}

// GetExpQrCodeUrlOk returns a tuple with the ExpQrCodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetExpQrCodeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExpQrCodeUrl) {
		return nil, false
	}
	return o.ExpQrCodeUrl, true
}

// HasExpQrCodeUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) HasExpQrCodeUrl() bool {
	if o != nil && !IsNil(o.ExpQrCodeUrl) {
		return true
	}

	return false
}

// SetExpQrCodeUrl gets a reference to the given string and assigns it to the ExpQrCodeUrl field.
func (o *AlipayOpenMiniExperienceQueryResponseModel) SetExpQrCodeUrl(v string) {
	o.ExpQrCodeUrl = &v
}

// GetExpSchemaUrl returns the ExpSchemaUrl field value if set, zero value otherwise.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetExpSchemaUrl() string {
	if o == nil || IsNil(o.ExpSchemaUrl) {
		var ret string
		return ret
	}
	return *o.ExpSchemaUrl
}

// GetExpSchemaUrlOk returns a tuple with the ExpSchemaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetExpSchemaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExpSchemaUrl) {
		return nil, false
	}
	return o.ExpSchemaUrl, true
}

// HasExpSchemaUrl returns a boolean if a field has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) HasExpSchemaUrl() bool {
	if o != nil && !IsNil(o.ExpSchemaUrl) {
		return true
	}

	return false
}

// SetExpSchemaUrl gets a reference to the given string and assigns it to the ExpSchemaUrl field.
func (o *AlipayOpenMiniExperienceQueryResponseModel) SetExpSchemaUrl(v string) {
	o.ExpSchemaUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayOpenMiniExperienceQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayOpenMiniExperienceQueryResponseModel) SetStatus(v string) {
	o.Status = &v
}

func (o AlipayOpenMiniExperienceQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniExperienceQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpQrCodeUrl) {
		toSerialize["exp_qr_code_url"] = o.ExpQrCodeUrl
	}
	if !IsNil(o.ExpSchemaUrl) {
		toSerialize["exp_schema_url"] = o.ExpSchemaUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniExperienceQueryResponseModel struct {
	value *AlipayOpenMiniExperienceQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniExperienceQueryResponseModel) Get() *AlipayOpenMiniExperienceQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniExperienceQueryResponseModel) Set(val *AlipayOpenMiniExperienceQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniExperienceQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniExperienceQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniExperienceQueryResponseModel(val *AlipayOpenMiniExperienceQueryResponseModel) *NullableAlipayOpenMiniExperienceQueryResponseModel {
	return &NullableAlipayOpenMiniExperienceQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniExperienceQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniExperienceQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
