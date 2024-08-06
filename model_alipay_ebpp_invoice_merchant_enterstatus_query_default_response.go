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

// AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse struct for AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse
type AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse struct {
	AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel *AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel
	CommonErrorType                                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceMerchantEnterstatusQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse struct {
	value *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) Get() *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) Set(val *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse(val *AlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) *NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceMerchantEnterstatusQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
