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


// AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse struct for AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse
type AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse struct {
	AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel)
		if string(jsonAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse struct {
	value *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) Get() *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) Set(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse(val *AlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse {
	return &NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseexctrlEmployertitleCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


