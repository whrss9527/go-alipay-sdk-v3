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

// checks if the AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel{}

// AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel struct for AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel struct {
	// 返回码 10000-成功 其他都是失败
	Code *string `json:"code,omitempty"`
	// 返回消息
	Msg *string `json:"msg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty"`
	// 抬头ID
	TitleId *string `json:"title_id,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) SetMsg(v string) {
	o.Msg = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

// GetTitleId returns the TitleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetTitleId() string {
	if o == nil || IsNil(o.TitleId) {
		var ret string
		return ret
	}
	return *o.TitleId
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) GetTitleIdOk() (*string, bool) {
	if o == nil || IsNil(o.TitleId) {
		return nil, false
	}
	return o.TitleId, true
}

// HasTitleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) HasTitleId() bool {
	if o != nil && !IsNil(o.TitleId) {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given string and assigns it to the TitleId field.
func (o *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) SetTitleId(v string) {
	o.TitleId = &v
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.TitleId) {
		toSerialize["title_id"] = o.TitleId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) Get() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) Set(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


