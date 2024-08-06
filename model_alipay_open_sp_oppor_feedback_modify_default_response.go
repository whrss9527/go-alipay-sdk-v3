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

// AlipayOpenSpOpporFeedbackModifyDefaultResponse struct for AlipayOpenSpOpporFeedbackModifyDefaultResponse
type AlipayOpenSpOpporFeedbackModifyDefaultResponse struct {
	AlipayOpenSpOpporFeedbackModifyErrorResponseModel *AlipayOpenSpOpporFeedbackModifyErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpOpporFeedbackModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpOpporFeedbackModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpOpporFeedbackModifyErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSpOpporFeedbackModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpOpporFeedbackModifyErrorResponseModel)
		if string(jsonAlipayOpenSpOpporFeedbackModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpOpporFeedbackModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpOpporFeedbackModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpOpporFeedbackModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpOpporFeedbackModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpOpporFeedbackModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpOpporFeedbackModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpOpporFeedbackModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse struct {
	value *AlipayOpenSpOpporFeedbackModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) Get() *AlipayOpenSpOpporFeedbackModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) Set(val *AlipayOpenSpOpporFeedbackModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpOpporFeedbackModifyDefaultResponse(val *AlipayOpenSpOpporFeedbackModifyDefaultResponse) *NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse {
	return &NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpOpporFeedbackModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
