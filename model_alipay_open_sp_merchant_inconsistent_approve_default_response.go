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

// AlipayOpenSpMerchantInconsistentApproveDefaultResponse struct for AlipayOpenSpMerchantInconsistentApproveDefaultResponse
type AlipayOpenSpMerchantInconsistentApproveDefaultResponse struct {
	AlipayOpenSpMerchantInconsistentApproveErrorResponseModel *AlipayOpenSpMerchantInconsistentApproveErrorResponseModel
	CommonErrorType                                           *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpMerchantInconsistentApproveDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpMerchantInconsistentApproveErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSpMerchantInconsistentApproveErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel)
		if string(jsonAlipayOpenSpMerchantInconsistentApproveErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpMerchantInconsistentApproveDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpMerchantInconsistentApproveDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpMerchantInconsistentApproveErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse struct {
	value *AlipayOpenSpMerchantInconsistentApproveDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) Get() *AlipayOpenSpMerchantInconsistentApproveDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) Set(val *AlipayOpenSpMerchantInconsistentApproveDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse(val *AlipayOpenSpMerchantInconsistentApproveDefaultResponse) *NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse {
	return &NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpMerchantInconsistentApproveDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
