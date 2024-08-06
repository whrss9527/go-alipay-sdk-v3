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

// AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse struct for AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse
type AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse struct {
	AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel *AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel
	CommonErrorType                                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionTemplatedUploadErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse struct {
	value *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) Get() *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) Set(val *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse(val *AlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) *NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionTemplatedUploadDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
