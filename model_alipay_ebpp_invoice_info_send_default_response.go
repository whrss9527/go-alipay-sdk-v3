/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayEbppInvoiceInfoSendDefaultResponse struct for AlipayEbppInvoiceInfoSendDefaultResponse
type AlipayEbppInvoiceInfoSendDefaultResponse struct {
	AlipayEbppInvoiceInfoSendErrorResponseModel *AlipayEbppInvoiceInfoSendErrorResponseModel
	CommonErrorType                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceInfoSendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceInfoSendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceInfoSendErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceInfoSendErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceInfoSendErrorResponseModel)
		if string(jsonAlipayEbppInvoiceInfoSendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceInfoSendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceInfoSendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceInfoSendErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
	if err == nil {
		jsonCommonErrorType, _ := json.Marshal(dst.CommonErrorType)
		if string(jsonCommonErrorType) == "{}" { // empty struct
			dst.CommonErrorType = nil
		} else {
			return nil // data stored in dst.CommonErrorType, return on the first match
		}
	} else {
		dst.CommonErrorType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceInfoSendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceInfoSendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceInfoSendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceInfoSendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceInfoSendDefaultResponse struct {
	value *AlipayEbppInvoiceInfoSendDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceInfoSendDefaultResponse) Get() *AlipayEbppInvoiceInfoSendDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInfoSendDefaultResponse) Set(val *AlipayEbppInvoiceInfoSendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInfoSendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInfoSendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInfoSendDefaultResponse(val *AlipayEbppInvoiceInfoSendDefaultResponse) *NullableAlipayEbppInvoiceInfoSendDefaultResponse {
	return &NullableAlipayEbppInvoiceInfoSendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInfoSendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInfoSendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
