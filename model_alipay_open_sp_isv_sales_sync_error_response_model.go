/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AlipayOpenSpIsvSalesSyncErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSpIsvSalesSyncErrorResponseModel{}

// AlipayOpenSpIsvSalesSyncErrorResponseModel struct for AlipayOpenSpIsvSalesSyncErrorResponseModel
type AlipayOpenSpIsvSalesSyncErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayOpenSpIsvSalesSyncErrorResponseModel AlipayOpenSpIsvSalesSyncErrorResponseModel

// NewAlipayOpenSpIsvSalesSyncErrorResponseModel instantiates a new AlipayOpenSpIsvSalesSyncErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSpIsvSalesSyncErrorResponseModel(code string, message string) *AlipayOpenSpIsvSalesSyncErrorResponseModel {
	this := AlipayOpenSpIsvSalesSyncErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayOpenSpIsvSalesSyncErrorResponseModelWithDefaults instantiates a new AlipayOpenSpIsvSalesSyncErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSpIsvSalesSyncErrorResponseModelWithDefaults() *AlipayOpenSpIsvSalesSyncErrorResponseModel {
	this := AlipayOpenSpIsvSalesSyncErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayOpenSpIsvSalesSyncErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSpIsvSalesSyncErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayOpenSpIsvSalesSyncErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayOpenSpIsvSalesSyncErrorResponseModel := _AlipayOpenSpIsvSalesSyncErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayOpenSpIsvSalesSyncErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayOpenSpIsvSalesSyncErrorResponseModel(varAlipayOpenSpIsvSalesSyncErrorResponseModel)

	return err
}

type NullableAlipayOpenSpIsvSalesSyncErrorResponseModel struct {
	value *AlipayOpenSpIsvSalesSyncErrorResponseModel
	isSet bool
}

func (v NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) Get() *AlipayOpenSpIsvSalesSyncErrorResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) Set(val *AlipayOpenSpIsvSalesSyncErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpIsvSalesSyncErrorResponseModel(val *AlipayOpenSpIsvSalesSyncErrorResponseModel) *NullableAlipayOpenSpIsvSalesSyncErrorResponseModel {
	return &NullableAlipayOpenSpIsvSalesSyncErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpIsvSalesSyncErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
