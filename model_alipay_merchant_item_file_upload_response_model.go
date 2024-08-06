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

// checks if the AlipayMerchantItemFileUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantItemFileUploadResponseModel{}

// AlipayMerchantItemFileUploadResponseModel struct for AlipayMerchantItemFileUploadResponseModel
type AlipayMerchantItemFileUploadResponseModel struct {
	// 文件在商品中心的素材标识（素材ID长期有效）
	MaterialId *string `json:"material_id,omitempty"`
	// 文件在商品中心的素材标示，创建/更新商品时使用
	MaterialKey *string `json:"material_key,omitempty"`
}

// NewAlipayMerchantItemFileUploadResponseModel instantiates a new AlipayMerchantItemFileUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantItemFileUploadResponseModel() *AlipayMerchantItemFileUploadResponseModel {
	this := AlipayMerchantItemFileUploadResponseModel{}
	return &this
}

// NewAlipayMerchantItemFileUploadResponseModelWithDefaults instantiates a new AlipayMerchantItemFileUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantItemFileUploadResponseModelWithDefaults() *AlipayMerchantItemFileUploadResponseModel {
	this := AlipayMerchantItemFileUploadResponseModel{}
	return &this
}

// GetMaterialId returns the MaterialId field value if set, zero value otherwise.
func (o *AlipayMerchantItemFileUploadResponseModel) GetMaterialId() string {
	if o == nil || IsNil(o.MaterialId) {
		var ret string
		return ret
	}
	return *o.MaterialId
}

// GetMaterialIdOk returns a tuple with the MaterialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantItemFileUploadResponseModel) GetMaterialIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialId) {
		return nil, false
	}
	return o.MaterialId, true
}

// HasMaterialId returns a boolean if a field has been set.
func (o *AlipayMerchantItemFileUploadResponseModel) HasMaterialId() bool {
	if o != nil && !IsNil(o.MaterialId) {
		return true
	}

	return false
}

// SetMaterialId gets a reference to the given string and assigns it to the MaterialId field.
func (o *AlipayMerchantItemFileUploadResponseModel) SetMaterialId(v string) {
	o.MaterialId = &v
}

// GetMaterialKey returns the MaterialKey field value if set, zero value otherwise.
func (o *AlipayMerchantItemFileUploadResponseModel) GetMaterialKey() string {
	if o == nil || IsNil(o.MaterialKey) {
		var ret string
		return ret
	}
	return *o.MaterialKey
}

// GetMaterialKeyOk returns a tuple with the MaterialKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantItemFileUploadResponseModel) GetMaterialKeyOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialKey) {
		return nil, false
	}
	return o.MaterialKey, true
}

// HasMaterialKey returns a boolean if a field has been set.
func (o *AlipayMerchantItemFileUploadResponseModel) HasMaterialKey() bool {
	if o != nil && !IsNil(o.MaterialKey) {
		return true
	}

	return false
}

// SetMaterialKey gets a reference to the given string and assigns it to the MaterialKey field.
func (o *AlipayMerchantItemFileUploadResponseModel) SetMaterialKey(v string) {
	o.MaterialKey = &v
}

func (o AlipayMerchantItemFileUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantItemFileUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaterialId) {
		toSerialize["material_id"] = o.MaterialId
	}
	if !IsNil(o.MaterialKey) {
		toSerialize["material_key"] = o.MaterialKey
	}
	return toSerialize, nil
}

type NullableAlipayMerchantItemFileUploadResponseModel struct {
	value *AlipayMerchantItemFileUploadResponseModel
	isSet bool
}

func (v NullableAlipayMerchantItemFileUploadResponseModel) Get() *AlipayMerchantItemFileUploadResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantItemFileUploadResponseModel) Set(val *AlipayMerchantItemFileUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantItemFileUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantItemFileUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantItemFileUploadResponseModel(val *AlipayMerchantItemFileUploadResponseModel) *NullableAlipayMerchantItemFileUploadResponseModel {
	return &NullableAlipayMerchantItemFileUploadResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantItemFileUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantItemFileUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

