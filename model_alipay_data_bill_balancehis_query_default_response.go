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


// AlipayDataBillBalancehisQueryDefaultResponse struct for AlipayDataBillBalancehisQueryDefaultResponse
type AlipayDataBillBalancehisQueryDefaultResponse struct {
	AlipayDataBillBalancehisQueryErrorResponseModel *AlipayDataBillBalancehisQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayDataBillBalancehisQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayDataBillBalancehisQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayDataBillBalancehisQueryErrorResponseModel);
	if err == nil {
		jsonAlipayDataBillBalancehisQueryErrorResponseModel, _ := json.Marshal(dst.AlipayDataBillBalancehisQueryErrorResponseModel)
		if string(jsonAlipayDataBillBalancehisQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayDataBillBalancehisQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayDataBillBalancehisQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayDataBillBalancehisQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayDataBillBalancehisQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayDataBillBalancehisQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayDataBillBalancehisQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayDataBillBalancehisQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayDataBillBalancehisQueryDefaultResponse struct {
	value *AlipayDataBillBalancehisQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayDataBillBalancehisQueryDefaultResponse) Get() *AlipayDataBillBalancehisQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayDataBillBalancehisQueryDefaultResponse) Set(val *AlipayDataBillBalancehisQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillBalancehisQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillBalancehisQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillBalancehisQueryDefaultResponse(val *AlipayDataBillBalancehisQueryDefaultResponse) *NullableAlipayDataBillBalancehisQueryDefaultResponse {
	return &NullableAlipayDataBillBalancehisQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayDataBillBalancehisQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillBalancehisQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


