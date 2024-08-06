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


// AlipayEcoEduKtParentQueryDefaultResponse struct for AlipayEcoEduKtParentQueryDefaultResponse
type AlipayEcoEduKtParentQueryDefaultResponse struct {
	AlipayEcoEduKtParentQueryErrorResponseModel *AlipayEcoEduKtParentQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoEduKtParentQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoEduKtParentQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoEduKtParentQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEcoEduKtParentQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEcoEduKtParentQueryErrorResponseModel)
		if string(jsonAlipayEcoEduKtParentQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoEduKtParentQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoEduKtParentQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoEduKtParentQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoEduKtParentQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoEduKtParentQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoEduKtParentQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoEduKtParentQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoEduKtParentQueryDefaultResponse struct {
	value *AlipayEcoEduKtParentQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoEduKtParentQueryDefaultResponse) Get() *AlipayEcoEduKtParentQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoEduKtParentQueryDefaultResponse) Set(val *AlipayEcoEduKtParentQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoEduKtParentQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoEduKtParentQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoEduKtParentQueryDefaultResponse(val *AlipayEcoEduKtParentQueryDefaultResponse) *NullableAlipayEcoEduKtParentQueryDefaultResponse {
	return &NullableAlipayEcoEduKtParentQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoEduKtParentQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoEduKtParentQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

