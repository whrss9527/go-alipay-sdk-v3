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


// AlipayFundAuthOperationCancelDefaultResponse struct for AlipayFundAuthOperationCancelDefaultResponse
type AlipayFundAuthOperationCancelDefaultResponse struct {
	AlipayFundAuthOperationCancelErrorResponseModel *AlipayFundAuthOperationCancelErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundAuthOperationCancelDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundAuthOperationCancelErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundAuthOperationCancelErrorResponseModel);
	if err == nil {
		jsonAlipayFundAuthOperationCancelErrorResponseModel, _ := json.Marshal(dst.AlipayFundAuthOperationCancelErrorResponseModel)
		if string(jsonAlipayFundAuthOperationCancelErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundAuthOperationCancelErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundAuthOperationCancelErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundAuthOperationCancelErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundAuthOperationCancelDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundAuthOperationCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundAuthOperationCancelErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundAuthOperationCancelErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayFundAuthOperationCancelDefaultResponse struct {
	value *AlipayFundAuthOperationCancelDefaultResponse
	isSet bool
}

func (v NullableAlipayFundAuthOperationCancelDefaultResponse) Get() *AlipayFundAuthOperationCancelDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundAuthOperationCancelDefaultResponse) Set(val *AlipayFundAuthOperationCancelDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOperationCancelDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOperationCancelDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOperationCancelDefaultResponse(val *AlipayFundAuthOperationCancelDefaultResponse) *NullableAlipayFundAuthOperationCancelDefaultResponse {
	return &NullableAlipayFundAuthOperationCancelDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOperationCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOperationCancelDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


