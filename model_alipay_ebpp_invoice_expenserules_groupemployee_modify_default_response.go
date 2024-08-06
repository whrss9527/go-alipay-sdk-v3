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

// AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse struct for AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse
type AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse struct {
	AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel *AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel
	CommonErrorType                                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel)
	if err == nil {
		jsonAlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel)
		if string(jsonAlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceExpenserulesGroupemployeeModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse struct {
	value *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) Get() *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) Set(val *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse(val *AlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse {
	return &NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesGroupemployeeModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
