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

// checks if the FaceExtParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaceExtParams{}

// FaceExtParams struct for FaceExtParams
type FaceExtParams struct {
	// 业务类型：7，基于1:N人脸搜索的刷脸支付场景；8，基于姓名和身份证号的刷脸支付场景。
	BizType *string `json:"biz_type,omitempty"`
}

// NewFaceExtParams instantiates a new FaceExtParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaceExtParams() *FaceExtParams {
	this := FaceExtParams{}
	return &this
}

// NewFaceExtParamsWithDefaults instantiates a new FaceExtParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaceExtParamsWithDefaults() *FaceExtParams {
	this := FaceExtParams{}
	return &this
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *FaceExtParams) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaceExtParams) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *FaceExtParams) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *FaceExtParams) SetBizType(v string) {
	o.BizType = &v
}

func (o FaceExtParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaceExtParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	return toSerialize, nil
}

type NullableFaceExtParams struct {
	value *FaceExtParams
	isSet bool
}

func (v NullableFaceExtParams) Get() *FaceExtParams {
	return v.value
}

func (v *NullableFaceExtParams) Set(val *FaceExtParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFaceExtParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFaceExtParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaceExtParams(val *FaceExtParams) *NullableFaceExtParams {
	return &NullableFaceExtParams{value: val, isSet: true}
}

func (v NullableFaceExtParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaceExtParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
