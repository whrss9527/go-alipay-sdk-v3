/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse struct for AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse
type AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse struct {
	AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel *AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel);
	if err == nil {
		jsonAlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel, _ := json.Marshal(dst.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel)
		if string(jsonAlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel) == "{}" { // empty struct
			dst.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel != nil {
		return json.Marshal(&src.AlipaySecurityRiskVerifyidentityMiniappConfirmErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse struct {
	value *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse
	isSet bool
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) Get() *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse {
	return v.value
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) Set(val *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse(val *AlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse {
	return &NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySecurityRiskVerifyidentityMiniappConfirmDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


