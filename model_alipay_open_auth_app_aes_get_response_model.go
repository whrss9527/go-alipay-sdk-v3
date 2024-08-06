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

// checks if the AlipayOpenAuthAppAesGetResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenAuthAppAesGetResponseModel{}

// AlipayOpenAuthAppAesGetResponseModel struct for AlipayOpenAuthAppAesGetResponseModel
type AlipayOpenAuthAppAesGetResponseModel struct {
	// 商家应用AES KEY密文，需要ISV使用三方应用配置的AES密钥内容进行解密。详情参见 <a href=\"https://opendocs.alipay.com/isv/grefvl/getaes\">应用AES密钥管理</a>。
	AesKey *string `json:"aes_key,omitempty"`
}

// NewAlipayOpenAuthAppAesGetResponseModel instantiates a new AlipayOpenAuthAppAesGetResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenAuthAppAesGetResponseModel() *AlipayOpenAuthAppAesGetResponseModel {
	this := AlipayOpenAuthAppAesGetResponseModel{}
	return &this
}

// NewAlipayOpenAuthAppAesGetResponseModelWithDefaults instantiates a new AlipayOpenAuthAppAesGetResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenAuthAppAesGetResponseModelWithDefaults() *AlipayOpenAuthAppAesGetResponseModel {
	this := AlipayOpenAuthAppAesGetResponseModel{}
	return &this
}

// GetAesKey returns the AesKey field value if set, zero value otherwise.
func (o *AlipayOpenAuthAppAesGetResponseModel) GetAesKey() string {
	if o == nil || IsNil(o.AesKey) {
		var ret string
		return ret
	}
	return *o.AesKey
}

// GetAesKeyOk returns a tuple with the AesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenAuthAppAesGetResponseModel) GetAesKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AesKey) {
		return nil, false
	}
	return o.AesKey, true
}

// HasAesKey returns a boolean if a field has been set.
func (o *AlipayOpenAuthAppAesGetResponseModel) HasAesKey() bool {
	if o != nil && !IsNil(o.AesKey) {
		return true
	}

	return false
}

// SetAesKey gets a reference to the given string and assigns it to the AesKey field.
func (o *AlipayOpenAuthAppAesGetResponseModel) SetAesKey(v string) {
	o.AesKey = &v
}

func (o AlipayOpenAuthAppAesGetResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenAuthAppAesGetResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AesKey) {
		toSerialize["aes_key"] = o.AesKey
	}
	return toSerialize, nil
}

type NullableAlipayOpenAuthAppAesGetResponseModel struct {
	value *AlipayOpenAuthAppAesGetResponseModel
	isSet bool
}

func (v NullableAlipayOpenAuthAppAesGetResponseModel) Get() *AlipayOpenAuthAppAesGetResponseModel {
	return v.value
}

func (v *NullableAlipayOpenAuthAppAesGetResponseModel) Set(val *AlipayOpenAuthAppAesGetResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAuthAppAesGetResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAuthAppAesGetResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAuthAppAesGetResponseModel(val *AlipayOpenAuthAppAesGetResponseModel) *NullableAlipayOpenAuthAppAesGetResponseModel {
	return &NullableAlipayOpenAuthAppAesGetResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenAuthAppAesGetResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAuthAppAesGetResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
