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

// checks if the AlipayDataDataserviceAdConversionUploadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayDataDataserviceAdConversionUploadModel{}

// AlipayDataDataserviceAdConversionUploadModel struct for AlipayDataDataserviceAdConversionUploadModel
type AlipayDataDataserviceAdConversionUploadModel struct {
	// 代理商访问灯火平台的token
	BizToken *string `json:"biz_token,omitempty"`
	// 转化数据列表, 最多1000条
	ConversionDataList []ConversionData `json:"conversion_data_list,omitempty"`
}

// NewAlipayDataDataserviceAdConversionUploadModel instantiates a new AlipayDataDataserviceAdConversionUploadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayDataDataserviceAdConversionUploadModel() *AlipayDataDataserviceAdConversionUploadModel {
	this := AlipayDataDataserviceAdConversionUploadModel{}
	return &this
}

// NewAlipayDataDataserviceAdConversionUploadModelWithDefaults instantiates a new AlipayDataDataserviceAdConversionUploadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayDataDataserviceAdConversionUploadModelWithDefaults() *AlipayDataDataserviceAdConversionUploadModel {
	this := AlipayDataDataserviceAdConversionUploadModel{}
	return &this
}

// GetBizToken returns the BizToken field value if set, zero value otherwise.
func (o *AlipayDataDataserviceAdConversionUploadModel) GetBizToken() string {
	if o == nil || IsNil(o.BizToken) {
		var ret string
		return ret
	}
	return *o.BizToken
}

// GetBizTokenOk returns a tuple with the BizToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataDataserviceAdConversionUploadModel) GetBizTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BizToken) {
		return nil, false
	}
	return o.BizToken, true
}

// HasBizToken returns a boolean if a field has been set.
func (o *AlipayDataDataserviceAdConversionUploadModel) HasBizToken() bool {
	if o != nil && !IsNil(o.BizToken) {
		return true
	}

	return false
}

// SetBizToken gets a reference to the given string and assigns it to the BizToken field.
func (o *AlipayDataDataserviceAdConversionUploadModel) SetBizToken(v string) {
	o.BizToken = &v
}

// GetConversionDataList returns the ConversionDataList field value if set, zero value otherwise.
func (o *AlipayDataDataserviceAdConversionUploadModel) GetConversionDataList() []ConversionData {
	if o == nil || IsNil(o.ConversionDataList) {
		var ret []ConversionData
		return ret
	}
	return o.ConversionDataList
}

// GetConversionDataListOk returns a tuple with the ConversionDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayDataDataserviceAdConversionUploadModel) GetConversionDataListOk() ([]ConversionData, bool) {
	if o == nil || IsNil(o.ConversionDataList) {
		return nil, false
	}
	return o.ConversionDataList, true
}

// HasConversionDataList returns a boolean if a field has been set.
func (o *AlipayDataDataserviceAdConversionUploadModel) HasConversionDataList() bool {
	if o != nil && !IsNil(o.ConversionDataList) {
		return true
	}

	return false
}

// SetConversionDataList gets a reference to the given []ConversionData and assigns it to the ConversionDataList field.
func (o *AlipayDataDataserviceAdConversionUploadModel) SetConversionDataList(v []ConversionData) {
	o.ConversionDataList = v
}

func (o AlipayDataDataserviceAdConversionUploadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayDataDataserviceAdConversionUploadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizToken) {
		toSerialize["biz_token"] = o.BizToken
	}
	if !IsNil(o.ConversionDataList) {
		toSerialize["conversion_data_list"] = o.ConversionDataList
	}
	return toSerialize, nil
}

type NullableAlipayDataDataserviceAdConversionUploadModel struct {
	value *AlipayDataDataserviceAdConversionUploadModel
	isSet bool
}

func (v NullableAlipayDataDataserviceAdConversionUploadModel) Get() *AlipayDataDataserviceAdConversionUploadModel {
	return v.value
}

func (v *NullableAlipayDataDataserviceAdConversionUploadModel) Set(val *AlipayDataDataserviceAdConversionUploadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataDataserviceAdConversionUploadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataDataserviceAdConversionUploadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataDataserviceAdConversionUploadModel(val *AlipayDataDataserviceAdConversionUploadModel) *NullableAlipayDataDataserviceAdConversionUploadModel {
	return &NullableAlipayDataDataserviceAdConversionUploadModel{value: val, isSet: true}
}

func (v NullableAlipayDataDataserviceAdConversionUploadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataDataserviceAdConversionUploadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
