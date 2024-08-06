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


// AlipayEcoSignFlowCreateDefaultResponse struct for AlipayEcoSignFlowCreateDefaultResponse
type AlipayEcoSignFlowCreateDefaultResponse struct {
	AlipayEcoSignFlowCreateErrorResponseModel *AlipayEcoSignFlowCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEcoSignFlowCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEcoSignFlowCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEcoSignFlowCreateErrorResponseModel);
	if err == nil {
		jsonAlipayEcoSignFlowCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEcoSignFlowCreateErrorResponseModel)
		if string(jsonAlipayEcoSignFlowCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEcoSignFlowCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEcoSignFlowCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEcoSignFlowCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEcoSignFlowCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEcoSignFlowCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEcoSignFlowCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEcoSignFlowCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEcoSignFlowCreateDefaultResponse struct {
	value *AlipayEcoSignFlowCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEcoSignFlowCreateDefaultResponse) Get() *AlipayEcoSignFlowCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEcoSignFlowCreateDefaultResponse) Set(val *AlipayEcoSignFlowCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoSignFlowCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoSignFlowCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoSignFlowCreateDefaultResponse(val *AlipayEcoSignFlowCreateDefaultResponse) *NullableAlipayEcoSignFlowCreateDefaultResponse {
	return &NullableAlipayEcoSignFlowCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEcoSignFlowCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoSignFlowCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


