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


// AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse struct for AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse
type AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse struct {
	AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel *AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceInstitutionDetailinfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse struct {
	value *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) Get() *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) Set(val *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse(val *AlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) *NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInstitutionDetailinfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

