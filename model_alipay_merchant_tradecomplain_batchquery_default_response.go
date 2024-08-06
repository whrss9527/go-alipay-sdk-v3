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

// AlipayMerchantTradecomplainBatchqueryDefaultResponse struct for AlipayMerchantTradecomplainBatchqueryDefaultResponse
type AlipayMerchantTradecomplainBatchqueryDefaultResponse struct {
	AlipayMerchantTradecomplainBatchqueryErrorResponseModel *AlipayMerchantTradecomplainBatchqueryErrorResponseModel
	CommonErrorType                                         *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMerchantTradecomplainBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMerchantTradecomplainBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMerchantTradecomplainBatchqueryErrorResponseModel)
	if err == nil {
		jsonAlipayMerchantTradecomplainBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayMerchantTradecomplainBatchqueryErrorResponseModel)
		if string(jsonAlipayMerchantTradecomplainBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMerchantTradecomplainBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMerchantTradecomplainBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMerchantTradecomplainBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMerchantTradecomplainBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMerchantTradecomplainBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMerchantTradecomplainBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMerchantTradecomplainBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse struct {
	value *AlipayMerchantTradecomplainBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) Get() *AlipayMerchantTradecomplainBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) Set(val *AlipayMerchantTradecomplainBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantTradecomplainBatchqueryDefaultResponse(val *AlipayMerchantTradecomplainBatchqueryDefaultResponse) *NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse {
	return &NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantTradecomplainBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
