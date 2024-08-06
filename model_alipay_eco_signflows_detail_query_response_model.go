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

// checks if the AlipayEcoSignflowsDetailQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoSignflowsDetailQueryResponseModel{}

// AlipayEcoSignflowsDetailQueryResponseModel struct for AlipayEcoSignflowsDetailQueryResponseModel
type AlipayEcoSignflowsDetailQueryResponseModel struct {
	Attachments *AttachmentDetail `json:"attachments,omitempty"`
	Docs        *DocInfo          `json:"docs,omitempty"`
}

// NewAlipayEcoSignflowsDetailQueryResponseModel instantiates a new AlipayEcoSignflowsDetailQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoSignflowsDetailQueryResponseModel() *AlipayEcoSignflowsDetailQueryResponseModel {
	this := AlipayEcoSignflowsDetailQueryResponseModel{}
	return &this
}

// NewAlipayEcoSignflowsDetailQueryResponseModelWithDefaults instantiates a new AlipayEcoSignflowsDetailQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoSignflowsDetailQueryResponseModelWithDefaults() *AlipayEcoSignflowsDetailQueryResponseModel {
	this := AlipayEcoSignflowsDetailQueryResponseModel{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) GetAttachments() AttachmentDetail {
	if o == nil || IsNil(o.Attachments) {
		var ret AttachmentDetail
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) GetAttachmentsOk() (*AttachmentDetail, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AttachmentDetail and assigns it to the Attachments field.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) SetAttachments(v AttachmentDetail) {
	o.Attachments = &v
}

// GetDocs returns the Docs field value if set, zero value otherwise.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) GetDocs() DocInfo {
	if o == nil || IsNil(o.Docs) {
		var ret DocInfo
		return ret
	}
	return *o.Docs
}

// GetDocsOk returns a tuple with the Docs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) GetDocsOk() (*DocInfo, bool) {
	if o == nil || IsNil(o.Docs) {
		return nil, false
	}
	return o.Docs, true
}

// HasDocs returns a boolean if a field has been set.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) HasDocs() bool {
	if o != nil && !IsNil(o.Docs) {
		return true
	}

	return false
}

// SetDocs gets a reference to the given DocInfo and assigns it to the Docs field.
func (o *AlipayEcoSignflowsDetailQueryResponseModel) SetDocs(v DocInfo) {
	o.Docs = &v
}

func (o AlipayEcoSignflowsDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoSignflowsDetailQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Docs) {
		toSerialize["docs"] = o.Docs
	}
	return toSerialize, nil
}

type NullableAlipayEcoSignflowsDetailQueryResponseModel struct {
	value *AlipayEcoSignflowsDetailQueryResponseModel
	isSet bool
}

func (v NullableAlipayEcoSignflowsDetailQueryResponseModel) Get() *AlipayEcoSignflowsDetailQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEcoSignflowsDetailQueryResponseModel) Set(val *AlipayEcoSignflowsDetailQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoSignflowsDetailQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoSignflowsDetailQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoSignflowsDetailQueryResponseModel(val *AlipayEcoSignflowsDetailQueryResponseModel) *NullableAlipayEcoSignflowsDetailQueryResponseModel {
	return &NullableAlipayEcoSignflowsDetailQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEcoSignflowsDetailQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoSignflowsDetailQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
