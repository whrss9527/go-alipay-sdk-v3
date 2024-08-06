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


// AlipayOpenMiniInnerversionAuditSubmitDefaultResponse struct for AlipayOpenMiniInnerversionAuditSubmitDefaultResponse
type AlipayOpenMiniInnerversionAuditSubmitDefaultResponse struct {
	AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel *AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionAuditSubmitErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionAuditSubmitDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionAuditSubmitErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse struct {
	value *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) Get() *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) Set(val *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse(val *AlipayOpenMiniInnerversionAuditSubmitDefaultResponse) *NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionAuditSubmitDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


