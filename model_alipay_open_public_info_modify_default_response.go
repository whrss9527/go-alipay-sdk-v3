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

// AlipayOpenPublicInfoModifyDefaultResponse struct for AlipayOpenPublicInfoModifyDefaultResponse
type AlipayOpenPublicInfoModifyDefaultResponse struct {
	AlipayOpenPublicInfoModifyErrorResponseModel *AlipayOpenPublicInfoModifyErrorResponseModel
	CommonErrorType                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicInfoModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicInfoModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicInfoModifyErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicInfoModifyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicInfoModifyErrorResponseModel)
		if string(jsonAlipayOpenPublicInfoModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicInfoModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicInfoModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicInfoModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicInfoModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicInfoModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicInfoModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicInfoModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicInfoModifyDefaultResponse struct {
	value *AlipayOpenPublicInfoModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicInfoModifyDefaultResponse) Get() *AlipayOpenPublicInfoModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicInfoModifyDefaultResponse) Set(val *AlipayOpenPublicInfoModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicInfoModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicInfoModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicInfoModifyDefaultResponse(val *AlipayOpenPublicInfoModifyDefaultResponse) *NullableAlipayOpenPublicInfoModifyDefaultResponse {
	return &NullableAlipayOpenPublicInfoModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicInfoModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicInfoModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
