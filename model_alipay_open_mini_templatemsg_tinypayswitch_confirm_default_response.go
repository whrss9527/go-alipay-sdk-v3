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

// AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse struct for AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse
type AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse struct {
	AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel *AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel
	CommonErrorType                                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel)
		if string(jsonAlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniTemplatemsgTinypayswitchConfirmErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse struct {
	value *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) Get() *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) Set(val *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse(val *AlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) *NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse {
	return &NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniTemplatemsgTinypayswitchConfirmDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
