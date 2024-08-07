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

// AlipayDataBillEreceiptQueryDefaultResponse struct for AlipayDataBillEreceiptQueryDefaultResponse
type AlipayDataBillEreceiptQueryDefaultResponse struct {
	AlipayDataBillEreceiptQueryErrorResponseModel *AlipayDataBillEreceiptQueryErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayDataBillEreceiptQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayDataBillEreceiptQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayDataBillEreceiptQueryErrorResponseModel)
	if err == nil {
		jsonAlipayDataBillEreceiptQueryErrorResponseModel, _ := json.Marshal(dst.AlipayDataBillEreceiptQueryErrorResponseModel)
		if string(jsonAlipayDataBillEreceiptQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayDataBillEreceiptQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayDataBillEreceiptQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayDataBillEreceiptQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayDataBillEreceiptQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayDataBillEreceiptQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayDataBillEreceiptQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayDataBillEreceiptQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayDataBillEreceiptQueryDefaultResponse struct {
	value *AlipayDataBillEreceiptQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayDataBillEreceiptQueryDefaultResponse) Get() *AlipayDataBillEreceiptQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayDataBillEreceiptQueryDefaultResponse) Set(val *AlipayDataBillEreceiptQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillEreceiptQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillEreceiptQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillEreceiptQueryDefaultResponse(val *AlipayDataBillEreceiptQueryDefaultResponse) *NullableAlipayDataBillEreceiptQueryDefaultResponse {
	return &NullableAlipayDataBillEreceiptQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayDataBillEreceiptQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillEreceiptQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
