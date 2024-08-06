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


// AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse struct for AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse
type AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse struct {
	AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel *AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceExpenserulesEmployeerulesQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse struct {
	value *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) Get() *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) Set(val *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse(val *AlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesEmployeerulesQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

