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


// AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse struct for AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse
type AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse struct {
	AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel *AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel);
	if err == nil {
		jsonAlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel, _ := json.Marshal(dst.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel)
		if string(jsonAlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayEbppInvoiceExpensesceneMerchantQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse struct {
	value *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) Get() *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) Set(val *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse(val *AlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) *NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse {
	return &NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensesceneMerchantQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


