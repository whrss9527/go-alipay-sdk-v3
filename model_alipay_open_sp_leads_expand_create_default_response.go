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


// AlipayOpenSpLeadsExpandCreateDefaultResponse struct for AlipayOpenSpLeadsExpandCreateDefaultResponse
type AlipayOpenSpLeadsExpandCreateDefaultResponse struct {
	AlipayOpenSpLeadsExpandCreateErrorResponseModel *AlipayOpenSpLeadsExpandCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpLeadsExpandCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpLeadsExpandCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpLeadsExpandCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSpLeadsExpandCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpLeadsExpandCreateErrorResponseModel)
		if string(jsonAlipayOpenSpLeadsExpandCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpLeadsExpandCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpLeadsExpandCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpLeadsExpandCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpLeadsExpandCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpLeadsExpandCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpLeadsExpandCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpLeadsExpandCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSpLeadsExpandCreateDefaultResponse struct {
	value *AlipayOpenSpLeadsExpandCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) Get() *AlipayOpenSpLeadsExpandCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) Set(val *AlipayOpenSpLeadsExpandCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpLeadsExpandCreateDefaultResponse(val *AlipayOpenSpLeadsExpandCreateDefaultResponse) *NullableAlipayOpenSpLeadsExpandCreateDefaultResponse {
	return &NullableAlipayOpenSpLeadsExpandCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpLeadsExpandCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

