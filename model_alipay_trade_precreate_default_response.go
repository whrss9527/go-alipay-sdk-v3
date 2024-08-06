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


// AlipayTradePrecreateDefaultResponse struct for AlipayTradePrecreateDefaultResponse
type AlipayTradePrecreateDefaultResponse struct {
	AlipayTradePrecreateErrorResponseModel *AlipayTradePrecreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradePrecreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradePrecreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradePrecreateErrorResponseModel);
	if err == nil {
		jsonAlipayTradePrecreateErrorResponseModel, _ := json.Marshal(dst.AlipayTradePrecreateErrorResponseModel)
		if string(jsonAlipayTradePrecreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradePrecreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradePrecreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradePrecreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradePrecreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradePrecreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradePrecreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradePrecreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayTradePrecreateDefaultResponse struct {
	value *AlipayTradePrecreateDefaultResponse
	isSet bool
}

func (v NullableAlipayTradePrecreateDefaultResponse) Get() *AlipayTradePrecreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradePrecreateDefaultResponse) Set(val *AlipayTradePrecreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradePrecreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradePrecreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradePrecreateDefaultResponse(val *AlipayTradePrecreateDefaultResponse) *NullableAlipayTradePrecreateDefaultResponse {
	return &NullableAlipayTradePrecreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradePrecreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradePrecreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


