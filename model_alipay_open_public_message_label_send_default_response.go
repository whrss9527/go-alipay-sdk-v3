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

// AlipayOpenPublicMessageLabelSendDefaultResponse struct for AlipayOpenPublicMessageLabelSendDefaultResponse
type AlipayOpenPublicMessageLabelSendDefaultResponse struct {
	AlipayOpenPublicMessageLabelSendErrorResponseModel *AlipayOpenPublicMessageLabelSendErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicMessageLabelSendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicMessageLabelSendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicMessageLabelSendErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicMessageLabelSendErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicMessageLabelSendErrorResponseModel)
		if string(jsonAlipayOpenPublicMessageLabelSendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicMessageLabelSendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicMessageLabelSendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicMessageLabelSendErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicMessageLabelSendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicMessageLabelSendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicMessageLabelSendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicMessageLabelSendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicMessageLabelSendDefaultResponse struct {
	value *AlipayOpenPublicMessageLabelSendDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicMessageLabelSendDefaultResponse) Get() *AlipayOpenPublicMessageLabelSendDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicMessageLabelSendDefaultResponse) Set(val *AlipayOpenPublicMessageLabelSendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMessageLabelSendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMessageLabelSendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMessageLabelSendDefaultResponse(val *AlipayOpenPublicMessageLabelSendDefaultResponse) *NullableAlipayOpenPublicMessageLabelSendDefaultResponse {
	return &NullableAlipayOpenPublicMessageLabelSendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMessageLabelSendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMessageLabelSendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
