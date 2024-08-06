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


// AlipayEbppInvoiceEinvpackageQueryDefaultResponse struct for AlipayEbppInvoiceEinvpackageQueryDefaultResponse
type AlipayEbppInvoiceEinvpackageQueryDefaultResponse struct {
	AlipayEbppInvoiceEinvpackageQueryErrorResponseModel *AlipayEbppInvoiceEinvpackageQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceEinvpackageQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceEinvpackageQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceEinvpackageQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceEinvpackageQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceEinvpackageQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceEinvpackageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceEinvpackageQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse struct {
	value *AlipayEbppInvoiceEinvpackageQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) Get() *AlipayEbppInvoiceEinvpackageQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) Set(val *AlipayEbppInvoiceEinvpackageQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse(val *AlipayEbppInvoiceEinvpackageQueryDefaultResponse) *NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEinvpackageQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


