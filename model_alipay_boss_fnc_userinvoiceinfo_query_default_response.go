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


// AlipayBossFncUserinvoiceinfoQueryDefaultResponse struct for AlipayBossFncUserinvoiceinfoQueryDefaultResponse
type AlipayBossFncUserinvoiceinfoQueryDefaultResponse struct {
	AlipayBossFncUserinvoiceinfoQueryErrorResponseModel *AlipayBossFncUserinvoiceinfoQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayBossFncUserinvoiceinfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayBossFncUserinvoiceinfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel);
	if err == nil {
		jsonAlipayBossFncUserinvoiceinfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel)
		if string(jsonAlipayBossFncUserinvoiceinfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayBossFncUserinvoiceinfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayBossFncUserinvoiceinfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayBossFncUserinvoiceinfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse struct {
	value *AlipayBossFncUserinvoiceinfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) Get() *AlipayBossFncUserinvoiceinfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) Set(val *AlipayBossFncUserinvoiceinfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse(val *AlipayBossFncUserinvoiceinfoQueryDefaultResponse) *NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse {
	return &NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayBossFncUserinvoiceinfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


