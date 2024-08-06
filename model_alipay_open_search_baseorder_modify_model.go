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

// checks if the AlipayOpenSearchBaseorderModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBaseorderModifyModel{}

// AlipayOpenSearchBaseorderModifyModel struct for AlipayOpenSearchBaseorderModifyModel
type AlipayOpenSearchBaseorderModifyModel struct {
	BizData *SearchBaseOrderCreateApiRequest `json:"biz_data,omitempty"`
	// 表示修改内容的业务类型
	BizType *string `json:"biz_type,omitempty"`
	// 操作的具体类型
	OptType *string `json:"opt_type,omitempty"`
}

// NewAlipayOpenSearchBaseorderModifyModel instantiates a new AlipayOpenSearchBaseorderModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBaseorderModifyModel() *AlipayOpenSearchBaseorderModifyModel {
	this := AlipayOpenSearchBaseorderModifyModel{}
	return &this
}

// NewAlipayOpenSearchBaseorderModifyModelWithDefaults instantiates a new AlipayOpenSearchBaseorderModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBaseorderModifyModelWithDefaults() *AlipayOpenSearchBaseorderModifyModel {
	this := AlipayOpenSearchBaseorderModifyModel{}
	return &this
}

// GetBizData returns the BizData field value if set, zero value otherwise.
func (o *AlipayOpenSearchBaseorderModifyModel) GetBizData() SearchBaseOrderCreateApiRequest {
	if o == nil || IsNil(o.BizData) {
		var ret SearchBaseOrderCreateApiRequest
		return ret
	}
	return *o.BizData
}

// GetBizDataOk returns a tuple with the BizData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) GetBizDataOk() (*SearchBaseOrderCreateApiRequest, bool) {
	if o == nil || IsNil(o.BizData) {
		return nil, false
	}
	return o.BizData, true
}

// HasBizData returns a boolean if a field has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) HasBizData() bool {
	if o != nil && !IsNil(o.BizData) {
		return true
	}

	return false
}

// SetBizData gets a reference to the given SearchBaseOrderCreateApiRequest and assigns it to the BizData field.
func (o *AlipayOpenSearchBaseorderModifyModel) SetBizData(v SearchBaseOrderCreateApiRequest) {
	o.BizData = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *AlipayOpenSearchBaseorderModifyModel) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *AlipayOpenSearchBaseorderModifyModel) SetBizType(v string) {
	o.BizType = &v
}

// GetOptType returns the OptType field value if set, zero value otherwise.
func (o *AlipayOpenSearchBaseorderModifyModel) GetOptType() string {
	if o == nil || IsNil(o.OptType) {
		var ret string
		return ret
	}
	return *o.OptType
}

// GetOptTypeOk returns a tuple with the OptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) GetOptTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OptType) {
		return nil, false
	}
	return o.OptType, true
}

// HasOptType returns a boolean if a field has been set.
func (o *AlipayOpenSearchBaseorderModifyModel) HasOptType() bool {
	if o != nil && !IsNil(o.OptType) {
		return true
	}

	return false
}

// SetOptType gets a reference to the given string and assigns it to the OptType field.
func (o *AlipayOpenSearchBaseorderModifyModel) SetOptType(v string) {
	o.OptType = &v
}

func (o AlipayOpenSearchBaseorderModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBaseorderModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizData) {
		toSerialize["biz_data"] = o.BizData
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.OptType) {
		toSerialize["opt_type"] = o.OptType
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBaseorderModifyModel struct {
	value *AlipayOpenSearchBaseorderModifyModel
	isSet bool
}

func (v NullableAlipayOpenSearchBaseorderModifyModel) Get() *AlipayOpenSearchBaseorderModifyModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBaseorderModifyModel) Set(val *AlipayOpenSearchBaseorderModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBaseorderModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBaseorderModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBaseorderModifyModel(val *AlipayOpenSearchBaseorderModifyModel) *NullableAlipayOpenSearchBaseorderModifyModel {
	return &NullableAlipayOpenSearchBaseorderModifyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBaseorderModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBaseorderModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


