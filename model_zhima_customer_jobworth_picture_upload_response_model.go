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

// checks if the ZhimaCustomerJobworthPictureUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZhimaCustomerJobworthPictureUploadResponseModel{}

// ZhimaCustomerJobworthPictureUploadResponseModel struct for ZhimaCustomerJobworthPictureUploadResponseModel
type ZhimaCustomerJobworthPictureUploadResponseModel struct {
	// 图片返回ID
	PicId *string `json:"pic_id,omitempty"`
	// 错误码
	SubCode *string `json:"sub_code,omitempty"`
	// 错误原因
	SubMsg *string `json:"sub_msg,omitempty"`
}

// NewZhimaCustomerJobworthPictureUploadResponseModel instantiates a new ZhimaCustomerJobworthPictureUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZhimaCustomerJobworthPictureUploadResponseModel() *ZhimaCustomerJobworthPictureUploadResponseModel {
	this := ZhimaCustomerJobworthPictureUploadResponseModel{}
	return &this
}

// NewZhimaCustomerJobworthPictureUploadResponseModelWithDefaults instantiates a new ZhimaCustomerJobworthPictureUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZhimaCustomerJobworthPictureUploadResponseModelWithDefaults() *ZhimaCustomerJobworthPictureUploadResponseModel {
	this := ZhimaCustomerJobworthPictureUploadResponseModel{}
	return &this
}

// GetPicId returns the PicId field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetPicId() string {
	if o == nil || IsNil(o.PicId) {
		var ret string
		return ret
	}
	return *o.PicId
}

// GetPicIdOk returns a tuple with the PicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetPicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PicId) {
		return nil, false
	}
	return o.PicId, true
}

// HasPicId returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) HasPicId() bool {
	if o != nil && !IsNil(o.PicId) {
		return true
	}

	return false
}

// SetPicId gets a reference to the given string and assigns it to the PicId field.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) SetPicId(v string) {
	o.PicId = &v
}

// GetSubCode returns the SubCode field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetSubCode() string {
	if o == nil || IsNil(o.SubCode) {
		var ret string
		return ret
	}
	return *o.SubCode
}

// GetSubCodeOk returns a tuple with the SubCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetSubCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SubCode) {
		return nil, false
	}
	return o.SubCode, true
}

// HasSubCode returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) HasSubCode() bool {
	if o != nil && !IsNil(o.SubCode) {
		return true
	}

	return false
}

// SetSubCode gets a reference to the given string and assigns it to the SubCode field.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) SetSubCode(v string) {
	o.SubCode = &v
}

// GetSubMsg returns the SubMsg field value if set, zero value otherwise.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetSubMsg() string {
	if o == nil || IsNil(o.SubMsg) {
		var ret string
		return ret
	}
	return *o.SubMsg
}

// GetSubMsgOk returns a tuple with the SubMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) GetSubMsgOk() (*string, bool) {
	if o == nil || IsNil(o.SubMsg) {
		return nil, false
	}
	return o.SubMsg, true
}

// HasSubMsg returns a boolean if a field has been set.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) HasSubMsg() bool {
	if o != nil && !IsNil(o.SubMsg) {
		return true
	}

	return false
}

// SetSubMsg gets a reference to the given string and assigns it to the SubMsg field.
func (o *ZhimaCustomerJobworthPictureUploadResponseModel) SetSubMsg(v string) {
	o.SubMsg = &v
}

func (o ZhimaCustomerJobworthPictureUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZhimaCustomerJobworthPictureUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PicId) {
		toSerialize["pic_id"] = o.PicId
	}
	if !IsNil(o.SubCode) {
		toSerialize["sub_code"] = o.SubCode
	}
	if !IsNil(o.SubMsg) {
		toSerialize["sub_msg"] = o.SubMsg
	}
	return toSerialize, nil
}

type NullableZhimaCustomerJobworthPictureUploadResponseModel struct {
	value *ZhimaCustomerJobworthPictureUploadResponseModel
	isSet bool
}

func (v NullableZhimaCustomerJobworthPictureUploadResponseModel) Get() *ZhimaCustomerJobworthPictureUploadResponseModel {
	return v.value
}

func (v *NullableZhimaCustomerJobworthPictureUploadResponseModel) Set(val *ZhimaCustomerJobworthPictureUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCustomerJobworthPictureUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCustomerJobworthPictureUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCustomerJobworthPictureUploadResponseModel(val *ZhimaCustomerJobworthPictureUploadResponseModel) *NullableZhimaCustomerJobworthPictureUploadResponseModel {
	return &NullableZhimaCustomerJobworthPictureUploadResponseModel{value: val, isSet: true}
}

func (v NullableZhimaCustomerJobworthPictureUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCustomerJobworthPictureUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
