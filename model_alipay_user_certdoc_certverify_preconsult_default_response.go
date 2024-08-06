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

// AlipayUserCertdocCertverifyPreconsultDefaultResponse struct for AlipayUserCertdocCertverifyPreconsultDefaultResponse
type AlipayUserCertdocCertverifyPreconsultDefaultResponse struct {
	AlipayUserCertdocCertverifyPreconsultErrorResponseModel *AlipayUserCertdocCertverifyPreconsultErrorResponseModel
	CommonErrorType                                         *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayUserCertdocCertverifyPreconsultDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayUserCertdocCertverifyPreconsultErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayUserCertdocCertverifyPreconsultErrorResponseModel)
	if err == nil {
		jsonAlipayUserCertdocCertverifyPreconsultErrorResponseModel, _ := json.Marshal(dst.AlipayUserCertdocCertverifyPreconsultErrorResponseModel)
		if string(jsonAlipayUserCertdocCertverifyPreconsultErrorResponseModel) == "{}" { // empty struct
			dst.AlipayUserCertdocCertverifyPreconsultErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayUserCertdocCertverifyPreconsultErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayUserCertdocCertverifyPreconsultErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayUserCertdocCertverifyPreconsultDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayUserCertdocCertverifyPreconsultDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayUserCertdocCertverifyPreconsultErrorResponseModel != nil {
		return json.Marshal(&src.AlipayUserCertdocCertverifyPreconsultErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse struct {
	value *AlipayUserCertdocCertverifyPreconsultDefaultResponse
	isSet bool
}

func (v NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) Get() *AlipayUserCertdocCertverifyPreconsultDefaultResponse {
	return v.value
}

func (v *NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) Set(val *AlipayUserCertdocCertverifyPreconsultDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserCertdocCertverifyPreconsultDefaultResponse(val *AlipayUserCertdocCertverifyPreconsultDefaultResponse) *NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse {
	return &NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserCertdocCertverifyPreconsultDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
