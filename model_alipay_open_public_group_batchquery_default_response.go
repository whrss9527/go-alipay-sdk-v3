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

// AlipayOpenPublicGroupBatchqueryDefaultResponse struct for AlipayOpenPublicGroupBatchqueryDefaultResponse
type AlipayOpenPublicGroupBatchqueryDefaultResponse struct {
	AlipayOpenPublicGroupBatchqueryErrorResponseModel *AlipayOpenPublicGroupBatchqueryErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicGroupBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicGroupBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicGroupBatchqueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicGroupBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicGroupBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenPublicGroupBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicGroupBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicGroupBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicGroupBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicGroupBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicGroupBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicGroupBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicGroupBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicGroupBatchqueryDefaultResponse struct {
	value *AlipayOpenPublicGroupBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) Get() *AlipayOpenPublicGroupBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) Set(val *AlipayOpenPublicGroupBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicGroupBatchqueryDefaultResponse(val *AlipayOpenPublicGroupBatchqueryDefaultResponse) *NullableAlipayOpenPublicGroupBatchqueryDefaultResponse {
	return &NullableAlipayOpenPublicGroupBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicGroupBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
