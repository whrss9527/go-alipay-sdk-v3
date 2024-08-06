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

// AlipayUserAgreementExecutionplanModifyDefaultResponse struct for AlipayUserAgreementExecutionplanModifyDefaultResponse
type AlipayUserAgreementExecutionplanModifyDefaultResponse struct {
	AlipayUserAgreementExecutionplanModifyErrorResponseModel *AlipayUserAgreementExecutionplanModifyErrorResponseModel
	CommonErrorType                                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayUserAgreementExecutionplanModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayUserAgreementExecutionplanModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayUserAgreementExecutionplanModifyErrorResponseModel)
	if err == nil {
		jsonAlipayUserAgreementExecutionplanModifyErrorResponseModel, _ := json.Marshal(dst.AlipayUserAgreementExecutionplanModifyErrorResponseModel)
		if string(jsonAlipayUserAgreementExecutionplanModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayUserAgreementExecutionplanModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayUserAgreementExecutionplanModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayUserAgreementExecutionplanModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayUserAgreementExecutionplanModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayUserAgreementExecutionplanModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayUserAgreementExecutionplanModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayUserAgreementExecutionplanModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayUserAgreementExecutionplanModifyDefaultResponse struct {
	value *AlipayUserAgreementExecutionplanModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) Get() *AlipayUserAgreementExecutionplanModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) Set(val *AlipayUserAgreementExecutionplanModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementExecutionplanModifyDefaultResponse(val *AlipayUserAgreementExecutionplanModifyDefaultResponse) *NullableAlipayUserAgreementExecutionplanModifyDefaultResponse {
	return &NullableAlipayUserAgreementExecutionplanModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementExecutionplanModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
