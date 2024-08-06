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


// AlipayEbppInvoiceDetailOutputQueryDefaultResponse struct for AlipayEbppInvoiceDetailOutputQueryDefaultResponse
type AlipayEbppInvoiceDetailOutputQueryDefaultResponse struct {
	AlipayEbppInvoiceDetailOutputQueryErrorResponseModel *AlipayEbppInvoiceDetailOutputQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceDetailOutputQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceDetailOutputQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceDetailOutputQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceDetailOutputQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceDetailOutputQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceDetailOutputQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceDetailOutputQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse struct {
	value *AlipayEbppInvoiceDetailOutputQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) Get() *AlipayEbppInvoiceDetailOutputQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) Set(val *AlipayEbppInvoiceDetailOutputQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse(val *AlipayEbppInvoiceDetailOutputQueryDefaultResponse) *NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceDetailOutputQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

