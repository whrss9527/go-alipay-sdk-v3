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


// AlipayFundAuthOperationDetailQueryDefaultResponse struct for AlipayFundAuthOperationDetailQueryDefaultResponse
type AlipayFundAuthOperationDetailQueryDefaultResponse struct {
	AlipayFundAuthOperationDetailQueryErrorResponseModel *AlipayFundAuthOperationDetailQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundAuthOperationDetailQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundAuthOperationDetailQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundAuthOperationDetailQueryErrorResponseModel);
	if err == nil {
		jsonAlipayFundAuthOperationDetailQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundAuthOperationDetailQueryErrorResponseModel)
		if string(jsonAlipayFundAuthOperationDetailQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundAuthOperationDetailQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundAuthOperationDetailQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundAuthOperationDetailQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundAuthOperationDetailQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundAuthOperationDetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundAuthOperationDetailQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundAuthOperationDetailQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayFundAuthOperationDetailQueryDefaultResponse struct {
	value *AlipayFundAuthOperationDetailQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundAuthOperationDetailQueryDefaultResponse) Get() *AlipayFundAuthOperationDetailQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundAuthOperationDetailQueryDefaultResponse) Set(val *AlipayFundAuthOperationDetailQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAuthOperationDetailQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAuthOperationDetailQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAuthOperationDetailQueryDefaultResponse(val *AlipayFundAuthOperationDetailQueryDefaultResponse) *NullableAlipayFundAuthOperationDetailQueryDefaultResponse {
	return &NullableAlipayFundAuthOperationDetailQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundAuthOperationDetailQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAuthOperationDetailQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


