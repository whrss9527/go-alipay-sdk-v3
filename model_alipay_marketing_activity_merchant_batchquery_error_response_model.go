/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AlipayMarketingActivityMerchantBatchqueryErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityMerchantBatchqueryErrorResponseModel{}

// AlipayMarketingActivityMerchantBatchqueryErrorResponseModel struct for AlipayMarketingActivityMerchantBatchqueryErrorResponseModel
type AlipayMarketingActivityMerchantBatchqueryErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayMarketingActivityMerchantBatchqueryErrorResponseModel AlipayMarketingActivityMerchantBatchqueryErrorResponseModel

// NewAlipayMarketingActivityMerchantBatchqueryErrorResponseModel instantiates a new AlipayMarketingActivityMerchantBatchqueryErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityMerchantBatchqueryErrorResponseModel(code string, message string) *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel {
	this := AlipayMarketingActivityMerchantBatchqueryErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayMarketingActivityMerchantBatchqueryErrorResponseModelWithDefaults instantiates a new AlipayMarketingActivityMerchantBatchqueryErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityMerchantBatchqueryErrorResponseModelWithDefaults() *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel {
	this := AlipayMarketingActivityMerchantBatchqueryErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayMarketingActivityMerchantBatchqueryErrorResponseModel := _AlipayMarketingActivityMerchantBatchqueryErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayMarketingActivityMerchantBatchqueryErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayMarketingActivityMerchantBatchqueryErrorResponseModel(varAlipayMarketingActivityMerchantBatchqueryErrorResponseModel)

	return err
}

type NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel struct {
	value *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) Get() *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) Set(val *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel(val *AlipayMarketingActivityMerchantBatchqueryErrorResponseModel) *NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel {
	return &NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityMerchantBatchqueryErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


